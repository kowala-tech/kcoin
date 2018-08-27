package kns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNameHash(t *testing.T) {
	tests := []struct {
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
			Name:     "kns",
			Expected: "0xbc25ce339c62a23a50c9bdae4aba9cb6dab4cefd53d1501dcbf2eec2583d200e",
			Input:    "kowala",
		},
		{
			Name:     "foo.kowala",
			Expected: "0x8b44e9e27ccae58fe2d6023ac02a4a4ee0bdea3ad1374151da2c1f0e99e883cc",
			Input:    "foo.kowala",
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
