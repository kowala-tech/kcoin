package token

import (
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/kowala-tech/kcoin/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/accounts/abi/bind/backends"
	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/contracts/token/testfiles"
	"github.com/kowala-tech/kcoin/core"
	"github.com/kowala-tech/kcoin/crypto"
	"github.com/kowala-tech/kcoin/params"
	"github.com/stretchr/testify/suite"
)

var (
	owner, _            = crypto.GenerateKey()
	user, _             = crypto.GenerateKey()
	initialBalance      = new(big.Int).Mul(new(big.Int).SetUint64(10), new(big.Int).SetUint64(params.Ether))         // 10 kUSD
	miningTokenCap      = new(big.Int).Mul(new(big.Int).SetUint64(1073741824), new(big.Int).SetUint64(params.Ether)) // 1073741824 mUSD
	miningToken         = "mUSD"
	miningTokenDecimals = uint8(18)
)

type MiningTokenSuite struct {
	suite.Suite
	backend     *backends.SimulatedBackend
	miningToken *MiningToken
}

func TestMiningTokenSuite(t *testing.T) {
	suite.Run(t, new(MiningTokenSuite))
}

func (suite *MiningTokenSuite) BeforeTest(suiteName, testName string) {
	req := suite.Require()

	alloc := make(core.GenesisAlloc)
	alloc[getAddress(owner)] = core.GenesisAccount{Balance: initialBalance}
	alloc[getAddress(user)] = core.GenesisAccount{Balance: initialBalance}

	switch {
	}

	backend := backends.NewSimulatedBackend(alloc)
	req.NotNil(backend)
	suite.backend = backend

	// MiningToken instance
	transactOpts := bind.NewKeyedTransactor(owner)
	_, _, mToken, err := DeployMiningToken(transactOpts, backend, miningToken, miningToken, miningTokenCap, miningTokenDecimals)
	req.NoError(err)
	req.NotNil(mToken)
	suite.miningToken = mToken

	suite.backend.Commit()
}

func (suite *MiningTokenSuite) TestDeploy() {
	req := suite.Require()

	storedCap, err := suite.miningToken.Cap(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedCap)
	req.Equal(miningTokenCap, storedCap)

	storedDecimals, err := suite.miningToken.Decimals(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedDecimals)
	req.Equal(miningTokenDecimals, storedDecimals)

	storedTotalSupply, err := suite.miningToken.TotalSupply(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedTotalSupply)
	req.Zero(storedTotalSupply.Uint64())

	storedName, err := suite.miningToken.Name(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedName)
	req.Equal(miningToken, storedName)

	storedSymbol, err := suite.miningToken.Symbol(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedSymbol)
	req.Equal(miningToken, storedSymbol)

	storedMintingFinished, err := suite.miningToken.MintingFinished(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedMintingFinished)
	req.False(storedMintingFinished)

	storedOwner, err := suite.miningToken.Owner(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(owner)
	req.Equal(getAddress(owner), storedOwner)
}

func (suite *MiningTokenSuite) TestFinishMinting() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)
	_, err := suite.miningToken.FinishMinting(transactOpts)
	req.NoError(err)

	suite.backend.Commit()

	finished, err := suite.miningToken.MintingFinished(&bind.CallOpts{})
	req.NoError(err)
	req.True(finished)
}

func (suite *MiningTokenSuite) TestFinishMinting_NotOwner() {
	req := suite.Require()

	mintingOpts := bind.NewKeyedTransactor(user)
	_, err := suite.miningToken.FinishMinting(mintingOpts)
	req.Error(err, "method FinishMinting is just available to the contract owner")
}

func (suite *MiningTokenSuite) TestFinishMinting_FinishedMinting() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)
	_, err := suite.miningToken.FinishMinting(transactOpts)
	req.NoError(err)

	_, err = suite.miningToken.FinishMinting(transactOpts)
	req.Error(err, "method FinishMinting is not available if the operation is finished")
}

func (suite *MiningTokenSuite) TestMint() {
	req := suite.Require()

	// mint all the tokens to the user
	numTokens := miningTokenCap
	transactOpts := bind.NewKeyedTransactor(owner)
	_, err := suite.miningToken.Mint(transactOpts, getAddress(user), numTokens)
	req.NoError(err)

	suite.backend.Commit()

	storedTotalSupply, err := suite.miningToken.TotalSupply(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedTotalSupply)
	req.Equal(numTokens, storedTotalSupply)

	balance, err := suite.miningToken.BalanceOf(&bind.CallOpts{}, getAddress(user))
	req.NoError(err)
	req.NotNil(balance)
	req.Equal(numTokens, balance)
}

func (suite *MiningTokenSuite) TestMint_NotOwner() {
	req := suite.Require()

	numTokens := miningTokenCap
	transactOpts := bind.NewKeyedTransactor(user)
	_, err := suite.miningToken.Mint(transactOpts, getAddress(user), numTokens)
	req.Error(err, "method Mint is just available to the contract owner")
}

func (suite *MiningTokenSuite) TestMint_OverCap() {
	req := suite.Require()

	numTokens := new(big.Int).Add(miningTokenCap, common.Big1)
	transactOpts := bind.NewKeyedTransactor(owner)
	_, err := suite.miningToken.Mint(transactOpts, getAddress(user), numTokens)
	req.Error(err, "cannot mint over the cap")
}

func (suite *MiningTokenSuite) TestTransfer_CustomFallback_Compatible() {
	req := suite.Require()

	// mint all the tokens to an user
	numTokens := miningTokenCap
	transactOpts := bind.NewKeyedTransactor(owner)
	_, err := suite.miningToken.Mint(transactOpts, getAddress(user), numTokens)
	req.NoError(err)

	// Compatible contract
	contractAddr, _, _, err := testfiles.DeployCompatible(transactOpts, suite.backend)
	req.NoError(err)
	req.NotZero(contractAddr)

	// transfer part of the funds
	value := new(big.Int).Div(numTokens, common.Big2)
	transferOpts := bind.NewKeyedTransactor(user)
	_, err = suite.miningToken.Transfer(transferOpts, contractAddr, value, defaultData, testfiles.CustomFallback)
	req.NoError(err)

	suite.backend.Commit()

	userBalance, err := suite.miningToken.BalanceOf(&bind.CallOpts{}, getAddress(user))
	req.NoError(err)
	req.NotNil(userBalance)
	req.Equal(value, userBalance)

	contractBalance, err := suite.miningToken.BalanceOf(&bind.CallOpts{}, contractAddr)
	req.NoError(err)
	req.NotNil(contractBalance)
	req.Equal(value, contractBalance)
}

func (suite *MiningTokenSuite) TestTransfer_CustomFallback_Incompatible() {
	req := suite.Require()

	// mint all the tokens to an user
	numTokens := miningTokenCap
	transactOpts := bind.NewKeyedTransactor(owner)
	_, err := suite.miningToken.Mint(transactOpts, getAddress(user), numTokens)
	req.NoError(err)

	// Incompatible contract
	contractAddr, _, _, err := testfiles.DeployIncompatible(transactOpts, suite.backend)
	req.NoError(err)
	req.NotZero(contractAddr)

	// transfer funds
	transferOpts := bind.NewKeyedTransactor(user)
	_, err = suite.miningToken.Transfer(transferOpts, contractAddr, numTokens, defaultData, testfiles.CustomFallback)
	req.Error(err, "The OracleMgr contract does not support the mining token")
}

// getAddress return the address of the given private key
func getAddress(privateKey *ecdsa.PrivateKey) common.Address {
	return crypto.PubkeyToAddress(privateKey.PublicKey)
}
