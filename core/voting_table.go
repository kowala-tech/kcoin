package core

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/core/types"
	"github.com/kowala-tech/kUSD/event"
	"github.com/kowala-tech/kUSD/log"
)

/*
	VoteSet helps collect signatures from validators at each height + round for a
	predefined vote type.

	We need VoteSet to be able to keep track of conflicting votes when validators
	double-sign.  Yet, we can't keep track of *all* the votes seen, as that could
	be a DoS attack vector.

	There are two storage areas for votes.
	1. voteSet.votes
	2. voteSet.votesByBlock

	`.votes` is the "canonical" list of votes.  It always has at least one vote,
	if a vote from a validator had been seen at all.  Usually it keeps track of
	the first vote seen, but when a 2/3 majority is found, votes for that get
	priority and are copied over from `.votesByBlock`.

	`.votesByBlock` keeps track of a list of votes for a particular block.  There
	are two ways a &blockVotes{} gets created in `.votesByBlock`.
	1. the first vote seen by a validator was for the particular block.
	2. a peer claims to have seen 2/3 majority for the particular block.

	Since the first vote from a validator will always get added in `.votesByBlock`
	, all votes in `.votes` will have a corresponding entry in `.votesByBlock`.

	When a &blockVotes{} in `.votesByBlock` reaches a 2/3 majority quorum, its
	votes are copied into `.votes`.

	All this is memory bounded because conflicting votes only get added if a peer
	told us to track that block, each peer only gets to tell us 1 such block, and,
	there's only a limited number of peers.

	NOTE: Assumes that the sum total of voting power does not exceed MaxUInt64.
*/
// Voting table stores the votes of an election round

type Votes struct {
	Voters *common.BitArray
	Votes  []*types.Vote
	// @TODO (rgeraldes) - review power
	power uint64
}

func (votes *Votes) Power() uint64 {
	return votes.power
}

type VotingTable struct {
	mtx sync.Mutex

	all map[common.Hash]*types.Vote // allow lookups

	blockNumber *big.Int
	round       uint64
	voteType    types.VoteType

	voters        *types.ValidatorSet
	received      *common.BitArray
	votes         []*types.Vote // Primary votes to share
	sum           int           // Sum of voting power for seen votes, discounting conflicts
	votesPerBlock map[common.Hash]*Votes

	signer types.Signer

	quorum int

	// events
	eventMux *event.TypeMux

	// cache
	addressToIndex map[common.Address]int

	//maj23         *common.Hash  // First 2/3 majority seen
	//peerMaj23s map[string]common.Hash // Maj23 for each peer
}

func NewVotingTable(eventMux *event.TypeMux, signer types.Signer, blockNumber *big.Int, round uint64, voteType types.VoteType, voters *types.ValidatorSet) *VotingTable {
	table := &VotingTable{
		blockNumber:   blockNumber,
		round:         round,
		voteType:      voteType,
		voters:        voters,
		received:      common.NewBitArray(uint64(1 /*validators.Size()*/)), // @TODO (rgeraldes)
		votes:         make([]*types.Vote, voters.Size()),
		sum:           0,
		all:           make(map[common.Hash]*types.Vote),
		votesPerBlock: make(map[common.Hash]*Votes, voters.Size()),
		eventMux:      eventMux,
		signer:        signer,
		// @TODO (rgeraldes) - replace with constants
		quorum:         voters.Size()*2/3 + 1,
		addressToIndex: make(map[common.Address]int, voters.Size()),
		//maj23:         nil,
		//peerMaj23s:    make(map[string]BlockID),
	}

	// cache voter index
	for i := 0; i < table.voters.Size(); i++ {
		table.addressToIndex[table.voters.AtIndex(i).Address()] = i
	}

	return table
}

func (table *VotingTable) validateVote(vote *types.Vote, local bool) error {

	/*

		valIndex := vote.ValidatorIndex
		valAddr := vote.ValidatorAddress
		blockKey := vote.BlockID.Key()


				// Ensure that validator index was set
				if valIndex < 0 {
					return false, errors.Wrap(ErrVoteInvalidValidatorIndex, "Index < 0")
				} else if len(valAddr) == 0 {
					return false, errors.Wrap(ErrVoteInvalidValidatorAddress, "Empty address")
				}

				// Make sure the step matches.
				if (vote.Height != voteSet.height) ||
					(vote.Round != voteSet.round) ||
					(vote.Type != voteSet.type_) {
					return false, errors.Wrapf(ErrVoteUnexpectedStep, "Got %d/%d/%d, expected %d/%d/%d",
						voteSet.height, voteSet.round, voteSet.type_,
						vote.Height, vote.Round, vote.Type)
				}


			// Ensure that signer is a validator.
			lookupAddr, val := voteSet.valSet.GetByIndex(valIndex)
			if val == nil {
				return false, errors.Wrapf(ErrVoteInvalidValidatorIndex,
					"Cannot find validator %d in valSet of size %d", valIndex, voteSet.valSet.Size())
			}

			// Ensure that the signer has the right address
			if !bytes.Equal(valAddr, lookupAddr) {
				return false, errors.Wrapf(ErrVoteInvalidValidatorAddress,
					"vote.ValidatorAddress (%X) does not match address (%X) for vote.ValidatorIndex (%d)",
					valAddr, lookupAddr, valIndex)
			}
	*/

	return nil
}

func (table *VotingTable) Add(vote *types.Vote, local bool) (bool, error) {
	//table.lock.Lock()
	//defer table.lock.Unlock()

	// If the vote is already known, discard it
	hash := vote.Hash()
	if table.all[hash] != nil {
		log.Trace("Discarding already known vote", "hash", hash)
		return false, fmt.Errorf("known vote: %x", hash)
	}
	// If the transaction fails basic validation, discard it
	if err := table.validateVote(vote, local); err != nil {
		log.Trace("Discarding invalid transaction", "hash", hash, "err", err)
		invalidTxCounter.Inc(1)
		return false, err
	}

	/*
		// Check signature.
		if err := vote.Verify(voteSet.chainID, val.PubKey); err != nil {
			return false, errors.Wrapf(err, "Failed to verify vote with ChainID %s and PubKey %s", voteSet.chainID, val.PubKey)
		}
	*/

	//added, conflict := table.add(vote)
	added, err := table.add(vote)
	if err != nil {
		return false, err
	}

	if added {
		go table.eventMux.Post(NewVoteEvent{Vote: vote})
	}

	/*
		// Add vote and get conflicting vote if any
		added, conflicting := voteSet.addVerifiedVote(vote, blockKey, val.VotingPower)
		if conflicting != nil {
			return added, NewConflictingVoteError(val, conflicting, vote)
		} else {
			if !added {
				cmn.PanicSanity("Expected to add non-conflicting vote")
			}
			return added, nil
		}
	*/
	return true, nil
}

func (table *VotingTable) add(vote *types.Vote) (bool, error) {
	// check signature
	from, err := types.VoteSender(table.signer, vote) // already validated & cached
	if err != nil {
		return false, err
	}

	index, ok := table.addressToIndex[from]
	if !ok {
		// @TODO (rgeraldes)
	}

	if vote := table.votes[index]; vote != nil {
		// @TODO (rgeraldes) - complete conflict code
	} else {
		table.votes[index] = vote
		//table.received.Set(0)
		table.sum++
	}

	if table.sum == table.quorum {
		go table.eventMux.Post(NewMajorityEvent{})
	}

	/*

		votes, ok := table.votesPerBlock[vote.BlockHash()]
		if ok {
			// @TODO (rgeraldes) - complete conflict code
		} else {
			// @TODO (rgeraldes) - complete conflict code
		}


			if ok := votes.Votes[index]; ok != nil {
				// @TODO (rgeraldes) - complete
				votes.Votes[index] = vote
			}
	*/

	return true, nil
}
