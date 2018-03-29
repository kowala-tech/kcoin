// This is a fake implementation of some structures used by the staterunner, faked temporarily until
// a proper implementation is in place.

package tests

import "github.com/kowala-tech/kcoin/core/state"

type StateTest struct{}
type StateSubtest struct {
	Fork string
}

func (t *StateTest) Subtests() []StateSubtest {
	return make([]StateSubtest, 0)
}

func (t *StateTest) Run(subtest StateSubtest, vmconfig interface{}) (*state.StateDB, error) {
	return &state.StateDB{}, nil
}
