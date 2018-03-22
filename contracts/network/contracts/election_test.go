package contracts

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"testing"

	"github.com/kowala-tech/kUSD/accounts/abi/bind"
	"github.com/kowala-tech/kUSD/accounts/abi/bind/backends"
	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/core"
	"github.com/kowala-tech/kUSD/crypto"
	"github.com/kowala-tech/kUSD/params"
	"github.com/stretchr/testify/suite"
)

const (
	initialBalance  = 10 // 10 kUSD
	baseDeposit     = 1  // 1 kUSD
	maxValidators   = 100
	unbondingPeriod = 10 // 10 days
)

var (
	errTransactionFailed = errors.New("failed to estimate gas needed: gas required exceeds allowance or always failing transaction")
)

type ElectionContractSuite struct {
	suite.Suite
	backend           *backends.SimulatedBackend
	contract          *ElectionContract
	owner, randomUser *ecdsa.PrivateKey
	genesisValidator  *ecdsa.PrivateKey
	initialBalance    *big.Int
	baseDeposit       *big.Int
	maxValidators     *big.Int
	unbondingPeriod   *big.Int
}

func TestElectionContractSuite(t *testing.T) {
	suite.Run(t, new(ElectionContractSuite))
}

func (suite *ElectionContractSuite) SetupSuite() {
	req := suite.Require()

	owner, err := crypto.GenerateKey()
	req.NoError(err)
	randomUser, err := crypto.GenerateKey()
	req.NoError(err)
	genesisValidator, err := crypto.GenerateKey()
	req.NoError(err)

	suite.owner = owner
	suite.randomUser = randomUser
	suite.genesisValidator = genesisValidator
	suite.initialBalance = new(big.Int).Mul(new(big.Int).SetUint64(initialBalance), new(big.Int).SetUint64(params.Ether))
	suite.baseDeposit = new(big.Int).Mul(new(big.Int).SetUint64(baseDeposit), new(big.Int).SetUint64(params.Ether))
	suite.maxValidators = new(big.Int).SetUint64(maxValidators)
	suite.unbondingPeriod = new(big.Int).SetUint64(unbondingPeriod)
}

func (suite *ElectionContractSuite) NewSimulatedBackend() *backends.SimulatedBackend {
	ownerAddr := crypto.PubkeyToAddress(suite.owner.PublicKey)
	randomUserAddr := crypto.PubkeyToAddress(suite.randomUser.PublicKey)
	genesisAddr := crypto.PubkeyToAddress(suite.genesisValidator.PublicKey)
	defaultAccount := core.GenesisAccount{Balance: suite.initialBalance}
	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		ownerAddr:      defaultAccount,
		randomUserAddr: defaultAccount,
		genesisAddr:    defaultAccount,
	})

	return backend
}

func (suite *ElectionContractSuite) DeployElectionContract(baseDeposit, maxValidators, unbondingPeriod *big.Int) error {
	opts := bind.NewKeyedTransactor(suite.owner)
	_, _, contract, err := DeployElectionContract(opts, suite.backend, baseDeposit, maxValidators, unbondingPeriod, crypto.PubkeyToAddress(suite.genesisValidator.PublicKey))
	if err != nil {
		return err
	}
	suite.contract = contract

	// NOTE (rgeraldes) - add balance to cover the base deposit
	// for the genesis validator. Eventually this could change
	// as soon as the token contracts are completed.
	opts.Value = suite.baseDeposit
	_, err = contract.ElectionContractTransactor.contract.Transfer(opts)
	if err != nil {
		return err
	}
	suite.backend.Commit()

	return nil
}

func (suite *ElectionContractSuite) SetupTest() {
	req := suite.Require()

	suite.backend = suite.NewSimulatedBackend()
	req.NoError(suite.DeployElectionContract(suite.baseDeposit, suite.maxValidators, suite.unbondingPeriod))
}

func (suite *ElectionContractSuite) TestDeployElectionContract() {
	req := suite.Require()

	latestBaseDeposit, err := suite.contract.BaseDeposit(&bind.CallOpts{})
	req.NoError(err)
	req.Equal(suite.baseDeposit, latestBaseDeposit)

	latestMaxValidators, err := suite.contract.MaxValidators(&bind.CallOpts{})
	req.NoError(err)
	req.Equal(suite.maxValidators, latestMaxValidators)

	/*
		@TODO (rgeraldes)
		latestUnbondingPeriod, err := suite.contract.UnbondingPeriod(&bind.CallOpts{})
		req.NoError(err)
		req.Equal(new(big.Int).SetUint64(suite.unbondingPeriod.Int64()*time.Hour.*24), latestUnbondingPeriod)
	*/

	genesisValidator, err := suite.contract.GenesisValidator(&bind.CallOpts{})
	req.NoError(err)
	req.Equal(crypto.PubkeyToAddress(suite.genesisValidator.PublicKey), genesisValidator)
}

func (suite *ElectionContractSuite) TestDeployElectionContract_MaxValidatorsEqualZero() {
	req := suite.Require()

	maxValidators := common.Big0
	req.Equal(errTransactionFailed, suite.DeployElectionContract(suite.baseDeposit, maxValidators, suite.unbondingPeriod))
}

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

func (suite *ElectionContractSuite) TestIsGenesis() {
	req := suite.Require()

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
		isGenesis, err := suite.contract.IsGenesisValidator(&bind.CallOpts{}, tc.input)
		req.NoError(err)
		req.Equal(tc.output, isGenesis)
	}
}

func (suite *ElectionContractSuite) TestIsValidator() {
	req := suite.Require()

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
		isValidator, err := suite.contract.IsValidator(&bind.CallOpts{}, tc.input)
		req.NoError(err)
		req.Equal(tc.output, isValidator)
	}
}

func (suite *ElectionContractSuite) TestGetMinimumDeposit_ElectionFull() {
	req := suite.Require()

	// leave a position available for the genesis validator - max validators = 1
	maxValidators := new(big.Int).SetUint64(1)
	suite.DeployElectionContract(suite.baseDeposit, maxValidators, suite.unbondingPeriod)

	minDeposit, err := suite.contract.GetMinimumDeposit(&bind.CallOpts{})
	req.NoError(err)
	// min deposit should be greater (+ 1) than the smallest stake
	// at play which is equal to the base deposit (genesis validator)
	req.Equal((new(big.Int).Add(suite.baseDeposit, common.Big1)), minDeposit)
}

func (suite *ElectionContractSuite) TestGetMinimumDeposit_ElectionNotFull() {
	// by default the contract has one validator (genesis) and 99 (100 - 1)
	// positions available
	req := suite.Require()

	minDeposit, err := suite.contract.GetMinimumDeposit(&bind.CallOpts{})
	req.NoError(err)
	// min deposit should be equal to the base deposit since
	// there are positions available
	req.Equal(suite.baseDeposit, minDeposit)
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

func getAddress(privateKey *ecdsa.PrivateKey) common.Address {
	return crypto.PubkeyToAddress(privateKey.PublicKey)
}
