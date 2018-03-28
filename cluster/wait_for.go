package cluster

import (
	"errors"
	"time"
)

var WaitForTimeout = errors.New("WaitFor timeout")

func WaitFor(tickTime, timeout time.Duration, condition func() bool) error {
	timeoutTime := time.After(timeout)
	tick := time.Tick(tickTime)

	for {
		select {
		case <-timeoutTime:
			return WaitForTimeout
		case <-tick:
			if condition() {
				return nil
			}
		}
	}
}
