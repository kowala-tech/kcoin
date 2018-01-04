package core

import (
	"math/big"

	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/core/types"
)

type VoteSet struct {
	blockNumber *big.Int
	round       int
	voteType    types.VoteType

	validators    Validators
	votesBitArray common.BitArray
	votes         Votes     // primary votes to share
	majority      common.Hash // first majority seen
	votesPerBlock map[common.Hash]
	majorityPerPeer  map[string]common.Hash
}

func NewVoteSet(blockNumber *big.Int, round int, voteType types.VoteType, validators Validators) {
	nVoters := len(validators)
	return &VoteSet{
		blockNumber: blockNumber,
		round: round,
		voteType: voteType,
		votesBitArray: common.NewBitArray(nVoters),
		votes : make([]*Vote, nVoters),
		majority: nil
		votesPerBlock: make(map[common.Hash]*, nVoters),
		majorityPerPeer: make(map[string]common.Hash),
	}
}


func (set *VoteSet) Add(vote *Vote) (bool, error) {
	
}