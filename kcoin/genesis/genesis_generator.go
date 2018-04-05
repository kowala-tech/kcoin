package genesis

import (
	"bytes"
	"github.com/kowala-tech/kcoin/accounts/abi"
	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/contracts/network"
	"github.com/kowala-tech/kcoin/contracts/network/contracts"
	"github.com/kowala-tech/kcoin/core"
	"github.com/kowala-tech/kcoin/core/vm"
	"github.com/kowala-tech/kcoin/core/vm/runtime"
	"github.com/kowala-tech/kcoin/params"
	"github.com/pkg/errors"
	"math/big"
	"math/rand"
	"strings"
	"time"
)

const (
	MainNetwork  = "main"
	TestNetwork  = "test"
	OtherNetwork = "other"

	TendermintConsensus = "tendermint"
)

var (
	DefaultSmartContractsOwner = network.MapChainIDToAddr[params.TestnetChainConfig.ChainID.Uint64()]

	availableNetworks = map[string]bool{
		MainNetwork:  true,
		TestNetwork:  true,
		OtherNetwork: true,
	}

	availableConsensusEngines = map[string]bool{
		TendermintConsensus: true,
	}

	ErrEmptyMaxNumValidators                        = errors.New("max number of validators is mandatory")
	ErrEmptyUnbondingPeriod                         = errors.New("unbonding period in days is mandatory")
	ErrEmptyWalletAddressValidator                  = errors.New("Wallet address of genesis validator is mandatory")
	ErrInvalidWalletAddressValidator                = errors.New("Wallet address of genesis validator is invalid")
	ErrEmptyPrefundedAccounts                       = errors.New("empty prefunded accounts, at least the validator wallet address should be included")
	ErrWalletAddressValidatorNotInPrefundedAccounts = errors.New("prefunded accounts should include genesis validator account")
	ErrInvalidAddressInPrefundedAccounts            = errors.New("address in prefunded accounts is invalid")
	ErrInvalidContractsOwnerAddress                 = errors.New("address used for smart contracts is invalid")
	ErrInvalidNetwork                               = errors.New("invalid Network, use main, test or other")
	ErrInvalidConsensusEngine                       = errors.New("invalid consensus engine")
)

type GenesisOptions struct {
	Network                       string
	MaxNumValidators              string
	UnbondingPeriod               string
	WalletAddressGenesisValidator string
	PrefundedAccounts             []PrefundedAccount
	ConsensusEngine               string
	SmartContractsOwner           string
	ExtraData                     string
}

type PrefundedAccount struct {
	WalletAddress string
	Balance       int64
}

type validPrefundedAccount struct {
	walletAddress *common.Address
	balance       *big.Int
}

type validGenesisOptions struct {
	network                       string
	maxNumValidators              *big.Int
	unbondingPeriod               *big.Int
	walletAddressGenesisValidator *common.Address
	prefundedAccounts             []*validPrefundedAccount
	consensusEngine               string
	smartContractsOwner           *common.Address
}

func GenerateGenesis(options GenesisOptions) (*core.Genesis, error) {
	validOptions, err := validateOptions(options)
	if err != nil {
		return nil, err
	}

	genesis := &core.Genesis{
		Timestamp: uint64(time.Now().Unix()),
		GasLimit:  4700000,
		Alloc:     make(core.GenesisAlloc),
		Config:    &params.ChainConfig{},
	}

	setNetwork(validOptions.network, genesis)
	setConsensusEngine(validOptions.consensusEngine, genesis)
	setExtraData(options.ExtraData, genesis)

	genesis.Alloc[*validOptions.smartContractsOwner] = core.GenesisAccount{Balance: new(big.Int).Mul(common.Big1, big.NewInt(params.Ether))}

	//TODO: This maybe will be need to be available to change by the parameters in the options in the future, right now is 0.
	baseDeposit := common.Big0

	electionABI, err := abi.JSON(strings.NewReader(contracts.ElectionContractABI))
	if err != nil {
		return nil, err
	}

	electionParams, err := electionABI.Pack(
		"",
		baseDeposit,
		validOptions.maxNumValidators,
		validOptions.unbondingPeriod,
		*validOptions.walletAddressGenesisValidator,
	)
	if err != nil {
		return nil, err
	}

	runtimeCfg := &runtime.Config{
		Origin: *validOptions.smartContractsOwner,
		EVMConfig: vm.Config{
			Debug:  true,
			Tracer: newVmTracer(),
		},
	}

	contract, err := createContract(runtimeCfg, append(common.FromHex(contracts.ElectionContractBin), electionParams...))
	if err != nil {
		return nil, err
	}

	genesis.Alloc[contract.addr] = core.GenesisAccount{
		Code:    contract.code,
		Storage: contract.storage,
		Balance: new(big.Int).Mul(baseDeposit, new(big.Int).SetUint64(params.Ether)),
	}

	setPrefundedAccounts(validOptions.prefundedAccounts, genesis)
	addBatchOfPrefundedAccounts(genesis)

	return genesis, nil
}

func setExtraData(extraData string, genesis *core.Genesis) {
	extra := ""
	if extraData != "" {
		extra = extraData
	}
	genesis.ExtraData = make([]byte, 32)
	if len(extra) > 32 {
		extra = extra[:32]
	}
	genesis.ExtraData = append([]byte(extra), genesis.ExtraData[len(extra):]...)
}

func setConsensusEngine(consensusEngine string, genesis *core.Genesis) {
	switch consensusEngine {
	case TendermintConsensus:
		genesis.Config.Tendermint = &params.TendermintConfig{Rewarded: true}
	}
}

func setNetwork(network string, genesis *core.Genesis) {
	switch network {
	case MainNetwork:
		genesis.Config.ChainID = params.MainnetChainConfig.ChainID
	case TestNetwork:
		genesis.Config.ChainID = params.TestnetChainConfig.ChainID
	case OtherNetwork:
		genesis.Config.ChainID = new(big.Int).SetUint64(uint64(rand.Intn(65536)))
	}
}

func validateOptions(options GenesisOptions) (*validGenesisOptions, error) {
	network, err := createNetwork(options.Network)
	if err != nil {
		return nil, err
	}

	maxNumValidators, err := createMaxNumValidators(options.MaxNumValidators)
	if err != nil {
		return nil, err
	}

	unbondingPeriod, err := createUnbondingPeriod(options.UnbondingPeriod)
	if err != nil {
		return nil, err
	}

	walletAddressValidator, err := createWalletAddress(options.WalletAddressGenesisValidator)
	if err != nil {
		return nil, err
	}

	validPrefundedAccounts, err := createPrefundedAccounts(options.PrefundedAccounts)
	if err != nil {
		return nil, err
	}

	if !prefundedIncludesValidatorWallet(validPrefundedAccounts, walletAddressValidator) {
		return nil, ErrWalletAddressValidatorNotInPrefundedAccounts
	}

	consensusEngine := TendermintConsensus
	if options.ConsensusEngine != "" {
		consensusEngine, err = createConsensusEngine(options.ConsensusEngine)
		if err != nil {
			return nil, err
		}
	}

	owner := &DefaultSmartContractsOwner
	if options.SmartContractsOwner != "" {
		strAddr := options.SmartContractsOwner

		owner, err = createWalletAddress(strAddr)
		if err != nil {
			return nil, ErrInvalidContractsOwnerAddress
		}
	}

	return &validGenesisOptions{
		network:                       network,
		maxNumValidators:              maxNumValidators,
		unbondingPeriod:               unbondingPeriod,
		walletAddressGenesisValidator: walletAddressValidator,
		prefundedAccounts:             validPrefundedAccounts,
		consensusEngine:               consensusEngine,
		smartContractsOwner:           owner,
	}, nil
}

func addBatchOfPrefundedAccounts(genesis *core.Genesis) {
	// Add a batch of precompile balances to avoid them getting deleted
	for i := int64(0); i < 256; i++ {
		genesis.Alloc[common.BigToAddress(big.NewInt(i))] = core.GenesisAccount{Balance: big.NewInt(1)}
	}
}

func setPrefundedAccounts(validPrefundedAccounts []*validPrefundedAccount, genesis *core.Genesis) {
	for _, vAccount := range validPrefundedAccounts {
		genesis.Alloc[*vAccount.walletAddress] = core.GenesisAccount{
			Balance: vAccount.balance,
		}
	}
}

func createNetwork(network string) (string, error) {
	if !availableNetworks[network] {
		return "", ErrInvalidNetwork
	}

	return network, nil
}

func createConsensusEngine(consensus string) (string, error) {
	if !availableConsensusEngines[consensus] {
		return "", ErrInvalidConsensusEngine
	}

	return consensus, nil
}

func createMaxNumValidators(s string) (*big.Int, error) {
	if s = strings.TrimSpace(s); s == "" {
		return nil, ErrEmptyMaxNumValidators
	}

	numValidators, ok := new(big.Int).SetString(s, 0)
	if !ok {
		//TODO: Create error
		return nil, errors.New("invalid max num of validators.")
	}

	return numValidators, nil
}

func createUnbondingPeriod(uP string) (*big.Int, error) {
	var text string
	if text = strings.TrimSpace(uP); text == "" {
		return nil, ErrEmptyUnbondingPeriod
	}

	unbondingPeriod, ok := new(big.Int).SetString(text, 0)
	if !ok {
		//TODO: Create error
		return nil, errors.New("invalid max num of validators.")
	}

	return unbondingPeriod, nil
}

func createWalletAddress(wA string) (*common.Address, error) {
	stringAddr := wA

	if text := strings.TrimSpace(wA); text == "" {
		return nil, ErrEmptyWalletAddressValidator
	}

	if strings.HasPrefix(stringAddr, "0x") {
		stringAddr = strings.TrimPrefix(stringAddr, "0x")
	}

	if len(stringAddr) != 40 {
		return nil, ErrInvalidWalletAddressValidator
	}

	bigaddr, _ := new(big.Int).SetString(stringAddr, 16)
	address := common.BigToAddress(bigaddr)

	return &address, nil
}

func createPrefundedAccounts(accounts []PrefundedAccount) ([]*validPrefundedAccount, error) {
	var validAccounts []*validPrefundedAccount

	if len(accounts) == 0 {
		return nil, ErrEmptyPrefundedAccounts
	}

	for _, a := range accounts {
		address, err := createWalletAddress(a.WalletAddress)
		if err != nil {
			return nil, ErrInvalidAddressInPrefundedAccounts
		}

		balance := big.NewInt(a.Balance)

		validAccount := &validPrefundedAccount{
			walletAddress: address,
			balance:       balance,
		}

		validAccounts = append(validAccounts, validAccount)
	}

	return validAccounts, nil
}

func prefundedIncludesValidatorWallet(
	accounts []*validPrefundedAccount,
	addresses *common.Address,
) bool {
	for _, account := range accounts {
		if bytes.Equal(account.walletAddress.Bytes(), addresses.Bytes()) {
			return true
		}
	}

	return false
}

type contractData struct {
	addr    common.Address
	code    []byte
	storage map[common.Hash]common.Hash
}

func createContract(cfg *runtime.Config, code []byte) (*contractData, error) {
	out, addr, _, err := runtime.Create(code, cfg)
	if err != nil {
		return nil, err
	}
	return &contractData{
		addr:    addr,
		code:    out,
		storage: cfg.EVMConfig.Tracer.(*vmTracer).data[addr],
	}, nil
}

//TODO:This vmTracer is an exact copy of the one in pupphet, can be unified? Or maybe it is not even needed? Someone
//with more knowledge on this part is welcome.
type vmTracer struct {
	data map[common.Address]map[common.Hash]common.Hash
}

func newVmTracer() *vmTracer {
	return &vmTracer{
		data: make(map[common.Address]map[common.Hash]common.Hash, 1024),
	}
}

func (vmt *vmTracer) CaptureState(env *vm.EVM, pc uint64, op vm.OpCode, gas, cost uint64, memory *vm.Memory, stack *vm.Stack, contract *vm.Contract, depth int, err error) error {
	if err != nil {
		return err
	}
	if op == vm.SSTORE {
		s := stack.Data()
		addrStorage, ok := vmt.data[contract.Address()]
		if !ok {
			addrStorage = make(map[common.Hash]common.Hash, 1024)
			vmt.data[contract.Address()] = addrStorage
		}
		addrStorage[common.BigToHash(s[len(s)-1])] = common.BigToHash(s[len(s)-2])
	}
	return nil
}

func (vmt *vmTracer) CaptureEnd(output []byte, gasUsed uint64, t time.Duration, err error) error {
	return nil
}
