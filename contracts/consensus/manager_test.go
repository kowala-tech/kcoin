package consensus

import (
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/kowala-tech/kcoin/accounts/abi/bind/backends"
	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/crypto"
	"github.com/kowala-tech/kcoin/kcoin/genesis"
	"github.com/kowala-tech/kcoin/params"
	"github.com/stretchr/testify/require"
)

const (
	customFallback = "registerValidator(address,uint256,bytes)"
	initialBalance = 10 // 10 kUSD
)

var (
	user, _ = crypto.GenerateKey()
)

func TestContractCreation(t *testing.T) {
	opts := genesis.GetDefaultOpts()

	// create genesis
	gen, err := genesis.New(opts)
	require.NoError(t, err)
	require.NotNil(t, gen)

	// create backend
	backend := backends.NewSimulatedBackend(gen.Alloc)

	// ValidatorMgr instance
	mgr, err := NewValidatorMgr(common.Address{}, backend)
	require.NoError(t, err)
	require.NotNil(t, err)

	storedBaseDeposit, err := mgr.BaseDeposit(&bind.CallOpts{})
	require.NoError(t, err)
	requite.NotNil(storedBaseDeposit)
	require.Equal(t, musd(opts.Consensus.BaseDeposit), storedBaseDeposit)

	storedBaseDeposit, err := mgr.BaseDeposit(&bind.CallOpts{})
	require.NoError(t, err)
	requite.NotNil(storedBaseDeposit)
	require.Equal(t, musd(opts.Consensus.BaseDeposit), storedBaseDeposit)

	storedFreezePeriod, err := mgr.FreezePeriod(&bind.CallOpts{})
	require.NoError(t, err)
	require.NotNil(t, storedFreezePeriod)
	require.Equal(t, opts.Consensus.FreezePeriod, storedFreezePeriod)

	storedMaxNumValidators, err := mgr.MaxNumValidators(&bind.CallOpts{})
	require.NoError(t, err)
	require.NotNil(t, storedMaxNumValidators)
	require.Equal(t, opts.Consesnsus.MaxNumValidators, storedMaxNumValidators)

	storedMiningTokenAddr, err := mgr.MiningTokenAddr(&bind.CallOpts{})
	require.NoError(t, err)
	require.NotNil(t, storedMiningTokenAddr)
	require.Equal(t, "", storedMiningTokenAddr)
}

func TestIsValidator(t *testing.T) {
	opts := genesis.GetDefaultOpts()

	// add a token holder
	holders := opts.Consensus.MiningToken.Holders
	holder := genesis.TokenHolder{
		Address:   getAddress(user).Str(),
		NumTokens: opts.Consensus.BaseDeposit,
	}

	holders = append(holders, holder)

	// prefund the token holder with kUSD to join the consensus
	opts.PrefundedAccounts = append(opts.PrefundedAccounts, genesis.PrefundedAccount{
		AccountAddress: holder.Address,
		Balance:        initialBalance,
	})

	// create genesis
	genesis, err := genesis.New(opts)
	require.NoError(t, err)
	require.NotNil(t, genesis)

	// create backend
	backend := backends.NewSimulatedBackend(alloc)(genesis.Alloc)

	// mUSD instance
	mUSD, err := token.NewMiningToken(common.Address{}, backend)
	require.NoError(t, err)
	require.NotNil(t, mUSD)

	// validatorMgr instance
	validatorMgrAddr := common.Address{}
	validatorMgr, err := NewValidatorMgr(validatorMgrAddr, backend)
	require.NoError(t, err)
	require.NotNil(t, validatorMgr)

	// register user as validator
	transactOpts := bind.NewKeyedTransactor(user)
	_, err = mUSD.Transfer(transactOpts, validatorMgrAddr, musd(new(big.Int).SetUint64(holder.NumTokens)), []byte("not_zero"), customFallback)
	require.NoError(t, err)

	backend.Commit()

	testCases := []struct {
		input  common.Address
		output bool
	}{
		{
			input:  getAddress(user),
			output: true,
		},
		{
			input:  getAddress(user2),
			output: false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Address %s", tc.input.Hex()), func(t *testing.T) {
			isValidator, err := validatorMgr.IsGenesisValidator(&bind.CallOpts{}, tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.output, isValidator)
		})
	}
}

/*

func (suite *ValidatorMgrSuite) TestIsValidator() {
	req := suite.Require()

	options := genesis.GetDefaultOpts(suite.contractOwner)
	options.Governance.Origin =
	options.Consensus.MiningToken.Holders = append(options.Consensus.MiningToken.Holders, genesis.TokenHolder{
		Address:   getAddress(suite.user).Str(),
		NumTokens: baseDeposit,
	})
	genesis, err := genesis.New(options)
	req.NoError(err)
	req.NotNil(genesis)

	backend := backends.NewSimulatedBackend(genesis.Alloc)

	mUSD, err := token.NewMiningToken(common.Address{}, backend)
	req.NoError(err)
	req.NotNil(mUSD)

	validatorMgrAddr := common.Address{}
	validatorMgr, err := NewValidatorMgr(validatorMgrAddr, backend)
	req.NoError(err)
	req.NotNil(validatorMgr)

	userTransactOpts := bind.NewKeyedTransactor(suite.user)
	minDeposit := suite.mgrArgs.baseDeposit
	_, err = mUSD.Transfer(userTransactOpts, validatorMgrAddr, minDeposit, []byte("not_zero_value"), suite.tokenArgs.customFallback)
	req.NoError(err)

	backend.Commit()

	randomUser, err := crypto.GenerateKey()
	req.NoError(err)
	req.NotNil(randomUser)


}


/*
const (
	// prefund
	initialBalance = 10 // 10 kUSD

	// ValidatorMgr
	baseDeposit      = 1 // 1 mUSD
	maxNumValidators = 100
	freezePeriod     = 10 // 10 days

	// MiningToken
	miningToken               = "mUSD"
	miningTokenCap            = 1073741824
	miningTokenDecimals       = uint8(18)
	miningTokenCustomFallback = "registerValidator(address,uint256,bytes)"

	secondsPerDay = 86400
)


var (
	errAlwaysFailingTransaction = errors.New("failed to estimate gas needed: gas required exceeds allowance or always failing transaction")
)


type ValidatorMgrSuite struct {
	suite.Suite
	contractOwner  *ecdsa.PrivateKey
	user           *ecdsa.PrivateKey
	initialBalance *big.Int
}

func TestValidatorMgrSuite(t *testing.T) {
	suite.Run(t, new(ValidatorMgrSuite))
}

func (suite *ValidatorMgrSuite) SetupSuite() {
	req := suite.Require()
	req.Error()

	contractOwner, err := crypto.GenerateKey()
	req.NoError(err)
	user, err := crypto.GenerateKey()
	req.NoError(err)

	suite.contractOwner = contractOwner
	suite.user = user
	suite.initialBalance = musd(new(big.Int).SetUint64(initialBalance))

	suite.mgrArgs = &validMgrArgs{
		baseDeposit:      musd(new(big.Int).SetUint64(baseDeposit)),
		maxNumValidators: new(big.Int).SetUint64(maxNumValidators),
		freezePeriod:     new(big.Int).SetUint64(freezePeriod),
	}

	suite.tokenArgs = &validTokenArgs{
		name:           miningToken,
		symbol:         miningToken,
		decimals:       miningTokenDecimals,
		cap:            musd(new(big.Int).SetUint64(miningTokenCap)),
		customFallback: miningTokenCustomFallback,
	}
}





func (suite *ValidatorMgrSuite) TestIsGenesisValidator() {
	req := suite.Require()

	options := genesis.GetDefaultOpts()
	options.Consensus.Validators = append(options.Consensus.Validators, genesis.Validator{
		Address: getAddress(suite.user).Str(),
		Deposit: baseDeposit,
	})
	options.Consensus.MiningToken.Holders = append(options.Consensus.MiningToken.Holders, genesis.TokenHolder{
		Address:   getAddress(suite.user).Str(),
		NumTokens: baseDeposit,
	})
	options.Consensus.MiningToken.Holders = append(options.Consensus.MiningToken.Holders, genesis.TokenHolder{
		Address:   getAddress(suite.user2).Str(),
		NumTokens: baseDeposit,
	})
	genesis, err := genesis.New(options)
	req.NoError(err)
	req.NotNil(genesis)

	backend := backends.NewSimulatedBackend(genesis.Alloc)

	ownerTransactOpts := bind.NewKeyedTransactor(suite.contractOwner)

	mUSD, err := token.NewMiningToken(common.Address{}, backend)
	req.NoError(err)
	req.NotNil(mUSD)

	validatorMgrAddr := common.Address{}
	validatorMgr, err := NewValidatorMgr(validatorMgrAddr, backend)
	req.NoError(err)
	req.NotNil(validatorMgr)

	userTransactOpts := bind.NewKeyedTransactor(suite.user)
	minDeposit := suite.mgrArgs.baseDeposit
	_, err = mUSD.Transfer(userTransactOpts, validatorMgrAddr, minDeposit, []byte("not_zero_value"), suite.tokenArgs.customFallback)
	req.NoError(err)

	backend.Commit()

	randomUser, err := crypto.GenerateKey()
	req.NoError(err)
	req.NotNil(randomUser)

	testCases := []struct {
		input  common.Address
		output bool
	}{
		{
			input:  getAddress(suite.user),
			output: true,
		},
		{
			input:  getAddress(randomUser),
			output: false,
		},
	}

	for _, tc := range testCases {
		suite.T().Run(fmt.Sprintf("Address %s", tc.input.Hex()), func(t *testing.T) {
			isValidator, err := validatorMgr.IsValidator(&bind.CallOpts{}, tc.input)
			req.NoError(err)
			req.Equal(tc.output, isValidator)
		})
	}
}

func (suite *ValidatorMgrSuite) TestGetMinimumDeposit_NotFull() {
	req := suite.Require()

	options := genesis.GetDefaultOpts()
	options.Consensus.Validators = append(options.Consensus.Validators, genesis.Validator{
		Address: getAddress(suite.user).Str(),
		Deposit: baseDeposit,
	})
	options.Consensus.MiningToken.Holders = append(options.Consensus.MiningToken.Holders, genesis.TokenHolder{
		Address:   getAddress(suite.user).Str(),
		NumTokens: baseDeposit,
	})
	options.Consensus.MiningToken.Holders = append(options.Consensus.MiningToken.Holders, genesis.TokenHolder{
		Address:   getAddress(suite.user2).Str(),
		NumTokens: baseDeposit,
	})
	genesis, err := genesis.New(options)
	req.NoError(err)
	req.NotNil(genesis)

	backend := backends.NewSimulatedBackend(genesis.Alloc)

	maxNumValidators := suite.maxNumValidators
	deploymentOpts := bind.NewKeyedTransactor(suite.contractOwner<)
	_, _, mgr, err := DeployValidatorMgr(deploymentOpts, suite.backend, suite.baseDeposit, maxNumValidators, suite.freezePeriod, getAddress(suite.genesisValidator))
	req.NoError(err)
	req.NotNil(mgr)

	suite.backend.Commit()

	minDeposit, err := mgr.GetMinimumDeposit(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(minDeposit)

	req.Equal(suite.baseDeposit, minDeposit)
}

/*

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

// mUSD converts the value to mUSD units
func musd(value *big.Int) *big.Int {
	return new(big.Int).Mul(value, new(big.Int).SetUint64(params.Ether))
}

// kUSD converts the value to mUSD units
func kusd(value *big.Int) *big.Int {
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


*/
