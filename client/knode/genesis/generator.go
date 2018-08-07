package genesis

import (
	"fmt"
	"math/big"
	"math/rand"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core"
	"github.com/kowala-tech/kcoin/client/core/state"
	"github.com/kowala-tech/kcoin/client/core/vm"
	"github.com/kowala-tech/kcoin/client/core/vm/runtime"
	"github.com/kowala-tech/kcoin/client/kcoindb"
	"github.com/kowala-tech/kcoin/client/params"
)

const genesisTimestamp = 1528988194

type Generator interface {
	Generate(opts *Options) (*core.Genesis, error)
}

type generator struct {
	contracts    []*contract
	sharedState  *state.StateDB
	sharedTracer vm.Tracer
	alloc        core.GenesisAlloc
}

func NewGenerator() *generator {
	stateDB, err := state.New(common.Hash{}, state.NewDatabase(kcoindb.NewMemDatabase()))
	if err != nil {
		panic(err)
	}

	gen := &generator{
		sharedState:  stateDB,
		sharedTracer: newVmTracer(),
		alloc:        make(core.GenesisAlloc),
	}

	return gen
}

func (gen *generator) AddContract(contract *contract) {
	gen.contracts = append(gen.contracts, contract)
}

func Generate(opts Options) (*core.Genesis, error) {
	gen := NewGenerator()
	gen.AddContract(MultiSigContract)
	gen.AddContract(MiningTokenContract)
	gen.AddContract(ValidatorMgrContract)
	gen.AddContract(OracleMgrContract)

	return gen.Generate(opts)
}

func (gen *generator) Generate(opts Options) (*core.Genesis, error) {
	validOptions, err := validateOptions(opts)
	if err != nil {
		return nil, err
	}

	if err := gen.genesisAllocFromOptions(validOptions); err != nil {
		return nil, err
	}

	genesis := &core.Genesis{
		Timestamp: uint64(genesisTimestamp),
		GasLimit:  4700000,
		Alloc:     gen.alloc,
		Config: &params.ChainConfig{
			ChainID:   getNetwork(validOptions.network),
			Konsensus: getConsensusEngine(validOptions.consensusEngine),
		},
		ExtraData: getExtraData(opts.ExtraData),
	}

	fmt.Println("Please update the codebase with the following addresses (go bindings):")
	for _, contract := range gen.contracts {
		fmt.Printf("Contract: %s, Address: %s\n", contract.name, contract.address.Hex())
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
	for _, contract := range gen.contracts {
		contract.runtimeCfg = gen.getDefaultRuntimeConfig()
		if err := contract.deploy(contract, opts); err != nil {
			return err
		}
		if contract.postDeploy != nil {
			if err := contract.postDeploy(contract, opts); err != nil {
				return err
			}
		}
	}

	for _, contract := range gen.contracts {
		// @NOTE (rgeraldes) - storage needs to be addressed in the end as
		// contracts can interact with each other modifying each other's state
		contract.storage = contract.runtimeCfg.EVMConfig.Tracer.(*vmTracer).data[contract.address]
		gen.alloc[contract.address] = contract.AsGenesisAccount()
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

func getConsensusEngine(consensusEngine string) *params.KonsensusConfig {
	var consensus *params.KonsensusConfig

	switch consensusEngine {
	case KonsensusConsensus:
		consensus = &params.KonsensusConfig{}
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
