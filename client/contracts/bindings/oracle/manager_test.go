package oracle_test

import (
	"math/big"
	"strings"
	"testing"

	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/accounts/abi/bind/backends"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/oracle"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/oracle/testfiles"
	"github.com/kowala-tech/kcoin/client/core"
	"github.com/kowala-tech/kcoin/client/crypto"
	"github.com/kowala-tech/kcoin/client/params"
	"github.com/stretchr/testify/suite"
)

const secondsPerDay = 86400

var (
	owner, _       = crypto.GenerateKey()
	user, _        = crypto.GenerateKey()
	initialBalance = new(big.Int).Mul(new(big.Int).SetUint64(100), new(big.Int).SetUint64(params.Kcoin))
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
	oracleMgr *oracle.OracleMgr
}

func TestOracleMgrSuite(t *testing.T) {
	suite.Run(t, new(OracleMgrSuite))
}

func (suite *OracleMgrSuite) BeforeTest(suiteName, testName string) {
	req := suite.Require()

	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		crypto.PubkeyToAddress(owner.PublicKey): core.GenesisAccount{
			Balance: initialBalance,
		},
		crypto.PubkeyToAddress(user.PublicKey): core.GenesisAccount{
			Balance: initialBalance,
		},
	})
	req.NotNil(backend)
	suite.backend = backend

	if strings.Contains(testName, "TestDeploy") {
		return
	}

	mockSuperNode := true
	maxNumOracles := 50
	switch {
	case strings.Contains(testName, "_NotSuperNode"):
		mockSuperNode = false
		fallthrough
	case strings.Contains(testName, "_Full"):
		maxNumOracles = 1
	}

	transactOpts := bind.NewKeyedTransactor(owner)

	// deploy consensus
	mockAddr, _, _, err := testfiles.DeployConsensusMock(transactOpts, suite.backend, mockSuperNode)
	req.NoError(err)
	req.NotZero(mockAddr)

	suite.backend.Commit()

	// deploy oracle mgr contract
	syncFreq := big.NewInt(900)
	updatePeriod := big.NewInt(50)
	_, _, oracleMgrContract, err := oracle.DeployOracleMgr(transactOpts, suite.backend, big.NewInt(int64(maxNumOracles)), syncFreq, updatePeriod, mockAddr)
	req.NoError(err)
	req.NotNil(oracleMgrContract)
	suite.oracleMgr = oracleMgrContract

	suite.backend.Commit()

}

func (suite *OracleMgrSuite) TestDeploy() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// deploy consensus
	mockSupernode := false
	mockAddr, _, _, err := testfiles.DeployConsensusMock(transactOpts, suite.backend, mockSupernode)
	req.NoError(err)
	req.NotZero(mockAddr)

	suite.backend.Commit()

	// deploy oracle mgr contract
	maxNumOracles := big.NewInt(50)
	syncFreq := big.NewInt(900)
	updatePeriod := big.NewInt(50)
	_, _, oracleMgrContract, err := oracle.DeployOracleMgr(transactOpts, suite.backend, maxNumOracles, syncFreq, updatePeriod, mockAddr)
	req.NoError(err)
	req.NotNil(oracleMgrContract)

	suite.backend.Commit()

	storedMaxNumOracles, err := oracleMgrContract.MaxNumOracles(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedMaxNumOracles)
	req.Equal(maxNumOracles, storedMaxNumOracles)

	storedSyncFreq, err := oracleMgrContract.SyncFrequency(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedSyncFreq)
	req.Equal(syncFreq, storedSyncFreq)

	storedUpdatePeriod, err := oracleMgrContract.UpdatePeriod(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedUpdatePeriod)
	req.Equal(updatePeriod, storedUpdatePeriod)
}

func (suite *OracleMgrSuite) TestDeploy_MaxNumOraclesEqualZero() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// deploy consensus
	mockSupernode := false
	mockAddr, _, _, err := testfiles.DeployConsensusMock(transactOpts, suite.backend, mockSupernode)
	req.NoError(err)
	req.NotZero(mockAddr)

	suite.backend.Commit()

	// deploy oracle mgr contract
	maxNumOracles := common.Big0
	syncFreq := big.NewInt(900)
	updatePeriod := big.NewInt(50)
	_, _, _, err = oracle.DeployOracleMgr(transactOpts, suite.backend, maxNumOracles, syncFreq, updatePeriod, mockAddr)
	req.Error(err, "max number of oracles must be greater than 0")
}

func (suite *OracleMgrSuite) TestDeploy_SyncFreqGreaterZero_UpdatePeriodZero() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// deploy consensus
	mockSupernode := false
	mockAddr, _, _, err := testfiles.DeployConsensusMock(transactOpts, suite.backend, mockSupernode)
	req.NoError(err)
	req.NotZero(mockAddr)

	suite.backend.Commit()

	// deploy oracle mgr contract
	maxNumOracles := big.NewInt(50)
	syncFreq := big.NewInt(900)
	updatePeriod := big.NewInt(0)
	_, _, _, err = oracle.DeployOracleMgr(transactOpts, suite.backend, maxNumOracles, syncFreq, updatePeriod, mockAddr)
	req.Error(err, "update period must be greater than 0 when sync is enabled")
}

func (suite *OracleMgrSuite) TestDeploy_SyncFreqGreaterZero_UpdatePeriodGreaterSyncFreq() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// deploy consensus
	mockSupernode := false
	mockAddr, _, _, err := testfiles.DeployConsensusMock(transactOpts, suite.backend, mockSupernode)
	req.NoError(err)
	req.NotZero(mockAddr)

	suite.backend.Commit()

	// deploy oracle mgr contract
	maxNumOracles := big.NewInt(50)
	syncFreq := big.NewInt(900)
	updatePeriod := big.NewInt(1000)
	_, _, _, err = oracle.DeployOracleMgr(transactOpts, suite.backend, maxNumOracles, syncFreq, updatePeriod, mockAddr)
	req.Error(err, "update period must be less or equal than sync freq")
}

func (suite *OracleMgrSuite) TestRegisterOracle_WhenPaused() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// pause service
	suite.oracleMgr.Pause(transactOpts)

	suite.backend.Commit()

	// register oracle must fail
	registerOpts := bind.NewKeyedTransactor(user)
	_, err := suite.oracleMgr.RegisterOracle(registerOpts)
	req.Error(err, "service is paused")
}

func (suite *OracleMgrSuite) TestRegisterOracle_NotPaused_Duplicate() {
	req := suite.Require()

	// register an oracle
	registerOpts := bind.NewKeyedTransactor(user)
	_, err := suite.oracleMgr.RegisterOracle(registerOpts)
	req.NoError(err)

	suite.backend.Commit()

	// register the same oracle again
	_, err = suite.oracleMgr.RegisterOracle(registerOpts)
	req.Error(err, "duplicate registration")
}

func (suite *OracleMgrSuite) TestRegisterOracle_NotPaused_NewCandidate_NotSuperNode() {
	req := suite.Require()

	// register an oracle
	registerOpts := bind.NewKeyedTransactor(user)
	_, err := suite.oracleMgr.RegisterOracle(registerOpts)
	req.Error(err, "user is not a super node")
}

func (suite *OracleMgrSuite) TestRegisterOracle_NotPaused_NewCandidate_SuperNode_Full() {
	req := suite.Require()

	// register a new oracle to match the max number of oracles
	registerOpts := bind.NewKeyedTransactor(owner)
	_, err := suite.oracleMgr.RegisterOracle(registerOpts)
	req.NoError(err)

	// register a new oracle
	registerOpts = bind.NewKeyedTransactor(user)
	_, err = suite.oracleMgr.RegisterOracle(registerOpts)
	req.Error(err, "no positions available")
}

func (suite *OracleMgrSuite) TestRegisterOracle_NotPaused_NewCandidate_SuperNode_NotFull() {
	req := suite.Require()

	// register a new oracle
	registerOpts := bind.NewKeyedTransactor(user)
	_, err := suite.oracleMgr.RegisterOracle(registerOpts)
	req.NoError(err)
}

func (suite *OracleMgrSuite) TestDeregisterOracle_WhenPaused() {
	req := suite.Require()

	// register oracle
	registerOpts := bind.NewKeyedTransactor(user)
	_, err := suite.oracleMgr.RegisterOracle(registerOpts)
	req.NoError(err)

	suite.backend.Commit()

	// pause service
	pauseOpts := bind.NewKeyedTransactor(owner)
	suite.oracleMgr.Pause(pauseOpts)

	suite.backend.Commit()

	// deregister oracle
	deregisterOpts := bind.NewKeyedTransactor(user)
	_, err = suite.oracleMgr.DeregisterOracle(deregisterOpts)
	req.Error(err, "service is paused")
}

func (suite *OracleMgrSuite) TestDeregisterOracle_NotPaused_NotOracle() {
	req := suite.Require()

	// deregister oracle
	deregisterOpts := bind.NewKeyedTransactor(user)
	_, err := suite.oracleMgr.DeregisterOracle(deregisterOpts)
	req.Error(err, "the user is not an oracle")
}

func (suite *OracleMgrSuite) TestDeregisterOracle_NotPaused_Oracle() {
	req := suite.Require()

	// register oracle
	registerOpts := bind.NewKeyedTransactor(user)
	_, err := suite.oracleMgr.RegisterOracle(registerOpts)
	req.NoError(err)

	suite.backend.Commit()

	// deregister oracle
	deregisterOpts := bind.NewKeyedTransactor(user)
	_, err = suite.oracleMgr.DeregisterOracle(deregisterOpts)
	req.NoError(err)

	suite.backend.Commit()

	// oracle count must be zero
	count, err := suite.oracleMgr.GetOracleCount(&bind.CallOpts{})
	req.NoError(err)
	req.Zero(count.Uint64())
}

func (suite *OracleMgrSuite) TestSubmitPrice_WhenPaused() {
	req := suite.Require()

	// register oracle
	registerOpts := bind.NewKeyedTransactor(user)
	_, err := suite.oracleMgr.RegisterOracle(registerOpts)
	req.NoError(err)

	suite.backend.Commit()

	// pause service
	pauseOpts := bind.NewKeyedTransactor(owner)
	suite.oracleMgr.Pause(pauseOpts)

	suite.backend.Commit()

	priceOpts := bind.NewKeyedTransactor(user)
	_, err = suite.oracleMgr.SubmitPrice(priceOpts, common.Big2)
	req.Error(err, "service is paused")
}

func (suite *OracleMgrSuite) TestSubmitPrice_NotPaused_NotOracle() {
	req := suite.Require()

	priceOpts := bind.NewKeyedTransactor(user)
	_, err := suite.oracleMgr.SubmitPrice(priceOpts, common.Big2)
	req.Error(err, "the user is not an oracle")
}

func (suite *OracleMgrSuite) TestSubmitPrice_NotPaused_Oracle() {
	req := suite.Require()

	// register oracle
	registerOpts := bind.NewKeyedTransactor(user)
	_, err := suite.oracleMgr.RegisterOracle(registerOpts)
	req.NoError(err)

	suite.backend.Commit()

	// submit price
	priceOpts := bind.NewKeyedTransactor(user)
	_, err = suite.oracleMgr.SubmitPrice(priceOpts, common.Big2)
	req.NoError(err)
}

func (suite *OracleMgrSuite) TestSubmitPrice_NotPaused_Oracle_SubmitPriceTwiceSameRound() {
	req := suite.Require()

	// register oracle
	registerOpts := bind.NewKeyedTransactor(user)
	_, err := suite.oracleMgr.RegisterOracle(registerOpts)
	req.NoError(err)

	suite.backend.Commit()

	// submit price
	priceOpts := bind.NewKeyedTransactor(user)
	_, err = suite.oracleMgr.SubmitPrice(priceOpts, common.Big2)
	req.NoError(err)

	suite.backend.Commit()

	// submit price
	_, err = suite.oracleMgr.SubmitPrice(priceOpts, common.Big2)
	req.Error(err, "cannot submit a price twice in the same round")
}

// dtos converts days to seconds
func dtos(days *big.Int) *big.Int {
	return new(big.Int).Mul(days, new(big.Int).SetUint64(secondsPerDay))
}
