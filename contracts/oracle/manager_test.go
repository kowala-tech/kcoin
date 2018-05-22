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
	if strings.Contains(testName, "_Full") {
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

func (suite *OracleMgrSuite) TestRegister_NotFull_GreaterThan() {
	req := suite.Require()

	// @TODO - register another oracle first

	deposit := new(big.Int).Add(baseDeposit, common.Big1)
	req.NoError(suite.registerOracle(user, baseDeposit))

	storedOracle := suite.getHighestBidder()
	req.NotZero(storedOracle)
	req.Equal(getAddress(user), storedOracle.Code)
	req.Equal(deposit, storedOracle.Deposit)

	storedDeposit := suite.getCurrentDeposit(user)
	req.NotZero(storedDeposit)
	req.Zero(storedDeposit.AvailableAt.Uint64())
	req.Equal(deposit, storedDeposit.Amount)
}

func (suite *OracleMgrSuite) TestRegister_NotFull_LessOrEqualTo() {
	req := suite.Require()

	// @TODO - register another oracle first

	deposit := new(big.Int).Add(baseDeposit, common.Big1)
	req.NoError(suite.registerOracle(user, baseDeposit))

	storedOracle := suite.getHighestBidder()
	req.NotZero(storedOracle)
	req.Equal(getAddress(user), storedOracle.Code)
	req.Equal(deposit, storedOracle.Deposit)

	storedDeposit := suite.getCurrentDeposit(user)
	req.NotZero(storedDeposit)
	req.Zero(storedDeposit.AvailableAt.Uint64())
	req.Equal(deposit, storedDeposit.Amount)
}

func (suite *OracleMgrSuite) TestRegister_Full_Replacement() {
	req := suite.Require()

	// @TODO - register another oracle first

	deposit := new(big.Int).Add(baseDeposit, common.Big1)
	req.NoError(suite.registerOracle(user, baseDeposit))

	storedOracle := suite.getHighestBidder()
	req.NotZero(storedOracle)
	req.Equal(getAddress(user), storedOracle.Code)
	req.Equal(deposit, storedOracle.Deposit)

	storedDeposit := suite.getCurrentDeposit(user)
	req.NotZero(storedDeposit)
	req.Zero(storedDeposit.AvailableAt.Uint64())
	req.Equal(deposit, storedDeposit.Amount)
}

func (suite *OracleMgrSuite) TestDeregisterOracle_Full_Replacement() {
	req := suite.Require()

	// @TODO - register another oracle first

	deposit := new(big.Int).Add(baseDeposit, common.Big1)
	req.NoError(suite.registerOracle(user, baseDeposit))

	storedOracle := suite.getHighestBidder()
	req.NotZero(storedOracle)
	req.Equal(getAddress(user), storedOracle.Code)
	req.Equal(deposit, storedOracle.Deposit)

	storedDeposit := suite.getCurrentDeposit(user)
	req.NotZero(storedDeposit)
	req.Zero(storedDeposit.AvailableAt.Uint64())
	req.Equal(deposit, storedDeposit.Amount)
}

/*
func (suite *OracleMgrSuite) TestDeregisterOracle_WhenPaused() {
	req := suite.Require()

	pauseOpts := bind.NewKeyedTransactor(owner)
	_, err := suite.oracleMgr.Pause(pauseOpts)
	req.NoError(err)

	req.Error(suite.deregisterOracle(user), "cannot deregister the oracle because the service is paused")
}

func (suite *OracleMgrSuite) TestDeregisterOracle_NotOracle() {
	req := suite.Require()

	pauseOpts := bind.NewKeyedTransactor(owner)
	_, err := suite.oracleMgr.Pause(pauseOpts)
	req.NoError(err)

	req.Error(suite.deregisterOracle(user), "cannot deregister a non-oracle")
}

func (suite *OracleMgrSuite) TestDeregisterOracle() {
	req := suite.Require()

}

func (suite *OracleMgrSuite) TestReleaseDeposits_WhenPaused() {
	req := suite.Require()

	pauseOpts := bind.NewKeyedTransactor(owner)
	_, err := suite.oracleMgr.Pause(pauseOpts)
	req.NoError(err)

}

func (suite *OracleMgrSuite) TestReleaseDeposits_NoDeposits() {
	req := suite.Require()

	pauseOpts := bind.NewKeyedTransactor(owner)
	_, err := suite.oracleMgr.Pause(pauseOpts)
	req.NoError(err)
}

func (suite *OracleMgrSuite) TestReleaseDeposits_LockedDeposit() {
	req := suite.Require()

}

func (suite *OracleMgrSuite) TestReleaseDeposits_UnlockedDeposit() {
	req := suite.Require()

}
*/

func (suite *OracleMgrSuite) registerOracle(user *ecdsa.PrivateKey, deposit *big.Int) error {
	transactOpts := bind.NewKeyedTransactor(user)
	transactOpts.Value = deposit
	_, err := suite.oracleMgr.RegisterOracle(transactOpts)
	return err
}

func (suite *OracleMgrSuite) getHighestBidder() struct {
	Code    common.Address
	Deposit *big.Int
} {
	req := suite.Require()

	registration, err := suite.oracleMgr.GetOracleAtIndex(&bind.CallOpts{}, common.Big0)
	req.NoError(err)
	req.NotZero(registration)

	return registration
}

func (suite *OracleMgrSuite) getCurrentDeposit(user *ecdsa.PrivateKey) struct {
	Amount      *big.Int
	AvailableAt *big.Int
} {
	req := suite.Require()

	deposit, err := suite.oracleMgr.GetDepositAtIndex(&bind.CallOpts{From: getAddress(user)}, common.Big0)
	req.NoError(err)
	req.NotZero(deposit)

	return deposit
}

// dtos converts days to seconds
func dtos(days *big.Int) *big.Int {
	return new(big.Int).Mul(days, new(big.Int).SetUint64(secondsPerDay))
}

// getAddress return the address of the given private key
func getAddress(privateKey *ecdsa.PrivateKey) common.Address {
	return crypto.PubkeyToAddress(privateKey.PublicKey)
}
