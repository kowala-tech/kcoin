package impl

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"strconv"
	"strings"

	"sync/atomic"

	"github.com/kowala-tech/kcoin/client/accounts"
	"github.com/kowala-tech/kcoin/client/accounts/keystore"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core/types"
	"github.com/kowala-tech/kcoin/client/kcoinclient"
	"github.com/kowala-tech/kcoin/client/log"
	"github.com/kowala-tech/kcoin/e2e/cluster"
)

type Context struct {
	Name            string
	AccountsStorage *keystore.KeyStore

	genesisOptions *e2eGenesisOptions

	logsToStdout bool

	// cluster config
	genesis  []byte
	bootnode string

	nodeRunner             cluster.NodeRunner
	genesisValidatorNodeID cluster.NodeID
	rpcNodeID              cluster.NodeID
	rpcPort                int32
	client                 *kcoinclient.Client
	chainID                *big.Int

	genesisValidatorAccount   accounts.Account
	mtokensSeederAccount      accounts.Account
	mtokensGovernanceAccounts []accounts.Account
	kusdSeederAccount         accounts.Account
	accounts                  map[string]accounts.Account

	lastTx              *types.Transaction
	lastTxErr           error
	lastTxStartingBlock *big.Int

	lastUnlockErr error

	scenarioNumber *int32
	nodeSuffix     string

	waiter doer
}

type e2eGenesisOptions struct {
	requiredGovernanceConfirmations uint64
}

func defaultGenesisOptions() *e2eGenesisOptions {
	return &e2eGenesisOptions{
		requiredGovernanceConfirmations: 2,
	}
}

func NewTestContext(chainID *big.Int, logsToStdout bool) *Context {
	tmpdir, _ := ioutil.TempDir("", "eth-keystore-test")
	accountsStorage := keystore.NewKeyStore(tmpdir, 2, 1)

	ctx := &Context{
		logsToStdout:    logsToStdout,
		AccountsStorage: accountsStorage,
		chainID:         chainID,

		accounts:       make(map[string]accounts.Account),
		scenarioNumber: new(int32),
		nodeSuffix:     common.RandomString(4),

		genesisOptions: defaultGenesisOptions(),
	}

	ctx.waiter = common.NewWaiter(ctx)

	return ctx
}

func (ctx *Context) GetScenarioNumber() int32 {
	return atomic.LoadInt32(ctx.scenarioNumber)
}

func (ctx *Context) IncreaseScenarioNumber() int32 {
	return atomic.AddInt32(ctx.scenarioNumber, 1)
}

func (ctx *Context) Reset() {
	ctx.accounts = make(map[string]accounts.Account)
	ctx.genesisOptions = defaultGenesisOptions()
	ctx.IncreaseScenarioNumber()
}

func (ctx *Context) execCommand(nodeID cluster.NodeID, command []string, response ...*cluster.ExecResponse) error {
	return ctx.makeExecFunc(nodeID, command, response...)()
}

func (ctx *Context) getTokenBalance(at common.Address) (*big.Int, error) {
	res := &cluster.ExecResponse{}
	if err := ctx.execCommand(ctx.rpcNodeID, getTokenBalance(at), res); err != nil {
		return nil, err
	}

	currentBalanceBig, ok := new(big.Int).SetString(res.StdOut, 10)
	if !ok {
		return nil, fmt.Errorf("incorrect mToken deposit %q of %s", res.StdOut, at.String())
	}

	return currentBalanceBig, nil
}

func (ctx *Context) Do(cmd []string, condFunc func() error) error {
	return ctx.waiter.Do(ctx.makeExecFunc(ctx.rpcNodeID, cmd), condFunc)
}

func (ctx *Context) CurrentBlock() (uint64, error) {
	res := &cluster.ExecResponse{}
	if err := ctx.execCommand(ctx.rpcNodeID, blockNumberCommand(), res); err != nil {
		return 0, err
	}

	return strconv.ParseUint(strings.TrimSpace(res.StdOut), 10, 64)
}

func (ctx *Context) makeExecFunc(nodeID cluster.NodeID, command []string, response ...*cluster.ExecResponse) func() error {
	return func() error {
		res, err := ctx.nodeRunner.Exec(nodeID, command)
		if err != nil {
			if res != nil {
				log.Debug(res.StdOut)
			}

			return fmt.Errorf("error while executing '%v': %q", command, err)
		}

		if len(response) != 0 {
			*response[0] = *res
		}

		if err = isError(res.StdOut); err != nil {
			return fmt.Errorf("error while executing '%v': %q", command, err)
		}

		return nil
	}
}

type doer interface {
	Do(execFunc func() error, condFunc ...func() error) error
}
