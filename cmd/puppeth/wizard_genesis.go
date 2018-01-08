// Copyright 2017 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/kowala-tech/kUSD/accounts/abi"
	"github.com/kowala-tech/kUSD/common"
	nc "github.com/kowala-tech/kUSD/contracts/network"
	"github.com/kowala-tech/kUSD/core"
	"github.com/kowala-tech/kUSD/core/state"
	"github.com/kowala-tech/kUSD/core/vm"
	"github.com/kowala-tech/kUSD/core/vm/runtime"
	"github.com/kowala-tech/kUSD/ethdb"
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

func (vmt *vmTracer) CaptureEnd(output []byte, gasUsed uint64, t time.Duration) error {
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
	memDb, err := ethdb.NewMemDatabase()
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
	priceOracleAbi, err := abi.JSON(strings.NewReader(nc.PriceOracleContractABI))
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
	netMapAbi, err := abi.JSON(strings.NewReader(nc.NetworkContractsMapContractABI))
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

func fromHex(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return b
}

// makeGenesis creates a new genesis struct based on some user input.
func (w *wizard) makeGenesis() {
	// Construct a default genesis block
	genesis := &core.Genesis{
		Timestamp:  uint64(time.Now().Unix()),
		GasLimit:   4700000,
		Difficulty: big.NewInt(1048576),
		Alloc:      make(core.GenesisAlloc),
		Config: &params.ChainConfig{
			HomesteadBlock: big.NewInt(1),
			EIP150Block:    big.NewInt(2),
			EIP155Block:    big.NewInt(3),
			EIP158Block:    big.NewInt(3),
		},
	}
	// Figure out which consensus engine to choose
	fmt.Println()
	fmt.Println("Which consensus engine to use? (default = clique)")
	fmt.Println(" 1. Ethash - proof-of-work")
	fmt.Println(" 2. Clique - proof-of-authority")

	choice := w.read()
	switch {
	case choice == "1":
		// In case of ethash, we're pretty much done
		genesis.Config.Ethash = new(params.EthashConfig)
		genesis.ExtraData = make([]byte, 32)

	case choice == "" || choice == "2":
		// In the case of clique, configure the consensus parameters
		genesis.Difficulty = big.NewInt(1)
		genesis.Config.Clique = &params.CliqueConfig{
			Period: 15,
			Epoch:  30000,
		}
		fmt.Println()
		fmt.Println("How many seconds should blocks take? (default = 15)")
		genesis.Config.Clique.Period = uint64(w.readDefaultInt(15))

		// We also need the initial list of signers
		fmt.Println()
		fmt.Println("Which accounts are allowed to seal? (mandatory at least one)")

		var signers []common.Address
		for {
			if address := w.readAddress(); address != nil {
				signers = append(signers, *address)
				continue
			}
			if len(signers) > 0 {
				break
			}
		}
		// Sort the signers and embed into the extra-data section
		for i := 0; i < len(signers); i++ {
			for j := i + 1; j < len(signers); j++ {
				if bytes.Compare(signers[i][:], signers[j][:]) > 0 {
					signers[i], signers[j] = signers[j], signers[i]
				}
			}
		}
		genesis.ExtraData = make([]byte, 32+len(signers)*common.AddressLength+65)
		for i, signer := range signers {
			copy(genesis.ExtraData[32+i*common.AddressLength:], signer[:])
		}

		// Run as rewarded clique ?
	Outer:
		for {
			fmt.Println()
			fmt.Println("Vanila or Rewarded clique ? [v/R]")
			choice := strings.TrimSpace(strings.ToLower(w.read()))
			var ownerAddr *common.Address
			switch choice {
			case "vanila", "v":
				break Outer
			case "rewarded", "r", "":
				genesis.Config.Clique.Rewarded = true
				for ownerAddr == nil {
					fmt.Println("the network contracts need an address to be set as owner")
					ownerAddr = w.readAddress()
				}
				contractsData, err := createContracts(*ownerAddr, map[contractType][]byte{
					ctNetworkMap:   fromHex(nc.NetworkContractsMapContractBin),
					ctMToken:       fromHex(nc.MusdContractBin),
					ctNetworkStats: fromHex(nc.NetworkStatsContractBin),
					ctPriceOracle:  fromHex(nc.PriceOracleContractBin),
				})
				if err != nil {
					fmt.Println("can't create contracts:", err)
					os.Exit(-1)
				}
				for _, cd := range contractsData {
					genesis.Alloc[cd.addr] = core.GenesisAccount{
						Code:    cd.code,
						Storage: cd.storage,
						Balance: common.Big0,
					}
				}
				break Outer
			default:
			}
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
	// // Add a batch of precompile balances to avoid them getting deleted
	// for i := int64(0); i < 256; i++ {
	// 	genesis.Alloc[common.BigToAddress(big.NewInt(i))] = core.GenesisAccount{Balance: big.NewInt(1)}
	// }
	fmt.Println()

	// Query the user for some custom extras
	fmt.Println()
	fmt.Println("Specify your chain/network ID if you want an explicit one (default = random)")
	genesis.Config.ChainId = new(big.Int).SetUint64(uint64(w.readDefaultInt(rand.Intn(65536))))

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
