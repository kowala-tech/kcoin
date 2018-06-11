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

	"github.com/kowala-tech/kcoin/accounts"
	"github.com/kowala-tech/kcoin/cluster"
	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/kcoinclient"
	"github.com/kowala-tech/kcoin/knode/genesis"
	"github.com/lazada/awg"
	"sync/atomic"
	"github.com/kowala-tech/kcoin/crypto"
	"crypto/ecdsa"
)

var (
	enodeSecretRegexp = regexp.MustCompile(`enode://([a-f0-9]*)@`)
)

var n int64
const LogsDir = "./logs"

func (ctx *Context) DeleteCluster() error {
	return ctx.nodeRunner.StopAll()
}

func (ctx *Context) PrepareCluster() error {
	var err error

	k := atomic.AddInt64(&n, 1)
	fmt.Println("PrepareCluster", ctx.Name)

	if ctx.nodeRunner, err = cluster.NewDockerNodeRunner(LogsDir, ctx.Name); err != nil {
		return err
	}

	if err := ctx.generateAccounts(); err != nil {
		return err
	}
	fmt.Println(k, "generateAccounts DONE")

	if err := ctx.buildGenesis(); err != nil {
		return err
	}
	fmt.Println(k, "buildGenesis DONE")

	if err := ctx.runNodes(); err != nil {
		return err
	}
	fmt.Println(k, "runNodes DONE")

	return nil
}

var initLogsOnce sync.Once

func InitLogs(logsDir string) error {
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

func BuildDockerImages() error {
	wg := awg.AdvancedWaitGroup{}

	initImagesOnce.Do(func() {
		fmt.Println("Building docker images")
		wg.Add(func() error {
			return cluster.BuildDockerImage("kowalatech/bootnode:dev", "bootnode.Dockerfile")
		})
		wg.Add(func() error {
			return cluster.BuildDockerImage("kowalatech/kusd:dev", "kcoin.Dockerfile")
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


var (
	validator, _     = crypto.GenerateKey()
	deregistered, _  = crypto.GenerateKey()
	user, _          = crypto.GenerateKey()
	governor, _      = crypto.GenerateKey()
	author, _        = crypto.HexToECDSA("bfef37ae9ac5d5e7ebbbefc19f4e1f572a7ca7aa0d28e527b7d62950951cc5eb")
	validatorMgrAddr = common.HexToAddress("0x161ad311F1D66381C17641b1B73042a4CA731F9f")
	multiSigAddr     = common.HexToAddress("0xA143ac5ec5D95f16aFD5Fc3B09e0aDaf360ffC9e")
	tokenAddr        = common.HexToAddress("0xB012F49629258C9c35b2bA80cD3dc3C841d9719D")
	secondsPerDay    = new(big.Int).SetUint64(86400)
)

func GetDefaultOpts(baseDeposit uint64, validatorAddress string) genesis.Options {
	tokenHolder := genesis.TokenHolder{
		Address:   validatorAddress,
		NumTokens: baseDeposit,
	}

	opts := genesis.Options{
		Network: "test",
		Consensus: &genesis.ConsensusOpts{
			Engine:           "tendermint",
			MaxNumValidators: 10,
			FreezePeriod:     30,
			BaseDeposit:      baseDeposit,
			Validators: []genesis.Validator{{
				Address: tokenHolder.Address,
				Deposit: tokenHolder.NumTokens,
			}},
			MiningToken: &genesis.MiningTokenOpts{
				Name:     "mUSD",
				Symbol:   "mUSD",
				Cap:      1000,
				Decimals: 18,
				Holders:  []genesis.TokenHolder{tokenHolder, {Address: getAddress(user).Hex(), NumTokens: baseDeposit * 3}},
			},
		},
		Governance: &genesis.GovernanceOpts{
			Origin:           getAddress(author).Hex(),
			Governors:        []string{getAddress(governor).Hex()},
			NumConfirmations: 1,
		},
		DataFeedSystem: &genesis.DataFeedSystemOpts{
			MaxNumOracles: 10,
			FreezePeriod:  0,
			BaseDeposit:   0,
		},
		PrefundedAccounts: []genesis.PrefundedAccount{
			{
				Address: tokenHolder.Address,
				Balance: 10,
			},
			{
				Address: getAddress(governor).Hex(),
				Balance: 10,
			},
			{
				Address: getAddress(user).Hex(),
				Balance: 10,
			},
			{
				Address: getAddress(deregistered).Hex(),
				Balance: 10,
			},
		},
	}

	return opts
}

// getAddress return the address of the given private key
func getAddress(privateKey *ecdsa.PrivateKey) common.Address {
	return crypto.PubkeyToAddress(privateKey.PublicKey)
}

func (ctx *Context) buildGenesis() error {
	genesisLock.Lock()
	fmt.Println("1")
	fmt.Println("LOCK")
	defer func() {
		fmt.Println("UNLOCK")
		genesisLock.Unlock()
	}()

	validatorAddr := ctx.genesisValidatorAccount.Address.Hex()
	baseDeposit := uint64(100)
	fmt.Println("2")

	opts := genesis.Options{
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
			Origin:           validatorAddr,
			Governors:        []string{validatorAddr},
			NumConfirmations: 1,
		},
		DataFeedSystem: &genesis.DataFeedSystemOpts{
			MaxNumOracles: 10,
			FreezePeriod:  0,
			BaseDeposit:   0,
		},
		PrefundedAccounts: []genesis.PrefundedAccount{
			{
				Address: validatorAddr,
				Balance: 1000,
			},
		},
	}
	fmt.Println("3")

	opts = GetDefaultOpts(baseDeposit, validatorAddr)
	fmt.Println("4")

	newGenesis, err := genesis.Generate(opts)
	fmt.Println("5")
	if err != nil {
		return err
	}

	rawJson, err := json.Marshal(newGenesis)
	if err != nil {
		return err
	}
	fmt.Println("6")
	ctx.genesis = rawJson
	fmt.Println("7")
	time.Sleep(5*time.Second)
	return nil
}
