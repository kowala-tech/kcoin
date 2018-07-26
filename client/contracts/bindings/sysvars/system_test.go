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
	"github.com/kowala-tech/kcoin/client/crypto"
	"github.com/kowala-tech/kcoin/client/params"
	"github.com/stretchr/testify/suite"
)

var (
	// roles
	user, _     = crypto.GenerateKey()
	governor, _ = crypto.GenerateKey()

	// contracts
	systemVarsAddr = common.HexToAddress("0x17C56D5aC0cddFd63aC860237197827cB4639CDA")
)

func getDefaultOpts() genesis.Options {
	baseDeposit := uint64(20)
	superNodeAmount := uint64(6000000)
	tokenHolder := genesis.TokenHolder{
		Address:   getAddress(user).Hex(),
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
			Validators: []genesis.Validator{
				{
					Address: tokenHolder.Address,
					Deposit: tokenHolder.NumTokens,
				},
			},
			MiningToken: &genesis.MiningTokenOpts{
				Name:     "mUSD",
				Symbol:   "mUSD",
				Cap:      20000000,
				Decimals: 18,
				Holders:  []genesis.TokenHolder{tokenHolder},
			},
		},
		Governance: &genesis.GovernanceOpts{
			// Origin needs to be the same as the testnet for now since we are using hardcoded addresses
			Origin:           "0x259be75d96876f2ada3d202722523e9cd4dd917d",
			Governors:        []string{getAddress(governor).Hex()},
			NumConfirmations: 1,
		},
		DataFeedSystem: &genesis.DataFeedSystemOpts{
			MaxNumOracles: 10,
			FreezePeriod:  32,
			BaseDeposit:   1,
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
	if strings.Contains(testName, "TestMintedAmount") {
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
	sysvars, err := sysvars.NewSystemVars(systemVarsAddr, backend)
	req.NoError(err)
	req.NotNil(sysvars)
	suite.sysvars = sysvars
}

func (suite *SystemVarsSuite) TestDeploySystemVars() {
	req := suite.Require()

	initialPrice := new(big.Int)
	new(big.Float).Mul(new(big.Float).SetFloat64(suite.opts.SystemVars.InitialPrice), big.NewFloat(params.Kcoin)).Int(initialPrice)

	storedPrice, err := suite.sysvars.CurrencyPrice(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedPrice)
	req.Equal(initialPrice, storedPrice)

	storedPrevPrice, err := suite.sysvars.PrevCurrencyPrice(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedPrevPrice)
	req.Equal(initialPrice, storedPrevPrice)

	mintedAmount := new(big.Int)
	for _, prefundedAccount := range suite.opts.PrefundedAccounts {
		mintedAmount.Add(mintedAmount, new(big.Int).Mul(new(big.Int).SetUint64(prefundedAccount.Balance), new(big.Int).SetUint64(params.Kcoin)))
	}

	storedCurrencySupply, err := suite.sysvars.CurrencySupply(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedCurrencySupply)
	req.Equal(mintedAmount, storedCurrencySupply)

	storedMintedReward, err := suite.sysvars.MintedReward(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedMintedReward)
	req.Equal(mintedAmount, storedMintedReward)
}

// getAddress return the address of the given private key
func getAddress(privateKey *ecdsa.PrivateKey) common.Address {
	return crypto.PubkeyToAddress(privateKey.PublicKey)
}
