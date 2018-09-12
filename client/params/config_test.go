package params

import (
	"math/big"
	"reflect"
	"testing"
)

func TestCheckCompatible(t *testing.T) {
	type test struct {
		stored, new *ChainConfig
		head        uint64
		wantErr     *ConfigCompatError
	}
	tests := []test{
		{
			stored:  AllKonsensusProtocolChanges,
			new:     AllKonsensusProtocolChanges,
			wantErr: nil,
		},
		{
			stored: AllKonsensusProtocolChanges,
			new:    MainnetChainConfig,
			wantErr: &ConfigCompatError{
				What:         "Chain ID",
				StoredConfig: big.NewInt(2),
				NewConfig:    big.NewInt(1),
				RewindTo:     0,
			},
		},
	}

	for _, test := range tests {
		err := test.stored.CheckCompatible(test.new, test.head)
		if !reflect.DeepEqual(err, test.wantErr) {
			t.Errorf("error mismatch:\nstored: %v\nnew: %v\nhead: %v\nerr: %v\nwant: %v", test.stored, test.new, test.head, err, test.wantErr)
		}
	}
}
