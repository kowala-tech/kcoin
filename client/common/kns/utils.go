package kns

import (
	"fmt"
	"github.com/kowala-tech/kcoin/client/params"
	"strings"

	"github.com/kowala-tech/kcoin/client/accounts/abi"
	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/contracts/bindings"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/kns"
	"github.com/kowala-tech/kcoin/client/core/vm"
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

// GetContractFromRegisteredDomains queries the EVM in order to get the contract from the
// KNS registered domains.
func GetContractFromRegisteredDomains(addr common.Address, env *vm.EVM) (string, error) {
	abiEnc, err := abi.JSON(strings.NewReader(kns.PublicResolverABI))
	if err != nil {
		return "", err
	}

	for _, domain := range params.KNSDomains {
		params, err := abiEnc.Pack("addr", NameHash(domain.FullDomain()))
		if err != nil {
			return "", err
		}

		ret, _, err := env.Call(
			vm.AccountRef(bindings.MultiSigWalletAddr),
			bindings.ProxyResolverAddr,
			params,
			600000,
			common.Big0,
		)
		if err != nil {
			return "", err
		}

		returnedAddress := common.BytesToAddress(ret)

		if addr == returnedAddress {
			return domain.FullDomain(), nil
		}
	}

	return "", fmt.Errorf("could not find domain for addr %s in the evm", addr.String())
}
