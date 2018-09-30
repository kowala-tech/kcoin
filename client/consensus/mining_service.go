package consensus

import (
	"sync"

	"github.com/kowala-tech/kcoin/client/consensus/validator"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/consensus"
)

type txPool interface {
	// SubscribeNewTxsEvent should return an event subscription of
	// NewTxsEvent and send events to the given channel.
	SubscribeNewTxsEvent(chan<- core.NewTxsEvent) event.Subscription
}

type blockChain interface {
	SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription
}

type votingSystem interface {
	// SubscribeNewVoteEvent should return an event subscription of
	// NewVoteEvent and send events to the given channel.
	SubscribeNewVoteEvent(chan<- core.NewVoteEvent) event.Subscription

	// SubscribeNewMajorityEvent should return an event subscription of
	// NewMajorityEvent and send events to the given channel.
	SubscribeNewMajorityEvent(chan<- core.NewMajorityEvent) event.Subscription
}

// MiningService represents mining as a service.
type MiningService struct {
	lock       sync.RWMutex
	
	validatorMgr *consensus.Consensus
	validator validator.Validator // consensus validator

	coinbase common.Address
	deposit  *big.Int
	
	eventMux       *event.TypeMux

	log log.Logger
}

// NewMiningService returns an instance of the mining service.
func NewMiningService(ctx *node.ServiceContext, kowalaServ *knode.Kowala, config *Config) (*Service, error) {
	validatorMgr, err := consensus.Bind(NewContractBackend(kcoin.apiBackend),kcoin.chainConfig.ChainID)
	if err != nil {
		return err
	}

	validator := validator.New(kcoin, kcoin.consensus, kcoin.chainConfig, kcoin.EventMux(), kcoin.engine, vmConfig)
	validator.SetExtra(makeExtraData(config.ExtraData))

	return &Service{
		validator: validator,
		validatorMgr: validatorMgr,
		coinbase:       config.Coinbase,
		deposit:        config.Deposit,
		eventMux: ctx.EventMux,
	}, nil
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
// Kowala protocol.
func (ms *MiningService) Stop() error {
	ms.StopValidating()
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

func (ms *MiningService) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
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
