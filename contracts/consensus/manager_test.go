package consensus

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
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
	// ValidatorMgr
	initialBalance   = 10 // 10 kUSD
	baseDeposit      = 1  // 1 mUSD
	maxNumValidators = 100
	freezePeriod     = 10 // 10 days
	secondsPerDay    = 86400

	// mUSD
	miningToken         = "mUSD"
	cap                 = 1073741824
	miningTokenDecimals = uint8(18)
	customFallback      = "registerValidator(address,uint256,bytes)"
)

var (
	errAlwaysFailingTransaction = errors.New("failed to estimate gas needed: gas required exceeds allowance or always failing transaction")
)

type ValidatorMgrSuite struct {
	suite.Suite
	backend                   *backends.SimulatedBackend
	contractOwner, randomUser *ecdsa.PrivateKey
	genesisValidator          *ecdsa.PrivateKey
	initialBalance            *big.Int
	baseDeposit               *big.Int
	maxNumValidators          *big.Int
	freezePeriod              *big.Int
}

func TestValidatorMgrSuite(t *testing.T) {
	suite.Run(t, new(ValidatorMgrSuite))
}

func (suite *ValidatorMgrSuite) SetupSuite() {
	req := suite.Require()

	contractOwner, err := crypto.GenerateKey()
	req.NoError(err)
	randomUser, err := crypto.GenerateKey()
	req.NoError(err)
	genesisValidator, err := crypto.GenerateKey()
	req.NoError(err)

	// users
	suite.contractOwner = contractOwner
	suite.randomUser = randomUser
	suite.genesisValidator = genesisValidator

	// valid params
	suite.initialBalance = musd(new(big.Int).SetUint64(initialBalance))
	suite.baseDeposit = musd(new(big.Int).SetUint64(baseDeposit))
	suite.maxNumValidators = new(big.Int).SetUint64(maxNumValidators)
	suite.freezePeriod = new(big.Int).SetUint64(freezePeriod)
}

func (suite *ValidatorMgrSuite) NewSimulatedBackend() *backends.SimulatedBackend {
	contractOwnerAddr := crypto.PubkeyToAddress(suite.contractOwner.PublicKey)
	randomUserAddr := crypto.PubkeyToAddress(suite.randomUser.PublicKey)
	genesisAddr := crypto.PubkeyToAddress(suite.genesisValidator.PublicKey)
	defaultAccount := core.GenesisAccount{Balance: suite.initialBalance}
	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		contractOwnerAddr: defaultAccount,
		randomUserAddr:    defaultAccount,
		genesisAddr:       defaultAccount,
	})

	return backend
}

func (suite *ValidatorMgrSuite) SetupTest() {
	suite.backend = suite.NewSimulatedBackend()
}

func (suite *ValidatorMgrSuite) TestDeployValidatorMgr() {
	req := suite.Require()

	deploymentOpts := bind.NewKeyedTransactor(suite.contractOwner)
	_, _, mgr, err := DeployValidatorMgr(deploymentOpts, suite.backend, suite.baseDeposit, suite.maxNumValidators, suite.freezePeriod, getAddress(suite.genesisValidator))
	req.NoError(err)
	req.NotNil(mgr)

	suite.backend.Commit()

	storedBaseDeposit, err := mgr.BaseDeposit(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedBaseDeposit)
	req.Equal(suite.baseDeposit, storedBaseDeposit)

	storedMaxNumValidators, err := mgr.MaxNumValidators(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedMaxNumValidators)
	req.Equal(suite.maxNumValidators, storedMaxNumValidators)

	storedFreezePeriod, err := mgr.FreezePeriod(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedFreezePeriod)
	req.Equal(dtos(suite.freezePeriod), storedFreezePeriod)

	req.True(mgr.IsGenesisValidator(&bind.CallOpts{}, getAddress(suite.genesisValidator)))
}

func (suite *ValidatorMgrSuite) TestDeployValidatorMgr_MaxNumValidators_Zero() {
	req := suite.Require()

	maxNumValidators := common.Big0
	deploymentOpts := bind.NewKeyedTransactor(suite.contractOwner)
	_, _, _, err := DeployValidatorMgr(deploymentOpts, suite.backend, suite.baseDeposit, maxNumValidators, suite.freezePeriod, getAddress(suite.genesisValidator))
	req.Equal(errAlwaysFailingTransaction, err)
}

func (suite *ValidatorMgrSuite) TestIsGenesis() {
	req := suite.Require()

	deploymentOpts := bind.NewKeyedTransactor(suite.contractOwner)
	_, _, mgr, err := DeployValidatorMgr(deploymentOpts, suite.backend, suite.baseDeposit, suite.maxNumValidators, suite.freezePeriod, getAddress(suite.genesisValidator))
	req.NoError(err)
	req.NotNil(mgr)

	suite.backend.Commit()

	testCases := []struct {
		input  common.Address
		output bool
	}{
		{
			input:  getAddress(suite.genesisValidator),
			output: true,
		},
		{
			input:  getAddress(suite.randomUser),
			output: false,
		},
	}

	for _, tc := range testCases {
		suite.T().Run(fmt.Sprintf("Address %s", tc.input.Hex()), func(t *testing.T) {
			isGenesis, err := mgr.IsGenesisValidator(&bind.CallOpts{}, tc.input)
			req.NoError(err)
			req.Equal(tc.output, isGenesis)
		})
	}
}

func (suite *ValidatorMgrSuite) TestIsValidator() {
	req := suite.Require()

	deploymentOpts := bind.NewKeyedTransactor(suite.contractOwner)
	_, _, mgr, err := DeployValidatorMgr(deploymentOpts, suite.backend, suite.baseDeposit, suite.maxNumValidators, suite.freezePeriod, getAddress(suite.genesisValidator))
	req.NoError(err)
	req.NotNil(mgr)

	suite.backend.Commit()

	testCases := []struct {
		input  common.Address
		output bool
	}{
		{
			input:  getAddress(suite.genesisValidator),
			output: true,
		},
		{
			input:  getAddress(suite.randomUser),
			output: false,
		},
	}

	for _, tc := range testCases {
		suite.T().Run(fmt.Sprintf("Address %s", tc.input.Hex()), func(t *testing.T) {
			isValidator, err := mgr.IsValidator(&bind.CallOpts{}, tc.input)
			req.NoError(err)
			req.Equal(tc.output, isValidator)
		})
	}
}

func (suite *ValidatorMgrSuite) TestGetMinimumDeposit_NotFull() {
	req := suite.Require()

	maxNumValidators := suite.maxNumValidators
	deploymentOpts := bind.NewKeyedTransactor(suite.contractOwner)
	_, _, mgr, err := DeployValidatorMgr(deploymentOpts, suite.backend, suite.baseDeposit, maxNumValidators, suite.freezePeriod, getAddress(suite.genesisValidator))
	req.NoError(err)
	req.NotNil(mgr)

	suite.backend.Commit()

	minDeposit, err := mgr.GetMinimumDeposit(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(minDeposit)

	req.Equal(suite.baseDeposit, minDeposit)
}

func (suite *ValidatorMgrSuite) TestGetMinimumDeposit_Full() {
	req := suite.Require()

	maxNumValidators := common.Big1
	deploymentOpts := bind.NewKeyedTransactor(suite.contractOwner)
	// contract includes one validator - genesis validator
	_, _, mgr, err := DeployValidatorMgr(deploymentOpts, suite.backend, suite.baseDeposit, maxNumValidators, suite.freezePeriod, getAddress(suite.genesisValidator))
	req.NoError(err)
	req.NotNil(mgr)

	suite.backend.Commit()

	minDeposit, err := mgr.GetMinimumDeposit(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(minDeposit)

	// genesis validator deposit equals base deposit
	// minimum deposit must be the smallest bid + 1
	req.Equal(new(big.Int).Add(suite.baseDeposit, common.Big1), minDeposit)
}

func (suite *ValidatorMgrSuite) TestGetValidatorAtIndex() {
	req := suite.Require()

	maxNumValidators := common.Big1
	deploymentOpts := bind.NewKeyedTransactor(suite.contractOwner)
	_, _, mgr, err := DeployValidatorMgr(deploymentOpts, suite.backend, suite.baseDeposit, maxNumValidators, suite.freezePeriod, getAddress(suite.genesisValidator))
	req.NoError(err)
	req.NotNil(mgr)

	suite.backend.Commit()

	// fetch the genesis validator
	index := common.Big0
	genesisValidator, err := mgr.GetValidatorAtIndex(&bind.CallOpts{}, index)
	req.NoError(err)
	req.NotNil(genesisValidator)
	req.Equal(suite.baseDeposit, genesisValidator.Deposit)
	req.Equal(getAddress(suite.genesisValidator), genesisValidator.Code)
}

func (suite *ValidatorMgrSuite) TestGetValidatorCount() {
	
}

/*





func (suite *ElectionContractSuite) TestGetOwner() {
	req := suite.Require()

	latestOwner, err := suite.contract.GetOwner(&bind.CallOpts{})
	req.NoError(err)
	req.Equal(crypto.PubkeyToAddress(suite.owner.PublicKey), latestOwner)
}

func (suite *ElectionContractSuite) TestTransferOwnership_NotOwner() {
	req := suite.Require()

	// future owner
	newOwnerPK, err := crypto.GenerateKey()
	req.NoError(err)
	newOwnerAddr := crypto.PubkeyToAddress(newOwnerPK.PublicKey)
	_, err = suite.contract.TransferOwnership(bind.NewKeyedTransactor(suite.randomUser), newOwnerAddr)
	req.Equal(errTransactionFailed, err)
}

func (suite *ElectionContractSuite) TestTransferOwnership_Owner() {
	req := suite.Require()

	newOwner := getAddress(suite.randomUser)
	_, err := suite.contract.TransferOwnership(bind.NewKeyedTransactor(suite.owner), newOwner)
	req.NoError(err)
	suite.backend.Commit()

	latestOwner, err := suite.contract.GetOwner(&bind.CallOpts{})
	req.NoError(err)
	req.Equal(newOwner, latestOwner)
}







func (suite *ElectionContractSuite) TestSetBaseDeposit_NotOwner() {
	req := suite.Require()

	_, err := suite.contract.SetBaseDeposit(bind.NewKeyedTransactor(suite.randomUser), common.Big0)
	req.Equal(errTransactionFailed, err)
}

func (suite *ElectionContractSuite) TestSetBaseDeposit_Owner() {
	req := suite.Require()

	deposit := new(big.Int).Add(suite.baseDeposit, common.Big1)
	_, err := suite.contract.SetBaseDeposit(bind.NewKeyedTransactor(suite.owner), deposit)
	req.NoError(err)
	suite.backend.Commit()

	latestDeposit, err := suite.contract.BaseDeposit(&bind.CallOpts{})
	req.NoError(err)
	req.Equal(deposit, latestDeposit)
}

func (suite *ElectionContractSuite) TestSetMaxValidators_NotOwner() {
	req := suite.Require()

	_, err := suite.contract.SetMaxValidators(bind.NewKeyedTransactor(suite.randomUser), common.Big2)
	req.Equal(errTransactionFailed, err)
}

func (suite *ElectionContractSuite) TestSetMaxValidators_Owner_GreaterOrEqualThanValidatorCount() {
	req := suite.Require()

	oldValidatorCount, err := suite.contract.GetValidatorCount(&bind.CallOpts{})
	req.NoError(err)

	maxValidators := new(big.Int).Add(oldValidatorCount, common.Big1)
	_, err = suite.contract.SetMaxValidators(bind.NewKeyedTransactor(suite.owner), maxValidators)
	req.NoError(err)
	suite.backend.Commit()

	// value should be updated
	latestMax, err := suite.contract.MaxValidators(&bind.CallOpts{})
	req.NoError(err)
	req.Equal(maxValidators, latestMax)

	// number of validators should remain the same
	latestValidatorCount, err := suite.contract.GetValidatorCount(&bind.CallOpts{})
	req.NoError(err)
	req.Equal(oldValidatorCount, latestValidatorCount)
}

func (suite *ElectionContractSuite) TestSetMaxValidators_Owner_LessThanValidatorCount() {
	req := suite.Require()

	oldValidatorCount, err := suite.contract.GetValidatorCount(&bind.CallOpts{})
	req.NoError(err)

	maxValidators := new(big.Int).Sub(oldValidatorCount, common.Big1)
	_, err = suite.contract.SetMaxValidators(bind.NewKeyedTransactor(suite.owner), maxValidators)
	req.NoError(err)
	suite.backend.Commit()

	// value should be updated
	latestMaxValidators, err := suite.contract.MaxValidators(&bind.CallOpts{})
	req.NoError(err)
	req.Equal(maxValidators, latestMaxValidators)

	// validator count should be equal to the new max
	latestValidatorCount, err := suite.contract.GetValidatorCount(&bind.CallOpts{})
	req.NoError(err)
	req.Equal(maxValidators, latestValidatorCount)

	// validator that left the election should be the
	// genesis validator - smallest bidder
	// @TODO (rgeraldes)
}

func (suite *ElectionContractSuite) TestJoin_AlreadyValidator() {
	req := suite.Require()

	// genesis validator is already part of the election
	_, err := suite.contract.Join(bind.NewKeyedTransactor(suite.genesisValidator))
	req.Equal(errTransactionFailed, err)
}

func (suite *ElectionContractSuite) TestJoin_InsufficientDeposit() {
	req := suite.Require()

	minDeposit, err := suite.contract.GetMinimumDeposit(&bind.CallOpts{})
	req.NoError(err)

	opts := bind.NewKeyedTransactor(suite.randomUser)
	opts.Value = new(big.Int).Sub(minDeposit, common.Big1)
	_, err = suite.contract.Join(opts)
	req.Equal(errTransactionFailed, err)
}

func (suite *ElectionContractSuite) TestJoin_InsertGreaterThan() {
	req := suite.Require()

	sender := suite.randomUser
	senderAddr := getAddress(sender)

	oldValidatorCount, err := suite.contract.GetValidatorCount(&bind.CallOpts{})
	req.NoError(err)

	// make sure that the deposit is greater than the genesis deposit
	genesisValidator, err := suite.contract.GetValidatorAtIndex(&bind.CallOpts{}, common.Big0)
	req.NoError(err)
	opts := bind.NewKeyedTransactor(sender)
	opts.Value = new(big.Int).Add(genesisValidator.Deposit, common.Big1)
	_, err = suite.contract.Join(opts)
	req.NoError(err)
	suite.backend.Commit()

	// the election should have one more validator
	latestValidatorCount, err := suite.contract.GetValidatorCount(&bind.CallOpts{})
	req.NoError(err)
	req.True(new(big.Int).Add(oldValidatorCount, common.Big1).Cmp(latestValidatorCount) == 0)

	// validator at index 0 should be the new candidate
	validator, err := suite.contract.GetValidatorAtIndex(&bind.CallOpts{}, common.Big0)
	req.NoError(err)
	req.Equal(validator.Code, senderAddr)
}

func (suite *ElectionContractSuite) TestJoin_InsertLessOrEqualTo() {
	req := suite.Require()

	sender := suite.randomUser
	senderAddr := getAddress(sender)

	oldValidatorCount, err := suite.contract.GetValidatorCount(&bind.CallOpts{})
	req.NoError(err)

	// make sure that the deposit is equal to the genesis
	// the deposit. In this scenario it cannot be lower than the
	// genesis deposit because the value would be less than the
	// base deposit
	genesisValidator, err := suite.contract.GetValidatorAtIndex(&bind.CallOpts{}, common.Big0)
	req.NoError(err)
	opts := bind.NewKeyedTransactor(sender)
	opts.Value = genesisValidator.Deposit
	_, err = suite.contract.Join(opts)
	req.NoError(err)
	suite.backend.Commit()

	// the election should have one more validator
	latestValidatorCount, err := suite.contract.GetValidatorCount(&bind.CallOpts{})
	req.NoError(err)
	req.True(new(big.Int).Add(oldValidatorCount, common.Big1).Cmp(latestValidatorCount) == 0)

	// validator at the end of the list should be the new candidate
	validator, err := suite.contract.GetValidatorAtIndex(&bind.CallOpts{}, new(big.Int).Sub(latestValidatorCount, common.Big1))
	req.NoError(err)
	req.Equal(validator.Code, senderAddr)
}

func (suite *ElectionContractSuite) TestLeave_NotValidator() {
	req := suite.Require()

	_, err := suite.contract.Leave(bind.NewKeyedTransactor(suite.randomUser))
	req.Equal(errTransactionFailed, err)
}

func (suite *ElectionContractSuite) TestLeave_Validator() {
	req := suite.Require()

	oldValidatorCount, err := suite.contract.GetValidatorCount(&bind.CallOpts{})
	req.NoError(err)

	sender := suite.genesisValidator
	senderAddr := getAddress(sender)
	_, err = suite.contract.Leave(bind.NewKeyedTransactor(sender))
	req.NoError(err)
	suite.backend.Commit()

	// user should not be a validator anymore
	isValidator, err := suite.contract.IsValidator(&bind.CallOpts{}, senderAddr)
	req.NoError(err)
	req.False(isValidator)

	// number of validators should be equal to old count minus one
	latestValidatorCount, err := suite.contract.GetValidatorCount(&bind.CallOpts{})
	req.NoError(err)
	req.Equal(new(big.Int).Sub(oldValidatorCount, common.Big1), latestValidatorCount)

	// latest deposit should have a release date (different than 0)
	latestDepositCount, err := suite.contract.GetDepositCount(&bind.CallOpts{From: senderAddr})
	req.NoError(err)
	currentDeposit, err := suite.contract.GetDepositAtIndex(&bind.CallOpts{From: senderAddr}, new(big.Int).Sub(latestDepositCount, common.Big1))
	req.NoError(err)
	req.NotZero(currentDeposit.Amount.Uint64())
}

func (suite *ElectionContractSuite) TestRedeemFunds_NoDeposits() {
	req := suite.Require()

	sender := suite.randomUser
	senderAddr := getAddress(sender)
	_, err := suite.contract.RedeemDeposits(bind.NewKeyedTransactor(sender))
	req.NoError(err)

	// deposit count should be zero
	depositCount, err := suite.contract.GetDepositCount(&bind.CallOpts{From: senderAddr})
	req.NoError(err)
	req.Zero(depositCount.Uint64())
	suite.backend.Commit()

	// user balance should be less than the initial balance =
	// cost of the previous transaction + no deposits
	balance, err := suite.backend.BalanceAt(context.TODO(), senderAddr, suite.backend.CurrentBlock().Number())
	req.NoError(err)
	req.True(balance.Cmp(suite.initialBalance) < 0)
}

func (suite *ElectionContractSuite) TestRedeemFunds_LockedDeposit() {
	// @NOTE (rgeraldes) - default unbonding period
	// is 10 days - deposit will remain locked for 10 days
	// as soon as the validator decides to leave the election.
	req := suite.Require()

	// leave the election - we will use the genesis
	// since he has one deposit with no release date
	sender := suite.genesisValidator
	senderAddr := getAddress(sender)

	oldDepositCount, err := suite.contract.GetDepositCount(&bind.CallOpts{From: senderAddr})
	req.NoError(err)

	opts := bind.NewKeyedTransactor(sender)
	_, err = suite.contract.Leave(opts)
	req.NoError(err)
	suite.backend.Commit()

	// redeem deposit
	_, err = suite.contract.RedeemDeposits(opts)
	req.NoError(err)
	suite.backend.Commit()

	// user balance should be less than the initial balance =
	// cost of previous transactions + no deposit (locked deposit)
	balance, err := suite.backend.BalanceAt(context.TODO(), senderAddr, suite.backend.CurrentBlock().Number())
	req.NoError(err)
	req.True(balance.Cmp(suite.initialBalance) < 0)

	// user deposit should be available as before
	depositCount, err := suite.contract.GetDepositCount(&bind.CallOpts{From: senderAddr})
	req.NoError(err)
	req.True(oldDepositCount.Cmp(depositCount) == 0)
}

func (suite *ElectionContractSuite) TestRedeemFunds_UnlockedDeposit() {
	// @NOTE (rgeraldes) - default unbonding period
	// is 10 days - deposit will remain locked for 10 days
	// as soon as the validator decides to leave the election.
	req := suite.Require()

	// deploy a new version of the contract
	// with an unbounding period of 0 days
	suite.DeployElectionContract(suite.baseDeposit, suite.maxValidators, common.Big0)
	suite.backend.Commit()

	// leave the election - we will use the genesis
	// since he has one deposit with no release date
	sender := suite.genesisValidator
	senderAddr := getAddress(sender)

	opts := bind.NewKeyedTransactor(sender)
	_, err := suite.contract.Leave(opts)
	req.NoError(err)
	suite.backend.Commit()

	// redeem deposit
	_, err = suite.contract.RedeemDeposits(opts)
	req.NoError(err)
	suite.backend.Commit()

	// user balance should be greater than the initial balance =
	// cost of previous transactions + unlocked deposit
	balance, err := suite.backend.BalanceAt(context.TODO(), senderAddr, suite.backend.CurrentBlock().Number())
	req.NoError(err)
	req.True(balance.Cmp(suite.initialBalance) > 0)

	// user deposit should not be available anymore
	depositCount, err := suite.contract.GetDepositCount(&bind.CallOpts{From: senderAddr})
	req.NoError(err)
	req.Zero(depositCount.Uint64())
}




*/

// dtos converts days to seconds
func dtos(days *big.Int) *big.Int {
	return new(big.Int).Mul(days, new(big.Int).SetUint64(secondsPerDay))
}

// musd converts the value to mUSD units
func musd(value *big.Int) *big.Int {
	return new(big.Int).Mul(value, new(big.Int).SetUint64(params.Ether))
}

// getAddress return the address of the given private key
func getAddress(privateKey *ecdsa.PrivateKey) common.Address {
	return crypto.PubkeyToAddress(privateKey.PublicKey)
}

/*
	_, _, mUSD, err := token.DeployMiningToken(ownerTransactOpts, suite.backend, miningToken, miningToken, musd(new(big.Int).SetUint64(cap)), miningTokenDecimals)
	req.NoError(err)
	req.NotNil(mUSD)

	_, err = mUSD.Mint(ownerTransactOpts, getAddress(suite.randomUser), suite.baseDeposit)
	req.NoError(err)

	suite.backend.Commit()

	balance, err := mUSD.BalanceOf(&bind.CallOpts{}, getAddress(suite.randomUser))
	req.NoError(err)


	deposit := suite.baseDeposit
	_, err = mUSD.Transfer(userTransactOpts, mgrAddr, deposit, []byte("non-zero"), customFallback)
	req.NoError(err)
*/
