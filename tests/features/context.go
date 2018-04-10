package features

import "github.com/kowala-tech/kcoin/cluster"

type Context struct {
	cluster              cluster.Cluster
	genesisValidatorName string

	accountsNodeNames map[string]string
	accountsCoinbase  map[string]string

	lastTxStdout string
}

func NewTestContext(k8sCluster cluster.Cluster, genesisValidatorName string) *Context {
	return &Context{
		cluster:              k8sCluster,
		genesisValidatorName: genesisValidatorName,

		accountsNodeNames: make(map[string]string),
		accountsCoinbase:  make(map[string]string),
	}
}
