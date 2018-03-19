package contract

import (
	"context"
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
	baseDeposit     = 100000
	maxValidators   = 100
	unbondingPeriod = 10000
)

var (
	errTransactionFailed = errors.New("failed to estimate gas needed: gas required exceeds allowance or always failing transaction")
)

type ElectionContractSuite struct {
	suite.Suite
	backend                  *backends.SimulatedBackend
	contract                 *ElectionContract
	owner, notOwner, genesis *ecdsa.PrivateKey
	balance                  *big.Int
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
	s.balance = new(big.Int).Mul(s.baseDeposit, new(big.Int).SetUint64(params.Ether))
}

func (s *ElectionContractSuite) NewSimulatedBackend() *backends.SimulatedBackend {
	// @NOTE (rgeraldes) - fund the owner account with enough balance to
	// deploy the contract and to cover the genesis validator deposit
	// The current value is more than enough.
	ownerAddr := crypto.PubkeyToAddress(s.owner.PublicKey)
	notOwnerAddr := crypto.PubkeyToAddress(s.notOwner.PublicKey)
	genesisAddr := crypto.PubkeyToAddress(s.genesis.PublicKey)
	simulatedBackend := backends.NewSimulatedBackend(core.GenesisAlloc{
		ownerAddr:    core.GenesisAccount{Balance: s.balance},
		notOwnerAddr: core.GenesisAccount{Balance: s.balance},
		genesisAddr:  core.GenesisAccount{Balance: s.balance},
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

func (s *ElectionContractSuite) TestIsValidator() {
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
		isValidator, err := s.contract.IsValidator(&bind.CallOpts{}, tc.input)
		r.NoError(err)
		r.Equal(tc.output, isValidator)
	}
}

func (s *ElectionContractSuite) TestGetMinimumDeposit_ElectionNotFull() {
	r := s.Require()

	minDeposit, err := s.contract.GetMinimumDeposit(&bind.CallOpts{})
	r.NoError(err)
	// min deposit should be equal to the base deposit
	r.Equal(s.baseDeposit, minDeposit)
}

func (s *ElectionContractSuite) TestGetMinimumDeposit_ElectionFull() {
	r := s.Require()

	// leave margin just for the genesis validator
	s.DeployElectionContract(s.baseDeposit, big.NewInt(1), s.unbondingPeriod)

	minDeposit, err := s.contract.GetMinimumDeposit(&bind.CallOpts{})
	r.NoError(err)
	// min deposit should be bigger than the base deposit
	r.NotEqual(s.baseDeposit, minDeposit)
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
	// @NOTE (rgeraldes) - Int64 is necessary for cases with value 0 like
	// this one.
	r.Equal(max.Int64(), latestMax.Int64())

	// the validator set should be empty
	latestCount, err := s.contract.GetValidatorCount(&bind.CallOpts{})
	r.NoError(err)
	r.Equal(max.Int64(), latestCount.Int64())
}

func (s *ElectionContractSuite) TestJoin_JoinMultipleTimes() {
	r := s.Require()

	opts := bind.NewKeyedTransactor(s.genesis)
	_, err := s.contract.Join(opts)
	r.Equal(errTransactionFailed, err)
}

func (s *ElectionContractSuite) TestJoin_NotEnoughDeposit() {
	r := s.Require()

	minDeposit, err := s.contract.GetMinimumDeposit(&bind.CallOpts{})
	r.NoError(err)

	opts := bind.NewKeyedTransactor(s.notOwner)
	opts.Value = big.NewInt(minDeposit.Int64() - 1)
	_, err = s.contract.Join(opts)
	r.Equal(errTransactionFailed, err)
}

func (s *ElectionContractSuite) TestJoin_OrderedInsert_GreaterThan() {
	r := s.Require()
	sender := s.notOwner

	genesis, err := s.contract.GetValidatorAtIndex(&bind.CallOpts{}, big.NewInt(0))
	r.NoError(err)

	oldCount, err := s.contract.GetValidatorCount(&bind.CallOpts{})
	r.NoError(err)

	opts := bind.NewKeyedTransactor(sender)
	// make sure that the deposit is bigger than the genesis deposit
	opts.Value = big.NewInt(genesis.Deposit.Int64() + 1)
	_, err = s.contract.Join(opts)
	r.NoError(err)

	s.backend.Commit()

	// we should have one more validator
	latestCount, err := s.contract.GetValidatorCount(&bind.CallOpts{})
	r.NoError(err)
	r.Equal(big.NewInt(oldCount.Int64()+1), latestCount)

	// validator at index 0 should be the new candidate
	validator, err := s.contract.GetValidatorAtIndex(&bind.CallOpts{}, big.NewInt(0))
	r.NoError(err)
	r.Equal(validator.Code, crypto.PubkeyToAddress(sender.PublicKey))
	r.Equal(opts.Value, validator.Deposit)
}

func (s *ElectionContractSuite) TestJoin_OrderedInsert_LessOrEqualThan() {
	r := s.Require()
	sender := s.notOwner

	genesis, err := s.contract.GetValidatorAtIndex(&bind.CallOpts{}, big.NewInt(0))
	r.NoError(err)

	oldCount, err := s.contract.GetValidatorCount(&bind.CallOpts{})
	r.NoError(err)

	opts := bind.NewKeyedTransactor(sender)
	// make sure that the deposit is bigger than the genesis deposit
	opts.Value = big.NewInt(genesis.Deposit.Int64())
	_, err = s.contract.Join(opts)
	r.NoError(err)

	s.backend.Commit()

	// we should have one more validator
	latestCount, err := s.contract.GetValidatorCount(&bind.CallOpts{})
	r.NoError(err)
	r.Equal(big.NewInt(oldCount.Int64()+1), latestCount)

	// validator at index 0 should be the new candidate
	validator, err := s.contract.GetValidatorAtIndex(&bind.CallOpts{}, big.NewInt(0))
	r.NoError(err)
	r.Equal(validator.Code, crypto.PubkeyToAddress(s.genesis.PublicKey))
	r.Equal(s.baseDeposit, validator.Deposit)
}

func (s *ElectionContractSuite) TestJoin_ElectionNotFull() {
	r := s.Require()
	sender := s.notOwner

	minDeposit, err := s.contract.GetMinimumDeposit(&bind.CallOpts{})
	r.NoError(err)

	oldCount, err := s.contract.GetValidatorCount(&bind.CallOpts{})
	r.NoError(err)

	opts := bind.NewKeyedTransactor(sender)
	opts.Value = minDeposit
	_, err = s.contract.Join(opts)
	r.NoError(err)

	s.backend.Commit()

	// we should have one more validator
	latestCount, err := s.contract.GetValidatorCount(&bind.CallOpts{})
	r.NoError(err)
	r.Equal(big.NewInt(oldCount.Int64()+1), latestCount)
}

func (s *ElectionContractSuite) TestJoin_ElectionFull() {
	r := s.Require()
	sender := s.notOwner

	// leave margin just for the genesis validator
	s.DeployElectionContract(s.baseDeposit, big.NewInt(1), s.unbondingPeriod)

	minDeposit, err := s.contract.GetMinimumDeposit(&bind.CallOpts{})
	r.NoError(err)

	oldCount, err := s.contract.GetValidatorCount(&bind.CallOpts{})
	r.NoError(err)

	opts := bind.NewKeyedTransactor(sender)
	opts.Value = minDeposit
	_, err = s.contract.Join(opts)
	r.NoError(err)

	s.backend.Commit()

	// number of validators should be the same (replacement)
	latestCount, err := s.contract.GetValidatorCount(&bind.CallOpts{})
	r.NoError(err)
	r.Equal(latestCount, oldCount)

	// validator at index (validator count - 1) should be the new candidate
	validator, err := s.contract.GetValidatorAtIndex(&bind.CallOpts{}, big.NewInt(oldCount.Int64()-1))
	r.NoError(err)
	r.Equal(validator.Code, crypto.PubkeyToAddress(sender.PublicKey))
	r.Equal(validator.Deposit, minDeposit)

	// minimum deposit should increase by 1
	latestMinDeposit, err := s.contract.GetMinimumDeposit(&bind.CallOpts{})
	r.NoError(err)
	r.Equal(big.NewInt(minDeposit.Int64()+1), latestMinDeposit)
}

func (s *ElectionContractSuite) TestLeave_NotValidator() {
	r := s.Require()
	sender := s.notOwner

	opts := bind.NewKeyedTransactor(sender)
	_, err := s.contract.Leave(opts)
	r.Equal(errTransactionFailed, err)
}

func (s *ElectionContractSuite) TestLeave_Validator() {
	r := s.Require()

	oldCount, err := s.contract.GetValidatorCount(&bind.CallOpts{})
	r.NoError(err)

	// genesis validator
	sender := s.genesis
	senderAddr := crypto.PubkeyToAddress(sender.PublicKey)
	opts := bind.NewKeyedTransactor(sender)
	_, err = s.contract.Leave(opts)
	r.NoError(err)

	s.backend.Commit()

	// should not be a validator anymore
	isValidator, err := s.contract.IsValidator(&bind.CallOpts{}, crypto.PubkeyToAddress(s.notOwner.PublicKey))
	r.NoError(err)
	r.False(isValidator)

	// number of validators should be equal to oldCount - 1
	latestCount, err := s.contract.GetValidatorCount(&bind.CallOpts{})
	r.NoError(err)
	r.Equal(big.NewInt(oldCount.Int64()-1).Int64(), latestCount.Int64())

	// latest deposit should have a release date
	depositCount, err := s.contract.GetDepositCount(&bind.CallOpts{From: senderAddr})
	r.NoError(err)
	currentDeposit, err := s.contract.GetDepositAtIndex(&bind.CallOpts{From: senderAddr}, big.NewInt(depositCount.Int64()-1))
	r.NoError(err)
	r.NotZero(currentDeposit.ReleasedAt)
}

func (s *ElectionContractSuite) TestRedeemFunds_NoDeposits() {
	r := s.Require()

	sender := s.notOwner
	opts := bind.NewKeyedTransactor(sender)
	_, err := s.contract.RedeemFunds(opts)
	r.NoError(err)

	s.backend.Commit()

	// balance should be < initial balance (gas used on this call + no deposits)
	senderBalance, err := s.backend.BalanceAt(context.TODO(), crypto.PubkeyToAddress(sender.PublicKey), s.backend.CurrentBlock().Number())
	r.NoError(err)
	r.True(senderBalance.Cmp(s.balance) < 0)
}

func (s ElectionContractSuite) TestRedeemFunds_HasLockedDeposit() {
	// @NOTE (rgeraldes) - default unbonding value is big enough to
	// keep the deposit locked for some time
	r := s.Require()

	// leave the election
	sender := s.genesis
	opts := bind.NewKeyedTransactor(sender)
	_, err := s.contract.Leave(opts)
	r.NoError(err)
	s.backend.Commit()

	// redeem funds
	_, err = s.contract.RedeemFunds(opts)
	r.NoError(err)
	s.backend.Commit()

	// balance should be < initial balance (gasted used
	// on both calls + no deposits past the unbond period)
	senderBalance, err := s.backend.BalanceAt(context.TODO(), crypto.PubkeyToAddress(sender.PublicKey), s.backend.CurrentBlock().Number())
	r.NoError(err)
	r.True(senderBalance.Cmp(s.balance) < 0)
}

func (s *ElectionContractSuite) TestRedeemFunds_HasUnlockedDeposit() {
	r := s.Require()

	// deploy a new version of the contract with unbonding period of 0
	s.DeployElectionContract(s.baseDeposit, s.maxValidators, big.NewInt(0))

	// initial number of deposits of the genesis validator

	// leave the election
	sender := s.genesis
	opts := bind.NewKeyedTransactor(sender)
	_, err := s.contract.Leave(opts)
	r.NoError(err)
	s.backend.Commit()

	// redeem funds
	_, err = s.contract.RedeemFunds(opts)
	r.NoError(err)
	s.backend.Commit()

	// balance should be > initial balance (assuming that deposit > gas used)
	senderBalance, err := s.backend.BalanceAt(context.TODO(), crypto.PubkeyToAddress(sender.PublicKey), s.backend.CurrentBlock().Number())
	r.NoError(err)
	r.True(senderBalance.Cmp(s.balance) > 0)

	// the number of deposits of the user should be the initial number - 1
}

func (s *ElectionContractSuite) TestRedeemFunds_HasLockedAndUnlockedDeposits() {
	r := s.Require()

	// deploy a new version of the contract with a unbonding period of 0 days
	s.DeployElectionContract(s.baseDeposit, s.maxValidators, big.NewInt(0))

	// leave the election in order to unlock the current deposit - unbond
	// period is 0
	sender := s.genesis
	senderAddr := crypto.PubkeyToAddress(sender.PublicKey)
	opts := bind.NewKeyedTransactor(sender)
	_, err := s.contract.Leave(opts)
	r.NoError(err)
	s.backend.Commit()

	// join the election again in order to have a deposit with no
	// release date (locked deposit)
	minDeposit, err := s.contract.GetMinimumDeposit(&bind.CallOpts{})
	r.NoError(err)
	opts = bind.NewKeyedTransactor(sender)
	opts.Value = minDeposit
	_, err = s.contract.Join(opts)
	r.NoError(err)
	s.backend.Commit()

	// the user should have two deposits
	count, err := s.contract.GetDepositCount(&bind.CallOpts{From: senderAddr})
	r.NoError(err)
	r.EqualValues(2, count)

	// redeem funds
	opts = bind.NewKeyedTransactor(sender)
	_, err = s.contract.RedeemFunds(opts)
	r.NoError(err)
	s.backend.Commit()

	// the user should have one deposit (the current one)
	count, err = s.contract.GetDepositCount(&bind.CallOpts{From: senderAddr})
	r.NoError(err)
	r.EqualValues(1, count)
	deposit, err := s.contract.GetDepositAtIndex(&bind.CallOpts{From: senderAddr}, big.NewInt(0))
	r.NoError(err)
	r.Equal(deposit.Amount, minDeposit)
	r.Equal(deposit.ReleasedAt.Int64(), 0)

}
