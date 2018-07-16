package tendermint

import (
	"math/big"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/params"
)

var (
	initialBlockReward = new(big.Int).Mul(new(big.Int).SetUint64(42), big.NewInt(params.Kcoin))
	initialCap         = new(big.Int).SetUint64(82)
	adjustmentFactor   = 1.0001
	lowSupplyMetric    = new(big.Int).SetUint64(1000000)
)

func cap(currency Currency, blockNumber *big.Int) (*big.Int, error) {
	prevSupply, err := currency.Supply(false)
	if err != nil {
		return nil, err
	}
	if (blockNumber.Cmp(common.Big1) > 0) && !hasLowCoinSupply(prevSupply) {
		return new(big.Int).Div(prevSupply, new(big.Int).SetUint64(1000)), nil
	}
	return initialCap, nil
}

func hasLowCoinSupply(supply *big.Int) bool {
	if supply.Cmp(lowSupplyMetric) < 0 {
		return true
	}
	return false
}
