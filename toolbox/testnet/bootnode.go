package testnet

import (
	"fmt"
	"regexp"

	"time"

	"math/rand"
	"strconv"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/pkg/errors"
)

var (
	enodeSecretRegexp = regexp.MustCompile(`enode://([a-f0-9]*)@`)
)

//BootNode represents a bootnode node running in a docker container.
type BootNode interface {
	Node
	Enode() string
}

//NewBootNode returns a bootnode running in a docker container.
func NewBootNode(dockerEngine DockerEngine, networkID string) (BootNode, error) {
	nodeSpecBuilder := NodeSpecBuilder{}

	nodeSpec, err := nodeSpecBuilder.
		AsBootnode().
		WithNetworkID(networkID).
		WithImage("kowalatech/bootnode").
		WithID("bootnode-" + strconv.Itoa(rand.Int())).
		Build()

	if err != nil {
		return nil, err
	}

	return &bootNode{
		node: node{
			nodeSpec:     nodeSpec,
			dockerEngine: dockerEngine,
		},
	}, nil
}

type bootNode struct {
	node
	enode string
}

//Enode returns the Enode address of the given bootnode.
func (g *bootNode) Enode() string {
	return g.enode
}

//Start starts the node as bootnode.
func (g *bootNode) Start() error {
	err := g.node.Start()
	if err != nil {
		return nil
	}

	err = g.setEnode()
	if err != nil {
		return err
	}

	return nil
}

func (g *bootNode) setEnode() error {
	err := common.WaitFor("error getting enode", 1*time.Second, 10*time.Second, func() error {
		bootnodeStdout, err := g.dockerEngine.GetLogs(g.nodeSpec.ID)
		if err != nil {
			return err
		}

		found := enodeSecretRegexp.FindStringSubmatch(bootnodeStdout)
		if len(found) != 2 {
			return errors.New("enode address not found")
		}
		enodeSecret := found[1]
		bootnodeInfo, err := g.dockerEngine.ContainerInspect(g.nodeSpec.ID)
		if err != nil {
			return err
		}

		bootnodeIP := bootnodeInfo.NetworkSettings.Networks[g.nodeSpec.NetworkID].IPAddress

		g.enode = fmt.Sprintf("enode://%v@%v:33445", enodeSecret, bootnodeIP)

		return nil
	})

	return err
}
