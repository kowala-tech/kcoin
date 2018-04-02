package main

import "testing"

func TestItFailsWithAnInvalidNetwork(t *testing.T) {
	tests := []struct{
		testName string
		network string
	} {
		{
			testName: "Empty string",
			network: "",
		},
		{
			testName: "Invalid network",
			network: "fakeNetwork",
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			_, err := NewNetwork("fakeNetwork")
			if err != ErrInvalidNetwork {
				t.Fatalf("Failed to throw exception with an invalid network value.")
			}
		})
	}
}