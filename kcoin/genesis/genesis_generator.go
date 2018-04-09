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
	ErrInvalidMaxNumValidators                      = errors.New("invalid max num of validators")
	ErrEmptyUnbondingPeriod                         = errors.New("unbonding period in days is mandatory")
	ErrInvalidUnbondingPeriod                       = errors.New("unbonding period is invalid")
	ErrEmptyWalletAddressValidator                  = errors.New("wallet address of genesis validator is mandatory")
	ErrInvalidWalletAddressValidator                = errors.New("wallet address of genesis validator is invalid")
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
	Balance       string
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
		Config: &params.ChainConfig{
			ChainID:    getNetwork(validOptions.network),
			Tendermint: getConsensusEngine(validOptions.consensusEngine),
		},
		ExtraData: getExtraData(options.ExtraData),
	}

	genesis.Alloc[*validOptions.smartContractsOwner] = core.GenesisAccount{Balance: new(big.Int).Mul(common.Big1, big.NewInt(params.Ether))}

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

	addPrefundedAccountsIntoGenesis(validOptions.prefundedAccounts, genesis)
	addBatchOfPrefundedAccountsIntoGenesis(genesis)

	return genesis, nil
}

func getExtraData(extraData string) []byte {
	extra := ""
	if extraData != "" {
		extra = extraData
	}
	extraSlice := make([]byte, 32)
	if len(extra) > 32 {
		extra = extra[:32]
	}
	return append([]byte(extra), extraSlice[len(extra):]...)
}

func getConsensusEngine(consensusEngine string) *params.TendermintConfig {
	var consensus *params.TendermintConfig

	switch consensusEngine {
	case TendermintConsensus:
		consensus = &params.TendermintConfig{Rewarded: true}
	}

	return consensus
}

func getNetwork(network string) *big.Int {
	var chainId *big.Int

	switch network {
	case MainNetwork:
		chainId = params.MainnetChainConfig.ChainID
	case TestNetwork:
		chainId = params.TestnetChainConfig.ChainID
	case OtherNetwork:
		chainId = new(big.Int).SetUint64(uint64(rand.Intn(65536)))
	}

	return chainId
}

func validateOptions(options GenesisOptions) (*validGenesisOptions, error) {
	network, err := mapNetwork(options.Network)
	if err != nil {
		return nil, err
	}

	maxNumValidators, err := mapMaxNumValidators(options.MaxNumValidators)
	if err != nil {
		return nil, err
	}

	unbondingPeriod, err := mapUnbondingPeriod(options.UnbondingPeriod)
	if err != nil {
		return nil, err
	}

	walletAddressValidator, err := mapWalletAddress(options.WalletAddressGenesisValidator)
	if err != nil {
		return nil, err
	}

	validPrefundedAccounts, err := mapPrefundedAccounts(options.PrefundedAccounts)
	if err != nil {
		return nil, err
	}

	if !prefundedIncludesValidatorWallet(validPrefundedAccounts, walletAddressValidator) {
		return nil, ErrWalletAddressValidatorNotInPrefundedAccounts
	}

	consensusEngine := TendermintConsensus
	if options.ConsensusEngine != "" {
		consensusEngine, err = mapConsensusEngine(options.ConsensusEngine)
		if err != nil {
			return nil, err
		}
	}

	owner := &DefaultSmartContractsOwner
	if options.SmartContractsOwner != "" {
		strAddr := options.SmartContractsOwner

		owner, err = mapWalletAddress(strAddr)
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

func addBatchOfPrefundedAccountsIntoGenesis(genesis *core.Genesis) {
	// Add a batch of precompile balances to avoid them getting deleted
	for i := int64(0); i < 256; i++ {
		genesis.Alloc[common.BigToAddress(big.NewInt(i))] = core.GenesisAccount{Balance: big.NewInt(1)}
	}
}

func addPrefundedAccountsIntoGenesis(validPrefundedAccounts []*validPrefundedAccount, genesis *core.Genesis) {
	for _, vAccount := range validPrefundedAccounts {
		genesis.Alloc[*vAccount.walletAddress] = core.GenesisAccount{
			Balance: vAccount.balance,
		}
	}
}

func mapNetwork(network string) (string, error) {
	if !availableNetworks[network] {
		return "", ErrInvalidNetwork
	}

	return network, nil
}

func mapConsensusEngine(consensus string) (string, error) {
	if !availableConsensusEngines[consensus] {
		return "", ErrInvalidConsensusEngine
	}

	return consensus, nil
}

func mapMaxNumValidators(s string) (*big.Int, error) {
	if s = strings.TrimSpace(s); s == "" {
		return nil, ErrEmptyMaxNumValidators
	}

	numValidators, ok := new(big.Int).SetString(s, 0)
	if !ok {
		return nil, ErrInvalidMaxNumValidators
	}

	return numValidators, nil
}

func mapUnbondingPeriod(uP string) (*big.Int, error) {
	var text string
	if text = strings.TrimSpace(uP); text == "" {
		return nil, ErrEmptyUnbondingPeriod
	}

	unbondingPeriod, ok := new(big.Int).SetString(text, 0)
	if !ok {
		return nil, ErrInvalidUnbondingPeriod
	}

	return unbondingPeriod, nil
}

func mapWalletAddress(a string) (*common.Address, error) {
	stringAddr := a

	if text := strings.TrimSpace(a); text == "" {
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

func mapPrefundedAccounts(accounts []PrefundedAccount) ([]*validPrefundedAccount, error) {
	var validAccounts []*validPrefundedAccount

	if len(accounts) == 0 {
		return nil, ErrEmptyPrefundedAccounts
	}

	for _, a := range accounts {
		address, err := mapWalletAddress(a.WalletAddress)
		if err != nil {
			return nil, ErrInvalidAddressInPrefundedAccounts
		}

		balance, _ := new(big.Int).SetString(a.Balance, 0)

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
