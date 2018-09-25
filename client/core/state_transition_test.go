package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntrisicGas(t *testing.T) {
	tests := []struct {
		name             string
		data             []byte
		contractCreation bool
		expectedGas      uint64
		expectedErr      error
	}{
		{
			name:             "simple transaction - kcoin transfer",
			data:             []byte(""),
			contractCreation: false,
			expectedGas:      21000,
			expectedErr:      nil,
		},
		{
			name:             "contract creation",
			data:             []byte("something"),
			contractCreation: true,
			expectedGas:      53612,
			expectedErr:      nil,
		},
		{
			name:             "contract call",
			data:             []byte("something2"),
			contractCreation: false,
			expectedGas:      21680,
			expectedErr:      nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gas, err := IntrinsicGas(test.data, test.contractCreation)
			require.Equal(t, test.expectedErr, err)
			assert.Equal(t, test.expectedGas, gas)
		})
	}
}
