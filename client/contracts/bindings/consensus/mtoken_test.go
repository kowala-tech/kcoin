package consensus_test

import (
	"math/big"
	"strings"
	"testing"

	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/accounts/abi/bind/backends"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/consensus"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/consensus/testfiles"
	"github.com/kowala-tech/kcoin/client/core"
	"github.com/kowala-tech/kcoin/client/crypto"
	"github.com/kowala-tech/kcoin/client/params"
	"github.com/stretchr/testify/suite"
)

var (
	owner, _            = crypto.GenerateKey()
	initialBalance      = new(big.Int).Mul(common.Big32, new(big.Int).SetUint64(params.Kcoin))                       // 10 Kcoin
	miningTokenCap      = new(big.Int).Mul(new(big.Int).SetUint64(1073741824), new(big.Int).SetUint64(params.Kcoin)) // 1073741824 Kcoin
	miningTokenName     = "mUSD"
	miningTokenDecimals = uint8(18)
)

type MiningTokenSuite struct {
	suite.Suite
	backend     *backends.SimulatedBackend
	miningToken *consensus.MiningToken
}

func TestMiningTokenSuite(t *testing.T) {
	suite.Run(t, new(MiningTokenSuite))
}

func (suite *MiningTokenSuite) BeforeTest(suiteName, testName string) {
	if strings.Contains(testName, "TestDeploy") {
		return
	}

	req := suite.Require()

	alloc := make(core.GenesisAlloc)
	alloc[getAddress(owner)] = core.GenesisAccount{Balance: initialBalance}
	alloc[getAddress(user)] = core.GenesisAccount{Balance: initialBalance}

	backend := backends.NewSimulatedBackend(alloc)
	req.NotNil(backend)
	suite.backend = backend

	// MiningToken instance
	transactOpts := bind.NewKeyedTransactor(owner)
	_, _, mToken, err := consensus.DeployMiningToken(transactOpts, backend, miningTokenName, miningTokenName, miningTokenCap, miningTokenDecimals)
	req.NoError(err)
	req.NotNil(mToken)
	suite.miningToken = mToken

	suite.backend.Commit()
}

func (suite *MiningTokenSuite) TestDeploy() {
	req := suite.Require()

	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		getAddress(owner): core.GenesisAccount{
			Balance: new(big.Int).Mul(new(big.Int).SetUint64(100), new(big.Int).SetUint64(params.Kcoin)),
		},
	})

	// MiningToken instance
	transactOpts := bind.NewKeyedTransactor(owner)
	_, _, miningToken, err := consensus.DeployMiningToken(transactOpts, backend, miningTokenName, miningTokenName, miningTokenCap, miningTokenDecimals)
	req.NoError(err)
	req.NotNil(miningToken)

	backend.Commit()

	storedCap, err := miningToken.Cap(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedCap)
	req.Equal(miningTokenCap, storedCap)

	storedDecimals, err := miningToken.Decimals(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedDecimals)
	req.Equal(miningTokenDecimals, storedDecimals)

	storedTotalSupply, err := miningToken.TotalSupply(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedTotalSupply)
	req.Zero(storedTotalSupply.Uint64())

	storedName, err := miningToken.Name(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedName)
	req.Equal(miningTokenName, storedName)

	storedSymbol, err := miningToken.Symbol(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedSymbol)
	req.Equal(miningTokenName, storedSymbol)

	storedMintingFinished, err := miningToken.MintingFinished(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedMintingFinished)
	req.False(storedMintingFinished)

	storedOwner, err := miningToken.Owner(&bind.CallOpts{})
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
	req.Error(err, "method Minter is just available to the contract owner")
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
	_, err = suite.miningToken.Transfer(transferOpts, contractAddr, value, consensus.DefaultData, testfiles.CustomFallback)
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
	_, err = suite.miningToken.Transfer(transferOpts, contractAddr, numTokens, consensus.DefaultData, testfiles.CustomFallback)
	req.Error(err, "The OracleMgr contract does not support the mining token")
}
