package features

import (
	"sync"
)

type atomicCounter struct {
	n int
	sync.Mutex
}

var portCounter = newCounter()

func newCounter() *atomicCounter {
	return &atomicCounter{}
}

func (c *atomicCounter) Get() int {
	c.Lock()
	c.n++
	v := c.n
	c.Unlock()
	return v
}
