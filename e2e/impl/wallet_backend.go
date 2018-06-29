package impl

import "github.com/DATA-DOG/godog"

type WalletBackendContext struct {
	globalCtx *Context
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
	return godog.ErrPending
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
