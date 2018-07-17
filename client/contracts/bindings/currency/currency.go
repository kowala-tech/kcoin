package currency

import (
	"math/big"

	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/contracts/bindings"
	"github.com/kowala-tech/kcoin/client/params"
)

//go:generate solc --allow-paths ., --abi --bin --overwrite -o build github.com/kowala-tech/kcoin/client/contracts/=../../truffle/contracts openzeppelin-solidity/=../../truffle/node_modules/openzeppelin-solidity/  ../../truffle/contracts/currency/Currency.sol
//go:generate ../../../build/bin/abigen -abi build/Currency.abi -bin build/Currency.bin -pkg currency -type Currency -out ./gen_currency.go

var mapCurrencyToAddr = map[uint64]common.Address{
	params.TestnetChainConfig.ChainID.Uint64(): common.HexToAddress("0x4C55B59340FF1398d6aaE362A140D6e93855D4A5"),
}

type Currency interface {
	PrevSupply() (*big.Int, error)
	PrevMintedAmount() (*big.Int, error)
	Address() common.Hash
}

type currency struct {
	addr common.Address
	Currency
}

// Binding returns a binding to the current currency contract
func Binding(contractBackend bind.ContractBackend, chainID *big.Int) (Currency, error) {
	addr, ok := mapCurrencyToAddr[chainID.Uint64()]
	if !ok {
		return nil, bindings.ErrNoAddress
	}

	curr, err := NewCurrency(addr, contractBackend)
	if err != nil {
		return nil, err
	}

	return nil, &currency{
		addr:     addr,
		Currency: curr,
	}
}

func (curr *currency) Address() common.Address {
	return curr.addr
}

func (curr *currency) PrevSupply() (*big.Int, error) {

}

func (curr *currency) PrevMintedAmount() (*big.Int, error) {

}
