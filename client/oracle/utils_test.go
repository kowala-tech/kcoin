package oracle

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsUpdatePeriod(t *testing.T) {
	testCases := []struct {
		Input  *big.Int
		Output bool
	}{
		{
			Input:  new(big.Int).SetUint64(884),
			Output: false,
		},
		{
			Input:  new(big.Int).SetUint64(885),
			Output: true,
		},
		{
			Input:  new(big.Int).SetUint64(886),
			Output: true,
		},
	}
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("blockNumber: %v", testCase.Input), func(t *testing.T) {
			require.Equal(t, testCase.Output, IsUpdatePeriod(testCase.Input))
		})
	}
}
