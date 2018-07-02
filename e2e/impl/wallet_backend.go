package impl

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"time"

	"github.com/DATA-DOG/godog"
	"github.com/kowala-tech/kcoin/client/accounts"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/e2e/cluster"
	"github.com/kowala-tech/kcoin/wallet-backend/application/command"
)

type WalletBackendContext struct {
	globalCtx           *Context
	nodeRunning         bool
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
	rpcIP, err := ctx.globalCtx.nodeRunner.IP(ctx.globalCtx.rpcNodeID)
	if err != nil {
		return err
	}
	rpcAddr := fmt.Sprintf("http://%v:%v", rpcIP, ctx.globalCtx.rpcPort)

	redisSpec, err := cluster.RedisSpec(ctx.globalCtx.nodeSuffix)
	if err != nil {
		return err
	}
	if err := ctx.globalCtx.nodeRunner.Run(redisSpec, ctx.globalCtx.GetScenarioNumber()); err != nil {
		return err
	}
	redisIP, err := ctx.globalCtx.nodeRunner.IP(redisSpec.ID)
	if err != nil {
		return err
	}
	redisAddr := fmt.Sprintf("%v:6379", redisIP)

	transactionsPersistanceSpec, err := cluster.TransactionsPersistanceSpec(ctx.globalCtx.nodeSuffix, rpcAddr, redisAddr)
	if err != nil {
		return err
	}
	if err := ctx.globalCtx.nodeRunner.Run(transactionsPersistanceSpec, ctx.globalCtx.GetScenarioNumber()); err != nil {
		return err
	}

	notificationsApiSpec, err := cluster.NotificationsApiSpec(ctx.globalCtx.nodeSuffix, redisAddr)
	if err != nil {
		return err
	}
	if err := ctx.globalCtx.nodeRunner.Run(notificationsApiSpec, ctx.globalCtx.GetScenarioNumber()); err != nil {
		return err
	}
	notificationsApiIP, err := ctx.globalCtx.nodeRunner.IP(notificationsApiSpec.ID)
	if err != nil {
		return err
	}
	notificationsApiAddr := fmt.Sprintf("%v:%v", notificationsApiIP, "3000")

	fmt.Printf("wallet_backend.go:68 ===> %[2]v: %[1]v\n", rpcAddr, `rpcAddr`)
	spec, err := cluster.WalletBackendSpec(ctx.globalCtx.nodeSuffix, rpcAddr, notificationsApiAddr)
	if err != nil {
		return err
	}

	if err := ctx.globalCtx.nodeRunner.Run(spec, ctx.globalCtx.GetScenarioNumber()); err != nil {
		return err
	}

	// Wait for some data to be meaningful
	err = common.WaitFor("wallet backend syncs up with RPC", time.Second, time.Second*20, func() error {
		block, err := ctx.getBlockHeight()
		if err != nil {
			return err
		}
		if block.Cmp(big.NewInt(0)) == 0 {
			return errors.New("Block height is still 0")
		}
		return nil
	})
	if err != nil {
		return err
	}

	ctx.nodeRunning = true

	return nil
}

func (ctx *WalletBackendContext) ICheckTheCurrentBlockHeightInTheWalletBackendAPI() error {
	blockHeight, err := ctx.getBlockHeight()
	if err != nil {
		return err
	}

	ctx.lastBlockRegistered = blockHeight

	return nil
}

func (ctx *WalletBackendContext) IWaitForBlocks(blocks int) error {
	baseHeight, err := ctx.getBlockHeight()
	if err != nil {
		return err
	}
	return common.WaitFor("wait for some blocks", time.Second, time.Second*20, func() error {
		height, err := ctx.getBlockHeight()
		if err != nil {
			return err
		}

		diff := new(big.Int).Sub(height, baseHeight)

		if diff.Cmp(big.NewInt(int64(blocks))) < 0 {
			return fmt.Errorf("block difference is %v, expected %v", diff.Int64, blocks)
		}
		return nil
	})
}

func (ctx *WalletBackendContext) TheNewBlockHeightInTheWalletBackendAPIHasIncreasedByAtLeast(arg1 int) error {
	actualBlockHeight, err := ctx.getBlockHeight()
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

func (ctx *WalletBackendContext) TheTransactionsOfAccountShouldContainLastTransaction(acc string) error {
	lastTx := ctx.globalCtx.lastTx
	account, ok := ctx.globalCtx.accounts[acc]
	if !ok {
		return fmt.Errorf("can't get account for %q", acc)
	}

	return common.WaitFor("find the transaction in the list", time.Second, time.Second*20, func() error {
		transactions, err := ctx.getTransactions(account)
		if err != nil {
			return err
		}
		for _, tx := range transactions.Transactions {
			if tx.Hash == lastTx.Hash().String() {
				return nil
			}
		}
		return errors.New("transaction not found")
	})
}

func (ctx *WalletBackendContext) TheBalanceUsingTheWalletBackendShouldBeAroundKcoins(account string, arg1 int) error {
	return godog.ErrPending
}

func (ctx *WalletBackendContext) TheBalanceUsingTheWalletBackendShouldBeKcoins(account string, arg1 int) error {
	return godog.ErrPending
}

func (ctx *WalletBackendContext) getBlockHeight() (*big.Int, error) {
	res, err := http.Get(
		fmt.Sprintf("http://%s:8080/api/blockheight", ctx.globalCtx.nodeRunner.HostIP()),
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

func (ctx *WalletBackendContext) getTransactions(acc accounts.Account) (*command.TransactionsResponse, error) {
	url := fmt.Sprintf("http://%s:8080/api/transactions/%s", ctx.globalCtx.nodeRunner.HostIP(), acc.Address.String())
	fmt.Printf("wallet_backend.go:204 ===> %[2]v: %[1]v\n", url, `url`)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error getting transactions. %s", err)
	}

	rawResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var txResp command.TransactionsResponse
	err = json.Unmarshal(rawResp, &txResp)
	if err != nil {
		return nil, err
	}
	fmt.Printf("wallet_backend.go:219 ===> %[2]v: %[1]v\n", txResp, `txResp`)

	return &txResp, nil
}
