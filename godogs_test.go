package kowala

import (
	"github.com/DATA-DOG/godog"
	"github.com/kowala-tech/kUSD/tests/features"
)

func FeatureContext(s *godog.Suite) {
	context := features.NewTestContext()
	s.Step(`^I have an account (\w+) with (\d+) kUSD$`, context.IHaveAnAccountWithKUSD)
	s.Step(`^I transfer (\d+) kUSD from (\w+) to (\w+)$`, context.ITransferKUSD)
	s.Step(`^the balance of (\w+) is (\d+) kUSD$`, context.TheBalanceIs)
}
