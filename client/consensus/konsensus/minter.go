package Konsensus

import (
	"math/big"

	"github.com/kowala-tech/kcoin/client/common"
)

// Minter represts the person who mints money
type Minter interface {
	Mint(account common.Address, amount *big.Int)
}

type MinterFunc func(common.Address, *big.Int)

func (fn MinterFunc) Mint(account common.Address, amount *big.Int) {
	fn(account, amount)
}

type MinterMiddleware func(Minter) Minter
