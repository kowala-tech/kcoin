package cluster

import (
	"fmt"
	"strings"
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
	IsReadyFn   func(runner NodeRunner) bool
}

func BootnodeSpec() (*NodeSpec, error) {
	id := NodeID("bootnode")
	spec := &NodeSpec{
		ID:    id,
		Image: "kowalatech/bootnode:dev",
		Cmd: []string{
			"--nodekeyhex", randStringBytes(64),
		},
		Files: map[string][]byte{},
	}
	return spec, nil
}

func kcoinIsReadyFn(nodeID NodeID) func(NodeRunner) bool {
	return func(runner NodeRunner) bool {
		randomStr := randStringBytes(64)
		res, err := runner.Exec(nodeID, KcoinExecCommand(fmt.Sprintf(`console.log("%v");`, randomStr)))
		return err == nil && strings.Contains(res.StdOut, randomStr)
	}
}
