// package oracle implements the network price reporting service
package oracle

import (
	"sync/atomic"

	"github.com/kowala-tech/kcoin/client/contracts/bindings/oracle"
	"github.com/kowala-tech/kcoin/client/core"
	"github.com/kowala-tech/kcoin/client/knode"
	"github.com/kowala-tech/kcoin/client/log"
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
	doneCh    chan struct{}
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
// oracle service
func (s *Service) APIs() []rpc.API {
	return []rpc.API{
		{
			Namespace: "oracle",
			Version:   "1.0",
			Service:   NewPrivateOracleAPI(s),
			Public:    false,
		},
	}
}

// Start implements node.Service, starting up the monitoring and reporting daemon.
func (s *Service) Start(server *p2p.Server) error {
	s.doneCh = make(chan struct{})
	log.Info("Oracle deamon started")

	return nil
}

// Stop implements node.Service, terminating the price.
func (s *Service) Stop() error {
	s.StopReporting()
	log.Info("Oracle deamon stopped")

	return nil
}

func (s *Service) reportPriceLoop() {
	blockChain := s.fullNode.BlockChain()
	chainHeadCh := make(chan core.ChainHeadEvent, chainHeadChanSize)
	headSub := blockChain.SubscribeChainHeadEvent(chainHeadCh)
	defer headSub.Unsubscribe()

	for {
		select {
		case <-s.doneCh:
			return
		case <-chainHeadCh:
			// head.Block.Number()
			// update price if inside update Period
		}
	}
}

func (s *Service) StartReporting() error {
	atomic.StoreInt32(&s.reporting, 1)
	go s.reportPriceLoop()
	return nil
}

func (s *Service) StopReporting() {
	atomic.StoreInt32(&s.reporting, 0)
	s.doneCh <- struct{}{}
}

func (s *Service) IsReporting() bool {
	return atomic.LoadInt32(&s.reporting) > 0
	close(s.doneCh)
	return true
}
