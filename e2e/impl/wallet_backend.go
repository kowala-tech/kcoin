package impl

import (
	"github.com/DATA-DOG/godog"
	"github.com/kowala-tech/kcoin/e2e/cluster"
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"github.com/kowala-tech/kcoin/wallet-backend/application/command"
	"math/big"
	"time"
)

type WalletBackendContext struct {
	globalCtx *Context
	nodeRunning bool
	lastBlockRegistered *big.Int
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
	blockHeight, err := getBlockHeight()
	if err != nil {
		return err
	}

	ctx.lastBlockRegistered = blockHeight

	return nil
}

func getBlockHeight() (*big.Int, error) {
	res, err := http.Get(
		fmt.Sprintf("http://%s:8080/api/blockheight", "localhost"),
	)
	if err != nil {
		return nil, fmt.Errorf("error connecting to wallet backend to get block height. %s", err)
	}

	rawResp, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("error parsing response from wallet backend to get block height. %s", err)
	}

	var blockHeightResponse *command.BlockHeightResponse
	err = json.Unmarshal(rawResp, &blockHeightResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling from json response. %s", err)
	}

	return blockHeightResponse.BlockHeight, nil
}

func (ctx *WalletBackendContext) IWaitForBlocks(arg1 int) error {
	return godog.ErrPending
}

func (ctx *WalletBackendContext) TheNewBlockHeightInTheWalletBackendAPIHasIncreasedByAtLeast(arg1 int) error {
	actualBlockHeight, err := getBlockHeight()
	if err != nil {
		return err
	}

	expectedBlockHeight := big.NewInt(0)
	expectedBlockHeight.Add(ctx.lastBlockRegistered, big.NewInt(int64(arg1)))

	if actualBlockHeight.Cmp(expectedBlockHeight) < 0 {
		return fmt.Errorf("block was expected to be bigger or equal than %d, %d instead", expectedBlockHeight, actualBlockHeight)
	}

	return nil
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

func (ctx *WalletBackendContext) IWaitForSeconds(arg1 int) error {
	time.Sleep(time.Second * time.Duration(arg1))

	return nil
}