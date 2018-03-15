package validator

import (
	"github.com/kowala-tech/kUSD/accounts"
	"github.com/kowala-tech/kUSD/contracts/network"
	"github.com/kowala-tech/kUSD/params"
	"github.com/kowala-tech/kUSD/event"
	"github.com/kowala-tech/kUSD/consensus"
	"github.com/kowala-tech/kUSD/core/vm"
	"github.com/kowala-tech/kUSD/core/types"
	"github.com/kowala-tech/kUSD/core"
)

// context represents consensus validator configuration
type context struct {
	extra         []byte
	walletAccount accounts.WalletAccount
	deposit       uint64
	backend       Backend
	chain         *core.BlockChain
	contract      *network.NetworkContract
	chainConfig   *params.ChainConfig
	eventMux      *event.TypeMux
	engine        consensus.Engine
	vmConfig      vm.Config
	signer        types.Signer
}

func NewValidatorContext(extra []byte, deposit uint64, walletAccount accounts.WalletAccount, backend Backend, contract *network.NetworkContract, config *params.ChainConfig, eventMux *event.TypeMux, engine consensus.Engine, vmConfig vm.Config) *context {
	return &context{
		extra:         extra,
		deposit:       deposit,
		walletAccount: walletAccount,
		backend:       backend,
		chain:         backend.BlockChain(),
		contract:      contract,
		chainConfig:   config,
		eventMux:      eventMux,
		engine:        engine,
		vmConfig:      vmConfig,
		signer:        types.NewAndromedaSigner(config.ChainID),
	}
}
