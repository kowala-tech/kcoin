package features

import (
	"encoding/json"
	"fmt"
	"math/big"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/kowala-tech/kcoin/accounts"
	"github.com/kowala-tech/kcoin/cluster"
	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/kcoin/genesis"
	"github.com/kowala-tech/kcoin/kcoinclient"
	"github.com/lazada/awg"
)

var (
	enodeSecretRegexp = regexp.MustCompile(`enode://([a-f0-9]*)@`)
)

func (ctx *Context) DeleteCluster() error {
	return ctx.nodeRunner.StopAll()
}

func (ctx *Context) PrepareCluster() error {
	nodeRunner, err := cluster.NewDockerNodeRunner()
	if err != nil {
		return err
	}
	ctx.nodeRunner = nodeRunner

	fmt.Println("Generating accounts")
	if err := ctx.generateAccounts(); err != nil {
		return err
	}
	fmt.Println("Building genesis")
	if err := ctx.buildGenesis(); err != nil {
		return err
	}
	fmt.Println("Building docker images")
	if err := ctx.buildDockerImages(); err != nil {
		return err
	}
	fmt.Println("Running bootnode")
	if err := ctx.runBootnode(); err != nil {
		return err
	}
	fmt.Println("Running genesis validator")
	if err := ctx.runGenesisValidator(); err != nil {
		fmt.Println(ctx.nodeRunner.Log(ctx.genesisValidatorNodeID))
		return err
	}
	fmt.Println("Triggering genesis validation")
	if err := ctx.triggerGenesisValidation(); err != nil {
		fmt.Println(ctx.nodeRunner.Log(ctx.genesisValidatorNodeID))
		return err
	}
	fmt.Println("Running RPC node")
	if err := ctx.runRpc(); err != nil {
		return err
	}
	return nil
}

func (ctx *Context) generateAccounts() error {
	seederAccount, err := ctx.newAccount()
	if err != nil {
		return err
	}
	ctx.seederAccount = *seederAccount

	genesisValidatorAccount, err := ctx.newAccount()
	if err != nil {
		return err
	}
	ctx.genesisValidatorAccount = *genesisValidatorAccount

	fmt.Println("Genesis validator", ctx.genesisValidatorAccount.Address.String())

	return nil
}

func (ctx *Context) newAccount() (*accounts.Account, error) {
	acc, err := ctx.AccountsStorage.NewAccount("test")
	if err != nil {
		return nil, err
	}
	if err := ctx.AccountsStorage.Unlock(acc, "test"); err != nil {
		return nil, err
	}
	return &acc, nil
}

func (ctx *Context) buildDockerImages() error {
	wg := awg.AdvancedWaitGroup{}
	wg.Add(func() error {
		return ctx.nodeRunner.BuildDockerImage("kowalatech/bootnode:dev", "bootnode.Dockerfile")
	})
	wg.Add(func() error {
		return ctx.nodeRunner.BuildDockerImage("kowalatech/kusd:dev", "kcoin.Dockerfile")
	})

	return wg.SetStopOnError(true).Start().GetLastError()
}

func (ctx *Context) runBootnode() error {
	bootnode, err := cluster.BootnodeSpec()
	if err != nil {
		return err
	}
	if err := ctx.nodeRunner.Run(bootnode); err != nil {
		return err
	}
	err = common.WaitFor("fetching bootnode enode", 1*time.Second, 20*time.Second, func() bool {
		bootnodeStdout, err := ctx.nodeRunner.Log(bootnode.ID)
		if err != nil {
			return false
		}
		found := enodeSecretRegexp.FindStringSubmatch(bootnodeStdout)
		if len(found) != 2 {
			return false
		}
		enodeSecret := found[1]
		bootnodeIP, err := ctx.nodeRunner.IP(bootnode.ID)
		if err != nil {
			return false
		}
		ctx.bootnode = fmt.Sprintf("enode://%v@%v:33445", enodeSecret, bootnodeIP)
		return true
	})

	if err != nil {
		return err
	}
	return nil
}

func (ctx *Context) runGenesisValidator() error {
	spec := cluster.NewKcoinNodeBuilder().
		WithBootnode(ctx.bootnode).
		WithLogLevel(3).
		WithID("genesis-validator").
		WithSyncMode("full").
		WithNetworkId(ctx.chainID.String()).
		WithGenesis(ctx.genesis).
		WithAccount(ctx.AccountsStorage, ctx.genesisValidatorAccount).
		WithValidation().
		WithDeposit(big.NewInt(1)).
		NodeSpec()

	if err := ctx.nodeRunner.Run(spec); err != nil {
		return err
	}

	ctx.genesisValidatorNodeID = spec.ID
	return nil
}

func (ctx *Context) runRpc() error {
	spec := cluster.NewKcoinNodeBuilder().
		WithBootnode(ctx.bootnode).
		WithLogLevel(3).
		WithID("rpc").
		WithSyncMode("full").
		WithNetworkId(ctx.chainID.String()).
		WithGenesis(ctx.genesis).
		WithRpc(8080).
		NodeSpec()

	if err := ctx.nodeRunner.Run(spec); err != nil {
		return err
	}

	rpcAddr := fmt.Sprintf("http://%v:%v", ctx.nodeRunner.HostIP(), 8080)
	client, err := kcoinclient.Dial(rpcAddr)
	if err != nil {
		return err
	}

	ctx.client = client
	return nil
}

func (ctx *Context) triggerGenesisValidation() error {
	command := fmt.Sprintf(`
		personal.unlockAccount(eth.coinbase, "test");
		eth.sendTransaction({from:eth.coinbase,to: "%v",value: 1})
	`, ctx.seederAccount.Address.Hex())
	_, err := ctx.nodeRunner.Exec(ctx.genesisValidatorNodeID, cluster.KcoinExecCommand(command))
	if err != nil {
		return err
	}

	return common.WaitFor("validation starts", 2*time.Second, 20*time.Second, func() bool {
		res, err := ctx.nodeRunner.Exec(ctx.genesisValidatorNodeID, cluster.KcoinExecCommand("eth.blockNumber"))
		if err != nil {
			return false
		}
		parsed, err := strconv.Atoi(strings.TrimSpace(res.StdOut))
		if err != nil {
			return false
		}
		return parsed > 0
	})
}

func (ctx *Context) buildGenesis() error {
	newGenesis, err := genesis.GenerateGenesis(
		genesis.Options{
			Network:                        "test",
			MaxNumValidators:               "5",
			UnbondingPeriod:                "5",
			AccountAddressGenesisValidator: ctx.genesisValidatorAccount.Address.Hex(),
			SmartContractsOwner:            "0x259be75d96876f2ada3d202722523e9cd4dd917d",
			PrefundedAccounts: []genesis.PrefundedAccount{
				{
					AccountAddress: ctx.genesisValidatorAccount.Address.Hex(),
					Balance:        "0x200000000000000000000000000000000000000000000000000000000000000",
				},
				{
					AccountAddress: ctx.seederAccount.Address.Hex(),
					Balance:        "0x200000000000000000000000000000000000000000000000000000000000000",
				},
			},
		},
	)
	if err != nil {
		return err
	}

	rawJson, err := json.Marshal(newGenesis)
	if err != nil {
		return err
	}
	ctx.genesis = rawJson

	return nil
}
