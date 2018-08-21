package knode

import (
	"compress/gzip"
	"context"
	"errors"
	"fmt"
	"io"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/kowala-tech/kcoin/client/accounts"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/common/hexutil"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/consensus"
	"github.com/kowala-tech/kcoin/client/core"
	"github.com/kowala-tech/kcoin/client/core/rawdb"
	"github.com/kowala-tech/kcoin/client/core/state"
	"github.com/kowala-tech/kcoin/client/core/types"
	"github.com/kowala-tech/kcoin/client/internal/kcoinapi"
	"github.com/kowala-tech/kcoin/client/knode/validator"
	"github.com/kowala-tech/kcoin/client/params"
	"github.com/kowala-tech/kcoin/client/rlp"
	"github.com/kowala-tech/kcoin/client/rpc"
	"github.com/kowala-tech/kcoin/client/trie"
)

// PublicKowalaAPI provides an API to access Kowala full node-related
// information.
type PublicKowalaAPI struct {
	kcoin *Kowala
}

// NewPublicKowalaAPI creates a new Kowala protocol API for full nodes.
func NewPublicKowalaAPI(kcoin *Kowala) *PublicKowalaAPI {
	return &PublicKowalaAPI{kcoin}
}

// Coinbase is the address that consensus rewards will be send to
func (api *PublicKowalaAPI) Coinbase() (common.Address, error) {
	return api.kcoin.Coinbase()
}

// PrivateValidatorAPI provides private RPC methods to control the validator.
// These methods can be abused by external users and must be considered insecure for use by untrusted users.
type PrivateValidatorAPI struct {
	kcoin *Kowala
}

// NewPrivateValidatorAPI create a new RPC service which controls the validator of this node.
func NewPrivateValidatorAPI(kcoin *Kowala) *PrivateValidatorAPI {
	return &PrivateValidatorAPI{kcoin: kcoin}
}

// Start the validator.
func (api *PrivateValidatorAPI) Start(deposit *big.Int) error {
	// Start the validator and return
	if !api.kcoin.IsValidating() {
		// Propagate the initial price point to the transaction pool
		api.kcoin.lock.RLock()
		price := api.kcoin.gasPrice
		api.kcoin.lock.RUnlock()

		if deposit != nil {
			err := api.kcoin.SetDeposit(deposit)
			if err != nil && err != validator.ErrIsNotRunning {
				return err
			}
		}

		api.kcoin.txPool.SetGasPrice(price)
		return api.kcoin.StartValidating()
	}
	return nil
}

// Stop the validator
func (api *PrivateValidatorAPI) Stop() bool {
	api.kcoin.StopValidating()
	return true
}

// SetExtra sets the extra data string that is included when this validator proposes a block.
func (api *PrivateValidatorAPI) SetExtra(extra string) (bool, error) {
	if err := api.kcoin.Validator().SetExtra([]byte(extra)); err != nil {
		return false, err
	}
	return true, nil
}

// SetGasPrice sets the minimum accepted gas price for the validator.
func (api *PrivateValidatorAPI) SetGasPrice(gasPrice hexutil.Big) bool {
	api.kcoin.lock.Lock()
	api.kcoin.gasPrice = (*big.Int)(&gasPrice)
	api.kcoin.lock.Unlock()

	api.kcoin.txPool.SetGasPrice((*big.Int)(&gasPrice))
	return true
}

// SetCoinbase sets the coinbase of the validator
func (api *PrivateValidatorAPI) SetCoinbase(coinbase common.Address) bool {
	api.kcoin.SetCoinbase(coinbase)
	return true
}

// GetMinimumDeposit gets the minimum deposit required to take a slot as a validator
func (api *PrivateValidatorAPI) GetMinimumDeposit() (*big.Int, error) {
	return api.kcoin.GetMinimumDeposit()
}

// GetDepositsResult is the result of a validator_getDeposits API call.
type GetDepositsResult struct {
	Deposits []depositEntry `json:"deposits"`
}

type depositEntry struct {
	Amount      *big.Int `json:"value"`
	AvailableAt string   `json:",omitempty"`
}

// GetDeposits returns the validator deposits
func (api *PrivateValidatorAPI) GetDeposits(address *common.Address) (GetDepositsResult, error) {
	rawDeposits, err := api.kcoin.Validator().Deposits(address)
	if err != nil {
		return GetDepositsResult{}, err
	}

	return depositsToResponse(rawDeposits), nil
}

func depositsToResponse(rawDeposits []*types.Deposit) GetDepositsResult {
	deposits := make([]depositEntry, len(rawDeposits))

	for i, deposit := range rawDeposits {
		// @NOTE (rgeraldes) - zero values are not shown for this field
		var availableAt string

		if deposit.AvailableAtTimeUnix() != 0 {
			availableAt = time.Unix(deposit.AvailableAtTimeUnix(), 0).String()
		}

		deposits[i] = depositEntry{
			Amount:      deposit.Amount(),
			AvailableAt: availableAt,
		}
	}

	return GetDepositsResult{Deposits: deposits}
}

// IsValidating returns the validator is currently validating
func (api *PrivateValidatorAPI) IsValidating() bool {
	return api.kcoin.IsValidating()
}

// IsValidating returns the validator is currently running
func (api *PrivateValidatorAPI) IsRunning() bool {
	return api.kcoin.IsRunning()
}

// RedeemDeposits requests a transfer of the unlocked deposits back
// to the validator account
func (api *PrivateValidatorAPI) RedeemDeposits() error {
	return api.kcoin.Validator().RedeemDeposits()
}

// TransferArgs represents the arguments to transfer tokens.
type TransferArgs struct {
	From           common.Address  `json:"from"`
	To             *common.Address `json:"to"`
	Value          *hexutil.Big    `json:"value"`
	Data           hexutil.Bytes   `json:"data"`
	CustomFallback string          `json:"fallback"`
}

// PublicTokenAPI exposes a collection of methods related to tokens
type PublicTokenAPI struct {
	accountMgr *accounts.Manager
	consensus  consensus.Consensus
	chainID    *big.Int
}

func NewPublicTokenAPI(accountMgr *accounts.Manager, c consensus.Consensus, chainID *big.Int) *PublicTokenAPI {
	return &PublicTokenAPI{
		accountMgr: accountMgr,
		consensus:  c,
		chainID:    chainID,
	}
}

func (api *PublicTokenAPI) GetBalance(target common.Address) (*big.Int, error) {
	return api.consensus.Token().BalanceOf(target)
}

func (api *PublicTokenAPI) Transfer(args TransferArgs) (common.Hash, error) {
	_, walletAccount, err := api.getWallet(args.From)
	if err != nil {
		return common.Hash{}, err
	}

	if args.Value == nil {
		args.Value = new(hexutil.Big)
	}

	return api.consensus.Token().Transfer(walletAccount, *args.To, (*big.Int)(args.Value), args.Data, args.CustomFallback)
}

func (api *PublicTokenAPI) Mint(from, to common.Address, value *hexutil.Big) (common.Hash, error) {
	if value == nil {
		return common.Hash{}, errors.New("a number of tokens should be specified")
	}

	account, walletAccount, err := api.getWallet(from)
	if err != nil {
		return common.Hash{}, err
	}

	tOpts := &accounts.TransactOpts{
		From: from,
		Signer: func(signer types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return walletAccount.SignTx(*account, tx, api.chainID)
		},
	}

	return api.consensus.Mint(tOpts, to, value.ToInt())
}

func (api *PublicTokenAPI) Confirm(from common.Address, transactionID *hexutil.Big) (common.Hash, error) {
	account, walletAccount, err := api.getWallet(from)
	if err != nil {
		return common.Hash{}, err
	}

	tOpts := &accounts.TransactOpts{
		From: from,
		Signer: func(signer types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return walletAccount.SignTx(*account, tx, api.chainID)
		},
	}

	return api.consensus.Confirm(tOpts, transactionID.ToInt())
}

func (api *PublicTokenAPI) getWallet(addr common.Address) (*accounts.Account, accounts.WalletAccount, error) {
	// Look up the wallet containing the requested signer
	for _, wallet := range api.accountMgr.Wallets() {
		for _, account := range wallet.Accounts() {
			if account.Address == addr {
				walletAccount, err := accounts.NewWalletAccount(wallet, account)
				if err != nil {
					return nil, nil, err
				}
				return &account, walletAccount, nil
			}
		}
	}
	return nil, nil, errors.New("account not found in any wallet")
}

// PrivateAdminAPI is the collection of Kowala full node-related APIs
// exposed over the private admin endpoint.
type PrivateAdminAPI struct {
	kcoin *Kowala
}

// NewPrivateAdminAPI creates a new API definition for the full node private
// admin methods of the Kowala service.
func NewPrivateAdminAPI(kcoin *Kowala) *PrivateAdminAPI {
	return &PrivateAdminAPI{kcoin: kcoin}
}

// ExportChain exports the current blockchain into a local file.
func (api *PrivateAdminAPI) ExportChain(file string) (bool, error) {
	// Make sure we can create the file to export into
	out, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return false, err
	}
	defer out.Close()

	var writer io.Writer = out
	if strings.HasSuffix(file, ".gz") {
		writer = gzip.NewWriter(writer)
		defer writer.(*gzip.Writer).Close()
	}

	// Export the blockchain
	if err := api.kcoin.BlockChain().Export(writer); err != nil {
		return false, err
	}
	return true, nil
}

func hasAllBlocks(chain *core.BlockChain, bs []*types.Block) bool {
	for _, b := range bs {
		if !chain.HasBlock(b.Hash(), b.NumberU64()) {
			return false
		}
	}

	return true
}

// ImportChain imports a blockchain from a local file.
func (api *PrivateAdminAPI) ImportChain(file string) (bool, error) {
	// Make sure the can access the file to import
	in, err := os.Open(file)
	if err != nil {
		return false, err
	}
	defer in.Close()

	var reader io.Reader = in
	if strings.HasSuffix(file, ".gz") {
		if reader, err = gzip.NewReader(reader); err != nil {
			return false, err
		}
	}

	// Run actual the import in pre-configured batches
	stream := rlp.NewStream(reader, 0)

	blocks, index := make([]*types.Block, 0, 2500), 0
	for batch := 0; ; batch++ {
		// Load a batch of blocks from the input file
		for len(blocks) < cap(blocks) {
			block := new(types.Block)
			if err := stream.Decode(block); err == io.EOF {
				break
			} else if err != nil {
				return false, fmt.Errorf("block %d: failed to parse: %v", index, err)
			}
			blocks = append(blocks, block)
			index++
		}
		if len(blocks) == 0 {
			break
		}

		if hasAllBlocks(api.kcoin.BlockChain(), blocks) {
			blocks = blocks[:0]
			continue
		}
		// Import the batch and reset the buffer
		if _, err := api.kcoin.BlockChain().InsertChain(blocks); err != nil {
			return false, fmt.Errorf("batch %d: failed to insert: %v", batch, err)
		}
		blocks = blocks[:0]
	}
	return true, nil
}

// PublicDebugAPI is the collection of Kowala full node APIs exposed
// over the public debugging endpoint.
type PublicDebugAPI struct {
	kcoin *Kowala
}

// NewPublicDebugAPI creates a new API definition for the full node-
// related public debug methods of the Kowala service.
func NewPublicDebugAPI(kcoin *Kowala) *PublicDebugAPI {
	return &PublicDebugAPI{kcoin: kcoin}
}

// DumpBlock retrieves the entire state of the database at a given block.
func (api *PublicDebugAPI) DumpBlock(blockNr rpc.BlockNumber) (state.Dump, error) {
	if blockNr == rpc.PendingBlockNumber {
		// If we're dumping the pending state, we need to request
		// both the pending block as well as the pending state from
		// the validator and operate on those
		_, stateDb := api.kcoin.validator.Pending()
		return stateDb.RawDump(), nil
	}
	var block *types.Block
	if blockNr == rpc.LatestBlockNumber {
		block = api.kcoin.blockchain.CurrentBlock()
	} else {
		block = api.kcoin.blockchain.GetBlockByNumber(uint64(blockNr))
	}
	if block == nil {
		return state.Dump{}, fmt.Errorf("block #%d not found", blockNr)
	}
	stateDb, err := api.kcoin.BlockChain().StateAt(block.Root())
	if err != nil {
		return state.Dump{}, err
	}
	return stateDb.RawDump(), nil
}

// PrivateDebugAPI is the collection of Kowala full node APIs exposed over
// the private debugging endpoint.
type PrivateDebugAPI struct {
	config *params.ChainConfig
	kcoin  *Kowala
}

// NewPrivateDebugAPI creates a new API definition for the full node-related
// private debug methods of the Kowala service.
func NewPrivateDebugAPI(config *params.ChainConfig, kcoin *Kowala) *PrivateDebugAPI {
	return &PrivateDebugAPI{config: config, kcoin: kcoin}
}

// Preimage is a debug API function that returns the preimage for a sha3 hash, if known.
func (api *PrivateDebugAPI) Preimage(ctx context.Context, hash common.Hash) (hexutil.Bytes, error) {
	if preimage := rawdb.ReadPreimage(api.kcoin.ChainDb(), hash); preimage != nil {
		return preimage, nil
	}
	return nil, errors.New("unknown preimage")
}

// BadBlockArgs represents the entries in the list returned when bad blocks are queried.
type BadBlockArgs struct {
	Hash  common.Hash            `json:"hash"`
	Block map[string]interface{} `json:"block"`
	RLP   string                 `json:"rlp"`
}

// GetBadBlocks returns a list of the last 'bad blocks' that the client has seen on the network
// and returns them as a JSON list of block-hashes
func (api *PrivateDebugAPI) GetBadBlocks(ctx context.Context) ([]*BadBlockArgs, error) {
	blocks := api.kcoin.BlockChain().BadBlocks()
	results := make([]*BadBlockArgs, len(blocks))

	var err error
	for i, block := range blocks {
		results[i] = &BadBlockArgs{
			Hash: block.Hash(),
		}
		if rlpBytes, err := rlp.EncodeToBytes(block); err != nil {
			results[i].RLP = err.Error() // Hacky, but hey, it works
		} else {
			results[i].RLP = fmt.Sprintf("0x%x", rlpBytes)
		}
		if results[i].Block, err = kcoinapi.RPCMarshalBlock(block, true, true); err != nil {
			results[i].Block = map[string]interface{}{"error": err.Error()}
		}
	}
	return results, nil
}

// StorageRangeResult is the result of a debug_storageRangeAt API call.
type StorageRangeResult struct {
	Storage storageMap   `json:"storage"`
	NextKey *common.Hash `json:"nextKey"` // nil if Storage includes the last key in the trie.
}

type storageMap map[common.Hash]storageEntry

type storageEntry struct {
	Key   *common.Hash `json:"key"`
	Value common.Hash  `json:"value"`
}

// StorageRangeAt returns the storage at the given block height and transaction index.
func (api *PrivateDebugAPI) StorageRangeAt(ctx context.Context, blockHash common.Hash, txIndex int, contractAddress common.Address, keyStart hexutil.Bytes, maxResult int) (StorageRangeResult, error) {
	_, _, statedb, err := api.computeTxEnv(blockHash, txIndex, 0)
	if err != nil {
		return StorageRangeResult{}, err
	}
	st := statedb.StorageTrie(contractAddress)
	if st == nil {
		return StorageRangeResult{}, fmt.Errorf("account %x doesn't exist", contractAddress)
	}
	return storageRangeAt(st, keyStart, maxResult)
}

func storageRangeAt(st state.Trie, start []byte, maxResult int) (StorageRangeResult, error) {
	it := trie.NewIterator(st.NodeIterator(start))
	result := StorageRangeResult{Storage: storageMap{}}
	for i := 0; i < maxResult && it.Next(); i++ {
		_, content, _, err := rlp.Split(it.Value)
		if err != nil {
			return StorageRangeResult{}, err
		}
		e := storageEntry{Value: common.BytesToHash(content)}
		if preimage := st.GetKey(it.Key); preimage != nil {
			preimage := common.BytesToHash(preimage)
			e.Key = &preimage
		}
		result.Storage[common.BytesToHash(it.Key)] = e
	}
	// Add the 'next key' so clients can continue downloading.
	if it.Next() {
		next := common.BytesToHash(it.Key)
		result.NextKey = &next
	}
	return result, nil
}

// GetModifiedAccountsByNumber returns all accounts that have changed between the
// two blocks specified. A change is defined as a difference in nonce, balance,
// code hash, or storage hash.
//
// With one parameter, returns the list of accounts modified in the specified block.
func (api *PrivateDebugAPI) GetModifiedAccountsByNumber(startNum uint64, endNum *uint64) ([]common.Address, error) {
	var startBlock, endBlock *types.Block

	startBlock = api.kcoin.blockchain.GetBlockByNumber(startNum)
	if startBlock == nil {
		return nil, fmt.Errorf("start block %x not found", startNum)
	}

	if endNum == nil {
		endBlock = startBlock
		startBlock = api.kcoin.blockchain.GetBlockByHash(startBlock.ParentHash())
		if startBlock == nil {
			return nil, fmt.Errorf("block %x has no parent", endBlock.Number())
		}
	} else {
		endBlock = api.kcoin.blockchain.GetBlockByNumber(*endNum)
		if endBlock == nil {
			return nil, fmt.Errorf("end block %d not found", *endNum)
		}
	}
	return api.getModifiedAccounts(startBlock, endBlock)
}

// GetModifiedAccountsByHash returns all accounts that have changed between the
// two blocks specified. A change is defined as a difference in nonce, balance,
// code hash, or storage hash.
//
// With one parameter, returns the list of accounts modified in the specified block.
func (api *PrivateDebugAPI) GetModifiedAccountsByHash(startHash common.Hash, endHash *common.Hash) ([]common.Address, error) {
	var startBlock, endBlock *types.Block
	startBlock = api.kcoin.blockchain.GetBlockByHash(startHash)
	if startBlock == nil {
		return nil, fmt.Errorf("start block %x not found", startHash)
	}

	if endHash == nil {
		endBlock = startBlock
		startBlock = api.kcoin.blockchain.GetBlockByHash(startBlock.ParentHash())
		if startBlock == nil {
			return nil, fmt.Errorf("block %x has no parent", endBlock.Number())
		}
	} else {
		endBlock = api.kcoin.blockchain.GetBlockByHash(*endHash)
		if endBlock == nil {
			return nil, fmt.Errorf("end block %x not found", *endHash)
		}
	}
	return api.getModifiedAccounts(startBlock, endBlock)
}

func (api *PrivateDebugAPI) getModifiedAccounts(startBlock, endBlock *types.Block) ([]common.Address, error) {
	if startBlock.Number().Uint64() >= endBlock.Number().Uint64() {
		return nil, fmt.Errorf("start block height (%d) must be less than end block height (%d)", startBlock.Number().Uint64(), endBlock.Number().Uint64())
	}

	oldTrie, err := trie.NewSecure(startBlock.Root(), trie.NewDatabase(api.kcoin.chainDb), 0)
	if err != nil {
		return nil, err
	}
	newTrie, err := trie.NewSecure(endBlock.Root(), trie.NewDatabase(api.kcoin.chainDb), 0)
	if err != nil {
		return nil, err
	}

	diff, _ := trie.NewDifferenceIterator(oldTrie.NodeIterator([]byte{}), newTrie.NodeIterator([]byte{}))
	iter := trie.NewIterator(diff)

	var dirty []common.Address
	for iter.Next() {
		key := newTrie.GetKey(iter.Key)
		if key == nil {
			return nil, fmt.Errorf("no preimage found for hash %x", iter.Key)
		}
		dirty = append(dirty, common.BytesToAddress(key))
	}
	return dirty, nil
}
