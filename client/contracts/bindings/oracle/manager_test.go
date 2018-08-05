package oracle_test

import (
	"math/big"
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

type OracleMgrSuite struct {
	suite.Suite
	backend *backends.SimulatedBackend
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

	// deploy consensus
	mockSupernode := true // super node must be true to register the oracle
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

	// pause service
	oracleMgrContract.Pause(transactOpts)

	suite.backend.Commit()

	// register oracle must fail
	registerOpts := bind.NewKeyedTransactor(user)
	_, err = oracleMgrContract.RegisterOracle(registerOpts)
	req.Error(err, "service is paused")
}

func (suite *OracleMgrSuite) TestRegisterOracle_NotPaused_Duplicate() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// deploy consensus
	mockSupernode := true // super node must be true to register the oracle
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

	// register an oracle
	registerOpts := bind.NewKeyedTransactor(user)
	_, err = oracleMgrContract.RegisterOracle(registerOpts)
	req.NoError(err)

	suite.backend.Commit()

	// register the same oracle again
	_, err = oracleMgrContract.RegisterOracle(registerOpts)
	req.Error(err, "duplicate registration")
}

func (suite *OracleMgrSuite) TestRegisterOracle_NotPaused_NewCandidate_NotSuperNode() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// deploy consensus
	mockSupernode := false // super node set to false
	mockAddr, _, _, err := testfiles.DeployConsensusMock(transactOpts, suite.backend, mockSupernode)
	req.NoError(err)
	req.NotZero(mockAddr)

	suite.backend.Commit()

	// deploy oracle mgr contract
	maxNumOracles := big.NewInt(1)
	syncFreq := big.NewInt(900)
	updatePeriod := big.NewInt(50)
	_, _, oracleMgrContract, err := oracle.DeployOracleMgr(transactOpts, suite.backend, maxNumOracles, syncFreq, updatePeriod, mockAddr)
	req.NoError(err)
	req.NotNil(oracleMgrContract)

	suite.backend.Commit()

	// register an oracle
	registerOpts := bind.NewKeyedTransactor(user)
	_, err = oracleMgrContract.RegisterOracle(registerOpts)
	req.Error(err, "user is not a super node")
}

func (suite *OracleMgrSuite) TestRegisterOracle_NotPaused_NewCandidate_SuperNode_Full() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// deploy consensus
	mockSupernode := true
	mockAddr, _, _, err := testfiles.DeployConsensusMock(transactOpts, suite.backend, mockSupernode)
	req.NoError(err)
	req.NotZero(mockAddr)

	suite.backend.Commit()

	// deploy oracle mgr contract
	maxNumOracles := big.NewInt(1)
	syncFreq := big.NewInt(900)
	updatePeriod := big.NewInt(50)
	_, _, oracleMgrContract, err := oracle.DeployOracleMgr(transactOpts, suite.backend, maxNumOracles, syncFreq, updatePeriod, mockAddr)
	req.NoError(err)
	req.NotNil(oracleMgrContract)

	suite.backend.Commit()

	// register a new oracle to match the max number of oracles
	registerOpts := bind.NewKeyedTransactor(owner)
	_, err = oracleMgrContract.RegisterOracle(registerOpts)
	req.NoError(err)

	// register a new oracle
	registerOpts = bind.NewKeyedTransactor(user)
	_, err = oracleMgrContract.RegisterOracle(registerOpts)
	req.Error(err, "no positions available")
}

func (suite *OracleMgrSuite) TestRegisterOracle_NotPaused_NewCandidate_SuperNode_NotFull() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// deploy consensus
	mockSupernode := true
	mockAddr, _, _, err := testfiles.DeployConsensusMock(transactOpts, suite.backend, mockSupernode)
	req.NoError(err)
	req.NotZero(mockAddr)

	suite.backend.Commit()

	// deploy oracle mgr contract
	maxNumOracles := big.NewInt(1)
	syncFreq := big.NewInt(900)
	updatePeriod := big.NewInt(50)
	_, _, oracleMgrContract, err := oracle.DeployOracleMgr(transactOpts, suite.backend, maxNumOracles, syncFreq, updatePeriod, mockAddr)
	req.NoError(err)
	req.NotNil(oracleMgrContract)

	suite.backend.Commit()

	// register a new oracle
	registerOpts := bind.NewKeyedTransactor(user)
	_, err = oracleMgrContract.RegisterOracle(registerOpts)
	req.NoError(err)
}

func (suite *OracleMgrSuite) TestDeregisterOracle_WhenPaused() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// deploy consensus
	mockSupernode := true // super node must be true to register the oracle
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

	// register oracle
	registerOpts := bind.NewKeyedTransactor(user)
	_, err = oracleMgrContract.RegisterOracle(registerOpts)
	req.NoError(err)

	suite.backend.Commit()

	// pause service
	oracleMgrContract.Pause(transactOpts)

	suite.backend.Commit()

	// deregister oracle
	deregisterOpts := bind.NewKeyedTransactor(user)
	_, err = oracleMgrContract.DeregisterOracle(deregisterOpts)
	req.Error(err, "service is paused")
}

func (suite *OracleMgrSuite) TestDeregisterOracle_NotPaused_NotOracle() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// deploy consensus
	mockSupernode := true // super node must be true to register the oracle
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

	// deregister oracle
	deregisterOpts := bind.NewKeyedTransactor(user)
	_, err = oracleMgrContract.DeregisterOracle(deregisterOpts)
	req.Error(err, "the user is not an oracle")
}

func (suite *OracleMgrSuite) TestDeregisterOracle_NotPaused_Oracle() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// deploy consensus
	mockSupernode := true // super node must be true to register the oracle
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

	// register oracle
	registerOpts := bind.NewKeyedTransactor(user)
	_, err = oracleMgrContract.RegisterOracle(registerOpts)
	req.NoError(err)

	suite.backend.Commit()

	// deregister oracle
	deregisterOpts := bind.NewKeyedTransactor(user)
	_, err = oracleMgrContract.DeregisterOracle(deregisterOpts)
	req.NoError(err)

	suite.backend.Commit()

	count, err := oracleMgrContract.GetOracleCount(&bind.CallOpts{})
	req.NoError(err)
	req.Zero(count.Uint64())
}

func (suite *OracleMgrSuite) TestSubmitPrice_WhenPaused() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// deploy consensus
	mockSupernode := true // super node must be true to register the oracle
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

	// register oracle
	registerOpts := bind.NewKeyedTransactor(user)
	_, err = oracleMgrContract.RegisterOracle(registerOpts)
	req.NoError(err)

	suite.backend.Commit()

	// pause service
	oracleMgrContract.Pause(transactOpts)

	suite.backend.Commit()

	priceOpts := bind.NewKeyedTransactor(user)
	_, err = oracleMgrContract.SubmitPrice(priceOpts, common.Big2)
	req.Error(err, "service is paused")
}

func (suite *OracleMgrSuite) TestSubmitPrice_NotPaused_NotOracle() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// deploy consensus
	mockSupernode := true // super node must be true to register the oracle
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

	priceOpts := bind.NewKeyedTransactor(user)
	_, err = oracleMgrContract.SubmitPrice(priceOpts, common.Big2)
	req.Error(err, "the user is not an oracle")
}

func (suite *OracleMgrSuite) TestSubmitPrice_NotPaused_Oracle() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// deploy consensus
	mockSupernode := true // super node must be true to register the oracle
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

	// register oracle
	registerOpts := bind.NewKeyedTransactor(user)
	_, err = oracleMgrContract.RegisterOracle(registerOpts)
	req.NoError(err)

	suite.backend.Commit()

	// submit price
	priceOpts := bind.NewKeyedTransactor(user)
	_, err = oracleMgrContract.SubmitPrice(priceOpts, common.Big2)
	req.NoError(err)
}

func (suite *OracleMgrSuite) TestSubmitPrice_NotPaused_Oracle_SubmitPriceTwiceSameRound() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// deploy consensus
	mockSupernode := true // super node must be true to register the oracle
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

	// register oracle
	registerOpts := bind.NewKeyedTransactor(user)
	_, err = oracleMgrContract.RegisterOracle(registerOpts)
	req.NoError(err)

	suite.backend.Commit()

	// submit price
	priceOpts := bind.NewKeyedTransactor(user)
	_, err = oracleMgrContract.SubmitPrice(priceOpts, common.Big2)
	req.NoError(err)

	suite.backend.Commit()

	// submit price
	_, err = oracleMgrContract.SubmitPrice(priceOpts, common.Big2)
	req.Error(err, "cannot submit a price twice in the same round")
}

// dtos converts days to seconds
func dtos(days *big.Int) *big.Int {
	return new(big.Int).Mul(days, new(big.Int).SetUint64(secondsPerDay))
}
