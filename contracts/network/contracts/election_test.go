package contracts

import (
	"crypto/ecdsa"
	"errors"
	"math/big"
	"testing"

	"github.com/kowala-tech/kUSD/accounts/abi/bind"
	"github.com/kowala-tech/kUSD/accounts/abi/bind/backends"
	"github.com/kowala-tech/kUSD/common"
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
	_, _, contract, err := DeployElectionContract(opts, suite.backend, baseDeposit, maxValidators, unbondingPeriod, crypto.PubkeyToAddress(suite.genesisValidator.PublicKey))
	if err != nil {
		return err
	}
	suite.contract = contract

	// NOTE (rgeraldes) - add balance to cover the base deposit
	// for the genesis validator. Eventually this could change
	// as soon as the token contracts are completed.
	opts.Value = suite.baseDeposit
	_, err = contract.ElectionContractTransactor.contract.Transfer(opts)
	if err != nil {
		return err
	}
	suite.backend.Commit()

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

func (suite *ElectionContractSuite) TestDeployElectionContract_MaxValidatorsEqualZero() {
	req := suite.Require()

	maxValidators := common.Big0
	req.Equal(errTransactionFailed, suite.DeployElectionContract(suite.baseDeposit, maxValidators, suite.unbondingPeriod))
}

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

func (suite *ElectionContractSuite) TestIsGenesis() {
	req := suite.Require()

	testCases := []struct {
		input  common.Address
		output bool
	}{
		{
			input:  getAddress(suite.genesisValidator),
			output: true,
		},
		{
			input:  getAddress(suite.randomUser),
			output: false,
		},
	}

	for _, tc := range testCases {
		isGenesis, err := suite.contract.IsGenesisValidator(&bind.CallOpts{}, tc.input)
		req.NoError(err)
		req.Equal(tc.output, isGenesis)
	}
}

func (suite *ElectionContractSuite) TestIsValidator() {
	req := suite.Require()

	testCases := []struct {
		input  common.Address
		output bool
	}{
		{
			input:  getAddress(suite.genesisValidator),
			output: true,
		},
		{
			input:  getAddress(suite.randomUser),
			output: false,
		},
	}

	for _, tc := range testCases {
		isValidator, err := suite.contract.IsValidator(&bind.CallOpts{}, tc.input)
		req.NoError(err)
		req.Equal(tc.output, isValidator)
	}
}

func (suite *ElectionContractSuite) TestGetMinimumDeposit_ElectionFull() {
	req := suite.Require()

	// leave a position available for the genesis validator - max validators = 1
	maxValidators := new(big.Int).SetUint64(1)
	suite.DeployElectionContract(suite.baseDeposit, maxValidators, suite.unbondingPeriod)

	minDeposit, err := suite.contract.GetMinimumDeposit(&bind.CallOpts{})
	req.NoError(err)
	// min deposit should be greater (+ 1) than the smallest stake
	// at play which is equal to the base deposit (genesis validator)
	req.Equal((new(big.Int).Add(suite.baseDeposit, common.Big1)), minDeposit)
}

func (suite *ElectionContractSuite) TestGetMinimumDeposit_ElectionNotFull() {
	// by default the contract has one validator (genesis) and 99 (100 - 1)
	// positions available
	req := suite.Require()

	minDeposit, err := suite.contract.GetMinimumDeposit(&bind.CallOpts{})
	req.NoError(err)
	// min deposit should be equal to the base deposit since
	// there are positions available
	req.Equal(suite.baseDeposit, minDeposit)
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

func getAddress(privateKey *ecdsa.PrivateKey) common.Address {
	return crypto.PubkeyToAddress(privateKey.PublicKey)
}
