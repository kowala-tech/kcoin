// package oracle implements the network price reporting service
package oracle

import (
	"sync/atomic"

	"github.com/kowala-tech/kcoin/client/contracts/bindings/oracle"
	"github.com/kowala-tech/kcoin/client/core"
	"github.com/kowala-tech/kcoin/client/knode"
	"github.com/kowala-tech/kcoin/client/p2p"
	"github.com/kowala-tech/kcoin/client/rpc"
	"github.com/pkg/errors"
)

// @TODO (rgeraldes) - add api

const (
	// chainHeadChanSize is the size of the channel listening to ChainHeadEvent
	chainHeadChanSize = 1024
)

// Service implements a kowala price reporting deamon that posts
// local transactions containing the latest average price provided
// by a set of pre-defined exchanges.
type Service struct {
	fullNode  *knode.Kowala
	oracleMgr oracle.Manager
	reporting int32
}

// New returns a price reporting service
func New(fullNode *knode.Kowala) (*Service, error) {
	oracleMgr, err := oracle.Binding(knode.NewContractBackend(fullNode.APIBackend()), fullNode.ChainConfig().ChainID)
	if err != nil {
		return nil, errors.New("failed to create oracle manager binding")
	}

	return &Service{
		fullNode:  fullNode,
		oracleMgr: oracleMgr,
	}, nil
}

// Protocols implements node.Service, returning the P2P network protocols used
// by the oralce service (nil as it doesn't use the devp2p overlay network).
func (s *Service) Protocols() []p2p.Protocol { return nil }

// APIs implements node.Service, returning the RPC API endpoints provided by the
// stats service (nil as it doesn't provide any user callable APIs).
func (s *Service) APIs() []rpc.API { return nil }

// Start implements node.Service, starting up the monitoring and reporting daemon.
func (s *Service) Start(server *p2p.Server) error {
	/*
		s.server = server
		go s.updatePriceLoop()

		log.Info("Oracle deamon started")
		return nil
	*/
	return nil
}

// Stop implements node.Service, terminating the price.
func (s *Service) Stop() error {
	/*
		log.Info("Oracle deamon stopped")
		return nil
	*/
	return nil
}

func (s *Service) updatePriceLoop() {
	blockChain := s.fullNode.BlockChain()
	chainHeadCh := make(chan core.ChainHeadEvent, chainHeadChanSize)
	headSub := blockChain.SubscribeChainHeadEvent(chainHeadCh)
	defer headSub.Unsubscribe()

	for {
		select {
		case <-chainHeadCh:
			// update price if inside update Period
		}
	}
}

func (s *Service) StartReporting() error {
	atomic.StoreInt32(&s.reporting, 1)
	go s.updatePriceLoop()
	return nil
}

func (s *Service) StopReporting() {
	atomic.StoreInt32(&s.reporting, 0)
}

func (s *Service) IsReporting() bool {
	return atomic.LoadInt32(&s.reporting) > 0
}
