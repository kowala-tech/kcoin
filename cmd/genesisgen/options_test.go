package main

import "testing"

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
			TestName: "Empty maximum num of Validators.",
			InvalidOpts: &Options{
				network: "fakeNet",
			},
			ExpectedError: ErrEmptyMaxNumValidators,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			err := validateOptions(test.InvalidOpts)
			if err != test.ExpectedError {
				t.Fatalf("Invalid options did not return error. Expected error %s", test.ExpectedError.Error())
			}
		})
	}
}
