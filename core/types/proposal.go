package types

import (
	"fmt"
	"io"
	"math/big"
	"sync/atomic"

	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/common/hexutil"
	"github.com/kowala-tech/kUSD/rlp"
)

//go:generate gencodec -type proposaldata -field-override proposaldataMarshaling -out gen_proposal_json.go

// Proposal represents a consensus block proposal
type Proposal struct {
	data proposaldata

	// caches
	hash atomic.Value
	size atomic.Value // @TODO (rgeraldes) - confirm if it's necessary
	from atomic.Value
}

type proposaldata struct {
	BlockNumber   *big.Int    `json:"number"			gencodec:"required"`
	Round         uint64      `json:"round"			gencodec:"required"`
	LockedRound   uint64      `json:"lockedRound"		gencodec:"required"`
	LockedBlock   common.Hash `json:"lockedBlock"		gencodec:"required"`
	BlockMetadata Metadata    `json:"block" 			gencoded:"required"`
	//Timestamp     time.Time      `json:"time"		gencoded:"required"` // @TODO(rgeraldes) confirm if it's necessary

	// signature values
	V *big.Int `json:"v"	gencodec:"required"`
	R *big.Int `json:"r"	gencodec:"required"`
	S *big.Int `json:"s"	gencodec:"required"`
}

// proposaldataMarshalling - field type overrides for gencodec
type proposaldataMarshalling struct {
	BlockNumber *hexutil.Big
	Round       hexutil.Uint64
	LockedRound hexutil.Uint64
	V           *hexutil.Big
	R           *hexutil.Big
	S           *hexutil.Big
}

// NewProposal returns a new proposal
func NewProposal(blockNumber *big.Int, round uint64, blockMetadata Metadata, lockedRound int, lockedBlock common.Hash) *Proposal {
	return newProposal(blockNumber, round, blockMetadata, lockedRound, lockedBlock)
}

func newProposal(blockNumber *big.Int, round uint64, blockMetadata Metadata, lockedRound int, lockedBlock common.Hash) *Proposal {
	d := proposaldata{
		BlockNumber:   new(big.Int),
		BlockMetadata: blockMetadata,
		Round:         round,
		V:             new(big.Int),
		R:             new(big.Int),
		S:             new(big.Int),
	}

	if blockNumber != nil {
		d.BlockNumber.Set(blockNumber)
	}

	return &Proposal{data: d}
}

// EncodeRLP implements rlp.Encoder
func (prop *Proposal) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, &prop.data)
}

// DecodeRLP implements rlp.Decoder
func (prop *Proposal) DecodeRLP(s *rlp.Stream) error {
	_, size, _ := s.Kind()
	err := s.Decode(&prop.data)
	if err == nil {
		prop.size.Store(common.StorageSize(rlp.ListSize(size)))
	}

	return err
}

func (prop *Proposal) BlockNumber() *big.Int    { return prop.data.BlockNumber }
func (prop *Proposal) Round() uint64            { return prop.data.Round }
func (prop *Proposal) LockedRound() uint64      { return prop.data.LockedRound }
func (prop *Proposal) LockedBlock() common.Hash { return prop.data.LockedBlock }
func (prop *Proposal) RawSignatureValues() (*big.Int, *big.Int, *big.Int) {
	return prop.data.V, prop.data.R, prop.data.S
}

//func (p *Proposal) Timestamp() time.Time          { return p.data.Timestamp }
//func (prop *Proposal) BlockMetaData() *core.Metadata { return prop.data.BlockMetaData }

// Hash hashes the RLP encoding of the proposal.
// It uniquely identifies the proposal.
func (prop *Proposal) Hash() common.Hash {
	if hash := prop.hash.Load(); hash != nil {
		return hash.(common.Hash)
	}
	v := rlpHash(prop)
	prop.hash.Store(v)
	return v
}

// ProtectedHash returns the hash to be signed by the sender.
// It does not uniquely identify the transaction.
func (prop *Proposal) ProtectedHash(chainID *big.Int) common.Hash {
	return rlpHash([]interface{}{
		prop.data.BlockNumber,
		prop.data.Round,
		prop.data.BlockMetadata,
		prop.data.LockedRound,
		prop.data.LockedBlock,
		chainID, uint(0), uint(0),
	})
}

// Size returns the proposal size
func (prop *Proposal) Size() common.StorageSize {
	if size := prop.size.Load(); size != nil {
		return size.(common.StorageSize)
	}
	c := writeCounter(0)
	rlp.Encode(&c, &prop.data)
	prop.size.Store(common.StorageSize(c))
	return common.StorageSize(c)
}

// WithSignature returns a new proposal with the given signature.
// This signature needs to be formatted as described in the yellow paper (v+27).
func (proposal *Proposal) WithSignature(signer Signer, sig []byte) (*Proposal, error) {
	cpy := &Proposal{data: proposal.data}
	V, R, S, err := signer.NewSignature(sig)
	if err != nil {
		return nil, err
	}
	cpy.data.V = V
	cpy.data.R = R
	cpy.data.S = S

	return cpy, nil
}

// @TODO (rgeraldes) - add metadata & timestamp
func (prop *Proposal) String() string {
	enc, _ := rlp.EncodeToBytes(&prop.data)
	return fmt.Sprintf(`
	Proposal(%x)
	Block Number:		%v
	Round:	  			%d
	Locked Block:		%x
	Locked Round:		%d
	V:        			%#x
	R:        			%#x
	S:        			%#x
	Hex:      			%x
`,
		prop.Hash(),
		prop.data.BlockNumber,
		prop.data.Round,
		prop.data.LockedBlock,
		prop.data.LockedRound,
		prop.data.V,
		prop.data.R,
		prop.data.S,
		enc,
	)
}
