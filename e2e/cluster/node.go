package cluster

import (
	"fmt"
	"strings"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/log"
)

type NodeID string

func (id NodeID) String() string {
	return string(id)
}

type NodeSpec struct {
	ID          NodeID
	Image       string
	Files       map[string][]byte
	PortMapping map[int32]int32
	Cmd         []string
	IsReadyFn   func(runner NodeRunner) error
}

func BootnodeSpec(nodeSuffix string) (*NodeSpec, error) {
	id := NodeID("bootnode-" + nodeSuffix)
	spec := &NodeSpec{
		ID:    id,
		Image: "kowalatech/bootnode:dev",
		Cmd: []string{
			"--nodekeyhex", randStringBytes(64),
			"--v5",
		},
		Files: map[string][]byte{},
	}
	return spec, nil
}

func WalletBackendSpec(nodeSuffix string) (*NodeSpec, error) {
	id := NodeID("wallet-backend-" + nodeSuffix)
	spec := &NodeSpec{
		ID:    id,
		Image: "kowalatech/wallet_backend:dev",
		Cmd:   []string{},
		PortMapping: map[int32]int32{
			8080 : 8080,
		},
		Files: map[string][]byte{},
	}
	return spec, nil
}

func kcoinIsReadyFn(nodeID NodeID) func(NodeRunner) error {
	return func(runner NodeRunner) error {
		randomStr := randStringBytes(64)
		res, err := runner.Exec(nodeID, KcoinExecCommand(fmt.Sprintf(`console.log("%v");`, randomStr)))
		if err != nil {
			log.Warn("node is not ready yet", "err", err)
			return common.ErrConditionNotMet
		}

		if !strings.Contains(res.StdOut, randomStr) {
			return fmt.Errorf("node returns a wrong result. expect %s, got %s", randomStr, res.StdOut)
		}

		return nil
	}
}
