package discv5

import (
	"sync"
	"github.com/kowala-tech/kcoin/client/log"
)

type nodeTickets struct {
	m map[*Node]*ticket
	sync.RWMutex
}

func newNodeTickets() *nodeTickets {
	return &nodeTickets{m: make(map[*Node]*ticket)}
}

func (r *nodeTickets) get(key *Node) (*ticket, bool) {
	r.RLock()
	v, ok := r.m[key]
	r.RUnlock()

	switch ok {
	case true:
		log.Trace("Retrieving node ticket", "node", key.ID, "serial", v.serial)
	case false:
		log.Trace("Retrieving node ticket", "node", key.ID, "serial", nil)
	}

	return v, ok
}

func (r *nodeTickets) set(key *Node, v *ticket) {
	r.Lock()
	r.m[key] = v
	r.Unlock()
}

func (r *nodeTickets) delete(key *Node) {
	r.Lock()
	delete(r.m, key)
	r.Unlock()
}
