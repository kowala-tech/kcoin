package kusd

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/core/bloombits"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/kowala-tech/kUSD/core"
	"github.com/kowala-tech/kUSD/core/vm"
	"github.com/kowala-tech/kUSD/eth/downloader"
	"github.com/kowala-tech/kUSD/event"
	"github.com/kowala-tech/kUSD/kusddb"
	"github.com/kowala-tech/kUSD/log"
	"github.com/kowala-tech/kUSD/node"
	"github.com/kowala-tech/kUSD/p2p"
	"github.com/kowala-tech/kUSD/params"
	"github.com/kowala-tech/kUSD/rpc"
)

const (
	ChainDBFilename = "chaindata"
)

var (
	ErrLightModeNotSupported = errors.New("can't run kusd.KUSD in light sync mode, use les.LightEthereum")
	ErrValidatorNotFound     = errors.New("the genesis file must have at least one validator")
)

// @TODO(rgeraldes) - check purpose of extra data - may be related to the miner

// KUSD implements the KUSD full node service.
type KUSD struct {
	config      *Config             // service config
	chainConfig *params.ChainConfig // chain config

	chainDB kusddb.Database // blockchain db

	blockchain      *core.BlockChain
	txPool          *core.TxPool     // tx handler
	protocolManager *ProtocolManager // msg handler

	validator *Validator // consensus validator

	eventMux *event.TypeMux // events

	bloomRequests chan chan *bloombits.Retrieval // Channel receiving bloom data retrieval requests
	bloomIndexer  *core.ChainIndexer             // Bloom indexer operating during block imports

}

// New creates a new KUSD object (including the
// initialisation of the common KUSD object)
func New(ctx *node.ServiceContext, config *Config) (*KUSD, error) {
	// @TODO(rgeraldes) - refactor config validations
	if !config.SyncMode.IsValid() {
		return nil, fmt.Errorf("invalid sync mode %d", config.SyncMode)
	}

	if config.SyncMode == downloader.LightSync {
		return nil, ErrLightModeNotSupported
	}

	if genesis := config.Genesis; genesis != nil {
		if len(genesis.Validators) == 0 {
			return nil, ErrValidatorNotFound
		}

		for _, validator := range genesis.Validators {
			if validator.Power == 0 {
				return nil, fmt.Errorf("The genesis file cannot contain validators with no voting power: %v", validator)
			}
		}
	}

	// open/create the chain db
	chainDB, err := CreateDB(ctx, config, ChainDBFilename)
	if err != nil {
		return nil, err
	}

	//stopDbUpgrade := upgradeDeduplicateData(chainDb)
	// genesis setup (if necessary)
	chainConfig, genesisHash, genesisErr := core.SetupGenesisBlock(chainDB, config.Genesis)
	if _, ok := genesisErr.(*params.ConfigCompatError); genesisErr != nil && !ok {
		return nil, genesisErr
	}
	log.Info("Initialised chain configuration", "config", chainConfig)

	kusd := &KUSD{
		config:        config,
		chainConfig:   chainConfig,
		chainDB:       chainDB,
		eventMux:      ctx.EventMux,
		bloomRequests: make(chan chan *bloombits.Retrieval),
		bloomIndexer:  NewBloomIndexer(chainDb, params.BloomBitsBlocks),
		//@TODO(rgeraldes) add other elements
		/*
			accountManager: ctx.AccountManager,
			engine:         CreateConsensusEngine(ctx, &config.Ethash, chainConfig, chainDb),
			shutdownChan:   make(chan bool),
			stopDbUpgrade:  stopDbUpgrade,
			networkId:      config.NetworkId,
			gasPrice:       config.GasPrice,
			etherbase:      config.Etherbase,*/
	}

	log.Info("Initialising KUSD protocol", "versions", ProtocolVersions, "network", config.NetworkID)

	/*
			if !config.SkipBcVersionCheck {
			bcVersion := core.GetBlockChainVersion(chainDb)
			if bcVersion != core.BlockChainVersion && bcVersion != 0 {
				return nil, fmt.Errorf("Blockchain DB version mismatch (%d / %d). Run geth upgradedb.\n", bcVersion, core.BlockChainVersion)
			}
			core.WriteBlockChainVersion(chainDb, core.BlockChainVersion)
		}
	*/

	vmConfig := vm.Config{EnablePreimageRecording: config.EnablePreimageRecording}
	kusd.blockchain, err = core.NewBlockChain(chainDb, kusd.chainConfig, nil /*eth.engine*/, vmConfig)
	if err != nil {
		return nil, err
	}

	//@TODO(rgeraldes) - analyze
	/*
		// Rewind the chain in case of an incompatible config upgrade.
		if compat, ok := genesisErr.(*params.ConfigCompatError); ok {
			log.Warn("Rewinding chain to upgrade configuration", "err", compat)
			kusd.blockchain.SetHead(compat.RewindTo)
			core.WriteChainConfig(chainDb, genesisHash, chainConfig)
		}
	*/

	kusd.bloomIndexer.Start(kusd.blockchain)

	// transaction pool
	if config.TxPool.Journal != "" {
		config.TxPool.Journal = ctx.ResolvePath(config.TxPool.Journal)
	}
	kusd.txPool = core.NewTxPool(config.TxPool, kusd.chainConfig, kusd.blockchain)

	// @TODO(rgeraldes) - add protocol manager, remove consensus argument?
	// msg handler
	/*
		if kusd.protocolManager, err = NewProtocolManager(kusd.chainConfig, config.SyncMode, config.NetworkID, kusd.eventMux, kusd.txPool, nil, kusd.blockchain, chainDb); err != nil {
			return nil, err
		}
	*/

	// @TODO(rgeraldes) - validator instead of a miner
	// kusd.validator = NewValidator()
	// eth.miner = miner.New(eth, eth.chainConfig, eth.EventMux(), eth.engine)
	// eth.miner.SetExtra(makeExtraData(config.ExtraData))

	/*
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
		db.Meter("eth/db/chaindata/")
	}
	return db, nil
}

// APIs returns a collection of RPC services offered by the kusd service
func (s *KUSD) APIs() []rpc.API {
	return []rpc.API{}
}

func (s *KUSD) BlockChain() *core.BlockChain { return s.blockchain }
func (s *KUSD) TxPool() *core.TxPool         { return s.txPool }
func (s *KUSD) EventMux() *event.TypeMux     { return s.eventMux }
func (s *KUSD) ChainDb() ethdb.Database      { return s.chainDb }
func (s *KUSD) IsListening() bool            { return true } // Always listening

// @TODO(rgeraldes) - version
//func (s *KUSD) KUSDVersion() int                   { return int(s.protocolManager.SubProtocols[0].Version) }
func (s *KUSD) NetVersion() uint64                 { return s.networkID }
func (s *KUSD) Downloader() *downloader.Downloader { return s.protocolManager.downloader }

// Protocols implements node.Service, returning all the currently configured
// network protocols to start.
func (s *KUSD) Protocols() []p2p.Protocol {
	if s.lesServer == nil {
		return s.protocolManager.SubProtocols
	}
	return append(s.protocolManager.SubProtocols, s.lesServer.Protocols()...)
}

/*

// Start implements node.Service, starting all internal goroutines needed by the
// Ethereum protocol implementation.
func (s *KUSD) Start(srvr *p2p.Server) error {
	// Start the bloom bits servicing goroutines
	s.startBloomHandlers()

	// Start the RPC service
	s.netRPCService = ethapi.NewPublicNetAPI(srvr, s.NetVersion())

	// Figure out a max peers count based on the server limits
	maxPeers := srvr.MaxPeers
	if s.config.LightServ > 0 {
		maxPeers -= s.config.LightPeers
		if maxPeers < srvr.MaxPeers/2 {
			maxPeers = srvr.MaxPeers / 2
		}
	}
	// Start the networking layer and the light server if requested
	s.protocolManager.Start(maxPeers)
	if s.lesServer != nil {
		s.lesServer.Start(srvr)
	}
	return nil
}

// Stop implements node.Service, terminating all internal goroutines used by the
// Ethereum protocol.
func (s *KUSD) Stop() error {
	if s.stopDbUpgrade != nil {
		s.stopDbUpgrade()
	}
	s.bloomIndexer.Close()
	s.blockchain.Stop()
	s.protocolManager.Stop()
	if s.lesServer != nil {
		s.lesServer.Stop()
	}
	s.txPool.Stop()
	s.miner.Stop()
	s.eventMux.Stop()

	s.chainDb.Close()
	close(s.shutdownChan)

	return nil
}

*/
