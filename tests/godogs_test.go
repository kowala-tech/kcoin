package tests

import (
	"math/big"

	"github.com/DATA-DOG/godog"
	"github.com/kowala-tech/kcoin/tests/features"
)

var (
	chainID = big.NewInt(519374298533)
)

func FeatureContext(s *godog.Suite) {
	context := features.NewTestContext(chainID)

	s.BeforeSuite(func() {
		if err := context.PrepareCluster(); err != nil {
			panic(err)
		}
	})

	s.AfterSuite(func() {
		if err := context.DeleteCluster(); err != nil {
			panic(err)
		}
	})

	s.BeforeScenario(func(interface{}) {
		context.Reset()
	})

	// Genesis validated
	s.Step(`^I have validated genesis:$`, context.IValidationSuccessed)

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
}
