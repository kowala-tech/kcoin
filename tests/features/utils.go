package features

import (
	"fmt"
	"math/big"
	"time"

	"github.com/kowala-tech/kcoin/params"
)

func toWei(kcoin int64) *big.Int {
	return new(big.Int).Mul(big.NewInt(kcoin), big.NewInt(params.Ether))
}

func waitFor(errorMessage string, tickTime, timeout time.Duration, condition func() bool) error {
	timeoutTime := time.After(timeout)
	tick := time.Tick(tickTime)

	for {
		select {
		case <-timeoutTime:
			return fmt.Errorf("Timeout error: %v", errorMessage)
		case <-tick:
			if condition() {
				return nil
			}
		}
	}
}
