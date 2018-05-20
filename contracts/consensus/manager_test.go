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
	"github.com/kowala-tech/kcoin/crypto"
	"github.com/kowala-tech/kcoin/kcoin/genesis"
	"github.com/kowala-tech/kcoin/params"
	"github.com/stretchr/testify/require"
)

const (
	initialBalance = 10 // 10 kUSD
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

func TestContractCreation(t *testing.T) {
	// create genesis
	opts := getDefaultOpts()
	require.NotNil(t, opts)
	gen, err := genesis.New(opts)
	require.NoError(t, err)
	require.NotNil(t, gen)

	// create backend
	backend := backends.NewSimulatedBackend(gen.Alloc)
	require.NotNil(t, backend)

	// ValidatorMgr instance
	mgr, err := NewValidatorMgr(validatorMgrAddr, backend)
	require.NoError(t, err)
	require.NotNil(t, mgr)

	// MiningToken instance
	mtoken, err := token.NewMiningToken(tokenAddr, backend)
	require.NoError(t, err)
	require.NotNil(t, mtoken)

	storedBaseDeposit, err := mgr.BaseDeposit(&bind.CallOpts{})
	require.NoError(t, err)
	require.NotNil(t, storedBaseDeposit)
	require.Equal(t, new(big.Int).Mul(new(big.Int).SetUint64(opts.Consensus.BaseDeposit), new(big.Int).SetUint64(params.Ether)), storedBaseDeposit)

	storedFreezePeriod, err := mgr.FreezePeriod(&bind.CallOpts{})
	require.NoError(t, err)
	require.NotNil(t, storedFreezePeriod)
	require.Equal(t, dtos(new(big.Int).SetUint64(opts.Consensus.FreezePeriod)), storedFreezePeriod)

	storedMaxNumValidators, err := mgr.MaxNumValidators(&bind.CallOpts{})
	require.NoError(t, err)
	require.NotNil(t, storedMaxNumValidators)
	require.Equal(t, opts.Consensus.MaxNumValidators, storedMaxNumValidators.Uint64())

	storedMiningTokenAddr, err := mgr.MiningTokenAddr(&bind.CallOpts{})
	require.NoError(t, err)
	require.NotNil(t, storedMiningTokenAddr)
	require.Equal(t, tokenAddr, storedMiningTokenAddr)

	validatorCount, err := mgr.GetValidatorCount(&bind.CallOpts{})
	require.NoError(t, err)
	require.NotNil(t, validatorCount)
	require.Equal(t, int64(len(opts.Consensus.Validators)), validatorCount.Int64())

	var deposits uint64
	for _, validator := range opts.Consensus.Validators {
		validatorAddr := common.HexToAddress(validator.Address)
		isValidator, err := mgr.IsValidator(&bind.CallOpts{}, validatorAddr)
		require.NoError(t, err)
		require.True(t, isValidator)

		storedDepositCount, err := mgr.GetDepositCount(&bind.CallOpts{From: validatorAddr})
		require.NoError(t, err)
		require.NotNil(t, storedDepositCount)
		require.Equal(t, common.Big1, storedDepositCount)

		deposit, err := mgr.GetDepositAtIndex(&bind.CallOpts{From: validatorAddr}, common.Big0)
		require.NoError(t, err)
		require.NotNil(t, deposit)
		require.Zero(t, deposit.AvailableAt.Uint64())
		require.Equal(t, new(big.Int).Mul(new(big.Int).SetUint64(opts.Consensus.BaseDeposit), new(big.Int).SetUint64(params.Ether)), deposit.Amount)

		deposits += validator.Deposit
	}

	mgrBalance, err := mtoken.BalanceOf(&bind.CallOpts{}, validatorMgrAddr)
	require.NoError(t, err)
	require.NotNil(t, mgrBalance)
	require.Equal(t, new(big.Int).Mul(new(big.Int).SetUint64(deposits), new(big.Int).SetUint64(params.Ether)), mgrBalance)
}

func TestIsValidator(t *testing.T) {
	// create genesis
	opts := getDefaultOpts()
	require.NotNil(t, opts)
	genesis, err := genesis.New(opts)
	require.NoError(t, err)
	require.NotNil(t, genesis)

	backend := backends.NewSimulatedBackend(genesis.Alloc)
	require.NotNil(t, backend)

	// ValidatorMgr instance
	mgr, err := NewValidatorMgr(validatorMgrAddr, backend)
	require.NoError(t, err)
	require.NotNil(t, mgr)

	testCases := []struct {
		input  common.Address
		output bool
	}{
		{
			input:  getAddress(validator),
			output: true,
		},
		{
			input:  getAddress(user),
			output: false,
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Address %s", tc.input.Hex()), func(t *testing.T) {
			isValidator, err := mgr.IsValidator(&bind.CallOpts{}, tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.output, isValidator)
		})
	}
}

func TestIsGenesisValidator(t *testing.T) {
	// create genesis
	opts := getDefaultOpts()
	require.NotNil(t, opts)
	genesis, err := genesis.New(opts)
	require.NoError(t, err)
	require.NotNil(t, genesis)

	backend := backends.NewSimulatedBackend(genesis.Alloc)

	// ValidatorMgr instance
	mgr, err := NewValidatorMgr(validatorMgrAddr, backend)
	require.NoError(t, err)
	require.NotNil(t, mgr)

	testCases := []struct {
		input  common.Address
		output bool
	}{
		{
			input:  getAddress(validator),
			output: true,
		},
		{
			input:  getAddress(user),
			output: false,
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Address %s", tc.input.Hex()), func(t *testing.T) {
			isValidator, err := mgr.IsGenesisValidator(&bind.CallOpts{}, tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.output, isValidator)
		})
	}
}

func TestGetMinimumDeposit_NotFull(t *testing.T) {
	// create genesis
	opts := getDefaultOpts()
	opts.Consensus.MaxNumValidators = 10
	require.NotNil(t, opts)
	genesis, err := genesis.New(opts)
	require.NoError(t, err)
	require.NotNil(t, genesis)

	backend := backends.NewSimulatedBackend(genesis.Alloc)

	// ValidatorMgr instance
	mgr, err := NewValidatorMgr(validatorMgrAddr, backend)
	require.NoError(t, err)
	require.NotNil(t, mgr)

	storedMinDeposit, err := mgr.GetMinimumDeposit(&bind.CallOpts{})
	require.NoError(t, err)
	require.NotNil(t, storedMinDeposit)
	require.Equal(t, new(big.Int).Mul(new(big.Int).SetUint64(opts.Consensus.BaseDeposit), new(big.Int).SetUint64(params.Ether)), storedMinDeposit)
}

func TestGetMinimumDeposit_Full(t *testing.T) {
	// create genesis
	opts := getDefaultOpts()
	opts.Consensus.MaxNumValidators = 1
	require.NotNil(t, opts)
	genesis, err := genesis.New(opts)
	require.NoError(t, err)
	require.NotNil(t, genesis)

	backend := backends.NewSimulatedBackend(genesis.Alloc)

	// ValidatorMgr instance
	mgr, err := NewValidatorMgr(validatorMgrAddr, backend)
	require.NoError(t, err)
	require.NotNil(t, mgr)

	storedMinDeposit, err := mgr.GetMinimumDeposit(&bind.CallOpts{})
	require.NoError(t, err)
	require.NotNil(t, storedMinDeposit)
	require.Equal(t, new(big.Int).Add(new(big.Int).Mul(new(big.Int).SetUint64(opts.Consensus.BaseDeposit), new(big.Int).SetUint64(params.Ether)), common.Big1), storedMinDeposit)
}

func TestRegisterValidator(t *testing.T) {
	// @TODO (rgeraldes) - already validator
	// @TODO (rgeraldes) - insufficient deposit
	// @TODO (rgeraldes) - insert greater than
	// @TODO (rgeraldes) - insert less or equal to

	// create genesis
	opts := getDefaultOpts()
	require.NotNil(t, opts)
	genesis, err := genesis.New(opts)
	require.NoError(t, err)
	require.NotNil(t, genesis)

	backend := backends.NewSimulatedBackend(genesis.Alloc)

	// ValidatorMgr instance
	mgr, err := NewValidatorMgr(validatorMgrAddr, backend)
	require.NoError(t, err)
	require.NotNil(t, mgr)

	// multiSig instance
	multiSig, err := ownership.NewMultiSigWallet(multiSigAddr, backend)
	require.NoError(t, err)
	require.NotNil(t, multiSig)

	// mint enough tokens to the new user (submit & confirm)
	tokenABI, err := abi.JSON(strings.NewReader(token.MiningTokenABI))
	require.NoError(t, err)
	require.NotNil(t, tokenABI)

	numTokens := new(big.Int).Mul(new(big.Int).SetUint64(opts.Consensus.BaseDeposit), new(big.Int).SetUint64(params.Ether))
	mintParams, err := tokenABI.Pack(
		"mint",
		getAddress(user),
		numTokens,
	)
	require.NoError(t, err)
	require.NotZero(t, mintParams)

	transactOpts := bind.NewKeyedTransactor(governor)
	_, err = multiSig.SubmitTransaction(transactOpts, tokenAddr, common.Big0, mintParams)
	require.NoError(t, err)

	// MiningToken instance
	mtoken, err := token.NewMiningToken(tokenAddr, backend)
	require.NoError(t, err)
	require.NotNil(t, mtoken)

	// register new validator
	transferOpts := bind.NewKeyedTransactor(user)
	_, err = mtoken.Transfer(transferOpts, validatorMgrAddr, numTokens, []byte("not_zero"), registrationHandler)
	require.NoError(t, err)

	backend.Commit()

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

func TestDeregisterValidator(t *testing.T) {
	// @TODO (rgeraldes) - deposit available at > 0
	// @TODO (rgeraldes) - not validator

	// create genesis
	opts := getDefaultOpts()
	require.NotNil(t, opts)
	genesis, err := genesis.New(opts)
	require.NoError(t, err)
	require.NotNil(t, genesis)

	backend := backends.NewSimulatedBackend(genesis.Alloc)

	// ValidatorMgr instance
	mgr, err := NewValidatorMgr(validatorMgrAddr, backend)
	require.NoError(t, err)
	require.NotNil(t, mgr)

	isValidator, err := mgr.IsValidator(&bind.CallOpts{}, getAddress(validator))
	require.NoError(t, err)
	require.True(t, isValidator)

	// deregister the genesis validator
	transactOpts := bind.NewKeyedTransactor(validator)
	_, err = mgr.DeregisterValidator(transactOpts)
	require.NoError(t, err)

	backend.Commit()

	isValidator, err = mgr.IsValidator(&bind.CallOpts{}, getAddress(validator))
	require.NoError(t, err)
	require.False(t, isValidator)
}

func TestReleaseDeposits(t *testing.T) {
	// @TODO (rgeraldes) - no deposits
	// @TODO (rgeraldes) - locked deposits
	// @TODO (rgeraldes) - unlocked deposits

	// create genesis
	opts := getDefaultOpts()
	opts.Consensus.FreezePeriod = 0
	require.NotNil(t, opts)
	genesis, err := genesis.New(opts)
	require.NoError(t, err)
	require.NotNil(t, genesis)

	backend := backends.NewSimulatedBackend(genesis.Alloc)

	// ValidatorMgr instance
	mgr, err := NewValidatorMgr(validatorMgrAddr, backend)
	require.NoError(t, err)
	require.NotNil(t, mgr)

	// MiningToken instance
	mtoken, err := token.NewMiningToken(tokenAddr, backend)
	require.NoError(t, err)
	require.NotNil(t, mtoken)

	// deregister the genesis validator
	transactOpts := bind.NewKeyedTransactor(validator)
	_, err = mgr.DeregisterValidator(transactOpts)
	require.NoError(t, err)

	backend.Commit()

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
	require.NotNil(t, balance)
}

// dtos converts days to seconds
func dtos(numOfDays *big.Int) *big.Int {
	return new(big.Int).Mul(numOfDays, secondsPerDay)
}

// getAddress return the address of the given private key
func getAddress(privateKey *ecdsa.PrivateKey) common.Address {
	return crypto.PubkeyToAddress(privateKey.PublicKey)
}
