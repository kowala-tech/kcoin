package impl

import (
	"fmt"
	"math/big"
	"strings"
)

type compareFunc func(x, y *big.Int) bool

const (
	around  = "around"
	equal   = "equal"
	greater = "greater"
	less    = "less"
)

func newCompare(s string) (compareFunc, error) {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	var cmpFunc compareFunc
	switch s {
	case around:
		cmpFunc = bigIntWithinDelta(big.NewInt(3))
	case equal:
		cmpFunc = bigIntCmp(0)
	case greater:
		cmpFunc = bigIntCmp(1)
	case less:
		cmpFunc = bigIntCmp(-1)
	default:
		return nil, fmt.Errorf("unknown compare operation %q", s)
	}

	return cmpFunc, nil
}

func bigIntWithin(x, y, delta *big.Int) bool {
	// diff = abs(100 - ((x * 100) / y)
	diff := &big.Int{}
	diff.Set(x)
	diff.Mul(diff, big.NewInt(100))
	diff.Div(diff, y)
	diff.Sub(big.NewInt(100), diff)
	diff.Abs(diff)

	return diff.Cmp(delta) <= 0
}

func bigIntWithinDelta(delta *big.Int) compareFunc {
	return func(x, y *big.Int) bool {
		return bigIntWithin(x, y, delta)
	}
}

// see math.big.Cmp comments for the cmpResult values
func bigIntCmp(cmpResult int) compareFunc {
	return func(x, y *big.Int) bool {
		return x.Cmp(y) == cmpResult
	}
}
