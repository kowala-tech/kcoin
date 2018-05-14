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
	"github.com/stretchr/testify/suite"
)

const (
	initialBalance = 10         // kUSD (18 decimals)
	cap            = 1073741824 // mUSD (18 decimals)
)

type MiningTokenSuite struct {
	suite.Suite
	backend          *backends.SimulatedBackend
	contract         *MiningToken
	contractOwner    *ecdsa.PrivateKey
	genesisValidator *ecdsa.PrivateKey
	initialBalance   *big.Int
	cap              *big.Int
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

	suite.contractOwner = contractOwner
	suite.genesisValidator = genesisValidator
	suite.initialBalance = new(big.Int).Mul(new(big.Int).SetUint64(initialBalance), new(big.Int).SetUint64(params.Ether))
	suite.cap = new(big.Int).Mul(new(big.Int).SetUint64(cap), new(big.Int).SetUint64(params.Ether))
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

func (suite *MiningTokenSuite) Transfer_CustomFallback_CompatibleContract_ValidatorMgr() {
	req := suite.Require()

	// create a capped mining token
	opts := bind.NewKeyedTransactor(suite.contractOwner)
	_, _, token, err := DeployMiningToken(opts, suite.backend, "mUSD", "mUSD", suite.cap, 18)
	req.NoError(err)

	// mint all the tokens to the contract owner
	token.Mint(opts, getAddress(suite.contractOwner), suite.cap)

	// the validator manager contract implements the custom fallback - receiveToken
	validatorMgrAddr, _, _, err := consensus.DeployValidatorManager(opts, suite.backend, common.Big0, common.Big32, common.Big0, getAddress(suite.genesisValidator))
	req.NoError(err)

	// transfer mUSD to the validator manager contract
	_, err = token.Transfer(opts, validatorMgrAddr, new(big.Int).SetUint64(100000), []byte(""), "receiveToken")
	req.NoError(err)

	suite.backend.Commit()

	// make sure that a validator has been registered

}

func (suite *MiningTokenSuite) Transfer_CustomFallback_CompatibleContract_OracleMgr() {
	req := suite.Require()

	opts := bind.NewKeyedTransactor(suite.contractOwner)
	_, _, token, err := DeployMiningToken(opts, suite.backend, "mUSD", "mUSD", suite.cap, 18)
	req.NoError(err)

	// mint all the tokens to the contract owner
	token.Mint(opts, getAddress(suite.contractOwner), suite.cap)

	// the oracle manager contract doesn't implement the custom fallback - receiveToken
	oracleMgrAddr, _, _, err := oracle.DeployOracleManager(opts, suite.backend, common.Big0, common.Big32, common.Big0)
	req.NoError(err)

	// transfer mUSD to the oracle manager contract
	// transaction must fail because the oracle manager does not implement the custom fallback - receiveToken
	_, err = token.Transfer(opts, oracleMgrAddr, new(big.Int).SetUint64(100000), []byte(""), "receiveToken")
	req.NoError(err)

	suite.backend.Commit()

	// make sure that a validator has been registered

}

func getAddress(privateKey *ecdsa.PrivateKey) common.Address {
	return crypto.PubkeyToAddress(privateKey.PublicKey)
}
