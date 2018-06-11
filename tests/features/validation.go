package features

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/DATA-DOG/godog"
	"github.com/kowala-tech/kcoin/cluster"
	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/log"
)

type ValidationContext struct {
	globalCtx       *Context
	accountPassword string
	nodeRunning     bool
	waiter          doer
}

func NewValidationContext(parentCtx *Context) *ValidationContext {
	ctx := &ValidationContext{
		globalCtx:       parentCtx,
		accountPassword: "test",
		nodeRunning:     false,
	}

	ctx.waiter = common.NewWaiter(ctx)
	return ctx
}

func (ctx *ValidationContext) nodeID() cluster.NodeID {
	return cluster.NodeID("validator-under-test-" + ctx.globalCtx.nodeSuffix)
}

func (ctx *ValidationContext) IStopValidation() error {
	return godog.ErrPending
}

func (ctx *ValidationContext) IWaitForTheUnbondingPeriodToBeOver() error {
	return godog.ErrPending
}

func (ctx *ValidationContext) IStartTheValidator(kcoin int64) error {
	err := ctx.waiter.Do(ctx.makeExecFunc(setDeposit(kcoin)))
	if err != nil {
		return err
	}

	return ctx.waiter.Do(
		ctx.makeExecFunc(validatorStartCommand()),
		func() error {
			res, err := ctx.globalCtx.nodeRunner.Exec(ctx.nodeID(), isRunningCommand())
			if err != nil {
				log.Debug(res.StdOut)
				return err
			}
			if strings.TrimSpace(res.StdOut) != "true" {
				log.Debug(res.StdOut)
				return errors.New("validator is not running")
			}
			return nil
		})
}

func (ctx *ValidationContext) IWaitForMyNodeToBeSynced() error {
	return common.WaitFor("timeout waiting for node sync", time.Second, time.Second*5, func() error {
		return ctx.MyNodeIsAlreadySynchronised()
	})
}

func (ctx *ValidationContext) IHaveMyNodeRunning(account string) error {
	if ctx.nodeRunning {
		return nil
	}

	spec := cluster.NewKcoinNodeBuilder().
		WithBootnode(ctx.globalCtx.bootnode).
		WithLogLevel(3).
		WithID(ctx.nodeID().String()).
		WithSyncMode("full").
		WithNetworkId(ctx.globalCtx.chainID.String()).
		WithGenesis(ctx.globalCtx.genesis).
		WithAccount(ctx.globalCtx.AccountsStorage, ctx.globalCtx.accounts[account]).
		NodeSpec()

	if err := ctx.globalCtx.nodeRunner.Run(spec, ctx.globalCtx.scenarioNumber); err != nil {
		return err
	}

	ctx.nodeRunning = true

	return nil
}

func (ctx *ValidationContext) IWithdrawMyNodeFromValidation() error {
	return ctx.waiter.Do(ctx.makeExecFunc(stopValidatingCommand()))
}

func (ctx *ValidationContext) ThereShouldBeTokensAvailableToMeAfterDays(expectedKcoins, days int) error {
	res, err := ctx.globalCtx.nodeRunner.Exec(ctx.nodeID(), getDepositsCommand())
	if err != nil {
		log.Debug(res.StdOut)
		return err
	}

	deposit, err := parseDepositResponse(res.StdOut)
	if err != nil {
		log.Debug(res.StdOut)
		return err
	}

	if expectedKcoins != *deposit.Value {
		return errors.New(fmt.Sprintf("kcoins don't match expected %d kcoins got %d", expectedKcoins, *deposit.Value))
	}

	daysExpected := time.Hour * 24 * time.Duration(days)
	expectedDate := time.Now().Add(daysExpected)
	if isSameDay(expectedDate, deposit.AvailableAt.Time()) {
		return errors.New(fmt.Sprintf("deposit available not within %d days(%f hours), available at %s",
			days, daysExpected.Hours(), deposit.AvailableAt.Time().String()))
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

func (ctx *ValidationContext) MyNodeShouldBeNotBeAValidator() error {
	res, err := ctx.globalCtx.nodeRunner.Exec(ctx.nodeID(), isRunningCommand())
	if err != nil {
		log.Debug(res.StdOut)
		return err
	}
	if strings.TrimSpace(res.StdOut) != "false" {
		log.Debug(res.StdOut)
		return errors.New("validator running")
	}
	return nil
}

func (ctx *ValidationContext) Reset() {
	ctx.nodeRunning = false
	ctx.globalCtx.nodeRunner.Stop(ctx.nodeID())
}

func (ctx *ValidationContext) MyNodeIsAlreadySynchronised() error {
	res, err := ctx.globalCtx.nodeRunner.Exec(ctx.nodeID(), isSyncedCommand())
	if err != nil {
		log.Debug(res.StdOut)
		return err
	}
	if strings.TrimSpace(res.StdOut) != "true" {
		log.Debug(res.StdOut)
		return errors.New("node is not synced")
	}
	return nil
}

func (ctx *ValidationContext) Do(cmd []string, condFunc func() error) error {
	return ctx.waiter.Do(ctx.makeExecFunc(cmd), condFunc)
}

func (ctx *ValidationContext) CurrentBlock() (uint64, error) {
	res, err := ctx.globalCtx.nodeRunner.Exec(ctx.nodeID(), blockNumberCommand())
	if err != nil {
		log.Debug(res.StdOut)
		return 0, err
	}

	return strconv.ParseUint(strings.TrimSpace(res.StdOut), 10, 64)
}

func (ctx *ValidationContext) makeExecFunc(command []string) func() error {
	return func() error {
		res, err := ctx.globalCtx.nodeRunner.Exec(ctx.nodeID(), command)
		if err != nil {
			log.Debug(res.StdOut)
		}
		return err
	}
}

func blockNumberCommand() []string {
	return cluster.KcoinExecCommand("eth.blockNumber")
}

func isSyncedCommand() []string {
	return cluster.KcoinExecCommand("eth.blockNumber > 0 && eth.syncing == false")
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
