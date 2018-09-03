package testnet

import (
	"math/big"

	"time"

	"math/rand"
	"strconv"

	"github.com/kowala-tech/kcoin/client/common"
)

var (
	chainID = big.NewInt(2)
)

//GenesisValidator represents a genesis validator node running in a docker container.
type GenesisValidator interface {
	Node
}

//NewGenesisValidator returns a node running in a docker container.
func NewGenesisValidator(dockerEngine DockerEngine, networkID string, bootNode string, genesis []byte, coinbase common.Address, rawCoinbase []byte) (GenesisValidator, error) {
	nodeSpecBuilder := NodeSpecBuilder{
		portMapping: make(map[int32]int32),
	}

	nodeSpec, err := nodeSpecBuilder.
		AsValidator().
		WithRPCAccess().
		WithID("validator-"+strconv.Itoa(rand.Int())).
		WithImage("kowalatech/kusd:dev").
		WithChainID(chainID.String()).
		WithNetworkID(networkID).
		WithBootnode(bootNode).
		WithGenesis(genesis).
		WithSyncMode("full").
		WithGasPrice("1").
		WithAccount(coinbase, rawCoinbase).
		WithDeposit(big.NewInt(1)).
		WithVerbosity("4").
		WithStats("genesis-validator:DuHLdsKV6cagynn9@zygote.kowala.tech").
		Build()

	if err != nil {
		return nil, err
	}

	return &genesisValidator{
		node: node{
			nodeSpec:     nodeSpec,
			dockerEngine: dockerEngine,
		},
	}, nil
}

type genesisValidator struct {
	node
}

//Start starts a node as a genesis validator.
func (g *genesisValidator) Start() error {
	err := g.node.Start()
	if err != nil {
		return err
	}

	time.Sleep(10 * time.Second)

	return nil
}
