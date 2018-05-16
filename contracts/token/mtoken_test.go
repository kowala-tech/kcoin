package token

import (
	"crypto/ecdsa"
	"errors"
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
	"github.com/stretchr/testify/suite"
)

const (
	initialBalance      = 10         // kUSD (18 decimals)
	cap                 = 1073741824 // mUSD (18 decimals)
	customFallback      = "registerValidator(address,uint256,bytes)"
	miningToken         = "mUSD"
	miningTokenDecimals = 18
)

var (
	errAlwaysFailingTransaction = errors.New("failed to estimate gas needed: gas required exceeds allowance or always failing transaction")
)

type MiningTokenSuite struct {
	suite.Suite
	backend                   *backends.SimulatedBackend
	contractOwner, randomUser *ecdsa.PrivateKey
	genesisValidator          *ecdsa.PrivateKey
	initialBalance            *big.Int
	cap                       *big.Int
}

func TestMiningTokenSuite(t *testing.T) {
	suite.Run(t, new(MiningTokenSuite))
}

func (suite *MiningTokenSuite) SetupSuite() {
	req := suite.Require()

	contractOwner, err := crypto.GenerateKey()
	req.NoError(err)
	genesisValidator, err := crypto.GenerateKey()
	req.NoError(err)
	randomUser, err := crypto.GenerateKey()
	req.NoError(err)

	suite.contractOwner = contractOwner
	suite.genesisValidator = genesisValidator
	suite.randomUser = randomUser
	suite.initialBalance = musd(new(big.Int).SetUint64(initialBalance))
	suite.cap = musd(new(big.Int).SetUint64(cap))
}

func (suite *MiningTokenSuite) NewSimulatedBackend() *backends.SimulatedBackend {
	contractOwnerAddr := crypto.PubkeyToAddress(suite.contractOwner.PublicKey)
	defaultAccount := core.GenesisAccount{Balance: suite.initialBalance}
	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		contractOwnerAddr: defaultAccount,
	})

	return backend
}

func (suite *MiningTokenSuite) SetupTest() {
	suite.backend = suite.NewSimulatedBackend()
}

func (suite *MiningTokenSuite) TestFinishMinting_Success() {
	req := suite.Require()

	// create a capped mining token
	opts := bind.NewKeyedTransactor(suite.contractOwner)
	_, _, token, err := DeployMiningToken(opts, suite.backend, miningToken, miningToken, suite.cap, miningTokenDecimals)
	req.NoError(err)
	req.NotNil(token)

	// terminate the minting process
	_, err = token.FinishMinting(opts)
	req.NoError(err)

	suite.backend.Commit()

	finished, err := token.MintingFinished(&bind.CallOpts{})
	req.NoError(err)
	req.True(finished)
}

func (suite *MiningTokenSuite) TestFinishMinting_NotOwner() {
	req := suite.Require()

	// create a capped mining token
	opts := bind.NewKeyedTransactor(suite.contractOwner)
	_, _, token, err := DeployMiningToken(opts, suite.backend, miningToken, miningToken, suite.cap, miningTokenDecimals)
	req.NoError(err)
	req.NotNil(token)

	// terminate the minting process
	opts = bind.NewKeyedTransactor(suite.randomUser)
	_, err = token.FinishMinting(opts)
	req.Equal(errAlwaysFailingTransaction, err)
}

func (suite *MiningTokenSuite) TestFinishMinting_MintingOver() {
	req := suite.Require()

	// create a capped mining token
	opts := bind.NewKeyedTransactor(suite.contractOwner)
	_, _, token, err := DeployMiningToken(opts, suite.backend, miningToken, miningToken, suite.cap, miningTokenDecimals)
	req.NoError(err)
	req.NotNil(token)

	// terminate the minting process
	_, err = token.FinishMinting(opts)
	req.NoError(err)

	// terminate the minting process
	_, err = token.FinishMinting(opts)
	req.Equal(errAlwaysFailingTransaction, err)
}

func (suite *MiningTokenSuite) TestMint_Success() {
	req := suite.Require()

	// create a capped mining token
	opts := bind.NewKeyedTransactor(suite.contractOwner)
	_, _, token, err := DeployMiningToken(opts, suite.backend, miningToken, miningToken, suite.cap, miningTokenDecimals)
	req.NoError(err)
	req.NotNil(token)

	// mint all the tokens to the contract owner
	_, err = token.Mint(opts, getAddress(suite.contractOwner), suite.cap)
	req.NoError(err)
}

func (suite *MiningTokenSuite) TestMint_NotOwner() {
	req := suite.Require()

	// create a capped mining token
	opts := bind.NewKeyedTransactor(suite.contractOwner)
	_, _, token, err := DeployMiningToken(opts, suite.backend, miningToken, miningToken, suite.cap, miningTokenDecimals)
	req.NoError(err)
	req.NotNil(token)

	// mint all the tokens to the contract owner
	opts = bind.NewKeyedTransactor(suite.randomUser)
	_, err = token.Mint(opts, getAddress(suite.contractOwner), suite.cap)
	req.Equal(errAlwaysFailingTransaction, err)
}

func (suite *MiningTokenSuite) TestMint_OverCap() {
	req := suite.Require()

	// create a capped mining token
	opts := bind.NewKeyedTransactor(suite.contractOwner)
	_, _, token, err := DeployMiningToken(opts, suite.backend, miningToken, miningToken, suite.cap, miningTokenDecimals)
	req.NoError(err)
	req.NotNil(token)

	// mint all the tokens plus one to the contract owner
	_, err = token.Mint(opts, getAddress(suite.contractOwner), new(big.Int).Add(suite.cap, common.Big1))
	req.Equal(errAlwaysFailingTransaction, err)
}

func (suite *MiningTokenSuite) TestMint_MintingOver() {
	req := suite.Require()

	// create a capped mining token
	opts := bind.NewKeyedTransactor(suite.contractOwner)
	_, _, token, err := DeployMiningToken(opts, suite.backend, miningToken, miningToken, suite.cap, miningTokenDecimals)
	req.NoError(err)
	req.NotNil(token)

	// terminate the minting process
	_, err = token.FinishMinting(opts)
	req.NoError(err)

	// mint all the tokens to the contract owner
	_, err = token.Mint(opts, getAddress(suite.contractOwner), suite.cap)
	req.Equal(errAlwaysFailingTransaction, err)
}

func (suite *MiningTokenSuite) TestTransfer_CustomFallback_CompatibleContract_ValidatorMgr_RegisterValidator() {
	req := suite.Require()

	// create a capped mining token
	opts := bind.NewKeyedTransactor(suite.contractOwner)
	_, _, token, err := DeployMiningToken(opts, suite.backend, miningToken, miningToken, suite.cap, miningTokenDecimals)
	req.NoError(err)
	req.NotNil(token)

	// mint all the tokens to the contract owner
	_, err = token.Mint(opts, getAddress(suite.contractOwner), suite.cap)
	req.NoError(err)

	// the validator manager contract implements the custom fallback
	validatorMgrAddr, _, validatorMgr, err := consensus.DeployValidatorMgr(opts, suite.backend, common.Big0, common.Big32, common.Big0, getAddress(suite.genesisValidator))
	req.NoError(err)

	// transfer mUSD to the validator manager contract
	_, err = token.Transfer(opts, validatorMgrAddr, common.Big0, []byte("non-zero"), customFallback)
	req.NoError(err)

	suite.backend.Commit()

	// make sure that the new validator has been registered
	count, err := validatorMgr.GetValidatorCount(&bind.CallOpts{})
	req.NoError(err)
	req.Equal(common.Big2, count)
}

func (suite *MiningTokenSuite) TestTransfer_CustomFallback_IncompatibleContract_OracleMgr() {
	req := suite.Require()

	opts := bind.NewKeyedTransactor(suite.contractOwner)
	_, _, token, err := DeployMiningToken(opts, suite.backend, "mUSD", "mUSD", suite.cap, 18)
	req.NoError(err)
	req.NotNil(token)

	// mint all the tokens to the contract owner
	_, err = token.Mint(opts, getAddress(suite.contractOwner), suite.cap)
	req.NoError(err)

	// the oracle manager contract doesn't implement the custom fallback - receiveToken
	oracleMgrAddr, _, _, err := oracle.DeployOracleMgr(opts, suite.backend, common.Big0, common.Big32, common.Big0)
	req.NoError(err)

	suite.backend.Commit()

	// transfer mUSD to the oracle manager contract
	// transaction must fail because the oracle mgr does not implement the custom fallback
	_, err = token.Transfer(opts, oracleMgrAddr, common.Big0, []byte("non-zero"), customFallback)
	req.Equal(errAlwaysFailingTransaction, err)
}

/*
func (suite *MiningTokenSuite) TestBalanceOf() {
	req := suite.Require()

	// create a capped mining token
	opts := bind.NewKeyedTransactor(suite.contractOwner)
	_, _, token, err := DeployMiningToken(opts, suite.backend, miningToken, miningToken, suite.cap, miningTokenDecimals)
	req.NoError(err)
	req.NotNil(token)

	// mint all the tokens to the contract owner
	_, err = token.Mint(opts, getAddress(suite.contractOwner), suite.cap)
	req.NoError(err)

	// transfer all mUSD from the owner to a random user
	_, err = token.Transfer(opts, getAddress(suite.randomUser), suite.cap, []byte("non-zero"), customFallback)
	req.NoError(err)

	suite.backend.Commit()

	randomUserBalance, err := token.BalanceOf(&bind.CallOpts{}, getAddress(suite.randomUser))
	req.NoError(err)
	req.Equal(suite.cap, randomUserBalance)

	ownerBalance, err := token.BalanceOf(&bind.CallOpts{}, getAddress(suite.contractOwner))
	req.NoError(err)
	req.Equal(common.Big0, ownerBalance)
}
*/

// getAddress return the address of the given private key
func getAddress(privateKey *ecdsa.PrivateKey) common.Address {
	return crypto.PubkeyToAddress(privateKey.PublicKey)
}

// musd converts the value to mUSD
func musd(value *big.Int) *big.Int {
	return new(big.Int).Mul(value, new(big.Int).SetUint64(params.Ether))
}
