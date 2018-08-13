package impl

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/DATA-DOG/godog"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/knode"
	"github.com/kowala-tech/kcoin/client/log"
	"github.com/kowala-tech/kcoin/e2e/cluster"
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

func (ctx *ValidationContext) CurrentBlock() (uint64, error) {
	return ctx.globalCtx.CurrentBlock()
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
	return ctx.waiter.Do(
		ctx.globalCtx.makeExecFunc(ctx.nodeID(), validatorStartCommand(kcoin)),
		func() error {
			pending, err := ctx.globalCtx.client.PendingTransactionCount(context.Background())
			if err != nil {
				return err
			}
			if pending != 0 {
				return errors.New("there are some pending transactions")
			}

			res := &cluster.ExecResponse{}
			if err := ctx.globalCtx.execCommand(ctx.nodeID(), isValidatingCommand(), res); err != nil {
				return err
			}
			if strings.TrimSpace(res.StdOut) != "true" {
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
	ctx.nodeRunning = true

	spec := cluster.NewKcoinNodeBuilder().
		WithBootnode(ctx.globalCtx.bootnode).
		WithLogLevel(6).
		WithID(ctx.nodeID().String()).
		WithSyncMode("full").
		WithNetworkId(ctx.globalCtx.chainID.String()).
		WithGenesis(ctx.globalCtx.genesis).
		WithCoinbase(ctx.globalCtx.accounts[account]).
		WithAccount(ctx.globalCtx.AccountsStorage, ctx.globalCtx.accounts[account]).
		WithAccount(ctx.globalCtx.AccountsStorage, ctx.globalCtx.mtokensSeederAccount).
		NodeSpec()

	if err := ctx.globalCtx.nodeRunner.Run(spec); err != nil {
		return err
	}

	return nil
}

func (ctx *ValidationContext) IWithdrawMyNodeFromValidation() error {
	return ctx.waiter.Do(
		ctx.globalCtx.makeExecFunc(ctx.nodeID(), stopValidatingCommand()),
		func() error {
			res := &cluster.ExecResponse{}
			if err := ctx.globalCtx.execCommand(ctx.nodeID(), isValidatingCommand(), res); err != nil {
				return err
			}
			if strings.TrimSpace(res.StdOut) != "false" {
				return errors.New("validator is not running")
			}
			return nil
		})
}

func (ctx *ValidationContext) ThereShouldBeTokensAvailableToMeAfterDays(expectedMTokens int64, days int) error {
	expectedWei := toWei(expectedMTokens)
	deposit, err := ctx.getMTokensDeposit()
	if err != nil {
		return err
	}

	err = ctx.isMTokensDepositExact(deposit, expectedWei)
	if err != nil {
		return err
	}

	daysExpected := time.Hour * 24 * time.Duration(days)
	expectedDate := time.Now().Add(daysExpected)
	if isSameDay(expectedDate, deposit.AvailableAt.Time()) {
		return errors.New(fmt.Sprintf("deposit available not within %d days(%f hours), available at %s",
			days, daysExpected.Hours(), deposit.AvailableAt.Time().String()))
	}

	return nil
}

func (ctx *ValidationContext) isMTokensDepositExact(deposit *Deposit, expectedMTokens *big.Int) error {
	if expectedMTokens.Cmp(deposit.Value) != 0 {
		return errors.New(fmt.Sprintf("kcoins don't match expected %d kcoins got %d", expectedMTokens, *deposit.Value))
	}

	return nil
}

func (ctx *ValidationContext) getMTokensDeposit() (*Deposit, error) {
	res := &cluster.ExecResponse{}
	if err := ctx.globalCtx.execCommand(ctx.nodeID(), getDepositsCommand(), res); err != nil {
		return nil, err
	}

	deposit, err := parseDepositResponse(res.StdOut)
	if err != nil {
		log.Debug(res.StdOut)
		return nil, err
	}

	return &deposit, nil
}

func isSameDay(date1, date2 time.Time) bool {
	expectedYear, expectedMonth, expectedDay := date1.Date()
	availableYear, availableMonth, availableDay := date2.Date()
	return expectedYear != availableYear ||
		expectedMonth != availableMonth ||
		expectedDay != availableDay
}

func (ctx *ValidationContext) MyNodeShouldBeNotBeAValidator() error {
	res := &cluster.ExecResponse{}
	if err := ctx.globalCtx.execCommand(ctx.nodeID(), isValidatingCommand(), res); err != nil {
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
}

func (ctx *ValidationContext) MyNodeIsAlreadySynchronised() error {
	return common.WaitFor("node is synchronised", 200*time.Millisecond, time.Second*20, func() error {
		res := &cluster.ExecResponse{}
		if err := ctx.globalCtx.execCommand(ctx.nodeID(), isSyncedCommand(), res); err != nil {
			return err
		}
		if strings.TrimSpace(res.StdOut) != "true" {
			log.Debug(res.StdOut)
			return errors.New("node is not synchronised")
		}
		return nil
	})
}

func isError(s string) error {
	if strings.HasPrefix(s, "Error: EOF") {
		return nil
	}
	if strings.HasPrefix(s, "Error:") {
		return errors.New(s)
	}
	return nil
}

func blockNumberCommand() []string {
	return cluster.KcoinExecCommand("eth.blockNumber")
}

func isSyncedCommand() []string {
	return cluster.KcoinExecCommand("eth.blockNumber > 1 && net.peerCount > 0 && eth.syncing == false")
}

func validatorStartCommand(mtokens int64) []string {
	return cluster.KcoinExecCommand(fmt.Sprintf("validator.start(%d)", toWei(mtokens)))
}

func stopValidatingCommand() []string {
	return cluster.KcoinExecCommand("validator.stop()")
}

func isValidatingCommand() []string {
	return cluster.KcoinExecCommand("validator.isValidating()")
}

func getDepositsCommand() []string {
	return cluster.KcoinExecCommand("validator.getDeposits()")
}

func getTokenBalance(at common.Address) []string {
	return cluster.KcoinExecCommand(fmt.Sprintf("mtoken.getBalance('%s')", at.String()))
}

func transferTokens(transferArgs knode.TransferArgs) []string {
	args, _ := json.Marshal(transferArgs)
	return cluster.KcoinExecCommand(fmt.Sprintf("mtoken.transfer(%s)", string(args)))
}

func mintTokens(transferArgs knode.TransferArgs, pass string) []string {
	args, _ := json.Marshal(transferArgs)
	return cluster.KcoinExecCommand(fmt.Sprintf("mtoken.mint(%s, %q)", string(args), pass))
}
