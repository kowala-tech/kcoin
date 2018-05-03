package cluster

import "strconv"

type KcoinNodeBuilder struct {
	image          string
	networkId      string
	genesisContent []byte
	name           string
	bootnode       string
	syncMode       string
	logLevel       int16
}

func NewKcoinNodeBuilder() *KcoinNodeBuilder {
	return &KcoinNodeBuilder{
		image: "kowalatech/kusd:dev",
	}
}

func (builder *KcoinNodeBuilder) Node() *Node {
	cmd := []string{
		"--gasprice", "1",
		"--networkid", builder.networkId,
		"--bootnodes", builder.bootnode,
		"--syncmode", builder.syncMode,
		"--verbosity", strconv.Itoa(int(builder.logLevel)),
	}

	node := &Node{
		Image: builder.image,
		Name:  builder.name,
		Cmd:   cmd,
		Files: map[string][]byte{
			"/kcoin/genesis.json": builder.genesisContent,
		},
	}
	return node
}

func (builder *KcoinNodeBuilder) WithNetworkId(networkID string) *KcoinNodeBuilder {
	builder.networkId = networkID
	return builder
}

func (builder *KcoinNodeBuilder) WithName(name string) *KcoinNodeBuilder {
	builder.name = name
	return builder
}

func (builder *KcoinNodeBuilder) WithBootnode(bootnode string) *KcoinNodeBuilder {
	builder.bootnode = bootnode
	return builder
}

func (builder *KcoinNodeBuilder) WithSyncMode(syncMode string) *KcoinNodeBuilder {
	builder.syncMode = syncMode
	return builder
}

func (builder *KcoinNodeBuilder) WithLogLevel(logLevel int16) *KcoinNodeBuilder {
	builder.logLevel = logLevel
	return builder
}

func (builder *KcoinNodeBuilder) WithGenesis(genesis []byte) *KcoinNodeBuilder {
	builder.genesisContent = genesis
	return builder
}
