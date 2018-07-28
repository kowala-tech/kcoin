package konsensus_test

import (
	"testing"
	
	"github.com/kowala-tech/kcoin/client/knode/genesis"
)

var (
	user, _     = crypto.GenerateKey()
	governor, _ = crypto.GenerateKey()
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
			Origin:           "0x259be75d96876f2ada3d202722523e9cd4dd917d",
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
		},
	}

	return opts
}

func TestFinalize(t *testing.T) {
	chainID := params.TestnetChainConfig.ChainID

	// create genesis
	opts := getDefaultOpts()
	require.NotNil(t, opts)
	genesis, err := genesis.Generate(opts)
	require.NoError(t, err)
	require.NotNil(t, genesis)

	// create backend
	backend := backends.NewSimulatedBackend(genesis.Alloc)
	require.NotNil(t, backend)

	// OracleMgr instance
	oracleBinding, err := oracle.Bind(backend, chainID)
	require.NoError(t, err)
	require.NotNil(t, oracleBinding)
	oracleMgr := oracleBinding.(oracle.Manager)

	// SystemVars instance
	sysBinding, err := sysvars.Bind(backend, chainID)
	require.NoError(t, err)
	require.NotNil(t, sysBinding)
	systemVars := sysBinding.(sysvars.System)

	// choose consensus engine
	engine := konsensus.New(&params.KonsensusConfig{}, oracleMgr, systemVars)

	// change backend consensus engine
	backend.WithEngine(engine)

	// register oracle
	_, err := oracleMgr.RegisterOracle(user)
	require.NoError(t, err)

	// oracle submits price
	submittedPrice := common.Big2
	_, err := oracleMgr.SubmitPrice(user, submittedPrice)
	require.NoError(t, err)
	
	backend.Commit()

	// minted reward must be > 0
	
	// supply must be equal to the initial funding + minted reward

}