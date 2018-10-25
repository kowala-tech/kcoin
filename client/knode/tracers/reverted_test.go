package tracers

import (
	"testing"

	"github.com/kowala-tech/kcoin/client/contracts/bindings/consensus"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/kns"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/oracle"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/ownership"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/proxy"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/sysvars"
	"github.com/kowala-tech/kcoin/client/params"
	"github.com/stretchr/testify/assert"
)

func TestSrcMapDomains(t *testing.T) {
	maps := []struct {
		expectedSrcMap []byte
		contractName   string
	}{
		{
			expectedSrcMap: []byte(proxy.UpgradeabilityProxyFactorySrcMap),
			contractName:   "Proxy Factory Contract",
		},
		{
			expectedSrcMap: []byte(kns.KNSRegistrySrcMap),
			contractName:   "Proxy KNS Registry Contract",
		},
		{
			expectedSrcMap: []byte(kns.FIFSRegistrarSrcMap),
			contractName:   "Proxy Registrar Contract",
		},
		{
			expectedSrcMap: []byte(kns.PublicResolverSrcMap),
			contractName:   "Proxy Resolver Contract",
		},
		{
			expectedSrcMap: []byte(ownership.MultiSigWalletSrcMap),
			contractName:   "Multisig Wallet Contract",
		},
		{
			expectedSrcMap: []byte(ownership.MultiSigWalletSrcMap),
			contractName:   params.KNSDomains[params.MultiSigDomain].FullDomain(),
		},
		{
			expectedSrcMap: []byte(oracle.OracleMgrSrcMap),
			contractName:   params.KNSDomains[params.OracleMgrDomain].FullDomain(),
		},
		{
			expectedSrcMap: []byte(consensus.ValidatorMgrSrcMap),
			contractName:   params.KNSDomains[params.ValidatorMgrDomain].FullDomain(),
		},
		{
			expectedSrcMap: []byte(consensus.MiningTokenSrcMap),
			contractName:   params.KNSDomains[params.MiningTokenDomain].FullDomain(),
		},
		{
			expectedSrcMap: []byte(sysvars.SystemVarsSrcMap),
			contractName:   params.KNSDomains[params.SystemVarsDomain].FullDomain(),
		},
	}

	for _, test := range maps {
		assert.Equal(t, test.expectedSrcMap, srcMapContractsData[test.contractName])
	}
}

func TestItReturnsFalseWhenSrcMapDoesNotExist(t *testing.T) {
	_, ok := srcMapContractsData["fake Name"]

	assert.False(t, ok)
}
