package tendermint

import (
	"math/big"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/params"
)

var (
	initialMintedAmount      = new(big.Int).Mul(new(big.Int).SetUint64(42), big.NewInt(params.Kcoin))
	initialCap               = new(big.Int).Mul(new(big.Int).SetUint64(82), big.NewInt(params.Kcoin))
	adjustmentFactor         = new(big.Int).SetUint64(10000)
	lowSupplyMetric          = new(big.Int).Mul(new(big.Int).SetUint64(1000000), big.NewInt(params.Kcoin))
	stabilizedPrice          = new(big.Int).Mul(common.Big1, big.NewInt(params.Kcoin))
	maxUnderNormalConditions = new(big.Int).SetUint64(1e12)
)

func mintedAmount(blockNumber, currentPrice, prevPrice, prevSupply, prevMintedAmount *big.Int) *big.Int {
	if blockNumber.Cmp(common.Big1) == 0 {
		return initialMintedAmount
	}

	var adjustedAmount *big.Int
	if currentPrice.Cmp(prevPrice) >= 0 && prevPrice.Cmp(stabilizedPrice) > 0 {
		adjustedAmount = new(big.Int).Add(prevMintedAmount, new(big.Int).Div(prevMintedAmount, adjustmentFactor))
		return common.Min(adjustedAmount, cap(blockNumber, prevSupply))
	}

	adjustedAmount = new(big.Int).Sub(prevMintedAmount, new(big.Int).Div(prevMintedAmount, adjustmentFactor))
	return common.Max(adjustedAmount, maxUnderNormalConditions)
}

func cap(blockNumber, prevSupply *big.Int) *big.Int {
	if (blockNumber.Cmp(common.Big1) > 0) && !hasLowSupply(prevSupply) {
		return new(big.Int).Div(prevSupply, new(big.Int).SetUint64(10000))
	}
	return initialCap
}

func hasLowSupply(supply *big.Int) bool {
	if supply.Cmp(lowSupplyMetric) < 0 {
		return true
	}
	return false
}
