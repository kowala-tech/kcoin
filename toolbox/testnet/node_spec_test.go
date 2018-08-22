package testnet

import (
	"testing"

	"math/big"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/stretchr/testify/assert"
)

var DefaultNode = &NodeSpec{
	ID:    "bootnode",
	Image: "theimage/image",
	Files: make(map[string][]byte, 0),
}

func TestNodeSpecBuilder(t *testing.T) {
	t.Run("Builder without name returns error", func(t *testing.T) {
		builder := NodeSpecBuilder{}

		builder.WithImage("theimage/image")

		_, err := builder.Build()
		assert.Error(t, err)
	})

	t.Run("Builder without image returns error", func(t *testing.T) {
		builder := NodeSpecBuilder{}

		builder.WithID("bootnode")

		_, err := builder.Build()
		assert.Error(t, err)
	})

	t.Run("We can create a node with name and image", func(t *testing.T) {
		builder := getBuilderWithDefaults()

		nodeSpec := buildNodeSpecWithoutError(t, builder)

		assert.Equal(t, DefaultNode, nodeSpec)
	})

	t.Run("We can set the network that the node will connect", func(t *testing.T) {
		builder := getBuilderWithDefaults()

		DefaultNode.NetworkID = "integration"
		builder.WithNetworkID("integration")

		nodeSpec := buildNodeSpecWithoutError(t, builder)

		assert.Equal(t, DefaultNode, nodeSpec)
	})

	t.Run("We can set a node as bootnode", func(t *testing.T) {
		builder := getBuilderWithDefaults()

		builder.AsBootnode()

		nodeSpec := buildNodeSpecWithoutError(t, builder)

		assert.Contains(t, nodeSpec.Cmd, "--nodekeyhex")
		assert.Len(t, nodeSpec.Cmd, 2)
	})

	t.Run("We can set a node with a bootnode address to connect", func(t *testing.T) {
		builder := getBuilderWithDefaults()

		bootnodeAddr := "enode://testnode"
		builder.WithBootnode(bootnodeAddr)

		nodeSpec := buildNodeSpecWithoutError(t, builder)

		assert.Contains(t, nodeSpec.Cmd, "--bootnodes")
		assert.Contains(t, nodeSpec.Cmd, bootnodeAddr)

		assert.Len(t, nodeSpec.Cmd, 2)
	})

	t.Run("We can specify the type of sync mode", func(t *testing.T) {
		builder := getBuilderWithDefaults()
		builder.WithSyncMode("full")

		nodeSpec := buildNodeSpecWithoutError(t, builder)

		assert.Contains(t, nodeSpec.Cmd, "--syncmode")
		assert.Contains(t, nodeSpec.Cmd, "full")

		assert.Len(t, nodeSpec.Cmd, 2)
	})

	t.Run("We can specify the gas price.", func(t *testing.T) {
		builder := getBuilderWithDefaults()
		builder.WithGasPrice("1")

		nodeSpec := buildNodeSpecWithoutError(t, builder)

		assert.Contains(t, nodeSpec.Cmd, "--gasprice")
		assert.Contains(t, nodeSpec.Cmd, "1")

		assert.Len(t, nodeSpec.Cmd, 2)
	})

	t.Run("We can set verbosity", func(t *testing.T) {
		builder := getBuilderWithDefaults()
		builder.WithVerbosity("3")

		nodeSpec := buildNodeSpecWithoutError(t, builder)

		assert.Contains(t, nodeSpec.Cmd, "--verbosity")
		assert.Contains(t, nodeSpec.Cmd, "3")

		assert.Len(t, nodeSpec.Cmd, 2)
	})

	t.Run("We can set Chain id", func(t *testing.T) {
		builder := getBuilderWithDefaults()
		builder.WithChainID("123456")

		nodeSpec := buildNodeSpecWithoutError(t, builder)

		assert.Contains(t, nodeSpec.Cmd, "--networkid")
		assert.Contains(t, nodeSpec.Cmd, "123456")

		assert.Len(t, nodeSpec.Cmd, 2)
	})

	t.Run("We can set a genesis block file", func(t *testing.T) {
		builder := getBuilderWithDefaults()
		builder.WithGenesis([]byte("theContent"))

		nodeSpec := buildNodeSpecWithoutError(t, builder)

		expectedFiles := make(map[string][]byte, 0)
		expectedFiles["/kcoin/genesis.json"] = []byte("theContent")

		assert.Contains(t, nodeSpec.Cmd, "--genesis-path=/kcoin/genesis.json")

		assert.Equal(
			t,
			nodeSpec.Files,
			expectedFiles,
		)
	})

	t.Run("We can specify a node as validator", func(t *testing.T) {
		builder := getBuilderWithDefaults()
		builder.AsValidator()

		nodeSpec := buildNodeSpecWithoutError(t, builder)

		assert.Contains(t, nodeSpec.Cmd, "--validate")
	})

	t.Run("With coinbase address", func(t *testing.T) {
		address := common.HexToAddress("0x09438E46Ea66647EA65E4b104C125c82076FDcE5")

		builder := getBuilderWithDefaults()
		builder.WithAccount(address, []byte("rawAccount"))

		nodeSpec := buildNodeSpecWithoutError(t, builder)

		assert.Contains(t, nodeSpec.Cmd, "--password")
		assert.Contains(t, nodeSpec.Cmd, "/kcoin/password.txt")
		assert.Contains(t, nodeSpec.Cmd, "--unlock")
		assert.Contains(t, nodeSpec.Cmd, address.Hex())

		expectedFiles := make(map[string][]byte, 0)
		expectedFiles["/kcoin/password.txt"] = []byte("test")
		expectedFiles["/root/.kcoin/kusd/keystore/0.json"] = []byte("rawAccount")

		assert.Equal(
			t,
			nodeSpec.Files,
			expectedFiles,
		)
	})

	t.Run("We can set the deposit", func(t *testing.T) {
		deposit := big.NewInt(1)

		builder := getBuilderWithDefaults()
		builder.WithDeposit(deposit)

		nodeSpec := buildNodeSpecWithoutError(t, builder)

		assert.Contains(t, nodeSpec.Cmd, "--deposit")
		assert.Contains(t, nodeSpec.Cmd, "1")
	})

	t.Run("We can activate rpc", func(t *testing.T) {
		builder := getBuilderWithDefaults()
		builder.WithRPCAccess()

		nodeSpec := buildNodeSpecWithoutError(t, builder)

		assert.Contains(t, nodeSpec.Cmd, "--rpc")
		assert.Contains(t, nodeSpec.Cmd, "--rpcaddr")
		assert.Contains(t, nodeSpec.Cmd, "0.0.0.0")
		assert.Contains(t, nodeSpec.Cmd, "--rpccorsdomain")
		assert.Contains(t, nodeSpec.Cmd, "*")

		assert.Equal(t, nodeSpec.PortMapping[22334], int32(22334))
	})
}

//buildNodeSpecWithoutError builds a nodeSpec and checks it does not return error
func buildNodeSpecWithoutError(t *testing.T, builder *NodeSpecBuilder) *NodeSpec {
	nodeSpec, err := builder.Build()
	assert.NoError(t, err)

	return nodeSpec
}

func getBuilderWithDefaults() *NodeSpecBuilder {
	builder := &NodeSpecBuilder{
		portMapping: make(map[int32]int32),
	}

	builder.
		WithID(DefaultNode.ID).
		WithImage(DefaultNode.Image)

	return builder
}
