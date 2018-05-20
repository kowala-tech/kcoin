package token

import (
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/kowala-tech/kcoin/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/accounts/abi/bind/backends"
	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/contracts/consensus"
	"github.com/kowala-tech/kcoin/contracts/oracle"
	"github.com/kowala-tech/kcoin/core"
	"github.com/kowala-tech/kcoin/crypto"
	"github.com/kowala-tech/kcoin/params"
	"github.com/stretchr/testify/require"
)

var (
	owner, _            = crypto.GenerateKey()
	user, _             = crypto.GenerateKey()
	initialBalance      = new(big.Int).Mul(new(big.Int).SetUint64(10), new(big.Int).SetUint64(params.Ether))         // 10 kUSD
	miningTokenCap      = new(big.Int).Mul(new(big.Int).SetUint64(1073741824), new(big.Int).SetUint64(params.Ether)) // 1073741824 mUSD
	miningToken         = "mUSD"
	miningTokenDecimals = uint8(18)
	customFallback      = "test()"
)

func TestDeploy(t *testing.T) {
	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		getAddress(owner): core.GenesisAccount{Balance: initialBalance},
	})

	transactOpts := bind.NewKeyedTransactor(owner)
	_, _, token, err := DeployMiningToken(transactOpts, backend, miningToken, miningToken, miningTokenCap, miningTokenDecimals)
	require.NoError(t, err)
	require.NotNil(t, token)

	backend.Commit()

	storedCap, err := token.Cap(&bind.CallOpts{})
	require.NoError(t, err)
	require.NotNil(t, storedCap)
	require.Equal(t, miningTokenCap, storedCap)

	storedDecimals, err := token.Decimals(&bind.CallOpts{})
	require.NoError(t, err)
	require.NotNil(t, storedDecimals)
	require.Equal(t, miningTokenDecimals, storedDecimals)

	storedTotalSupply, err := token.TotalSupply(&bind.CallOpts{})
	require.NoError(t, err)
	require.NotNil(t, storedTotalSupply)
	require.Zero(t, storedTotalSupply.Uint64())

	storedName, err := token.Name(&bind.CallOpts{})
	require.NoError(t, err)
	require.NotNil(t, storedName)
	require.Equal(t, miningToken, storedName)

	storedSymbol, err := token.Symbol(&bind.CallOpts{})
	require.NoError(t, err)
	require.NotNil(t, storedSymbol)
	require.Equal(t, miningToken, storedSymbol)

	storedMintingFinished, err := token.MintingFinished(&bind.CallOpts{})
	require.NoError(t, err)
	require.NotNil(t, storedMintingFinished)
	require.False(t, storedMintingFinished)

	storedOwner, err := token.Owner(&bind.CallOpts{})
	require.NoError(t, err)
	require.NotNil(t, owner)
	require.Equal(t, getAddress(owner), storedOwner)
}

func TestFinishMinting(t *testing.T) {
	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		getAddress(owner): core.GenesisAccount{Balance: initialBalance},
	})

	transactOpts := bind.NewKeyedTransactor(owner)
	_, _, token, err := DeployMiningToken(transactOpts, backend, miningToken, miningToken, miningTokenCap, miningTokenDecimals)
	require.NoError(t, err)
	require.NotNil(t, token)

	_, err = token.FinishMinting(transactOpts)
	require.NoError(t, err)

	backend.Commit()

	finished, err := token.MintingFinished(&bind.CallOpts{})
	require.NoError(t, err)
	require.True(t, finished)
}

func TestFinishMinting_NotOwner(t *testing.T) {
	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		getAddress(owner): core.GenesisAccount{Balance: initialBalance},
	})

	transactOpts := bind.NewKeyedTransactor(owner)
	_, _, token, err := DeployMiningToken(transactOpts, backend, miningToken, miningToken, miningTokenCap, miningTokenDecimals)
	require.NoError(t, err)
	require.NotNil(t, token)

	mintingOpts := bind.NewKeyedTransactor(user)
	_, err = token.FinishMinting(mintingOpts)
	require.Error(t, err, "Method FinishMinting is just available to the contract owner")
}

func TestFinishMinting_FinishedMinting(t *testing.T) {
	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		getAddress(owner): core.GenesisAccount{Balance: initialBalance},
	})

	transactOpts := bind.NewKeyedTransactor(owner)
	_, _, token, err := DeployMiningToken(transactOpts, backend, miningToken, miningToken, miningTokenCap, miningTokenDecimals)
	require.NoError(t, err)
	require.NotNil(t, token)

	_, err = token.FinishMinting(transactOpts)
	require.NoError(t, err)

	_, err = token.FinishMinting(transactOpts)
	require.Error(t, err, "The FinishMinting operation is not available if the operation is finished")
}

func TestMint(t *testing.T) {
	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		getAddress(owner): core.GenesisAccount{Balance: initialBalance},
	})

	transactOpts := bind.NewKeyedTransactor(owner)
	_, _, token, err := DeployMiningToken(transactOpts, backend, miningToken, miningToken, miningTokenCap, miningTokenDecimals)
	require.NoError(t, err)
	require.NotNil(t, token)

	// mint all the tokens to an user
	numTokens := miningTokenCap
	_, err = token.Mint(transactOpts, getAddress(user), numTokens)
	require.NoError(t, err)

	backend.Commit()

	storedTotalSupply, err := token.TotalSupply(&bind.CallOpts{})
	require.NoError(t, err)
	require.NotNil(t, storedTotalSupply)
	require.Equal(t, numTokens, storedTotalSupply)

	userBalance, err := token.BalanceOf(&bind.CallOpts{}, getAddress(user))
	require.NoError(t, err)
	require.NotNil(t, userBalance)
	require.Equal(t, numTokens, userBalance)
}

func TestMint_NotOwner(t *testing.T) {
	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		getAddress(owner): core.GenesisAccount{Balance: initialBalance},
	})

	transactOpts := bind.NewKeyedTransactor(owner)
	_, _, token, err := DeployMiningToken(transactOpts, backend, miningToken, miningToken, miningTokenCap, miningTokenDecimals)
	require.NoError(t, err)
	require.NotNil(t, token)

	numTokens := miningTokenCap
	mintingOpts := bind.NewKeyedTransactor(user)
	_, err = token.Mint(mintingOpts, getAddress(user), numTokens)
	require.Error(t, err, "Method Mint is just available to the contract owner")
}

func TestMint_OverCap(t *testing.T) {
	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		getAddress(owner): core.GenesisAccount{Balance: initialBalance},
	})

	transactOpts := bind.NewKeyedTransactor(owner)
	_, _, token, err := DeployMiningToken(transactOpts, backend, miningToken, miningToken, miningTokenCap, miningTokenDecimals)
	require.NoError(t, err)
	require.NotNil(t, token)

	numTokens := new(big.Int).Add(miningTokenCap, common.Big1)
	_, err = token.Mint(transactOpts, getAddress(user), numTokens)
	require.Error(t, err, "Cannot mint over the cap")
}

func TestTransfer_CustomFallback_Compatible(t *testing.T) {
	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		getAddress(owner): core.GenesisAccount{Balance: initialBalance},
		getAddress(user):  core.GenesisAccount{Balance: initialBalance},
	})

	transactOpts := bind.NewKeyedTransactor(owner)
	tokenAddr, _, token, err := DeployMiningToken(transactOpts, backend, miningToken, miningToken, miningTokenCap, miningTokenDecimals)
	require.NoError(t, err)
	require.NotNil(t, token)

	// mint all the tokens to an user
	numTokens := miningTokenCap
	_, err = token.Mint(transactOpts, getAddress(user), numTokens)
	require.NoError(t, err)

	// validatorMgr instance
	baseDeposit := numTokens
	maxNumValidators := new(big.Int).SetUint64(10)
	freezePeriod := new(big.Int).SetUint64(0)
	mgrAddr, _, mgr, err := consensus.DeployValidatorMgr(transactOpts, backend, baseDeposit, maxNumValidators, freezePeriod, tokenAddr)
	require.NoError(t, err)
	require.NotNil(t, mgr)
	require.NotZero(t, mgrAddr)

	// register validator
	transferOpts := bind.NewKeyedTransactor(user)
	_, err = token.Transfer(transferOpts, mgrAddr, numTokens, defaultData, customFallback)
	require.NoError(t, err)

	backend.Commit()

	// user must be a validator
	isValidator, err := mgr.IsValidator(&bind.CallOpts{}, getAddress(user))
	require.NoError(t, err)
	require.True(t, isValidator)
}

func TestTransfer_CustomFallback_Incompatible(t *testing.T) {
	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		getAddress(owner): core.GenesisAccount{Balance: initialBalance},
		getAddress(user):  core.GenesisAccount{Balance: initialBalance},
	})

	transactOpts := bind.NewKeyedTransactor(owner)
	_, _, token, err := DeployMiningToken(transactOpts, backend, miningToken, miningToken, miningTokenCap, miningTokenDecimals)
	require.NoError(t, err)
	require.NotNil(t, token)

	// mint all the tokens to an user
	numTokens := miningTokenCap
	_, err = token.Mint(transactOpts, getAddress(user), numTokens)
	require.NoError(t, err)

	// oracleMgr instance
	baseDeposit := numTokens
	maxNumOracles := new(big.Int).SetUint64(10)
	freezePeriod := new(big.Int).SetUint64(0)
	mgrAddr, _, mgr, err := oracle.DeployOracleMgr(transactOpts, backend, baseDeposit, maxNumOracles, freezePeriod)
	require.NoError(t, err)
	require.NotNil(t, mgr)
	require.NotZero(t, mgrAddr)

	// transfer the funds to a contract that does not support the mining token
	transferOpts := bind.NewKeyedTransactor(user)
	_, err = token.Transfer(transferOpts, mgrAddr, numTokens, defaultData, customFallback)
	require.Error(t, err, "The OracleMgr contract does not support the mining token")
}

func TestBalanceOf(t *testing.T) {
	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		getAddress(owner): core.GenesisAccount{Balance: initialBalance},
	})

	transactOpts := bind.NewKeyedTransactor(owner)
	_, _, token, err := DeployMiningToken(transactOpts, backend, miningToken, miningToken, miningTokenCap, miningTokenDecimals)
	require.NoError(t, err)
	require.NotNil(t, token)

	backend.Commit()

	userBalance, err := token.BalanceOf(&bind.CallOpts{}, getAddress(user))
	require.NoError(t, err)
	require.NotNil(t, userBalance)
	require.Zero(t, userBalance.Uint64())

	// mint all the tokens to an user
	numTokens := miningTokenCap
	_, err = token.Mint(transactOpts, getAddress(user), numTokens)
	require.NoError(t, err)

	backend.Commit()

	userBalance, err = token.BalanceOf(&bind.CallOpts{}, getAddress(user))
	require.NoError(t, err)
	require.NotNil(t, userBalance)
	require.Equal(t, numTokens, userBalance)
}

// getAddress return the address of the given private key
func getAddress(privateKey *ecdsa.PrivateKey) common.Address {
	return crypto.PubkeyToAddress(privateKey.PublicKey)
}
