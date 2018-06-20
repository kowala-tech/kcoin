package features

import (
	"io/ioutil"
	"math/big"

	"github.com/kowala-tech/kcoin/client/accounts"
	"github.com/kowala-tech/kcoin/client/accounts/keystore"
	"github.com/kowala-tech/kcoin/client/cluster"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core/types"
	"github.com/kowala-tech/kcoin/client/kcoinclient"
)

type Context struct {
	Name            string
	AccountsStorage *keystore.KeyStore

	// cluster config
	genesis  []byte
	bootnode string

	nodeRunner             cluster.NodeRunner
	genesisValidatorNodeID cluster.NodeID
	rpcPort                int32
	client                 *kcoinclient.Client
	chainID                *big.Int

	genesisValidatorAccount accounts.Account
	seederAccount           accounts.Account
	accounts                map[string]accounts.Account

	lastTx              *types.Transaction
	lastTxErr           error
	lastTxStartingBlock *big.Int

	lastUnlockErr error

	scenarioNumber int
	nodeSuffix     string

	waiter doer
}

func NewTestContext(chainID *big.Int) *Context {
	tmpdir, _ := ioutil.TempDir("", "eth-keystore-test")
	accountsStorage := keystore.NewKeyStore(tmpdir, 2, 1)

	ctx := &Context{
		AccountsStorage: accountsStorage,
		chainID:         chainID,

		accounts:   make(map[string]accounts.Account),
		nodeSuffix: common.RandomString(4),
	}

	ctx.waiter = common.NewWaiter(ctx)

	return ctx
}

func (ctx *Context) Reset() {
	ctx.accounts = make(map[string]accounts.Account)
	ctx.nodeRunner.StopAll()

	ctx.scenarioNumber++
	ctx.runNodes()
}

type doer interface {
	Do(execFunc func() error, condFunc ...func() error) error
}
