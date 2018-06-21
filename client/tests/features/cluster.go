package features

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/kowala-tech/kcoin/client/accounts"
	"github.com/kowala-tech/kcoin/client/cluster"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/kcoinclient"
	"github.com/kowala-tech/kcoin/client/knode/genesis"
	"github.com/lazada/awg"
)

var (
	enodeSecretRegexp = regexp.MustCompile(`enode://([a-f0-9]*)@`)
)

func (ctx *Context) DeleteCluster() error {
	return ctx.nodeRunner.StopAll()
}

func (ctx *Context) PrepareCluster() error {
	var err error
	logsDir := "./logs"

	if err := ctx.initLogs(logsDir); err != nil {
		return err
	}

	if ctx.nodeRunner, err = cluster.NewDockerNodeRunner(logsDir, ctx.Name); err != nil {
		return err
	}

	if err = ctx.buildDockerImages(); err != nil {
		return err
	}

	if err := ctx.generateAccounts(); err != nil {
		return err
	}

	if err := ctx.buildGenesis(); err != nil {
		return err
	}

	if err := ctx.runNodes(); err != nil {
		return err
	}

	return nil
}

var initLogsOnce sync.Once

func (ctx *Context) initLogs(logsDir string) error {
	var err error
	initLogsOnce.Do(func() {
		if err = os.RemoveAll(logsDir); err != nil {
			return
		}

		if err = os.Mkdir(logsDir, 0700); err != nil {
			return
		}
	})

	return err
}

func (ctx *Context) runNodes() error {
	if err := ctx.runBootnode(); err != nil {
		return err
	}

	if err := ctx.runGenesisValidator(); err != nil {
		return err
	}

	if err := ctx.triggerGenesisValidation(); err != nil {
		return err
	}

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

var initImagesOnce sync.Once

func (ctx *Context) buildDockerImages() error {
	wg := awg.AdvancedWaitGroup{}

	initImagesOnce.Do(func() {
		fmt.Println("Building docker images")
		wg.Add(func() error {
			return ctx.nodeRunner.BuildDockerImage("kowalatech/bootnode:dev", "bootnode.Dockerfile")
		})
		wg.Add(func() error {
			return ctx.nodeRunner.BuildDockerImage("kowalatech/kusd:dev", "kcoin.Dockerfile")
		})
	})

	return wg.SetStopOnError(true).Start().GetLastError()
}

func (ctx *Context) runBootnode() error {
	bootnode, err := cluster.BootnodeSpec(ctx.nodeSuffix)
	if err != nil {
		return err
	}

	if err := ctx.nodeRunner.Run(bootnode, ctx.scenarioNumber); err != nil {
		return err
	}
	err = common.WaitFor("fetching bootnode enode", 1*time.Second, 20*time.Second, func() error {
		bootnodeStdout, err := ctx.nodeRunner.Log(bootnode.ID)
		if err != nil {
			return err
		}

		found := enodeSecretRegexp.FindStringSubmatch(bootnodeStdout)
		if len(found) != 2 {
			return fmt.Errorf("can't start a bootnode %q", bootnodeStdout)
		}

		enodeSecret := found[1]
		bootnodeIP, err := ctx.nodeRunner.IP(bootnode.ID)
		if err != nil {
			return err
		}
		ctx.bootnode = fmt.Sprintf("enode://%v@%v:33445", enodeSecret, bootnodeIP)

		return nil
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
		WithID("genesis-validator-"+ctx.nodeSuffix).
		WithSyncMode("full").
		WithNetworkId(ctx.chainID.String()).
		WithGenesis(ctx.genesis).
		WithAccount(ctx.AccountsStorage, ctx.genesisValidatorAccount).
		WithValidation().
		WithDeposit(big.NewInt(1)).
		NodeSpec()

	if err := ctx.nodeRunner.Run(spec, ctx.scenarioNumber); err != nil {
		return err
	}

	ctx.genesisValidatorNodeID = spec.ID

	return nil
}

func (ctx *Context) runRpc() error {
	if ctx.rpcPort == 0 {
		ctx.rpcPort = 8080 + portCounter.Get()
	}

	spec := cluster.NewKcoinNodeBuilder().
		WithBootnode(ctx.bootnode).
		WithLogLevel(3).
		WithID("rpc-" + ctx.nodeSuffix).
		WithSyncMode("full").
		WithNetworkId(ctx.chainID.String()).
		WithGenesis(ctx.genesis).
		WithRpc(ctx.rpcPort).
		NodeSpec()

	if err := ctx.nodeRunner.Run(spec, ctx.scenarioNumber); err != nil {
		return err
	}

	rpcAddr := fmt.Sprintf("http://%v:%v", ctx.nodeRunner.HostIP(), ctx.rpcPort)
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

	return common.WaitFor("validation starts", 2*time.Second, 20*time.Second, func() error {
		res, err := ctx.nodeRunner.Exec(ctx.genesisValidatorNodeID, cluster.KcoinExecCommand("eth.blockNumber"))
		if err != nil {
			return err
		}

		parsed, err := strconv.Atoi(strings.TrimSpace(res.StdOut))
		if err != nil {
			return err
		}

		if parsed <= 0 {
			return fmt.Errorf("can't start validation %q", res.StdOut)
		}

		return nil
	})
}

func (ctx *Context) buildGenesis() error {
	validatorAddr := ctx.genesisValidatorAccount.Address.Hex()
	baseDeposit := uint64(10000000) // 10000000 mUSD

	newGenesis, err := genesis.Generate(genesis.Options{
		Network: "test",
		Consensus: &genesis.ConsensusOpts{
			Engine:           "tendermint",
			MaxNumValidators: 10,
			FreezePeriod:     30,
			BaseDeposit:      baseDeposit,
			Validators: []genesis.Validator{{
				Address: validatorAddr,
				Deposit: baseDeposit,
			}},
			MiningToken: &genesis.MiningTokenOpts{
				Name:     "mUSD",
				Symbol:   "mUSD",
				Cap:      1000,
				Decimals: 18,
				Holders:  []genesis.TokenHolder{{Address: validatorAddr, NumTokens: baseDeposit}},
			},
		},
		Governance: &genesis.GovernanceOpts{
			Origin:           "0x259be75d96876f2ada3d202722523e9cd4dd917d",
			Governors:        []string{"0x259be75d96876f2ada3d202722523e9cd4dd917d"},
			NumConfirmations: 1,
		},
		DataFeedSystem: &genesis.DataFeedSystemOpts{
			MaxNumOracles: 10,
			FreezePeriod:  0,
			BaseDeposit:   0,
		},
	})
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
