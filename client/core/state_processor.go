package core

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/consensus"
	"github.com/kowala-tech/kcoin/client/core/state"
	"github.com/kowala-tech/kcoin/client/core/types"
	"github.com/kowala-tech/kcoin/client/core/vm"
	"github.com/kowala-tech/kcoin/client/crypto"
	"github.com/kowala-tech/kcoin/client/log"
	"github.com/kowala-tech/kcoin/client/params"
)

// StateProcessor is a basic Processor, which takes care of transitioning
// state from one point to another.
//
// StateProcessor implements Processor.
type StateProcessor struct {
	config *params.ChainConfig // Chain configuration options
	bc     *BlockChain         // Canonical block chain
	engine consensus.Engine    // Consensus engine used for block rewards
}

// NewStateProcessor initialises a new StateProcessor.
func NewStateProcessor(config *params.ChainConfig, bc *BlockChain, engine consensus.Engine) *StateProcessor {
	return &StateProcessor{
		config: config,
		bc:     bc,
		engine: engine,
	}
}

// Process processes the state changes according to the Kowala rules by running
// the transaction messages using the statedb and applying any rewards to both
// the processor (coinbase) and any included uncles.
//
// Process returns the receipts and logs accumulated during the process and
// returns the amount of computational resource that was used in the process. If any of the
// transactions failed to execute due to insufficient computational resource it will return an error.
func (p *StateProcessor) Process(block *types.Block, statedb *state.StateDB, cfg vm.Config) (types.Receipts, []*types.Log, uint64, error) {
	var (
		receipts      types.Receipts
		resourceUsage = new(uint64)
		header        = block.Header()
		allLogs       []*types.Log
		crpool        = new(ComputationalResourcePool).AddResource(params.ComputeCapacity)
	)

	// Iterate over and process the individual transactions
	for i, tx := range block.Transactions() {
		statedb.Prepare(tx.Hash(), block.Hash(), i)
		receipt, _, err := ApplyTransaction(p.config, p.bc, nil, crpool, statedb, header, tx, resourceUsage, cfg)
		if err != nil {
			log.Debug("failed StateProcessor.Process", "data", spew.Sdump(
				header.Number,
				header.Root,
				header.ParentHash,
				header.TxHash,
				header.ValidatorsHash,
				header.LastCommitHash,
				tx, usedGas, allLogs))

			return nil, nil, 0, err
		}
		receipts = append(receipts, receipt)
		allLogs = append(allLogs, receipt.Logs...)
	}
	// Finalize the block, applying any consensus engine specific extras (e.g. block rewards)
	p.engine.Finalize(p.bc, header, statedb, block.Transactions(), block.LastCommit(), receipts)

	return receipts, allLogs, *resourceUsage, nil
}

// ApplyTransaction attempts to apply a transaction to the given state database
// and uses the input parameters for its environment. It returns the receipt
// for the transaction, computational resource used and an error if the transaction failed,
// indicating the block was invalid.
func ApplyTransaction(config *params.ChainConfig, bc ChainContext, author *common.Address, crpool *ComputationalResourcePool, statedb *state.StateDB, header *types.Header, tx *types.Transaction, resourceUsage *uint64, cfg vm.Config) (*types.Receipt, uint64, error) {
	msg, err := tx.AsMessage(types.MakeSigner(config, header.Number))
	if err != nil {
		return nil, 0, err
	}
	// Create a new context to be used in the VM environment
	context := NewVMContext(msg, header, bc, author)
	// Create a new environment which holds all relevant information
	// about the transaction and calling mechanisms.
	env := vm.New(context, statedb, config, cfg)
	// Apply the transaction to the current state (included in the env)
	_, resourceUsed, failed, err := ApplyMessage(env, msg, crpool)
	if err != nil {
		return nil, 0, err
	}
	// Update the state with pending changes
	var root []byte
	statedb.Finalise(true)

	*resourceUsage += resourceUsed

	// Create a new receipt for the transaction, storing the intermediate root and computational resource used by the tx
	// based on the eip phase, we're passing wether the root touch-delete accounts.
	receipt := types.NewReceipt(root, failed, *resourceUsage)
	receipt.TxHash = tx.Hash()
	receipt.ResourceUsage = resourceUsed
	// if the transaction created a contract, store the creation address in the receipt.
	if msg.To() == nil {
		receipt.ContractAddress = crypto.CreateAddress(env.Context.Origin, tx.Nonce())
	}
	// Set the receipt logs and create a bloom for filtering
	receipt.Logs = statedb.GetLogs(tx.Hash())
	receipt.Bloom = types.CreateBloom(types.Receipts{receipt})

	return receipt, resourceUsed, err
}
