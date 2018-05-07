package common

import (
	"fmt"
	"time"
)

func WaitFor(errorMessage string, tickTime, timeout time.Duration, condition func() bool) error {
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
