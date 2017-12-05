package kusd

import (
	"errors"
	"fmt"
	"math/big"
	"sync"

	"github.com/kowala-tech/kUSD/accounts"
	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/core"
	"github.com/kowala-tech/kUSD/core/bloombits"
	"github.com/kowala-tech/kUSD/core/types"
	"github.com/kowala-tech/kUSD/core/vm"
	"github.com/kowala-tech/kUSD/eth/downloader"
	"github.com/kowala-tech/kUSD/event"
	"github.com/kowala-tech/kUSD/kusddb"
	"github.com/kowala-tech/kUSD/log"
	"github.com/kowala-tech/kUSD/node"
	"github.com/kowala-tech/kUSD/p2p"
	"github.com/kowala-tech/kUSD/params"
	"github.com/kowala-tech/kUSD/rpc"
	"github.com/kowala-tech/kUSD/validator"
)

// NOTE(rgeraldes) - removed references to the light service
// we can add them later as soon as we need them

const (
	ChainDBFilename        = "chaindata"
	MetricsCollectorPrefix = "kusd/db/chaindata/"
)

var (
	ErrLightModeNotSupported = errors.New("can't run kusd.KUSD in light sync mode, use ls.LightKUSD")
	ErrValidatorNotFound     = errors.New("the genesis file must have at least one validator")
)

// KUSD implements the KUSD full node service.
type KUSD struct {
	config      *Config             // service config
	chainConfig *params.ChainConfig // chain config

	chainDB kusddb.Database // blockchain db

	// handlers
	blockchain      *core.BlockChain
	txPool          *core.TxPool     // tx handler
	protocolManager *ProtocolManager // msg handler

	// consensus
	validator *validator.Validator // consensus validator
	coinbase  common.Address       // validator's KUSD account address
	gasPrice  *big.Int

	eventMux       *event.TypeMux // kusd events
	accountManager *accounts.Manager

	bloomRequests chan chan *bloombits.Retrieval // Channel receiving bloom data retrieval requests
	bloomIndexer  *core.ChainIndexer             // Bloom indexer operating during block imports

	networkID uint64

	lock sync.RWMutex // Protects the variadic fields (e.g. gas price and mUSDBase)
}

// New creates a new KUSD object (including the
// initialisation of the common KUSD object)
func New(ctx *node.ServiceContext, config *Config) (*KUSD, error) {
	// config validation
	if !config.SyncMode.IsValid() {
		return nil, fmt.Errorf("invalid sync mode %d", config.SyncMode)
	}
	if config.SyncMode == downloader.LightSync {
		return nil, ErrLightModeNotSupported
	}
	/*
		// @TODO(rgeraldes) - need to review
		if genesis := config.Genesis; genesis != nil {
			if len(genesis.Validators) == 0 {
				return nil, ErrValidatorNotFound
			}
			for _, validator := range genesis.Validators {
				if validator.Power == 0 {
					return nil, fmt.Errorf("The genesis file cannot contain validators with no voting power: %v", validator)
				}
			}
	}*/

	// setup blockchain db
	chainDB, err := CreateDB(ctx, config, ChainDBFilename)
	if err != nil {
		return nil, err
	}

	// genesis setup (if necessary)
	chainConfig, genesisHash, genesisErr := core.SetupGenesisBlock(chainDB, config.Genesis)
	if _, ok := genesisErr.(*params.ConfigCompatError); genesisErr != nil && !ok {
		return nil, genesisErr
	}
	log.Info("Initialised chain configuration", "config", chainConfig)

	kusd := &KUSD{
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

	log.Info("Initialising KUSD protocol", "versions", ProtocolVersions, "network", config.NetworkID)

	vmConfig := vm.Config{EnablePreimageRecording: config.EnablePreimageRecording}

	// @TODO(rgeraldes) - engine set to nil (analyze)
	kusd.blockchain, err = core.NewBlockChain(chainDb, kusd.chainConfig, nil, vmConfig)
	if err != nil {
		return nil, err
	}

	kusd.bloomIndexer.Start(kusd.blockchain)

	// transaction pool journal
	if config.TxPool.Journal != "" {
		config.TxPool.Journal = ctx.ResolvePath(config.TxPool.Journal)
	}
	kusd.txPool = core.NewTxPool(config.TxPool, kusd.chainConfig, kusd.blockchain)

	// msg handler
	if kusd.protocolManager, err = NewProtocolManager(kusd.chainConfig, config.SyncMode, config.NetworkID, kusd.eventMux, kusd.txPool, kusd.blockchain, chainDb); err != nil {
		return nil, err
	}

	// consensus validator
	kusd.validator = validator.New(kusd)

	/* @TODO(rgeraldes) - analyze
	eth.ApiBackend = &EthApiBackend{eth, nil}
	gpoParams := config.GPO
	if gpoParams.Default == nil {
		gpoParams.Default = config.GasPrice
	}
	eth.ApiBackend.gpo = gasprice.NewOracle(eth.ApiBackend, gpoParams)
	*/

	return kusd, nil
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

// @TODO(rgeraldes) - review (#11)
// APIs returns a collection of RPC services offered by the kusd service
func (s *KUSD) APIs() []rpc.API {
	return []rpc.API{}
}

func (s *KUSD) ResetWithGenesisBlock(gb *types.Block) {
	s.blockchain.ResetWithGenesisBlock(gb)
}

func (s *KUSD) Coinbase() (ommon.Address, error) {
	s.lock.RLock()
	coinbase := s.coinbase
	s.lock.RUnlock()

	if coinbase != (common.Address{}) {
		return coinbase, nil
	}
	if wallets := s.AccountManager().Wallets(); len(wallets) > 0 {
		if accounts := wallets[0].Accounts(); len(accounts) > 0 {
			return accounts[0].Address, nil
		}
	}
	return common.Address{}, fmt.Errorf("coinbase must be explicitly specified")
}

// set in js console via admin interface or wrapper from cli flags
func (s *KUSD) SetCoinbase(coinbase common.Address) {
	s.lock.Lock()
	s.coinbase = coinbase
	s.lock.Unlock()

	self.validator.SetCoinbase(coinbase)
}

func (s *KUSD) StartValidating(local bool) error {
	addr, err := s.Coinbase()
	if err != nil {
		log.Error("Cannot start validating without the coinbase", "err", err)
		return fmt.Errorf("coinbase missing: %v", err)
	}

	// @TODO(rgeraldes) - verify in tendermint if it makes sense to accept transactions during sync
	// it probably does not make sense.
	/*
		if local {
			// If local (CPU) mining is started, we can disable the transaction rejection
			// mechanism introduced to speed sync times. CPU mining on mainnet is ludicrous
			// so noone will ever hit this path, whereas marking sync done on CPU mining
			// will ensure that private networks work in single miner mode too.
			atomic.StoreUint32(&s.protocolManager.acceptTxs, 1)
		}
	*/

	go s.validator.Start(eb)
	return nil
}

func (s *KUSD) StopValidating()                 { s.validating.Stop() }
func (s *KUSD) IsValidating() bool              { return s.validator.IsValidating() }
func (s *KUSD) Validator() *validator.Validator { return s.validator }

func (s *KUSD) BlockChain() *core.BlockChain { return s.blockchain }
func (s *KUSD) TxPool() *core.TxPool         { return s.txPool }
func (s *KUSD) EventMux() *event.TypeMux     { return s.eventMux }
func (s *KUSD) ChainDb() kusddb.Database     { return s.chainDb }

// @TODO(rgeraldes) - analyze
func (s *KUSD) IsListening() bool                  { return true } // Always listening
func (s *KUSD) KUSDVersion() int                   { return int(s.protocolManager.SubProtocols[0].Version) }
func (s *KUSD) NetVersion() uint64                 { return s.networkID }
func (s *KUSD) Downloader() *downloader.Downloader { return s.protocolManager.downloader }

// Protocols implements node.Service, returning all the currently configured
// network protocols to start.
func (s *KUSD) Protocols() []p2p.Protocol {
	if s.lesServer == nil {
		return s.protocolManager.SubProtocols
	}
}

// Start implements node.Service, starting all internal goroutines needed by the
// KUSD protocol implementation.
func (s *KUSD) Start(srvr *p2p.Server) error {
	// Start the bloom bits servicing goroutines
	s.startBloomHandlers()

	// @TODO(rgeraldes) - grpc topic (#11)
	// Start the RPC service
	//s.netRPCService = ethapi.NewPublicNetAPI(srvr, s.NetVersion())

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
