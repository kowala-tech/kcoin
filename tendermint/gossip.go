package tendermint

import (
	"fmt"
	"time"

	"github.com/kowala-tech/kUSD/node"
	"github.com/kowala-tech/kUSD/p2p"
	"github.com/kowala-tech/kUSD/rpc"
)

type GossipService struct{ stop chan struct{} }

var _ node.Service = ((*GossipService)(nil))

func New(cfg *Config, node *node.Node) *GossipService {
	fmt.Println("New tendermint gossip network")
	return &GossipService{stop: make(chan struct{}, 0)}
}

func (tg *GossipService) Start(server *p2p.Server) error {
	fmt.Println("tendermint_gossip.Start()")
	return nil
}

func (tg *GossipService) Stop() error {
	fmt.Println("tendermint_gossip.Stop()")
	return nil
}

func (tg *GossipService) APIs() []rpc.API {
	fmt.Println("tendermint_gossip.APIs()")
	return nil
}

func (tg *GossipService) Protocols() []p2p.Protocol {
	fmt.Println("tendermint_gossip.Protocols()")
	// return nil
	return []p2p.Protocol{
		p2p.Protocol{
			Name:    "KTG",
			Version: 0,
			Length:  2,
			Run:     tg.p2pRun,
		},
	}
}

func (tg *GossipService) p2pRun(peer *p2p.Peer, rw p2p.MsgReadWriter) error {
	fmt.Printf("tendermint_gossip.p2pRun(%s)\n", peer)
	for {
		select {
		case <-tg.stop:
			return nil
		case <-time.After(100 * time.Millisecond):
			continue
		}
	}
}
