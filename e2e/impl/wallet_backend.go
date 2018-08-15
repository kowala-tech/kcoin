package impl

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"time"

	"github.com/kowala-tech/kcoin/client/accounts"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/rlp"
	"github.com/kowala-tech/kcoin/e2e/cluster"
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

func (ctx *WalletBackendContext) runRedis() (redisAddr string, err error) {
	redisSpec, err := cluster.RedisSpec(ctx.globalCtx.nodeSuffix)
	if err != nil {
		return "", err
	}
	if err := ctx.globalCtx.nodeRunner.Run(redisSpec); err != nil {
		return "", err
	}
	redisIP, err := ctx.globalCtx.nodeRunner.IP(redisSpec.ID)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v:6379", redisIP), nil
}

func (ctx *WalletBackendContext) runNsqlookupd() (nsqlookupdAddr string, err error) {
	nsqlookupdSpec, err := cluster.NsqlookupdSpec(ctx.globalCtx.nodeSuffix)
	if err != nil {
		return "", err
	}
	if err := ctx.globalCtx.nodeRunner.Run(nsqlookupdSpec); err != nil {
		return "", err
	}

	nsqlookupdIP, err := ctx.globalCtx.nodeRunner.IP(nsqlookupdSpec.ID)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v:4160", nsqlookupdIP), nil
}

func (ctx *WalletBackendContext) runNsqd(nsqlookupdAddr string) (nsqdAddr string, err error) {
	nsqdSpec, err := cluster.NsqdSpec(ctx.globalCtx.nodeSuffix, nsqlookupdAddr)
	if err != nil {
		return "", nil
	}
	if err := ctx.globalCtx.nodeRunner.Run(nsqdSpec); err != nil {
		return "", err
	}
	nsqdIP, err := ctx.globalCtx.nodeRunner.IP(nsqdSpec.ID)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v:4150", nsqdIP), nil
}

func (ctx *WalletBackendContext) runTransactionsPersistance(nsqdAddr, redisAddr string) error {
	transactionsPersistanceSpec, err := cluster.TransactionsPersistanceSpec(ctx.globalCtx.nodeSuffix, nsqdAddr, redisAddr)
	if err != nil {
		return err
	}
	if err := ctx.globalCtx.nodeRunner.Run(transactionsPersistanceSpec); err != nil {
		return err
	}

	return nil
}

func (ctx *WalletBackendContext) runNotificationsApi(redisAddr string) (notificationsApiAddr string, err error) {
	notificationsApiSpec, err := cluster.NotificationsApiSpec(ctx.globalCtx.nodeSuffix, redisAddr)
	if err != nil {
		return "", err
	}
	if err := ctx.globalCtx.nodeRunner.Run(notificationsApiSpec); err != nil {
		return "", err
	}
	notificationsApiIP, err := ctx.globalCtx.nodeRunner.IP(notificationsApiSpec.ID)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v:%v", notificationsApiIP, "3000"), nil
}

func (ctx *WalletBackendContext) runTransactionsPublisher(nsqdAddr, redisAddr, rpcAddr string) error {
	txPublisherSpec, err := cluster.TransactionsPublisherSpec(ctx.globalCtx.nodeSuffix, nsqdAddr, redisAddr, rpcAddr)
	if err != nil {
		return err
	}
	if err := ctx.globalCtx.nodeRunner.Run(txPublisherSpec); err != nil {
		return err
	}

	return nil
}

func (ctx *WalletBackendContext) runWalletBackend(rpcAddr, notificationsApiAddr string) error {
	spec, err := cluster.WalletBackendSpec(ctx.globalCtx.nodeSuffix, rpcAddr, notificationsApiAddr)
	if err != nil {
		return err
	}
	if err := ctx.globalCtx.nodeRunner.Run(spec); err != nil {
		return err
	}

	return nil
}

func (ctx *WalletBackendContext) getRpcAddr() (rpcAddr string, err error) {
	rpcIP, err := ctx.globalCtx.nodeRunner.IP(ctx.globalCtx.rpcNodeID)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("http://%v:%v", rpcIP, ctx.globalCtx.rpcPort), nil
}

func (ctx *WalletBackendContext) TheWalletBackendNodeIsRunning() error {
	rpcAddr, err := ctx.getRpcAddr()
	if err != nil {
		return err
	}

	redisAddr, err := ctx.runRedis()
	if err != nil {
		return err
	}

	nsqlookupdAddr, err := ctx.runNsqlookupd()
	if err != nil {
		return nil
	}

	nsqdAddr, err := ctx.runNsqd(nsqlookupdAddr)
	if err != nil {
		return err
	}

	err = ctx.runTransactionsPersistance(nsqdAddr, redisAddr)
	if err != nil {
		return err
	}

	notificationsApiAddr, err := ctx.runNotificationsApi(redisAddr)
	if err != nil {
		return err
	}

	err = ctx.runTransactionsPublisher(nsqdAddr, redisAddr, rpcAddr)
	if err != nil {
		return err
	}

	err = ctx.runWalletBackend(rpcAddr, notificationsApiAddr)
	if err != nil {
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

func (ctx *WalletBackendContext) ITransferKcoin(kcoin int64, from, to string) error {
	fromAccount, ok := ctx.globalCtx.accounts[from]
	if !ok {
		return fmt.Errorf("can't get account for %q", from)
	}

	toAccount, ok := ctx.globalCtx.accounts[to]
	if !ok {
		return fmt.Errorf("can't get account for %q", to)
	}

	tx, err := ctx.globalCtx.buildTx(fromAccount, toAccount, kcoin)
	if err != nil {
		return err
	}

	rawTx, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return err
	}

	httpResp, err := http.Get(
		fmt.Sprintf("http://%s:8080/api/broadcasttx/%x", ctx.globalCtx.nodeRunner.HostIP(), rawTx),
	)
	if err != nil {
		return fmt.Errorf("error sending signed transaction. %s", err)
	}

	rawResp, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return err
	}

	var resp map[string]interface{}

	err = json.Unmarshal(rawResp, &resp)
	if err != nil {
		return err
	}

	return nil
}

func (ctx *WalletBackendContext) TheTransactionsOfAccountShouldContainLastTransaction(acc string) error {
	lastTx := ctx.globalCtx.lastTx
	account, ok := ctx.globalCtx.accounts[acc]
	if !ok {
		return fmt.Errorf("can't get account for %q", acc)
	}

	return common.WaitFor("find the transaction in the list", time.Second, time.Second*20, func() error {
		transactions, err := ctx.getTransactionsHashes(account)
		if err != nil {
			return err
		}
		for _, hash := range transactions {
			if hash == lastTx.Hash().String() {
				return nil
			}
		}
		return errors.New("transaction not found")
	})
}

func (ctx *WalletBackendContext) TheBalanceIsAround(accountName string, expectedKcoin int64) error {
	return ctx.theBalanceIs(accountName, "around", expectedKcoin)
}

func (ctx *WalletBackendContext) TheBalanceIsGreater(accountName string, expectedKcoin int64) error {
	return ctx.theBalanceIs(accountName, "greater", expectedKcoin)
}

func (ctx *WalletBackendContext) TheBalanceIsExactly(accountName string, expectedKcoin int64) error {
	return ctx.theBalanceIs(accountName, "equal", expectedKcoin)
}

func (ctx *WalletBackendContext) theBalanceIs(accountName string, cmp string, expectedKcoin int64) error {
	return common.WaitFor("the balance satisfies a condition", time.Second, 10*time.Second, func() error {
		expected := toWei(expectedKcoin)
		account := ctx.globalCtx.accounts[accountName]

		res, err := http.Get(
			fmt.Sprintf("http://%s:8080/api/balance/%s", ctx.globalCtx.nodeRunner.HostIP(), account.Address.String()),
		)
		if err != nil {
			return fmt.Errorf("error connecting to wallet backend to get balance. %s", err)
		}

		rawResp, err := ioutil.ReadAll(res.Body)
		defer res.Body.Close()
		if err != nil {
			return fmt.Errorf("error parsing response from wallet backend to get block height. %s", err)
		}

		var resp struct {
			Balance *big.Int `json:"balance"`
		}

		err = json.Unmarshal(rawResp, &resp)
		if err != nil {
			return fmt.Errorf("error unmarshalling from json response. %s", err)
		}

		cmpFunc, err := newCompare(cmp)
		if err != nil {
			return err
		}
		if !cmpFunc(expected, resp.Balance) {
			return fmt.Errorf("balance expected to be %s %v but is %v", cmp, expected, resp.Balance)
		}

		return nil
	})
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

	var resp struct {
		BlockHeight *big.Int `json:"block_height"`
	}

	err = json.Unmarshal(rawResp, &resp)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling from json response. %s", err)
	}

	return resp.BlockHeight, nil
}

func (ctx *WalletBackendContext) getTransactionsHashes(acc accounts.Account) ([]string, error) {
	url := fmt.Sprintf("http://%s:8080/api/transactions/%s", ctx.globalCtx.nodeRunner.HostIP(), acc.Address.String())
	httpRes, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error getting transactions. %s", err)
	}

	rawResp, err := ioutil.ReadAll(httpRes.Body)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Transactions []struct {
			Hash string `json:"hash"`
		} `json:"transactions"`
	}
	err = json.Unmarshal(rawResp, &resp)
	if err != nil {
		return nil, err
	}

	hashes := make([]string, len(resp.Transactions))
	for i, tx := range resp.Transactions {
		hashes[i] = tx.Hash
	}

	return hashes, nil
}
