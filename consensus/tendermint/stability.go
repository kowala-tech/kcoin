package stability

import (
	"math/big"

	"github.com/kowala-tech/kcoin/common"
)

// repeat is used to track how many times we need to increment
// the stability fee
func repeat(blockNumber *big.Int) *big.Int {
	if blockNumber.Cmp(common.Big1) == 0 {
		return common.BigMinus1
	}
}

// CalcStabilityFee returns the stability fee
func CalcStabilityFee(blockNumber, currentPrice, txFees *big.Int) *big.Int {
	if repeat := repeat(blockNumber); repeat.Cmp(common.BigMinus1) == 0 {
		return common.Big0
	}


	if blockNumber.Cmp(common.Big1) || txFees.Cmp(common.Big0) == 0 {
		return common.Big0
	}

	

	return math.
}
