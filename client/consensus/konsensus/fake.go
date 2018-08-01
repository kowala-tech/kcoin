package konsensus

func NewFaker() *Konsensus {
	return &Konsensus{}
}

func (ks *Konsensus) Author(header *types.Header) (common.Address, error) {
	return header.Coinbase, nil
}

func (ks *Konsensus) Finalize(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, commit *types.Commit, receipts []*types.Receipt) (*types.Block, error) {
	// commit the final state root
	header.Root = state.IntermediateRoot(true)

	// Header seems complete, assemble into a block and return
	return types.NewBlock(header, txs, receipts, commit), nil
}

func (ks *Konsensus) VerifyHeader(chain consensus.ChainReader, header *types.Header, seal bool) error {
	return nil
}

func (ks *Konsensus) VerifyHeaders(chain consensus.ChainReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {
	// @TODO (rgeraldes) - temporary work around
	abort, results := make(chan struct{}), make(chan error, len(headers))
	for i := 0; i < len(headers); i++ {
		results <- nil
	}
	return abort, results
}

func (ks *Konsensus) VerifySeal(chain consensus.ChainReader, header *types.Header) error {
	return nil
}

func (ks *Konsensus) Prepare(chain consensus.ChainReader, header *types.Header) error {
	return nil
}

func (ks *Konsensus) Seal(chain consensus.ChainReader, block *types.Block, stop <-chan struct{}) (*types.Block, error) {
	return nil, nil
}

func (ks *Konsensus) APIs(chain consensus.ChainReader) []rpc.API {
	return nil
}
