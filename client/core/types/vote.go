package types

import (
	"fmt"
	"io"
	"math/big"
	"sync"
	"sync/atomic"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/common/hexutil"
	"github.com/kowala-tech/kcoin/client/rlp"
)

//go:generate gencodec -type votedata -field-override votedataMarshalling -out gen_vote_json.go

// VoteType represents the different kinds of consensus votes
type VoteType byte

const (
	// PreVote represents a vote in the first consensus election
	PreVote VoteType = iota
	// PreCommit represents a vote in the second consensus election
	PreCommit
)

// IsValid indicates whether a vote type is valid or not
func (typ VoteType) IsValid() bool {
	return typ >= PreVote && typ <= PreCommit
}

type AddressVote interface {
	Address() common.Address
	Vote() *Vote
}

func NewAddressVote(signer Signer, vote *Vote) (*addressVote, error) {
	address, err := VoteSender(signer, vote)
	if err != nil {
		return nil, err
	}

	return &addressVote{
		vote,
		address,
	}, nil
}

type addressVote struct {
	vote    *Vote
	address common.Address
}

func (addressVote *addressVote) Address() common.Address {
	return addressVote.address
}

func (addressVote *addressVote) Vote() *Vote {
	return addressVote.vote
}

// Vote represents a consensus vote
type Vote struct {
	data votedata

	// cache
	hash atomic.Value
	size atomic.Value // @TODO (rgeraldes) - confirm if it's necessary
	from atomic.Value
}

type votedata struct {
	BlockHash   common.Hash `json:"blockHash"    gencodec:"required"`
	BlockNumber *big.Int    `json:"blockNumber"  gencodec:"required"`
	Round       uint64      `json:"round"        gencodec:"required"`
	Type        VoteType    `json:"type"         gencodec:"required"`
	// Timestamp     time.Time      `json:"time"		gencoded:"required"` // @TODO (rgeraldes) confirm if it's necessary

	// signature values
	V *big.Int `json:"v"   gencodec:"required"`
	R *big.Int `json:"r"   gencodec:"required"`
	S *big.Int `json:"s"   gencodec:"required"`
}

// votedataMarshalling - field type overrides for gencodec
type votedataMarshalling struct {
	BlockNumber *hexutil.Big
	Round       hexutil.Uint64
	V           *hexutil.Big
	R           *hexutil.Big
	S           *hexutil.Big
}

// NewVote returns a new consensus vote
func NewVote(blockNumber *big.Int, blockHash common.Hash, round uint64, voteType VoteType) *Vote {
	return newVote(blockNumber, blockHash, round, voteType)
}

func newVote(blockNumber *big.Int, blockHash common.Hash, round uint64, voteType VoteType) *Vote {
	d := votedata{
		BlockNumber: new(big.Int),
		BlockHash:   blockHash,
		Round:       round,
		Type:        voteType,
		V:           new(big.Int),
		R:           new(big.Int),
		S:           new(big.Int),
	}

	if blockNumber != nil {
		d.BlockNumber.Set(blockNumber)
	}

	return &Vote{data: d}
}

// EncodeRLP implements rlp.Encoder
func (vote *Vote) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, &vote.data)
}

// DecodeRLP implements rlp.Decoder
func (vote *Vote) DecodeRLP(s *rlp.Stream) error {
	_, size, _ := s.Kind()
	err := s.Decode(&vote.data)
	if err == nil {
		vote.size.Store(common.StorageSize(rlp.ListSize(size)))
	}

	return err
}

func (vote *Vote) BlockNumber() *big.Int  { return vote.data.BlockNumber }
func (vote *Vote) BlockHash() common.Hash { return vote.data.BlockHash }
func (vote *Vote) Round() uint64          { return vote.data.Round }
func (vote *Vote) Type() VoteType         { return vote.data.Type }
func (vote *Vote) RawSignatureValues() (*big.Int, *big.Int, *big.Int) {
	return vote.data.R, vote.data.S, vote.data.V
}

// Hash hashes the RLP encoding of the vote.
// It uniquely identifies the vote.
func (vote *Vote) Hash() common.Hash {
	if hash := vote.hash.Load(); hash != nil {
		return hash.(common.Hash)
	}
	v := rlpHash(vote)
	vote.hash.Store(v)
	return v
}

// ProtectedHash returns the hash to be signed by the sender.
// It does not uniquely identify the transaction.
func (vote *Vote) ProtectedHash(chainID *big.Int) common.Hash {
	return vote.HashWithData(chainID, uint(0), uint(0))
}

func (vote *Vote) HashWithData(data ...interface{}) common.Hash {
	voteData := []interface{}{
		vote.data.BlockHash,
		vote.data.BlockNumber,
		vote.data.Round,
		vote.data.Type,
	}
	return rlpHash(append(voteData, data...))
}

// WithSignature returns a new vote with the given signature.
// This signature needs to be formatted as described in the yellow paper (v+27).
func (vote *Vote) WithSignature(signer Signer, sig []byte) (*Vote, error) {
	r, s, v, err := signer.SignatureValues(sig)
	if err != nil {
		return nil, err
	}

	cpy := &Vote{data: vote.data}
	cpy.data.R, cpy.data.S, cpy.data.V = r, s, v

	return cpy, nil
}

func (vote *Vote) Protected() bool {
	return true
}

func (vote *Vote) ChainID() *big.Int {
	return deriveChainID(vote.data.V)
}

func (vote *Vote) SignatureValues() (R, S, V *big.Int) {
	R, S, V = vote.data.R, vote.data.S, vote.data.V
	return
}

func (vote *Vote) String() string {
	enc, _ := rlp.EncodeToBytes(&vote.data)
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
		vote.Hash(),
		vote.data.BlockNumber,
		vote.data.BlockHash,
		vote.data.Round,
		vote.data.Type,
		vote.data.V,
		vote.data.R,
		vote.data.S,
		enc,
	)
}

type Votes []*Vote

type VotesSet struct {
	m map[common.Hash]*Vote
	l sync.RWMutex
}

func NewVotesSet() *VotesSet {
	return &VotesSet{m: make(map[common.Hash]*Vote)}
}

func (v *VotesSet) Add(vote *Vote) {
	v.l.Lock()
	v.m[vote.Hash()] = vote
	v.l.Unlock()
}

func (v *VotesSet) Contains(h common.Hash) bool {
	v.l.RLock()
	_, res := v.m[h]
	v.l.RUnlock()
	return res
}

func (v *VotesSet) Len() int {
	v.l.RLock()
	res := len(v.m)
	v.l.RUnlock()
	return res
}

func (v *VotesSet) Get(h common.Hash) (*Vote, bool) {
	v.l.RLock()
	vote, ok := v.m[h]
	v.l.RUnlock()
	return vote, ok
}
