package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"strings"
	"time"

	"github.com/kowala-tech/kUSD/accounts/abi"
	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/contracts/network/contracts"
	"github.com/kowala-tech/kUSD/core"
	"github.com/kowala-tech/kUSD/core/vm"
	"github.com/kowala-tech/kUSD/core/vm/runtime"
	"github.com/kowala-tech/kUSD/log"
	"github.com/kowala-tech/kUSD/params"
)

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

func (w *wizard) createElectionContract(genesis *core.Genesis, owner common.Address) error {
	fmt.Println()
	var baseDeposit *big.Int
	for baseDeposit == nil {
		fmt.Println("How much is a deposit to secure a place in the election? (default=0)")
		fmt.Println("Note: Multiplier for the kUSD denomination: 1e18 (ex: 10 kUSD = 10 * 1e18)")
		baseDeposit = w.readDefaultBigInt(common.Big0)
	}

	fmt.Println()
	var genesisAddr *common.Address
	for genesisAddr == nil {
		fmt.Println("Which account will be used as the genesis validator? (mandatory at least one)")
		genesisAddr = w.readAddress()
	}

	fmt.Println()
	var maxVal *big.Int
	for maxVal == nil {
		fmt.Println("What is the maximum number of validators? (default=1)")
		maxVal = w.readDefaultBigInt(common.Big1)
	}

	fmt.Println()
	var unbondingPeriod *big.Int
	for unbondingPeriod == nil {
		fmt.Println("How long should the unbonding period be? (default=0)")
		unbondingPeriod = w.readDefaultBigInt(common.Big0)
	}

	electionABI, err := abi.JSON(strings.NewReader(contracts.ElectionContractABI))
	if err != nil {
		log.Error("can't parse the election contract ABI:", err)
		return err
	}

	electionParams, err := electionABI.Pack("", baseDeposit, maxVal, unbondingPeriod, *genesisAddr)
	if err != nil {
		log.Error("can't pack the election contract params", err)
		return err
	}

	runtimeCfg := &runtime.Config{
		Origin: owner,
		EVMConfig: vm.Config{
			Debug:  true,
			Tracer: newVmTracer(),
		},
	}

	contract, err := createContract(runtimeCfg, append(common.FromHex(contracts.ElectionContractBin), electionParams...))
	if err != nil {
		log.Error("can't create the election contract", err)
		return err
	}

	genesis.Alloc[contract.addr] = core.GenesisAccount{
		Code:    contract.code,
		Storage: contract.storage,
		Balance: baseDeposit,
	}

	return nil
}

// makeGenesis creates a new genesis struct based on some user input.
func (w *wizard) makeGenesis() {
	// Construct a default genesis block
	genesis := &core.Genesis{
		Timestamp: uint64(time.Now().Unix()),
		GasLimit:  4700000,
		Alloc:     make(core.GenesisAlloc),
		Config:    &params.ChainConfig{},
	}

	fmt.Println()
	fmt.Println("Which network to use? (default = Test Network)")
	fmt.Println(" 1. Main Network")
	fmt.Println(" 2. Test Network")
	fmt.Println(" 3. Other Network")
	choice := w.read()
	switch {
	case choice == "1":
		genesis.Config.ChainID = params.MainnetChainConfig.ChainID
	case choice == "2" || choice == "":
		genesis.Config.ChainID = params.TestnetChainConfig.ChainID
	case choice == "3":
		fmt.Println("Specify your chain/network ID if you want an explicit one (default = random)")
		genesis.Config.ChainID = new(big.Int).SetUint64(uint64(w.readDefaultInt(rand.Intn(65536))))
	default:
		log.Crit("Invalid network choice", "choice", choice)
	}

	// Figure out which consensus engine to choose
	fmt.Println()
	fmt.Println("Which consensus engine to use? (default = Tendermint)")
	fmt.Println(" 1. Tendermint - proof-of-stake")

	choice := w.read()
	var owner *common.Address
	switch {
	case choice == "" || choice == "1":
		genesis.Config.Tendermint = &params.TendermintConfig{Rewarded: true}
		genesis.ExtraData = make([]byte, 32)

		fmt.Println()
		for owner == nil {
			fmt.Println("Which account will be used as the owner of the network contracts? (mandatory at least one)")
			owner = w.readAddress()
		}

		log.Info("the owner account will be pre-funded with 1 coin to cover the gas used", "address", owner)
		genesis.Alloc[*owner] = core.GenesisAccount{Balance: new(big.Int).Mul(common.Big1, big.NewInt(params.Ether))}

		if err := w.createElectionContract(genesis, *owner); err != nil {
			log.Crit("Failed to create", err)
		}

	default:
		log.Crit("Invalid consensus engine choice", "choice", choice)
	}

	// Consensus all set, just ask for initial funds and go
	fmt.Println()
	fmt.Println("Which accounts should be pre-funded? (advisable at least one)")
	for {
		// Read the address of the account to fund
		if address := w.readAddress(); address != nil {
			genesis.Alloc[*address] = core.GenesisAccount{
				Balance: new(big.Int).Lsh(big.NewInt(1), 256-7), // 2^256 / 128 (allow many pre-funds without balance overflows)
			}
			continue
		}
		break
	}

	// Add a batch of precompile balances to avoid them getting deleted
	for i := int64(0); i < 256; i++ {
		genesis.Alloc[common.BigToAddress(big.NewInt(i))] = core.GenesisAccount{Balance: big.NewInt(1)}
	}
	fmt.Println()

	// Query the user for some custom extras
	fmt.Println()
	fmt.Println("Anything fun to embed into the genesis block? (max 32 bytes)")

	extra := w.read()
	if len(extra) > 32 {
		extra = extra[:32]
	}
	genesis.ExtraData = append([]byte(extra), genesis.ExtraData[len(extra):]...)

	// All done, store the genesis and flush to disk
	w.conf.genesis = genesis
}
