package features

import (
	"github.com/DATA-DOG/godog"
)

func (ctx *Context) IStartTheValidator(kcoin int64, accountName string) error {
	err := ctx.cluster.RunNode("somevalidator")
	if err != nil {
		return err
	}

	_, err = ctx.cluster.Exec("somevalidator", "validator.start()")

	return err
}

func (ctx *Context) IShouldBeAValidator() error {
	return godog.ErrPending
}
