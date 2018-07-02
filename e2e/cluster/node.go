package cluster

import (
	"fmt"
	"net/http"
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
	Env         []string
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
		Env:   []string{},
		Cmd: []string{
			"--nodekeyhex", randStringBytes(64),
			"--v5",
		},
		Files: map[string][]byte{},
	}
	return spec, nil
}

func WalletBackendSpec(nodeSuffix, rpcAddr, notificationsAddr string) (*NodeSpec, error) {
	id := NodeID("wallet-backend-" + nodeSuffix)
	spec := &NodeSpec{
		ID:    id,
		Image: "kowalatech/wallet_backend:dev",
		Cmd: []string{
			"--node-endpoint", rpcAddr,
			"--notifications-endpoint", notificationsAddr,
		},
		Env: []string{},
		PortMapping: map[int32]int32{
			8080: 8080,
		},
		IsReadyFn: func(runner NodeRunner) error {
			res, err := http.Get(
				fmt.Sprintf("http://%s:8080/api/blockheight", runner.HostIP()),
			)
			if err != nil {
				return err
			}
			if res.StatusCode != 200 {
				return fmt.Errorf("wallet backend API response code is %v", res.StatusCode)
			}
			return nil
		},
	}
	return spec, nil
}

func TransactionsPersistanceSpec(nodeSuffix, rpcAddr, redisAddr string) (*NodeSpec, error) {
	id := NodeID("transactions-persistance-" + nodeSuffix)
	spec := &NodeSpec{
		ID:    id,
		Image: "kowalatech/transactions_persistance:dev",
		Cmd:   []string{},
		Env: []string{
			fmt.Sprintf("REDIS_ADDR=%v", redisAddr),
			fmt.Sprintf("TESTNET_RPC_ADDR=%v", rpcAddr),
		},
		PortMapping: map[int32]int32{},
	}
	return spec, nil
}

func NotificationsApiSpec(nodeSuffix, redisAddr string) (*NodeSpec, error) {
	id := NodeID("notifications-api-" + nodeSuffix)
	spec := &NodeSpec{
		ID:    id,
		Image: "kowalatech/backend_api:dev",
		Cmd:   []string{},
		Env: []string{
			"PORT=3000",
			fmt.Sprintf("REDIS_ADDR=%v", redisAddr),
		},
		PortMapping: map[int32]int32{
			3000: 3000,
		},
	}
	return spec, nil
}

func RedisSpec(nodeSuffix string) (*NodeSpec, error) {
	id := NodeID("redis-" + nodeSuffix)
	spec := &NodeSpec{
		ID:    id,
		Image: "redis:alpine",
		Cmd:   []string{},
		Env:   []string{},
		PortMapping: map[int32]int32{
			6379: 6379,
		},
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
