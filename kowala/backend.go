package kowala

import (
	"errors"
	"fmt"
	"math/big"
	"sync"
	"sync/atomic"

	"github.com/kowala-tech/kUSD/accounts"
	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/consensus/downloader"
	"github.com/kowala-tech/kUSD/consensus/validator"
	"github.com/kowala-tech/kUSD/core"
	"github.com/kowala-tech/kUSD/core/bloombits"
	"github.com/kowala-tech/kUSD/core/types"
	"github.com/kowala-tech/kUSD/core/vm"
	"github.com/kowala-tech/kUSD/event"
	"github.com/kowala-tech/kUSD/kusddb"
	"github.com/kowala-tech/kUSD/log"
	"github.com/kowala-tech/kUSD/node"
	"github.com/kowala-tech/kUSD/p2p"
	"github.com/kowala-tech/kUSD/params"
	"github.com/kowala-tech/kUSD/rpc"
)

const (
	ChainDBFilename        = "chaindata"
	MetricsCollectorPrefix = "kusd/db/chaindata/"
)

var (
	ErrLightModeNotSupported = errors.New("can't run kusd.Kowala in light sync mode, use ls.LightKowala")
	//ErrValidatorNotFound     = errors.New("the genesis file must have at least one validator")
)

// Kowala implements the Kowala full node service.
type Kowala struct {
	networkID   uint64
	config      *Config
	chainConfig *params.ChainConfig

	// handlers
	blockchain      *core.BlockChain
	txPool          *core.TxPool
	protocolManager *ProtocolManager

	// db
	chainDB kusddb.Database

	// consensus
	lock      sync.RWMutex // Protects the variadic fields (e.g. gas price and mUSDBase)
	coinbase  common.Address
	gasPrice  *big.Int
	validator *validator.Validator

	// events
	eventMux *event.TypeMux // events

	accountManager *accounts.Manager

	// bloom
	bloomRequests chan chan *bloombits.Retrieval // Channel receiving bloom data retrieval requests
	bloomIndexer  *core.ChainIndexer             // Bloom indexer operating during block imports
}

// New creates a new kowala object (including the
// initialisation of the common kowala object)
func New(ctx *node.ServiceContext, config *Config) (*Kowala, error) {
	if !config.SyncMode.IsValid() {
		return nil, fmt.Errorf("invalid sync mode %d", config.SyncMode)
	}
	if config.SyncMode == downloader.LightSync {
		return nil, ErrLightModeNotSupported
	}

	chainDB, err := CreateDB(ctx, config, ChainDBFilename)
	if err != nil {
		return nil, err
	}

	chainConfig, genesisHash, genesisErr := core.SetupGenesisBlock(chainDB, config.Genesis)
	if _, ok := genesisErr.(*params.ConfigCompatError); genesisErr != nil && !ok {
		return nil, genesisErr
	}
	log.Info("Initialised chain configuration", "config", chainConfig)

	kowala := &Kowala{
		config:         config,
		chainConfig:    chainConfig,
		chainDB:        chainDB,
		eventMux:       ctx.EventMux,
		accountManager: ctx.AccountManager,
		networkID:      config.NetworkID,
		gasPrice:       config.GasPrice,
		coinbase:       config.coinbase,
		bloomRequests:  make(chan chan *bloombits.Retrieval),
		bloomIndexer:   NewBloomIndexer(chainDb, params.BloomBitsBlocks),
	}

	log.Info("Initialising kowala protocol", "versions", ProtocolVersions, "network", config.NetworkID)

	vmConfig := vm.Config{EnablePreimageRecording: config.EnablePreimageRecording}
	kowala.blockchain, err = core.NewBlockChain(chainDb, kowala.chainConfig, vmConfig)
	if err != nil {
		return nil, err
	}

	kowala.bloomIndexer.Start(kowala.blockchain)

	if config.TxPool.Journal != "" {
		config.TxPool.Journal = ctx.ResolvePath(config.TxPool.Journal)
	}
	kowala.txPool = core.NewTxPool(config.TxPool, kowala.chainConfig, kowala.blockchain)

	if kowala.protocolManager, err = NewProtocolManager(kowala.chainConfig, config.SyncMode, config.NetworkID, kowala.eventMux, kowala.txPool, kowala.blockchain, chainDb); err != nil {
		return nil, err
	}

	kowala.validator = validator.New(kowala, kowala.chainConfig, kowala.EventMux())

	return kowala, nil
}

// CreateDB creates the chain database.
func CreateDB(ctx *node.ServiceContext, config *Config, name string) (kusddb.Database, error) {
	db, err := ctx.OpenDatabase(name, config.DatabaseCache, config.DatabaseHandles)
	if err != nil {
		return nil, err
	}
	if db, ok := db.(*kusddb.LDBDatabase); ok {
		db.Meter(MetricsCollectorPrefix)
	}
	return db, nil
}

// APIs returns a collection of RPC services offered by the kowala service
func (s *Kowala) APIs() []rpc.API {
	return []rpc.API{}
}

func (s *Kowala) ResetWithGenesisBlock(gb *types.Block) {
	s.blockchain.ResetWithGenesisBlock(gb)
}

func (s *Kowala) Coinbase() (common.Address, error) {
	s.lock.RLock()
	coinbase := s.coinbase
	s.lock.RUnlock()

	if coinbase != (common.Address{}) {
		return coinbase, nil
	}
	if wallets := s.AccountManager().Wallets(); len(wallets) > 0 {
		if accounts := wallets[0].Accounts(); len(accounts) > 0 {
			coinbase = accounts[0].Address

			s.lock.Lock()
			s.coinbase = coinbase
			s.lock.Unlock()

			log.Info("Coinbase automatically configured", "address", coinbase)
			return coinbase, nil
		}
	}
	return common.Address{}, fmt.Errorf("coinbase must be explicitly specified")
}

// set in js console via admin interface or wrapper from cli flags
func (s *Kowala) SetCoinbase(coinbase common.Address) {
	s.lock.Lock()
	s.coinbase = coinbase
	s.lock.Unlock()

	self.validator.SetCoinbase(coinbase)
}

func (s *Kowala) StartValidating(local bool) error {
	cb, err := s.Coinbase()
	if err != nil {
		log.Error("Cannot start validation without the coinbase", "err", err)
		return fmt.Errorf("coinbase missing: %v", err)
	}

	// disable the transaction rejection mechanism introduced to speed sync times.
	atomic.StoreUint32(&s.protocolManager.acceptTxs, 1)

	go s.validator.Start(cb)
	return nil
}

func (s *Kowala) StopValidating()                 { s.validating.Stop() }
func (s *Kowala) IsValidating() bool              { return s.validator.IsValidating() }
func (s *Kowala) Validator() *validator.Validator { return s.validator }

func (s *Kowala) AccountManager() *accounts.Manager  { return s.accountManager }
func (s *Kowala) BlockChain() *core.BlockChain       { return s.blockchain }
func (s *Kowala) TxPool() *core.TxPool               { return s.txPool }
func (s *Kowala) EventMux() *event.TypeMux           { return s.eventMux }
func (s *Kowala) ChainDb() kusddb.Database           { return s.chainDb }
func (s *Kowala) IsListening() bool                  { return true } // Always listening
func (s *Kowala) KowalaVersion() int                 { return int(s.protocolManager.SubProtocols[0].Version) }
func (s *Kowala) NetworkVersion() uint64             { return s.networkId }
func (s *Kowala) Downloader() *downloader.Downloader { return s.protocolManager.downloader }

// Protocols implements node.Service, returning all the currently configured
// network protocols to start.
func (s *Kowala) Protocols() []p2p.Protocol {
	return s.protocolManager.SubProtocols
}

// Start implements node.Service, starting all internal goroutines needed by the
// Kowala protocol implementation.
func (s *Kowala) Start(srvr *p2p.Server) error {
	// Start the bloom bits servicing goroutines
	s.startBloomHandlers()

	// Figure out a max peers count based on the server limits
	maxPeers := srvr.MaxPeers

	// Start the networking layer and the light server if requested
	s.protocolManager.Start(maxPeers)

	return nil
}

// Stop implements node.Service, terminating all internal goroutines used by the
// Ethereum protocol.
func (s *Ethereum) Stop() error {
	s.bloomIndexer.Close()
	s.blockchain.Stop()
	s.protocolManager.Stop()
	s.txPool.Stop()
	s.validator.Stop()
	s.eventMux.Stop()
	s.chainDb.Close()
	return nil
}

func RegisterService(node *node.Node, cfg *kowala.Config) err {
	return node.Register(func(ctx *node.ServiceContext) (node.Service, error) {
		return Kowala.New(ctx, cfg)
	})
}
