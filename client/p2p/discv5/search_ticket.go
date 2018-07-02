package discv5

import "sync"

type searchTopics struct {
	m map[Topic]searchTopic
	sync.RWMutex
}

func newSearchTopics() *searchTopics {
	return &searchTopics{m: make(map[Topic]searchTopic)}
}

func (r *searchTopics) get(key Topic) (searchTopic, bool) {
	r.RLock()
	v, ok := r.m[key]
	r.RUnlock()
	return v, ok
}

func (r *searchTopics) has(key Topic) bool {
	r.RLock()
	v, ok := r.m[key]
	r.RUnlock()
	return ok && v.foundChn != nil
}

func (r *searchTopics) set(key Topic, v searchTopic) {
	r.Lock()
	r.m[key] = v
	r.Unlock()
}

func (r *searchTopics) len() int {
	r.RLock()
	l := len(r.m)
	r.RUnlock()
	return l
}

func (r *searchTopics) delete(key Topic) {
	r.Lock()
	delete(r.m, key)
	r.Unlock()
}
