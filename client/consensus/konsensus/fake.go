package konsensus

import (
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/consensus"
	"github.com/kowala-tech/kcoin/client/core/state"
	"github.com/kowala-tech/kcoin/client/core/types"
	"github.com/kowala-tech/kcoin/client/rpc"
)

type FakeKonsensus struct{}

func NewFaker() *FakeKonsensus {
	return &FakeKonsensus{}
}

func (fk *FakeKonsensus) Author(header *types.Header) (common.Address, error) {
	return header.Coinbase, nil
}

func (fk *FakeKonsensus) VerifyHeader(chain consensus.ChainReader, header *types.Header, seal bool) error {
	return nil
}

func (fk *FakeKonsensus) VerifyHeaders(chain consensus.ChainReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {
	// @TODO (rgeraldes) - temporary work around
	abort, results := make(chan struct{}), make(chan error, len(headers))
	for i := 0; i < len(headers); i++ {
		results <- nil
	}
	return abort, results
}

func (fk *FakeKonsensus) VerifySeal(chain consensus.ChainReader, header *types.Header) error {
	return nil
}

func (fk *FakeKonsensus) Prepare(chain consensus.ChainReader, header *types.Header) error {
	return nil
}

func (fk *FakeKonsensus) Finalize(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, commit *types.Commit, receipts []*types.Receipt) (*types.Block, error) {
	header.Root = state.IntermediateRoot(true)

	// Header seems complete, assemble into a block and return
	return types.NewBlock(header, txs, receipts, commit), nil
}

func (fk *FakeKonsensus) Seal(chain consensus.ChainReader, block *types.Block, stop <-chan struct{}) (*types.Block, error) {
	return nil, nil
}

func (fk *FakeKonsensus) APIs(chain consensus.ChainReader) []rpc.API {
	return nil
}
