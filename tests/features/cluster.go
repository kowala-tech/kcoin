package features

import (
	"encoding/json"
	"fmt"
	"regexp"
	"time"

	"github.com/kowala-tech/kcoin/accounts"
	"github.com/kowala-tech/kcoin/cluster"
	"github.com/kowala-tech/kcoin/kcoin/genesis"
)

var (
	enodeSecretRegexp = regexp.MustCompile(`enode://([a-f0-9]*)@`)
)

func (ctx *Context) PrepareCluster() error {
	nodeRunner, err := cluster.NewDockerNodeRunner()
	if err != nil {
		return err
	}
	ctx.nodeRunner = nodeRunner

	if err := ctx.generateAccounts(); err != nil {
		return err
	}
	if err := ctx.buildGenesis(); err != nil {
		return err
	}
	if err := ctx.buildDockerImages(); err != nil {
		return err
	}
	if err := ctx.runBootnode(); err != nil {
		return err
	}
	if err := ctx.runGenesisValidator(); err != nil {
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

	contractOwnerAccount, err := ctx.newAccount()
	if err != nil {
		return err
	}
	ctx.contractOwnerAccount = *contractOwnerAccount

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
	if err := ctx.nodeRunner.BuildDockerImage("kowalatech/bootnode:dev", "bootnode.Dockerfile"); err != nil {
		return err
	}
	if err := ctx.nodeRunner.BuildDockerImage("kowalatech/kusd:dev", "kcoin.Dockerfile"); err != nil {
		return err
	}
	return nil
}

func (ctx *Context) runBootnode() error {
	bootnode, err := cluster.BootnodeNode()
	if err != nil {
		return err
	}
	if err := ctx.nodeRunner.Run(bootnode); err != nil {
		return err
	}
	err = waitFor("fetching bootnode enode", 1*time.Second, 20*time.Second, func() bool {
		bootnodeStdout, err := ctx.nodeRunner.Log(bootnode)
		if err != nil {
			return false
		}
		found := enodeSecretRegexp.FindStringSubmatch(bootnodeStdout)
		if len(found) != 2 {
			return false
		}
		enodeSecret := found[1]
		bootnodeIP, err := ctx.nodeRunner.IP(bootnode)
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

	genesisValidatorNode := cluster.NewKcoinNodeBuilder().
		WithBootnode(ctx.bootnode).
		WithLogLevel(3).
		WithName("genesis-validator").
		WithSyncMode("full").
		WithNetworkId(ctx.chainID.String()).
		WithGenesis(ctx.genesis).
		Node()

	if err := ctx.nodeRunner.Run(genesisValidatorNode); err != nil {
		return err
	}
	return nil
}

func (ctx *Context) buildGenesis() error {
	newGenesis, err := genesis.GenerateGenesis(
		genesis.Options{
			Network:                        "test",
			MaxNumValidators:               "1",
			UnbondingPeriod:                "0",
			AccountAddressGenesisValidator: ctx.genesisValidatorAccount.Address.Hex(),
			SmartContractsOwner:            ctx.contractOwnerAccount.Address.Hex(),
			PrefundedAccounts: []genesis.PrefundedAccount{
				{
					AccountAddress: ctx.genesisValidatorAccount.Address.Hex(),
					Balance:        "0x200000000000000000000000000000000000000000000000000000000000000",
				},
				{
					AccountAddress: ctx.seederAccount.Address.Hex(),
					Balance:        "0x200000000000000000000000000000000000000000000000000000000000000",
				},
				{
					AccountAddress: ctx.contractOwnerAccount.Address.Hex(),
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
