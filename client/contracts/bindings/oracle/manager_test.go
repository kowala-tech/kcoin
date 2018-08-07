package oracle_test

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"strings"
	"testing"

	"github.com/kowala-tech/kcoin/client/accounts/abi"
	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/accounts/abi/bind/backends"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/oracle"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/ownership"
	"github.com/kowala-tech/kcoin/client/core"
	"github.com/kowala-tech/kcoin/client/crypto"
	"github.com/kowala-tech/kcoin/client/knode/genesis"
	"github.com/kowala-tech/kcoin/client/params"
	"github.com/stretchr/testify/suite"
)

const secondsPerDay = 86400

var (
	user, _            = crypto.GenerateKey()
	superNode, _       = crypto.GenerateKey()
	userWithoutMUSD, _ = crypto.GenerateKey()
	governor, _        = crypto.GenerateKey()
	author, _          = crypto.HexToECDSA("bfef37ae9ac5d5e7ebbbefc19f4e1f572a7ca7aa0d28e527b7d62950951cc5eb")
	validatorMgrAddr   = common.HexToAddress("0x161ad311F1D66381C17641b1B73042a4CA731F9f")
	multiSigAddr       = common.HexToAddress("0xA143ac5ec5D95f16aFD5Fc3B09e0aDaf360ffC9e")
	oracleMgrAddr      = common.HexToAddress("0x2c3DA02A82D11D649857AaE537920D8cA368cAB5")
)

func getDefaultOpts() genesis.Options {
	baseDeposit := uint64(20)
	superNodeAmount := uint64(6000000)
	tokenHolder := genesis.TokenHolder{
		Address:   getAddress(user).Hex(),
		NumTokens: superNodeAmount,
	}

	opts := genesis.Options{
		Network: "test",
		Consensus: &genesis.ConsensusOpts{
			Engine:           "konsensus",
			MaxNumValidators: 10,
			FreezePeriod:     30,
			BaseDeposit:      baseDeposit,
			SuperNodeAmount:  superNodeAmount,
			Validators: []genesis.Validator{
				{
					Address: tokenHolder.Address,
					Deposit: tokenHolder.NumTokens,
				},
				{
					Address: getAddress(superNode).Hex(),
					Deposit: superNodeAmount,
				},
			},
			MiningToken: &genesis.MiningTokenOpts{
				Name:     "mUSD",
				Symbol:   "mUSD",
				Cap:      20000000,
				Decimals: 18,
				Holders:  []genesis.TokenHolder{tokenHolder, {Address: getAddress(superNode).Hex(), NumTokens: superNodeAmount}},
			},
		},
		Governance: &genesis.GovernanceOpts{
			Origin:           getAddress(author).Hex(),
			Governors:        []string{getAddress(governor).Hex()},
			NumConfirmations: 1,
		},
		DataFeedSystem: &genesis.DataFeedSystemOpts{
			MaxNumOracles: 10,
			FreezePeriod:  32,
			BaseDeposit:   1,
			Price: genesis.PriceOpts{
				InitialPrice:  1,
				SyncFrequency: 600,
				UpdatePeriod:  30,
			},
		},
		PrefundedAccounts: []genesis.PrefundedAccount{
			{
				Address: tokenHolder.Address,
				Balance: 10,
			},
			{
				Address: getAddress(superNode).Hex(),
				Balance: 10,
			},
			{
				Address: getAddress(governor).Hex(),
				Balance: 10,
			},
			{
				Address: getAddress(userWithoutMUSD).Hex(),
				Balance: 10,
			},
		},
	}

	return opts
}

type OracleMgrSuite struct {
	suite.Suite
	backend   *backends.SimulatedBackend
	opts      genesis.Options
	multiSig  *ownership.MultiSigWallet
	oracleMgr *oracle.OracleMgr
}

func TestOracleMgrSuite(t *testing.T) {
	suite.Run(t, new(OracleMgrSuite))
}

func (suite *OracleMgrSuite) BeforeTest(suiteName, testName string) {
	if strings.Contains(testName, "TestDeploy") {
		return
	}

	req := suite.Require()

	// create genesis
	opts := getDefaultOpts()
	req.NotNil(opts)

	switch {
	case strings.Contains(testName, "_Full"):
		opts.DataFeedSystem.MaxNumOracles = 1
	case testName == "TestReleaseDeposits_UnlockedDeposit":
		opts.DataFeedSystem.FreezePeriod = 0
	}
	suite.opts = opts

	genesis, err := genesis.Generate(opts)
	req.NoError(err)
	req.NotNil(genesis)

	backend := backends.NewSimulatedBackend(genesis.Alloc)
	req.NotNil(backend)
	suite.backend = backend

	// multiSig instance
	multiSig, err := ownership.NewMultiSigWallet(multiSigAddr, backend)
	req.NoError(err)
	req.NotNil(multiSig)
	suite.multiSig = multiSig

	// OracleMgr instance
	oracleMgr, err := oracle.NewOracleMgr(oracleMgrAddr, backend)
	req.NoError(err)
	req.NotNil(oracleMgr)
	suite.oracleMgr = oracleMgr
}

func (suite *OracleMgrSuite) TestDeployOracleMgr() {
	req := suite.Require()

	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		getAddress(governor): core.GenesisAccount{
			Balance: new(big.Int).Mul(new(big.Int).SetUint64(100), new(big.Int).SetUint64(params.Kcoin)),
		},
	})
	req.NotNil(backend)

	initialPrice := new(big.Int).SetUint64(1)
	baseDeposit := new(big.Int).SetUint64(100)
	maxNumOracles := new(big.Int).SetUint64(100)
	freezePeriod := new(big.Int).SetUint64(10)
	syncFrequency := new(big.Int).SetUint64(20)
	updatePeriod := new(big.Int).SetUint64(5)

	transactOpts := bind.NewKeyedTransactor(governor)
	_, _, mgr, err := oracle.DeployOracleMgr(transactOpts, backend, initialPrice, baseDeposit, maxNumOracles, freezePeriod, syncFrequency, updatePeriod, validatorMgrAddr)
	req.NoError(err)
	req.NotNil(mgr)

	backend.Commit()

	storedBaseDeposit, err := mgr.BaseDeposit(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedBaseDeposit)
	req.Equal(baseDeposit, storedBaseDeposit)

	storedMaxNumOracles, err := mgr.MaxNumOracles(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedMaxNumOracles)
	req.Equal(maxNumOracles, storedMaxNumOracles)

	storedFreezePeriod, err := mgr.FreezePeriod(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedFreezePeriod)
	req.Equal(dtos(freezePeriod), storedFreezePeriod)

	storedInitialPrice, err := mgr.Price(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedInitialPrice)
	req.Equal(initialPrice, storedInitialPrice)

	storedSyncFrequency, err := mgr.SyncFrequency(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedSyncFrequency)
	req.Equal(syncFrequency, storedSyncFrequency)

	storedUpdatePeriod, err := mgr.UpdatePeriod(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedUpdatePeriod)
	req.Equal(updatePeriod, storedUpdatePeriod)
}

func (suite *OracleMgrSuite) TestDeployOracleMgr_MaxNumOraclesEqualZero() {
	req := suite.Require()

	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		getAddress(governor): core.GenesisAccount{
			Balance: new(big.Int).Mul(new(big.Int).SetUint64(100), new(big.Int).SetUint64(params.Kcoin)),
		},
	})
	req.NotNil(backend)

	initialPrice := new(big.Int).SetUint64(1)
	baseDeposit := new(big.Int).SetUint64(100)
	maxNumOracles := common.Big0
	freezePeriod := new(big.Int).SetUint64(10)
	syncFrequency := new(big.Int).SetUint64(20)
	updatePeriod := new(big.Int).SetUint64(5)

	transactOpts := bind.NewKeyedTransactor(governor)
	_, _, _, err := oracle.DeployOracleMgr(transactOpts, backend, initialPrice, baseDeposit, maxNumOracles, freezePeriod, syncFrequency, updatePeriod, validatorMgrAddr)
	req.Error(err, "maximum number of oracles cannot be zero")
}

func (suite *OracleMgrSuite) TestDeployOracleMgr_InitialPriceEqualsZero() {
	req := suite.Require()

	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		getAddress(governor): core.GenesisAccount{
			Balance: new(big.Int).Mul(new(big.Int).SetUint64(100), new(big.Int).SetUint64(params.Kcoin)),
		},
	})
	req.NotNil(backend)

	initialPrice := common.Big0
	baseDeposit := new(big.Int).SetUint64(100)
	maxNumOracles := new(big.Int).SetUint64(100)
	freezePeriod := new(big.Int).SetUint64(10)
	syncFrequency := new(big.Int).SetUint64(20)
	updatePeriod := new(big.Int).SetUint64(5)

	transactOpts := bind.NewKeyedTransactor(governor)
	_, _, _, err := oracle.DeployOracleMgr(transactOpts, backend, initialPrice, baseDeposit, maxNumOracles, freezePeriod, syncFrequency, updatePeriod, validatorMgrAddr)
	req.Error(err, "initial price cannot be zero")
}

func (suite *OracleMgrSuite) TestDeployOracleMgr_SyncEnabled_UpdatePeriodEqualsZero() {
	req := suite.Require()

	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		getAddress(governor): core.GenesisAccount{
			Balance: new(big.Int).Mul(new(big.Int).SetUint64(100), new(big.Int).SetUint64(params.Kcoin)),
		},
	})
	req.NotNil(backend)

	initialPrice := new(big.Int).SetUint64(1)
	baseDeposit := new(big.Int).SetUint64(100)
	maxNumOracles := new(big.Int).SetUint64(100)
	freezePeriod := new(big.Int).SetUint64(10)
	syncFrequency := new(big.Int).SetUint64(20)
	updatePeriod := common.Big0

	transactOpts := bind.NewKeyedTransactor(governor)
	_, _, _, err := oracle.DeployOracleMgr(transactOpts, backend, initialPrice, baseDeposit, maxNumOracles, freezePeriod, syncFrequency, updatePeriod, validatorMgrAddr)
	req.Error(err, "update period cannot be zero if sync is enabled")
}

func (suite *OracleMgrSuite) TestDeployOracleMgr_SyncEnabled_UpdatePeriodGreaterThanSyncFrequency() {
	req := suite.Require()

	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		getAddress(governor): core.GenesisAccount{
			Balance: new(big.Int).Mul(new(big.Int).SetUint64(100), new(big.Int).SetUint64(params.Kcoin)),
		},
	})
	req.NotNil(backend)

	initialPrice := new(big.Int).SetUint64(1)
	baseDeposit := new(big.Int).SetUint64(100)
	maxNumOracles := new(big.Int).SetUint64(100)
	freezePeriod := new(big.Int).SetUint64(10)
	syncFrequency := new(big.Int).SetUint64(20)
	updatePeriod := new(big.Int).Add(syncFrequency, common.Big1)

	transactOpts := bind.NewKeyedTransactor(governor)
	_, _, _, err := oracle.DeployOracleMgr(transactOpts, backend, initialPrice, baseDeposit, maxNumOracles, freezePeriod, syncFrequency, updatePeriod, validatorMgrAddr)
	req.Error(err, "update period cannot be greater that sync frequency if sync is enabled")
}

func (suite *OracleMgrSuite) TestGetMinimumDeposit_NotFull() {
	req := suite.Require()

	storedMinDeposit, err := suite.oracleMgr.GetMinimumDeposit(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedMinDeposit)
	req.Equal(new(big.Int).Mul(new(big.Int).SetUint64(suite.opts.DataFeedSystem.BaseDeposit), big.NewInt(params.Kcoin)), storedMinDeposit)
}

func (suite *OracleMgrSuite) TestGetMinimumDeposit_Full() {
	req := suite.Require()

	deposit := new(big.Int).Mul(new(big.Int).SetUint64(suite.opts.DataFeedSystem.BaseDeposit), big.NewInt(params.Kcoin))
	req.NoError(suite.registerOracle(user, deposit))

	suite.backend.Commit()

	storedMinDeposit, err := suite.oracleMgr.GetMinimumDeposit(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedMinDeposit)
	req.Equal(new(big.Int).Add(deposit, common.Big1), storedMinDeposit)
}

func (suite *OracleMgrSuite) TestRegisterOracle_WhenPaused() {
	req := suite.Require()

	suite.pauseService()

	req.Error(suite.registerOracle(user, new(big.Int).Mul(new(big.Int).SetUint64(suite.opts.DataFeedSystem.BaseDeposit), big.NewInt(params.Kcoin))), "cannot register the oracle because the service is paused")
}

func (suite *OracleMgrSuite) TestRegisterOracle_Duplicate() {
	req := suite.Require()

	baseDeposit := new(big.Int).Mul(new(big.Int).SetUint64(suite.opts.DataFeedSystem.BaseDeposit), big.NewInt(params.Kcoin))
	req.NoError(suite.registerOracle(user, baseDeposit))
	req.Error(suite.registerOracle(user, baseDeposit), "cannot register the same oracle twice")
}

func (suite *OracleMgrSuite) TestRegisterOracle_WithoutMinDeposit() {
	req := suite.Require()

	deposit := new(big.Int).Sub(new(big.Int).Mul(new(big.Int).SetUint64(suite.opts.DataFeedSystem.BaseDeposit), big.NewInt(params.Kcoin)), common.Big1)
	req.Error(suite.registerOracle(user, deposit), "registerOracle requires the minimum deposit")
}

func (suite *OracleMgrSuite) TestRegisterOracle_NotSuperNode() {
	req := suite.Require()

	req.Error(suite.registerOracle(userWithoutMUSD, new(big.Int).Mul(new(big.Int).SetUint64(suite.opts.DataFeedSystem.BaseDeposit), big.NewInt(params.Kcoin))), "registerOracle requires a super node")
}

func (suite *OracleMgrSuite) TestRegister_NotFull_GreaterThan() {
	req := suite.Require()

	req.NoError(suite.registerOracle(superNode, new(big.Int).Mul(new(big.Int).SetUint64(suite.opts.DataFeedSystem.BaseDeposit), big.NewInt(params.Kcoin))))

	suite.backend.Commit()

	initialOracleCount, err := suite.oracleMgr.GetOracleCount(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(initialOracleCount)

	deposit := new(big.Int).Add(new(big.Int).Mul(new(big.Int).SetUint64(suite.opts.DataFeedSystem.BaseDeposit), big.NewInt(params.Kcoin)), common.Big1)
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

	baseDeposit := new(big.Int).Mul(new(big.Int).SetUint64(suite.opts.DataFeedSystem.BaseDeposit), big.NewInt(params.Kcoin))
	req.NoError(suite.registerOracle(superNode, baseDeposit))

	suite.backend.Commit()

	initialOracleCount, err := suite.oracleMgr.GetOracleCount(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(initialOracleCount)

	req.NoError(suite.registerOracle(user, baseDeposit))

	suite.backend.Commit()

	storedOracle := suite.getHighestBidder()
	req.NotZero(storedOracle)
	req.Equal(getAddress(superNode), storedOracle.Code)
	req.Equal(baseDeposit, storedOracle.Deposit)

	storedDeposit := suite.getCurrentDeposit(superNode)
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

	baseDeposit := new(big.Int).Mul(new(big.Int).SetUint64(suite.opts.DataFeedSystem.BaseDeposit), big.NewInt(params.Kcoin))
	req.NoError(suite.registerOracle(superNode, baseDeposit))

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

	suite.pauseService()

	req.Error(suite.deregisterOracle(user), "cannot deregister the oracle because the service is paused")
}

func (suite *OracleMgrSuite) TestDeregisterOracle_NotOracle() {
	req := suite.Require()

	suite.pauseService()

	req.Error(suite.deregisterOracle(user), "cannot deregister a non-oracle")
}

func (suite *OracleMgrSuite) TestDeregisterOracle() {
	req := suite.Require()

	baseDeposit := new(big.Int).Mul(new(big.Int).SetUint64(suite.opts.DataFeedSystem.BaseDeposit), big.NewInt(params.Kcoin))
	req.NoError(suite.registerOracle(user, baseDeposit))
	req.NoError(suite.deregisterOracle(user))

	suite.backend.Commit()

	deposit := suite.getCurrentDeposit(user)
	req.NotNil(deposit)

	req.True(deposit.AvailableAt.Cmp(common.Big0) > 0)
}

func (suite *OracleMgrSuite) TestReleaseDeposits_WhenPaused() {
	req := suite.Require()

	suite.pauseService()

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

	// @NOTE (rgeraldes) - leave funds for the gas costs (1 Kcoin)
	deposit := new(big.Int).Sub(initialBalance, new(big.Int).Sub(initialBalance, new(big.Int).Mul(common.Big1, new(big.Int).SetUint64(params.Kcoin))))
	req.NoError(suite.registerOracle(user, deposit))
	req.NoError(suite.deregisterOracle(user))
	req.NoError(suite.releaseDeposits(user))

	suite.backend.Commit()

	depositCount := suite.getDepositCount(user)
	req.NotNil(depositCount)
	req.Equal(common.Big1, depositCount)

	finalBalance := suite.balanceOf(user)
	req.NotNil(finalBalance)
	// @NOTE (rgeraldes) - final balance should be 1 Kcoin minus the gas costs
	req.Zero(finalBalance.Cmp(common.Big1) < 0)
}

func (suite *OracleMgrSuite) TestReleaseDeposits_UnlockedDeposit() {
	req := suite.Require()

	initialBalance := suite.balanceOf(user)
	req.NotNil(initialBalance)

	// @NOTE (rgeraldes) - leave funds for the gas costs (1 Kcoin)
	deposit := new(big.Int).Sub(initialBalance, new(big.Int).Sub(initialBalance, new(big.Int).Mul(common.Big1, new(big.Int).SetUint64(params.Kcoin))))
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
	// @NOTE (rgeraldes) - final balance should be the deposit + 1 Kcoin - the gas costs
	req.True(finalBalance.Cmp(common.Big1) > 0)
}

func (suite *OracleMgrSuite) pauseService() {
	req := suite.Require()

	// pause the service
	oracleMgrABI, err := abi.JSON(strings.NewReader(oracle.OracleMgrABI))
	req.NoError(err)
	req.NotNil(oracleMgrABI)

	pauseParams, err := oracleMgrABI.Pack("pause")
	req.NoError(err)
	req.NotZero(pauseParams)

	transactOpts := bind.NewKeyedTransactor(governor)
	_, err = suite.multiSig.SubmitTransaction(transactOpts, oracleMgrAddr, common.Big0, pauseParams)
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
