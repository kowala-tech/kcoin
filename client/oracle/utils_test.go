package oracle

import (
	"fmt"
	"math/big"
	"testing"
)

func TestIsUpdatePeriod(t *testing.T) {
	testCases := struct{
		input *big.Int
		output bool
	}
	{
		{
			input: new(big.Int).SetUint64(869),
			output: false,
		},
		{
			input: new(big.Int).SetUint64(870),
			output: true,
		},
		{
			input: new(big.Int).SetUint64(871),
			output: true,
		},
	}
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("blockNumber: %v", testCase.input), func (t *testing.T){
			require.Equal(t, testCase.output, IsUpdatePeriod(testCase.input))
		})
	}
}
