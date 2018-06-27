package discv5

import "sync"

type radius struct {
	m map[Topic]*topicRadius
	sync.RWMutex
}

func newRadius() *radius {
	return &radius{m: make(map[Topic]*topicRadius)}
}

func (r *radius) get(key Topic) (*topicRadius, bool) {
	r.RLock()
	v, ok := r.m[key]
	r.RUnlock()
	return v, ok
}

func (r *radius) set(key Topic, v *topicRadius) {
	r.Lock()
	r.m[key] = v
	r.Unlock()
}

func (r *radius) delete(key Topic) {
	r.Lock()
	delete(r.m, key)
	r.Unlock()
}
