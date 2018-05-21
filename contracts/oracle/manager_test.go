package oracle

import (
	"crypto/ecdsa"
	"math/big"
	"strings"
	"testing"

	"github.com/kowala-tech/kcoin/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/accounts/abi/bind/backends"
	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/core"
	"github.com/kowala-tech/kcoin/crypto"
	"github.com/kowala-tech/kcoin/params"
	"github.com/stretchr/testify/suite"
)

const secondsPerDay = 86400

var (
	owner, _       = crypto.GenerateKey()
	user, _        = crypto.GenerateKey()
	initialBalance = new(big.Int).Mul(common.Big32, new(big.Int).SetUint64(params.Ether)) // 10 kUSD
	baseDeposit    = new(big.Int).Mul(common.Big1, new(big.Int).SetUint64(params.Ether))  // 1 kUSD
	maxNumOracles  = common.Big256
	freezePeriod   = common.Big32
	oneDollar      = new(big.Int).Mul(common.Big1, new(big.Int).SetUint64(params.Ether)) // $1
)

type OracleMgrSuite struct {
	suite.Suite
	backend   *backends.SimulatedBackend
	oracleMgr *OracleMgr
}

func TestOracleMgrSuite(t *testing.T) {
	suite.Run(t, new(OracleMgrSuite))
}

func (suite *OracleMgrSuite) BeforeTest(suiteName, testName string) {
	if strings.Contains(testName, "TestDeploy") {
		return
	}

	req := suite.Require()

	alloc := make(core.GenesisAlloc)
	alloc[getAddress(owner)] = core.GenesisAccount{Balance: initialBalance}
	alloc[getAddress(user)] = core.GenesisAccount{Balance: initialBalance}

	backend := backends.NewSimulatedBackend(alloc)
	req.NotNil(backend)
	suite.backend = backend

	var finalMaxNumOracles *big.Int
	if testName == "TestGetMinimumDeposit_Full" {
		finalMaxNumOracles = common.Big1
	} else {
		finalMaxNumOracles = maxNumOracles
	}

	// OracleMgr instance
	transactOpts := bind.NewKeyedTransactor(owner)
	_, _, oracleMgr, err := DeployOracleMgr(transactOpts, backend, baseDeposit, finalMaxNumOracles, freezePeriod)
	req.NoError(err)
	req.NotNil(oracleMgr)
	suite.oracleMgr = oracleMgr

	suite.backend.Commit()
}

func (suite *OracleMgrSuite) TestDeployOracleMgr() {
	req := suite.Require()

	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		getAddress(owner): core.GenesisAccount{
			Balance: new(big.Int).Mul(new(big.Int).SetUint64(100), new(big.Int).SetUint64(params.Ether)),
		},
	})

	transactOpts := bind.NewKeyedTransactor(owner)
	_, _, oracleMgr, err := DeployOracleMgr(transactOpts, backend, baseDeposit, maxNumOracles, freezePeriod)
	req.NoError(err)
	req.NotNil(oracleMgr)

	backend.Commit()

	storedBaseDeposit, err := oracleMgr.BaseDeposit(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedBaseDeposit)
	req.Equal(baseDeposit, storedBaseDeposit)

	storedMaxNumOracles, err := oracleMgr.MaxNumOracles(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedMaxNumOracles)
	req.Equal(maxNumOracles, storedMaxNumOracles)

	storedFreezePeriod, err := oracleMgr.FreezePeriod(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedFreezePeriod)
	req.Equal(dtos(freezePeriod), storedFreezePeriod)

	storedPrice, err := oracleMgr.Price(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedPrice)

	req.Equal(oneDollar, storedPrice)
}

func (suite *OracleMgrSuite) TestDeployOracleMgr_MaxNumOraclesEqualZero() {
	req := suite.Require()

	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		getAddress(owner): core.GenesisAccount{
			Balance: new(big.Int).Mul(new(big.Int).SetUint64(100), new(big.Int).SetUint64(params.Ether)),
		},
	})

	transactOpts := bind.NewKeyedTransactor(owner)
	_, _, _, err := DeployOracleMgr(transactOpts, backend, baseDeposit, maxNumOracles, freezePeriod)
	req.NoError(err, "maximum number of oracles cannot be zero")
}

func (suite *OracleMgrSuite) TestGetMinimumDeposit_NotFull() {
	req := suite.Require()

	storedMinDeposit, err := suite.oracleMgr.GetMinimumDeposit(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedMinDeposit)
	req.Equal(baseDeposit, storedMinDeposit)
}

func (suite *OracleMgrSuite) TestGetMinimumDeposit_Full() {
	req := suite.Require()

	// register one oracle to fill the available place
	req.NoError(suite.registerOracle(user, baseDeposit))

	suite.backend.Commit()

	storedMinDeposit, err := suite.oracleMgr.GetMinimumDeposit(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedMinDeposit)
	req.Equal(new(big.Int).Add(baseDeposit, common.Big1), storedMinDeposit)
}

func (suite *OracleMgrSuite) TestRegisterOracle_WhenPaused() {
	req := suite.Require()

	pauseOpts := bind.NewKeyedTransactor(owner)
	_, err := suite.oracleMgr.Pause(pauseOpts)
	req.NoError(err)

	req.Error(suite.registerOracle(user, baseDeposit), "cannot register the oracle because the service is paused")
}

func (suite *OracleMgrSuite) TestRegisterOracle_Duplicate() {
	req := suite.Require()

	req.NoError(suite.registerOracle(user, baseDeposit))
	req.Error(suite.registerOracle(user, baseDeposit), "cannot register the same oracle twice")
}

func (suite *OracleMgrSuite) TestRegisterOracle_WithoutMinDeposit() {
	req := suite.Require()

	deposit := new(big.Int).Sub(baseDeposit, common.Big1)
	req.Error(suite.registerOracle(user, deposit), "requires the minimum deposit")
}

/*

func (suite *OracleMgrSuite) TestRegisterOracle() {
	req := suite.Require()

	req.NoError(suite.registerOracle(user, baseDeposit))

	suite.backend.Commit()


}

/*
func (suite *OracleMgrSuite) TestRegisterOracle_GreaterThan() {
	req := suite.Require()

	// deploy oracle manager with one spot available
	opts := bind.NewKeyedTransactor(suite.contractOwner)
	_, _, mgr, err := DeployOracleMgr(opts, suite.backend, suite.baseDeposit, common.Big1, suite.freezePeriod)
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
	_, _, mgr, err := DeployOracleMgr(ownerAuth, suite.backend, suite.baseDeposit, common.Big1, suite.freezePeriod)
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
	_, _, mgr, err := DeployOracleMgr(ownerAuth, suite.backend, suite.baseDeposit, common.Big1, suite.freezePeriod)
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
	_, _, mgr, err := DeployOracleMgr(ownerAuth, suite.backend, suite.baseDeposit, common.Big1, suite.freezePeriod)
	req.NoError(err)

	// deregister oracle
	userAuth := bind.NewKeyedTransactor(suite.randomUser)
	_, err = mgr.DeregisterOracle(userAuth)
	req.Equal(errAlwaysFailingTransaction, err)
}


func (suite *OracleMgrSuite) TestReleaseDeposits_WhenNotPaused() {
	req := suite.Require()

	// deploy oracle manager
	deployOpts := bind.NewKeyedTransactor(suite.contractOwner)
	_, _, mgr, err := DeployOracleMgr(deployOpts, suite.backend, suite.baseDeposit, common.Big32, suite.freezePeriod)
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
	_, _, mgr, err := DeployOracleMgr(deployOpts, suite.backend, suite.baseDeposit, common.Big32, suite.freezePeriod)
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



func (suite *OracleMgrSuite) TestAddPrice_Paused() {
	req := suite.Require()

	deploymentOpts := bind.NewKeyedTransactor(suite.contractOwner)
	_, _, mgr, err := DeployOracleMgr(deploymentOpts, suite.backend, suite.baseDeposit, suite.maxNumOracles, suite.freezePeriod)
	req.NoError(err)
	req.NotNil(mgr)

	deposit := suite.baseDeposit
	registrationOpts := bind.NewKeyedTransactor(suite.randomUser)
	registrationOpts.Value = deposit
	_, err = mgr.RegisterOracle(registrationOpts)
	req.NoError(err)

	pauseOpts := bind.NewKeyedTransactor(suite.contractOwner)
	_, err = mgr.Pause(pauseOpts)
	req.NoError(err)

	// must fail because the service is paused
	newPrice := common.Big1
	priceOpts := bind.NewKeyedTransactor(suite.randomUser)
	_, err = mgr.AddPrice(priceOpts, newPrice)
	req.Equal(errAlwaysFailingTransaction, err)
}

func (suite *OracleMgrSuite) TestAddPrice_InvalidUser() {
	req := suite.Require()

	deploymentOpts := bind.NewKeyedTransactor(suite.contractOwner)
	_, _, mgr, err := DeployOracleMgr(deploymentOpts, suite.backend, suite.baseDeposit, suite.maxNumOracles, suite.freezePeriod)
	req.NoError(err)
	req.NotNil(mgr)

	// must fail because the user is not an oracle
	newPrice := common.Big1
	priceOpts := bind.NewKeyedTransactor(suite.randomUser)
	_, err = mgr.AddPrice(priceOpts, newPrice)
	req.Equal(errAlwaysFailingTransaction, err)
}

func (suite *OracleMgrSuite) TestAddPrice_InvalidPrice() {
	req := suite.Require()

	deploymentOpts := bind.NewKeyedTransactor(suite.contractOwner)
	_, _, mgr, err := DeployOracleMgr(deploymentOpts, suite.backend, suite.baseDeposit, suite.maxNumOracles, suite.freezePeriod)
	req.NoError(err)
	req.NotNil(mgr)

	deposit := suite.baseDeposit
	registrationOpts := bind.NewKeyedTransactor(suite.randomUser)
	registrationOpts.Value = deposit
	_, err = mgr.RegisterOracle(registrationOpts)
	req.NoError(err)

	invalidPrice := common.Big0 // price must be > 0
	priceOpts := bind.NewKeyedTransactor(suite.randomUser)
	_, err = mgr.AddPrice(priceOpts, invalidPrice)
	req.Equal(errAlwaysFailingTransaction, err)
}

func (suite *OracleMgrSuite) TestAddPrice() {
	req := suite.Require()

	deploymentOpts := bind.NewKeyedTransactor(suite.contractOwner)
	_, _, mgr, err := DeployOracleMgr(deploymentOpts, suite.backend, suite.baseDeposit, suite.maxNumOracles, suite.freezePeriod)
	req.NoError(err)
	req.NotNil(mgr)

	deposit := suite.baseDeposit
	registrationOpts := bind.NewKeyedTransactor(suite.randomUser)
	registrationOpts.Value = deposit
	_, err = mgr.RegisterOracle(registrationOpts)
	req.NoError(err)

	newPrice := common.Big2
	priceOpts := bind.NewKeyedTransactor(suite.randomUser)
	_, err = mgr.AddPrice(priceOpts, newPrice)
	req.NoError(err)

	suite.backend.Commit()

	// price must be equal to newPrice
	storedPrice, err := mgr.Price(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedPrice)

	req.Equal(newPrice, storedPrice)
}


// kusd converts the value to kUSD units
func kusd(value *big.Int) *big.Int {
	return new(big.Int).Mul(value, new(big.Int).SetUint64(params.Ether))
}

*/

func (suite *OracleMgrSuite) registerOracle(user *ecdsa.PrivateKey, deposit *big.Int) error {
	transactOpts := bind.NewKeyedTransactor(user)
	transactOpts.Value = deposit
	_, err := suite.oracleMgr.RegisterOracle(transactOpts)
	return err
}

// dtos converts days to seconds
func dtos(days *big.Int) *big.Int {
	return new(big.Int).Mul(days, new(big.Int).SetUint64(secondsPerDay))
}

// getAddress return the address of the given private key
func getAddress(privateKey *ecdsa.PrivateKey) common.Address {
	return crypto.PubkeyToAddress(privateKey.PublicKey)
}
