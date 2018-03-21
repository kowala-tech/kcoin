package contracts

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"testing"

	"github.com/kowala-tech/kUSD/accounts/abi/bind"
	"github.com/kowala-tech/kUSD/accounts/abi/bind/backends"
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
	addr, _, contract, err := DeployElectionContract(opts, suite.backend, baseDeposit, maxValidators, unbondingPeriod, crypto.PubkeyToAddress(suite.genesisValidator.PublicKey))
	if err != nil {
		return err
	}
	suite.contract = contract

	// NOTE (rgeraldes) - add balance to cover the base deposit
	// for the genesis validator. Eventually this could change
	// as soon as the token contracts are completed.
	//suite.backend.SendTransaction(context.TODO(), types.NewTransaction(uint64(0), addr, suite.baseDeposit, nil, nil, []byte("")))
	suite.backend.Commit()

	balance, err := suite.backend.BalanceAt(context.TODO(), addr, suite.backend.CurrentBlock().Number())
	if err != nil {
		return err
	}
	fmt.Println(balance)

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
		latestUnbondingPeriod, err := suite.contract.UnbondingPeriod(&bind.CallOpts{})
		req.NoError(err)
		req.Equal(new(big.Int).SetUint64(suite.unbondingPeriod.Int64()*time.Hour.*24), latestUnbondingPeriod)
	*/

	genesisValidator, err := suite.contract.GenesisValidator(&bind.CallOpts{})
	req.NoError(err)
	req.Equal(crypto.PubkeyToAddress(suite.genesisValidator.PublicKey), genesisValidator)
}
