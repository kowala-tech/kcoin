package sysvars_test

import (
	"crypto/ecdsa"
	"math/big"
	"strings"
	"testing"

	"github.com/kowala-tech/kcoin/client/knode/genesis"

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
	owner, _       = crypto.GenerateKey()
	initialBalance = new(big.Int).Mul(common.Big32, new(big.Int).SetUint64(params.Kcoin)) // 32 Kcoin
	initialPrice   = new(big.Int).Mul(common.Big1, new(big.Int).SetUint64(params.Kcoin))  // $1
)

func getDefaultOpts() genesis.Options {
	baseDeposit := uint64(20)
	superNodeAmount := uint64(6000000)
	tokenHolder := genesis.TokenHolder{
		Address:   getAddress(validator).Hex(),
		NumTokens: superNodeAmount,
	}

	opts := genesis.Options{
		Network: "test",
		SystemVars: &genesis.SystemVarsOpts{
			InitialPrice: 1,
		},
		Consensus: &genesis.ConsensusOpts{
			Engine:           "konsensus",
			MaxNumValidators: 10,
			FreezePeriod:     30,
			BaseDeposit:      baseDeposit,
			SuperNodeAmount:  superNodeAmount,
			Validators: []genesis.Validator{{
				Address: tokenHolder.Address,
				Deposit: tokenHolder.NumTokens,
			}},
			MiningToken: &genesis.MiningTokenOpts{
				Name:     "mUSD",
				Symbol:   "mUSD",
				Cap:      20000000,
				Decimals: 18,
				Holders:  []genesis.TokenHolder{tokenHolder, {Address: getAddress(user).Hex(), NumTokens: 10000000}},
			},
		},
		Governance: &genesis.GovernanceOpts{
			Origin:           getAddress(author).Hex(),
			Governors:        []string{getAddress(governor).Hex()},
			NumConfirmations: 1,
		},
		DataFeedSystem: &genesis.DataFeedSystemOpts{
			MaxNumOracles: 10,
			FreezePeriod:  0,
			BaseDeposit:   0,
			Price: genesis.PriceOpts{
				SyncFrequency: 600,
				UpdatePeriod:  30,
			},
		},
		PrefundedAccounts: []genesis.PrefundedAccount{
			{
				Address: tokenHolder.Address,
				Balance: 10,
			},
			{
				Address: getAddress(governor).Hex(),
				Balance: 10,
			},
			{
				Address: getAddress(user).Hex(),
				Balance: 10,
			},
			{
				Address: getAddress(deregistered).Hex(),
				Balance: 10,
			},
		},
	}

	return opts
}

type SystemVarsSuite struct {
	suite.Suite
	backend *backends.SimulatedBackend
	opts    genesis.Options
	sysvars *sysvars.SystemVars
}

func TestSystemVarsSuite(t *testing.T) {
	suite.Run(t, new(SystemVarsSuite))
}

func (suite *SystemVarsSuite) BeforeTest(suiteName, testName string) {
	if strings.Contains(testName, "TestDeploy") {
		return
	}

	req := suite.Require()

	// create genesis
	opts := getDefaultOpts()
	req.NotNil(opts)
	suite.opts = opts

	genesis, err := genesis.Generate(opts)
	req.NoError(err)
	req.NotNil(genesis)

	backend := backends.NewSimulatedBackend(genesis.Alloc)
	req.NotNil(backend)
	suite.backend = backend

	// SystemVars instance
	_, _, sysvars, err := sysvars.DeploySystemVars(bind.NewKeyedTransactor(owner), backend, initialPrice)
	req.NoError(err)
	req.NotNil(sysvars)
	suite.sysvars = sysvars

	suite.backend.Commit()
}

func (suite *SystemVarsSuite) TestDeploy() {
	req := suite.Require()

	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		getAddress(owner): core.GenesisAccount{
			Balance: new(big.Int).Mul(new(big.Int).SetUint64(100), new(big.Int).SetUint64(params.Kcoin)),
		},
	})

	// SystemVars instance
	transactOpts := bind.NewKeyedTransactor(owner)
	_, _, sysvars, err := sysvars.DeploySystemVars(transactOpts, backend, initialPrice)
	req.NoError(err)
	req.NotNil(sysvars)

	backend.Commit()

	storedPrice
}

// getAddress return the address of the given private key
func getAddress(privateKey *ecdsa.PrivateKey) common.Address {
	return crypto.PubkeyToAddress(privateKey.PublicKey)
}

/*
func (suite *OracleMgrSuite) TestDeployOracleMgr_InitialPriceEqualsZero() {
	req := suite.Require()

	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		getAddress(governor): core.GenesisAccount{
			Balance: new(big.Int).Mul(new(big.Int).SetUint64(100), new(big.Int).SetUint64(params.Kcoin)),
		},
	})
	req.NotNil(backend)

	initialPrice := common.Big0
	baseDeposit := new(big.Int).SetUint64(100)
	maxNumOracles := new(big.Int).SetUint64(100)
	freezePeriod := new(big.Int).SetUint64(10)
	syncFrequency := new(big.Int).SetUint64(20)
	updatePeriod := new(big.Int).SetUint64(5)

	transactOpts := bind.NewKeyedTransactor(governor)
	_, _, _, err := oracle.DeployOracleMgr(transactOpts, backend, initialPrice, baseDeposit, maxNumOracles, freezePeriod, syncFrequency, updatePeriod, validatorMgrAddr)
	req.Error(err, "initial price cannot be zero")
}
*/
