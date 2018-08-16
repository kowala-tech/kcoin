package cluster

import (
	"fmt"
	"strconv"

	"github.com/kowala-tech/kcoin/client/accounts"
	"github.com/kowala-tech/kcoin/client/accounts/keystore"
)

type KcoinNodeBuilder struct {
	image          string
	networkId      string
	genesisContent []byte
	accounts       [][]byte
	id             NodeID
	bootnode       string
	syncMode       string
	logLevel       int16
	validate       bool
	rpcPort        *int32
	coinbase       string
	unlockAccounts string

	err error
}

func NewKcoinNodeBuilder() *KcoinNodeBuilder {
	return &KcoinNodeBuilder{
		image:    "kowalatech/kusd:dev",
		logLevel: 3,
		syncMode: "fast",
		accounts: make([][]byte, 0),
	}
}

func (builder *KcoinNodeBuilder) NodeSpec() *NodeSpec {
	cmd := []string{
		"--testnet",
		"--gasprice", "1",
		"--networkid", builder.networkId,
		"--bootnodes", builder.bootnode,
		"--syncmode", builder.syncMode,
		"--verbosity", strconv.Itoa(int(builder.logLevel)),
	}
	files := make(map[string][]byte, 0)
	portMapping := make(map[int32]int32, 0)

	if builder.validate {
		cmd = append(cmd, "--validate")
	}
	if builder.rpcPort != nil {
		cmd = append(cmd, "--rpc")
		cmd = append(cmd, "--rpcaddr", "0.0.0.0")
		cmd = append(cmd, "--rpccorsdomain", `"*"`)
		cmd = append(cmd, "--rpcport", fmt.Sprintf("%v", *builder.rpcPort))
		cmd = append(cmd, "--rpcvhosts=*")
		portMapping[*builder.rpcPort] = *builder.rpcPort
	}
	if builder.coinbase != "" {
		cmd = append(cmd, "--coinbase", builder.coinbase)
	}
	if builder.unlockAccounts != "" {
		files["/kcoin/password.txt"] = []byte("test")
		cmd = append(cmd, "--password", "/kcoin/password.txt", "--unlock", builder.unlockAccounts)
	}

	if len(builder.genesisContent) > 0 {
		cmd = append(cmd, "--genesis-path=/kcoin/genesis.json")
		files["/kcoin/genesis.json"] = builder.genesisContent
	}

	for i, account := range builder.accounts {
		file := fmt.Sprintf("/root/.kcoin/kusd/keystore/%v.json", i)
		files[file] = account
	}

	spec := &NodeSpec{
		ID:          builder.id,
		Image:       builder.image,
		Cmd:         cmd,
		Env:         []string{},
		Files:       files,
		IsReadyFn:   builder.isReadyFn(),
		PortMapping: portMapping,
	}
	return spec
}
func (builder *KcoinNodeBuilder) isReadyFn() func(NodeRunner) error {
	if builder.rpcPort != nil {
		return rpcIsReadyFn(builder.id, *builder.rpcPort)
	}
	return kcoinIsReadyFn(builder.id)
}

func (builder *KcoinNodeBuilder) WithNetworkId(networkID string) *KcoinNodeBuilder {
	builder.networkId = networkID
	return builder
}

func (builder *KcoinNodeBuilder) WithID(id string) *KcoinNodeBuilder {
	builder.id = NodeID(id)
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

func (builder *KcoinNodeBuilder) WithValidation() *KcoinNodeBuilder {
	builder.validate = true
	return builder
}

func (builder *KcoinNodeBuilder) WithCoinbase(account accounts.Account) *KcoinNodeBuilder {
	builder.coinbase = account.Address.Hex()
	return builder
}

func (builder *KcoinNodeBuilder) WithAccount(ks *keystore.KeyStore, account accounts.Account) *KcoinNodeBuilder {
	raw, err := ks.Export(account, "test", "test")
	if err != nil {
		builder.err = err
		return builder
	}
	builder.accounts = append(builder.accounts, raw)
	if builder.unlockAccounts == "" {
		builder.unlockAccounts = account.Address.Hex()
	} else {
		builder.unlockAccounts = fmt.Sprintf("%s,%s", builder.unlockAccounts, account.Address.Hex())
	}
	return builder
}

func (builder *KcoinNodeBuilder) WithRpc(port int32) *KcoinNodeBuilder {
	builder.rpcPort = &port
	return builder
}
