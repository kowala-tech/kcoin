package oracle_test

import (
	"math/big"
	"strings"
	"testing"

	"crypto/ecdsa"

	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/accounts/abi/bind/backends"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/oracle/testfiles"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/utils"
	"github.com/kowala-tech/kcoin/client/core"
	"github.com/kowala-tech/kcoin/client/crypto"
	"github.com/kowala-tech/kcoin/client/params"
	"github.com/stretchr/testify/suite"
)

var (
	owner, _       = ecdsa.GenerateKey(crypto.S256(), strings.NewReader("ownerAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")) // Needed deterministic in order to link NameHash lib in contract.
	user, _        = ecdsa.GenerateKey(crypto.S256(), strings.NewReader("userAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"))
	initialBalance = new(big.Int).Mul(new(big.Int).SetUint64(100), new(big.Int).SetUint64(params.Kcoin))
)

type OracleMgrSuite struct {
	utils.ContractTestSuite
	oracleMgr *testfiles.OracleMgr
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
	suite.Backend = backend

	if strings.Contains(testName, "TestDeploy") {
		return
	}

	isSuperNode := true
	maxNumOracles := 50
	switch {
	case strings.Contains(testName, "_NotSuperNode"):
		isSuperNode = false
		fallthrough
	case strings.Contains(testName, "_Full"):
		maxNumOracles = 1
	}

	transactOpts := bind.NewKeyedTransactor(owner)

	suite.DeployStringsLibrary(transactOpts)
	suite.DeployNameHashLibrary(transactOpts)
	consensusAddr := suite.DeployConsensusMock(transactOpts, isSuperNode)
	resolverMockAddress := suite.DeployResolverMock(transactOpts, consensusAddr)

	// deploy oracle mgr contract
	syncFreq := big.NewInt(900)
	updatePeriod := big.NewInt(50)
	_, _, oracleMgrContract, err := testfiles.DeployOracleMgr(transactOpts, suite.Backend, big.NewInt(int64(maxNumOracles)), syncFreq, updatePeriod, resolverMockAddress)
	req.NoError(err)
	req.NotNil(oracleMgrContract)
	suite.oracleMgr = oracleMgrContract

	suite.Backend.Commit()
}

func (suite *OracleMgrSuite) TestDeploy() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	suite.DeployStringsLibrary(transactOpts)
	suite.DeployNameHashLibrary(transactOpts)
	consensusAddr := suite.DeployConsensusMock(transactOpts, false)
	suite.DeployResolverMock(transactOpts, consensusAddr)

	//deploy oracle mgr contract
	maxNumOracles := big.NewInt(50)
	syncFreq := big.NewInt(900)
	updatePeriod := big.NewInt(50)
	_, _, oracleMgrContract, err := testfiles.DeployOracleMgr(transactOpts, suite.Backend, maxNumOracles, syncFreq, updatePeriod, suite.ResolverMockAddr)
	req.NoError(err)
	req.NotNil(oracleMgrContract)

	suite.Backend.Commit()

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

	// deploy oracle mgr contract
	maxNumOracles := common.Big0
	syncFreq := big.NewInt(900)
	updatePeriod := big.NewInt(50)
	_, _, _, err := testfiles.DeployOracleMgr(transactOpts, suite.Backend, maxNumOracles, syncFreq, updatePeriod, suite.ResolverMockAddr)
	req.Error(err, "max number of oracles must be greater than 0")
}

func (suite *OracleMgrSuite) TestDeploy_SyncFreqGreaterZero_UpdatePeriodZero() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// deploy oracle mgr contract
	maxNumOracles := big.NewInt(50)
	syncFreq := big.NewInt(900)
	updatePeriod := big.NewInt(0)
	_, _, _, err := testfiles.DeployOracleMgr(transactOpts, suite.Backend, maxNumOracles, syncFreq, updatePeriod, suite.ResolverMockAddr)
	req.Error(err, "update period must be greater than 0 when sync is enabled")
}

func (suite *OracleMgrSuite) TestDeploy_SyncFreqGreaterZero_UpdatePeriodGreaterSyncFreq() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// deploy oracle mgr contract
	maxNumOracles := big.NewInt(50)
	syncFreq := big.NewInt(900)
	updatePeriod := big.NewInt(1000)
	_, _, _, err := testfiles.DeployOracleMgr(transactOpts, suite.Backend, maxNumOracles, syncFreq, updatePeriod, suite.ResolverMockAddr)
	req.Error(err, "update period must be less or equal than sync freq")
}

func (suite *OracleMgrSuite) TestRegisterOracle_Paused() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// pause service
	_, err := suite.oracleMgr.Pause(transactOpts)
	req.NoError(err)
	suite.Backend.Commit()

	// register oracle must fail
	_, err = suite.oracleMgr.RegisterOracle(transactOpts, crypto.PubkeyToAddress(user.PublicKey))
	req.Error(err, "service is paused")
}

func (suite *OracleMgrSuite) TestRegisterOracle_NotPaused_NotOwner() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(user)

	// register oracle must fail
	_, err := suite.oracleMgr.RegisterOracle(transactOpts, crypto.PubkeyToAddress(owner.PublicKey))
	req.Error(err, "not owner")
}

func (suite *OracleMgrSuite) TestRegisterOracle_NotPaused_Owner_Duplicate() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// register an oracle
	_, err := suite.oracleMgr.RegisterOracle(transactOpts, crypto.PubkeyToAddress(user.PublicKey))
	req.NoError(err)

	suite.Backend.Commit()

	// register the same oracle again
	_, err = suite.oracleMgr.RegisterOracle(transactOpts, crypto.PubkeyToAddress(user.PublicKey))
	req.Error(err, "duplicate registration")
}

func (suite *OracleMgrSuite) TestRegisterOracle_NotPaused_Owner_NewCandidate_Full() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// register an oracle
	_, err := suite.oracleMgr.RegisterOracle(transactOpts, crypto.PubkeyToAddress(owner.PublicKey))
	req.NoError(err)

	suite.Backend.Commit()

	// register the same oracle again
	_, err = suite.oracleMgr.RegisterOracle(transactOpts, crypto.PubkeyToAddress(user.PublicKey))
	req.Error(err, "duplicate registration")
}

func (suite *OracleMgrSuite) TestRegisterOracle_NotPaused_Owner_NewCandidate_NotFull() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// register an oracle
	userAddr := crypto.PubkeyToAddress(user.PublicKey)
	_, err := suite.oracleMgr.RegisterOracle(transactOpts, userAddr)
	req.NoError(err)

	suite.Backend.Commit()

	oracleCount, err := suite.oracleMgr.GetOracleCount(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(oracleCount)
	req.Equal(common.Big1, oracleCount)

	storedOracle, err := suite.oracleMgr.GetOracleAtIndex(&bind.CallOpts{}, common.Big0)
	req.NoError(err)
	req.NotZero(storedOracle)
	req.Equal(storedOracle, userAddr)
}

func (suite *OracleMgrSuite) TestDeregisterOracle_Paused() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)
	userAddr := crypto.PubkeyToAddress(user.PublicKey)

	// register oracle
	_, err := suite.oracleMgr.RegisterOracle(transactOpts, userAddr)
	req.NoError(err)

	suite.Backend.Commit()

	// pause service
	_, err = suite.oracleMgr.Pause(transactOpts)
	req.NoError(err)
	suite.Backend.Commit()

	// deregister oracle
	_, err = suite.oracleMgr.DeregisterOracle(transactOpts, userAddr)
	req.Error(err, "service is paused")
}

func (suite *OracleMgrSuite) TestDeregisterOracle_NotPaused_NotOwner() {
	req := suite.Require()

	// deregister oracle
	transactOpts := bind.NewKeyedTransactor(user)
	_, err := suite.oracleMgr.DeregisterOracle(transactOpts, crypto.PubkeyToAddress(owner.PublicKey))
	req.Error(err, "not owner")
}

func (suite *OracleMgrSuite) TestDeregisterOracle_NotPaused_Owner_NotOracle() {
	req := suite.Require()

	// deregister oracle
	transactOpts := bind.NewKeyedTransactor(owner)
	_, err := suite.oracleMgr.DeregisterOracle(transactOpts, crypto.PubkeyToAddress(user.PublicKey))
	req.Error(err, "the user is not an oracle")
}

func (suite *OracleMgrSuite) TestDeregisterOracle_NotPaused_Owner_Oracle() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)
	userAddr := crypto.PubkeyToAddress(user.PublicKey)

	// register oracle
	_, err := suite.oracleMgr.RegisterOracle(transactOpts, userAddr)
	req.NoError(err)

	suite.Backend.Commit()

	// deregister oracle
	_, err = suite.oracleMgr.DeregisterOracle(transactOpts, userAddr)
	req.NoError(err)
}

/*
func (suite *OracleMgrSuite) TestRegisterOracle_Paused() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// pause service
	_, err := suite.oracleMgr.Pause(transactOpts)
	req.NoError(err)
	suite.Backend.Commit()

	// register oracle must fail
	registerOpts := bind.NewKeyedTransactor(user)
	_, err = suite.oracleMgr.RegisterOracle(registerOpts)
	req.Error(err, "service is paused")
}

func (suite *OracleMgrSuite) TestRegisterOracle_NotPaused_Duplicate() {
	req := suite.Require()

	// register an oracle
	registerOpts := bind.NewKeyedTransactor(user)
	_, err := suite.oracleMgr.RegisterOracle(registerOpts)
	req.NoError(err)

	suite.Backend.Commit()

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

func (suite *OracleMgrSuite) TestDeregisterOracle_Paused() {
	req := suite.Require()

	// register oracle
	registerOpts := bind.NewKeyedTransactor(user)
	_, err := suite.oracleMgr.RegisterOracle(registerOpts)
	req.NoError(err)

	suite.Backend.Commit()

	// pause service
	pauseOpts := bind.NewKeyedTransactor(owner)
	_, err = suite.oracleMgr.Pause(pauseOpts)
	req.NoError(err)
	suite.Backend.Commit()

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

	suite.Backend.Commit()

	// deregister oracle
	deregisterOpts := bind.NewKeyedTransactor(user)
	_, err = suite.oracleMgr.DeregisterOracle(deregisterOpts)
	req.NoError(err)

	suite.Backend.Commit()

	// oracle count must be zero
	count, err := suite.oracleMgr.GetOracleCount(&bind.CallOpts{})
	req.NoError(err)
	req.Zero(count.Uint64())
}
*/

func (suite *OracleMgrSuite) TestSubmitPrice_Paused() {
	req := suite.Require()

	// register oracle
	registerOpts := bind.NewKeyedTransactor(owner)
	_, err := suite.oracleMgr.RegisterOracle(registerOpts, crypto.PubkeyToAddress(user.PublicKey))
	req.NoError(err)

	suite.Backend.Commit()

	// pause service
	pauseOpts := bind.NewKeyedTransactor(owner)
	_, err = suite.oracleMgr.Pause(pauseOpts)
	req.NoError(err)
	suite.Backend.Commit()

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
	registerOpts := bind.NewKeyedTransactor(owner)
	_, err := suite.oracleMgr.RegisterOracle(registerOpts, crypto.PubkeyToAddress(user.PublicKey))
	req.NoError(err)

	suite.Backend.Commit()

	// submit price
	priceOpts := bind.NewKeyedTransactor(user)
	_, err = suite.oracleMgr.SubmitPrice(priceOpts, common.Big2)
	req.NoError(err)
}

func (suite *OracleMgrSuite) TestSubmitPrice_NotPaused_Oracle_SubmitPriceTwiceSameRound() {
	req := suite.Require()

	// register oracle
	registerOpts := bind.NewKeyedTransactor(owner)
	_, err := suite.oracleMgr.RegisterOracle(registerOpts, crypto.PubkeyToAddress(user.PublicKey))
	req.NoError(err)

	suite.Backend.Commit()

	// submit price
	priceOpts := bind.NewKeyedTransactor(user)
	_, err = suite.oracleMgr.SubmitPrice(priceOpts, common.Big2)
	req.NoError(err)

	suite.Backend.Commit()

	// submit price
	_, err = suite.oracleMgr.SubmitPrice(priceOpts, common.Big2)
	req.Error(err, "cannot submit a price twice in the same round")
}
