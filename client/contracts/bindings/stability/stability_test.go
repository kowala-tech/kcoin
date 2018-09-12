package stability_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/accounts/abi/bind/backends"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/stability"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/stability/testfiles"
	"github.com/kowala-tech/kcoin/client/core"
	"github.com/kowala-tech/kcoin/client/crypto"
	"github.com/kowala-tech/kcoin/client/params"
	"github.com/stretchr/testify/suite"
)

var (
	owner, _       = crypto.GenerateKey()
	initialBalance = new(big.Int).Mul(new(big.Int).SetUint64(100), new(big.Int).SetUint64(params.Kcoin))
)

type StabilityContractSuite struct {
	suite.Suite
	backend *backends.SimulatedBackend
}

func TestStabilityContractSuite(t *testing.T) {
	suite.Run(t, new(StabilityContractSuite))
}

func (suite *StabilityContractSuite) BeforeTest(suiteName, testName string) {
	req := suite.Require()

	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		crypto.PubkeyToAddress(owner.PublicKey): core.GenesisAccount{
			Balance: initialBalance,
		},
	})
	req.NotNil(backend)
	suite.backend = backend
}

func (suite *StabilityContractSuite) TestDeploy() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// deploy price provider
	mockPrice := new(big.Int).Mul(common.Big1, big.NewInt(params.Kcoin))
	mockAddr, _, _, err := testfiles.DeployPriceProviderMock(transactOpts, suite.backend, mockPrice)
	req.NoError(err)
	req.NotZero(mockAddr)

	suite.backend.Commit()

	// deploy stability contract
	minDeposit := new(big.Int).Mul(common.Big32, big.NewInt(params.Kcoin))
	_, _, stabilityContract, err := stability.DeployStability(transactOpts, suite.backend, minDeposit, mockAddr)
	req.NoError(err)
	req.NotNil(stabilityContract)

	suite.backend.Commit()

	storedMinDeposit, err := stabilityContract.MinDeposit(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedMinDeposit)
	req.Equal(minDeposit, storedMinDeposit)
}

func (suite *StabilityContractSuite) TestSubscribe_NoSubscription_LessThanMinDeposit() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// deploy price provider
	mockPrice := new(big.Int).Mul(common.Big1, big.NewInt(params.Kcoin))
	mockAddr, _, _, err := testfiles.DeployPriceProviderMock(transactOpts, suite.backend, mockPrice)
	req.NoError(err)
	req.NotZero(mockAddr)

	suite.backend.Commit()

	// deploy stability contract
	minDeposit := new(big.Int).Mul(common.Big32, big.NewInt(params.Kcoin))
	_, _, stabilityContract, err := stability.DeployStability(transactOpts, suite.backend, minDeposit, mockAddr)
	req.NoError(err)
	req.NotNil(stabilityContract)

	suite.backend.Commit()

	subscribeOpts := bind.NewKeyedTransactor(owner)
	subscribeOpts.Value = new(big.Int).Sub(minDeposit, common.Big1)
	_, err = stabilityContract.Subscribe(subscribeOpts)
	req.Error(err, "value does not cover the min deposit")
}

func (suite *StabilityContractSuite) TestSubscribe_NoSubscription_GreaterOrEqualMinDeposit() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// deploy price provider
	mockPrice := new(big.Int).Mul(common.Big1, big.NewInt(params.Kcoin))
	mockAddr, _, _, err := testfiles.DeployPriceProviderMock(transactOpts, suite.backend, mockPrice)
	req.NoError(err)
	req.NotZero(mockAddr)

	suite.backend.Commit()

	// deploy stability contract
	minDeposit := new(big.Int).Mul(common.Big32, big.NewInt(params.Kcoin))
	_, _, stabilityContract, err := stability.DeployStability(transactOpts, suite.backend, minDeposit, mockAddr)
	req.NoError(err)
	req.NotNil(stabilityContract)

	suite.backend.Commit()

	// subscribe service
	subscribeOpts := bind.NewKeyedTransactor(owner)
	subscribeOpts.Value = minDeposit
	_, err = stabilityContract.Subscribe(subscribeOpts)
	req.NoError(err)

	suite.backend.Commit()

	// one subscriber
	count, err := stabilityContract.GetSubscriptionCount(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(count)
	req.Equal(common.Big1, count)

	// subscriber code and deposit must match the details provided earlier
	subscriptionDetails, err := stabilityContract.GetSubscriptionAtIndex(&bind.CallOpts{}, common.Big0)
	req.NoError(err)
	req.NotZero(subscriptionDetails)
	req.Equal(crypto.PubkeyToAddress(owner.PublicKey), subscriptionDetails.Code)
	req.Equal(minDeposit, subscriptionDetails.Deposit)
}

func (suite *StabilityContractSuite) TestSubscribe_HasSubscription_AnyValue() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// deploy price provider
	mockPrice := new(big.Int).Mul(common.Big1, big.NewInt(params.Kcoin))
	mockAddr, _, _, err := testfiles.DeployPriceProviderMock(transactOpts, suite.backend, mockPrice)
	req.NoError(err)
	req.NotZero(mockAddr)

	suite.backend.Commit()

	// deploy stability contract
	minDeposit := new(big.Int).Mul(common.Big32, big.NewInt(params.Kcoin))
	_, _, stabilityContract, err := stability.DeployStability(transactOpts, suite.backend, minDeposit, mockAddr)
	req.NoError(err)
	req.NotNil(stabilityContract)

	suite.backend.Commit()

	// subscribe service
	subscribeOpts := bind.NewKeyedTransactor(owner)
	subscribeOpts.Value = minDeposit
	_, err = stabilityContract.Subscribe(subscribeOpts)
	req.NoError(err)

	suite.backend.Commit()

	// subscribe with the same account for a new value
	subscribeOpts.Value = common.Big1
	_, err = stabilityContract.Subscribe(subscribeOpts)
	req.NoError(err)

	suite.backend.Commit()

	// still one subscriber
	count, err := stabilityContract.GetSubscriptionCount(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(count)
	req.Equal(common.Big1, count)

	// deposit should now be minDeposit + value of the second subscription
	subscriptionDetails, err := stabilityContract.GetSubscriptionAtIndex(&bind.CallOpts{}, common.Big0)
	req.NoError(err)
	req.NotZero(subscriptionDetails)
	req.Equal(crypto.PubkeyToAddress(owner.PublicKey), subscriptionDetails.Code)
	req.Equal(new(big.Int).Add(minDeposit, common.Big1), subscriptionDetails.Deposit)
}

func (suite *StabilityContractSuite) TestUnsubscribe_NotSubscriber() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// deploy price provider
	mockPrice := new(big.Int).Mul(common.Big1, big.NewInt(params.Kcoin))
	mockAddr, _, _, err := testfiles.DeployPriceProviderMock(transactOpts, suite.backend, mockPrice)
	req.NoError(err)
	req.NotZero(mockAddr)

	suite.backend.Commit()

	// deploy stability contract
	minDeposit := new(big.Int).Mul(common.Big32, big.NewInt(params.Kcoin))
	_, _, stabilityContract, err := stability.DeployStability(transactOpts, suite.backend, minDeposit, mockAddr)
	req.NoError(err)
	req.NotNil(stabilityContract)

	suite.backend.Commit()

	// unsubscribe service
	_, err = stabilityContract.Unsubscribe(transactOpts)
	req.Error(err, "user does not have a subscription")
}

func (suite *StabilityContractSuite) TestUnsubscribe_Subscriber_PriceLessThanOne() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// deploy price provider
	mockPrice := new(big.Int).Mul(common.Big0, big.NewInt(params.Kcoin))
	mockAddr, _, _, err := testfiles.DeployPriceProviderMock(transactOpts, suite.backend, mockPrice)
	req.NoError(err)
	req.NotZero(mockAddr)

	suite.backend.Commit()

	// deploy stability contract
	minDeposit := new(big.Int).Mul(common.Big32, big.NewInt(params.Kcoin))
	_, _, stabilityContract, err := stability.DeployStability(transactOpts, suite.backend, minDeposit, mockAddr)
	req.NoError(err)
	req.NotNil(stabilityContract)

	suite.backend.Commit()

	// subscribe service
	subscribeOpts := bind.NewKeyedTransactor(owner)
	subscribeOpts.Value = minDeposit
	_, err = stabilityContract.Subscribe(subscribeOpts)
	req.NoError(err)

	suite.backend.Commit()

	// unsubscribe service
	_, err = stabilityContract.Unsubscribe(transactOpts)
	req.Error(err, "price is less than one")
}

func (suite *StabilityContractSuite) TestUnsubscribe_Subscriber_PriceGreaterOrEqualOne() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// deploy price provider
	mockPrice := new(big.Int).Mul(common.Big1, big.NewInt(params.Kcoin))
	mockAddr, priceProviderTx, _, err := testfiles.DeployPriceProviderMock(transactOpts, suite.backend, mockPrice)
	req.NoError(err)
	req.NotNil(priceProviderTx)
	req.NotZero(mockAddr)

	suite.backend.Commit()

	// deploy stability contract
	minDeposit := new(big.Int).Mul(common.Big32, big.NewInt(params.Kcoin))
	_, stabilityTx, stabilityContract, err := stability.DeployStability(transactOpts, suite.backend, minDeposit, mockAddr)
	req.NoError(err)
	req.NotNil(stabilityTx)
	req.NotNil(stabilityContract)

	suite.backend.Commit()

	// subscribe service
	subscribeOpts := bind.NewKeyedTransactor(owner)
	subscribeOpts.Value = minDeposit
	subscribeTx, err := stabilityContract.Subscribe(subscribeOpts)
	req.NoError(err)
	req.NotNil(subscribeTx)

	suite.backend.Commit()

	// unsubscribe service
	unsubscribeTx, err := stabilityContract.Unsubscribe(transactOpts)
	req.NoError(err)
	req.NotNil(unsubscribeTx)

	suite.backend.Commit()

	// zero subscribers by now
	count, err := stabilityContract.GetSubscriptionCount(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(count)
	req.Zero(count.Uint64())

	// user balance must be equal to initial balance - deployTx cost - subscribeTx cost - unsubscribeTx cost
	priceProviderReceipt, err := suite.backend.TransactionReceipt(context.TODO(), priceProviderTx.Hash())
	req.NoError(err)
	req.NotNil(priceProviderReceipt)
	stabilityReceipt, err := suite.backend.TransactionReceipt(context.TODO(), stabilityTx.Hash())
	req.NoError(err)
	req.NotNil(stabilityReceipt)
	subscribeReceipt, err := suite.backend.TransactionReceipt(context.TODO(), subscribeTx.Hash())
	req.NoError(err)
	req.NotNil(subscribeReceipt)
	unsubscribeReceipt, err := suite.backend.TransactionReceipt(context.TODO(), unsubscribeTx.Hash())
	req.NoError(err)
	req.NotNil(unsubscribeReceipt)

	priceProviderCost := new(big.Int).Mul(params.ComputeUnitPrice, new(big.Int).SetUint64(priceProviderReceipt.ResourceUsage))
	stabilityCost := new(big.Int).Mul(params.ComputeUnitPrice, new(big.Int).SetUint64(stabilityReceipt.ResourceUsage))
	subscribeCost := new(big.Int).Mul(params.ComputeUnitPrice, new(big.Int).SetUint64(subscribeReceipt.ResourceUsage))
	unsubscribeCost := new(big.Int).Mul(params.ComputeUnitPrice, new(big.Int).SetUint64(unsubscribeReceipt.ResourceUsage))

	finalBalance := new(big.Int).Sub(initialBalance, priceProviderCost)
	finalBalance.Sub(finalBalance, stabilityCost)
	finalBalance.Sub(finalBalance, subscribeCost)
	finalBalance.Sub(finalBalance, unsubscribeCost)

	currentBalance, err := suite.backend.BalanceAt(context.TODO(), crypto.PubkeyToAddress(owner.PublicKey), suite.backend.CurrentBlock().Number())
	req.NoError(err)
	req.NotNil(currentBalance)
	req.Equal(finalBalance, currentBalance)
}
