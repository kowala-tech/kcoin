package features

import (
	"os"
	"time"

	"github.com/kowala-tech/kcoin/cluster"
)

func (ctx *Context) IsClusterReady() bool {
	return ctx.cluster != nil
}

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
	if os.Getenv("CI") == "drone" {
		backend = cluster.NewStaticCluster("172.31.94.205", 2376)
	} else {
		backend = cluster.NewMinikubeCluster("testing")
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

	ctx.cluster.Cleanup() // Just in case the previous run didn't finish gracefully

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
