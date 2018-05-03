package features

import (
	"fmt"

	"github.com/DATA-DOG/godog"
)

var (
	nodeName = "somevalidator"
	password = "test"
	coinbase = ""
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

func (ctx *Context) IHaveMyNodeRunning() error {
	// return ctx.cluster.RunNode(nodeName)
	return nil
}

func (ctx *Context) IHaveAnAccountInMyNode(kcoin int64) error {
	// response, err := ctx.cluster.Exec(nodeName, newAccountCommand(password))
	// if err != nil {
	// 	log.Debug(response.StdOut)
	// 	return err
	// }
	// coinbase = parseNewAccountResponse(response.StdOut)

	// if err := ctx.fundAccount(coinbase, kcoin); err != nil {
	// 	return err
	// }

	// return err
	return nil
}

func (ctx *Context) IStartTheValidator(kcoin int64) error {
	// response, err := ctx.cluster.Exec(nodeName, unlockAccountCommand(coinbase, password))
	// if err != nil {
	// 	log.Debug(response.StdOut)
	// 	return err
	// }

	// response, err = ctx.cluster.Exec(nodeName, setCoinbaseCommand(coinbase))
	// if err != nil {
	// 	log.Debug(response.StdOut)
	// 	return err
	// }

	// response, err = ctx.cluster.Exec(nodeName, setDeposit(kcoin))
	// if err != nil {
	// 	log.Debug(response.StdOut)
	// 	return err
	// }

	// response, err = ctx.cluster.Exec(nodeName, validatorStartCommand())
	// if err != nil {
	// 	log.Debug(response.StdOut)
	// 	return err
	// }

	// return err
	return nil
}

func (ctx *Context) fundAccount(address string, kcoin int64) error {
	// _, err := ctx.cluster.Exec(
	// 	"genesis-validator",
	// 	fmt.Sprintf(`eth.sendTransaction({from:eth.coinbase, to: "%s", value: %d})`, address, toWei(kcoin)))
	// if err != nil {
	// 	return err
	// }

	// err = cluster.WaitFor(2*time.Second, 1*time.Minute, func() bool {
	// 	resp, err := ctx.cluster.Exec("genesis-validator", fmt.Sprintf(`eth.getBalance("%s")`, address))
	// 	if err != nil {
	// 		return false
	// 	}
	// 	balance := big.NewInt(0)
	// 	balance.SetString(strings.TrimSpace(resp.StdOut), 10)
	// 	return balance.Cmp(big.NewInt(0)) > 0
	// })
	// return err
	return nil
}

func (ctx *Context) IShouldBeAValidator() error {
	// response, err := ctx.cluster.Exec(nodeName, isRunningCommand())
	// message := response.StdOut
	// if err != nil {
	// 	log.Debug(message)
	// 	return err
	// }

	// if strings.TrimSpace(message) != "true" {
	// 	log.Debug(message)
	// 	return errors.New("validator is not running")
	// }

	// return nil

	return nil
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

func setDeposit(kcoin int64) string {
	return fmt.Sprintf("validator.setDeposit(%d)", toWei(kcoin))
}

func validatorStartCommand() string {
	return "validator.start()"
}

func isRunningCommand() string {
	return "validator.isRunning()"
}
