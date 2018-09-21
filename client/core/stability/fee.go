package stability

import (
	"math/big"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/params"
)

var (
	stabilityIncrease     = new(big.Int).SetUint64(109)
	stabilityTxPercentage = new(big.Int).SetUint64(params.StabilityFeeTxPercentage)
)

// CalcFee returns the stability fee given a compute fee, stability level and tx amount.
func CalcFee(computeFee *big.Int, stabilizationLevel uint64, txAmount *big.Int) *big.Int {
	if stabilizationLevel == 0 {
		return common.Big0
	}

	if txAmount.Cmp(common.Big0) == 0 {
		return computeFee
	}

	// fee = compute fee  * 1.09^r(b)
	lvl := new(big.Int).SetUint64(stabilizationLevel)
	mul := new(big.Int).Exp(stabilityIncrease, lvl, nil)
	div := new(big.Int).Exp(common.Big100, lvl, nil)
	fee := new(big.Int).Div(new(big.Int).Mul(computeFee, mul), div)

	// percentage of tx amount
	maxFee := new(big.Int).Div(new(big.Int).Mul(txAmount, stabilityTxPercentage), common.Big100)

	return common.Min(fee, maxFee)
}
