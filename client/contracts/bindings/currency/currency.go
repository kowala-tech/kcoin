package currency

import (
	"math/big"

	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/contracts/bindings"
	"github.com/kowala-tech/kcoin/client/params"
)

var mapCurrencyToAddr = map[uint64]common.Address{
	params.TestnetChainConfig.ChainID.Uint64(): common.HexToAddress("0x4C55B59340FF1398d6aaE362A140D6e93855D4A5"),
}

type Kcoin interface {
	Supply(pending bool) (*big.Int, error)
}

type kcoin struct {
	addr     common.Address
	currency Currency
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
		currency: curr,
	}
}

func (curr *currency) Address() common.Address {
	return curr.addr
}

func (curr *currency) Supply() (*big.Int, error) {

}
