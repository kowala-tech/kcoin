package types

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/crypto"
	"github.com/kowala-tech/kcoin/params"
)

var (
	ErrInvalidSig = errors.New("invalid v, r, s values")

	big8 = big.NewInt(8)
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

	var hash common.Hash
	V := tx.data.V
	if tx.Protected() {
		V = new(big.Int).Sub(V, signer.ChainIDMul())
		V.Sub(V, big8)
		hash = tx.ProtectedHash(signer.ChainID())
	} else {
		hash = tx.Hash()
	}

	addr, err := signer.Sender(hash, tx.data.R, tx.data.S, tx.data.V)
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

	V := new(big.Int).Sub(proposal.data.V, signer.ChainIDMul())
	V.Sub(V, big8)
	addr, err := signer.Sender(proposal.ProtectedHash(signer.ChainID()), proposal.data.R, proposal.data.S, V)
	if err != nil {
		return common.Address{}, err
	}

	proposal.from.Store(sigCache{signer: signer, from: addr})

	return addr, nil
}

func VoteSender(signer Signer, vote *Vote) (common.Address, error) {
	if sc := vote.from.Load(); sc != nil {
		sigCache := sc.(sigCache)
		// If the signer used to derive from in a previous
		// call is not the same as used current, invalidate
		// the cache.
		if sigCache.signer.Equal(signer) {
			return sigCache.from, nil
		}
	}

	V := new(big.Int).Sub(vote.data.V, signer.ChainIDMul())
	V.Sub(V, big8)
	addr, err := signer.Sender(vote.ProtectedHash(signer.ChainID()), vote.data.R, vote.data.S, V)
	if err != nil {
		return common.Address{}, err
	}

	vote.from.Store(sigCache{signer: signer, from: addr})

	return addr, nil
}

type Signer interface {
	// PubilcKey returns the public key derived from the signature
	Sender(hash common.Hash, R, S, V *big.Int) (common.Address, error)
	// SignatureValues returns the raw R, S, V values corresponding to the
	// given signature.
	SignatureValues(sig []byte) (R, S, V *big.Int, err error)
	// Checks for equality on the signers
	Equal(Signer) bool
	// Returns the current network ID
	ChainID() *big.Int
	// Returns the current network ID
	ChainIDMul() *big.Int
}

type UnprotectedSigner struct{}

func (s UnprotectedSigner) ChainID() *big.Int    { return nil }
func (s UnprotectedSigner) ChainIDMul() *big.Int { return nil }

func (s UnprotectedSigner) Equal(s2 Signer) bool {
	_, ok := s2.(UnprotectedSigner)
	return ok
}

func (s UnprotectedSigner) Sender(hash common.Hash, R, S, V *big.Int) (common.Address, error) {
	return recoverPlain(hash, R, S, V, true)
}

func (s UnprotectedSigner) SignatureValues(sig []byte) (R, S, V *big.Int, err error) {
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

// SignatureValues returns a new signature. This signature
// needs to be in the [R || S || V] format where V is 0 or 1.
func (s AndromedaSigner) SignatureValues(sig []byte) (R, S, V *big.Int, err error) {
	R, S, V, err = UnprotectedSigner{}.SignatureValues(sig)
	if err != nil {
		return nil, nil, nil, err
	}

	if s.chainID.Sign() != 0 {
		V = big.NewInt(int64(sig[64] + 35))
		V.Add(V, s.chainIDMul)
	}

	return R, S, V, nil
}

func (s AndromedaSigner) ChainID() *big.Int {
	return s.chainID
}

func (s AndromedaSigner) ChainIDMul() *big.Int {
	return s.chainIDMul
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
	// recover the public key from the signature
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
