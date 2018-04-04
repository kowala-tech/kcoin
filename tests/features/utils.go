package features

import (
	"math/big"
)

var kcoinWei = big.NewInt(1000000000000000000)

func toWei(kcoin int64) *big.Int {
	res := big.NewInt(kcoin)
	return res.Mul(res, kcoinWei)
}
