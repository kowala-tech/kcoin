package mining

import (
	"fmt"
	"math/big"
	"runtime"
	"sync"

	"github.com/kowala-tech/kcoin/client/accounts"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/common/hexutil"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/consensus"
	"github.com/kowala-tech/kcoin/client/core"
	"github.com/kowala-tech/kcoin/client/event"
	"github.com/kowala-tech/kcoin/client/log"
	"github.com/kowala-tech/kcoin/client/node"
	"github.com/kowala-tech/kcoin/client/p2p"
	"github.com/kowala-tech/kcoin/client/params"
	"github.com/kowala-tech/kcoin/client/rlp"
	"github.com/kowala-tech/kcoin/client/rpc"
	root "github.com/kowala-tech/kcoin/client/services"
	"github.com/kowala-tech/kcoin/client/services/knode"
	"github.com/kowala-tech/kcoin/client/services/mining/validator"
)

// Service provides mining as a service.
type Service struct {
	lock     sync.RWMutex
	coinbase common.Address
	deposit  *big.Int
	gasPrice *big.Int

	validator validator.Validator

	validatorMgr *consensus.Consensus
	protocolMgr  *ProtocolManager
	globalEvents *event.TypeMux
	txPool       *core.TxPool
	accountMgr   *accounts.Manager

	kowalaService *knode.Kowala

	log log.Logger

	shutdownChan chan bool // Channel for shutting down the service
}

// New returns an instance of the mining service.
func New(ctx *node.ServiceContext, config *Config) (*Service, error) {
	// @TODO - logger nil?

	service := &Service{
		coinbase:     config.Coinbase,
		deposit:      config.Deposit,
		gasPrice:     config.GasPrice,
		globalEvents: ctx.GlobalEventMux,
		accountMgr:   ctx.AccountManager,
		shutdownChan: make(chan bool),
	}

	service.log.Info("Initialising Mining Service")

	if err := ctx.Service(&service.kowalaService); err != nil {
		return nil, err
	}

	binding, err := consensus.Bind(root.NewContractBackend(service.kowalaService.APIBackend()), service.kowalaService.BlockChain().Config().ChainID)
	if err != nil {
		return nil, err
	}
	service.validatorMgr = binding.(*consensus.Consensus)

	service.validator = validator.New(service.kowalaService, service.validatorMgr, service.globalEvents)
	service.validator.SetExtra(makeExtraData(config.ExtraData))

	if service.protocolMgr, err = NewProtocolManager(service.kowalaService.ChainID().Uint64(), service.validator, service.validator.VotingSystem(), service.kowalaService.BlockChain(), service.log); err != nil {
		return nil, err
	}

	service.txPool = service.kowalaService.TxPool()

	return service, nil
}

func (s *Service) Deposit() (*big.Int, error) {
	s.lock.RLock()
	deposit := s.deposit
	s.lock.RUnlock()

	return deposit, nil
}

// set in js console via admin interface or wrapper from cli flags
func (s *Service) SetCoinbase(coinbase common.Address) {
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

func (s *Service) getWalletAccount() (accounts.WalletAccount, error) {
	account := accounts.Account{Address: s.coinbase}
	wallet, err := s.accountMgr.Find(account)
	if err != nil {
		return nil, err
	}
	return accounts.NewWalletAccount(wallet, account)
}

// GetMinimumDeposit return minimum amount required to join the validators
func (s *Service) GetMinimumDeposit() (*big.Int, error) {
	return s.validatorMgr.MinimumDeposit()
}

// set in js console via admin interface or wrapper from cli flags
func (s *Service) SetDeposit(deposit *big.Int) error {
	s.lock.Lock()
	s.deposit = deposit
	s.lock.Unlock()

	return s.validator.SetDeposit(deposit)
}

func (s *Service) StartValidating() error {
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
	// atomic.StoreUint32(&s.protocolMgr.acceptTxs, 1)

	walletAccount, err := s.getWalletAccount()
	if err != nil {
		return fmt.Errorf("error starting validating: %v", err)
	}

	s.validator.Start(walletAccount, deposit)
	return nil
}

func (s *Service) StopValidating() {
	if err := s.validator.Stop(); err != nil {
		log.Error("Error stopping Consensus", "err", err)
	}
}

func (s *Service) IsValidating() bool                 { return s.validator.Validating() }
func (s *Service) IsRunning() bool                    { return s.validator.Running() }
func (s *Service) Validator() validator.Validator     { return s.validator }
func (s *Service) ValidatorMgr() *consensus.Consensus { return s.validatorMgr }
func (s *Service) AccountMgr() *accounts.Manager      { return s.accountMgr }

func (s *Service) Coinbase() (eb common.Address, err error) {
	s.lock.RLock()
	coinbase := s.coinbase
	s.lock.RUnlock()

	if coinbase != (common.Address{}) {
		return coinbase, nil
	}
	if wallets := s.AccountMgr().Wallets(); len(wallets) > 0 {
		if accounts := wallets[0].Accounts(); len(accounts) > 0 {
			return accounts[0].Address, nil
		}
	}
	return common.Address{}, fmt.Errorf("coinbase address must be explicitly specified")
}

// Start implements node.Service, starting all internal goroutines needed by the
// Consensus protocol implementation.
func (s *Service) Start(srvr *p2p.Server) error {
	// Start the networking layer and the light server if requested
	s.protocolMgr.Start(srvr.MaxPeers)

	return nil
}

// Stop implements node.Service, terminating all internal goroutines used by the
// Mining service protocol.
func (s *Service) Stop() error {
	s.StopValidating()
	s.protocolMgr.Stop()
	close(s.shutdownChan)
	return nil
}

// APIs returns the collection of RPC services the mining service offers.
func (s *Service) APIs() []rpc.API {
	return []rpc.API{
		{
			Namespace: "validator",
			Version:   "1.0",
			Service:   NewPrivateValidatorAPI(s),
			Public:    false,
		},
		{
			Namespace: "mtoken",
			Version:   "1.0",
			Service:   NewPublicTokenAPI(s.accountMgr, s.validatorMgr, s.kowalaService.ChainID()),
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
