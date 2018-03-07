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

var (
	// minimum deposit
	minDeposit = big.NewInt(100000)
	// genesis address

)

// NetworkContractSuite contains the test suite of the network contract
type NetworkContractSuite struct {
	suite.Suite
	owner    *ecdsa.PrivateKey
	contract *NetworkContract
}

func TestNetworkContractSuite(t *testing.T) {
	suite.Run(t, new(NetworkContractSuite))
}

// SetupTest is executed once for every test
func (s *NetworkContractSuite) SetupTest() {
	require := s.Require()

	// generate a new owner
	owner, err := crypto.GenerateKey()
	require.NoError(err)
	require.NotNil(owner)
	s.owner = owner

	// @NOTE (rgeraldes) - fund the owner account with enough balance to
	// deploy the contract and to cover the genesis validator deposit
	ownerAddr := crypto.PubkeyToAddress(owner.PublicKey)
	simulatedBackend := backends.NewSimulatedBackend(core.GenesisAlloc{
		ownerAddr: core.GenesisAccount{
			// 100 ether
			Balance: new(big.Int).Mul(big.NewInt(100), new(big.Int).SetUint64(params.Ether)),
		},
	})

	// deploy the network
	_, _, contract, err := DeployNetworkContract(bind.NewKeyedTransactor(owner), simulatedBackend, minDeposit, common.Address{})
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
	require.Equal(minDeposit, deployedMinDeposit)

	deployedMinDepositUB, err := s.contract.MinDepositUpperBound(&bind.CallOpts{})
	require.NoError(err)

	deployedMinDepositLB, err := s.contract.MinDepositLowerBound(&bind.CallOpts{})
	require.NoError(err)

}

func (s *NetworkContractSuite) TestSetMinDeposit() {
	require := s.Require()

	testCases := []struct {
		name          string
		input         *big.Int
		expectedError error
	}{
		{
			"user is not the contract owner",
			big.NewInt(1),
		},
	}

	for _, testCase := range testCases {
		s.T().Run(testCase.name, func(*testing.T) {
			reply, err := s.backend.addComment(tc.comment, tc.userID)
			s.EqualValues(tc.expectedError, err)
			if err == nil {
				s.NotNil(reply)
				s.NotZero(reply.CommentID)
			}
		})
	}
}

func (s *NetworkContractSuite) TestSetMinDepositUpperBound() {}

func (s *NetworkContractSuite) TestSetMinDepositLowerBound() {}
