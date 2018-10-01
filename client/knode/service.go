// Package kcoin implements the Kowala protocol.
package knode

import (
	"context"
	"fmt"
	"math/big"
	"reflect"
	"sync"

	"github.com/kowala-tech/kcoin/client/accounts"
	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/common"
	engine "github.com/kowala-tech/kcoin/client/consensus"
	consensus "github.com/kowala-tech/kcoin/client/consensus/protocol"
	"github.com/kowala-tech/kcoin/client/contracts/bindings"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/oracle"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/stability"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/sysvars"
	"github.com/kowala-tech/kcoin/client/core"
	"github.com/kowala-tech/kcoin/client/core/bloombits"
	"github.com/kowala-tech/kcoin/client/core/rawdb"
	"github.com/kowala-tech/kcoin/client/core/types"
	"github.com/kowala-tech/kcoin/client/core/vm"
	"github.com/kowala-tech/kcoin/client/event"
	"github.com/kowala-tech/kcoin/client/internal/kcoinapi"
	"github.com/kowala-tech/kcoin/client/kcoindb"
	"github.com/kowala-tech/kcoin/client/knode/downloader"
	"github.com/kowala-tech/kcoin/client/knode/filters"
	"github.com/kowala-tech/kcoin/client/knode/gasprice"
	"github.com/kowala-tech/kcoin/client/knode/protocol"
	"github.com/kowala-tech/kcoin/client/log"
	"github.com/kowala-tech/kcoin/client/node"
	"github.com/kowala-tech/kcoin/client/p2p"
	"github.com/kowala-tech/kcoin/client/p2p/discv5"
	"github.com/kowala-tech/kcoin/client/params"
	"github.com/kowala-tech/kcoin/client/rpc"
	"github.com/pkg/errors"
)

// Binding constructor creates a new contract binding
type BindingConstructor func(contractBackend bind.ContractBackend, chainID *big.Int) (bindings.Binding, error)

// Kowala implements the Kowala full node service.
type Kowala struct {
	config      *Config
	chainConfig *params.ChainConfig

	shutdownChan chan bool // Channel for shutting down the service

	// Handlers
	txPool          *core.TxPool
	blockchain      *core.BlockChain
	protocolManager *ProtocolManager

	// DB interfaces
	chainDb kcoindb.Database // Block chain database

	eventMux       *event.TypeMux
	engine         engine.Engine
	accountManager *accounts.Manager

	bloomRequests chan chan *bloombits.Retrieval // Channel receiving bloom data retrieval requests
	bloomIndexer  *core.ChainIndexer             // Bloom indexer operating during block imports

	apiBackend *KowalaAPIBackend

	consensus *consensus.Consensus

	bindingFuncs []BindingConstructor // binding constructors (in dependency order)
	contracts    map[reflect.Type]bindings.Binding

	gasPrice *big.Int

	networkID     uint64
	netRPCService *kcoinapi.PublicNetAPI

	lock       sync.RWMutex // Protects the variadic fields (e.g. gas price )
	serverPool *serverPool
}

// New creates a new Kowala object (including the
// initialisation of the common Kowala object)
func New(ctx *node.ServiceContext, config *Config) (*Kowala, error) {
	if config.SyncMode == downloader.LightSync {
		return nil, errors.New("can't run kcoin.Kowala in light sync mode")
	}
	if !config.SyncMode.IsValid() {
		return nil, fmt.Errorf("invalid sync mode %d", config.SyncMode)
	}
	chainDb, err := CreateDB(ctx, config, "chaindata")
	if err != nil {
		return nil, err
	}
	chainConfig, genesisHash, genesisErr := core.SetupGenesisBlock(chainDb, config.Genesis)
	if _, ok := genesisErr.(*params.ConfigCompatError); genesisErr != nil && !ok {
		return nil, genesisErr
	}
	log.Info("Initialised chain configuration", "config", chainConfig)

	kcoin := &Kowala{
		config:         config,
		chainDb:        chainDb,
		chainConfig:    chainConfig,
		eventMux:       ctx.EventMux,
		accountManager: ctx.AccountManager,
		shutdownChan:   make(chan bool),
		networkID:      config.NetworkId,
		gasPrice:       config.GasPrice,
		bloomRequests:  make(chan chan *bloombits.Retrieval),
		bloomIndexer:   NewBloomIndexer(chainDb, params.BloomBitsBlocks),
		bindingFuncs: []BindingConstructor{
			oracle.Bind,
			sysvars.Bind,
			stability.Bind,
		},
		contracts: make(map[reflect.Type]bindings.Binding),
	}

	log.Info("Initialising Kowala protocol", "versions", protocol.Constants.Versions, "network", config.NetworkId)

	kcoin.apiBackend = &KowalaAPIBackend{kcoin, nil}

	// consensus engine
	kcoin.engine = CreateConsensusEngine(ctx, kcoin.config, kcoin.chainConfig, kcoin.chainDb)

	if !config.SkipBcVersionCheck {
		bcVersion := rawdb.ReadDatabaseVersion(chainDb)
		if bcVersion != core.BlockChainVersion && bcVersion != 0 {
			return nil, fmt.Errorf("Blockchain DB version mismatch (%d / %d). Run kcoin upgradedb.\n", bcVersion, core.BlockChainVersion)
		}
		rawdb.WriteDatabaseVersion(chainDb, core.BlockChainVersion)
	}

	vmConfig := vm.Config{EnablePreimageRecording: config.EnablePreimageRecording}
	cacheConfig := &core.CacheConfig{Disabled: config.NoPruning, TrieNodeLimit: config.TrieCache, TrieTimeLimit: config.TrieTimeout}
	kcoin.blockchain, err = core.NewBlockChain(chainDb, cacheConfig, kcoin.chainConfig, kcoin.engine, vmConfig)
	if err != nil {
		return nil, err
	}

	for _, constructor := range kcoin.bindingFuncs {
		contract, err := constructor(NewContractBackend(kcoin.apiBackend), kcoin.chainConfig.ChainID)
		if err != nil {
			return nil, err
		}
		// build and save the binding
		kind := reflect.TypeOf(contract)
		if _, exists := kcoin.contracts[kind]; exists {
			return nil, errors.New("duplicate contract")
		}
		kcoin.contracts[kind] = contract
	}

	var oracleMgr *oracle.Manager
	if err := kcoin.Contract(&oracleMgr); err != nil {
		return nil, err
	}

	var systemVars *sysvars.Vars
	if err := kcoin.Contract(&systemVars); err != nil {
		return nil, err
	}

	if err := kcoin.Contract(&kcoin.consensus); err != nil {
		return nil, err
	}

	// Rewind the chain in case of an incompatible config upgrade.
	if compat, ok := genesisErr.(*params.ConfigCompatError); ok {
		log.Warn("Rewinding chain to upgrade configuration", "err", compat)
		kcoin.blockchain.SetHead(compat.RewindTo)
		rawdb.WriteChainConfig(chainDb, genesisHash, chainConfig)
	}
	kcoin.bloomIndexer.Start(kcoin.blockchain)

	if config.TxPool.Journal != "" {
		config.TxPool.Journal = ctx.ResolvePath(config.TxPool.Journal)
	}
	kcoin.txPool = core.NewTxPool(config.TxPool, kcoin.chainConfig, kcoin.blockchain)

	gpoParams := config.GPO
	if gpoParams.Default == nil {
		gpoParams.Default = config.GasPrice
	}
	kcoin.apiBackend.gpo = gasprice.NewOracle(kcoin.apiBackend, gpoParams)

	if kcoin.protocolManager, err = NewProtocolManager(kcoin.chainConfig, config.SyncMode, config.NetworkId, kcoin.eventMux, kcoin.txPool, kcoin.engine, kcoin.blockchain, chainDb); err != nil {
		return nil, err
	}

	kcoin.serverPool = newServerPool(chainDb, kcoin.shutdownChan, new(sync.WaitGroup))

	return kcoin, nil
}

// CreateDB creates the chain database.
func CreateDB(ctx *node.ServiceContext, config *Config, name string) (kcoindb.Database, error) {
	db, err := ctx.OpenDatabase(name, config.DatabaseCache, config.DatabaseHandles)
	if err != nil {
		return nil, err
	}
	if db, ok := db.(*kcoindb.LDBDatabase); ok {
		db.Meter("kcoin/db/chaindata/")
	}
	return db, nil
}

// CreateConsensusEngine creates the required type of consensus engine instance for an Kowala service
func CreateConsensusEngine(ctx *node.ServiceContext, config *Config, chainConfig *params.ChainConfig, db kcoindb.Database) engine.Engine {
	engine := consensus.New(&params.KonsensusConfig{})
	return engine
}

// APIs returns the collection of RPC services the kowala package offers.
// NOTE, some of these services probably need to be moved to somewhere else.
func (s *Kowala) APIs() []rpc.API {
	apis := kcoinapi.GetAPIs(s.apiBackend)

	// Append any APIs exposed explicitly by the consensus engine
	apis = append(apis, s.engine.APIs(s.BlockChain())...)

	// Append all the local APIs and return
	return append(apis, []rpc.API{
		{
			Namespace: "eth",
			Version:   "1.0",
			Service:   NewPublicKowalaAPI(s),
			Public:    true,
		}, {
			Namespace: "eth",
			Version:   "1.0",
			Service:   downloader.NewPublicDownloaderAPI(s.protocolManager.downloader, s.eventMux),
			Public:    true,
		}, {
			Namespace: "mtoken",
			Version:   "1.0",
			Service:   NewPublicTokenAPI(s.accountManager, s.consensus, s.chainConfig.ChainID),
			Public:    false,
		}, {
			Namespace: "eth",
			Version:   "1.0",
			Service:   filters.NewPublicFilterAPI(s.apiBackend, false),
			Public:    true,
		}, {
			Namespace: "admin",
			Version:   "1.0",
			Service:   NewPrivateAdminAPI(s),
		}, {
			Namespace: "debug",
			Version:   "1.0",
			Service:   NewPublicDebugAPI(s),
			Public:    true,
		}, {
			Namespace: "debug",
			Version:   "1.0",
			Service:   NewPrivateDebugAPI(s.chainConfig, s),
		}, {
			Namespace: "net",
			Version:   "1.0",
			Service:   s.netRPCService,
			Public:    true,
		},
	}...)
}

func (s *Kowala) ResetWithGenesisBlock(gb *types.Block) {
	s.blockchain.ResetWithGenesisBlock(gb)
}

func (s *Kowala) AccountManager() *accounts.Manager  { return s.accountManager }
func (s *Kowala) BlockChain() *core.BlockChain       { return s.blockchain }
func (s *Kowala) TxPool() *core.TxPool               { return s.txPool }
func (s *Kowala) EventMux() *event.TypeMux           { return s.eventMux }
func (s *Kowala) Engine() engine.Engine              { return s.engine }
func (s *Kowala) ChainDb() kcoindb.Database          { return s.chainDb }
func (s *Kowala) IsListening() bool                  { return true } // Always listening
func (s *Kowala) EthVersion() int                    { return int(s.protocolManager.SubProtocols[0].Version) }
func (s *Kowala) NetVersion() uint64                 { return s.networkID }
func (s *Kowala) Downloader() *downloader.Downloader { return s.protocolManager.downloader }
func (s *Kowala) Consensus() *consensus.Consensus    { return s.consensus }
func (s *Kowala) APIBackend() *KowalaAPIBackend      { return s.apiBackend }
func (s *Kowala) ChainConfig() *params.ChainConfig   { return s.chainConfig }
func (s *Kowala) ChainID() *big.Int                  { return s.chainConfig.ChainID }

func (s *Kowala) Contract(contract interface{}) error {
	element := reflect.ValueOf(contract).Elem()
	if c, ok := s.contracts[element.Type()]; ok {
		element.Set(reflect.ValueOf(c))
		return nil
	}
	return errors.New("contract unknown")
}

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

	// Start the RPC service
	s.netRPCService = kcoinapi.NewPublicNetAPI(srvr, s.NetVersion())

	// Figure out a max peers count based on the server limits
	maxPeers := srvr.MaxPeers
	if s.config.LightServ > 0 {
		maxPeers -= s.config.LightPeers
		if maxPeers < srvr.MaxPeers/2 {
			maxPeers = srvr.MaxPeers / 2
		}
	}

	//fixme: should be removed after develop light client
	if srvr.DiscoveryV5 {
		protocolTopic := discv5.DiscoveryTopic(s.blockchain.Genesis().Hash(), protocol.ProtocolName, protocol.Kcoin1)

		go func() {
			srvr.DiscV5.RegisterTopic(protocolTopic, s.shutdownChan)
		}()

		s.serverPool.start(srvr, protocolTopic)
	}

	// Start the networking layer and the light server if requested
	s.protocolManager.Start(maxPeers)

	return nil
}

func (s *kowala) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	tx, blockHash, _, index := rawdb.ReadTransaction(s.chainDb, txHash)
	if tx == nil {
		return nil, nil
	}
	receipts, err := s.apiBackend.GetReceipts(ctx, blockHash)
	if err != nil {
		return nil, err
	}
	if len(receipts) <= int(index) {
		return nil, nil
	}
	return receipts[index], nil
}

// Stop implements node.Service, terminating all internal goroutines used by the
// Kowala protocol.
func (s *Kowala) Stop() error {
	s.bloomIndexer.Close()
	s.blockchain.Stop()
	s.protocolManager.Stop()
	s.txPool.Stop()
	s.eventMux.Stop()

	s.chainDb.Close()
	close(s.shutdownChan)

	return nil
}
