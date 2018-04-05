package kcoin

import "testing"

func TestItFailsWithAnInvalidNetwork(t *testing.T) {
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
			_, err := NewNetwork(test.network)
			if err != ErrInvalidNetwork {
				t.Fatalf("Failed to throw exception with an invalid Network value.")
			}
		})
	}
}

func TestItFailsWithInvalidConsensusEngine(t *testing.T) {
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
			_, err := NewConsensusEngine(test.consensus)
			if err != ErrInvalidConsensusEngine {
				t.Fatalf("Failed to throw exception with an invalid Network value.")
			}
		})
	}
}
