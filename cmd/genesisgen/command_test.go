package main

import (
	"testing"
	"bytes"
	"bufio"
	"github.com/stretchr/testify/assert"
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
			{
				walletAddress: "0xe286cbae1bbbb47d157516d334e70859a1bee4ff",
				balance:       15,
			},
		},
	}

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)

	handler := GenerateGenesisCommandHandler{w:writer}

	err := handler.Handle(cmd)
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	assert.Equal(t, "", b.String())
}
