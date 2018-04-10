package tests

import (
	"time"

	"github.com/DATA-DOG/godog"
	"github.com/kowala-tech/kcoin/cluster"
	"github.com/kowala-tech/kcoin/tests/features"
)

var (
	k8sCluster           cluster.Cluster
	genesisValidatorName string
)

func FeatureContext(s *godog.Suite) {
	// Use this approach instead of `BeforeSuite` because we need it right away,
	//  the `BeforeSuite` runs after executing the current function
	if k8sCluster == nil {
		prepareCluster()
	}
	s.AfterSuite(cleanupCluster)

	context := features.NewTestContext(k8sCluster, genesisValidatorName)
	s.Step(`^I have the following accounts:$`, context.IHaveTheFollowingAccounts)
	s.Step(`^I transfer (\d+) kcoins? from (\w+) to (\w+)$`, context.ITransferKUSD)
	s.Step(`^I try to transfer (\d+) kcoins? from (\w+) to (\w+)$`, context.ITryTransferKUSD)
	s.Step(`^the balance of (\w+) should be (\d+) kcoins?$`, context.TheBalanceIsExactly)
	s.Step(`^the balance of (\w+) should be around (\d+) kcoins?$`, context.TheBalanceIsAround)
	s.Step(`^the transaction should fail$`, context.LastTransactionFailed)
}

func prepareCluster() {
	backend := cluster.NewMinikubeCluster("testing")
	if !backend.Exists() {
		if err := backend.Create(); err != nil {
			panic(err)
		}
	}
	k8sCluster = cluster.NewCluster(backend)

	if err := k8sCluster.Connect(); err != nil {
		panic(err)
	}

	k8sCluster.Cleanup() // Just in case the previous run didn't finish gracefully

	if err := k8sCluster.Initialize("519374298533"); err != nil {
		panic(err)
	}
	if err := k8sCluster.RunBootnode(); err != nil {
		panic(err)
	}
	name, err := k8sCluster.RunGenesisValidator()
	if err := err; err != nil {
		panic(err)
	}
	genesisValidatorName = name
	if err := k8sCluster.TriggerGenesisValidation(); err != nil {
		panic(err)
	}

	time.Sleep(3 * time.Second) // let the genesis validator generate some blocks
}

func cleanupCluster() {
	k8sCluster.Cleanup()
}
