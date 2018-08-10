package kns

import (
	"testing"
	"gotest.tools/assert"
)

func TestNameHash(t *testing.T) {
	tests := []struct{
		Name     string
		Expected string
		Input    string
	}{
		{
			Name:     "empty name",
			Expected: "0x0000000000000000000000000000000000000000000000000000000000000000",
			Input:    "",
		},
		{
			Name:     "eth",
			Expected: "0x93cdeb708b7545dc668eb9280176169d1c33cfd8ed6f04690a0bcc88a93fc4ae",
			Input:    "eth",
		},
		{
			Name:     "foo.eth",
			Expected: "0xde9b09fd7c5f901e23a3f19fecc54828e9c848539801e86591bd9801b019f84f",
			Input:    "foo.eth",
		},
	}

	for _, test := range tests {
		assert.Equal(
			t,
			test.Expected,
			NameHash(test.Input).String(),
			"failed generating namehash with %s",
			test.Name,
		)
	}
}
