package knode

import (
	"context"
	"math/big"

	"github.com/kowala-tech/kcoin/client"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/common/hexutil"
	"github.com/kowala-tech/kcoin/client/core/types"
	"github.com/kowala-tech/kcoin/client/internal/kcoinapi"
	"github.com/kowala-tech/kcoin/client/rlp"
	"github.com/kowala-tech/kcoin/client/rpc"
)

// ContractBackend implements bind.ContractBackend with direct calls to Kowala
// internals to support operating on contracts within subprotocols like kcoin and
// swarm.
//
// Internally this backend uses the already exposed API endpoints of the Kowala
// object. These should be rewritten to internal Go method calls when the Go API
// is refactored to support a clean library use.
type ContractBackend struct {
	eapi  *kcoinapi.PublicKowalaAPI          // Wrapper around the Kowala object to access metadata
	bcapi *kcoinapi.PublicBlockChainAPI      // Wrapper around the blockchain to access chain data
	txapi *kcoinapi.PublicTransactionPoolAPI // Wrapper around the transaction pool to access transaction data
}

// NewContractBackend creates a new native contract backend using an existing
// Kowala object.
func NewContractBackend(apiBackend kcoinapi.Backend) *ContractBackend {
	return &ContractBackend{
		eapi:  kcoinapi.NewPublicKowalaAPI(apiBackend),
		bcapi: kcoinapi.NewPublicBlockChainAPI(apiBackend),
		txapi: kcoinapi.NewPublicTransactionPoolAPI(apiBackend, new(kcoinapi.AddrLocker)),
	}
}

// CodeAt retrieves any code associated with the contract from the local API.
func (b *ContractBackend) CodeAt(ctx context.Context, contract common.Address, blockNum *big.Int) ([]byte, error) {
	return b.bcapi.GetCode(ctx, contract, toBlockNumber(blockNum))
}

// PendingCodeAt retrieves any code associated with the contract from the local API.
func (b *ContractBackend) PendingCodeAt(ctx context.Context, contract common.Address) ([]byte, error) {
	return b.bcapi.GetCode(ctx, contract, rpc.PendingBlockNumber)
}

// CallContract implements bind.ContractCaller executing an Kowala contract
// call with the specified data as the input. The pending flag requests execution
// against the pending block, not the stable head of the chain.
func (b *ContractBackend) CallContract(ctx context.Context, msg kowala.CallMsg, blockNum *big.Int) ([]byte, error) {
	out, err := b.bcapi.Call(ctx, toCallArgs(msg), toBlockNumber(blockNum))
	return out, err
}

// PendingCallContract implements bind.ContractCaller executing an Kowala contract
// call with the specified data as the input. The pending flag requests execution
// against the pending block, not the stable head of the chain.
func (b *ContractBackend) PendingCallContract(ctx context.Context, msg kowala.CallMsg) ([]byte, error) {
	out, err := b.bcapi.Call(ctx, toCallArgs(msg), rpc.PendingBlockNumber)
	return out, err
}

func toCallArgs(msg kowala.CallMsg) kcoinapi.CallArgs {
	args := kcoinapi.CallArgs{
		To:           msg.To,
		From:         msg.From,
		Data:         msg.Data,
		ComputeLimit: hexutil.Uint64(msg.ComputeLimit),
	}
	if msg.Value != nil {
		args.Value = hexutil.Big(*msg.Value)
	}
	return args
}

func toBlockNumber(num *big.Int) rpc.BlockNumber {
	if num == nil {
		return rpc.LatestBlockNumber
	}
	return rpc.BlockNumber(num.Int64())
}

// PendingNonceAt implements bind.ContractTransactor retrieving the current
// pending nonce associated with an account.
func (b *ContractBackend) PendingNonceAt(ctx context.Context, account common.Address) (nonce uint64, err error) {
	out, err := b.txapi.GetTransactionCount(ctx, account, rpc.PendingBlockNumber)
	if out != nil {
		nonce = uint64(*out)
	}
	return nonce, err
}

// EstimateComputationalEffort implements bind.ContractTransactor trying to estimate the required
// computational effort to execute a specific transaction based on the current pending state of
// the backend blockchain. There is no guarantee that this is the true computational effort
// requirement as other transactions may be added or removed by validators, but it
// should provide a basis for setting a reasonable default.
func (b *ContractBackend) EstimateComputationalEffort(ctx context.Context, msg kowala.CallMsg) (uint64, error) {
	out, err := b.bcapi.EstimateComputationalEffort(ctx, toCallArgs(msg))
	if err != nil {
		return 0, err
	}

	return hexutil.MustDecodeUint64(out.String()), nil
}

// SendTransaction implements bind.ContractTransactor injects the transaction
// into the pending pool for execution.
func (b *ContractBackend) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	raw, _ := rlp.EncodeToBytes(tx)
	_, err := b.txapi.SendRawTransaction(ctx, raw)
	return err
}

func (b *ContractBackend) FilterLogs(ctx context.Context, query kowala.FilterQuery) ([]types.Log, error) {
	return nil, nil
}

func (b *ContractBackend) SubscribeFilterLogs(ctx context.Context, query kowala.FilterQuery, ch chan<- types.Log) (kowala.Subscription, error) {
	return nil, nil
}
