package state

import (
	"reflect"

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

type DomainState struct {
	*StateDB
	DomainResolver
}

func NewDomainState(state *StateDB, resolver DomainResolver) *DomainState {
	if resolver == nil {
		resolver = ResolverFunc(hardcodedResolver)
	}

	return &DomainState{
		StateDB:        state,
		DomainResolver: resolver,
	}
}

func (ds *DomainState) SetState(value interface{}) error {
	addr, err := ds.Resolve(getTypeName(value))
	if err != nil {
		return errors.Wrap(err, "failed to resolve domain")
	}

	stateObject := ds.GetOrNewStateObject(addr)
	if stateObject != nil {
		if err := Map(stateObject, value); err != nil {
			return errors.Wrap(err, "failed to map value to state object")
		}
	}
}

func getTypeName(val interface{}) string {
	if t := reflect.TypeOf(val); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}
