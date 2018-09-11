package impl

import (
	"github.com/DATA-DOG/godog"
)

func (ctx *Context) IStartANewNode() error {
	return godog.ErrPending
}

func (ctx *Context) MyNodeShouldSyncWithTheNetwork() error {
	return godog.ErrPending
}

func (ctx *Context) MyNodeIsAlreadySynchronised() error {
	return godog.ErrPending
}

func (ctx *Context) IDisconnectMyNodeForBlocksAndReconnectIt(blocks int) error {
	return godog.ErrPending
}

func (ctx *Context) IStartANewNodeWithADifferentNetworkID() error {
	return godog.ErrPending
}

func (ctx *Context) MyNodeShouldNotSyncWithTheNetwork() error {
	return godog.ErrPending
}

func (ctx *Context) IStartANewNodeWithADifferentChainID() error {
	return godog.ErrPending
}

func (ctx *Context) IStartValidatorWithDepositAndCoinbaseA(deposit int) error {
	return godog.ErrPending
}

func (ctx *Context) CrashMyNode() error {
	return ctx.nodeRunner.Stop(ctx.genesisValidatorNodeID)
}

func (ctx *Context) IRestartTheValidator() error {
	err := ctx.runGenesisValidator()
	if err != nil {
		return err
	}

	return nil
}
