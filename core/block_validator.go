package core

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/kowala-tech/kUSD/common/math"
	"github.com/kowala-tech/kUSD/core/state"
	"github.com/kowala-tech/kUSD/core/types"
	"github.com/kowala-tech/kUSD/params"
)

var (
	// ErrUnknownAncestor is returned when validating a block requires an ancestor
	// that is unknown.
	ErrUnknownAncestor = errors.New("unknown ancestor")
)

// BlockValidator is responsible for validating block headers and
// processed state.
//
// BlockValidator implements Validator.
type BlockValidator struct {
	config *params.ChainConfig // Chain configuration options
	bc     *BlockChain         // Canonical block chain
}

// NewBlockValidator returns a new block validator which is safe for re-use
func NewBlockValidator(config *params.ChainConfig, blockchain *BlockChain) *BlockValidator {
	validator := &BlockValidator{
		config: config,
		bc:     blockchain,
	}
	return validator
}

// ValidateBody verifies the the block header's transaction roots.
// The headers are assumed to be already validated at this point.
func (v *BlockValidator) ValidateBody(block *types.Block) error {
	// Check whether the block's known, and if not, that it's linkable
	if v.bc.HasBlockAndState(block.Hash()) {
		return ErrKnownBlock
	}
	if !v.bc.HasBlockAndState(block.ParentHash()) {
		return ErrUnknownAncestor
	}
	// Header validity is known at this point, check the transactions
	if hash := types.DeriveSha(block.Transactions()); hash != header.TxHash {
		return fmt.Errorf("transaction root hash mismatch: have %x, want %x", hash, header.TxHash)
	}
	return nil
}

// ValidateState validates the various changes that happen after a state
// transition, such as amount of used gas, the receipt roots and the state root
// itself. ValidateState returns a database batch if the validation was a success
// otherwise nil and an error is returned.
func (v *BlockValidator) ValidateState(block, parent *types.Block, statedb *state.StateDB, receipts types.Receipts, usedGas *big.Int) error {
	header := block.Header()
	if block.GasUsed().Cmp(usedGas) != 0 {
		return fmt.Errorf("invalid gas used (remote: %v local: %v)", block.GasUsed(), usedGas)
	}
	// Validate the received block's bloom with the one derived from the generated receipts.
	// For valid blocks this should always validate to true.
	rbloom := types.CreateBloom(receipts)
	if rbloom != header.Bloom {
		return fmt.Errorf("invalid bloom (remote: %x  local: %x)", header.Bloom, rbloom)
	}
	// Tre receipt Trie's root (R = (Tr [[H1, R1], ... [Hn, R1]]))
	receiptSha := types.DeriveSha(receipts)
	if receiptSha != header.ReceiptHash {
		return fmt.Errorf("invalid receipt root hash (remote: %x local: %x)", header.ReceiptHash, receiptSha)
	}
	// Validate the state root against the received state root and throw
	// an error if they don't match.
	if root := statedb.IntermediateRoot(true); header.Root != root {
		return fmt.Errorf("invalid merkle root (remote: %x local: %x)", header.Root, root)
	}
	return nil
}

// CalcGasLimit computes the gas limit of the next block after parent.
// The result may be modified by the caller.
// This is miner strategy, not consensus protocol.
func CalcGasLimit(parent *types.Block) *big.Int {
	// contrib = (parentGasUsed * 3 / 2) / 1024
	contrib := new(big.Int).Mul(parent.GasUsed(), big.NewInt(3))
	contrib = contrib.Div(contrib, big.NewInt(2))
	contrib = contrib.Div(contrib, params.GasLimitBoundDivisor)

	// decay = parentGasLimit / 1024 -1
	decay := new(big.Int).Div(parent.GasLimit(), params.GasLimitBoundDivisor)
	decay.Sub(decay, big.NewInt(1))

	/*
		strategy: gasLimit of block-to-mine is set based on parent's
		gasUsed value.  if parentGasUsed > parentGasLimit * (2/3) then we
		increase it, otherwise lower it (or leave it unchanged if it's right
		at that usage) the amount increased/decreased depends on how far away
		from parentGasLimit * (2/3) parentGasUsed is.
	*/
	gl := new(big.Int).Sub(parent.GasLimit(), decay)
	gl = gl.Add(gl, contrib)
	gl.Set(math.BigMax(gl, params.MinGasLimit))

	// however, if we're now below the target (TargetGasLimit) we increase the
	// limit as much as we can (parentGasLimit / 1024 -1)
	if gl.Cmp(params.TargetGasLimit) < 0 {
		gl.Add(parent.GasLimit(), decay)
		gl.Set(math.BigMin(gl, params.TargetGasLimit))
	}
	return gl
}
