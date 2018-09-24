package impl

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/kowala-tech/kcoin/e2e/cluster"
)

const faucetPort = 8080

type FaucetContext struct {
	globalCtx *Context

	nodeRunning bool
	nodeSpec    *cluster.NodeSpec

	lastResponse *http.Response
}

func NewFaucetContext(parentCtx *Context) *FaucetContext {
	ctx := &FaucetContext{
		globalCtx: parentCtx,
	}
	return ctx
}

func (ctx *FaucetContext) Reset() {
}

func (ctx *FaucetContext) TheFaucetNodeIsRunning(account, password string) error {
	acct := ctx.globalCtx.accounts[account]
	acctRaw, err := ctx.globalCtx.AccountsStorage.Export(acct, password, password)
	if err != nil {
		return err
	}
	spec, err := cluster.FaucetSpec(ctx.globalCtx.nodeSuffix, ctx.globalCtx.bootnode, ctx.globalCtx.genesis, acctRaw, password, faucetPort)
	if err != nil {
		return err
	}

	if err := ctx.globalCtx.nodeRunner.Run(spec); err != nil {
		return err
	}

	ctx.nodeRunning = true
	ctx.nodeSpec = spec

	return nil
}

func (ctx *FaucetContext) IFetchOnTheFaucet(url string) error {
	ip := ctx.globalCtx.nodeRunner.HostIP()
	resp, err := http.Get(fmt.Sprintf("http://%v:%v/", ip, faucetPort))
	if err != nil {
		return err
	}
	ctx.lastResponse = resp

	return nil
}

func (ctx *FaucetContext) TheStatusCodeIs(expectedCode int) error {
	if ctx.lastResponse == nil {
		return errors.New("No http request performed")
	}
	if ctx.lastResponse.StatusCode != expectedCode {
		return fmt.Errorf("Status code doesn't match. Expected %v, received %v", expectedCode, ctx.lastResponse.StatusCode)
	}
	return nil
}
