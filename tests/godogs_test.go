package tests

import (
	"log"
	"math/big"
	"regexp"
	"strings"

	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"github.com/kowala-tech/kcoin/tests/features"
	"sync"
)

var (
	chainID = big.NewInt(1000)

	// contexts are different for each feature
	contextMapper = newContextMapper()
)

func init() {
	if err := features.InitLogs(features.LogsDir); err != nil {
		panic(err)
	}
	fmt.Println("initLogs DONE")

	if err := features.BuildDockerImages(); err != nil {
		panic(err)
	}
	fmt.Println("buildDockerImages DONE")
}

type contexts struct{
	ctx *features.Context
	validationCtx *features.ValidationContext
}

type ctxMapper struct {
	m map[string]contexts
	l sync.RWMutex
}

func newContextMapper() *ctxMapper {
	return &ctxMapper{make(map[string]contexts), sync.RWMutex{}}
}

func (mapper *ctxMapper) getContext(name string) (contexts, bool) {
	mapper.l.RLock()
	ctxs, ok := mapper.m[name]
	mapper.l.RUnlock()

	return ctxs, ok
}

func (mapper *ctxMapper) setContext(name string, ctxs contexts) {
	mapper.l.RLock()
	mapper.m[name] = ctxs
	mapper.l.RUnlock()
}

func FeatureContext(s *godog.Suite) {
	var ctxs contexts
	s.BeforeFeature(func(ft *gherkin.Feature) {
		ctxs.ctx = features.NewTestContext(chainID, ft.Name)
		ctxs.ctx.Name = getFeatureName(ft.Name)

		ctxs.validationCtx = features.NewValidationContext(ctxs.ctx)

		if err := ctxs.ctx.PrepareCluster(); err != nil {
			log.Fatal(err)
		}

		setFeatureSteps(s.Step, ctxs)
		contextMapper.setContext(ft.Name, ctxs)
	})

	s.AfterFeature(func(ft *gherkin.Feature) {
		context, _ := contextMapper.getContext(ft.Name)
		if err := context.ctx.DeleteCluster(); err != nil {
			log.Fatal(err)
		}
	})

	s.BeforeScenario(func(interface{}) {
		ctxs.ctx.Reset()
		ctxs.validationCtx.Reset()
	})
}

func setFeatureSteps(setStep func(expr interface{}, stepFunc interface{}), ctxs contexts) {
	// Creating accounts
	setStep(`^I have the following accounts:$`, ctxs.ctx.IHaveTheFollowingAccounts)
	setStep(`^I created an account with password '(\w+)'$`, ctxs.ctx.ICreatedAnAccountWithPassword)

	// Unlocking accounts
	setStep(`^I unlock the account (\w+) with password '(\w+)'$`, ctxs.ctx.IUnlockAccountWithPassword)
	setStep(`^I try to unlock the account (\w+) with password '(\w+)'$`, ctxs.ctx.ITryUnlockAccountWithPassword)
	setStep(`^I try to unlock my account with password '(\w+)'$`, ctxs.ctx.ITryUnlockMyAccountWithPassword)

	setStep(`^I should get my account unlocked$`, ctxs.ctx.IGotAccountUnlocked)
	setStep(`^I should get an error unlocking the account$`, ctxs.ctx.IGotErrorUnlocking)

	// Transactions
	setStep(`^I transfer (\d+) kcoins? from (\w+) to (\w+)$`, ctxs.ctx.ITransferKUSD)
	setStep(`^I try to transfer (\d+) kcoins? from (\w+) to (\w+)$`, ctxs.ctx.ITryTransferKUSD)
	setStep(`^the transaction should fail$`, ctxs.ctx.LastTransactionFailed)
	setStep(`^only one transaction should be done$`, ctxs.ctx.OnlyOneTransactionIsDone)
	setStep(`^the transaction hash the same$`, ctxs.ctx.TransactionHashTheSame)

	// Balances
	setStep(`^the balance of (\w+) should be (\d+) kcoins?$`, ctxs.ctx.TheBalanceIsExactly)
	setStep(`^the balance of (\w+) should be around (\d+) kcoins?$`, ctxs.ctx.TheBalanceIsAround)
	setStep(`^the balance of (\w+) should be greater (\d+) kcoins?$`, ctxs.ctx.TheBalanceIsGreater)
	setStep(`^the transaction should fail$`, ctxs.ctx.LastTransactionFailed)

	// validation
	setStep(`^I start validator with (\d+) kcoins deposit$`, ctxs.validationCtx.IStartTheValidator)
	setStep(`^I have my node running using account (\w+)$`, ctxs.validationCtx.IHaveMyNodeRunning)
	setStep(`^I stop validation$`, ctxs.validationCtx.IStopValidation)
	setStep(`^I wait for the unbonding period to be over$`, ctxs.validationCtx.IWaitForTheUnbondingPeriodToBeOver)
	setStep(`^I withdraw my node from validation$`, ctxs.validationCtx.IWithdrawMyNodeFromValidation)
	setStep(`^there should be (\d+) kcoins available to me after (\d+) days$`, ctxs.validationCtx.ThereShouldBeTokensAvailableToMeAfterDays)
	setStep(`^My node should be not be a validator$`, ctxs.validationCtx.MyNodeShouldBeNotBeAValidator)
	setStep(`^I wait for my node to be synced$`, ctxs.validationCtx.IWaitForMyNodeToBeSynced)

	// Nodes
	setStep(`^I start a new node$`, ctxs.ctx.IStartANewNode)
	setStep(`^my node should sync with the network$`, ctxs.ctx.MyNodeShouldSyncWithTheNetwork)
	setStep(`^my node is already synchronised$`, ctxs.validationCtx.MyNodeIsAlreadySynchronised)
	setStep(`^I disconnect my node for (\d+) blocks and reconnect it$`, ctxs.ctx.IDisconnectMyNodeForBlocksAndReconnectIt)
	setStep(`^I start a new node with a different network ID$`, ctxs.ctx.IStartANewNodeWithADifferentNetworkID)
	setStep(`^my node should not sync with the network$`, ctxs.ctx.MyNodeShouldNotSyncWithTheNetwork)
	setStep(`^I start a new node with a different chain ID$`, ctxs.ctx.IStartANewNodeWithADifferentChainID)
	setStep(`^I start validator with (\d+) deposit and coinbase A$`, ctxs.ctx.IStartValidatorWithDepositAndCoinbaseA)
}

func getFeatureName(feature string) string {
	feature = strings.ToLower(feature)
	reg, _ := regexp.Compile("[^a-z0-9 ]+")
	feature = reg.ReplaceAllString(feature, "")
	feature = strings.Replace(feature, " ", "_", -1)

	return feature
}
