package tendermint

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	. "math/big"
	"testing"
)

func TestBigMax(t *testing.T) {
	testCases := []struct {
		expected *Int
		left     *Int
		right    *Int
	}{
		{NewInt(0), NewInt(0), NewInt(0)},
		{NewInt(1), NewInt(1), NewInt(0)},
		{NewInt(5), NewInt(0), NewInt(5)},
		{NewInt(-1), NewInt(-1), NewInt(-5)},
		{NewInt(-5), NewInt(-10), NewInt(-5)},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("max %s,%s expected %s", tc.left, tc.right, tc.expected), func(t *testing.T) {
			assert.Equal(t, tc.expected, bigMax(tc.left, tc.right))
		})
	}
}

func TestBigMin(t *testing.T) {
	testCases := []struct {
		expected *Int
		left     *Int
		right    *Int
	}{
		{NewInt(0), NewInt(0), NewInt(0)},
		{NewInt(0), NewInt(1), NewInt(0)},
		{NewInt(0), NewInt(0), NewInt(5)},
		{NewInt(-5), NewInt(-1), NewInt(-5)},
		{NewInt(-10), NewInt(-10), NewInt(-5)},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("min %s,%s expected %s", tc.left, tc.right, tc.expected), func(t *testing.T) {
			assert.Equal(t, tc.expected, bigMin(tc.left, tc.right))
		})
	}
}
