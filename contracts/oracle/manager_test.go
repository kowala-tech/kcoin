package oracle

import (
	"context"
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

	// create backend
	alloc := make(core.GenesisAlloc)
	alloc[getAddress(owner)] = core.GenesisAccount{Balance: initialBalance}
	alloc[getAddress(user)] = core.GenesisAccount{Balance: initialBalance}
	backend := backends.NewSimulatedBackend(alloc)
	req.NotNil(backend)
	suite.backend = backend

	finalMaxNumOracles := maxNumOracles
	finalFreezePeriod := freezePeriod

	switch {
	case strings.Contains(testName, "_Full"):
		finalMaxNumOracles = common.Big1
	case testName == "TestReleaseDeposits_UnlockedDeposit":
		finalFreezePeriod = common.Big0
	}

	// OracleMgr instance
	oracleMgr, err := suite.deployOracleMgr(owner, backend, baseDeposit, finalMaxNumOracles, finalFreezePeriod)
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
	req.NotNil(backend)

	oracleMgr, err := suite.deployOracleMgr(owner, backend, baseDeposit, maxNumOracles, freezePeriod)
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
	req.NotNil(backend)

	_, err := suite.deployOracleMgr(owner, backend, baseDeposit, maxNumOracles, freezePeriod)
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

	deposit := baseDeposit
	req.NoError(suite.registerOracle(user, deposit))

	suite.backend.Commit()

	storedMinDeposit, err := suite.oracleMgr.GetMinimumDeposit(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedMinDeposit)
	req.Equal(new(big.Int).Add(deposit, common.Big1), storedMinDeposit)
}

func (suite *OracleMgrSuite) TestRegisterOracle_WhenPaused() {
	req := suite.Require()

	suite.pauseOracleMgr(owner)

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
	req.Error(suite.registerOracle(user, deposit), "registerOracle requires the minimum deposit")
}

func (suite *OracleMgrSuite) TestRegister_NotFull_GreaterThan() {
	req := suite.Require()

	req.NoError(suite.registerOracle(owner, baseDeposit))

	suite.backend.Commit()

	initialOracleCount, err := suite.oracleMgr.GetOracleCount(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(initialOracleCount)

	deposit := new(big.Int).Add(baseDeposit, common.Big1)
	req.NoError(suite.registerOracle(user, deposit))

	suite.backend.Commit()

	storedOracle := suite.getHighestBidder()
	req.NotZero(storedOracle)
	req.Equal(getAddress(user), storedOracle.Code)
	req.Equal(deposit, storedOracle.Deposit)

	storedDeposit := suite.getCurrentDeposit(user)
	req.NotZero(storedDeposit)
	req.Zero(storedDeposit.AvailableAt.Uint64())
	req.Equal(deposit, storedDeposit.Amount)

	finalOracleCount, err := suite.oracleMgr.GetOracleCount(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(finalOracleCount)
	req.Equal(new(big.Int).Add(initialOracleCount, common.Big1), finalOracleCount)
}

func (suite *OracleMgrSuite) TestRegister_NotFull_LessOrEqualTo() {
	req := suite.Require()

	req.NoError(suite.registerOracle(owner, baseDeposit))

	suite.backend.Commit()

	initialOracleCount, err := suite.oracleMgr.GetOracleCount(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(initialOracleCount)

	req.NoError(suite.registerOracle(user, baseDeposit))

	suite.backend.Commit()

	storedOracle := suite.getHighestBidder()
	req.NotZero(storedOracle)
	req.Equal(getAddress(owner), storedOracle.Code)
	req.Equal(baseDeposit, storedOracle.Deposit)

	storedDeposit := suite.getCurrentDeposit(owner)
	req.NotZero(storedDeposit)
	req.Zero(storedDeposit.AvailableAt.Uint64())
	req.Equal(baseDeposit, storedDeposit.Amount)

	finalOracleCount, err := suite.oracleMgr.GetOracleCount(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(finalOracleCount)
	req.Equal(new(big.Int).Add(initialOracleCount, common.Big1), finalOracleCount)
}

func (suite *OracleMgrSuite) TestRegister_Full_Replacement() {
	req := suite.Require()

	req.NoError(suite.registerOracle(owner, baseDeposit))

	suite.backend.Commit()

	initialOracleCount, err := suite.oracleMgr.GetOracleCount(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(initialOracleCount)

	minDeposit, err := suite.oracleMgr.GetMinimumDeposit(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(minDeposit)

	req.NoError(suite.registerOracle(user, minDeposit))

	suite.backend.Commit()

	storedOracle := suite.getHighestBidder()
	req.NotZero(storedOracle)
	req.Equal(getAddress(user), storedOracle.Code)
	req.Equal(minDeposit, storedOracle.Deposit)

	storedDeposit := suite.getCurrentDeposit(user)
	req.NotZero(storedDeposit)
	req.Zero(storedDeposit.AvailableAt.Uint64())
	req.Equal(minDeposit, storedDeposit.Amount)

	finalOracleCount, err := suite.oracleMgr.GetOracleCount(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(finalOracleCount)
	req.Equal(initialOracleCount, finalOracleCount)

}

func (suite *OracleMgrSuite) TestDeregisterOracle_WhenPaused() {
	req := suite.Require()

	suite.pauseOracleMgr(owner)

	req.Error(suite.deregisterOracle(user), "cannot deregister the oracle because the service is paused")
}

func (suite *OracleMgrSuite) TestDeregisterOracle_NotOracle() {
	req := suite.Require()

	suite.pauseOracleMgr(owner)

	req.Error(suite.deregisterOracle(user), "cannot deregister a non-oracle")
}

func (suite *OracleMgrSuite) TestDeregisterOracle() {
	req := suite.Require()

	req.NoError(suite.registerOracle(user, baseDeposit))
	req.NoError(suite.deregisterOracle(user))

	suite.backend.Commit()

	deposit := suite.getCurrentDeposit(user)
	req.NotNil(deposit)

	req.True(deposit.AvailableAt.Cmp(common.Big0) > 0)
}

func (suite *OracleMgrSuite) TestReleaseDeposits_WhenPaused() {
	req := suite.Require()

	suite.pauseOracleMgr(owner)

	req.Error(suite.releaseDeposits(user), "cannot release deposits because the service is paused")
}

func (suite *OracleMgrSuite) TestReleaseDeposits_NoDeposits() {
	req := suite.Require()

	initialBalance := suite.balanceOf(user)
	req.NotNil(initialBalance)

	req.NoError(suite.releaseDeposits(user))

	suite.backend.Commit()

	depositCount := suite.getDepositCount(user)
	req.NotNil(depositCount)
	req.Zero(depositCount.Uint64())

	finalBalance := suite.balanceOf(user)
	req.NotNil(finalBalance)
	// @NOTE (rgeraldes) - gas costs
	req.True(finalBalance.Cmp(initialBalance) < 0)
}

func (suite *OracleMgrSuite) TestReleaseDeposits_LockedDeposits() {
	req := suite.Require()

	initialBalance := suite.balanceOf(user)
	req.NotNil(initialBalance)

	// @NOTE (rgeraldes) - leave funds for the gas costs (1 kUSD)
	deposit := new(big.Int).Sub(initialBalance, new(big.Int).Sub(initialBalance, new(big.Int).Mul(common.Big1, new(big.Int).SetUint64(params.Ether))))
	req.NoError(suite.registerOracle(user, deposit))
	req.NoError(suite.deregisterOracle(user))
	req.NoError(suite.releaseDeposits(user))

	suite.backend.Commit()

	depositCount := suite.getDepositCount(user)
	req.NotNil(depositCount)
	req.Equal(common.Big1, depositCount)

	finalBalance := suite.balanceOf(user)
	req.NotNil(finalBalance)
	// @NOTE (rgeraldes) - final balance should be 1 kUSD minus the gas costs
	req.Zero(finalBalance.Cmp(common.Big1) < 0)
}

func (suite *OracleMgrSuite) TestReleaseDeposits_UnlockedDeposit() {
	req := suite.Require()

	initialBalance := suite.balanceOf(user)
	req.NotNil(initialBalance)

	// @NOTE (rgeraldes) - leave funds for the gas costs (1 kUSD)
	deposit := new(big.Int).Sub(initialBalance, new(big.Int).Sub(initialBalance, new(big.Int).Mul(common.Big1, new(big.Int).SetUint64(params.Ether))))
	req.NoError(suite.registerOracle(user, deposit))
	req.NoError(suite.deregisterOracle(user))

	suite.backend.Commit()

	req.NoError(suite.releaseDeposits(user))

	suite.backend.Commit()

	depositCount := suite.getDepositCount(user)
	req.NotNil(depositCount)
	req.Zero(depositCount.Uint64())

	finalBalance := suite.balanceOf(user)
	req.NotNil(finalBalance)
	// @NOTE (rgeraldes) - final balance should be the deposit + 1 kUSD - the gas costs
	req.True(finalBalance.Cmp(common.Big1) > 0)
}

func (suite *OracleMgrSuite) deployOracleMgr(user *ecdsa.PrivateKey, backend bind.ContractBackend, _baseDeposit *big.Int, _maxNumOracles *big.Int, _freezePeriod *big.Int) (*OracleMgr, error) {
	transactOpts := bind.NewKeyedTransactor(user)
	_, _, oracleMgr, err := DeployOracleMgr(transactOpts, backend, _baseDeposit, _maxNumOracles, _freezePeriod)
	return oracleMgr, err
}

func (suite *OracleMgrSuite) pauseOracleMgr(user *ecdsa.PrivateKey) {
	req := suite.Require()

	pauseOpts := bind.NewKeyedTransactor(user)
	_, err := suite.oracleMgr.Pause(pauseOpts)
	req.NoError(err)
}

func (suite *OracleMgrSuite) registerOracle(user *ecdsa.PrivateKey, deposit *big.Int) error {
	transactOpts := bind.NewKeyedTransactor(user)
	transactOpts.Value = deposit
	_, err := suite.oracleMgr.RegisterOracle(transactOpts)
	return err
}

func (suite *OracleMgrSuite) deregisterOracle(user *ecdsa.PrivateKey) error {
	transactOpts := bind.NewKeyedTransactor(user)
	_, err := suite.oracleMgr.DeregisterOracle(transactOpts)
	return err
}

func (suite *OracleMgrSuite) releaseDeposits(user *ecdsa.PrivateKey) error {
	transactOpts := bind.NewKeyedTransactor(user)
	_, err := suite.oracleMgr.ReleaseDeposits(transactOpts)
	return err
}

func (suite *OracleMgrSuite) balanceOf(user *ecdsa.PrivateKey) *big.Int {
	req := suite.Require()

	balance, err := suite.backend.BalanceAt(context.TODO(), getAddress(user), suite.backend.CurrentBlock().Number())
	req.NoError(err)
	req.NotNil(balance)

	return balance
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

func (suite *OracleMgrSuite) getDepositCount(user *ecdsa.PrivateKey) *big.Int {
	req := suite.Require()

	depositCount, err := suite.oracleMgr.GetDepositCount(&bind.CallOpts{From: getAddress(user)})
	req.NoError(err)
	req.NotNil(depositCount)

	return depositCount
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
