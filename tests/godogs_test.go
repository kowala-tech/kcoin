package tests

import (
	"math/big"

	"log"

	"github.com/DATA-DOG/godog"
	"github.com/kowala-tech/kcoin/tests/features"
)

var (
	chainID = big.NewInt(519374298533)
)

func FeatureContext(s *godog.Suite) {
	context := features.NewTestContext(chainID)
	validationCtx := features.NewValidationContext(context)

	s.BeforeSuite(func() {
		if err := context.PrepareCluster(); err != nil {
			log.Fatal(err)
		}
	})

	s.AfterSuite(func() {
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

	// Balances
	s.Step(`^the balance of (\w+) should be (\d+) kcoins?$`, context.TheBalanceIsExactly)
	s.Step(`^the balance of (\w+) should be around (\d+) kcoins?$`, context.TheBalanceIsAround)
	s.Step(`^the transaction should fail$`, context.LastTransactionFailed)

	// validation
	s.Step(`^I start validator with (\d+) kcoins deposit$`, validationCtx.IStartTheValidator)
	s.Step(`^I should be a validator$`, validationCtx.IShouldBeAValidator)
	s.Step(`^I have my node running using account (\w+)$`, validationCtx.IHaveMyNodeRunning)
	s.Step(`^I should be a validator$`, validationCtx.IShouldBeAValidator)
	s.Step(`^I stop validation$`, validationCtx.IStopValidation)
	s.Step(`^I wait for the unbonding period to be over$`, validationCtx.IWaitForTheUnbondingPeriodToBeOver)
	s.Step(`^I withdraw my node from validation$`, validationCtx.IWithdrawMyNodeFromValidation)
	s.Step(`^There should be (\d+) kcoins available to me after (\d+) days$`, validationCtx.ThereShouldBeTokensAvailableToMeAfterDays)
	s.Step(`^My node should be not be a validator$`, validationCtx.MyNodeShouldBeNotBeAValidator)
	s.Step(`^I wait for my node to be synced$`, validationCtx.IWaitForMyNodeToBeSynced)

	// Nodes
	s.Step(`^I start a new node$`, context.IStartANewNode)
	s.Step(`^My node should sync with the network$`, context.MyNodeShouldSyncWithTheNetwork)
	s.Step(`^My node is already synchronised$`, validationCtx.MyNodeIsAlreadySynchronised)
	s.Step(`^I disconnect my node for (\d+) blocks and reconnect it$`, context.IDisconnectMyNodeForBlocksAndReconnectIt)
	s.Step(`^I start a new node with a different network ID$`, context.IStartANewNodeWithADifferentNetworkID)
	s.Step(`^My node should not sync with the network$`, context.MyNodeShouldNotSyncWithTheNetwork)
	s.Step(`^I start a new node with a different chain ID$`, context.IStartANewNodeWithADifferentChainID)
	s.Step(`^I start validator with (\d+) deposit and coinbase A$`, context.IStartValidatorWithDepositAndCoinbaseA)
}
