package features

import (
	"github.com/DATA-DOG/godog"
	"fmt"
)

func (ctx *Context) IStartTheValidator(kcoin int64, accountName string) error {
	nodeName := "somevalidator"
	password := "test"
	coinbase := ""

	err := ctx.cluster.RunNode(nodeName)
	if err != nil {
		return err
	}

	response, err := ctx.cluster.Exec(nodeName, newAccountCommand(password))
	if err != nil {
		return err
	}
	coinbase = parseNewAccountResponse(response.StdOut)

	_, err = ctx.cluster.Exec(nodeName, unlockAccountCommand(coinbase, password))
	if err != nil {
		return err
	}

	_, err = ctx.cluster.Exec(nodeName, setCoinbaseCommand(coinbase))
	if err != nil {
		return err
	}

	response, err = ctx.cluster.Exec(nodeName, setDeposit())
	print(response)
	if err != nil {
		return err
	}

	_, err = ctx.cluster.Exec(nodeName, validatorStartCommand())
	if err != nil {
		return err
	}

	return err
}

func (ctx *Context) IShouldBeAValidator() error {
	return godog.ErrPending
}

// parseNewAccountResponse remove first and last char, response comes in format
// "0x7ddba4b656cd3b537f208973bb6f6957e2d3750d"
func parseNewAccountResponse(response string) string {
	return response[1 : len(response)-2]
}

func newAccountCommand(password string) string {
	return fmt.Sprintf("personal.newAccount(\"%s\")", password)
}

func unlockAccountCommand(account, password string) string {
	return fmt.Sprintf("personal.unlockAccount(\"%s\", \"%s\")", account, password)
}

func setCoinbaseCommand(coinbase string) string {
	return fmt.Sprintf("validator.setCoinbase(\"%s\")", coinbase)
}

func setDeposit() string {
	return "validator.setDeposit(100000)"
}

func validatorStartCommand() string {
	return "validator.start()"
}
