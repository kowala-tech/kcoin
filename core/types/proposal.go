package types

import (
	"crypto/ecdsa"
	"fmt"
	"io"
	"math/big"
	"sync/atomic"

	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/common/hexutil"
	"github.com/kowala-tech/kUSD/core"
	"github.com/kowala-tech/kUSD/crypto"
	"github.com/kowala-tech/kUSD/rlp"
)

//go:generate gencodec -type proposaldata -field-override proposaldataMarshaling -out gen_proposal_json.go

// Proposal defines a block proposal for the consensus.
type Proposal struct {
	data proposaldata
	// caches
	hash atomic.Value
	size atomic.Value
}

type proposaldata struct {
	BlockNumber   *big.Int       `json:"number"		gencodec:"required"`
	BlockMetaData *core.Metadata `json:"block" 		gencoded:"required"`
	Round         int            `json:"round"		gencodec:"required"`
	POLRound      int            `json:"polround"	gencodec:"required"`
	POLBlock      common.Hash    `json: polblock	gencodec:"required"`
	//Timestamp     time.Time      `json:"time"		gencoded:"required"` // @TODO(rgeraldes) confirm if it's necessary

	// signature values
	V *big.Int `json:"v"	gencodec:"required"`
	R *big.Int `json:"r"	gencodec:"required"`
	S *big.Int `json:"s"	gencodec:"required"`
}

// proposaldataMarshalling - field type overrides for gencodec
type proposaldataMarshalling struct {
	BlockNumber *hexutil.Big
	V           *hexutil.Big
	R           *hexutil.Big
	S           *hexutil.Big
}

// NewProposal returns a new proposal
func NewProposal(blockNumber *big.Int, round int, blockMetadata *core.Metadata, polRound int, polBlock common.Hash) *Proposal {
	return newProposal(blockNumber, round, blockMetadata, polRound, polBlock)
}

func newProposal(blockNumber *big.Int, round int, blockData *core.Metadata, polRound int, polBlock common.Hash) *Proposal {
	d := proposaldata{
		BlockNumber: new(big.Int),
		BlockData:   blockData,
		Round:       round,
		V:           new(big.Int),
		R:           new(big.Int),
		S:           new(big.Int),
	}

	if blockNumber != nil {
		d.BlockNumber.Set(blockNumber)
	}

	return &Proposal{data: d}
}

// EncodeRLP implements rlp.Encoder
func (p *Proposal) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, &p.data)
}

// DecodeRLP implements rlp.Decoder
func (p *Proposal) DecodeRLP(s *rlp.Stream) error {
	_, size, _ := s.Kind()
	err := s.Decode(&p.data)
	if err == nil {
		p.size.Store(common.StorageSize(rlp.ListSize(size)))
	}

	return err
}

func (p *Proposal) BlockNumber() *big.Int         { return p.data.BlockNumber }
func (p *Proposal) BlockMetaData() *core.Metadata { return p.data.BlockMetaData }
func (p *Proposal) Round() int                    { return p.data.Round }

//func (p *Proposal) Timestamp() time.Time          { return p.data.Timestamp }
func (p *Proposal) POLRound() int         { return p.data.POLRound }
func (p *Proposal) POLBlock() common.Hash { return p.data.POLBlock }
func (p *Proposal) RawSignatureValues() (*big.Int, *big.Int, *big.Int) {
	return p.data.V, p.data.R, p.data.S
}

// Hash hashes the RLP encoding of tx.
// It uniquely identifies the proposal.
func (p *Proposal) Hash() common.Hash {
	if hash := p.hash.Load(); hash != nil {
		return hash.(common.Hash)
	}
	v := rlpHash(p)
	p.hash.Store(v)
	return v
}

// Size returns the proposal size
func (p *Proposal) Size() common.StorageSize {
	if size := p.size.Load(); size != nil {
		return size.(common.StorageSize)
	}
	c := writeCounter(0)
	rlp.Encode(&c, &p.data)
	p.size.Store(common.StorageSize(c))
	return common.StorageSize(c)
}

// WithSignature returns a new proposal with the given signature.
// This signature needs to be formatted as described in the yellow paper (v+27).
func (p *Proposal) WithSignature(signer ProposalSigner, sig []byte) (*Proposal, error) {
	r, s, v, err := signer.SignatureValues(p, sig)
	if err != nil {
		return nil, err
	}
	cpy := &Proposal{data: p.data}
	cpy.data.R, cpy.data.S, cpy.data.V = r, s, v
	return cpy, nil
}

// SignProposal signs the proposal using the given signer and private key
func SignProposal(p *Proposal, s Signer, prv *ecdsa.PrivateKey) (*Proposal, error) {
	h := s.SigHash(p, s.ChainID())
	sig, err := crypto.Sign(h[:], prv)
	if err != nil {
		return nil, err
	}
	return s.WithSignature(p, sig)
}

// SenderHash returns the hash to be signed by the sender.
// It does not uniquely identify the proposal.
func (p *Proposal) SenderHash(chainID *big.Int) common.Hash {
	return rlpHash([]interface{}{
		p.data.BlockNumber,
		p.data.BlockMetaData,
		p.data.Round,
		//p.data.Timestamp,
		p.data.POLRound,
		p.data.POLRound,
		chainID, uint(0), uint(0),
	})
}

func (p *Proposal) String() string {
	enc, _ := rlp.EncodeToBytes(&p.data)
	return fmt.Sprintf(`
	Proposal(%x)
	Block Number:		%v
	Block Metadata:		%v
	Round:	  			%d
	POLBlock:			%x
	POLRound:			%d
	V:        			%#x
	R:        			%#x
	S:        			%#x
	Hex:      			%x
`,
		p.Hash(),
		p.data.BlockNumber,
		p.data.BlockMetaData,
		p.data.Round,
		//p.data.Timestamp
		p.data.V,
		p.data.R,
		p.data.S,
		enc,
	)
}
