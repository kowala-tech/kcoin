package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/kowala-tech/kUSD/accounts/abi"
	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/contracts/network"
	"github.com/kowala-tech/kUSD/core"
	"github.com/kowala-tech/kUSD/core/state"
	"github.com/kowala-tech/kUSD/core/vm"
	"github.com/kowala-tech/kUSD/core/vm/runtime"
	"github.com/kowala-tech/kUSD/kusddb"
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

type contractType byte

const (
	ctNetworkMap contractType = iota
	ctMToken
	ctPriceOracle
	ctNetworkStats
)

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

func createContracts(owner common.Address, contractsCode map[contractType][]byte) ([]*contractData, error) {
	// memdb
	memDb, err := kusddb.NewMemDatabase()
	if err != nil {
		return nil, err
	}
	// statedb
	sdb, err := state.New(common.Hash{}, state.NewDatabase(memDb))
	if err != nil {
		return nil, err
	}
	// tracer
	tracer := newVmTracer()
	// evm runtime config
	runtimeConfig := &runtime.Config{
		Origin: owner,
		State:  sdb,
		EVMConfig: vm.Config{
			Debug:  true,
			Tracer: tracer,
		},
	}
	// run contracts
	r := make([]*contractData, 0, len(contractsCode))
	// first the network stats contract
	c, err := createContract(runtimeConfig, contractsCode[ctNetworkStats])
	if err != nil {
		fmt.Println("can't create network stats contract:", err)
		os.Exit(-5)
	}
	r = append(r, c)
	// create mToken contract
	if c, err = createContract(runtimeConfig, contractsCode[ctMToken]); err != nil {
		fmt.Println("can't create mToken contract:", err)
		os.Exit(-6)
	}
	r = append(r, c)
	// create price oracle contract
	priceOracleAbi, err := abi.JSON(strings.NewReader(network.PriceOracleContractABI))
	if err != nil {
		fmt.Println("can't parse price oracle contract ABI:", err)
		os.Exit(-7)
	}
	priceOracleParams, err := priceOracleAbi.Pack("",
		"kUSD", "kUSD", uint8(18), big.NewInt(1000000000000000000),
		"US Dollar", "USD", uint8(4), big.NewInt(10000),
	)
	if err != nil {
		fmt.Println("can't pack price oracle contract params:", err)
		os.Exit(-8)
	}
	if c, err = createContract(runtimeConfig, append(contractsCode[ctPriceOracle], priceOracleParams...)); err != nil {
		fmt.Println("can't create price oracle contract:", err)
		os.Exit(-9)
	}
	r = append(r, c)
	// create network map contract
	netMapAbi, err := abi.JSON(strings.NewReader(network.ContractsContractABI))
	if err != nil {
		fmt.Println("can't parse network stats abi:", err)
		os.Exit(-10)
	}
	netMapParams, err := netMapAbi.Pack("", r[1].addr, r[2].addr, r[0].addr)
	if err != nil {
		fmt.Println("can't pack network map contract params:", err)
		os.Exit(-11)
	}
	if c, err = createContract(runtimeConfig, append(contractsCode[ctNetworkMap], netMapParams...)); err != nil {
		fmt.Println("can't create network map contract:", err)
		os.Exit(-12)
	}
	return append(r, c), nil
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

	choice = w.read()
	var ownerAddr *common.Address
	switch {
	case choice == "" || choice == "1":
		genesis.Config.Tendermint = &params.TendermintConfig{Rewarded: true}
		genesis.ExtraData = make([]byte, 32)

		fmt.Println()
		for ownerAddr == nil {
			fmt.Println("Which account will be used as the owner of the network contracts? (mandatory at least one)")
			ownerAddr = w.readAddress()
		}

		contractsData, err := createContracts(*ownerAddr, map[contractType][]byte{
			ctNetworkMap:   common.FromHex(network.ContractsContractBin),
			ctMToken:       common.FromHex(network.MusdContractBin),
			ctNetworkStats: common.FromHex(network.NetworkContractBin),
			ctPriceOracle:  common.FromHex(network.PriceOracleContractBin),
		})
		if err != nil {
			log.Crit("Failed to create contracts", "err", err)
		}

		for _, contract := range contractsData {
			genesis.Alloc[contract.addr] = core.GenesisAccount{
				Code:    contract.code,
				Storage: contract.storage,
				Balance: common.Big0,
			}
		}

		log.Info("the owner account will be pre-funded with 1 coin", "address", ownerAddr)
		genesis.Alloc[*ownerAddr] = core.GenesisAccount{Balance: new(big.Int).SetUint64(1000000000000000000)}

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
	fmt.Println("Specify your chain/network ID if you want an explicit one (default = random)")
	genesis.Config.ChainID = new(big.Int).SetUint64(uint64(w.readDefaultInt(rand.Intn(65536))))

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
