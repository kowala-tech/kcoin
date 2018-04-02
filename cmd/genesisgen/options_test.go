package main

import (
	"testing"
	"math/big"
)

func TestItFailsWhenCreatingOptionsWithInvalidValues(t *testing.T) {
	tests := []struct{
		TestName string
		InvalidOpts *Options
		ExpectedError error
	} {
		{
			TestName: "Invalid Network",
			InvalidOpts: &Options{
				network: "fakeNet",
			},
			ExpectedError: ErrInvalidNetwork,
		},
		{
			TestName: "Empty maximum num of Validators",
			InvalidOpts: &Options{
				network: "test",
			},
			ExpectedError: ErrEmptyMaxNumValidators,
		},
		{
			TestName: "Zero max num validators",
			InvalidOpts: &Options{
				network: "test",
				maxValidators: big.NewInt(0),
			},
			ExpectedError: ErrEmptyMaxNumValidators,
		},
		{
			TestName: "Empty Unbound Period",
			InvalidOpts: &Options{
				network: "test",
				maxValidators: big.NewInt(3),
			},
			ExpectedError: ErrEmptyUnbondingPeriod,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			err := validateOptions(test.InvalidOpts)
			if err != test.ExpectedError {
				t.Fatalf("Invalid options did not return error. Expected error: %s", test.ExpectedError.Error())
			}
		})
	}
}
