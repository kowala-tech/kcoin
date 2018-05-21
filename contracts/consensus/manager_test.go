package consensus

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/kowala-tech/kcoin/accounts/abi"
	"github.com/kowala-tech/kcoin/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/accounts/abi/bind/backends"
	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/contracts/ownership"
	"github.com/kowala-tech/kcoin/contracts/token"
	"github.com/kowala-tech/kcoin/core"
	"github.com/kowala-tech/kcoin/crypto"
	"github.com/kowala-tech/kcoin/kcoin/genesis"
	"github.com/kowala-tech/kcoin/params"
	"github.com/stretchr/testify/suite"
)

var (
	validator, _     = crypto.GenerateKey()
	user, _          = crypto.GenerateKey()
	author, _        = crypto.HexToECDSA("bfef37ae9ac5d5e7ebbbefc19f4e1f572a7ca7aa0d28e527b7d62950951cc5eb")
	governor, _      = crypto.GenerateKey()
	validatorMgrAddr = common.HexToAddress("0x161ad311F1D66381C17641b1B73042a4CA731F9f")
	multiSigAddr     = common.HexToAddress("0xA143ac5ec5D95f16aFD5Fc3B09e0aDaf360ffC9e")
	tokenAddr        = common.HexToAddress("0xB012F49629258C9c35b2bA80cD3dc3C841d9719D")
	secondsPerDay    = new(big.Int).SetUint64(86400)
)

func getDefaultOpts() *genesis.Options {
	baseDeposit := uint64(20)
	tokenHolder := genesis.TokenHolder{
		Address:   getAddress(validator).Hex(),
		NumTokens: baseDeposit,
	}

	opts := &genesis.Options{
		Network: "test",
		Consensus: &genesis.ConsensusOpts{
			Engine:           "tendermint",
			MaxNumValidators: 10,
			FreezePeriod:     30,
			BaseDeposit:      baseDeposit,
			Validators: []genesis.Validator{genesis.Validator{
				Address: tokenHolder.Address,
				Deposit: tokenHolder.NumTokens,
			}},
			MiningToken: &genesis.MiningTokenOpts{
				Name:     "mUSD",
				Symbol:   "mUSD",
				Cap:      1000,
				Decimals: 18,
				Holders:  []genesis.TokenHolder{tokenHolder},
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
		},
		PrefundedAccounts: []genesis.PrefundedAccount{
			genesis.PrefundedAccount{
				Address: tokenHolder.Address,
				Balance: 10,
			},
			genesis.PrefundedAccount{
				Address: getAddress(governor).Hex(),
				Balance: 10,
			},
			genesis.PrefundedAccount{
				Address: getAddress(user).Hex(),
				Balance: 10,
			},
		},
	}

	return opts
}

type ValidatorMgrSuite struct {
	suite.Suite
	backend      *backends.SimulatedBackend
	opts         *genesis.Options
	validatorMgr *ValidatorMgr
	multiSig     *ownership.MultiSigWallet
	miningToken  *token.MiningToken
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
	case testName == "TestGetMinimumDeposit_Full":
		opts.Consensus.MaxNumValidators = 1
	}
	suite.opts = opts

	genesis, err := genesis.New(opts)
	req.NoError(err)
	req.NotNil(genesis)

	backend := backends.NewSimulatedBackend(genesis.Alloc)
	req.NotNil(backend)
	suite.backend = backend

	// ValidatorMgr instance
	mgr, err := NewValidatorMgr(validatorMgrAddr, backend)
	req.NoError(err)
	req.NotNil(mgr)
	suite.validatorMgr = mgr

	// multiSig instance
	multiSig, err := ownership.NewMultiSigWallet(multiSigAddr, backend)
	req.NoError(err)
	req.NotNil(multiSig)
	suite.multiSig = multiSig

	// MiningToken instance
	mToken, err := token.NewMiningToken(tokenAddr, backend)
	req.NoError(err)
	req.NotNil(mToken)
	suite.miningToken = mToken

}

func (suite *ValidatorMgrSuite) TestDeploy() {
	req := suite.Require()

	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		getAddress(governor): core.GenesisAccount{
			Balance: new(big.Int).Mul(new(big.Int).SetUint64(100), new(big.Int).SetUint64(params.Ether)),
		},
	})

	baseDeposit := new(big.Int).SetUint64(100)
	maxNumValidators := new(big.Int).SetUint64(100)
	freezePeriod := new(big.Int).SetUint64(100)

	transactOpts := bind.NewKeyedTransactor(governor)
	_, _, mgr, err := DeployValidatorMgr(transactOpts, backend, baseDeposit, maxNumValidators, freezePeriod, tokenAddr)
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

func (suite *ValidatorMgrSuite) TestDeploy_MaxNumValidators_Zero() {
	req := suite.Require()

	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		getAddress(governor): core.GenesisAccount{
			Balance: new(big.Int).Mul(new(big.Int).SetUint64(100), new(big.Int).SetUint64(params.Ether)),
		},
	})

	baseDeposit := new(big.Int).SetUint64(100)
	maxNumValidators := common.Big0
	freezePeriod := new(big.Int).SetUint64(100)

	transactOpts := bind.NewKeyedTransactor(governor)
	_, _, _, err := DeployValidatorMgr(transactOpts, backend, baseDeposit, maxNumValidators, freezePeriod, tokenAddr)
	req.Error(err, "maximum number of validators cannot be zero")
}

func (suite *ValidatorMgrSuite) TestIsValidator() {
	req := suite.Require()

	// @TODO (add case where user registers and deregisters)

	numTokens := new(big.Int).Mul(new(big.Int).SetUint64(suite.opts.Consensus.BaseDeposit), new(big.Int).SetUint64(params.Ether))
	suite.mintTokens(governor, user, numTokens)
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
			output: true,
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

	numTokens := new(big.Int).Mul(new(big.Int).SetUint64(suite.opts.Consensus.BaseDeposit), new(big.Int).SetUint64(params.Ether))
	suite.mintTokens(governor, user, numTokens)
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

func (suite *ValidatorMgrSuite) TestGetMinimumDeposit_NotFull() {
	req := suite.Require()

	storedMinDeposit, err := suite.validatorMgr.GetMinimumDeposit(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedMinDeposit)
	req.Equal(new(big.Int).Mul(new(big.Int).SetUint64(suite.opts.Consensus.BaseDeposit), new(big.Int).SetUint64(params.Ether)), storedMinDeposit)
}

func (suite *ValidatorMgrSuite) TestGetMinimumDeposit_Full() {
	req := suite.Require()

	storedMinDeposit, err := suite.validatorMgr.GetMinimumDeposit(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedMinDeposit)
	req.Equal(new(big.Int).Add(new(big.Int).Mul(new(big.Int).SetUint64(suite.opts.Consensus.BaseDeposit), new(big.Int).SetUint64(params.Ether)), common.Big1), storedMinDeposit)
}

/*
func (suite *ValidatorMgrSuite) TestRegisterValidator() {
	// @TODO (rgeraldes) - insert greater than
	// @TODO (rgeraldes) - insert less or equal to
	// @TODO (rgeraldes) - normal call

	numTokens := new(big.Int).Mul(new(big.Int).SetUint64(suite.opts.Consensus.BaseDeposit), new(big.Int).SetUint64(params.Ether))
	suite.mintTokens(governor, user, numTokens)
	suite.registerValidator(user, numTokens)

	suite.backend.Commit()

	isValidator, err := mgr.IsValidator(&bind.CallOpts{}, getAddress(user))
	require.NoError(t, err)
	require.True(t, isValidator)

	isGenesis, err := mgr.IsGenesisValidator(&bind.CallOpts{}, getAddress(user))
	require.NoError(t, err)
	require.False(t, isGenesis)

	//balance, err := mtoken.BalanceOf(&bind.CallOpts{}, validatorMgrAddr)
	//require.NoError(t, err)
	//require.NotNil(t, balance)
	//require.Equal(t, new(big.Int).Mul(new(big.Int).SetUint64(opts.Consensus.Validators[0].Deposit), new(big.Int).SetUint64(params.Ether)), balance)
}
*/

func (suite *ValidatorMgrSuite) TestRegisterValidator_Duplicate() {
	req := suite.Require()

	numTokens := new(big.Int).Mul(new(big.Int).SetUint64(suite.opts.Consensus.BaseDeposit), new(big.Int).SetUint64(params.Ether))
	suite.mintTokens(governor, user, new(big.Int).Mul(numTokens, common.Big2))
	req.NoError(suite.registerValidator(user, numTokens))
	req.Error(suite.registerValidator(user, numTokens), "cannot register the same validator twice")
}

func (suite *ValidatorMgrSuite) TestRegisterValidator_InsufficientDeposit() {
	req := suite.Require()

	numTokens := new(big.Int).Mul(new(big.Int).SetUint64(suite.opts.Consensus.BaseDeposit), new(big.Int).SetUint64(params.Ether))
	suite.mintTokens(governor, user, new(big.Int).Sub(numTokens, common.Big1))
	req.Error(suite.registerValidator(user, numTokens), "requires a minimum deposit")
}

func (suite *ValidatorMgrSuite) TestDeregisterValidator() {
	req := suite.Require()

	req.NoError(suite.deregisterValidator(validator))

	suite.backend.Commit()

	isValidator, err := suite.validatorMgr.IsValidator(&bind.CallOpts{}, getAddress(validator))
	req.NoError(err)
	req.False(isValidator)

	// @TODO (rgeraldes) - deposit available at > 0
}

func (suite *ValidatorMgrSuite) TestDeregisterValidator_NotValidator() {
	req := suite.Require()

	req.Error(suite.deregisterValidator(user), "cannot deregister a non validator user")
}

func (suite *ValidatorMgrSuite) TestReleaseDeposits() {
	// @TODO (rgeraldes) - no deposits
	// @TODO (rgeraldes) - locked deposits
	// @TODO (rgeraldes) - unlocked deposits

	req := suite.Require()

	req.NoError(suite.deregisterValidator(validator))

	suite.backend.Commit()

	/*
		deposit, err := mgr.GetDepositAtIndex(&bind.CallOpts{From: getAddress(validator)}, common.Big0)
		require.NoError(t, err)
		require.NotNil(t, deposit)
		require.Equal(t, new(big.Int).Mul(new(big.Int).SetUint64(opts.Consensus.BaseDeposit), new(big.Int).SetUint64(params.Ether)), deposit.Amount)

		// release deposit
		_, err = mgr.ReleaseDeposits(transactOpts)
		require.NoError(t, err)

		backend.Commit()

		balance, err := mtoken.BalanceOf(&bind.CallOpts{}, getAddress(validator))
		require.NoError(t, err)
		require.NotNil(t, balance)*/
}

func (suite *ValidatorMgrSuite) mintTokens(governor *ecdsa.PrivateKey, to *ecdsa.PrivateKey, numTokens *big.Int) {
	req := suite.Require()

	// mint enough tokens to the new user (submit & confirm)
	tokenABI, err := abi.JSON(strings.NewReader(token.MiningTokenABI))
	req.NoError(err)
	req.NotNil(tokenABI)

	mintParams, err := tokenABI.Pack(
		"mint",
		getAddress(user),
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
	_, err := suite.miningToken.Transfer(transferOpts, validatorMgrAddr, deposit, []byte("not_zero"), registrationHandler)
	return err
}

func (suite *ValidatorMgrSuite) deregisterValidator(user *ecdsa.PrivateKey) error {
	transactOpts := bind.NewKeyedTransactor(user)
	_, err := suite.validatorMgr.DeregisterValidator(transactOpts)
	return err
}

// dtos converts days to seconds
func dtos(numberOfDays *big.Int) *big.Int {
	return new(big.Int).Mul(numberOfDays, secondsPerDay)
}

// getAddress return the address of the given private key
func getAddress(privateKey *ecdsa.PrivateKey) common.Address {
	return crypto.PubkeyToAddress(privateKey.PublicKey)
}
