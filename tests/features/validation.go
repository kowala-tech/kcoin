package features

import (
	"github.com/DATA-DOG/godog"
)

func (ctx *Context) IStopValidation() error {
	return godog.ErrPending
}

func (ctx *Context) IWaitForTheUnbondingPeriodToBeOver() error {
	return godog.ErrPending
}

func (ctx *Context) IShouldNotBeAValidator() error {
	return godog.ErrPending
}
