package tests

import (
	"log"
	"math/big"
	"regexp"
	"strings"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"github.com/kowala-tech/kcoin/tests/features"
)

var (
	chainID = big.NewInt(1000)
)

func FeatureContext(s *godog.Suite) {
	context := features.NewTestContext(chainID)
	validationCtx := features.NewValidationContext(context)

	s.BeforeFeature(func(ft *gherkin.Feature) {
		context.Name = getFeatureName(ft.Name)

		if err := context.PrepareCluster(); err != nil {
			log.Fatal(err)
		}
	})

	s.AfterFeature(func(ft *gherkin.Feature) {
		if err := context.DeleteCluster(); err != nil {
			log.Fatal(err)
		}
	})

	s.BeforeScenario(func(interface{}) {
		context.Reset()
		validationCtx.Reset()
	})

	// Creating accounts
	s.Step(`^I have the following accounts:$`, context.IHaveTheFollowingAccounts)
	s.Step(`^I created an account with password '(\w+)'$`, context.ICreatedAnAccountWithPassword)

	// Unlocking accounts
	s.Step(`^I unlock the account (\w+) with password '(\w+)'$`, context.IUnlockAccountWithPassword)
	s.Step(`^I try to unlock the account (\w+) with password '(\w+)'$`, context.ITryUnlockAccountWithPassword)
	s.Step(`^I try to unlock my account with password '(\w+)'$`, context.ITryUnlockMyAccountWithPassword)

	s.Step(`^I should get my account unlocked$`, context.IGotAccountUnlocked)
	s.Step(`^I should get an error unlocking the account$`, context.IGotErrorUnlocking)

	// Transactions
	s.Step(`^I transfer (\d+) kcoins? from (\w+) to (\w+)$`, context.ITransferKUSD)
	s.Step(`^I try to transfer (\d+) kcoins? from (\w+) to (\w+)$`, context.ITryTransferKUSD)
	s.Step(`^the transaction should fail$`, context.LastTransactionFailed)
	s.Step(`^only one transaction should be done$`, context.OnlyOneTransactionIsDone)
	s.Step(`^the transaction hash the same$`, context.TransactionHashTheSame)

	// Balances
	s.Step(`^the balance of (\w+) should be (\d+) kcoins?$`, context.TheBalanceIsExactly)
	s.Step(`^the balance of (\w+) should be around (\d+) kcoins?$`, context.TheBalanceIsAround)
	s.Step(`^the balance of (\w+) should be greater (\d+) kcoins?$`, context.TheBalanceIsGreater)
	s.Step(`^the transaction should fail$`, context.LastTransactionFailed)

	// validation
	s.Step(`^I start validator with (\d+) kcoins deposit$`, validationCtx.IStartTheValidator)
	s.Step(`^I have my node running using account (\w+)$`, validationCtx.IHaveMyNodeRunning)
	s.Step(`^I stop validation$`, validationCtx.IStopValidation)
	s.Step(`^I wait for the unbonding period to be over$`, validationCtx.IWaitForTheUnbondingPeriodToBeOver)
	s.Step(`^I withdraw my node from validation$`, validationCtx.IWithdrawMyNodeFromValidation)
	s.Step(`^there should be (\d+) kcoins available to me after (\d+) days$`, validationCtx.ThereShouldBeTokensAvailableToMeAfterDays)
	s.Step(`^My node should be not be a validator$`, validationCtx.MyNodeShouldBeNotBeAValidator)
	s.Step(`^I wait for my node to be synced$`, validationCtx.IWaitForMyNodeToBeSynced)

	// Nodes
	s.Step(`^I start a new node$`, context.IStartANewNode)
	s.Step(`^my node should sync with the network$`, context.MyNodeShouldSyncWithTheNetwork)
	s.Step(`^my node is already synchronised$`, validationCtx.MyNodeIsAlreadySynchronised)
	s.Step(`^I disconnect my node for (\d+) blocks and reconnect it$`, context.IDisconnectMyNodeForBlocksAndReconnectIt)
	s.Step(`^I start a new node with a different network ID$`, context.IStartANewNodeWithADifferentNetworkID)
	s.Step(`^my node should not sync with the network$`, context.MyNodeShouldNotSyncWithTheNetwork)
	s.Step(`^I start a new node with a different chain ID$`, context.IStartANewNodeWithADifferentChainID)
	s.Step(`^I start validator with (\d+) deposit and coinbase A$`, context.IStartValidatorWithDepositAndCoinbaseA)
}

func getFeatureName(feature string) string {
	feature = strings.ToLower(feature)
	reg, _ := regexp.Compile("[^a-z0-9 ]+")
	feature = reg.ReplaceAllString(feature, "")
	feature = strings.Replace(feature, " ", "_", -1)

	return feature
}
