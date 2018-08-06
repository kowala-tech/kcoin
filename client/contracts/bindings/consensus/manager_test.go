package consensus_test

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/kowala-tech/kcoin/client/accounts/abi"
	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/accounts/abi/bind/backends"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/consensus"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/ownership"
	"github.com/kowala-tech/kcoin/client/core"
	"github.com/kowala-tech/kcoin/client/crypto"
	"github.com/kowala-tech/kcoin/client/knode/genesis"
	"github.com/kowala-tech/kcoin/client/params"
	"github.com/stretchr/testify/suite"
)

var (
	validator, _     = crypto.GenerateKey()
	deregistered, _  = crypto.GenerateKey()
	user, _          = crypto.GenerateKey()
	governor, _      = crypto.GenerateKey()
	author, _        = crypto.HexToECDSA("bfef37ae9ac5d5e7ebbbefc19f4e1f572a7ca7aa0d28e527b7d62950951cc5eb")
	validatorMgrAddr = common.HexToAddress("0x161ad311F1D66381C17641b1B73042a4CA731F9f")
	multiSigAddr     = common.HexToAddress("0xA143ac5ec5D95f16aFD5Fc3B09e0aDaf360ffC9e")
	tokenAddr        = common.HexToAddress("0xB012F49629258C9c35b2bA80cD3dc3C841d9719D")
	secondsPerDay    = new(big.Int).SetUint64(86400)
)

func getDefaultOpts() genesis.Options {
	baseDeposit := uint64(20)
	superNodeAmount := uint64(6000000)
	tokenHolder := genesis.TokenHolder{
		Address:   getAddress(validator).Hex(),
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
			Validators: []genesis.Validator{{
				Address: tokenHolder.Address,
				Deposit: tokenHolder.NumTokens,
			}},
			MiningToken: &genesis.MiningTokenOpts{
				Name:     "mUSD",
				Symbol:   "mUSD",
				Cap:      20000000,
				Decimals: 18,
				Holders:  []genesis.TokenHolder{tokenHolder, {Address: getAddress(user).Hex(), NumTokens: 10000000}},
			},
		},
		Governance: &genesis.GovernanceOpts{
			Origin:           getAddress(author).Hex(),
			Governors:        []string{getAddress(governor).Hex()},
			NumConfirmations: 1,
		},
		DataFeedSystem: &genesis.DataFeedSystemOpts{
			MaxNumOracles: 10,
			FreezePeriod:  0,
			BaseDeposit:   0,
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
				Address: getAddress(governor).Hex(),
				Balance: 10,
			},
			{
				Address: getAddress(user).Hex(),
				Balance: 10,
			},
			{
				Address: getAddress(deregistered).Hex(),
				Balance: 10,
			},
		},
	}

	return opts
}

type ValidatorMgrSuite struct {
	suite.Suite
	backend      *backends.SimulatedBackend
	opts         genesis.Options
	validatorMgr *consensus.ValidatorMgr
	multiSig     *ownership.MultiSigWallet
	miningToken  *consensus.MiningToken
}

func TestValidatorMgrSuite(t *testing.T) {
	suite.Run(t, new(ValidatorMgrSuite))
}

func (suite *ValidatorMgrSuite) BeforeTest(suiteName, testName string) {
	// TestDeploy does not rely on the genesis utility
	if strings.Contains(testName, "TestDeploy") {
		return
	}

	req := suite.Require()

	// create genesis
	opts := getDefaultOpts()
	req.NotNil(opts)

	switch {
	case strings.Contains(testName, "_Full"):
		opts.Consensus.MaxNumValidators = 1
	case testName == "TestReleaseDeposits_UnlockedDeposit":
		opts.Consensus.FreezePeriod = 0
	}
	suite.opts = opts

	genesis, err := genesis.Generate(opts)
	req.NoError(err)
	req.NotNil(genesis)

	backend := backends.NewSimulatedBackend(genesis.Alloc)
	req.NotNil(backend)
	suite.backend = backend

	// ValidatorMgr instance
	mgr, err := consensus.NewValidatorMgr(validatorMgrAddr, backend)
	req.NoError(err)
	req.NotNil(mgr)
	suite.validatorMgr = mgr

	// multiSig instance
	multiSig, err := ownership.NewMultiSigWallet(multiSigAddr, backend)
	req.NoError(err)
	req.NotNil(multiSig)
	suite.multiSig = multiSig

	// MiningToken instance
	mToken, err := consensus.NewMiningToken(tokenAddr, backend)
	req.NoError(err)
	req.NotNil(mToken)
	suite.miningToken = mToken
}

func (suite *ValidatorMgrSuite) TestDeploy() {
	req := suite.Require()

	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		getAddress(governor): core.GenesisAccount{
			Balance: new(big.Int).Mul(new(big.Int).SetUint64(100), new(big.Int).SetUint64(params.Kcoin)),
		},
	})

	baseDeposit := new(big.Int).SetUint64(100)
	maxNumValidators := new(big.Int).SetUint64(100)
	freezePeriod := new(big.Int).SetUint64(100)
	superNodeAmount := new(big.Int).SetUint64(100)

	transactOpts := bind.NewKeyedTransactor(governor)
	_, _, mgr, err := consensus.DeployValidatorMgr(transactOpts, backend, baseDeposit, maxNumValidators, freezePeriod, tokenAddr, superNodeAmount)
	req.NoError(err)
	req.NotNil(mgr)

	backend.Commit()

	storedBaseDeposit, err := mgr.BaseDeposit(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedBaseDeposit)
	req.Equal(baseDeposit, storedBaseDeposit)

	storedFreezePeriod, err := mgr.FreezePeriod(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedFreezePeriod)
	req.Equal(dtos(freezePeriod), storedFreezePeriod)

	storedMaxNumValidators, err := mgr.MaxNumValidators(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedMaxNumValidators)
	req.Equal(maxNumValidators, storedMaxNumValidators)

	storedMiningTokenAddr, err := mgr.MiningTokenAddr(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedMiningTokenAddr)
	req.Equal(tokenAddr, storedMiningTokenAddr)
}

func (suite *ValidatorMgrSuite) TestDeploy_MaxNumValidatorsEqualZero() {
	req := suite.Require()

	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		getAddress(governor): core.GenesisAccount{
			Balance: new(big.Int).Mul(new(big.Int).SetUint64(100), new(big.Int).SetUint64(params.Kcoin)),
		},
	})

	baseDeposit := new(big.Int).SetUint64(100)
	maxNumValidators := common.Big0
	freezePeriod := new(big.Int).SetUint64(100)
	superNodeAmount := new(big.Int).SetUint64(100)

	transactOpts := bind.NewKeyedTransactor(governor)
	_, _, _, err := consensus.DeployValidatorMgr(transactOpts, backend, baseDeposit, maxNumValidators, freezePeriod, tokenAddr, superNodeAmount)
	req.Error(err, "maximum number of validators cannot be zero")
}

func (suite *ValidatorMgrSuite) TestIsValidator() {
	req := suite.Require()

	numTokens := new(big.Int).Mul(new(big.Int).SetUint64(suite.opts.Consensus.BaseDeposit), new(big.Int).SetUint64(params.Kcoin))
	req.NoError(suite.registerValidator(user, numTokens))

	// register and deregister validator
	suite.mintTokens(governor, deregistered, numTokens)
	req.NoError(suite.registerValidator(deregistered, numTokens))
	req.NoError(suite.deregisterValidator(deregistered))

	suite.backend.Commit()

	// generate a random user
	randomUser, err := crypto.GenerateKey()
	req.NoError(err)
	req.NotNil(randomUser)

	testCases := []struct {
		name   string
		input  common.Address
		output bool
	}{
		{
			name:   "genesis validator",
			input:  getAddress(validator),
			output: true,
		},
		{
			name:   "non-genesis validator",
			input:  getAddress(user),
			output: true,
		},
		{
			name:   "deregistered validator",
			input:  getAddress(deregistered),
			output: false,
		},
		{
			name:   "random user",
			input:  getAddress(randomUser),
			output: false,
		},
	}
	for _, tc := range testCases {
		suite.T().Run(fmt.Sprintf("role: %s, address: %s", tc.name, tc.input.Hex()), func(t *testing.T) {
			isValidator, err := suite.validatorMgr.IsValidator(&bind.CallOpts{}, tc.input)
			req.NoError(err)
			req.Equal(tc.output, isValidator)
		})
	}
}

func (suite *ValidatorMgrSuite) TestIsGenesisValidator() {
	req := suite.Require()

	numTokens := new(big.Int).Mul(new(big.Int).SetUint64(suite.opts.Consensus.BaseDeposit), new(big.Int).SetUint64(params.Kcoin))
	req.NoError(suite.registerValidator(user, numTokens))

	suite.backend.Commit()

	// generate a random user
	randomUser, err := crypto.GenerateKey()
	req.NoError(err)
	req.NotNil(randomUser)

	testCases := []struct {
		name   string
		input  common.Address
		output bool
	}{
		{
			name:   "genesis validator",
			input:  getAddress(validator),
			output: true,
		},
		{
			name:   "non-genesis validator",
			input:  getAddress(user),
			output: false,
		},
		{
			name:   "random user",
			input:  getAddress(randomUser),
			output: false,
		},
	}
	for _, tc := range testCases {
		suite.T().Run(fmt.Sprintf("role: %s, address %s", tc.name, tc.input.Hex()), func(t *testing.T) {
			isValidator, err := suite.validatorMgr.IsGenesisValidator(&bind.CallOpts{}, tc.input)
			req.NoError(err)
			req.Equal(tc.output, isValidator)
		})
	}
}

func (suite *ValidatorMgrSuite) TestIsSuperNode() {
	req := suite.Require()

	numTokens := new(big.Int).Mul(new(big.Int).SetUint64(suite.opts.Consensus.BaseDeposit), new(big.Int).SetUint64(params.Kcoin))
	req.NoError(suite.registerValidator(user, numTokens))

	suite.backend.Commit()

	// generate a random user
	randomUser, err := crypto.GenerateKey()
	req.NoError(err)
	req.NotNil(randomUser)

	testCases := []struct {
		name   string
		input  common.Address
		output bool
	}{
		{
			name:   "super node - genesis validator",
			input:  getAddress(validator),
			output: true,
		},
		{
			name:   "validator, not super node",
			input:  getAddress(user),
			output: false,
		},
		{
			name:   "not validator - random user",
			input:  getAddress(randomUser),
			output: false,
		},
	}
	for _, tc := range testCases {
		suite.T().Run(fmt.Sprintf("name: %s, address %s", tc.name, tc.input.Hex()), func(t *testing.T) {
			isSuperNode, err := suite.validatorMgr.IsSuperNode(&bind.CallOpts{}, tc.input)
			req.NoError(err)
			req.Equal(tc.output, isSuperNode)
		})
	}
}

func (suite *ValidatorMgrSuite) TestGetMinimumDeposit_NotFull() {
	req := suite.Require()

	storedMinDeposit, err := suite.validatorMgr.GetMinimumDeposit(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedMinDeposit)
	req.Equal(new(big.Int).Mul(new(big.Int).SetUint64(suite.opts.Consensus.BaseDeposit), new(big.Int).SetUint64(params.Kcoin)), storedMinDeposit)
}

func (suite *ValidatorMgrSuite) TestGetMinimumDeposit_Full() {
	req := suite.Require()

	storedMinDeposit, err := suite.validatorMgr.GetMinimumDeposit(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedMinDeposit)
	req.Equal(new(big.Int).Add(new(big.Int).Mul(new(big.Int).SetUint64(suite.opts.Consensus.Validators[0].Deposit), new(big.Int).SetUint64(params.Kcoin)), common.Big1), storedMinDeposit)
}

func (suite *ValidatorMgrSuite) TestRegisterValidator_WhenPaused() {
	req := suite.Require()

	suite.pauseService()

	deposit := new(big.Int).Mul(new(big.Int).SetUint64(suite.opts.Consensus.BaseDeposit), new(big.Int).SetUint64(params.Kcoin))
	req.Error(suite.registerValidator(user, deposit), "cannot register the validator because the service is paused")
}

func (suite *ValidatorMgrSuite) TestRegisterValidator_Duplicate() {
	req := suite.Require()

	deposit := new(big.Int).Mul(new(big.Int).SetUint64(suite.opts.Consensus.BaseDeposit), new(big.Int).SetUint64(params.Kcoin))
	req.NoError(suite.registerValidator(user, deposit))
	req.Error(suite.registerValidator(user, deposit), "cannot register the same validator twice")
}

func (suite *ValidatorMgrSuite) TestRegisterValidator_WithoutMinDeposit() {
	req := suite.Require()

	deposit := new(big.Int).Sub(new(big.Int).Mul(new(big.Int).SetUint64(suite.opts.Consensus.BaseDeposit), new(big.Int).SetUint64(params.Kcoin)), common.Big1)
	req.Error(suite.registerValidator(user, deposit), "requires the minimum deposit")
}

func (suite *ValidatorMgrSuite) TestRegister_NotFull_GreaterThan() {
	req := suite.Require()

	initialValidatorCount, err := suite.validatorMgr.GetValidatorCount(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(initialValidatorCount)

	deposit := new(big.Int).Add(new(big.Int).Mul(new(big.Int).SetUint64(suite.opts.Consensus.Validators[0].Deposit), new(big.Int).SetUint64(params.Kcoin)), common.Big1)
	req.NoError(suite.registerValidator(user, deposit))

	suite.backend.Commit()

	req.Equal(new(big.Int).Add(new(big.Int).SetUint64(uint64(len(suite.opts.Consensus.Validators))), common.Big1), suite.getValidatorCount())

	storedValidator := suite.getHighestBidder()
	req.NotZero(storedValidator)
	req.Equal(getAddress(user), storedValidator.Code)
	req.Equal(deposit, storedValidator.Deposit)

	storedDeposit := suite.getCurrentDeposit(user)
	req.NotZero(storedDeposit)
	req.Zero(storedDeposit.AvailableAt.Uint64())
	req.Equal(deposit, storedDeposit.Amount)

	finalValidatorCount, err := suite.validatorMgr.GetValidatorCount(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(finalValidatorCount)
	req.Equal(new(big.Int).Add(initialValidatorCount, common.Big1), finalValidatorCount)
}

func (suite *ValidatorMgrSuite) TestRegister_NotFull_LessOrEqualTo() {
	req := suite.Require()

	initialValidatorCount, err := suite.validatorMgr.GetValidatorCount(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(initialValidatorCount)

	deposit := new(big.Int).Mul(new(big.Int).SetUint64(suite.opts.Consensus.BaseDeposit), new(big.Int).SetUint64(params.Kcoin))
	req.NoError(suite.registerValidator(user, deposit))

	suite.backend.Commit()

	req.Equal(new(big.Int).Add(new(big.Int).SetUint64(uint64(len(suite.opts.Consensus.Validators))), common.Big1), suite.getValidatorCount())

	storedValidator := suite.getHighestBidder()
	req.NotZero(storedValidator)
	req.Equal(getAddress(validator), storedValidator.Code)
	req.Equal(new(big.Int).Mul(new(big.Int).SetUint64(suite.opts.Consensus.Validators[0].Deposit), big.NewInt(params.Kcoin)), storedValidator.Deposit)

	storedDeposit := suite.getCurrentDeposit(user)
	req.NotZero(storedDeposit)
	req.Zero(storedDeposit.AvailableAt.Uint64())
	req.Equal(deposit, storedDeposit.Amount)

	finalValidatorCount, err := suite.validatorMgr.GetValidatorCount(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(finalValidatorCount)
	req.Equal(new(big.Int).Add(initialValidatorCount, common.Big1), finalValidatorCount)
}

func (suite *ValidatorMgrSuite) TestRegister_Full_Replacement() {
	req := suite.Require()

	initialValidatorCount, err := suite.validatorMgr.GetValidatorCount(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(initialValidatorCount)

	minDeposit, err := suite.validatorMgr.GetMinimumDeposit(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(minDeposit)
	req.NoError(suite.registerValidator(user, minDeposit))

	suite.backend.Commit()

	req.Equal(new(big.Int).SetUint64(suite.opts.Consensus.MaxNumValidators), suite.getValidatorCount())

	storedValidator := suite.getHighestBidder()
	req.NotZero(storedValidator)
	req.Equal(getAddress(user), storedValidator.Code)
	req.Equal(minDeposit, storedValidator.Deposit)

	storedDeposit := suite.getCurrentDeposit(user)
	req.NotZero(storedDeposit)
	req.Zero(storedDeposit.AvailableAt.Uint64())
	req.Equal(minDeposit, storedDeposit.Amount)

	finalValidatorCount, err := suite.validatorMgr.GetValidatorCount(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(finalValidatorCount)
	req.Equal(initialValidatorCount, finalValidatorCount)

}

func (suite *ValidatorMgrSuite) TestDeregisterValidator_WhenPaused() {
	req := suite.Require()

	suite.pauseService()

	req.Error(suite.deregisterValidator(validator), "cannot deregister the validator because the service is paused")
}

func (suite *ValidatorMgrSuite) TestDeregisterValidator_NotValidator() {
	req := suite.Require()

	req.Error(suite.deregisterValidator(user), "cannot deregister a non-validator")
}

func (suite *ValidatorMgrSuite) TestDeregisterValidator() {
	req := suite.Require()

	req.NoError(suite.deregisterValidator(validator))

	suite.backend.Commit()

	deposit := suite.getCurrentDeposit(validator)
	req.NotNil(deposit)

	req.True(deposit.AvailableAt.Cmp(common.Big0) > 0)
}

func (suite *ValidatorMgrSuite) TestReleaseDeposits_WhenPaused() {
	req := suite.Require()

	suite.pauseService()

	req.Error(suite.releaseDeposits(user), "cannot release deposits because the service is paused")
}

func (suite *ValidatorMgrSuite) TestReleaseDeposits_NoDeposits() {
	req := suite.Require()

	initialBalance := suite.balanceOf(getAddress(user))
	req.NotNil(initialBalance)

	req.NoError(suite.releaseDeposits(user))

	suite.backend.Commit()

	depositCount := suite.getDepositCount(user)
	req.NotNil(depositCount)
	req.Zero(depositCount.Uint64())

	finalBalance := suite.balanceOf(getAddress(user))
	req.NotNil(finalBalance)
	req.Equal(initialBalance, finalBalance)
}

func (suite *ValidatorMgrSuite) TestReleaseDeposits_LockedDeposit() {
	req := suite.Require()

	initialBalance := suite.balanceOf(getAddress(validator))
	req.NotNil(initialBalance)

	req.NoError(suite.deregisterValidator(validator))
	req.NoError(suite.releaseDeposits(validator))

	suite.backend.Commit()

	depositCount := suite.getDepositCount(validator)
	req.NotNil(depositCount)
	req.Equal(common.Big1, depositCount)

	finalBalance := suite.balanceOf(getAddress(validator))
	req.NotNil(finalBalance)
	req.Equal(initialBalance, finalBalance)
}

func (suite *ValidatorMgrSuite) TestReleaseDeposits_UnlockedDeposit() {
	req := suite.Require()

	initialBalance := suite.balanceOf(getAddress(validator))
	req.NotNil(initialBalance)

	deposit := suite.getCurrentDeposit(validator)

	req.NoError(suite.deregisterValidator(validator))
	req.NoError(suite.releaseDeposits(validator))

	suite.backend.Commit()

	depositCount := suite.getDepositCount(validator)
	req.NotNil(depositCount)
	req.Zero(depositCount.Uint64())

	finalBalance := suite.balanceOf(getAddress(validator))
	req.NotNil(finalBalance)
	req.Equal(new(big.Int).Add(initialBalance, deposit.Amount), finalBalance)
}

func (suite *ValidatorMgrSuite) mintTokens(governor *ecdsa.PrivateKey, to *ecdsa.PrivateKey, numTokens *big.Int) {
	req := suite.Require()

	// mint enough tokens to the new user (submit & confirm)
	tokenABI, err := abi.JSON(strings.NewReader(consensus.MiningTokenABI))
	req.NoError(err)
	req.NotNil(tokenABI)

	mintParams, err := tokenABI.Pack(
		"mint",
		getAddress(to),
		numTokens,
	)
	req.NoError(err)
	req.NotZero(mintParams)

	transactOpts := bind.NewKeyedTransactor(governor)
	_, err = suite.multiSig.SubmitTransaction(transactOpts, tokenAddr, common.Big0, mintParams)
	req.NoError(err)
}

func (suite *ValidatorMgrSuite) registerValidator(user *ecdsa.PrivateKey, deposit *big.Int) error {
	transferOpts := bind.NewKeyedTransactor(user)
	_, err := suite.miningToken.Transfer(transferOpts, validatorMgrAddr, deposit, []byte("not_zero"), consensus.RegistrationHandler)
	return err
}

func (suite *ValidatorMgrSuite) deregisterValidator(user *ecdsa.PrivateKey) error {
	transactOpts := bind.NewKeyedTransactor(user)
	_, err := suite.validatorMgr.DeregisterValidator(transactOpts)
	return err
}

func (suite *ValidatorMgrSuite) releaseDeposits(user *ecdsa.PrivateKey) error {
	transactOpts := bind.NewKeyedTransactor(user)
	_, err := suite.validatorMgr.ReleaseDeposits(transactOpts)
	return err
}

func (suite *ValidatorMgrSuite) pauseService() {
	req := suite.Require()

	// pause the service
	validatorMgrABI, err := abi.JSON(strings.NewReader(consensus.ValidatorMgrABI))
	req.NoError(err)
	req.NotNil(validatorMgrABI)

	pauseParams, err := validatorMgrABI.Pack("pause")
	req.NoError(err)
	req.NotZero(pauseParams)

	transactOpts := bind.NewKeyedTransactor(governor)
	_, err = suite.multiSig.SubmitTransaction(transactOpts, validatorMgrAddr, common.Big0, pauseParams)
	req.NoError(err)
}

func (suite *ValidatorMgrSuite) balanceOf(user common.Address) *big.Int {
	req := suite.Require()

	balance, err := suite.miningToken.BalanceOf(&bind.CallOpts{}, user)
	req.NoError(err)
	req.NotNil(balance)

	return balance
}

func (suite *ValidatorMgrSuite) getValidatorCount() *big.Int {
	req := suite.Require()

	validatorCount, err := suite.validatorMgr.GetValidatorCount(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(validatorCount)

	return validatorCount
}

func (suite *ValidatorMgrSuite) getHighestBidder() struct {
	Code    common.Address
	Deposit *big.Int
} {
	req := suite.Require()

	registration, err := suite.validatorMgr.GetValidatorAtIndex(&bind.CallOpts{}, common.Big0)
	req.NoError(err)
	req.NotZero(registration)

	return registration
}

func (suite *ValidatorMgrSuite) getDepositCount(user *ecdsa.PrivateKey) *big.Int {
	req := suite.Require()

	depositCount, err := suite.validatorMgr.GetDepositCount(&bind.CallOpts{From: getAddress(user)})
	req.NoError(err)
	req.NotNil(depositCount)

	return depositCount
}

func (suite *ValidatorMgrSuite) getCurrentDeposit(user *ecdsa.PrivateKey) struct {
	Amount      *big.Int
	AvailableAt *big.Int
} {
	req := suite.Require()

	deposit, err := suite.validatorMgr.GetDepositAtIndex(&bind.CallOpts{From: getAddress(user)}, common.Big0)
	req.NoError(err)
	req.NotZero(deposit)

	return deposit
}

// dtos converts days to seconds
func dtos(numberOfDays *big.Int) *big.Int {
	return new(big.Int).Mul(numberOfDays, secondsPerDay)
}

// getAddress return the address of the given private key
func getAddress(privateKey *ecdsa.PrivateKey) common.Address {
	return crypto.PubkeyToAddress(privateKey.PublicKey)
}
