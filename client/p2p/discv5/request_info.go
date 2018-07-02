package discv5

import "sync"

type requestInfo struct {
	m map[*Node]reqInfo
	sync.RWMutex
}

func newRequestInfo() *requestInfo {
	return &requestInfo{m: make(map[*Node]reqInfo)}
}

func (r *requestInfo) get(key *Node) (reqInfo, bool) {
	r.RLock()
	v, ok := r.m[key]
	r.RUnlock()
	return v, ok
}

func (r *requestInfo) set(key *Node, req reqInfo) {
	r.Lock()
	r.m[key] = req
	r.Unlock()
}

func (r *requestInfo) delete(key *Node) {
	r.Lock()
	delete(r.m, key)
	r.Unlock()
}
