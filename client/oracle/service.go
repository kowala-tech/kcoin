// package oracle implements the network price reporting service
package oracle

import (
	"context"
	"sync"

	"github.com/kowala-tech/kcoin/client/accounts"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/oracle"
	"github.com/kowala-tech/kcoin/client/core"
	"github.com/kowala-tech/kcoin/client/core/rawdb"
	"github.com/kowala-tech/kcoin/client/core/types"
	"github.com/kowala-tech/kcoin/client/internal/kcoinapi"
	"github.com/kowala-tech/kcoin/client/knode"
	"github.com/kowala-tech/kcoin/client/log"
	"github.com/kowala-tech/kcoin/client/oracle/scraper"
	"github.com/kowala-tech/kcoin/client/p2p"
	"github.com/kowala-tech/kcoin/client/rpc"
	"github.com/pkg/errors"
)

const (
	// chainHeadChanSize is the size of the channel listening to ChainHeadEvent
	chainHeadChanSize = 1024
)

var (
	errNotWhitelisted = errors.New("user has not been whitelisted by Kowala")
)

// Service implements a kowala price reporting deamon that posts
// local transactions containing the latest average price provided
// by a set of pre-defined exchanges.
type Service struct {
	fullNode  *knode.Kowala
	oracleMgr *oracle.Manager

	reportingMu   sync.RWMutex
	reporting     bool
	walletAccount accounts.WalletAccount
	txPoolAPI     *kcoinapi.PublicTransactionPoolAPI

	doneCh chan struct{}
}

// New returns a price reporting service
func New(fullNode *knode.Kowala) (*Service, error) {
	binding, err := oracle.Bind(knode.NewContractBackend(fullNode.APIBackend()), fullNode.ChainConfig().ChainID)
	if err != nil {
		return nil, errors.New("failed to create oracle manager binding")
	}

	oracleMgr, ok := binding.(*oracle.Manager)
	if !ok {
		return nil, err
	}

	return &Service{
		fullNode:  fullNode,
		oracleMgr: oracleMgr,
		txPoolAPI: kcoinapi.NewPublicTransactionPoolAPI(fullNode.APIBackend(), nil),
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
	close(s.doneCh)
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
		case head := <-chainHeadCh:
			// fixme: initially, Kowala controls the oracle registration
			isOracle, err := s.oracleMgr.IsOracle(s.walletAccount.Account().Address)
			if err != nil {
				return
			}
			if !isOracle {
				return
			}

			submitted, err := s.oracleMgr.HasSubmittedPrice(s.walletAccount.Account().Address)
			if err != nil {
				return
			}
			if IsUpdatePeriod(head.Block.Number()) && !submitted {
				rawTx, err := scraper.GetPrice()
				if err != nil {
					log.Error("failed to secure a price")
					continue
				}
				hash, err := s.txPoolAPI.SendRawTransaction(context.TODO(), rawTx)
				if err != nil {
					log.Error("failed to send the transaction")
					continue
				}
				receipt, err := waitMined(context.TODO(), s, hash)
				if err != nil {
					log.Error("failed to identify transaction")
					continue
				}
				if receipt.Status == types.ReceiptStatusFailed {
					log.Error("receipt status: failed")
					continue
				}
			}
		}
	}
}

func (s *Service) StartReporting() error {
	s.reportingMu.Lock()
	defer s.reportingMu.Unlock()

	walletAccount, err := s.fullNode.GetWalletAccount()
	if err != nil {
		return err
	}
	s.walletAccount = walletAccount

	isOracle, err := s.oracleMgr.IsOracle(walletAccount.Account().Address)
	if err != nil {
		return err
	}
	if !isOracle {
		return errNotWhitelisted
	}
	// fixme: initially, Kowala controls the oracle registration
	/*
		// register oracle
		if !isOracle {
			tx, err := s.oracleMgr.RegisterOracle(walletAccount)
			if err != nil {
				return err
			}
			receipt, err := waitMined(context.TODO(), s, tx.Hash())
			if err != nil {
				return err
			}
			if receipt.Status == types.ReceiptStatusFailed {
				return errors.New("receipt status: failed")
			}
		}
	*/

	scraper.Init()

	go s.reportPriceLoop()
	s.reporting = true

	return nil
}

func (s *Service) StopReporting() error {
	s.reportingMu.Lock()
	defer s.reportingMu.Unlock()

	isOracle, err := s.oracleMgr.IsOracle(s.walletAccount.Account().Address)
	if err != nil {
		return err
	}
	if !isOracle {
		return nil
	}

	// fixme: initially, Kowala controls the oracle deregistration
	/*
		tx, err := s.oracleMgr.DeregisterOracle(s.walletAccount)
		if err != nil {
			return err
		}
		receipt, err := waitMined(context.TODO(), s, tx.Hash())
		if err != nil {
			return err
		}
		if receipt.Status == types.ReceiptStatusFailed {
			return errors.New("receipt status: failed")
		}
	*/

	s.doneCh <- struct{}{}
	scraper.Free()
	s.reporting = false

	return nil
}

func (s *Service) IsReporting() bool {
	s.reportingMu.RLock()
	defer s.reportingMu.RUnlock()

	return s.reporting
}

func (s *Service) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	tx, blockHash, _, index := rawdb.ReadTransaction(s.fullNode.APIBackend().ChainDb(), txHash)
	if tx == nil {
		return nil, nil
	}
	receipts, err := s.fullNode.APIBackend().GetReceipts(ctx, blockHash)
	if err != nil {
		return nil, err
	}
	if len(receipts) <= int(index) {
		return nil, nil
	}
	return receipts[index], nil
}
