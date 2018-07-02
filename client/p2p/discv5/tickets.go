package discv5

import "sync"

type tickets struct {
	m map[Topic]*topicTickets
	sync.RWMutex
}

func newTickets() *tickets {
	return &tickets{m: make(map[Topic]*topicTickets)}
}

func (r *tickets) get(key Topic) (*topicTickets, bool) {
	r.RLock()
	v, ok := r.m[key]
	r.RUnlock()
	return v, ok
}

func (r *tickets) set(key Topic, v *topicTickets) {
	r.Lock()
	r.m[key] = v
	r.Unlock()
}

func (r *tickets) len() int {
	r.RLock()
	l := len(r.m)
	r.RUnlock()
	return l
}

func (r *tickets) delete(key Topic) {
	r.Lock()
	delete(r.m, key)
	r.Unlock()
}
