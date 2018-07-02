package set

import (
	"sync"
)

type memorySet struct {
	mtx  sync.RWMutex
	data map[string]interface{}
}

func NewMemorySet() Set {
	return &memorySet{
		data: make(map[string]interface{}),
	}
}

func (set *memorySet) Add(value string) error {
	set.mtx.Lock()
	defer set.mtx.Unlock()
	set.data[value] = true
	return nil
}

func (set *memorySet) Remove(value string) error {
	set.mtx.Lock()
	defer set.mtx.Unlock()
	set.data[value] = nil
	return nil
}

func (set *memorySet) Contains(value string) (bool, error) {
	set.mtx.RLock()
	defer set.mtx.RUnlock()
	return set.data[value] != nil, nil
}
