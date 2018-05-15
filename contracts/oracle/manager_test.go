package oracle

import (
	"crypto/ecdsa"
	"errors"
	"math/big"
	"testing"

	"github.com/kowala-tech/kcoin/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/accounts/abi/bind/backends"
	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/core"
	"github.com/kowala-tech/kcoin/crypto"
	"github.com/kowala-tech/kcoin/params"
	"github.com/stretchr/testify/suite"
)

const (
	initialBalance = 10 // kUSD
	baseDeposit    = 1  // kUSD
	maxNumOracles  = 100
	freezePeriod   = 10 // days
	secondsPerDay  = 86400
)

var (
	errAlwaysFailingTransaction = errors.New("failed to estimate gas needed: gas required exceeds allowance or always failing transaction")
)

type OracleMgrSuite struct {
	suite.Suite
	backend                   *backends.SimulatedBackend
	contractOwner, randomUser *ecdsa.PrivateKey
	initialBalance            *big.Int
	baseDeposit               *big.Int
	maxNumOracles             *big.Int
	freezePeriod              *big.Int
}

func TestOracleMgrSuite(t *testing.T) {
	suite.Run(t, new(OracleMgrSuite))
}

func (suite *OracleMgrSuite) SetupSuite() {
	req := suite.Require()

	contractOwner, err := crypto.GenerateKey()
	req.NoError(err)
	randomUser, err := crypto.GenerateKey()
	req.NoError(err)

	suite.contractOwner = contractOwner
	suite.randomUser = randomUser
	suite.initialBalance = kusd(new(big.Int).SetUint64(initialBalance))
	suite.baseDeposit = kusd(new(big.Int).SetUint64(baseDeposit))
	suite.maxNumOracles = new(big.Int).SetUint64(maxNumOracles)
	suite.freezePeriod = new(big.Int).SetUint64(freezePeriod)
}

func (suite *OracleMgrSuite) NewSimulatedBackend() *backends.SimulatedBackend {
	contractOwnerAddr := getAddress(suite.contractOwner)
	randomUserAddr := getAddress(suite.randomUser)
	defaultAccount := core.GenesisAccount{Balance: suite.initialBalance}
	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		contractOwnerAddr: defaultAccount,
		randomUserAddr:    defaultAccount,
	})

	backend.Commit()

	return backend
}

func (suite *OracleMgrSuite) SetupTest() {
	suite.backend = suite.NewSimulatedBackend()
}

/*

func (suite *OracleMgrSuite) TestDeployOracleManager() {
	req := suite.Require()

	opts := bind.NewKeyedTransactor(suite.contractOwner)
	_, _, mgr, err := DeployOracleManager(opts, suite.backend, suite.baseDeposit, suite.maxNumOracles, suite.freezePeriod)
	req.NoError(err)
	req.NotNil(mgr)

	suite.backend.Commit()

	storedBaseDeposit, err := mgr.BaseDeposit(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedBaseDeposit)

	req.Equal(kusd(suite.baseDeposit), storedBaseDeposit)

	storedMaxNumOracles, err := mgr.MaxNumOracles(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedMaxNumOracles)

	req.Equal(suite.maxNumOracles, storedMaxNumOracles)

	storedFreezePeriod, err := mgr.FreezePeriod(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedFreezePeriod)

	req.Equal(dtos(suite.freezePeriod), storedFreezePeriod)
}

func (suite *OracleMgrSuite) TestDeployOracleManager_MaxNumOraclesEqualZero() {
	req := suite.Require()

	opts := bind.NewKeyedTransactor(suite.contractOwner)
	_, _, _, err := DeployOracleManager(opts, suite.backend, suite.baseDeposit, common.Big0, suite.freezePeriod)
	req.Equal(errAlwaysFailingTransaction, err)
}

func (suite *OracleMgrSuite) TestGetMinimumDeposit_NotFull() {
	req := suite.Require()

	// deploy oracle manager
	opts := bind.NewKeyedTransactor(suite.contractOwner)
	_, _, mgr, err := DeployOracleManager(opts, suite.backend, suite.baseDeposit, common.Big1, suite.freezePeriod)
	req.NoError(err)

	suite.backend.Commit()

	minDeposit, err := mgr.GetMinimumDeposit(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(minDeposit)

	req.Equal(kusd(suite.baseDeposit), minDeposit)
}

func (suite *OracleMgrSuite) TestGetMinimumDeposit_Full() {
	req := suite.Require()

	// deploy oracle manager with one spot available
	opts := bind.NewKeyedTransactor(suite.contractOwner)
	_, _, mgr, err := DeployOracleManager(opts, suite.backend, suite.baseDeposit, common.Big1, suite.freezePeriod)
	req.NoError(err)

	// register oracle
	deposit := kusd(suite.baseDeposit)
	opts = bind.NewKeyedTransactor(suite.randomUser)
	opts.Value = deposit
	_, err = mgr.RegisterOracle(opts)
	req.NoError(err)

	suite.backend.Commit()

	minDeposit, err := mgr.GetMinimumDeposit(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(minDeposit)

	// minimum deposit must be the smallest bid + 1
	req.Equal(new(big.Int).Add(deposit, common.Big1), minDeposit)
}

func (suite *OracleMgrSuite) TestRegisterOracle_WhenNotPaused() {
	req := suite.Require()

	// deploy oracle manager with one spot available
	opts := bind.NewKeyedTransactor(suite.contractOwner)
	_, _, mgr, err := DeployOracleManager(opts, suite.backend, suite.baseDeposit, common.Big1, suite.freezePeriod)
	req.NoError(err)

	// pause service
	_, err = mgr.Pause(opts)
	req.NoError(err)

	// register oracle must fail
	deposit := kusd(suite.baseDeposit)
	opts = bind.NewKeyedTransactor(suite.randomUser)
	opts.Value = deposit
	_, err = mgr.RegisterOracle(opts)
	req.Equal(errAlwaysFailingTransaction, err)
}

func (suite *OracleMgrSuite) TestRegisterOracle_OnlyNewCandidate() {
	req := suite.Require()

	// deploy oracle manager
	opts := bind.NewKeyedTransactor(suite.contractOwner)
	_, _, mgr, err := DeployOracleManager(opts, suite.backend, suite.baseDeposit, common.Big2, suite.freezePeriod)
	req.NoError(err)

	// register oracle
	deposit := kusd(suite.baseDeposit)
	opts = bind.NewKeyedTransactor(suite.randomUser)
	opts.Value = deposit
	_, err = mgr.RegisterOracle(opts)
	req.NoError(err)

	// registration must fail for the same user
	_, err = mgr.RegisterOracle(opts)
	req.Equal(errAlwaysFailingTransaction, err)
}

func (suite *OracleMgrSuite) TestRegisterOracle_OnlyWithMinDeposit() {
	req := suite.Require()

	// deploy oracle manager
	opts := bind.NewKeyedTransactor(suite.contractOwner)
	_, _, mgr, err := DeployOracleManager(opts, suite.backend, suite.baseDeposit, common.Big1, suite.freezePeriod)
	req.NoError(err)
	req.NotNil(mgr)

	// register oracle must fail
	deposit := kusd(new(big.Int).Sub(suite.baseDeposit, common.Big1))
	opts = bind.NewKeyedTransactor(suite.randomUser)
	opts.Value = deposit
	_, err = mgr.RegisterOracle(opts)
	req.Equal(errAlwaysFailingTransaction, err)
}

/*
func (suite *OracleMgrSuite) TestRegisterOracle_GreaterThan() {
	req := suite.Require()

	// deploy oracle manager with one spot available
	opts := bind.NewKeyedTransactor(suite.contractOwner)
	_, _, mgr, err := DeployOracleManager(opts, suite.backend, suite.baseDeposit, common.Big1, suite.freezePeriod)
	req.NoError(err)

	// register oracle
	deposit := kusd(suite.baseDeposit)
	opts = bind.NewKeyedTransactor(suite.contractOwner)
	opts.Value = deposit
	_, err = mgr.RegisterOracle(opts)
	req.NoError(err)
}

func (suite *OracleMgrSuite) TestRegisterOracle_LessOrEqualTo() {

}

func (suite *OracleMgrSuite) TestRegisterOracle_Replacement() {

}


func (suite *OracleMgrSuite) TestDeregisterOracle() {
	req := suite.Require()

	// deploy oracle manager
	ownerAuth := bind.NewKeyedTransactor(suite.contractOwner)
	_, _, mgr, err := DeployOracleManager(ownerAuth, suite.backend, suite.baseDeposit, common.Big1, suite.freezePeriod)
	req.NoError(err)

	// register oracle
	deposit := kusd(suite.baseDeposit)
	userAuth := bind.NewKeyedTransactor(suite.randomUser)
	userAuth.Value = deposit
	_, err = mgr.RegisterOracle(userAuth)
	req.NoError(err)

	// deregister oracle
	_, err = mgr.DeregisterOracle(userAuth)
	req.NoError(err)
}


func (suite *OracleMgrSuite) TestDeregisterOracle_WhenNotPaused() {
	req := suite.Require()

	// deploy oracle manager
	ownerAuth := bind.NewKeyedTransactor(suite.contractOwner)
	_, _, mgr, err := DeployOracleManager(ownerAuth, suite.backend, suite.baseDeposit, common.Big1, suite.freezePeriod)
	req.NoError(err)

	// register oracle
	deposit := kusd(suite.baseDeposit)
	userAuth := bind.NewKeyedTransactor(suite.randomUser)
	userAuth.Value = deposit
	_, err = mgr.RegisterOracle(userAuth)
	req.NoError(err)

	// pause service
	_, err = mgr.Pause(ownerAuth)
	req.NoError(err)

	// deregister oracle
	_, err = mgr.DeregisterOracle(userAuth)
	req.Equal(errAlwaysFailingTransaction, err)
}

func (suite *OracleMgrSuite) TestDeregisterOracle_OnlyOracle() {
	req := suite.Require()

	// deploy oracle manager
	ownerAuth := bind.NewKeyedTransactor(suite.contractOwner)
	_, _, mgr, err := DeployOracleManager(ownerAuth, suite.backend, suite.baseDeposit, common.Big1, suite.freezePeriod)
	req.NoError(err)

	// deregister oracle
	userAuth := bind.NewKeyedTransactor(suite.randomUser)
	_, err = mgr.DeregisterOracle(userAuth)
	req.Equal(errAlwaysFailingTransaction, err)
}
*/

func (suite *OracleMgrSuite) TestReleaseDeposits_WhenNotPaused() {
	req := suite.Require()

	// deploy oracle manager
	deployOpts := bind.NewKeyedTransactor(suite.contractOwner)
	_, _, mgr, err := DeployOracleManager(deployOpts, suite.backend, suite.baseDeposit, common.Big32, suite.freezePeriod)
	req.NoError(err)

	// pause service
	pauseOpts := bind.NewKeyedTransactor(suite.contractOwner)
	_, err = mgr.Pause(pauseOpts)
	req.NoError(err)

	// release deposits
	depositsOpts := bind.NewKeyedTransactor(suite.randomUser)
	_, err = mgr.ReleaseDeposits(depositsOpts)
	req.Equal(errAlwaysFailingTransaction, err)
}

func (suite *OracleMgrSuite) TestReleaseDeposits_NoAssets() {
	// deploy oracle manager
	deployOpts := bind.NewKeyedTransactor(suite.contractOwner)
	_, _, mgr, err := DeployOracleManager(deployOpts, suite.backend, suite.baseDeposit, common.Big32, suite.freezePeriod)
	req.NoError(err)

	// release deposits
	depositsOpts := bind.NewKeyedTransactor(suite.randomUser)
	_, err = mgr.ReleaseDeposits(depositsOpts)
	req.Equal(errAlwaysFailingTransaction, err)

	suite.backend.Commit()

	// @TODO (rgeraldes)
	// deposit count must be zero

	// balance must be <= initial balance (tx fees and no deposits)
	userBalance, err := suite.backend.BalanceAt(context.TODO(), getAddress(suite.randomUser), suite.backend.CurrentBlock().Number())
	req.NoError(err)
	
	req.True(userBalance.Cmp(suite.initialBalance) < 0)

	// @TODO (rgeraldes)
	// contract balance must be the same

}

func (suite *OracleMgrSuite) TestReleaseDeposits_FrozenAssets() {}
func (suite *OracleMgrSuite) TestReleaseDeposits_UnfrozenAssets() {}

func (suite *OracleMgrSuite) TestAddPrice_WhenNotPaused() {
	req := suite.Require()

	// deploy oracle manager
	deployOpts := bind.NewKeyedTransactor(suite.contractOwner)
	_, _, mgr, err := DeployOracleManager(deployOpts, suite.backend, suite.baseDeposit, common.Big32, suite.freezePeriod)
	req.NoError(err)

	// register oracle
	deposit := suite.baseDeposit
	registrationOpts := bind.NewKeyedTransactor(suite.randomUser)
	registrationOpts.Value = deposit
	_, err = mgr.RegisterOracle(registrationOpts)
	req.NoError(err)

	// pause service
	pauseOpts := bind.NewKeyedTransactor(suite.contractOwner)
	_, err = mgr.Pause(pauseOpts)
	req.NoError(err)

	// oracle submits price
	newPrice := common.Big1
	priceOpts := bind.NewKeyedTransactor(suite.randomUser)
	_, err = mgr.AddPrice(priceOpts, newPrice)
	req.Equal(errAlwaysFailingTransaction, err)
}

func (suite *OracleMgrSuite) TestAddPrice_OnlyOracle() {
	req := suite.Require()

	// deploy oracle manager
	deployOpts := bind.NewKeyedTransactor(suite.contractOwner)
	_, _, mgr, err := DeployOracleManager(deployOpts, suite.backend, suite.baseDeposit, common.Big32, suite.freezePeriod)
	req.NoError(err)

	// not-oracle submits price
	invalidPrice := common.Big1
	priceOpts := bind.NewKeyedTransactor(suite.randomUser)
	_, err = mgr.AddPrice(priceOpts, invalidPrice)
	req.Equal(errAlwaysFailingTransaction, err)
}

func (suite *OracleMgrSuite) TestAddPrice_OnlyValidPrice() {
	req := suite.Require()

	// deploy oracle manager
	deployOpts := bind.NewKeyedTransactor(suite.contractOwner)
	_, _, mgr, err := DeployOracleManager(deployOpts, suite.backend, suite.baseDeposit, common.Big32, suite.freezePeriod)
	req.NoError(err)

	// register oracle
	deposit := suite.baseDeposit
	registrationOpts := bind.NewKeyedTransactor(suite.randomUser)
	registrationOpts.Value = deposit
	_, err = mgr.RegisterOracle(registrationOpts)
	req.NoError(err)

	// oracle submits price
	invalidPrice := common.Big0
	priceOpts := bind.NewKeyedTransactor(suite.randomUser)
	_, err = mgr.AddPrice(priceOpts, invalidPrice)
	req.Equal(errAlwaysFailingTransaction, err)
}

func (suite *OracleMgrSuite) TestAddPrice() {
	req := suite.Require()

	// deploy oracle manager
	deployOpts := bind.NewKeyedTransactor(suite.contractOwner)
	_, _, mgr, err := DeployOracleManager(deployOpts, suite.backend, suite.baseDeposit, common.Big32, suite.freezePeriod)
	req.NoError(err)

	// register oracle
	deposit := suite.baseDeposit
	registrationOpts := bind.NewKeyedTransactor(suite.randomUser)
	registrationOpts.Value = deposit
	_, err = mgr.RegisterOracle(registrationOpts)
	req.NoError(err)

	// oracle submits price
	newPrice := common.Big1
	priceOpts := bind.NewKeyedTransactor(suite.randomUser)
	_, err = mgr.AddPrice(priceOpts, newPrice)
	req.NoError(err)

	suite.backend.Commit()

	// confirm price update
	storedPrice, err := mgr.Price(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedPrice)

	req.Equal(newPrice, storedPrice)
}

// dtos converts days to seconds
func dtos(days *big.Int) *big.Int {
	return new(big.Int).Mul(days, new(big.Int).SetUint64(secondsPerDay))
}

// kusd converts the value to kUSD
func kusd(value *big.Int) *big.Int {
	return new(big.Int).Mul(value, new(big.Int).SetUint64(params.Ether))
}

// getAddress return the address of the given private key
func getAddress(privateKey *ecdsa.PrivateKey) common.Address {
	return crypto.PubkeyToAddress(privateKey.PublicKey)
}
