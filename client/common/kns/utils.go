package kns

import (
	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/contracts/bindings"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/kns"
)

func GetAddressFromDomain(domain string, caller bind.ContractCaller) (common.Address, error) {
	resolver, err := kns.NewPublicResolverCaller(
		bindings.ProxyResolverAddr,
		caller,
	)
	if err != nil {
		return common.Address{}, err
	}

	return resolver.Addr(nil, NameHash(domain))
}
