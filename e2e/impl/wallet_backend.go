package impl

import (
	"github.com/DATA-DOG/godog"
	"github.com/kowala-tech/kcoin/e2e/cluster"
)

type WalletBackendContext struct {
	globalCtx *Context
	nodeRunning bool
}

func NewWalletBackendContext(parentCtx *Context) *WalletBackendContext {
	ctx := &WalletBackendContext{
		globalCtx: parentCtx,
	}
	return ctx
}

func (ctx *WalletBackendContext) Reset() {
}

func (ctx *WalletBackendContext) TheWalletBackendNodeIsRunning() error {
	if ctx.nodeRunning {
		return nil
	}

	spec, err := cluster.WalletBackendSpec(ctx.globalCtx.nodeSuffix)
	if err != nil {
		return err
	}

	if err := ctx.globalCtx.nodeRunner.Run(spec, ctx.globalCtx.GetScenarioNumber()); err != nil {
		return err
	}

	ctx.nodeRunning = true

	return nil
}

func (ctx *WalletBackendContext) ICheckTheCurrentBlockHeightInTheWalletBackendAPI() error {
	return godog.ErrPending
}

func (ctx *WalletBackendContext) IWaitForBlocks(arg1 int) error {
	return godog.ErrPending
}

func (ctx *WalletBackendContext) TheNewBlockHeightInTheWalletBackendAPIHasIncreasedByAtLeast(arg1 int) error {
	return godog.ErrPending
}

func (ctx *WalletBackendContext) ITransferKcoinFromAToBUsingTheWalletAPI(arg1 int) error {
	return godog.ErrPending
}

func (ctx *WalletBackendContext) TheTransactionIsListedInTheWalletBackendAPI() error {
	return godog.ErrPending
}

func (ctx *WalletBackendContext) TheBalanceOfAUsingTheWalletBackendShouldBeAroundKcoins(arg1 int) error {
	return godog.ErrPending
}

func (ctx *WalletBackendContext) TheBalanceOfBUsingTheWalletBackendShouldBeKcoins(arg1 int) error {
	return godog.ErrPending
}
