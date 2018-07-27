package konsensus_test

import (
	"testing"

	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/accounts/abi/bind/backends"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/consensus/konsensus"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/oracle"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/sysvars"
	"github.com/kowala-tech/kcoin/client/crypto"
	"github.com/kowala-tech/kcoin/client/knode/genesis"
	"github.com/stretchr/testify/require"
)

var (
	// roles
	validator, _    = crypto.GenerateKey()
	deregistered, _ = crypto.GenerateKey()
	user, _         = crypto.GenerateKey()
	governor, _     = crypto.GenerateKey()
	author, _       = crypto.HexToECDSA("bfef37ae9ac5d5e7ebbbefc19f4e1f572a7ca7aa0d28e527b7d62950951cc5eb")

	// contracts
	sysVarsAddr   = common.HexToAddress("0x")
	oracleMgrAddr = common.HexToAddress("0x")
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
				InitialPrice:  1,
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

func TestSetPrice(t *testing.T) {
	// create genesis
	opts := getDefaultOpts()
	require.NotNil(t, opts)
	genesis, err := genesis.Generate(opts)
	require.NoError(t, err)
	require.NotNil(t, genesis)

	// create backend
	backend := backends.NewSimulatedBackend(genesis.Alloc)
	require.NotNil(t, backend)

	// SystemVars instance
	vars, err := sysvars.NewSystemVars(sysVarsAddr, backend)
	require.NoError(t, err)
	require.NotNil(t, vars)

	// OracleMgr instance
	oracleMgr, err := oracle.NewOracleMgr(oracleMgrAddr, backend)
	require.NoError(t, err)
	require.NotNil(t, oracleMgr)

	state, err := backend.State()
	require.NoError(t, err)
	require.NotNil(t, state)

	// register oracle
	transactOpts := bind.NewKeyedTransactor(user)
	oracleMgr.RegisterOracle(transactOpts)
	backend.Commit()

	// submit price
	newPrice := common.Big2
	oracleMgr.SubmitPrice(transactOpts, newPrice)
	backend.Commit()

	// system wrapper
	sys := konsensus.Sys(state, vars, oracleMgr)
	require.NotNil(t, system)

	newPrice, err := oracleMgr.AveragePrice()
	require.NoError(t, err)
	require.NotNil(t, avgPrice)

	price, err := vars.CurrencyPrice()
	require.NoError(t, err)
	require.NotNil(t, price)

	// set system price
	sys.SetPrice(newPrice)

	// prev price must have the old price
	prevPrice, err := vars.PrevCurrencyPrice()
	require.NoError(t, err)
	require.NotNil(t, prevPrice)
	require.Equal(t, price, prevPrice)

	// current price must be equal to the new price
	currentPrice, err := vars.CurrencyPrice()
	require.NoError(t, err)
	require.NotNil(t, currentPrice)
	require.Equal(t, newPrice, currentPrice)

	// oracle's hasSubmittedPrice must be false
	oracle, err := oracleMgr.GetOracleAtIndex(&bind.CallOpts{}, common.Big0)
	require.NoError(t, err)
	require.NotNil(t, oracle)
	require.False(t, oracle.Price)

	// oracleMgr average price must be 0
	avgPrice, err := oracleMgr.AveragePrice()
	require.NoError(t, err)
	require.NotNil(t, avgPrice)
	require.Zero(t, avgPrice.Uint64())
}

func TestMint(t *testing.T) {
	// create genesis
	opts := getDefaultOpts()
	require.NotNil(t, opts)
	genesis, err := genesis.Generate(opts)
	require.NoError(t, err)
	require.NotNil(t, genesis)

	// create backend
	backend := backends.NewSimulatedBackend(genesis.Alloc)
	require.NotNil(t, backend)

	// SystemVars instance
	vars, err := sysvars.NewSystemVars(sysVarsAddr, backend)
	require.NoError(t, err)
	require.NotNil(t, vars)

	state, err := backend.State()
	require.NoError(t, err)
	require.NotNil(t, state)

	// system wrapper
	sys := konsensus.Sys(state, vars, oracleMgr)
	require.NotNil(t, system)

	// mint kUSD
	sys.Mint(newPrice)

	// prev price must have the old price
	prevPrice, err := vars.PrevCurrencyPrice()
	require.NoError(t, err)
	require.NotNil(t, prevPrice)
	require.Equal(t, price, prevPrice)

	// current price must be equal to the new price
	currentPrice, err := vars.CurrencyPrice()
	require.NoError(t, err)
	require.NotNil(t, currentPrice)
	require.Equal(t, newPrice, currentPrice)

	// oracle's hasSubmittedPrice must be false
	oracle, err := oracleMgr.GetOracleAtIndex(&bind.CallOpts{}, common.Big0)
	require.NoError(t, err)
	require.NotNil(t, oracle)
	require.False(t, oracle.Price)

	// oracleMgr average price must be 0
	avgPrice, err := oracleMgr.AveragePrice()
	require.NoError(t, err)
	require.NotNil(t, avgPrice)
	require.Zero(t, avgPrice.Uint64())
}
