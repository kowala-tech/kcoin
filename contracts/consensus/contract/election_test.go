package contract

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
	baseDeposit   = 100000
	maxValidators = 100
	unbondingPeriod
)

var (
	errTransactionFailed = errors.New("failed to estimate gas needed: gas required exceeds allowance or always failing transaction")
)

type ElectionContractSuite struct {
	suite.Suite
	backend                  *backends.SimulatedBackend
	contract                 *ElectionContract
	owner, notOwner, genesis *ecdsa.PrivateKey
	baseDeposit              *big.Int
	maxValidators            *big.Int
	unbondingPeriod          *big.Int
}

func TestElectionContractSuite(t *testing.T) {
	suite.Run(t, new(ElectionContractSuite))
}

func (s *ElectionContractSuite) SetupSuite() {
	r := s.Require()

	owner, err := crypto.GenerateKey()
	r.NoError(err)
	notOwner, err := crypto.GenerateKey()
	r.NoError(err)
	genesis, err := crypto.GenerateKey()
	r.NoError(err)
	s.owner, s.notOwner, s.genesis = owner, notOwner, genesis
	s.baseDeposit = big.NewInt(baseDeposit)
	s.maxValidators = big.NewInt(maxValidators)
	s.unbondingPeriod = big.NewInt(unbondingPeriod)
}

func (s *ElectionContractSuite) NewSimulatedBackend() *backends.SimulatedBackend {
	// @NOTE (rgeraldes) - fund the owner account with enough balance to
	// deploy the contract and to cover the genesis validator deposit
	// The current value is more than enough
	ownerAddr := crypto.PubkeyToAddress(s.owner.PublicKey)
	notOwnerAddr := crypto.PubkeyToAddress(s.notOwner.PublicKey)
	simulatedBackend := backends.NewSimulatedBackend(core.GenesisAlloc{
		ownerAddr:    core.GenesisAccount{Balance: new(big.Int).Mul(big.NewInt(baseDeposit), new(big.Int).SetUint64(params.Ether))},
		notOwnerAddr: core.GenesisAccount{Balance: new(big.Int).Mul(big.NewInt(baseDeposit), new(big.Int).SetUint64(params.Ether))},
	})
	return simulatedBackend
}

func (s *ElectionContractSuite) DeployElectionContract(deposit, maxValidators, unbondingPeriod *big.Int) error {
	opts := bind.NewKeyedTransactor(s.owner)
	opts.Value = deposit
	_, _, contract, err := DeployElectionContract(opts, s.backend, s.baseDeposit, maxValidators, unbondingPeriod, crypto.PubkeyToAddress(s.genesis.PublicKey))
	if err != nil {
		return err
	}
	s.contract = contract

	s.backend.Commit()

	return nil
}

func (s *ElectionContractSuite) SetupTest() {
	require := s.Require()

	s.backend = s.NewSimulatedBackend()
	require.NoError(s.DeployElectionContract(s.baseDeposit, s.maxValidators, s.unbondingPeriod))
}

func (s *ElectionContractSuite) TestDeploy() {
	r := s.Require()

	err := s.DeployElectionContract(s.baseDeposit, s.maxValidators, s.unbondingPeriod)
	r.NoError(err)

	s.backend.Commit()

	latestBaseDeposit, err := s.contract.BaseDeposit(&bind.CallOpts{})
	r.NoError(err)
	r.Equal(s.baseDeposit, latestBaseDeposit)

	latestMax, err := s.contract.MaxValidators(&bind.CallOpts{})
	r.NoError(err)
	r.Equal(s.maxValidators, latestMax)

	latestPeriod, err := s.contract.UnbondingPeriod(&bind.CallOpts{})
	r.NoError(err)
	r.Equal(s.unbondingPeriod, latestPeriod)

	genesis, err := s.contract.Genesis(&bind.CallOpts{})
	r.NoError(err)
	r.Equal(crypto.PubkeyToAddress(s.genesis.PublicKey), genesis)
}

func (s *ElectionContractSuite) TestDeployDepositIsNotEnough() {
	r := s.Require()

	err := s.DeployElectionContract(big.NewInt(0).Sub(s.baseDeposit, big.NewInt(1)), s.maxValidators, s.unbondingPeriod)
	r.Equal(errTransactionFailed, err)
}

func (s *ElectionContractSuite) TestDeployMaxValidatorsEqualZero() {
	r := s.Require()

	err := s.DeployElectionContract(s.baseDeposit, big.NewInt(0), s.unbondingPeriod)
	r.Equal(errTransactionFailed, err)
}

func (s *ElectionContractSuite) TestDeployMaxValidatorsGreaterThanZero() {
	r := s.Require()

	max := big.NewInt(1)
	err := s.DeployElectionContract(s.baseDeposit, max, s.unbondingPeriod)
	r.NoError(err)

	s.backend.Commit()

	latestMax, err := s.contract.MaxValidators(&bind.CallOpts{})
	r.NoError(err)
	r.Equal(max, latestMax)
}

func (s *ElectionContractSuite) TestGetOwner() {
	r := s.Require()

	latestOwner, err := s.contract.GetOwner(&bind.CallOpts{})
	r.NoError(err)
	owner := crypto.PubkeyToAddress(s.owner.PublicKey)
	r.Equal(owner, latestOwner)
}

func (s *ElectionContractSuite) TestTransferOwnershipNotOwner() {
	r := s.Require()

	newOwnerPK, err := crypto.GenerateKey()
	r.NoError(err)
	newOwner := crypto.PubkeyToAddress(newOwnerPK.PublicKey)
	_, err = s.contract.TransferOwnership(bind.NewKeyedTransactor(s.notOwner), newOwner)
	r.Equal(errTransactionFailed, err)
}

func (s *ElectionContractSuite) TestTransferOwnershipOwner() {
	r := s.Require()

	newOwnerPK, err := crypto.GenerateKey()
	r.NoError(err)
	newOwner := crypto.PubkeyToAddress(newOwnerPK.PublicKey)
	_, err = s.contract.TransferOwnership(bind.NewKeyedTransactor(s.owner), newOwner)
	r.NoError(err)

	s.backend.Commit()

	latestOwner, err := s.contract.GetOwner(&bind.CallOpts{})
	r.NoError(err)
	r.Equal(newOwner, latestOwner)
}

func (s *ElectionContractSuite) TestSetBaseDepositNotOwner() {
	r := s.Require()

	_, err := s.contract.SetBaseDeposit(bind.NewKeyedTransactor(s.notOwner), big.NewInt(0))
	r.Equal(errTransactionFailed, err)
}

func (s *ElectionContractSuite) TestSetBaseDepositOwner() {
	r := s.Require()

	deposit := big.NewInt(s.baseDeposit.Int64() + 1)
	_, err := s.contract.SetBaseDeposit(bind.NewKeyedTransactor(s.owner), deposit)
	r.NoError(err)

	s.backend.Commit()

	latestDeposit, err := s.contract.BaseDeposit(&bind.CallOpts{})
	r.NoError(err)
	r.Equal(deposit, latestDeposit)
}

func (s *ElectionContractSuite) TestSetMaxValidatorsNotOwner() {
	r := s.Require()

	_, err := s.contract.SetMaxValidators(bind.NewKeyedTransactor(s.notOwner), big.NewInt(0))
	r.Equal(errTransactionFailed, err)
}

/*
func (s *ElectionContractSuite) TestSetMaxValidatorsLessThanValidatorCount() {
	r := s.Require()

	oldMax, err := s.contract.MaxValidators(&bind.CallOpts{})
	r.NoError(err)

	oldCount, err := s.contract.GetValidatorCount(&bind.CallOpts{})
	r.NoError(err)

	// make sure that we remove at least one validator
	max := big.NewInt(oldCount.Int64() - 1)
	r.NotEqual(oldMax, max)
	_, err = s.contract.SetMaxValidators(bind.NewKeyedTransactor(s.owner), max)
	r.NoError(err)

	s.backend.Commit()

	latestMax, err := s.contract.MaxValidators(&bind.CallOpts{})
	r.NoError(err)
	r.Equal(max, latestMax)

	// the validator set should be empty
	latestCount, err := s.contract.GetValidatorCount(&bind.CallOpts{})
	r.NoError(err)
	r.Equal(max, latestCount)
}

func (s *ElectionContractSuite) TestSetMaxValidatorsGreaterOrEqualThanValidatorCount() {
	r := s.Require()

	oldMax, err := s.contract.MaxValidators(&bind.CallOpts{})
	r.NoError(err)

	oldCount, err := s.contract.GetValidatorCount(&bind.CallOpts{})
	r.NoError(err)

	max := big.NewInt(oldCount.Int64() + 1)
	r.NotEqual(oldMax, max)
	_, err = s.contract.SetMaxValidators(bind.NewKeyedTransactor(s.owner), max)
	r.NoError(err)

	s.backend.Commit()

	latestMax, err := s.contract.MaxValidators(&bind.CallOpts{})
	r.NoError(err)
	r.Equal(max, latestMax)

	latestCount, err := s.contract.GetValidatorCount(&bind.CallOpts{})
	r.NoError(err)
	r.Equal(oldCount, latestCount)
}
*/

func (s *ElectionContractSuite) TestIsGenesis() {
	r := s.Require()

	testCases := []struct {
		input  common.Address
		output bool
	}{
		{
			crypto.PubkeyToAddress(s.genesis.PublicKey),
			true,
		},
		{
			crypto.PubkeyToAddress(s.notOwner.PublicKey),
			false,
		},
	}

	for _, tc := range testCases {
		isGenesis, err := s.contract.IsGenesisValidator(&bind.CallOpts{}, tc.input)
		r.NoError(err)
		r.Equal(tc.output, isGenesis)
	}
}

func (s *ElectionContractSuite) TestGetMinimumDeposit_ElectionNotFull() {
	r := s.Require()

	minDeposit, err := s.contract.GetMinimumDeposit(&bind.CallOpts{})
	r.NoError(err)
	r.Equal(s.baseDeposit, minDeposit)
}

func (s *ElectionContractSuite) TestGetMinimumDeposit_ElectionFull() {
	r := s.Require()

	// leave margin just for the genesis validator
	s.DeployElectionContract(s.baseDeposit, big.NewInt(1), s.unbondingPeriod)

	minDeposit, err := s.contract.GetMinimumDeposit(&bind.CallOpts{})
	r.NoError(err)
	r.NotEqual(s.baseDeposit, minDeposit)
}

func (s *ElectionContractSuite) TestJoinElection_NotEnoughDeposit() {
	r := s.Require()

	minDeposit, err := s.contract.GetMinimumDeposit(&bind.CallOpts{})
	r.NoError(err)

	opts := bind.NewKeyedTransactor(s.notOwner)
	opts.Value = big.NewInt(minDeposit.Int64() - 1)
	_, err = s.contract.Join(opts)
	r.Equal(errTransactionFailed, err)
}

func (s *ElectionContractSuite) TestJoinElection_ElectionNotFull() {
	r := s.Require()

	minDeposit, err := s.contract.GetMinimumDeposit(&bind.CallOpts{})
	r.NoError(err)

	oldCount, err := s.contract.GetValidatorCount(&bind.CallOpts{})
	r.NoError(err)

	opts := bind.NewKeyedTransactor(s.notOwner)
	opts.Value = minDeposit
	_, err = s.contract.Join(opts)
	r.NoError(err)

	s.backend.Commit()

	// there's one more validator
	latestCount, err := s.contract.GetValidatorCount(&bind.CallOpts{})
	r.NoError(err)
	r.Equal(latestCount, big.NewInt(oldCount.Int64()+1))

}

func (s *ElectionContractSuite) TestJoinElection_ElectionFull() {
	r := s.Require()

	// leave margin just for the genesis validator
	s.DeployElectionContract(s.baseDeposit, big.NewInt(1), s.unbondingPeriod)

	minDeposit, err := s.contract.GetMinimumDeposit(&bind.CallOpts{})
	r.NoError(err)

	oldCount, err := s.contract.GetValidatorCount(&bind.CallOpts{})
	r.NoError(err)

	opts := bind.NewKeyedTransactor(s.notOwner)
	opts.Value = minDeposit
	_, err = s.contract.Join(opts)
	r.NoError(err)

	s.backend.Commit()

	// number of validators should be the same
	latestCount, err := s.contract.GetValidatorCount(&bind.CallOpts{})
	r.NoError(err)
	r.Equal(latestCount, oldCount)

	// minimum deposit should increase by 1
	latestMinDeposit, err := s.contract.GetMinimumDeposit(&bind.CallOpts{})
	r.NoError(err)
	r.NotEqual(minDeposit, latestMinDeposit)
}

func (s *ElectionContractSuite) TestLeave_NotValidator() {
	r := s.Require()

	opts := bind.NewKeyedTransactor(s.notOwner)
	_, err := s.contract.Leave(opts)
	r.Equal(errTransactionFailed, err)
}

func (s *ElectionContractSuite) TestLeave_Validator() {
	r := s.Require()

	opts := bind.NewKeyedTransactor(s.genesis)
	_, err := s.contract.Leave(opts)
	r.NoError(err)

	s.backend.Commit()

	isValidator, err := s.contract.IsValidator(&bind.CallOpts{}, crypto.PubkeyToAddress(s.notOwner.PublicKey))
	r.NoError(err)
	r.False(isValidator)
}

func (s *ElectionContractSuite) TestRedeemFunds_NoDeposits() {
	r := s.Require()

	opts := bind.NewKeyedTransactor(s.notOwner)
	_, err := s.contract.RedeemFunds(opts)

	// check balance before commiting transaction

	s.backend.Commit()
}

func (s *ElectionContractSuite) TestRedeemFunds_NoDepositsPastUnbondingPeriod() {
	r := s.Require()

	// genesis has not left yet the election
	opts := bind.NewKeyedTransactor(s.genesis)

}

func (s *ElectionContractSuite) TestRedeemFunds_HasDepositsPastUnbondingPeriod() {}
