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
	"github.com/kowala-tech/kcoin/contracts/oracle"
	"github.com/kowala-tech/kcoin/contracts/ownership"
	"github.com/kowala-tech/kcoin/contracts/token"
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

	tokenCustomFallback = "registerValidator(address,uint256,bytes)"
)

var (
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

type Generator interface {
	Generate(opts *Options) (*core.Genesis, error)
}

type generator struct {
	sharedState  *state.StateDB
	sharedTracer vm.Tracer
	alloc        core.GenesisAlloc
}

func New(opts *Options) (*core.Genesis, error) {
	db, err := kcoindb.NewMemDatabase()
	if err != nil {
		return nil, err
	}
	stateDB, err := state.New(common.Hash{}, state.NewDatabase(db))
	if err != nil {
		return nil, err
	}

	gen := &generator{
		sharedState:  stateDB,
		sharedTracer: newVmTracer(),
		alloc:        make(core.GenesisAlloc),
	}

	return gen.Generate(opts)
}

func (gen *generator) Generate(opts *Options) (*core.Genesis, error) {
	validOptions, err := validateOptions(opts)
	if err != nil {
		return nil, err
	}

	if err := gen.genesisAllocFromOptions(validOptions); err != nil {
		return nil, err
	}

	genesis := &core.Genesis{
		Timestamp: uint64(time.Now().Unix()),
		GasLimit:  4700000,
		Alloc:     gen.alloc,
		Config: &params.ChainConfig{
			ChainID:    getNetwork(validOptions.network),
			Tendermint: getConsensusEngine(validOptions.consensusEngine),
		},
		// @TODO (rgeraldes)
		ExtraData: getExtraData(opts.ExtraData),
	}

	return genesis, nil
}

func (gen *generator) genesisAllocFromOptions(opts *validGenesisOptions) error {
	if err := gen.deployContracts(opts); err != nil {
		return err
	}

	gen.prefundAccounts(opts.prefundedAccounts)
	gen.addBatchOfPrefundedAccountsIntoGenesis()

	return nil
}

func (gen *generator) deployContracts(opts *validGenesisOptions) error {
	ownerAddr, err := gen.deployMultiSigWallet(opts.multiSig)
	if err != nil {
		return err
	}

	opts.miningToken.owner = *ownerAddr
	miningTokenAddr, err := gen.deployMiningToken(opts.miningToken)
	if err != nil {
		return err
	}

	// @NOTE (rgeraldes) - validator manager must know the mining token addr in order to transfer back the funds
	opts.validatorMgr.miningTokenAddr = *miningTokenAddr
	_, err = gen.deployValidatorMgr(opts.validatorMgr, ownerAddr)
	if err != nil {
		return err
	}

	_, err = gen.deployOracleMgr(opts.oracleMgr, ownerAddr)
	if err != nil {
		return err
	}

	return nil
}

func (gen *generator) getDefaultRuntimeConfig() *runtime.Config {
	return &runtime.Config{
		State:       gen.sharedState,
		BlockNumber: common.Big0,
		EVMConfig: vm.Config{
			Debug:  true,
			Tracer: gen.sharedTracer,
		},
	}
}

// deployMultiSigWallet includes kowala's multi sig wallet in the genesis state. The creator - tx originator - has no influence
// in this specific contract (MultiSigWallet does not satisfy the Ownable contract) but we should use the same one all the
// time in order to have the same contract addresses.
func (gen *generator) deployMultiSigWallet(opts *validMultiSigOpts) (*common.Address, error) {
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

	runtimeCfg := gen.getDefaultRuntimeConfig()
	runtimeCfg.Origin = *opts.multiSigCreator
	contractCode, contractAddr, _, err := runtime.Create(append(common.FromHex(ownership.MultiSigWalletBin), multiSigParams...), runtimeCfg)
	if err != nil {
		return nil, err
	}

	gen.alloc[contractAddr] = core.GenesisAccount{
		Storage: runtimeCfg.EVMConfig.Tracer.(*vmTracer).data[contractAddr],
		Code:    contractCode,
		Balance: new(big.Int),
	}

	return &contractAddr, nil
}

// deployValidatorMgr includes the validator manager in the genesis block. The contract creator
// is also the owner - Oracle Manager satisfies the Ownable interface.
func (gen *generator) deployValidatorMgr(opts *validValidatorMgrOpts, owner *common.Address) (*common.Address, error) {
	managerABI, err := abi.JSON(strings.NewReader(consensus.ValidatorMgrABI))
	if err != nil {
		return nil, err
	}

	managerParams, err := managerABI.Pack(
		"",
		opts.baseDeposit,
		opts.maxNumValidators,
		opts.freezePeriod,
		opts.miningTokenAddr,
	)
	if err != nil {
		return nil, err
	}

	runtimeCfg := gen.getDefaultRuntimeConfig()
	runtimeCfg.Origin = *owner
	contractCode, contractAddr, _, err := runtime.Create(append(common.FromHex(consensus.ValidatorMgrBin), managerParams...), runtimeCfg)
	if err != nil {
		return nil, err
	}

	// register genesis validators
	tokenABI, err := abi.JSON(strings.NewReader(token.MiningTokenABI))
	if err != nil {
		return nil, err
	}

	for _, validator := range opts.validators {
		runtimeCfg.Origin = validator.address

		registrationParams, err := tokenABI.Pack(
			"transfer",
			contractAddr,
			validator.deposit,
			[]byte("not_zero"), // @NOTE (rgeraldes) - https://github.com/kowala-tech/kcoin/issues/285
			tokenCustomFallback,
		)
		if err != nil {
			return nil, err
		}

		_, _, err = runtime.Call(opts.miningTokenAddr, registrationParams, runtimeCfg)
		if err != nil {
			return nil, fmt.Errorf("%s:%s", "Failed to register validator", err)
		}
	}

	gen.alloc[contractAddr] = core.GenesisAccount{
		Storage: runtimeCfg.EVMConfig.Tracer.(*vmTracer).data[contractAddr],
		Code:    contractCode,
		Balance: new(big.Int),
	}

	// NOTE(rgeraldes) - mtoken contract has been modified and it's storage needs to be updated
	// @TODO (rgeraldes) - trace the state just in the end of the main method
	gen.alloc[opts.miningTokenAddr] = core.GenesisAccount{
		Storage: runtimeCfg.EVMConfig.Tracer.(*vmTracer).data[opts.miningTokenAddr],
		Code:    gen.alloc[opts.miningTokenAddr].Code,
		Balance: gen.alloc[opts.miningTokenAddr].Balance,
	}

	return &contractAddr, nil
}

// deployMiningToken includes the mUSD token contract in the genesis block
func (gen *generator) deployMiningToken(opts *validMiningTokenOpts) (*common.Address, error) {
	tokenABI, err := abi.JSON(strings.NewReader(token.MiningTokenABI))
	if err != nil {
		return nil, err
	}

	tokenParams, err := tokenABI.Pack(
		"",
		opts.name,
		opts.symbol,
		opts.cap,
		// @TODO (rgeraldes) - modify type
		uint8(opts.decimals.Uint64()),
	)
	if err != nil {
		return nil, err
	}

	runtimeCfg := gen.getDefaultRuntimeConfig()
	runtimeCfg.Origin = opts.owner
	contractCode, contractAddr, _, err := runtime.Create(append(common.FromHex(token.MiningTokenBin), tokenParams...), runtimeCfg)
	if err != nil {
		return nil, err
	}

	for _, holder := range opts.holders {
		mintParams, err := tokenABI.Pack(
			"mint",
			holder.address,
			holder.balance,
		)
		if err != nil {
			return nil, err
		}
		// mint tokens
		_, _, err = runtime.Call(contractAddr, mintParams, runtimeCfg)
		if err != nil {
			return nil, fmt.Errorf("%s:%s", "Failed to mint tokens", err)
		}
	}

	gen.alloc[contractAddr] = core.GenesisAccount{
		Storage: runtimeCfg.EVMConfig.Tracer.(*vmTracer).data[contractAddr],
		Code:    contractCode,
		Balance: new(big.Int),
	}

	return &contractAddr, nil
}

// deployOracleMgr includes the oracle manager in the genesis block. The contract creator
// is also the owner - Oracle Manager satisfies the Ownable interface.
func (gen *generator) deployOracleMgr(opts *validOracleMgrOpts, owner *common.Address) (*common.Address, error) {
	managerABI, err := abi.JSON(strings.NewReader(oracle.OracleMgrABI))
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

	runtimeCfg := gen.getDefaultRuntimeConfig()
	runtimeCfg.Origin = *owner
	contractCode, contractAddr, _, err := runtime.Create(append(common.FromHex(oracle.OracleMgrBin), managerParams...), runtimeCfg)
	if err != nil {
		return nil, err
	}

	gen.alloc[contractAddr] = core.GenesisAccount{
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

func (gen *generator) addBatchOfPrefundedAccountsIntoGenesis() {
	// Add a batch of precompile balances to avoid them getting deleted
	for i := int64(0); i < 256; i++ {
		gen.alloc[common.BigToAddress(big.NewInt(i))] = core.GenesisAccount{Balance: big.NewInt(1)}
	}
}

func (gen *generator) prefundAccounts(validPrefundedAccounts []*validPrefundedAccount) {
	for _, vAccount := range validPrefundedAccounts {
		gen.alloc[*vAccount.accountAddress] = core.GenesisAccount{
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
		address, err := mapWalletAddress(a.Address)
		if err != nil {
			return nil, ErrInvalidAddressInPrefundedAccounts
		}

		balance := new(big.Int).Mul(new(big.Int).SetUint64(a.Balance), new(big.Int).SetUint64(params.Ether))

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
