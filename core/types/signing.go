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
	h := tx.ProtectedHash(signer.ChainID())
	sig, err := crypto.Sign(h[:], prv)
	if err != nil {
		return nil, err
	}

	return tx.WithSignature(signer, sig)
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

	addr, err := Sender(signer, tx.ProtectedHash(signer.ChainID()), signer.ChainID(), tx.data.V, tx.data.R, tx.data.S)
	if err != nil {
		return common.Address{}, err
	}

	tx.from.Store(sigCache{signer: signer, from: addr})

	return addr, nil
}

// Sender derives the sender from the tx using the signer derivation
// functions.

// Sender returns the address derived from the signature (V, R, S) using secp256k1
// elliptic curve and an error if it failed deriving or upon an incorrect
// signature.
//
// Sender may cache the address, allowing it to be used regardless of
// signing method. The cache is invalidated if the cached signer does
// not match the signer used in the current call.
func Sender(signer Signer, hash common.Hash, chainID, V, R, S *big.Int) (common.Address, error) {
	pubkey, err := signer.PublicKey(hash, chainID, V, R, S)
	if err != nil {
		return common.Address{}, err
	}
	var addr common.Address
	copy(addr[:], crypto.Keccak256(pubkey[1:])[12:])
	return addr, nil
}

type Signer interface {
	// PubilcKey returns the public key derived from the signature
	PublicKey(hash common.Hash, chainID, V, R, S *big.Int) ([]byte, error)
	// NewSignature returns a new signature.
	// The signature must be encoded in [R || S || V] format where V is 0 or 1.
	NewSignature(sig []byte) (V, R, S *big.Int, err error)
	// Checks for equality on the signers
	Equal(Signer) bool
	// Returns the current network ID
	ChainID() *big.Int
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

func (s AndromedaSigner) PublicKey(hash common.Hash, chainID, V, R, S *big.Int) ([]byte, error) {
	if chainID.Cmp(s.chainID) != 0 {
		return nil, ErrInvalidChainId
	}

	rawV := byte(new(big.Int).Sub(V, s.chainIDMul).Uint64() - 35)
	if !crypto.ValidateSignatureValues(rawV, R, S, true) {
		return nil, ErrInvalidSig
	}

	// encode the signature in uncompressed format
	rawR, rawS := R.Bytes(), S.Bytes()
	sig := make([]byte, 65)
	copy(sig[32-len(rawR):32], rawR)
	copy(sig[64-len(rawS):64], rawS)
	sig[64] = rawV

	// recover the public key from the signature
	pub, err := crypto.Ecrecover(hash[:], sig)
	if err != nil {
		return nil, err
	}
	if len(pub) == 0 || pub[0] != 4 {
		return nil, errors.New("invalid public key")
	}
	return pub, nil
}

// NewSignature returns a new signature. This signature
// needs to be in the [R || S || V] format where V is 0 or 1.
func (s AndromedaSigner) NewSignature(sig []byte) (V, R, S *big.Int, err error) {
	if len(sig) != 65 {
		panic(fmt.Sprintf("wrong size for signature: got %d, want 65", len(sig)))
	}

	R = new(big.Int).SetBytes(sig[:32])
	S = new(big.Int).SetBytes(sig[32:64])
	V = new(big.Int).SetBytes([]byte{sig[64]})
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
