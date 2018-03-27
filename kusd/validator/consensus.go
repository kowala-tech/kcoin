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
	"github.com/kowala-tech/kUSD/internal/kusdapi"
	"github.com/kowala-tech/kUSD/log"
)

type Election interface {
	Join(walletAccount accounts.WalletAccount, amount uint64) error
	AddProposal(proposal *types.Proposal) error
	AddVote(vote *types.Vote) error
	AddBlockFragment(blockNumber *big.Int, round uint64, fragment *types.BlockFragment) error
	IsGenesisValidator(address common.Address) (bool, error)
	IsValidator(address common.Address) (bool, error)
	MinimumDeposit() (uint64, error)
}

func NewElection(contract network.Election) *election {
	return &election{
		Election: contract,
	}
}

// election retains the consensus state for a specific block election
type election struct {
	network.Election // consensus election
	kusdapi.Backend

	blockNumber *big.Int
	round       uint64

	validators         types.ValidatorList
	validatorsChecksum [32]byte

	proposal       *types.Proposal
	block          *types.Block
	blockFragments *types.BlockFragments
	votingSystem   *VotingSystem // election votes since round 1

	lockedRound uint64
	lockedBlock *types.Block

	start time.Time // used to sync the validator nodes

	commitRound int

	// inputs
	blockCh  chan *types.Block
	majority *event.TypeMuxSubscription

	// state changes related to the election
	*work
}

func (election *election) AddVote(vote *types.Vote) error {
	if err := election.votingSystem.Add(vote); err != nil {
		switch err {
		}
	}
	return nil
}

func (election *election) AddProposal(proposal *types.Proposal) error {
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

	if err := transactionWaiter(tx.Hash(), nil, nil); err != nil {
		return err
	}

	return nil
}

func (election *election) Leave(walletAccount accounts.WalletAccount) error {
	tx, err := election.Election.Leave(walletAccount)
	if err != nil {
		return err
	}

	if err := transactionWaiter(tx.Hash(), nil, nil); err != nil {
		return err
	}

	return nil
}

// transactionWaiter subscribes to the chain head events and
// verifies if a mined transaction failed or not
func transactionWaiter(hash common.Hash, blockchain *core.BlockChain, db core.DatabaseReader) error {
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

// VotingTables represents the voting tables available for each election round
type VotingTables = [2]core.VotingTable

func NewVotingTables(eventMux *event.TypeMux, voters types.ValidatorList) VotingTables {
	majorityFunc := func() {
		go eventMux.Post(core.NewMajorityEvent{})
	}
	tables := VotingTables{}
	tables[0] = core.NewVotingTable(types.PreVote, voters, majorityFunc)
	tables[1] = core.NewVotingTable(types.PreCommit, voters, majorityFunc)
	return tables
}

// VotingSystem records the election votes since round 1
type VotingSystem struct {
	voters         types.ValidatorList
	electionNumber *big.Int // election number
	round          uint64
	votesPerRound  map[uint64]VotingTables
	signer         types.Signer

	eventMux *event.TypeMux
}

// NewVotingSystem returns a new voting system
// @TODO (rgeraldes) - in the future replace eventMux with a subscription method
func NewVotingSystem(eventMux *event.TypeMux, signer types.Signer, electionNumber *big.Int, voters types.ValidatorList) *VotingSystem {
	system := &VotingSystem{
		voters:         voters,
		electionNumber: electionNumber,
		round:          0,
		votesPerRound:  make(map[uint64]VotingTables),
		eventMux:       eventMux,
		signer:         signer,
	}

	system.NewRound()

	return system
}

func (vs *VotingSystem) NewRound() {
	vs.votesPerRound[vs.round] = NewVotingTables(vs.eventMux, vs.voters)
}

// Add registers a vote
func (vs *VotingSystem) Add(vote *types.Vote) error {
	votingTable := vs.getVoteSet(vote.Round(), vote.Type())

	signedVote, err := types.NewSignedVote(vs.signer, vote)
	if err != nil {
		return err
	}

	err = votingTable.Add(signedVote)
	if err != nil {
		return err
	}

	go vs.eventMux.Post(core.NewVoteEvent{Vote: vote})

	return nil
}

func (vs *VotingSystem) getVoteSet(round uint64, voteType types.VoteType) core.VotingTable {
	votingTables, ok := vs.votesPerRound[round]
	if !ok {
		// @TODO (rgeraldes) - critical
		return nil
	}

	return votingTables[int(voteType)]
}
