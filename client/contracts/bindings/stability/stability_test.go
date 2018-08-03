package stability_test

import (
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/accounts/abi/bind/backends"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/stability"
	"github.com/kowala-tech/kcoin/client/core"
	"github.com/kowala-tech/kcoin/client/crypto"
	"github.com/kowala-tech/kcoin/client/params"
	"github.com/stretchr/testify/suite"
)

var (
	user, _  = crypto.GenerateKey()
	owner, _ = crypto.GenerateKey()
)

type StabilityContractSuite struct {
	suite.Suite
}

func TestStabilityContractSuite(t *testing.T) {
	suite.Run(t, new(StabilityContractSuite))
}

func (suite *StabilityContractSuite) TestDeploy() {
	req := suite.Require()

	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		getAddress(owner): core.GenesisAccount{
			Balance: new(big.Int).Mul(new(big.Int).SetUint64(1), new(big.Int).SetUint64(params.Kcoin)),
		},
	})

	minDeposit := new(big.Int).SetUint64(100)
	transactOpts := bind.NewKeyedTransactor(owner)
	_, _, stabilityContract, err := stability.DeployStability(transactOpts, backend, minDeposit, common.HexToAddress("0x161ad311F1D66381C17641b1B73042a4CA731F9f"))
	req.NoError(err)
	req.NotNil(stabilityContract)

	storedMinDeposit, err := stabilityContract.MinDeposit(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedMinDeposit)
	req.Equal(minDeposit, storedMinDeposit)

}

// getAddress return the address of the given private key
func getAddress(privateKey *ecdsa.PrivateKey) common.Address {
	return crypto.PubkeyToAddress(privateKey.PublicKey)
}
