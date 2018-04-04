package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"github.com/kowala-tech/kcoin/core"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
	"math/big"
	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/params"
)

func TestItFailsWhenRunningHandlerWithInvalidCommandValues(t *testing.T) {
	tests := []struct {
		TestName       string
		InvalidCommand GenerateGenesisCommand
		ExpectedError  error
	}{
		{
			TestName: "Invalid Network",
			InvalidCommand: GenerateGenesisCommand{
				network: "fakeNetwork",
			},
			ExpectedError: ErrInvalidNetwork,
		},
		{
			TestName: "Empty max number of validators",
			InvalidCommand: GenerateGenesisCommand{
				network:          "test",
				maxNumValidators: "",
			},
			ExpectedError: ErrEmptyMaxNumValidators,
		},
		{
			TestName: "Empty unbonding period of days",
			InvalidCommand: GenerateGenesisCommand{
				network:          "test",
				maxNumValidators: "5",
				unbondingPeriod:  "",
			},
			ExpectedError: ErrEmptyUnbondingPeriod,
		},
		{
			TestName: "Empty wallet address of genesis validator",
			InvalidCommand: GenerateGenesisCommand{
				network:                       "test",
				maxNumValidators:              "5",
				unbondingPeriod:               "5",
				walletAddressGenesisValidator: "",
			},
			ExpectedError: ErrEmptyWalletAddressValidator,
		},
		{
			TestName: "Invalid wallet address less than 20 bytes with Hex prefix",
			InvalidCommand: GenerateGenesisCommand{
				network:                       "test",
				maxNumValidators:              "5",
				unbondingPeriod:               "5",
				walletAddressGenesisValidator: "0xe2ac86cbae1bbbb47d157516d334e70859a1be",
			},
			ExpectedError: ErrInvalidWalletAddressValidator,
		},
		{
			TestName: "Invalid wallet address less than 20 bytes without Hex prefix",
			InvalidCommand: GenerateGenesisCommand{
				network:                       "test",
				maxNumValidators:              "5",
				unbondingPeriod:               "5",
				walletAddressGenesisValidator: "e2ac86cbae1bbbb47d157516d334e70859a1be",
			},
			ExpectedError: ErrInvalidWalletAddressValidator,
		},
		{
			TestName: "Empty prefunded accounts",
			InvalidCommand: GenerateGenesisCommand{
				network:                       "test",
				maxNumValidators:              "5",
				unbondingPeriod:               "5",
				walletAddressGenesisValidator: "0xe2ac86cbae1bbbb47d157516d334e70859a1bee4",
				prefundedAccounts:             []PrefundedAccount{},
			},
			ExpectedError: ErrEmptyPrefundedAccounts,
		},
		{
			TestName: "Prefunded accounts does not include validator address",
			InvalidCommand: GenerateGenesisCommand{
				network:                       "test",
				maxNumValidators:              "5",
				unbondingPeriod:               "5",
				walletAddressGenesisValidator: "0xe2ac86cbae1bbbb47d157516d334e70859a1bee4",
				prefundedAccounts: []PrefundedAccount{
					{
						walletAddress: "0xaaaaaacbae1bbbb47d157516d334e70859a1bee4",
						balance:       15,
					},
				},
			},
			ExpectedError: ErrWalletAddressValidatorNotInPrefundedAccounts,
		},
		{
			TestName: "Prefunded accounts has invalid account.",
			InvalidCommand: GenerateGenesisCommand{
				network:                       "test",
				maxNumValidators:              "5",
				unbondingPeriod:               "5",
				walletAddressGenesisValidator: "0xe2ac86cbae1bbbb47d157516d334e70859a1bee4",
				prefundedAccounts: []PrefundedAccount{
					{
						walletAddress: "0xe2ac86cbae1bbbb47d157516d334e70859a1bee4",
						balance:       15,
					},
					{
						walletAddress: "0xe286cbae1bbbb47d157516d334e70859a1bee4",
						balance:       15,
					},
				},
			},
			ExpectedError: ErrInvalidAddressInPrefundedAccounts,
		},
		{
			TestName: "Invalid consensus engine.",
			InvalidCommand: GenerateGenesisCommand{
				network:                       "test",
				maxNumValidators:              "5",
				unbondingPeriod:               "5",
				walletAddressGenesisValidator: "0xe2ac86cbae1bbbb47d157516d334e70859a1bee4",
				prefundedAccounts: []PrefundedAccount{
					{
						walletAddress: "0xe2ac86cbae1bbbb47d157516d334e70859a1bee4",
						balance:       15,
					},
				},
				consensusEngine: "fakeConsensus",
			},
			ExpectedError: ErrInvalidConsensusEngine,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			handler := GenerateGenesisCommandHandler{}
			err := handler.Handle(test.InvalidCommand)
			if err != test.ExpectedError {
				t.Fatalf(
					"Invalid options did not return error. Expected error: %s, received error: %s",
					test.ExpectedError.Error(),
					err.Error(),
				)
			}
		})
	}
}

func TestItWritesTheGeneratedFileToAWriter(t *testing.T) {
	cmd := GenerateGenesisCommand{
		network:                       "test",
		maxNumValidators:              "5",
		unbondingPeriod:               "5",
		walletAddressGenesisValidator: "0xe2ac86cbae1bbbb47d157516d334e70859a1bee4",
		prefundedAccounts: []PrefundedAccount{
			{
				walletAddress: "0xe2ac86cbae1bbbb47d157516d334e70859a1bee4",
				balance:       15,
			},
		},
	}

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)

	handler := GenerateGenesisCommandHandler{w: writer}

	err := handler.Handle(cmd)
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	fileName := "testfiles/testnet_default.json"
	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		t.Fatalf("Failed to read file %s", fileName)
	}

	var expectedGenesis = new(core.Genesis)
	err = json.Unmarshal(contents, expectedGenesis)
	if err != nil {
		t.Fatalf("Error unmarshalling json genesis with error: %s", err.Error())
	}

	var generatedGenesis = new(core.Genesis)
	err = json.Unmarshal(b.Bytes(), generatedGenesis)

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
}

func TestOptionalValues(t *testing.T) {
	t.Run("Consensus engine value", func(t *testing.T) {
		cmd := GenerateGenesisCommand{
			network:                       "test",
			maxNumValidators:              "5",
			unbondingPeriod:               "5",
			walletAddressGenesisValidator: "0xe2ac86cbae1bbbb47d157516d334e70859a1bee4",
			prefundedAccounts: []PrefundedAccount{
				{
					walletAddress: "0xe2ac86cbae1bbbb47d157516d334e70859a1bee4",
					balance:       15,
				},
			},
			consensusEngine: "tendermint",
		}

		var b bytes.Buffer
		handler := GenerateGenesisCommandHandler{w: &b}

		err := handler.Handle(cmd)
		if err != nil {
			t.Fatalf("Error: %s", err.Error())
		}

		var generatedGenesis = new(core.Genesis)
		err = json.Unmarshal(b.Bytes(), generatedGenesis)
		if err != nil {
			t.Fatal("Error unmarshaling genesis.")
		}

		assert.NotNil(t, generatedGenesis.Config.Tendermint)
	})

	t.Run("Smart contracts owner", func(t *testing.T) {
		cmd := GenerateGenesisCommand{
			network:                       "test",
			maxNumValidators:              "5",
			unbondingPeriod:               "5",
			walletAddressGenesisValidator: "0xe2ac86cbae1bbbb47d157516d334e70859a1bee4",
			prefundedAccounts: []PrefundedAccount{
				{
					walletAddress: "0xe2ac86cbae1bbbb47d157516d334e70859a1bee4",
					balance:       15,
				},
			},
			smartContractsOwner: "0xe2ac86cbae1bbbb47d157516d334e70859a1aaaa",
		}

		var b bytes.Buffer
		handler := GenerateGenesisCommandHandler{w: &b}

		err := handler.Handle(cmd)
		if err != nil {
			t.Fatalf("Error: %s", err.Error())
		}

		var generatedGenesis = new(core.Genesis)
		err = json.Unmarshal(b.Bytes(), generatedGenesis)
		if err != nil {
			t.Fatal("Error unmarshaling genesis.")
		}

		bigaddr, _ := new(big.Int).SetString("0xe2ac86cbae1bbbb47d157516d334e70859a1aaaa", 0)
		address := common.BigToAddress(bigaddr)

		expectedAlloc := core.GenesisAccount{Balance: new(big.Int).Mul(common.Big1, big.NewInt(params.Ether))}

		assert.Equal(t, generatedGenesis.Alloc[address], expectedAlloc)
	})
}
