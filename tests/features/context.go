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
	cluster              cluster.Cluster
	client               *kcoinclient.Client
	genesisValidatorName string
	chainID              *big.Int

	accountsStorage *keystore.KeyStore

	accounts map[string]accounts.Account

	lastTx    *types.Transaction
	lastTxErr error
}

func NewTestContext(k8sCluster cluster.Cluster, genesisValidatorName string, client *kcoinclient.Client, chainID *big.Int) *Context {
	tmpdir, _ := ioutil.TempDir("", "eth-keystore-test")
	accountsStorage := keystore.NewKeyStore(tmpdir, 2, 1)

	return &Context{
		cluster:              k8sCluster,
		client:               client,
		genesisValidatorName: genesisValidatorName,

		accountsStorage: accountsStorage,

		accounts: make(map[string]accounts.Account),
	}
}
