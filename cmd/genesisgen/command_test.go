package main

import "testing"

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
			TestName: "Empty max number of validators.",
			InvalidCommand: GenerateGenesisCommand{
				network: "test",
				maxNumValidators: "",
			},
			ExpectedError: ErrEmptyMaxNumValidators,
		},
		/*{
			TestName: "Empty wallet address of the genesis validator.",
			InvalidOpts: &Options{
				network: "test",
				maxValidators: big.NewInt(3),
				unbondingPeriod: big.NewInt(0),
			},
			ExpectedError: ErrEmptyUnbondingPeriod,
		},*/
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			handler := GenerateGenesisCommandHandler{}
			err := handler.Handle(test.InvalidCommand)
			if err != test.ExpectedError {
				t.Fatalf("Invalid options did not return error. Expected error: %s", test.ExpectedError.Error())
			}
		})
	}
}

