package features

import (
	"math/big"
)

var kUSDWei = big.NewInt(1000000000000000000)

func toWei(kusd int64) *big.Int {
	res := big.NewInt(kusd)
	return res.Mul(res, kUSDWei)
}
