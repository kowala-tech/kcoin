package features

import "sync/atomic"

type atomicCounter int32

var portCounter = new(atomicCounter)

func (c *atomicCounter) Get() int32 {
	return atomic.AddInt32((*int32)(c), 1)
}
