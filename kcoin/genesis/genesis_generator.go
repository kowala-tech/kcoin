package genesis

import (
	"bytes"
	"fmt"
	"math/big"
	"math/rand"
	"strings"
	"time"

	"github.com/kowala-tech/kcoin/accounts/abi"
	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/contracts/consensus"
	"github.com/kowala-tech/kcoin/contracts/nameservice"
	"github.com/kowala-tech/kcoin/contracts/oracle"
	"github.com/kowala-tech/kcoin/contracts/ownership"
	"github.com/kowala-tech/kcoin/core"
	"github.com/kowala-tech/kcoin/core/state"
	"github.com/kowala-tech/kcoin/core/vm"
	"github.com/kowala-tech/kcoin/core/vm/runtime"
	"github.com/kowala-tech/kcoin/kcoindb"
	"github.com/kowala-tech/kcoin/params"
	"github.com/pkg/errors"
)

const (
	MainNetwork  = "main"
	TestNetwork  = "test"
	OtherNetwork = "other"

	TendermintConsensus = "tendermint"
)

var (
	//DefaultSmartContractsOwner = network.MapChainIDToAddr[params.TestnetChainConfig.ChainID.Uint64()]

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
	ErrEmptyFreezePeriod                            = errors.New("freeze period in days is mandatory")
	ErrInvalidFreezePeriod                          = errors.New("freeze period is invalid")
	ErrEmptyWalletAddressValidator                  = errors.New("wallet address of genesis validator is mandatory")
	ErrInvalidWalletAddressValidator                = errors.New("wallet address of genesis validator is invalid")
	ErrEmptyPrefundedAccounts                       = errors.New("empty prefunded accounts, at least the validator wallet address should be included")
	ErrWalletAddressValidatorNotInPrefundedAccounts = errors.New("prefunded accounts should include genesis validator account")
	ErrInvalidAddressInPrefundedAccounts            = errors.New("address in prefunded accounts is invalid")
	ErrInvalidContractsOwnerAddress                 = errors.New("address used for smart contracts is invalid")
	ErrInvalidNetwork                               = errors.New("invalid Network, use main, test or other")
	ErrInvalidConsensusEngine                       = errors.New("invalid consensus engine")
)

type domain struct {
	name string
	addr common.Address
}

func GenerateGenesis(opts Options) (*core.Genesis, error) {
	validOptions, err := validateOptions(opts)
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
		// @TODO (rgeraldes) -
		ExtraData: getExtraData(opts.ExtraData),
	}

	// @NOTE (rgeraldes) - state needs to be shared between contracts because of the address
	// generation for multiple contracts with the same origin.
	db, err := kcoindb.NewMemDatabase()
	if err != nil {
		return nil, err
	}
	stateDB, err := state.New(common.Hash{}, state.NewDatabase(db))
	if err != nil {
		return nil, err
	} 

	multiSigAddr, err := addMultiSigWallet(stateDB, validOptions.multiSig, genesis)
	if err != nil {
		return nil, err
	}
	contractsOwner := multiSigAddr

	oracleMgrAddr, err := addOracleMgr(stateDB, validOptions.oracleMgr, genesis, contractsOwner)
	if err != nil {
		return nil, err
	}

	validatorMgrAddr, err := addValidatorMgr(stateDB, validOptions.validatorMgr, genesis, contractsOwner)
	if err != nil {
		return nil, err
	}

	domains := []*domain{&domain{params.ConsensusServiceDomain, *validatorMgrAddr}, &domain{params.OracleServiceDomain, *oracleMgrAddr}}
	_, err = addNameServiceWithDomains(stateDB, genesis, contractsOwner, domains)
	if err != nil {
		return nil, err
	}

	addPrefundedAccountsIntoGenesis(validOptions.prefundedAccounts, genesis)
	addBatchOfPrefundedAccountsIntoGenesis(genesis)

	return genesis, nil
}

func getDefaultRuntimeConfig(sharedState *state.StateDB) *runtime.Config {
	return &runtime.Config{
		State: sharedState,
		EVMConfig: vm.Config{
			Debug:  true,
			Tracer: newVmTracer(),
		},
	}
}

func addNameServiceWithDomains(stateDB *state.StateDB, genesis *core.Genesis, owner *common.Address, domains []*domain) (*common.Address, error) {
	runtimeCfg := getDefaultRuntimeConfig(stateDB)
	runtimeCfg.Origin = *owner

	nameServiceABI, err := abi.JSON(strings.NewReader(nameservice.NameServiceABI))
	if err != nil {
		return nil, err
	}

	// nameServiceParams params for the name service creation
	nameServiceParams, err := nameServiceABI.Pack("")
	if err != nil {
		return nil, err
	}

	// create contract
	contractCode, contractAddr, _, err := runtime.Create(append(common.FromHex(nameservice.NameServiceBin), nameServiceParams...), runtimeCfg)
	if err != nil {
		return nil, err
	}

	for _, domain := range domains {
		nameServiceRegParams, err := nameServiceABI.Pack(
			"register",
			domain.name,
			domain.addr,
		)
		if err != nil {
			return nil, err
		}
		// register domain
		runtime.Call(contractAddr, append(common.FromHex(nameservice.NameServiceBin), nameServiceRegParams...), runtimeCfg)
	}

	genesis.Alloc[contractAddr] = core.GenesisAccount{
		Storage: runtimeCfg.EVMConfig.Tracer.(*vmTracer).data[contractAddr],
		Code:    contractCode,
		Balance: new(big.Int),
	}

	return &contractAddr, nil
}

// addMultiSigWallet includes kowala's multi sig wallet in the genesis state. The creator - tx originator - has no influence
// in this specific contract (MultiSigWallet does not satisfy the Ownable contract) but we should use the same one all the
// time in order to have the same contract addresses.
func addMultiSigWallet(stateDB *state.StateDB, opts *validMultiSigOpts, genesis *core.Genesis) (*common.Address, error) {
	multiSigWalletABI, err := abi.JSON(strings.NewReader(ownership.MultiSigWalletABI))
	if err != nil {
		return nil, err
	}

	multiSigParams, err := multiSigWalletABI.Pack(
		"",
		opts.multiSigOwners,
		opts.numConfirmations,
	)
	if err != nil {
		return nil, err
	}

	runtimeCfg := getDefaultRuntimeConfig(stateDB)
	runtimeCfg.Origin = *opts.multiSigCreator
	contractCode, contractAddr, _, err := runtime.Create(append(common.FromHex(ownership.MultiSigWalletBin), multiSigParams...), runtimeCfg)
	if err != nil {
		return nil, err
	}

	genesis.Alloc[contractAddr] = core.GenesisAccount{
		Storage: runtimeCfg.EVMConfig.Tracer.(*vmTracer).data[contractAddr],
		Code:    contractCode,
		Balance: new(big.Int),
	}

	return &contractAddr, nil
}

// addValidatorManager includes the validator manager in the genesis block. The contract creator
// is also the owner - Oracle Manager satisfies the Ownable interface.
func addValidatorMgr(stateDB *state.StateDB, opts *validValidatorMgrOpts, genesis *core.Genesis, owner *common.Address) (*common.Address, error) {
	managerABI, err := abi.JSON(strings.NewReader(consensus.ValidatorManagerABI))
	if err != nil {
		return nil, err
	}

	managerParams, err := managerABI.Pack(
		"",
		opts.baseDeposit,
		opts.maxNumValidators,
		opts.freezePeriod,
		opts.validators[0],
	)
	if err != nil {
		return nil, err
	}

	runtimeCfg := getDefaultRuntimeConfig(stateDB)
	runtimeCfg.Origin = *owner
	contractCode, contractAddr, _, err := runtime.Create(append(common.FromHex(consensus.ValidatorManagerBin), managerParams...), runtimeCfg)
	if err != nil {
		return nil, err
	}

	genesis.Alloc[contractAddr] = core.GenesisAccount{
		Storage: runtimeCfg.EVMConfig.Tracer.(*vmTracer).data[contractAddr],
		Code:    contractCode,
		Balance: new(big.Int),
	}
	fmt.Printf("Do not forget to replace the hardcoded contract address @ contracts/consensus/consensus.go to %s!\n", contractAddr.Hex())

	return &contractAddr, nil
}

// addOracleMgr includes the oracle manager in the genesis block. The contract creator
// is also the owner - Oracle Manager satisfies the Ownable interface.
func addOracleMgr(stateDB *state.StateDB, opts *validOracleMgrOpts, genesis *core.Genesis, owner *common.Address) (*common.Address, error) {
	managerABI, err := abi.JSON(strings.NewReader(oracle.OracleManagerABI))
	if err != nil {
		return nil, err
	}

	managerParams, err := managerABI.Pack(
		"",
		opts.baseDeposit,
		opts.maxNumOracles,
		opts.freezePeriod,
	)
	if err != nil {
		return nil, err
	}

	runtimeCfg := getDefaultRuntimeConfig(stateDB)
	runtimeCfg.Origin = *owner
	// create contract
	contractCode, contractAddr, _, err := runtime.Create(append(common.FromHex(oracle.OracleManagerBin), managerParams...), runtimeCfg)
	if err != nil {
		return nil, err
	}

	genesis.Alloc[contractAddr] = core.GenesisAccount{
		Storage: runtimeCfg.EVMConfig.Tracer.(*vmTracer).data[contractAddr],
		Code:    contractCode,
		Balance: new(big.Int),
	}

	return &contractAddr, nil
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

func addBatchOfPrefundedAccountsIntoGenesis(genesis *core.Genesis) {
	// Add a batch of precompile balances to avoid them getting deleted
	for i := int64(0); i < 256; i++ {
		genesis.Alloc[common.BigToAddress(big.NewInt(i))] = core.GenesisAccount{Balance: big.NewInt(1)}
	}
}

func addPrefundedAccountsIntoGenesis(validPrefundedAccounts []*validPrefundedAccount, genesis *core.Genesis) {
	for _, vAccount := range validPrefundedAccounts {
		genesis.Alloc[*vAccount.accountAddress] = core.GenesisAccount{
			Balance: vAccount.balance,
		}
	}
}

func mapNetwork(network string) (string, error) {
	if !availableNetworks[network] {
		return "", fmt.Errorf("%v:%s", ErrInvalidNetwork, network)
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
		return nil, ErrEmptyFreezePeriod
	}

	unbondingPeriod, ok := new(big.Int).SetString(text, 0)
	if !ok {
		return nil, ErrInvalidFreezePeriod
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
		address, err := mapWalletAddress(a.AccountAddress)
		if err != nil {
			return nil, ErrInvalidAddressInPrefundedAccounts
		}

		balance, _ := new(big.Int).SetString(a.Balance, 0)

		validAccount := &validPrefundedAccount{
			accountAddress: address,
			balance:        balance,
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
		if bytes.Equal(account.accountAddress.Bytes(), addresses.Bytes()) {
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
