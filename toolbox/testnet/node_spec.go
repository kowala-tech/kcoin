package testnet

import (
	"math/big"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/pkg/errors"
)

//NodeSpec represents an object with several options to configure a Node running in a docker container.
type NodeSpec struct {
	ID          string
	Image       string
	NetworkID   string
	Cmd         []string
	Files       map[string][]byte
	PortMapping map[int32]int32
}

//NodeSpecBuilder is a builder to help creating a NodeSpec object.
type NodeSpecBuilder struct {
	id             string
	image          string
	networkID      string
	cmd            []string
	genesisContent []byte
	coinbase       *common.Address
	rawcoinbase    []byte
	portMapping    map[int32]int32
}

//Build returns a NodeSpec configured by the given options.
func (n *NodeSpecBuilder) Build() (*NodeSpec, error) {
	if n.id == "" {
		return nil, errors.New("node id is needed")
	}

	if n.image == "" {
		return nil, errors.New("docker images is needed")
	}

	files := make(map[string][]byte, 0)

	if len(n.genesisContent) > 0 {
		files["/kcoin/genesis.json"] = n.genesisContent
	}

	if n.coinbase != nil && len(n.rawcoinbase) > 0 {
		files["/kcoin/password.txt"] = []byte("test")
		n.addCmdArgs([]string{"--password", "/kcoin/password.txt", "--unlock", n.coinbase.Hex()})
		files["/root/.kcoin/kusd/keystore/0.json"] = n.rawcoinbase
	}

	return &NodeSpec{
		ID:        n.id,
		Image:     n.image,
		NetworkID: n.networkID,
		Cmd:       n.cmd,
		Files:     files,
		PortMapping: n.portMapping,
	}, nil
}

//WithID specifies the id the container node will have as a docker container.
func (n *NodeSpecBuilder) WithID(nodeID string) *NodeSpecBuilder {
	n.id = nodeID
	return n
}

//WithImage sets the docker image that will be used to launch the node.
func (n *NodeSpecBuilder) WithImage(image string) *NodeSpecBuilder {
	n.image = image
	return n
}

//WithNetworkID sets the docker container in the given network.
func (n *NodeSpecBuilder) WithNetworkID(networkID string) *NodeSpecBuilder {
	n.networkID = networkID
	return n
}

//WithSyncMode sets the sync mode to specified
func (n *NodeSpecBuilder) WithSyncMode(syncMode string) *NodeSpecBuilder {
	n.addCmdArgs([]string{"--syncmode", syncMode})
	return n
}

//AsBootnode sets a node as a bootnode.
func (n *NodeSpecBuilder) AsBootnode() *NodeSpecBuilder {
	n.addCmdArgs([]string{"--nodekeyhex", randStringBytes(64)})
	return n
}

//WithBootnode uses bootnode address as the bootnode to connect.
func (n *NodeSpecBuilder) WithBootnode(bootnodeAddr string) *NodeSpecBuilder {
	n.addCmdArgs([]string{"--bootnodes", bootnodeAddr})
	return n
}

//WithGasPrice sets the price of the given node.
func (n *NodeSpecBuilder) WithGasPrice(gasPrice string) *NodeSpecBuilder {
	n.addCmdArgs([]string{"--gasprice", gasPrice})
	return n
}

//WithVerbosity sets the level of verbosity that the node will have.
func (n *NodeSpecBuilder) WithVerbosity(level string) *NodeSpecBuilder {
	n.addCmdArgs([]string{"--verbosity", level})
	return n
}

//WithChainID sets the chain id that the knode will run.
func (n *NodeSpecBuilder) WithChainID(chainID string) *NodeSpecBuilder {
	n.addCmdArgs([]string{"--networkid", chainID})
	return n
}

//AsValidator sets the node as a validator of blocks.
func (n *NodeSpecBuilder) AsValidator() *NodeSpecBuilder {
	n.addCmdArgs([]string{"--validate"})
	return n
}

//WithDeposit sets the deposit the node will have.
func (n *NodeSpecBuilder) WithDeposit(deposit *big.Int) *NodeSpecBuilder {
	n.addCmdArgs([]string{"--deposit", deposit.String()})
	return n
}

//WithRPCAccess activates rpc access to the given node.
func (n *NodeSpecBuilder) WithRPCAccess() *NodeSpecBuilder {
	n.addCmdArgs([]string{"--rpc"})
	n.addCmdArgs([]string{"--rpcaddr", "0.0.0.0"})
	n.addCmdArgs([]string{"--rpccorsdomain=*"})
	n.addCmdArgs([]string{"--rpcport", "30503"})
	n.addCmdArgs([]string{"--rpcvhosts=*"})

	n.portMapping[30503] = 30503

	return n
}

func (n *NodeSpecBuilder) addCmdArgs(args []string) {
	n.cmd = append(n.cmd, args...)
}

//WithGenesis sets the content of the genesis.json file for the given node.
func (n *NodeSpecBuilder) WithGenesis(content []byte) *NodeSpecBuilder {
	n.addCmdArgs([]string{"--genesis-path=/kcoin/genesis.json"})
	n.genesisContent = content
	return n
}

//WithAccount sets the given account as the coinbase of the node.
func (n *NodeSpecBuilder) WithAccount(address common.Address, rawAccount []byte) *NodeSpecBuilder {
	n.coinbase = &address
	n.rawcoinbase = rawAccount

	return n
}
