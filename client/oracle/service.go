// package oracle implements the network price reporting service
package oracle

import (
	"sync"

	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/oracle"
	"github.com/kowala-tech/kcoin/client/core"
	"github.com/kowala-tech/kcoin/client/knode"
	"github.com/kowala-tech/kcoin/client/log"
	"github.com/kowala-tech/kcoin/client/p2p"
	"github.com/kowala-tech/kcoin/client/rpc"
	"github.com/pkg/errors"
)

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

	reportingMu sync.RWMutex
	reporting   bool
	doneCh      chan struct{}
	sgx         *SGX
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
		sgx:       new(SGX),
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
	s.sgx.init()
	log.Info("Oracle deamon started")

	return nil
}

// Stop implements node.Service, terminating the price.
func (s *Service) Stop() error {
	close(s.doneCh)
	s.StopReporting()
	s.sgx.free()
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
			rawTx := s.sgx.assemblePriceTx()
		}
	}
}

func (s *Service) StartReporting() error {
	s.reportingMu.Lock()
	defer s.reportingMu.Unlock()

	coinbase, err := s.fullNode.Coinbase()
	if err != nil {
		return err
	}

	// register oracle
	isOracle, err := s.oracleMgr.IsOracle(coinbase)
	if err != nil {
		return err
	}

	if !isOracle {
		tx, err := s.oracleMgr.RegisterOracle(bind.NewKeyedTransactor(nil))
		if err != nil {
			return err
		}

		/*
			receipt, err := bind.WaitMined(context.TODO(), s.fullNode.APIBackend(), tx)
			if err != nil {
				return err
			}
		*/

		// @TODO - receipt status
	}

	go s.reportPriceLoop()
	s.reporting = true

	return nil
}

func (s *Service) StopReporting() error {
	s.reportingMu.Lock()
	defer s.reportingMu.Unlock()

	coinbase, err := s.fullNode.Coinbase()
	if err != nil {
		return err
	}

	isOracle, err := s.oracleMgr.IsOracle(coinbase)
	if err != nil {
		return nil
	}

	if isOracle {
		tx, err := s.oracleMgr.DeregisterOracle(bind.NewKeyedTransactor(nil))
		if err != nil {
			return err
		}

		/*
			receipt, err := bind.WaitMined(context.TODO(), knode.NewContractBackend(fullNode.APIBackend()), tx)
			if err != nil {
				return err
			}
		*/

		// @TODO - receipt status

		s.doneCh <- struct{}{}
		s.reporting = false
	}

	return nil
}

func (s *Service) IsReporting() bool {
	s.reportingMu.RLock()
	defer s.reportingMu.RUnlock()

	return s.reporting
}
