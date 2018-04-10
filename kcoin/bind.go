package kcoin

import (
	"context"
	"math/big"

	"github.com/kowala-tech/kcoin"
	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/common/hexutil"
	"github.com/kowala-tech/kcoin/core/types"
	"github.com/kowala-tech/kcoin/internal/kcoinapi"
	"github.com/kowala-tech/kcoin/rlp"
	"github.com/kowala-tech/kcoin/rpc"
)
//FilterLogs(ctx context.Context, query kowala.FilterQuery) ([]types.Log, error)
//SubscribeFilterLogs(ctx context.Context, query kowala.FilterQuery, ch chan<- types.Log) (kowala.Subscription, error)
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
	nilLogger
}

// NewContractBackend creates a new native contract backend using an existing
// Kowala object.
func NewContractBackend(apiBackend kcoinapi.Backend) *ContractBackend {
	return &ContractBackend{
		eapi:      kcoinapi.NewPublicKowalaAPI(apiBackend),
		bcapi:     kcoinapi.NewPublicBlockChainAPI(apiBackend),
		txapi:     kcoinapi.NewPublicTransactionPoolAPI(apiBackend, new(kcoinapi.AddrLocker)),
		nilLogger: newNilLogger(),
	}
}

// CodeAt retrieves any code associated with the contract from the local API.
func (b *ContractBackend) CodeAt(ctx context.Context, contract common.Address, blockNum *big.Int) ([]byte, error) {
	return b.bcapi.GetCode(ctx, contract, toBlockNumber(blockNum))
}

// CodeAt retrieves any code associated with the contract from the local API.
func (b *ContractBackend) PendingCodeAt(ctx context.Context, contract common.Address) ([]byte, error) {
	return b.bcapi.GetCode(ctx, contract, rpc.PendingBlockNumber)
}

// ContractCall implements bind.ContractCaller executing an Kowala contract
// call with the specified data as the input. The pending flag requests execution
// against the pending block, not the stable head of the chain.
func (b *ContractBackend) CallContract(ctx context.Context, msg kowala.CallMsg, blockNum *big.Int) ([]byte, error) {
	out, err := b.bcapi.Call(ctx, toCallArgs(msg), toBlockNumber(blockNum))
	return out, err
}

// ContractCall implements bind.ContractCaller executing an Kowala contract
// call with the specified data as the input. The pending flag requests execution
// against the pending block, not the stable head of the chain.
func (b *ContractBackend) PendingCallContract(ctx context.Context, msg kowala.CallMsg) ([]byte, error) {
	out, err := b.bcapi.Call(ctx, toCallArgs(msg), rpc.PendingBlockNumber)
	return out, err
}

func toCallArgs(msg kowala.CallMsg) kcoinapi.CallArgs {
	args := kcoinapi.CallArgs{
		To:   msg.To,
		From: msg.From,
		Data: msg.Data,
	}
	if msg.Gas != 0 {
		args.Gas = hexutil.Uint64(msg.Gas)
	}
	if msg.GasPrice != nil {
		args.GasPrice = hexutil.Big(*msg.GasPrice)
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

// PendingAccountNonce implements bind.ContractTransactor retrieving the current
// pending nonce associated with an account.
func (b *ContractBackend) PendingNonceAt(ctx context.Context, account common.Address) (nonce uint64, err error) {
	out, err := b.txapi.GetTransactionCount(ctx, account, rpc.PendingBlockNumber)
	if out != nil {
		nonce = uint64(*out)
	}
	return nonce, err
}

// SuggestGasPrice implements bind.ContractTransactor retrieving the currently
// suggested gas price to allow a timely execution of a transaction.
func (b *ContractBackend) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return b.eapi.GasPrice(ctx)
}

// EstimateGasLimit implements bind.ContractTransactor triing to estimate the gas
// needed to execute a specific transaction based on the current pending state of
// the backend blockchain. There is no guarantee that this is the true gas limit
// requirement as other transactions may be added or removed by validators, but it
// should provide a basis for setting a reasonable default.
func (b *ContractBackend) EstimateGas(ctx context.Context, msg kowala.CallMsg) (uint64, error) {
	out, err := b.bcapi.EstimateGas(ctx, toCallArgs(msg))
	return uint64(out), err
}

// SendTransaction implements bind.ContractTransactor injects the transaction
// into the pending pool for execution.
func (b *ContractBackend) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	raw, _ := rlp.EncodeToBytes(tx)
	_, err := b.txapi.SendRawTransaction(ctx, raw)
	return err
}

// FilterLogs executes a filter query.
func (b *ContractBackend) FilterLogs(ctx context.Context, q kowala.FilterQuery) ([]types.Log, error) {
	return b.FilterLogs(ctx, q)
}

// SubscribeFilterLogs subscribes to the results of a streaming filter query.
func (b *ContractBackend) SubscribeFilterLogs(ctx context.Context, q kowala.FilterQuery, ch chan<- types.Log) (kowala.Subscription, error) {
	return b.SubscribeFilterLogs(ctx, q, ch)
}