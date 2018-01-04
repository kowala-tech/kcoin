package types

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/crypto"
	"github.com/kowala-tech/kUSD/params"
)

var (
	ErrInvalidChainID = errors.New("invalid chain id for signer")
	ErrInvalidSig     = errors.New("invalid transaction v, r, s values")
	ErrNoSigner       = errors.New("missing signing methods")
)

// MakeSigner returns a Signer based on the given chain config and block number.
func MakeSigner(config *params.ChainConfig, blockNumber *big.Int) Signer {
	return NewAndromedaSigner(config.ChainId)
}

// sigCache is used to cache the derived sender and contains
// the signer used to derive it.
type sigCache struct {
	signer Signer
	from   common.Address
}

// Signer encapsulates signature handling.
type Signer interface {
	// Sender returns the sender address
	Sender(sigHash common.Hash, V, R, S *big.Int) (common.Address, error)
	// SignatureValues returns the raw R, S, V values corresponding to the
	// given signature.
	SignatureValues(sig []byte) (r, s, v *big.Int, err error)
	// Equal returns true if the given signer is the same as the receiver.
	Equal(Signer) bool
	//
	ChainID() *big.Int
}

// AndromedaSigner
type AndromedaSigner struct {
	chainID, chainIDMul *big.Int
}

// NewAndromedaSigner
func NewAndromedaSigner(chainID *big.Int) *AndromedaSigner {
	if chainID == nil {
		chainID = new(big.Int)
	}
	return &AndromedaSigner{
		chainID:    chainID,
		chainIDMul: new(big.Int).Mul(chainID, big.NewInt(2)),
	}
}

// Sender
func (s *AndromedaSigner) Sender(sigHash common.Hash, V, R, S *big.Int) (common.Address, error) {
	if tx.ChainId().Cmp(s.chainId) != 0 {
		return common.Address{}, ErrInvalidChainID
	}
	V = new(big.Int).Sub(V, s.chainIDMul)
	V.Sub(V, big8)
	return recoverPlain(sigHash, R, S, V, true)
}

// Equal
func (s AndromedaSigner) Equal(s2 Signer) bool {
	signer, ok := s2.(AndromedaSigner)
	return ok && signer.chainId.Cmp(s.chainId) == 0
}

// SignatureValues returns signature values. This signature
// needs to be in the [R || S || V] format where V is 0 or 1.
func (s *AndromedaSigner) SignatureValues(sig []byte) (*big.Int, *big.Int, *big.Int, error) {
	// @TODO (rgeraldes) - review panic
	if len(sig) != 65 {
		panic(fmt.Sprintf("wrong size for signature: got %d, want 65", len(sig)))
	}
	R := new(big.Int).SetBytes(sig[:32])
	S := new(big.Int).SetBytes(sig[32:64])
	V := new(big.Int).SetBytes([]byte{sig[64] + 27})
	return R, S, V, nil
}

/*
// PublicKey
func (s *AndromedaSigner) PublicKey(sigHash common.Hash, V, R, S *big.Int) ([]byte, error) {
	if s.ChainID().Cmp(s.chainID) != 0 {
		return nil, ErrInvalidChainId
	}

	V = byte(new(big.Int).Sub(V, s.chainIDMul).Uint64() - 35)
	if !crypto.ValidateSignatureValues(V, R, S) {
		return nil, ErrInvalidSig
	}
	// encode the signature in uncompressed format
	R, S = R.Bytes(), S.Bytes()
	sig := make([]byte, 65)
	copy(sig[32-len(R):32], R)
	copy(sig[64-len(S):64], S)
	sig[64] = V

	// recover the public key from the signature
	pub, err := crypto.Ecrecover(sigHash[:], sig)
	if err != nil {
		return nil, err
	}
	if len(pub) == 0 || pub[0] != 4 {
		return nil, ErrInvalidPubKey
	}
	return pub, nil
}
*/

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

// recoverPlain
func recoverPlain(sighash common.Hash, R, S, Vb *big.Int) (common.Address, error) {
	if Vb.BitLen() > 8 {
		return common.Address{}, ErrInvalidSig
	}
	V := byte(Vb.Uint64() - 27)
	if !crypto.ValidateSignatureValues(V, R, S) {
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
