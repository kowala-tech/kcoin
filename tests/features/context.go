package features

import (
	"github.com/kowala-tech/kcoin/cluster"
	"github.com/kowala-tech/kcoin/kcoinclient"
)

type Context struct {
	cluster              cluster.Cluster
	client               *kcoinclient.Client
	genesisValidatorName string

	accountsNodeNames map[string]string
	accountsCoinbase  map[string]string
}

func NewTestContext(k8sCluster cluster.Cluster, genesisValidatorName string, client *kcoinclient.Client) *Context {
	return &Context{
		cluster:              k8sCluster,
		client:               client,
		genesisValidatorName: genesisValidatorName,

		accountsNodeNames: make(map[string]string),
		accountsCoinbase:  make(map[string]string),
	}
}
