package impl

import (
	"io/ioutil"
	"math/big"

	"sync/atomic"

	"github.com/kowala-tech/kcoin/client/accounts"
	"github.com/kowala-tech/kcoin/client/accounts/keystore"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core/types"
	"github.com/kowala-tech/kcoin/client/kcoinclient"
	"github.com/kowala-tech/kcoin/e2e/cluster"
)

type Context struct {
	Name            string
	AccountsStorage *keystore.KeyStore

	// cluster config
	genesis  []byte
	bootnode string

	nodeRunner             cluster.NodeRunner
	genesisValidatorNodeID cluster.NodeID
	rpcNodeID              cluster.NodeID
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

	scenarioNumber *int32
	nodeSuffix     string

	waiter doer
}

func NewTestContext(chainID *big.Int) *Context {
	tmpdir, _ := ioutil.TempDir("", "eth-keystore-test")
	accountsStorage := keystore.NewKeyStore(tmpdir, 2, 1)

	ctx := &Context{
		AccountsStorage: accountsStorage,
		chainID:         chainID,

		accounts:       make(map[string]accounts.Account),
		scenarioNumber: new(int32),
		nodeSuffix:     common.RandomString(4),
	}

	ctx.waiter = common.NewWaiter(ctx)

	return ctx
}

func (ctx *Context) GetScenarioNumber() int32 {
	return atomic.LoadInt32(ctx.scenarioNumber)
}

func (ctx *Context) IncreaseScenarioNumber() int32 {
	return atomic.AddInt32(ctx.scenarioNumber, 1)
}

func (ctx *Context) Reset() {
	ctx.accounts = make(map[string]accounts.Account)

	ctx.IncreaseScenarioNumber()
}

type doer interface {
	Do(execFunc func() error, condFunc ...func() error) error
}
