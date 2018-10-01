package consensus

import (
	"github.com/kowala-tech/kcoin/client/internal/kcoinapi"
	"sync"

	"github.com/kowala-tech/kcoin/client/consensus/validator"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/consensus"
	"github.com/kowala-tech/kcoin/client/knode"
	"github.com/kowala-tech/kcoin/client/params"
)

type Core interface {
	APIBackend() kcoinapi.Backend
	BlockChain() *core.BlockChain
	TxPool() *core.TxPool
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
}

// MiningService provides mining as a service.
type MiningService struct {
	lock       sync.RWMutex
	validator validator.Validator
	coinbase common.Address
	deposit  *big.Int

	validatorMgr *consensus.Consensus 
	protocolMgr *ProtocolManager
	serverPool *serverPool

	log log.Logger

	shutdownChan chan bool // Channel for shutting down the service
}

// NewMiningService returns an instance of the mining service.
func NewMiningService(core Core, config *Config) (*Service, error) {
	service := &MiningService{
		log: config.Logger,	
		coinbase: config.Coinbase,
		deposit:  config.Deposit,
	}

	service.log.Info("Initialising Mining Service")

	validatorMgr, err := consensus.Bind(NewContractBackend(core.APIBackend()), core.BlockChain().Config().ChainID)
	if err != nil {
		return nil, err
	}
	service.validatorMgr = validatorMgr

	service.validator = validator.New(core, service.validatorMgr)
	service.validator.SetExtra(makeExtraData(config.ExtraData))
	
	if service.protocolManager, err = NewProtocolManager(core.BlockChain().Config(), config.SyncMode, config.NetworkId, kcoin.eventMux, kcoin.txPool, kcoin.engine, kcoin.blockchain, chainDb); err != nil {
		return nil, err
	}

	kcoin.serverPool = newServerPool(chainDb, kcoin.shutdownChan, new(sync.WaitGroup))

	service, nil
}

func (ms *MiningService) Deposit() (*big.Int, error) {
	s.lock.RLock()
	deposit := s.deposit
	s.lock.RUnlock()

	return deposit, nil
}

// set in js console via admin interface or wrapper from cli flags
func (ms *MiningService) SetCoinbase(coinbase common.Address) {
	s.lock.Lock()
	s.coinbase = coinbase
	s.lock.Unlock()

	walletAccount, err := s.getWalletAccount()
	if err != nil {
		log.Error("error SetCoinbase on validator getWalletAccount", "err", err)
	}

	if err := s.validator.SetCoinbase(walletAccount); err != nil {
		log.Error("error SetCoinbase on validator setCoinbase", "err", err)
	}
}

func (ms *MiningService) getWalletAccount() (accounts.WalletAccount, error) {
	account := accounts.Account{Address: s.coinbase}
	wallet, err := s.accountManager.Find(account)
	if err != nil {
		return nil, err
	}
	return accounts.NewWalletAccount(wallet, account)
}

// GetMinimumDeposit return minimum amount required to join the validators
func (ms *MiningService) GetMinimumDeposit() (*big.Int, error) {
	return s.consensus.MinimumDeposit()
}

// set in js console via admin interface or wrapper from cli flags
func (ms *MiningService) SetDeposit(deposit *big.Int) error {
	s.lock.Lock()
	s.deposit = deposit
	s.lock.Unlock()

	return s.validator.SetDeposit(deposit)
}

func (ms *MiningService) StartValidating() error {
	_, err := s.Coinbase()
	if err != nil {
		log.Error("Cannot start consensus validation without coinbase", "err", err)
		return fmt.Errorf("coinbase missing: %v", err)
	}

	deposit, err := s.Deposit()
	if err != nil {
		log.Error("Cannot start consensus validation with insufficient funds", "err", err)
		return fmt.Errorf("insufficient funds: %v", err)
	}

	// @NOTE (rgeraldes) - ignored transaction rejection mechanism introduced to speed sync times
	// @TODO (rgeraldes) - review (does it make sense to have a list of transactions before the election or not)
	atomic.StoreUint32(&s.protocolManager.acceptTxs, 1)

	walletAccount, err := s.getWalletAccount()
	if err != nil {
		return fmt.Errorf("error starting validating: %v", err)
	}

	s.validator.Start(walletAccount, deposit)
	return nil
}

func (ms *MiningService) StopValidating() {
	if err := s.validator.Stop(); err != nil {
		log.Error("Error stopping Consensus", "err", err)
	}
}

func (ms *MiningService) IsValidating() bool             { return s.validator.Validating() }
func (ms *MiningService) IsRunning() bool                { return s.validator.Running() }
func (ms *MiningService) Validator() validator.Validator { return s.validator }

// Start implements node.Service, starting all internal goroutines needed by the
// Consensus protocol implementation.
func (ms *MiningService) Start(srvr *p2p.Server) error {
	//fixme: should be removed after develop light client
	if srvr.DiscoveryV5 {
		protocolTopic := discv5.DiscoveryTopic(s.blockchain.Genesis().Hash(), protocol.ProtocolName, protocol.Kcoin1)

		go func() {
			srvr.DiscV5.RegisterTopic(protocolTopic, s.shutdownChan)
		}()

		s.serverPool.start(srvr, protocolTopic)
	}

	// Start the networking layer and the light server if requested
	s.protocolManager.Start(srvr.MaxPeers)

	return nil
}

// Stop implements node.Service, terminating all internal goroutines used by the
// Mining service protocol.
func (ms *MiningService) Stop() error {
	ms.StopValidating()
	s.protocolManager.Stop()
	close(s.shutdownChan)
}

func (ms *MiningService) Coinbase() (eb common.Address, err error) {
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



// APIs returns the collection of RPC services the mining service offers.
func (ms *MiningService) APIs() []rpc.API {
	return []rpc.API{
		{
			Namespace: "validator",
			Version:   "1.0",
			Service:   NewPrivateValidatorAPI(ms),
			Public:    false,
		},
	}
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