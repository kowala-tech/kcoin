package features

import (
	"fmt"
	"math/big"
	"time"

	"github.com/kowala-tech/kcoin/cluster"
	"github.com/kowala-tech/kcoin/params"
)

func toWei(kcoin int64) *big.Int {
	return new(big.Int).Mul(big.NewInt(kcoin), big.NewInt(params.Ether))
}

func waitFor(errorMessage string, tickTime, timeout time.Duration, condition func() bool) error {
	err := cluster.WaitFor(tickTime, timeout, condition)
	if err == cluster.WaitForTimeout {
		return fmt.Errorf("Timeout error at: %v", errorMessage)
	}
	return err
}
