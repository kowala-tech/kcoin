package network

import (
	"crypto/ecdsa"
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

// NetworkContractSuite contains the test suite of the network contract
type NetworkContractSuite struct {
	suite.Suite
	minDeposit      *big.Int
	owner, notOwner *ecdsa.PrivateKey
	contract        *NetworkContract
}

func TestNetworkContractSuite(t *testing.T) {
	suite.Run(t, new(NetworkContractSuite))
}

// SetupSuite is executed once for every suite
func (s *NetworkContractSuite) SetupSuite() {
	require := s.Require()

	// contract owners
	owner, err := crypto.GenerateKey()
	require.NoError(err)
	require.NotNil(owner)
	notOwner, err := crypto.GenerateKey()
	require.NoError(err)
	require.NotNil(notOwner)
	s.owner, s.notOwner = owner, notOwner

	// minimum deposit
	s.minDeposit = big.NewInt(100000)
}

// SetupTest is executed once for every test
func (s *NetworkContractSuite) SetupTest() {
	require := s.Require()

	// @NOTE (rgeraldes) - fund the owner account with enough balance to
	// deploy the contract and to cover the genesis validator deposit
	ownerAddr := crypto.PubkeyToAddress(s.owner.PublicKey)
	notOwnerAddr := crypto.PubkeyToAddress(s.notOwner.PublicKey)
	simulatedBackend := backends.NewSimulatedBackend(core.GenesisAlloc{
		ownerAddr:    core.GenesisAccount{Balance: new(big.Int).Mul(big.NewInt(100), new(big.Int).SetUint64(params.Ether))},
		notOwnerAddr: core.GenesisAccount{Balance: new(big.Int).Mul(big.NewInt(100), new(big.Int).SetUint64(params.Ether))},
	})

	// deploy the network
	_, _, contract, err := DeployNetworkContract(bind.NewKeyedTransactor(s.owner), simulatedBackend, s.minDeposit, common.Address{})
	require.NoError(err)
	require.NotZero(contract)
	s.contract = contract

	// commit
	simulatedBackend.Commit()
}

func (s *NetworkContractSuite) TestConstructor() {
	require := s.Require()

	deployedMinDeposit, err := s.contract.MinDeposit(&bind.CallOpts{})
	require.NoError(err)
	require.Equal(s.minDeposit, deployedMinDeposit)

	deployedMinDepositUB, err := s.contract.MinDepositUpperBound(&bind.CallOpts{})
	require.NoError(err)

	deployedMinDepositLB, err := s.contract.MinDepositLowerBound(&bind.CallOpts{})
	require.NoError(err)

}

func (s *NetworkContractSuite) TestSetMinDeposit() {
	require := s.Require()

	deployedMinDepositLB, err := s.contract.MinDepositLowerBound(&bind.CallOpts{})
	require.NoError(err)

	deployedMinDepositUB, err := s.contract.MinDepositUpperBound(&bind.CallOpts{})
	require.NoError(err)

	testCases := []struct {
		name          string
		user          *ecdsa.PrivateKey
		deposit       *big.Int
		expectedError error
	}{
		{
			"user is not the contract owner",
			s.notOwner,
			nil,
			nil,
		},
		{
			"user is the contract owner but the new minimum deposit value is lower then the lower bound",
			s.owner,
			deployedMinDepositLB.Sub(deployedMinDepositLB, big.NewInt(1)),
			nil,
		},
		{
			"user is the contract owner but the new minimum deposit value is bigger then the upper bound",
			s.owner,
			deployedMinDepositUB.Add(deployedMinDepositUB, big.NewInt(1)),
			nil,
		},
		{
			"user is the contract owner and the new deposit minimum is within the valid range",
			s.owner,
			nil,
			nil,
		},
	}

	for _, testCase := range testCases {
		s.T().Run(testCase.name, func(*testing.T) {
			_, err := s.contract.SetMinDeposit(bind.NewKeyedTransactor(testCase.user), testCase.deposit)
			s.EqualValues(testCase.expectedError, err)

			if err == nil {
				// make sure that the value was modified

				// make sure that the bounds were not modified
				currentMinDepositLB, err := s.contract.MinDepositLowerBound(&bind.CallOpts{})
				require.NoError(err)
				currentMinDepositUB, err := s.contract.MinDepositUpperBound(&bind.CallOpts{})
				require.NoError(err)
				s.EqualValues(deployedMinDepositLB, currentMinDepositLB)
				s.EqualValues(deployedMinDepositUB, currentMinDepositUB)
			}
		})
	}
}

func (s *NetworkContractSuite) TestSetMinDepositLowerBound() {
	require := s.Require()

	deployedMinDepositUB, err := s.contract.MinDepositUpperBound(&bind.CallOpts{})
	require.NoError(err)

	testCases := []struct {
		name          string
		user          *ecdsa.PrivateKey
		limit         *big.Int
		expectedError error
	}{
		{
			"user is not the contract owner",
			s.notOwner,
			nil,
			nil,
		},
		{
			"user is the contract owner but the new lower bound is not smaller than the current upper bound",
			s.owner,
			deployedMinDepositUB,
			nil,
		},
		{
			"user is the contract owner and the limit is valid",
			s.owner,
			deployedMinDepositUB.Sub(deployedMinDepositUB, big.NewInt(1)),
			nil,
		},
	}

	for _, testCase := range testCases {
		s.T().Run(testCase.name, func(*testing.T) {
			_, err := s.contract.SetMinDepositLowerBound(bind.NewKeyedTransactor(testCase.user), testCase.limit)
			s.EqualValues(testCase.expectedError, err)

			if err == nil {
				// make sure that the value was updated
			}
		})
	}
}

func (s *NetworkContractSuite) TestSetMinDepositUpperBound() {
	require := s.Require()

	deployedMinDepositLB, err := s.contract.MinDepositLowerBound(&bind.CallOpts{})
	require.NoError(err)

	testCases := []struct {
		name          string
		user          *ecdsa.PrivateKey
		limit         *big.Int
		expectedError error
	}{
		{
			"user is not the contract owner",
			s.notOwner,
			nil,
			nil,
		},
		{
			"user is the contract owner but the new upper bound is not bigger than the current lower bound",
			s.owner,
			deployedMinDepositLB,
			nil,
		},
		{
			"user is the contract owner and the limit is valid",
			s.owner,
			deployedMinDepositLB.Add(deployedMinDepositLB, big.NewInt(1)),
			nil,
		},
	}

	for _, testCase := range testCases {
		s.T().Run(testCase.name, func(*testing.T) {
			_, err := s.contract.SetMinDepositUpperBound(bind.NewKeyedTransactor(testCase.user), testCase.limit)
			s.EqualValues(testCase.expectedError, err)

			if err == nil {
				// make sure that the value was updated
			}
		})
	}
}
