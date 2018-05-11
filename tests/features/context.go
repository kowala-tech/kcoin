package features

import (
	"io/ioutil"
	"math/big"

	"github.com/kowala-tech/kcoin/accounts"
	"github.com/kowala-tech/kcoin/accounts/keystore"
	"github.com/kowala-tech/kcoin/cluster"
	"github.com/kowala-tech/kcoin/core/types"
	"github.com/kowala-tech/kcoin/kcoinclient"
)

type Context struct {
	AccountsStorage *keystore.KeyStore

	// cluster config
	genesis  []byte
	bootnode string

	nodeRunner             cluster.NodeRunner
	genesisValidatorNodeID cluster.NodeID
	client                 *kcoinclient.Client
	chainID                *big.Int

	genesisValidatorAccount accounts.Account
	seederAccount           accounts.Account
	accounts                map[string]accounts.Account

	lastTx    *types.Transaction
	lastTxErr error

	lastUnlockErr error

	scenarioNumber int
}

func NewTestContext(chainID *big.Int) *Context {
	tmpdir, _ := ioutil.TempDir("", "eth-keystore-test")
	accountsStorage := keystore.NewKeyStore(tmpdir, 2, 1)

	return &Context{
		AccountsStorage: accountsStorage,
		chainID:         chainID,

		accounts: make(map[string]accounts.Account),
	}
}

func (ctx *Context) Reset() {
	ctx.accounts = make(map[string]accounts.Account)
	ctx.nodeRunner.StopAll()

	ctx.scenarioNumber++
	ctx.runNodes()
}
