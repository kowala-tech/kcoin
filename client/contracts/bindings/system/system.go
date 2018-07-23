package system

import (
	"math/big"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/params"
)

//go:generate solc --allow-paths ., --abi --bin --overwrite -o build github.com/kowala-tech/kcoin/client/contracts/=../../truffle/contracts openzeppelin-solidity/=../../truffle/node_modules/openzeppelin-solidity/  ../../truffle/contracts/system/System.sol
//go:generate ../../../build/bin/abigen -abi build/System.abi -bin build/System.bin -pkg system -type System -out ./gen_system.go

var mapSystemToAddr = map[uint64]common.Address{
	params.TestnetChainConfig.ChainID.Uint64(): common.HexToAddress("0x4C55B59340FF1398d6aaE362A140D6e93855D4A5"),
}

type Storage struct {
	PrevCurrencyPrice *big.Int
	CurrencyPrice     *big.Int
	CurrencySupply    *big.Int
	PrevMintedAmount  *big.Int
}

func (st * Storage)

type System interface {
	Variables() Storage
	Address() common.Address
}

type system struct {
}
