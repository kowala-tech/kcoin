package core

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/stretchr/testify/assert"
)

func TestCalcStabilityFee(t *testing.T) {
	testCases := []struct {
		computeFee           *big.Int
		stabilityLevel       uint64
		txAmount             *big.Int
		expectedStabilityFee *big.Int
	}{
		{
			computeFee:           common.Big0,
			stabilityLevel:       0,
			txAmount:             common.Big0,
			expectedStabilityFee: common.Big0,
		},
		{
			computeFee:           common.Big256,
			stabilityLevel:       0,
			txAmount:             common.Big256,
			expectedStabilityFee: common.Big0,
		},
		{
			computeFee:           common.Big256,
			stabilityLevel:       4,
			txAmount:             common.Big0,
			expectedStabilityFee: common.Big256,
		},
		{
			computeFee:           common.Big3,
			stabilityLevel:       10,
			txAmount:             common.Big256,
			expectedStabilityFee: new(big.Int).SetUint64(5),
		},
		{
			computeFee:           common.Big3,
			stabilityLevel:       2,
			txAmount:             common.Big256,
			expectedStabilityFee: new(big.Int).SetUint64(4),
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("compute fee %#x stability level %d tx amount %#x expected stability fee %#x", tc.computeFee, tc.stabilityLevel, tc.txAmount, tc.expectedStabilityFee), func(t *testing.T) {
			assert.Equal(t, tc.expectedStabilityFee, calcStabilityFee(tc.computeFee, tc.stabilityLevel, tc.txAmount))
		})
	}
}
