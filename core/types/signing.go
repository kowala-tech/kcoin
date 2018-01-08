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
func SignTx(tx *Transaction, s Signer, prv *ecdsa.PrivateKey) (*Transaction, error) {
	h := s.Hash(tx)
	sig, err := crypto.Sign(h[:], prv)
	if err != nil {
		return nil, err
	}
	return s.WithSignature(tx, sig)
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
func Sender(signer Signer, tx *Transaction) (common.Address, error) {
	if sc := tx.from.Load(); sc != nil {
		sigCache := sc.(sigCache)
		// If the signer used to derive from in a previous
		// call is not the same as used current, invalidate
		// the cache.
		if sigCache.signer.Equal(signer) {
			return sigCache.from, nil
		}
	}

	pubkey, err := signer.PublicKey(tx)
	if err != nil {
		return common.Address{}, err
	}
	var addr common.Address
	copy(addr[:], crypto.Keccak256(pubkey[1:])[12:])
	tx.from.Store(sigCache{signer: signer, from: addr})
	return addr, nil
}

type Signer interface {
	// Hash returns the rlp encoded hash for signatures
	Hash(tx *Transaction) common.Hash
	// PubilcKey returns the public key derived from the signature
	PublicKey(tx *Transaction) ([]byte, error)
	// WithSignature returns a copy of the transaction with the given signature.
	// The signature must be encoded in [R || S || V] format where V is 0 or 1.
	WithSignature(tx *Transaction, sig []byte) (*Transaction, error)
	// Checks for equality on the signers
	Equal(Signer) bool
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

func (s AndromedaSigner) PublicKey(tx *Transaction) ([]byte, error) {
	if tx.ChainId().Cmp(s.chainID) != 0 {
		return nil, ErrInvalidChainId
	}

	V := byte(new(big.Int).Sub(tx.data.V, s.chainIDMul).Uint64() - 35)
	if !crypto.ValidateSignatureValues(V, tx.data.R, tx.data.S, true) {
		return nil, ErrInvalidSig
	}
	// encode the signature in uncompressed format
	R, S := tx.data.R.Bytes(), tx.data.S.Bytes()
	sig := make([]byte, 65)
	copy(sig[32-len(R):32], R)
	copy(sig[64-len(S):64], S)
	sig[64] = V

	// recover the public key from the signature
	hash := s.Hash(tx)
	pub, err := crypto.Ecrecover(hash[:], sig)
	if err != nil {
		return nil, err
	}
	if len(pub) == 0 || pub[0] != 4 {
		return nil, errors.New("invalid public key")
	}
	return pub, nil
}

// WithSignature returns a new transaction with the given signature. This signature
// needs to be in the [R || S || V] format where V is 0 or 1.
func (s AndromedaSigner) WithSignature(tx *Transaction, sig []byte) (*Transaction, error) {
	if len(sig) != 65 {
		panic(fmt.Sprintf("wrong size for signature: got %d, want 65", len(sig)))
	}

	cpy := &Transaction{data: tx.data}
	cpy.data.R = new(big.Int).SetBytes(sig[:32])
	cpy.data.S = new(big.Int).SetBytes(sig[32:64])
	cpy.data.V = new(big.Int).SetBytes([]byte{sig[64]})
	if s.chainID.Sign() != 0 {
		cpy.data.V = big.NewInt(int64(sig[64] + 35))
		cpy.data.V.Add(cpy.data.V, s.chainIDMul)
	}
	return cpy, nil
}

// Hash returns the hash to be signed by the sender.
// It does not uniquely identify the transaction.
func (s AndromedaSigner) Hash(tx *Transaction) common.Hash {
	return rlpHash([]interface{}{
		tx.data.AccountNonce,
		tx.data.Price,
		tx.data.GasLimit,
		tx.data.Recipient,
		tx.data.Amount,
		tx.data.Payload,
		s.chainID, uint(0), uint(0),
	})
}

// deriveChainId derives the chain id from the given v parameter
func deriveChainId(v *big.Int) *big.Int {
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
