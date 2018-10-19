package validator

import (
	"sync/atomic"

	"github.com/kowala-tech/kcoin/client/core"
	"github.com/kowala-tech/kcoin/client/event"
	"github.com/kowala-tech/kcoin/client/log"
)

type poster struct {
	*event.TypeMux
	events    chan interface{}
	isStarted *int32
}

type eventPoster interface {
	EventPost(event interface{})
}

func newPoster(eventMux *event.TypeMux) *poster {
	return &poster{eventMux, make(chan interface{}, 1000), new(int32)}
}

func (p *poster) EventPost(event interface{}) {
	p.events <- event

	if atomic.CompareAndSwapInt32(p.isStarted, 0, 1) {
		go func() {
			for event := range p.events {
				switch event.(type) {
				case core.NewMinedBlockEvent, core.NewVoteEvent, core.NewMajorityEvent:
					if err := p.Post(event); err != nil {
						log.Warn("can't post an event", "err", err, "event", event)
					}
				default:
					log.Debug("unknown validator event", "event", event)
				}
			}
		}()
	}
}
