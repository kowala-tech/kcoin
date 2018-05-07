package features

import (
	"errors"
	"fmt"
	"strings"

	"regexp"
	"strconv"
	"time"

	"github.com/DATA-DOG/godog"
	"github.com/kowala-tech/kcoin/cluster"
	"github.com/kowala-tech/kcoin/log"
)

type ValidationContext struct {
	globalCtx       *Context
	nodeID          cluster.NodeID
	accountPassword string
	account         string
	nodeRunning     bool
}

func NewValidationContext(parentCtx *Context) *ValidationContext {
	return &ValidationContext{
		globalCtx:       parentCtx,
		nodeID:          cluster.NodeID("validator-under-test"),
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

func (ctx *ValidationContext) IStartTheValidator(kcoin int64) error {
	res, err := ctx.globalCtx.nodeRunner.Exec(ctx.nodeID, setDeposit(kcoin))
	if err != nil {
		log.Debug(res.StdOut)
		return err
	}
	res, err = ctx.globalCtx.nodeRunner.Exec(ctx.nodeID, validatorStartCommand())
	if err != nil {
		log.Debug(res.StdOut)
		return err
	}
	return nil
}

func (ctx *ValidationContext) IShouldBeAValidator() error {
	res, err := ctx.globalCtx.nodeRunner.Exec(ctx.nodeID, isRunningCommand())
	if err != nil {
		log.Debug(res.StdOut)
		return err
	}
	if strings.TrimSpace(res.StdOut) != "true" {
		log.Debug(res.StdOut)
		return errors.New("validator is not running")
	}
	return nil
}

func (ctx *ValidationContext) IHaveMyNodeRunning(account string) error {
	if ctx.nodeRunning {
		return nil
	}
	ctx.nodeRunning = true

	spec := cluster.NewKcoinNodeBuilder().
		WithBootnode(ctx.globalCtx.bootnode).
		WithLogLevel(3).
		WithID(ctx.nodeID.String()).
		WithSyncMode("full").
		WithNetworkId(ctx.globalCtx.chainID.String()).
		WithGenesis(ctx.globalCtx.genesis).
		WithAccount(ctx.globalCtx.AccountsStorage, ctx.globalCtx.accounts[account]).
		NodeSpec()

	if err := ctx.globalCtx.nodeRunner.Run(spec); err != nil {
		return err
	}

	return nil
}

func (ctx *ValidationContext) stopNode() error {
	if !ctx.nodeRunning {
		return nil
	}
	return ctx.globalCtx.nodeRunner.Stop(ctx.nodeID)
}

func (ctx *ValidationContext) IWithdrawMyNodeFromValidation() error {
	res, err := ctx.globalCtx.nodeRunner.Exec(ctx.nodeID, stopValidatingCommand())
	if err != nil {
		log.Debug(res.StdOut)
		return err
	}
	return nil
}

func (ctx *ValidationContext) ThereShouldBeTokensAvailableToMeAfterDays(expectedKcoins, days int) error {
	res, err := ctx.globalCtx.nodeRunner.Exec(ctx.nodeID, getDepositsCommand())
	if err != nil {
		log.Debug(res.StdOut)
		return err
	}
	availableAt, kcoins, err := parseDepositsResponse(res.StdOut)
	if err != nil {
		log.Debug(res.StdOut)
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
	// isValidator, err := ctx.isValidator()
	// if err != nil {
	// 	return err
	// }

	// if isValidator {
	// 	return errors.New("validator is still running")
	// }

	// return nil
	return godog.ErrPending
}

func (ctx *ValidationContext) Reset() {
	ctx.nodeRunning = false
	ctx.globalCtx.nodeRunner.Stop(ctx.nodeID)
}

func setDeposit(kcoin int64) []string {
	return cluster.KcoinExecCommand(fmt.Sprintf("validator.setDeposit(%d)", toWei(kcoin)))
}

func validatorStartCommand() []string {
	return cluster.KcoinExecCommand("validator.start()")
}

func stopValidatingCommand() []string {
	return cluster.KcoinExecCommand("validator.stop()")
}

func isRunningCommand() []string {
	return cluster.KcoinExecCommand("validator.isRunning()")
}

func getDepositsCommand() []string {
	return cluster.KcoinExecCommand("validator.getDeposits()")
}
