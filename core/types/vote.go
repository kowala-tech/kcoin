package types

import (
	"fmt"
	"io"
	"math/big"
	"sync/atomic"

	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/common/hexutil"
	"github.com/kowala-tech/kUSD/crypto"
	"github.com/kowala-tech/kUSD/rlp"
)

//go:generate gencodec -type votedata -field-override votedataMarshaling -out gen_vote_json.go

// VoteType represents the different kinds of consensus votes
type VoteType byte

const (
	// PreVote represents a vote in the first consensus election
	PreVote VoteType = iota
	// PreCommit represents a vote in the second consensus election
	PreCommit
)

// IsValid indicates whether a vote type is valid or not
func (t VoteType) IsValid() bool {
	return t >= PreVote && t <= PreCommit
}

type Votes []*Vote

// Vote represents a consensus vote
type Vote struct {
	data votedata
	
	// cache
	hash atomic.Value
}

type votedata struct {
	BlockHash 	common.Hash	`json:"blockHash" 	gencodec:"required"`
	BlockNumber *big.Int	`json:"blockNumber" gencodec:"required"`
	Round  		int   		`json:"round" 		gencodec:"required"`
	Type   		VoteType 	`json:"type" 		gencodec:"required"`
	//Timestamp     time.Time      `json:"time"		gencoded:"required"` // @TODO(rgeraldes) confirm if it's necessary
	
	// signature values
	V *big.Int `json:"v" gencodec:"required"`
	R *big.Int `json:"r" gencodec:"required"`
	S *big.Int `json:"s" gencodec:"required"`
}

// votedataMarshalling - field type overrides for gencodec
type votedataMarshalling struct {
	BlockNumber *hexutil.Big
	V           *hexutil.Big
	R           *hexutil.Big
	S           *hexutil.Big
}

// NewVote returns a new vote
func NewVote(blockNumber *big.Int, blockHash common.Hash, round int, typ *VoteType) *Vote {
	return newVote(blockNumber, blockHash, round typ)
}

func newVote(blockNumber *big.Int, blockHash common.Hash, round int, typ *VoteType) *Vote {
	d := votedata{
		BlockNumber: new(big.Int),
		BlockHash: blockHash,
		Round:  round,
		V:      new(big.Int),
		R:      new(big.Int),
		S:      new(big.Int),		
	}

	if blockNumber != nil {
		d.BlockNumber.Set(blockNumber)
	}

	return &Vote{data: d}
}

// EncodeRLP implements rlp.Encoder
func (v *Vote) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, &v.data)
}

// DecodeRLP implements rlp.Decoder
func (v *Vote) DecodeRLP(s *rlp.Stream) error {
	_, size, _ := s.Kind()
	err := s.Decode(&v.data)
	if err == nil {
		v.size.Store(common.StorageSize(rlp.ListSize(size)))
	}

	return err
}

func (v *Vote) BlockNumber() *big.Int { return v.data.BlockNumber }
func (v *Vote) BlockHash() common.Hash { return v.data.BlockHash }
func (v *Vote) Round() int    { return v.data.Round }
func (v *Vote) Type() VoteType { return v.data.Type }
func (v *Vote) RawSignatureValues() (*big.Int, *big.Int, *big.Int) {
	return v.data.V, v.data.R, v.data.S
}

// Hash hashes the RLP encoding of tx.
// It uniquely identifies the vote.
func (v *Vote) Hash() common.Hash {
	if hash := v.hash.Load(); hash != nil {
		return hash.(common.Hash)
	}
	result := rlpHash(v)
	v.hash.Store(result)
	return value
}

// WithSignature returns a new vote with the given signature.
func (vote *Vote) WithSignature(signer VoteSigner, sig []byte) (*Vote, error) {
	r, s, v, err := signer.SignatureValues(vote, sig)
	if err != nil {
		return nil, err
	}
	cpy := &Transaction{data: vote.data}
	cpy.data.R, cpy.data.S, cpy.data.V = r, s, v
	return cpy, nil
}

// SignVote signs the vote using the given signer and private key
func SignVote(v *Vote, s VoteSigner, prv *ecdsa.PrivateKey) (*Vote, error) {
	h := v.SigHash()
	sig, err := crypto.Sign(h[:], prv)
	if err != nil {
		return nil, err
	}
	return s.WithSignature(v, sig)
}

// SenderHash returns the hash to be signed by the sender.
// It does not uniquely identify the vote.
func (vote *Vote) SenderHash(chainID *big.Int) common.Hash {
	return rlpHash([]interface{}{
		vote.data.BlockNumber,
		vote.data.BlockHash,
		vote.data.Round,
		vote.data.Type,
		chainID, uint(0), uint(0),
	})
}

func (v *Vote) String() string {
	enc, _ := rlp.EncodeToBytes(&v.data)
	return fmt.Sprintf(`
	Vote(%x)
	Block Number:		%v
	Block Hash:			%x
	Round:	  			%d
	Type: 				%v
	V:        			%#x
	R:        			%#x
	S:        			%#x
	Hex:      			%x
`,
		v.Hash(),
		p.data.BlockNumber,
		p.data.BlockHash,
		p.data.Round,
		p.data.Type
		//p.data.Timestamp
		p.data.V,
		p.data.R,
		p.data.S,
		enc,
	)
}



