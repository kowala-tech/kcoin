package bindings

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kowala-tech/kcoin/client/common"
)

func TestGetContractNameByAddr(t *testing.T) {
	tests := []struct {
		addr     common.Address
		expected string
	}{
		{
			addr:     ProxyFactoryAddr,
			expected: "Proxy Factory Contract",
		},
		{
			addr:     ProxyKNSRegistryAddr,
			expected: "Proxy KNS Registry Contract",
		},
		{
			addr:     ProxyRegistrarAddr,
			expected: "Proxy Registrar Contract",
		},
		{
			addr:     ProxyResolverAddr,
			expected: "Proxy Resolver Contract",
		},
		{
			addr:     MultiSigWalletAddr,
			expected: "Multisig Wallet Contract",
		},
	}

	for _, test := range tests {
		contractName, _ := GetContractByAddr(test.addr)
		assert.Equal(t, test.expected, contractName)
	}
}
