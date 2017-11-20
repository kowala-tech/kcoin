package tendermint

import (
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/consensus"
	"github.com/kowala-tech/kUSD/core/state"
	"github.com/kowala-tech/kUSD/core/types"
	"github.com/kowala-tech/kUSD/node"
	"github.com/kowala-tech/kUSD/rpc"
	"github.com/kowala-tech/kUSD/tendermint"
)

var _ consensus.Engine = ((*Tendermint)(nil))

type Tendermint struct {
	ctx     *node.ServiceContext
	signers []common.Address
}

// New creates a tendermint consensus engine
func New(ctx *node.ServiceContext) *Tendermint {
	fmt.Println("New() tendermint consensus engine")
	return &Tendermint{ctx: ctx}
}

func (t *Tendermint) gossipService() (*tendermint.GossipService, error) {
	w := &tendermint.GossipService{}
	err := t.ctx.Service(&w)
	return w, err
}

// Author retrieves the Ethereum address of the account that minted the given
// block, which may be different from the header's coinbase if a consensus
// engine is based on signatures.
func (t *Tendermint) Author(header *types.Header) (common.Address, error) {
	fmt.Println("tendermint_consensus.Author()")
	return header.Coinbase, nil
}

// VerifyHeader checks whether a header conforms to the consensus rules of a
// given engine. Verifying the seal may be done optionally here, or explicitly
// via the VerifySeal method.
func (t *Tendermint) VerifyHeader(chain consensus.ChainReader, header *types.Header, seal bool) error {
	fmt.Println("tendermint_consensus.VerifyHeader()")
	// unknown block
	if header.Number == nil {
		return ErrUnknownBlock
	}
	// block in the future
	if header.Time.Cmp(big.NewInt(time.Now().Unix())) > 0 {
		return consensus.ErrFutureBlock
	}
	// check seal
	if seal {
		if err := t.VerifySeal(chain, header); err != nil {
			return err
		}
	}
	return nil
}

// VerifyHeaders is similar to VerifyHeader, but verifies a batch of headers
// concurrently. The method returns a quit channel to abort the operations and
// a results channel to retrieve the async verifications (the order is that of
// the input slice).
func (t *Tendermint) VerifyHeaders(chain consensus.ChainReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {
	fmt.Println("tendermint_consensus.VerifyHeaders()")
	return make(chan struct{}, 0), make(chan error, 0)
}

var (
	ErrNoUnclesAllowed = errors.New("uncle blocks not allowed")
	ErrUnknownBlock    = errors.New("unknown block")
)

// VerifyUncles verifies that the given block's uncles conform to the consensus
// rules of a given engine.
func (t *Tendermint) VerifyUncles(chain consensus.ChainReader, block *types.Block) error {
	fmt.Println("tendermint_consensus.VerifyUncles()")
	if len(block.Uncles()) > 0 {
		return ErrNoUnclesAllowed
	}
	return nil
}

// VerifySeal checks whether the crypto seal on a header is valid according to
// the consensus rules of the given engine.
func (t *Tendermint) VerifySeal(chain consensus.ChainReader, header *types.Header) error {
	fmt.Println("tendermint_consensus.VerifySeal()")
	return nil
}

// Prepare initializes the consensus fields of a block header according to the
// rules of a particular engine. The changes are executed inline.
func (t *Tendermint) Prepare(chain consensus.ChainReader, header *types.Header) error {
	return nil
}

// Finalize runs any post-transaction state modifications (e.g. block rewards)
// and assembles the final block.
// Note: The block header and state database might be updated to reflect any
// consensus rules that happen at finalization (e.g. block rewards).
func (t *Tendermint) Finalize(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error) {
	return nil, nil
}

// Seal generates a new block for the given input block with the local miner's
// seal place on top.
func (t *Tendermint) Seal(chain consensus.ChainReader, block *types.Block, stop <-chan struct{}) (*types.Block, error) {
	fmt.Println("tendermint_consensus.Seal()")
	return nil, nil
}

// APIs returns the RPC APIs this consensus engine provides.
func (t *Tendermint) APIs(chain consensus.ChainReader) []rpc.API {
	fmt.Println("tendermint_consensus.APIs()")
	return nil
}
