package discv5

import "sync"

type requestSet struct {
	m map[Topic]struct{}
	sync.RWMutex
}

func newRequestSet() *requestSet {
	return &requestSet{m: make(map[Topic]struct{})}
}

func (r *requestSet) get(key Topic) bool {
	r.RLock()
	_, ok := r.m[key]
	r.RUnlock()
	return ok
}

func (r *requestSet) set(key Topic) {
	r.Lock()
	r.m[key] = struct{}{}
	r.Unlock()
}

func (r *requestSet) len() int {
	r.RLock()
	l := len(r.m)
	r.RUnlock()
	return l
}

func (r *requestSet) delete(key Topic) {
	r.Lock()
	delete(r.m, key)
	r.Unlock()
}
