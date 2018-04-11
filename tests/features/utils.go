package features

import (
	"fmt"
	"math/big"
	"time"

	"github.com/kowala-tech/kcoin/cluster"
)

var kcoinWei = big.NewInt(1000000000000000000)

func toWei(kcoin int64) *big.Int {
	res := big.NewInt(kcoin)
	return res.Mul(res, kcoinWei)
}

func waitFor(errorMessage string, tickTime, timeout time.Duration, condition func() bool) error {
	err := cluster.WaitFor(tickTime, timeout, condition)
	if err == cluster.WaitForTimeout {
		return fmt.Errorf("Timeout error at: %v", errorMessage)
	}
	return err
}
