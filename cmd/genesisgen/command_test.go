package main

import (
	"testing"
)

func TestItFailsWhenRunningHandlerWithInvalidCommandValues(t *testing.T) {
	tests := []struct{
		TestName string
		InvalidCommand GenerateGenesisCommand
		ExpectedError error
	} {
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
				network: "test",
				maxNumValidators: "",
			},
			ExpectedError: ErrEmptyMaxNumValidators,
		},
		{
			TestName: "Empty unbonding period of days",
			InvalidCommand: GenerateGenesisCommand{
				network: "test",
				maxNumValidators: "5",
				unbondingPeriod: "",
			},
			ExpectedError: ErrEmptyUnbondingPeriod,
		},
		{
			TestName: "Empty wallet address of genesis validator",
			InvalidCommand: GenerateGenesisCommand{
				network: "test",
				maxNumValidators: "5",
				unbondingPeriod: "5",
				walletAddressGenesisValidator: "",
			},
			ExpectedError: ErrEmptyWalletAddressValidator,
		},
		{
			TestName: "Invalid wallet address less than 20 bytes with Hex prefix",
			InvalidCommand: GenerateGenesisCommand{
				network: "test",
				maxNumValidators: "5",
				unbondingPeriod: "5",
				walletAddressGenesisValidator: "0xe2ac86cbae1bbbb47d157516d334e70859a1be",
			},
			ExpectedError: ErrInvalidWalletAddressValidator,
		},
		{
			TestName: "Invalid wallet address less than 20 bytes without Hex prefix",
			InvalidCommand: GenerateGenesisCommand{
				network: "test",
				maxNumValidators: "5",
				unbondingPeriod: "5",
				walletAddressGenesisValidator: "e2ac86cbae1bbbb47d157516d334e70859a1be",
			},
			ExpectedError: ErrInvalidWalletAddressValidator,
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

