package common

import (
	"errors"
	"fmt"
	"time"

	"github.com/kowala-tech/kcoin/client/log"
)

func WaitFor(errorMessage string, tickTime, timeout time.Duration, condition func() error) error {
	err := condition()
	if err == nil {
		return nil
	}
	logError(err)

	timeoutTime := time.After(timeout)
	tick := time.Tick(tickTime)

	for {
		select {
		case <-timeoutTime:
			return fmt.Errorf("timeout error: %v", errorMessage)
		case <-tick:
			if err := condition(); err != nil {
				logError(err)
				continue
			}

			return nil
		}
	}
}

var ErrConditionNotMet = errors.New("the condition is not met")

func logError(err error) {
	if err != ErrConditionNotMet {
		log.Warn(fmt.Sprintf("error while executing tha condition: %q", err.Error()))
	}
}

type nodeAPI interface {
	CurrentBlock() (block uint64, err error)
}

type waiter struct {
	api nodeAPI
}

func NewWaiter(api nodeAPI) *waiter {
	return &waiter{api}
}

// Do executes the command on the node and waits 1 block then
func (w *waiter) Do(execFunc func() error, condFuncs ...func() error) error {
	currentBlock, err := w.api.CurrentBlock()
	if err != nil {
		return err
	}

	if err = execFunc(); err != nil {
		return err
	}

	err = w.waitBlocksFrom(currentBlock, 10, condFuncs...)
	if err != nil {
		return err
	}

	return nil
}

func (w *waiter) waitBlocksFrom(block, n uint64, condFuncs ...func() error) error {
	t := time.NewTicker(500 * time.Millisecond)
	timeout := time.NewTimer(2 * time.Duration(n) * time.Second)
	defer t.Stop()

	var (
		err      error
		lastErr  error
		newBlock uint64
	)

	for {
		select {
		case <-timeout.C:
			return fmt.Errorf("timeout. started with block %d, finished with %d. Last error: %v", block, newBlock, lastErr.Error())
		case <-t.C:
			newBlock, err = w.api.CurrentBlock()
			if err != nil {
				return err
			}

			lastErr = runConditions(block, newBlock, condFuncs...)
			if lastErr == nil {
				return nil
			}
		}
	}

	return nil
}

func runConditions(block, newBlock uint64, condFuncs ...func() error) error {
	if len(condFuncs) == 0 {
		return errors.New("No conditions specified")
	}
	for _, condFunc := range condFuncs {
		err := condFunc()
		if err != nil && err != ErrConditionNotMet {
			log.Warn(fmt.Sprintf("error while executing the condition: %q. "+
				"started with block %d, finished with %d",
				err.Error(), block, newBlock))
		}
		if err != nil {
			return err
		}
	}
	return nil
}
