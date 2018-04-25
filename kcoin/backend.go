// Package kcoin implements the Kowala protocol.
package kcoin

import (
	"errors"
	"fmt"
	"math/big"
	"runtime"
	"sync"
	"sync/atomic"

	"github.com/kowala-tech/kcoin/accounts"
	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/common/hexutil"
	"github.com/kowala-tech/kcoin/consensus"
	"github.com/kowala-tech/kcoin/consensus/tendermint"
	"github.com/kowala-tech/kcoin/contracts/network"
	"github.com/kowala-tech/kcoin/core"
	"github.com/kowala-tech/kcoin/core/bloombits"
	"github.com/kowala-tech/kcoin/core/types"
	"github.com/kowala-tech/kcoin/core/vm"
	"github.com/kowala-tech/kcoin/event"
	"github.com/kowala-tech/kcoin/internal/kcoinapi"
	"github.com/kowala-tech/kcoin/kcoin/downloader"
	"github.com/kowala-tech/kcoin/kcoin/filters"
	"github.com/kowala-tech/kcoin/kcoin/gasprice"
	"github.com/kowala-tech/kcoin/kcoin/validator"
	"github.com/kowala-tech/kcoin/kcoindb"
	"github.com/kowala-tech/kcoin/log"
	"github.com/kowala-tech/kcoin/node"
	"github.com/kowala-tech/kcoin/p2p"
	"github.com/kowala-tech/kcoin/params"
	"github.com/kowala-tech/kcoin/rlp"
	"github.com/kowala-tech/kcoin/rpc"
	"github.com/kowala-tech/kcoin/kcoin/wal/wal"
)

// @TODO(rgeraldes) - we may need to enable transaction syncing right from the beginning (in StartValidating - check previous version)

// Kowala implements the Kowala full node service.
type Kowala struct {
	config      *Config
	chainConfig *params.ChainConfig
	// Channel for shutting down the service
	shutdownChan chan bool // Channel for shutting down the service

	// Handlers
	txPool          *core.TxPool
	blockchain      *core.BlockChain
	protocolManager *ProtocolManager
	// DB interfaces
	chainDb kcoindb.Database // Block chain database

	eventMux       *event.TypeMux
	engine         consensus.Engine
	accountManager *accounts.Manager

	bloomRequests chan chan *bloombits.Retrieval // Channel receiving bloom data retrieval requests
	bloomIndexer  *core.ChainIndexer             // Bloom indexer operating during block imports

	ApiBackend *KowalaApiBackend

	validator validator.Validator // consensus validator
	election  network.Election    // consensus election
	gasPrice  *big.Int
	coinbase  common.Address
	deposit   uint64

	networkId     uint64
	netRPCService *kcoinapi.PublicNetAPI

	lock sync.RWMutex // Protects the variadic fields (e.g. gas price and coinbase)
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
		engine:         CreateConsensusEngine(ctx, config, chainConfig, chainDb),
		shutdownChan:   make(chan bool),
		networkId:      config.NetworkId,
		gasPrice:       config.GasPrice,
		coinbase:       config.Coinbase,
		deposit:        config.Deposit,
		bloomRequests:  make(chan chan *bloombits.Retrieval),
		bloomIndexer:   NewBloomIndexer(chainDb, params.BloomBitsBlocks),
	}

	log.Info("Initialising Kowala protocol", "versions", ProtocolVersions, "network", config.NetworkId)

	if !config.SkipBcVersionCheck {
		bcVersion := core.GetBlockChainVersion(chainDb)
		if bcVersion != core.BlockChainVersion && bcVersion != 0 {
			return nil, fmt.Errorf("Blockchain DB version mismatch (%d / %d). Run kcoin upgradedb.\n", bcVersion, core.BlockChainVersion)
		}
		core.WriteBlockChainVersion(chainDb, core.BlockChainVersion)
	}

	vmConfig := vm.Config{EnablePreimageRecording: config.EnablePreimageRecording}
	kcoin.blockchain, err = core.NewBlockChain(chainDb, kcoin.chainConfig, kcoin.engine, vmConfig)
	if err != nil {
		return nil, err
	}
	// Rewind the chain in case of an incompatible config upgrade.
	if compat, ok := genesisErr.(*params.ConfigCompatError); ok {
		log.Warn("Rewinding chain to upgrade configuration", "err", compat)
		kcoin.blockchain.SetHead(compat.RewindTo)
		core.WriteChainConfig(chainDb, genesisHash, chainConfig)
	}
	kcoin.bloomIndexer.Start(kcoin.blockchain)

	if config.TxPool.Journal != "" {
		config.TxPool.Journal = ctx.ResolvePath(config.TxPool.Journal)
	}
	kcoin.txPool = core.NewTxPool(config.TxPool, kcoin.chainConfig, kcoin.blockchain)

	kcoin.ApiBackend = &KowalaApiBackend{kcoin, nil}
	gpoParams := config.GPO
	if gpoParams.Default == nil {
		gpoParams.Default = config.GasPrice
	}
	fmt.Println("Chain CONFIG", *config.Genesis.Config, config.GasPrice.String(), config.Genesis.Number, config.Genesis.GasLimit, config.Genesis.GasUsed)
	fmt.Println("ORACLE CONFIG", gpoParams, config.Genesis.Alloc)
	kcoin.ApiBackend.gpo = gasprice.NewOracle(kcoin.ApiBackend, gpoParams)

	fmt.Println("Chain Config", chainConfig)
	// consensus validator
	election, err := network.NewElection(NewContractBackend(kcoin.ApiBackend), chainConfig.ChainID)
	if err != nil {
		log.Crit("Failed to load the network contract", "err", err)
	}
	kcoin.election = election

	/*
	walletAccount, err := kcoin.getWalletAccount()
	if err != nil {
		log.Warn("failed to get wallet account", "err", err)
	}
	*/

	userWal, err := wal.New(ctx.ResolvePath("wal"))
	if err != nil {
		log.Warn("failed to get WAL", "err", err)
	}

	kcoin.validator = validator.New(kcoin, kcoin.election, kcoin.chainConfig, kcoin.EventMux(), kcoin.engine, vmConfig, userWal)
	kcoin.validator.SetExtra(makeExtraData(config.ExtraData))

	if kcoin.protocolManager, err = NewProtocolManager(kcoin.chainConfig, config.SyncMode, config.NetworkId, kcoin.eventMux, kcoin.txPool, kcoin.engine, kcoin.blockchain, chainDb, kcoin.validator); err != nil {
		return nil, err
	}

	return kcoin, nil
}

func makeExtraData(extra []byte) []byte {
	if len(extra) == 0 {
		// create default extradata
		extra, _ = rlp.EncodeToBytes([]interface{}{
			uint(params.VersionMajor<<16 | params.VersionMinor<<8 | params.VersionPatch),
			"kcoin",
			runtime.Version(),
			runtime.GOOS,
		})
	}
	if uint64(len(extra)) > params.MaximumExtraDataSize {
		log.Warn("Validator extra data exceed limit", "extra", hexutil.Bytes(extra), "limit", params.MaximumExtraDataSize)
		extra = nil
	}
	return extra
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
func CreateConsensusEngine(ctx *node.ServiceContext, config *Config, chainConfig *params.ChainConfig, db kcoindb.Database) consensus.Engine {
	// @TODO (rgeraldes) - complete with tendermint config if necessary, set rewarded to true
	engine := tendermint.New(&params.TendermintConfig{Rewarded: false})
	return engine
}

// APIs returns the collection of RPC services the kowala package offers.
// NOTE, some of these services probably need to be moved to somewhere else.
func (s *Kowala) APIs() []rpc.API {
	apis := kcoinapi.GetAPIs(s.ApiBackend)

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
			Namespace: "validator",
			Version:   "1.0",
			Service:   NewPrivateValidatorAPI(s),
			Public:    false,
		}, {
			Namespace: "eth",
			Version:   "1.0",
			Service:   filters.NewPublicFilterAPI(s.ApiBackend, false),
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

func (s *Kowala) Coinbase() (eb common.Address, err error) {
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
	return common.Address{}, fmt.Errorf("coinbase address must be explicitly specified")
}

func (s *Kowala) Deposit() (uint64, error) {
	s.lock.RLock()
	deposit := s.deposit
	s.lock.RUnlock()

	// @TODO(rgeraldes) - as soon as we have the dynamic validator set contract
	// if there are spots available for validators & value > min value
	// else if there are no spots available check if deposit is bigger than the the
	// smallest one

	return deposit, nil
}

// set in js console via admin interface or wrapper from cli flags
func (s *Kowala) SetCoinbase(coinbase common.Address) {
	fmt.Println("BACKEND SET_COINBASE", coinbase.String())
	s.lock.Lock()
	s.coinbase = coinbase
	s.lock.Unlock()

	walletAccount, err := s.GetWalletAccount()
	if err != nil {
		log.Error("Error setting Coinbase on validator", "err", err)
	}

	if err := s.validator.SetCoinbase(walletAccount); err != nil {
		log.Error("Error setting Coinbase on validator", "err", err)
	}
}

func (s *Kowala) GetWalletAccount() (accounts.WalletAccount, error) {
	fmt.Println("BACKEND DATA", s.coinbase.String())
	account := accounts.Account{Address: s.coinbase}
	wallet, err := s.accountManager.Find(account)
	if err != nil {
		return nil, err
	}
	return accounts.NewWalletAccount(wallet, account)
}

// GetMinimumDeposit return minimum amount required to join the validators
func (s *Kowala) GetMinimumDeposit() (uint64, error) {
	return s.election.MinimumDeposit()
}

// set in js console via admin interface or wrapper from cli flags
func (s *Kowala) SetDeposit(deposit uint64) {
	s.lock.Lock()
	s.deposit = deposit
	s.lock.Unlock()

	s.validator.SetDeposit(deposit)
}

func (s *Kowala) StartValidating() error {
	_, err := s.Coinbase()
	if err != nil {
		fmt.Println("KOWALA 1", err)
		log.Error("Cannot start consensus validation without coinbase", "err", err)
		return fmt.Errorf("coinbase missing: %v", err)
	}

	deposit, err := s.Deposit()
	if err != nil {
		fmt.Println("KOWALA 2", err)
		log.Error("Cannot start consensus validation with insufficient funds", "err", err)
		return fmt.Errorf("insufficient funds: %v", err)
	}

	// @NOTE (rgeraldes) - ignored transaction rejection mechanism introduced to speed sync times
	// @TODO (rgeraldes) - review (does it make sense to have a list of transactions before the election or not)
	atomic.StoreUint32(&s.protocolManager.acceptTxs, 1)

	walletAccount, err := s.GetWalletAccount()
	if err != nil {
		fmt.Println("KOWALA 3", err)
		return fmt.Errorf("error starting validating: %v", err)
	}

	s.validator.Start(walletAccount, deposit)
	return nil
}

func (s *Kowala) StopValidating() {
	if err := s.validator.Stop(); err != nil {
		log.Error("Error stopping Consensus", "err", err)
	}
}

func (s *Kowala) IsValidating() bool             { return s.validator.Validating() }
func (s *Kowala) Validator() validator.Validator { return s.validator }

func (s *Kowala) AccountManager() *accounts.Manager  { return s.accountManager }
func (s *Kowala) BlockChain() *core.BlockChain       { return s.blockchain }
func (s *Kowala) TxPool() *core.TxPool               { return s.txPool }
func (s *Kowala) EventMux() *event.TypeMux           { return s.eventMux }
func (s *Kowala) Engine() consensus.Engine           { return s.engine }
func (s *Kowala) ChainDb() kcoindb.Database          { return s.chainDb }
func (s *Kowala) IsListening() bool                  { return true } // Always listening
func (s *Kowala) EthVersion() int                    { return int(s.protocolManager.SubProtocols[0].Version) }
func (s *Kowala) NetVersion() uint64                 { return s.networkId }
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
	// Start the networking layer and the light server if requested
	s.protocolManager.Start(maxPeers)
	return nil
}

// Stop implements node.Service, terminating all internal goroutines used by the
// Kowala protocol.
func (s *Kowala) Stop() error {
	fmt.Println("^^^^^^^^^^^^^^^^^^^ Backend CLOSE()")
	// @NOTE (rgeraldes) - validator needs to be the first process
	// otherwise it might not be able to finish an election and
	// could be punished
	s.StopValidating()
	s.bloomIndexer.Close()
	s.blockchain.Stop()
	s.protocolManager.Stop()
	s.txPool.Stop()
	s.eventMux.Stop()

	s.chainDb.Close()
	close(s.shutdownChan)

	return nil
}
