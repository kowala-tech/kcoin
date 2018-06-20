package features

import (
	"math/big"

	"github.com/kowala-tech/kcoin/client/params"
)

func toWei(kcoin int64) *big.Int {
	return new(big.Int).Mul(big.NewInt(kcoin), big.NewInt(params.KUSD))
}
