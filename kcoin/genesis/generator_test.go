package genesis

import (
	"flag"
)

var update = flag.Bool("update", false, "update .golden files")

/*
func TestItWritesTheGeneratedFileToAWriter(t *testing.T) {
	opt := Options{
		Network:                        "test",
		MaxNumValidators:               "5",
		UnbondingPeriod:                "5",
		AccountAddressGenesisValidator: "0xe2ac86cbae1bbbb47d157516d334e70859a1bee4",
		PrefundedAccounts: []PrefundedAccount{
			{
				AccountAddress: "0xe2ac86cbae1bbbb47d157516d334e70859a1bee4",
				Balance:        "15",
			},
		},
	}

	generatedGenesis, err := GenerateGenesis(opt)

	assert.NoError(t, err)

	fileName := filepath.Join("testfiles", "testnet_default.json")
	if *update {
		t.Log("update golden file")

		out, err := json.MarshalIndent(generatedGenesis, "", "  ")
		if err != nil {
			t.Fatal("Error marshaling generated genesis to create golden file.")
		}

		if err := ioutil.WriteFile(fileName, out, 0644); err != nil {
			t.Fatal("Error saving golden file.")
		}
	}

	contents, err := ioutil.ReadFile(fileName)
	assert.NoError(t, err)

	var expectedGenesis = new(core.Genesis)
	err = json.Unmarshal(contents, expectedGenesis)
	assert.NoError(t, err)

	assertEqualGenesis(t, expectedGenesis, generatedGenesis)
}

//assertEqualGenesis checks if two genesis are the same, it ignores some fields as the timestamp that
//will be always different when it is generated.
func assertEqualGenesis(t *testing.T, expectedGenesis *core.Genesis, generatedGenesis *core.Genesis) {
	assert.Equal(t, expectedGenesis.ExtraData, generatedGenesis.ExtraData)
	assert.Equal(t, expectedGenesis.Config, generatedGenesis.Config)
	assert.Equal(t, expectedGenesis.GasLimit, generatedGenesis.GasLimit)
	assert.Equal(t, expectedGenesis.GasUsed, generatedGenesis.GasUsed)
	assert.Equal(t, expectedGenesis.Coinbase, generatedGenesis.Coinbase)
	assert.Equal(t, expectedGenesis.ParentHash, generatedGenesis.ParentHash)

	assert.Len(t, expectedGenesis.Alloc, len(generatedGenesis.Alloc))

	address := DefaultSmartContractsOwner
	expectedAlloc := core.GenesisAccount{Balance: new(big.Int).Mul(common.Big1, big.NewInt(params.Ether))}
	assert.Equal(t, generatedGenesis.Alloc[address], expectedAlloc)
}
func TestItFailsWhenRunningHandlerWithInvalidCommandValues(t *testing.T) {
	baseValidCommand := Options{
		Network:                        "test",
		MaxNumValidators:               "1",
		UnbondingPeriod:                "1",
		AccountAddressGenesisValidator: "0xe2ac86cbae1bbbb47d157516d334e70859a1bee4",
		PrefundedAccounts: []PrefundedAccount{
			{
				AccountAddress: "0xe2ac86cbae1bbbb47d157516d334e70859a1bee4",
				Balance:        "15",
			},
		},
	}

	tests := []struct {
		TestName                string
		InvalidCommandFromValid func(command Options) Options
		ExpectedError           error
	}{
		{
			TestName: "Invalid Network",
			InvalidCommandFromValid: func(command Options) Options {
				command.Network = "fakeNetwork"
				return command
			},
			ExpectedError: ErrInvalidNetwork,
		},
		{
			TestName: "Empty max number of validators",
			InvalidCommandFromValid: func(command Options) Options {
				command.MaxNumValidators = ""
				return command
			},
			ExpectedError: ErrEmptyMaxNumValidators,
		},
		{
			TestName: "Empty unbonding period of days",
			InvalidCommandFromValid: func(command Options) Options {
				command.UnbondingPeriod = ""
				return command
			},
			ExpectedError: ErrEmptyUnbondingPeriod,
		},
		{
			TestName: "Empty wallet address of genesis validator",
			InvalidCommandFromValid: func(command Options) Options {
				command.AccountAddressGenesisValidator = ""
				return command
			},
			ExpectedError: ErrEmptyWalletAddressValidator,
		},
		{
			TestName: "Invalid wallet address less than 20 bytes with Hex prefix",
			InvalidCommandFromValid: func(command Options) Options {
				command.AccountAddressGenesisValidator = "0xe2ac86cbae1bbbb47d157516d334e70859a1be"
				return command
			},
			ExpectedError: ErrInvalidWalletAddressValidator,
		},
		{
			TestName: "Invalid wallet address less than 20 bytes without Hex prefix",
			InvalidCommandFromValid: func(command Options) Options {
				command.AccountAddressGenesisValidator = "e2ac86cbae1bbbb47d157516d334e70859a1be"
				return command
			},
			ExpectedError: ErrInvalidWalletAddressValidator,
		},
		{
			TestName: "Empty prefunded accounts",
			InvalidCommandFromValid: func(command Options) Options {
				command.PrefundedAccounts = []PrefundedAccount{}
				return command
			},
			ExpectedError: ErrEmptyPrefundedAccounts,
		},
		{
			TestName: "Prefunded accounts does not include validator address",
			InvalidCommandFromValid: func(command Options) Options {
				command.PrefundedAccounts = []PrefundedAccount{
					{
						AccountAddress: "0xaaaaaacbae1bbbb47d157516d334e70859a1bee4",
						Balance:        "15",
					},
				}
				return command
			},
			ExpectedError: ErrWalletAddressValidatorNotInPrefundedAccounts,
		},
		{
			TestName: "Prefunded accounts has invalid account.",
			InvalidCommandFromValid: func(command Options) Options {
				command.PrefundedAccounts = []PrefundedAccount{
					{
						AccountAddress: "0xe2ac86cbae1bbbb47d157516d334e70859a1bee4",
						Balance:        "15",
					},
					{
						AccountAddress: "0xe286cbae1bbbb47d157516d334e70859a1bee4",
						Balance:        "15",
					},
				}
				return command
			},
			ExpectedError: ErrInvalidAddressInPrefundedAccounts,
		},
		{
			TestName: "Invalid consensus engine.",
			InvalidCommandFromValid: func(command Options) Options {
				command.ConsensusEngine = "fakeConsensus"
				return command
			},
			ExpectedError: ErrInvalidConsensusEngine,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			_, err := GenerateGenesis(test.InvalidCommandFromValid(baseValidCommand))
			assert.EqualError(t, test.ExpectedError, err.Error())
		})
	}
}

func TestOptionalValues(t *testing.T) {
	baseCommand := Options{
		Network:                        "test",
		MaxNumValidators:               "5",
		UnbondingPeriod:                "5",
		AccountAddressGenesisValidator: "0xe2ac86cbae1bbbb47d157516d334e70859a1bee4",
		PrefundedAccounts: []PrefundedAccount{
			{
				AccountAddress: "0xe2ac86cbae1bbbb47d157516d334e70859a1bee4",
				Balance:        "15",
			},
		},
	}

	t.Run("Consensus engine value", func(t *testing.T) {
		baseCommand.ConsensusEngine = "tendermint"

		generatedGenesis, err := GenerateGenesis(baseCommand)
		if err != nil {
			t.Fatalf("Error: %s", err.Error())
		}

		assert.NotNil(t, generatedGenesis.Config.Tendermint)
	})

	t.Run("Smart contracts owner", func(t *testing.T) {
		customSmartContractOwner := "0xe2ac86cbae1bbbb47d157516d334e70859a1aaaa"
		baseCommand.SmartContractsOwner = customSmartContractOwner

		generatedGenesis, err := GenerateGenesis(baseCommand)
		if err != nil {
			t.Fatalf("Error: %s", err.Error())
		}

		bigaddr, _ := new(big.Int).SetString(customSmartContractOwner, 0)
		address := common.BigToAddress(bigaddr)
		expectedAlloc := core.GenesisAccount{Balance: new(big.Int).Mul(common.Big1, big.NewInt(params.Ether))}

		assert.Equal(t, generatedGenesis.Alloc[address], expectedAlloc)
	})

	t.Run("Extra data", func(t *testing.T) {
		extraDataStr := "TheExtradata"
		baseCommand.ExtraData = extraDataStr

		generatedGenesis, err := GenerateGenesis(baseCommand)
		if err != nil {
			t.Fatalf("Error: %s", err.Error())
		}

		expectedExtradata := make([]byte, 32)
		expectedExtradata = append([]byte(extraDataStr), expectedExtradata[len(extraDataStr):]...)

		assert.Equal(t, expectedExtradata, generatedGenesis.ExtraData)
	})
}

func TestMapNetwork(t *testing.T) {
	tests := []struct {
		testName string
		network  string
	}{
		{
			testName: "Empty string",
			network:  "",
		},
		{
			testName: "Invalid Network",
			network:  "fakeNetwork",
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			_, err := mapNetwork(test.network)
			if err != ErrInvalidNetwork {
				t.Fatalf("Failed to throw exception with an invalid Network value.")
			}
		})
	}
}

func TestMapConsensusEngine(t *testing.T) {
	tests := []struct {
		testName  string
		consensus string
	}{
		{
			testName:  "Empty string",
			consensus: "",
		},
		{
			testName:  "Invalid consensus engine",
			consensus: "invalidConsensus",
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			_, err := mapConsensusEngine(test.consensus)
			if err != ErrInvalidConsensusEngine {
				t.Fatalf("Failed to throw exception with an invalid Network value.")
			}
		})
	}
}

/*
func TestNew(t *testing.T) {
	// create genesis
	opts := getDefaultOpts()
	require.NotNil(t, opts)
	gen, err := genesis.New(opts)
	require.NoError(t, err)
	require.NotNil(t, gen)

	// create backend
	backend := backends.NewSimulatedBackend(gen.Alloc)
	require.NotNil(t, backend)

	// ValidatorMgr instance
	mgr, err := NewValidatorMgr(validatorMgrAddr, backend)
	require.NoError(t, err)
	require.NotNil(t, mgr)

	// MiningToken instance
	mtoken, err := token.NewMiningToken(tokenAddr, backend)
	require.NoError(t, err)
	require.NotNil(t, mtoken)

	validatorCount, err := mgr.GetValidatorCount(&bind.CallOpts{})
	require.NoError(t, err)
	require.NotNil(t, validatorCount)
	require.Equal(t, int64(len(opts.Consensus.Validators)), validatorCount.Int64())

	var deposits uint64
	for _, validator := range opts.Consensus.Validators {
		validatorAddr := common.HexToAddress(validator.Address)
		isValidator, err := mgr.IsValidator(&bind.CallOpts{}, validatorAddr)
		require.NoError(t, err)
		require.True(t, isValidator)

		storedDepositCount, err := mgr.GetDepositCount(&bind.CallOpts{From: validatorAddr})
		require.NoError(t, err)
		require.NotNil(t, storedDepositCount)
		require.Equal(t, common.Big1, storedDepositCount)

		deposit, err := mgr.GetDepositAtIndex(&bind.CallOpts{From: validatorAddr}, common.Big0)
		require.NoError(t, err)
		require.NotNil(t, deposit)
		require.Zero(t, deposit.AvailableAt.Uint64())
		require.Equal(t, new(big.Int).Mul(new(big.Int).SetUint64(opts.Consensus.BaseDeposit), new(big.Int).SetUint64(params.Ether)), deposit.Amount)

		deposits += validator.Deposit
	}

	mgrBalance, err := mtoken.BalanceOf(&bind.CallOpts{}, validatorMgrAddr)
	require.NoError(t, err)
	require.NotNil(t, mgrBalance)
	require.Equal(t, new(big.Int).Mul(new(big.Int).SetUint64(deposits), new(big.Int).SetUint64(params.Ether)), mgrBalance)
}
*/
