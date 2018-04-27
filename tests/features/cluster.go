package features

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/kowala-tech/kcoin/cluster"
)

var showLogs = os.Getenv("KOWALA_LOGS")

func (ctx *Context) PrepareCluster() error {
	seederAccount, err := ctx.AccountsStorage.NewAccount("test")
	if err != nil {
		return err
	}
	if err := ctx.AccountsStorage.Unlock(seederAccount, "test"); err != nil {
		return err
	}

	ctx.seederAccount = seederAccount
	var backend cluster.Backend
	if ip, port := getStaticClusterConfig(); ip != "" && port != 0 {
		backend = cluster.NewStaticCluster(ip, port)
	} else {
		backend = cluster.NewMinikubeCluster("minikube")
	}
	if !backend.Exists() {
		if err := backend.Create(); err != nil {
			return err
		}
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

func (ctx *Context) GetFile(podName, fileName string) string {
	res, err := ctx.cluster.ExecCMD(podName, []string{"cat", fileName})
	if res != nil {
		return res.StdOut
	}

	fmt.Printf("can't read a file %q: %s", fileName, err)

	return ""
}

func (ctx *Context) DeleteCluster() error {
	var logToFiles *bool
	switch showLogs {
	case "stdout":
		logToFiles = new(bool)
	case "files":
		res := true
		logToFiles = &res
	default:
		// nothing to do
	}
	if logToFiles != nil {
		ctx.PrintLogs(*logToFiles)
	}

	return ctx.cluster.Cleanup()
}

func (ctx *Context) PrintLogs(toFiles bool) {
	err := ctx.cluster.PrintLogs(toFiles)
	if err != nil {
		fmt.Println("Error on getting logs", err)
	}
}

func getStaticClusterConfig() (string, int) {
	rawPort := os.Getenv("K8S_DOCKER_PORT")
	rawIp := os.Getenv("K8S_CUSTER_IP")
	if rawIp == "" || rawPort == "" {
		return "", 0
	}
	parsedPort, err := strconv.Atoi(rawPort)
	if err != nil {
		fmt.Println("Invalid K8S_DOCKER_PORT, must be just a number")
		return "", 0
	}
	return rawIp, parsedPort
}
