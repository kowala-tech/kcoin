package validator

import (
	"fmt"
	"math/big"
	"time"

	"github.com/kowala-tech/kUSD/accounts"
	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/contracts/network"
	"github.com/kowala-tech/kUSD/core"
	"github.com/kowala-tech/kUSD/core/types"
	"github.com/kowala-tech/kUSD/event"
	"github.com/kowala-tech/kUSD/log"
	"github.com/kowala-tech/kUSD/params"
)

type Election interface {
	Join(walletAccount accounts.WalletAccount, amount uint64) error
	Leave(alletAccount accounts.WalletAccount) error
	SubmitProposal(proposal *types.Proposal) error
	Vote(vote *types.Vote) error
	// @TODO (rgeraldes) - I think that this should be moved to the protocol manager
	AddBlockFragment(blockNumber *big.Int, round uint64, fragment *types.BlockFragment) error
	IsGenesisValidator(address common.Address) (bool, error)
	IsValidator(address common.Address) (bool, error)
	MinimumDeposit() (uint64, error)
	Proposer() common.Address
	Number() *big.Int
	Round() uint64
	SubscribeMajorityEvent(ch chan<- core.NewMajorityEvent) event.Subscription
	SubscribeProposalEvent(ch chan<- core.ProposalEvent) event.Subscription
}

// election retains the consensus state for a specific block election
type election struct {
	network.Election // consensus election

	majorityFeed event.Feed
	proposalFeed event.Feed
	scope        event.SubscriptionScope

	chain    *core.BlockChain
	eventMux *event.TypeMux
	signer   types.Signer

	blockNumber *big.Int
	round       uint64

	validators         types.ValidatorList
	validatorsChecksum [32]byte

	proposal       *types.Proposal
	block          *types.Block
	blockFragments *types.BlockFragments
	votingSystem   *VotingSystem // election votes since round 1

	start time.Time // used to sync the validator nodes
}

func NewElection(contract network.Election, config *params.ChainConfig, blockchain *core.BlockChain, eventMux *event.TypeMux) *election {
	return &election{
		Election: contract,
		chain:    blockchain,
		eventMux: eventMux,
		signer:   types.NewAndromedaSigner(config.ChainID),
	}
}

func (election *election) updateValidators(checksum [32]byte, genesis bool) error {
	validators, err := election.Validators()
	if err != nil {
		return err
	}

	election.validators = validators
	election.validatorsChecksum = checksum

	return nil
}

func (election *election) init() error {
	parent := election.chain.CurrentBlock()

	checksum, err := election.ValidatorsChecksum()
	if err != nil {
		log.Crit("Failed to access the voters checksum", "err", err)
	}

	if election.validatorsChecksum != checksum {
		if err := election.updateValidators(checksum, true); err != nil {
			log.Crit("Failed to update the validator set", "err", err)
		}
	}

	start := time.Unix(parent.Time().Int64(), 0)
	election.start = start.Add(time.Duration(params.BlockTime) * time.Millisecond)
	election.blockNumber = parent.Number().Add(parent.Number(), big.NewInt(1))
	election.round = 0

	election.proposal = nil
	election.block = nil
	election.blockFragments = nil

	election.votingSystem = NewVotingSystem(election.eventMux, election.signer, election.blockNumber, election.validators)

	return nil
}

func (election *election) Vote(vote *types.Vote) error {
	if err := election.votingSystem.Add(vote); err != nil {
		switch err {
		}
	}
	return nil
}

func (election *election) SubmitProposal(proposal *types.Proposal) error {
	log.Info("Received Proposal")

	election.proposal = proposal
	election.blockFragments = types.NewDataSetFromMeta(proposal.BlockMetadata())

	return nil
}

// @TODO (rgeraldes) - move this to a different place
func (election *election) AddBlockFragment(blockNumber *big.Int, round uint64, fragment *types.BlockFragment) error {
	/*
		election.blockFragments.Add(fragment)

		if election.blockFragments.HasAll() {
			block, err := election.blockFragments.Assemble()
			if err != nil {
				log.Crit("Failed to assemble the block", "err", err)
			}

			// Start the parallel header verifier
			nBlocks := 1
			headers := make([]*types.Header, nBlocks)
			seals := make([]bool, nBlocks)
			headers[nBlocks-1] = block.Header()
			seals[nBlocks-1] = true

			abort, results := val.engine.VerifyHeaders(val.chain, headers, seals)
			defer close(abort)

			err = <-results
			if err == nil {
				err = val.chain.Validator().ValidateBody(block)
			}

			parent := val.chain.GetBlock(block.ParentHash(), block.NumberU64()-1)

			// Process block using the parent state as reference point.
			receipts, _, usedGas, err := val.chain.Processor().Process(block, val.state, val.vmConfig)
			if err != nil {
				log.Crit("Failed to process the block", "err", err)
			}
			val.receipts = receipts

			// Validate the state using the default validator
			err = val.chain.Validator().ValidateState(block, parent, val.state, receipts, usedGas)
			if err != nil {
				log.Crit("Failed to validate the state", "err", err)
			}

			val.block = block

			go func() { val.blockCh <- block }()
		}
	*/
	return nil
}

func (election *election) Join(walletAccount accounts.WalletAccount, amount uint64) error {
	tx, err := election.Election.Join(walletAccount, amount)
	if err != nil {
		return err
	}

	if err := txVerifier(tx.Hash(), nil, nil); err != nil {
		return err
	}

	return nil
}

func (election *election) Leave(walletAccount accounts.WalletAccount) error {
	tx, err := election.Election.Leave(walletAccount)
	if err != nil {
		return err
	}

	if err := txVerifier(tx.Hash(), nil, nil); err != nil {
		return err
	}

	return nil
}

func (election *election) Proposer() common.Address {
	return election.validators.Proposer().Address()
}

// confirmTransaction subscribes to the chain head events and
// verifies if a mined transaction failed or not
func txVerifier(hash common.Hash, blockchain *core.BlockChain, db core.DatabaseReader) error {
	chainHeadCh := make(chan core.ChainHeadEvent)
	chainHeadSub := blockchain.SubscribeChainHeadEvent(chainHeadCh)
	defer chainHeadSub.Unsubscribe()

	for {
		select {
		case _, ok := <-chainHeadCh:
			if !ok {
				return nil
			}

			minedTx, _, _, _ := core.GetTransaction(db, hash)
			if minedTx == nil {
				continue
			}

			receipt, _, _, _ := core.GetReceipt(db, hash)
			if receipt.Status == types.ReceiptStatusFailed {
				return fmt.Errorf("transaction failed: %v", hash)
			}
			return nil
		}
	}
}

func (election *election) Number() *big.Int {
	return election.blockNumber
}

func (election *election) Round() uint64 {
	return election.round
}

// SubscribeMajorityEvent registers a subscription of NewMajorityEvent and
// starts sending event to the given channel.
func (election *election) SubscribeMajorityEvent(ch chan<- core.NewMajorityEvent) event.Subscription {
	return election.scope.Track(election.majorityFeed.Subscribe(ch))
}

// SubscribeProposalEvent registers a subscription of NewProposalEvent and
// starts sending event to the given channel.
func (election *election) SubscribeProposalEvent(ch chan<- core.ProposalEvent) event.Subscription {
	return election.scope.Track(election.majorityFeed.Subscribe(ch))
}
