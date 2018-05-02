package features

import (
	"errors"
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/kowala-tech/kcoin/cluster"
	"github.com/kowala-tech/kcoin/log"
	"math/big"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type ValidationContext struct {
	globalCtx       *Context
	nodeName        string
	accountPassword string
	account         string
	nodeRunning     bool
}

func NewValidationContext(parentCtx *Context) *ValidationContext {
	return &ValidationContext{
		globalCtx:       parentCtx,
		nodeName:        "validator-under-test",
		accountPassword: "test",
		nodeRunning:     false,
	}
}

func (ctx *ValidationContext) IStopValidation() error {
	return godog.ErrPending
}

func (ctx *ValidationContext) IWaitForTheUnbondingPeriodToBeOver() error {
	return godog.ErrPending
}

func (ctx *ValidationContext) IHaveMyNodeRunning() error {
	if ctx.nodeRunning {
		return nil
	}
	ctx.nodeRunning = true
	return ctx.globalCtx.cluster.RunNode(ctx.nodeName)
}

func (ctx *ValidationContext) stopNode() error {
	if !ctx.nodeRunning {
		return nil
	}
	return ctx.globalCtx.cluster.DeletePod(ctx.nodeName)
}

func (ctx *ValidationContext) IWithdrawMyNodeFromValidation() error {
	response, err := ctx.globalCtx.cluster.Exec(ctx.nodeName, stopValidatingCommand())
	if err != nil {
		log.Debug(response.StdOut)
		return err
	}
	return nil
}

func (ctx *ValidationContext) ThereShouldBeTokensAvailableToMeAfterDays(expectedKcoins, days int) error {
	response, err := ctx.globalCtx.cluster.Exec(ctx.nodeName, getDepositsCommand())
	if err != nil {
		log.Debug(response.StdOut)
		return err
	}

	availableAt, kcoins, err := parseDepositsResponse(response.StdOut)
	if err != nil {
		log.Debug(response.StdOut)
		return err
	}

	if expectedKcoins != kcoins {
		return errors.New(fmt.Sprintf("kcoins don't match expected %d kcoins got %d", expectedKcoins, kcoins))
	}

	daysExpected := time.Hour * 24 * time.Duration(days)
	expectedDate := time.Now().Add(daysExpected)
	if isSameDay(expectedDate, availableAt) {
		return errors.New(fmt.Sprintf("deposit available not within %d days, available at %s", daysExpected, availableAt))
	}

	return nil
}

func isSameDay(date1, date2 time.Time) bool {
	expectedYear, expectedMonth, expectedDay := date1.Date()
	availableYear, availableMonth, availableDay := date2.Date()
	return expectedYear != availableYear ||
		expectedMonth != availableMonth ||
		expectedDay != availableDay
}

func parseDepositsResponse(value string) (time.Time, int, error) {
	re := regexp.MustCompile("\"(.+)\",\\s+value:\\s(\\d+)")
	matches := re.FindAllStringSubmatch(value, -1)
	if len(matches) == 0 || len(matches[0]) < 3 {
		return time.Now(), 0, errors.New("cant find AvailableAt and Value on response")
	}
	return parseDate(matches[0][1]), parseKCoins(matches[0][2]), nil
}

func parseKCoins(kcoins string) int {
	result, _ := strconv.Atoi(kcoins)
	return result
}

func parseDate(date string) time.Time {
	const longForm = "2006-01-02 15:04:05 -0700 MST"
	t, _ := time.Parse(longForm, date)
	return t
}

func (ctx *ValidationContext) MyNodeShouldBeNotBeAValidator() error {
	isValidator, err := ctx.isValidator()
	if err != nil {
		return err
	}

	if isValidator {
		return errors.New("validator is still running")
	}

	return nil
}

func (ctx *ValidationContext) IHaveAnAccountInMyNode(kcoin int64) error {
	response, err := ctx.globalCtx.cluster.Exec(ctx.nodeName, newAccountCommand(ctx.accountPassword))
	if err != nil {
		log.Debug(response.StdOut)
		return err
	}
	ctx.account = parseNewAccountResponse(response.StdOut)

	if err := ctx.fundAccount(ctx.account, kcoin); err != nil {
		return err
	}

	return err
}

func (ctx *ValidationContext) IStartTheValidator(kcoin int64) error {
	response, err := ctx.globalCtx.cluster.Exec(ctx.nodeName, unlockAccountCommand(ctx.account, ctx.accountPassword))
	if err != nil {
		log.Debug(response.StdOut)
		return err
	}

	response, err = ctx.globalCtx.cluster.Exec(ctx.nodeName, setCoinbaseCommand(ctx.account))
	if err != nil {
		log.Debug(response.StdOut)
		return err
	}

	response, err = ctx.globalCtx.cluster.Exec(ctx.nodeName, setDeposit(kcoin))
	if err != nil {
		log.Debug(response.StdOut)
		return err
	}

	response, err = ctx.globalCtx.cluster.Exec(ctx.nodeName, validatorStartCommand())
	if err != nil {
		log.Debug(response.StdOut)
		return err
	}

	return err
}

func (ctx *ValidationContext) fundAccount(address string, kcoin int64) error {
	_, err := ctx.globalCtx.cluster.Exec(
		"genesis-validator",
		fmt.Sprintf(`eth.sendTransaction({from:eth.coinbase, to: "%s", value: %d})`, address, toWei(kcoin)))
	if err != nil {
		return err
	}

	err = cluster.WaitFor(2*time.Second, 1*time.Minute, func() bool {
		resp, err := ctx.globalCtx.cluster.Exec("genesis-validator", fmt.Sprintf(`eth.getBalance("%s")`, address))
		if err != nil {
			return false
		}
		balance := big.NewInt(0)
		balance.SetString(strings.TrimSpace(resp.StdOut), 10)
		return balance.Cmp(big.NewInt(0)) > 0
	})
	return err
}

func (ctx *ValidationContext) IShouldBeAValidator() error {
	isValidator, err := ctx.isValidator()
	if err != nil {
		return err
	}

	if !isValidator {
		return errors.New("validator is not running")
	}

	return nil
}

func (ctx *ValidationContext) isValidator() (bool, error) {
	response, err := ctx.globalCtx.cluster.Exec(ctx.nodeName, isRunningCommand())
	message := response.StdOut
	if err != nil {
		log.Debug(message)
		return false, err
	}

	return strings.TrimSpace(message) == "true", nil
}

func (ctx *ValidationContext) Reset() {
	ctx.stopNode()
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

func stopValidatingCommand() string {
	return "validator.stop()"
}

func isRunningCommand() string {
	return "validator.isRunning()"
}

func getDepositsCommand() string {
	return "validator.getDeposits()"
}
