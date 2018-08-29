package sysvars_test

import (
	"math/big"
	"testing"

	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/accounts/abi/bind/backends"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/sysvars"
	"github.com/kowala-tech/kcoin/client/core"
	"github.com/kowala-tech/kcoin/client/crypto"
	"github.com/kowala-tech/kcoin/client/params"
	"github.com/stretchr/testify/suite"
)

var (
	owner, _ = crypto.GenerateKey()
)

type SystemVarsSuite struct {
	suite.Suite
	backend *backends.SimulatedBackend
}

func TestSystemVarsSuite(t *testing.T) {
	suite.Run(t, new(SystemVarsSuite))
}

func (suite *SystemVarsSuite) BeforeTest(suiteName, testName string) {
	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		crypto.PubkeyToAddress(owner.PublicKey): core.GenesisAccount{
			Balance: new(big.Int).Mul(new(big.Int).SetUint64(100), new(big.Int).SetUint64(params.Kcoin)),
		},
	})
	suite.backend = backend
}

func (suite *SystemVarsSuite) TestDeploy() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// deploy system vars contract
	initialPrice := common.Big1
	initialSupply := common.Big32
	_, _, systemVarsContract, err := sysvars.DeploySystemVars(transactOpts, suite.backend, initialPrice, initialSupply)
	req.NoError(err)
	req.NotNil(systemVarsContract)

	suite.backend.Commit()

	storedCurrencyPrice, err := systemVarsContract.CurrencyPrice(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedCurrencyPrice)

	storedPrevCurrencyPrice, err := systemVarsContract.PrevCurrencyPrice(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedPrevCurrencyPrice)

	storedCurrencySupply, err := systemVarsContract.CurrencySupply(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedCurrencySupply)

	// @TODO (rgeraldes) - minted reward?
}
