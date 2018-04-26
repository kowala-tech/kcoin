package features

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/kowala-tech/kcoin/cluster"
)

var (
	k8sClusterNotConfiguredErr = errors.New("kubernetes cluster not configured")
	dockerK8SEnvVars           = []string{
		"DOCKER_TLS_VERIFY",
		"DOCKER_HOST",
		"DOCKER_CERT_PATH",
		"DOCKER_API_VERSION",
	}
)

func (ctx *Context) PrepareCluster() error {
	seederAccount, err := ctx.AccountsStorage.NewAccount("test")
	if err != nil {
		return err
	}
	if err := ctx.AccountsStorage.Unlock(seederAccount, "test"); err != nil {
		return err
	}

	ctx.seederAccount = seederAccount
	backend, err := getBackend()
	if err != nil {
		return err
	}
	ctx.cluster = cluster.NewCluster(backend)

	if err := ctx.cluster.Connect(); err != nil {
		return err
	}

	if err := ctx.cluster.Initialize(ctx.chainID.String(), ctx.seederAccount.Address); err != nil {
		return err
	}
	if err := ctx.cluster.RunBootnode(); err != nil {
		return err
	}
	_, err = ctx.cluster.RunGenesisValidator()
	if err := err; err != nil {
		return err
	}
	if err := ctx.cluster.TriggerGenesisValidation(); err != nil {
		return err
	}
	_, err = ctx.cluster.RunRpcNode()
	if err != nil {
		return err
	}

	newClient, err := ctx.cluster.RpcClient()
	if err != nil {
		return err
	}
	ctx.client = newClient

	time.Sleep(3 * time.Second) // let the genesis validator generate some blocks
	return nil
}

func (ctx *Context) DeleteCluster() error {
	return ctx.cluster.Cleanup()
}

func getBackend() (cluster.Backend, error) {
	ip := os.Getenv("K8S_CLUSTER_IP")
	if ip == "" {
		return nil, k8sClusterNotConfiguredErr
	}
	masterUrl := os.Getenv("K8S_CLUSTER_MASTER_URL")
	if masterUrl == "" {
		return nil, k8sClusterNotConfiguredErr
	}
	token := os.Getenv("K8S_CLUSTER_TOKEN")
	if token == "" {
		return nil, k8sClusterNotConfiguredErr
	}
	extraEnv := make([]string, 0)

	for _, envVar := range dockerK8SEnvVars {
		prefixedEnvVar := fmt.Sprintf("K8S_%v", envVar)
		if value := os.Getenv(prefixedEnvVar); value != "" {
			extraEnv = append(extraEnv, fmt.Sprintf(`%v=%v`, envVar, value))
		}
	}
	if len(extraEnv) == 0 {
		return nil, k8sClusterNotConfiguredErr
	}

	return cluster.NewK8SBackend(ip, masterUrl, token, extraEnv), nil
}
