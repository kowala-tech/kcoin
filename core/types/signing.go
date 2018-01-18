package types

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/crypto"
	"github.com/kowala-tech/kUSD/params"
)

var (
	ErrInvalidSig     = errors.New("invalid v, r, s values")
	errNoSigner       = errors.New("missing signing methods")
	ErrInvalidChainId = errors.New("invalid chain id for signer")

	errAbstractSigner     = errors.New("abstract signer")
	abstractSignerAddress = common.HexToAddress("ffffffffffffffffffffffffffffffffffffffff")
)

// sigCache is used to cache the derived sender and contains
// the signer used to derive it.
type sigCache struct {
	signer Signer
	from   common.Address
}

// MakeSigner returns a Signer based on the given chain config and block number.
func MakeSigner(config *params.ChainConfig, blockNumber *big.Int) Signer {
	var signer Signer
	switch {
	default:
		signer = NewAndromedaSigner(config.ChainID)
	}
	return signer
}

// SignTx signs the transaction using the given signer and private key
func SignTx(tx *Transaction, signer Signer, prv *ecdsa.PrivateKey) (*Transaction, error) {
	var h common.Hash
	if signer.ChainID() == nil {
		h = tx.UnprotectedHash()
	} else {
		h = tx.ProtectedHash(signer.ChainID())
	}

	sig, err := crypto.Sign(h.Bytes(), prv)
	if err != nil {
		return nil, err
	}

	return tx.WithSignature(signer, sig)
}

// SignProposal signs the proposal using the given signer and private key
func SignProposal(proposal *Proposal, signer Signer, prv *ecdsa.PrivateKey) (*Proposal, error) {
	h := proposal.ProtectedHash(signer.ChainID())
	sig, err := crypto.Sign(h[:], prv)
	if err != nil {
		return nil, err
	}

	return proposal.WithSignature(signer, sig)

}

// SignVote signs the vote using the given signer and private key
func SignVote(vote *Vote, signer Signer, prv *ecdsa.PrivateKey) (*Vote, error) {
	h := vote.ProtectedHash(signer.ChainID())
	sig, err := crypto.Sign(h[:], prv)
	if err != nil {
		return nil, err
	}

	return vote.WithSignature(signer, sig)
}

func TxSender(signer Signer, tx *Transaction) (common.Address, error) {
	if sc := tx.from.Load(); sc != nil {
		sigCache := sc.(sigCache)
		// If the signer used to derive from in a previous
		// call is not the same as used current, invalidate
		// the cache.
		if sigCache.signer.Equal(signer) {
			return sigCache.from, nil
		}
	}

	var h common.Hash
	if !tx.Protected() {
		h = tx.UnprotectedHash()
	} else {
		h = tx.ProtectedHash(signer.ChainID())
	}

	addr, err := signer.Sender(h, tx.data.R, tx.data.S, tx.data.V)
	if err != nil {
		return common.Address{}, err
	}

	tx.from.Store(sigCache{signer: signer, from: addr})

	return addr, nil
}

func ProposalSender(signer Signer, proposal *Proposal) (common.Address, error) {
	if sc := proposal.from.Load(); sc != nil {
		sigCache := sc.(sigCache)
		// If the signer used to derive from in a previous
		// call is not the same as used current, invalidate
		// the cache.
		if sigCache.signer.Equal(signer) {
			return sigCache.from, nil
		}
	}

	addr, err := signer.Sender(proposal.ProtectedHash(signer.ChainID()), proposal.data.R, proposal.data.S, proposal.data.V)
	if err != nil {
		return common.Address{}, err
	}

	proposal.from.Store(sigCache{signer: signer, from: addr})

	return addr, nil
}

func VoteSender(signer Signer, proposal *Proposal) (common.Address, error) {
	if sc := proposal.from.Load(); sc != nil {
		sigCache := sc.(sigCache)
		// If the signer used to derive from in a previous
		// call is not the same as used current, invalidate
		// the cache.
		if sigCache.signer.Equal(signer) {
			return sigCache.from, nil
		}
	}

	addr, err := signer.Sender(proposal.ProtectedHash(signer.ChainID()), proposal.data.R, proposal.data.S, proposal.data.V)
	if err != nil {
		return common.Address{}, err
	}

	proposal.from.Store(sigCache{signer: signer, from: addr})

	return addr, nil
}

type Signer interface {
	// PubilcKey returns the public key derived from the signature
	Sender(hash common.Hash, R, S, V *big.Int) (common.Address, error)
	// NewSignature returns a new signature.
	// The signature must be encoded in [R || S || V] format where V is 0 or 1.
	NewSignature(sig []byte) (R, S, V *big.Int, err error)
	// Checks for equality on the signers
	Equal(Signer) bool
	// Returns the current network ID
	ChainID() *big.Int
}

type UnprotectedSigner struct{}

func (s UnprotectedSigner) ChainID() *big.Int { return nil }

func (s UnprotectedSigner) Equal(s2 Signer) bool {
	_, ok := s2.(UnprotectedSigner)
	return ok
}

func (s UnprotectedSigner) Sender(hash common.Hash, R, S, V *big.Int) (common.Address, error) {
	return recoverPlain(hash, R, S, V, true)
}

func (s UnprotectedSigner) NewSignature(sig []byte) (R, S, V *big.Int, err error) {
	if len(sig) != 65 {
		panic(fmt.Sprintf("wrong size for signature: got %d, want 65", len(sig)))
	}

	R = new(big.Int).SetBytes(sig[:32])
	S = new(big.Int).SetBytes(sig[32:64])
	V = new(big.Int).SetBytes([]byte{sig[64] + 27})
	return
}

// AndromedaSigner implements the Signer interface using andromeda's rules
type AndromedaSigner struct {
	chainID, chainIDMul *big.Int
}

func NewAndromedaSigner(chainID *big.Int) *AndromedaSigner {
	if chainID == nil {
		chainID = new(big.Int)
	}
	return &AndromedaSigner{
		chainID:    chainID,
		chainIDMul: new(big.Int).Mul(chainID, big.NewInt(2)),
	}
}

func (s AndromedaSigner) Equal(s2 Signer) bool {
	andromeda, ok := s2.(AndromedaSigner)
	return ok && andromeda.chainID.Cmp(s.chainID) == 0
}

func (s AndromedaSigner) Sender(hash common.Hash, R, S, V *big.Int) (common.Address, error) {
	return recoverPlain(hash, R, S, V, true)
}

// NewSignature returns a new signature. This signature
// needs to be in the [R || S || V] format where V is 0 or 1.
func (s AndromedaSigner) NewSignature(sig []byte) (R, S, V *big.Int, err error) {
	R, S, V, err = UnprotectedSigner{}.NewSignature(sig)
	if err != nil {
		return nil, nil, nil, err
	}

	if s.chainID.Sign() != 0 {
		V = big.NewInt(int64(sig[64] + 35))
		V.Add(V, s.chainIDMul)
	}
	return
}

func (s AndromedaSigner) ChainID() *big.Int {
	return s.chainID
}

// deriveChainID derives the chain id from the given v parameter
func deriveChainID(v *big.Int) *big.Int {
	if v.BitLen() <= 64 {
		v := v.Uint64()
		if v == 27 || v == 28 {
			return new(big.Int)
		}
		return new(big.Int).SetUint64((v - 35) / 2)
	}
	v = new(big.Int).Sub(v, big.NewInt(35))
	return v.Div(v, big.NewInt(2))
}

func recoverPlain(sighash common.Hash, R, S, Vb *big.Int, homestead bool) (common.Address, error) {
	if Vb.BitLen() > 8 {
		return common.Address{}, ErrInvalidSig
	}
	V := byte(Vb.Uint64() - 27)
	if !crypto.ValidateSignatureValues(V, R, S, homestead) {
		return common.Address{}, ErrInvalidSig
	}
	// encode the snature in uncompressed format
	r, s := R.Bytes(), S.Bytes()
	sig := make([]byte, 65)
	copy(sig[32-len(r):32], r)
	copy(sig[64-len(s):64], s)
	sig[64] = V
	// recover the public key from the snature
	pub, err := crypto.Ecrecover(sighash[:], sig)
	if err != nil {
		return common.Address{}, err
	}
	if len(pub) == 0 || pub[0] != 4 {
		return common.Address{}, errors.New("invalid public key")
	}
	var addr common.Address
	copy(addr[:], crypto.Keccak256(pub[1:])[12:])
	return addr, nil
}
