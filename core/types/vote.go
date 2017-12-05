package types

import "github.com/kowala-tech/kUSD/common/hexutil"

type Election uint

const (
	PreVote Election = iota
	PreCommit
)

// @TODO - add gencodec details for each vote field and verify if we need overrides

//go:generate gencodec -type Vote -field-override voteMarshaling -out gen_vote_json.go

// Vote represents a pre-vote or a pre-commit vote from validators for consensus
type Vote struct {
	// @TODO(rgeraldes) - analyze
	//ValidatorAddress data.Bytes       `json:"validator_address"`
	//ValidatorIndex   int              `json:"validator_index"`
	//BlockID BlockID  `json:"block_id" gencodec:"required"` // zero if vote is nil.
	Height   int      `json:"height" gencodec:"required"`
	Round    int      `json:"round" gencodec:"required"`
	Election Election `json:"type" gencodec:"required"`
	Sig      []byte   `json:"sig" gencodec:"required"`
}

// field type overrides for gencodec
type voteMarshalling struct {
	Height *hexutil.Big
}

func (t VoteType) IsValid() bool {
	return t >= PreVote && t <= PreCommit
}
