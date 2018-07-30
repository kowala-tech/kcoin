package konsensus

import (
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/pkg/errors"
)

type DomainResolver interface {
	Resolve(domain string) (common.Address, error)
}

type ResolverFunc func(string) (common.Address, error)

func (fn ResolverFunc) Resolve(domain string) (common.Address, error) {
	return fn(domain)
}

func hardcodedResolver(domain string) (common.Address, error) {
	switch domain {
	case "systemvars":
		// @TODO (rgeraldes) - hardcoded address
	case "oraclemgr":
		// @TODO (rgeraldes) - hardcoded address
	}
	return common.Address{}, errors.New("invalid domain")
}
