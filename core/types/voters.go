package types

import (
	"errors"
	"math/big"

	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/rlp"
	"io"
)

// Validator represents a consensus validator
type Voter struct {
	address common.Address
	deposit uint64
	weight  *big.Int
}

// NewVoter returns a new validator instance
func NewVoter(address common.Address, deposit uint64, weight *big.Int) *Voter {
	return &Voter{
		address: address,
		deposit: deposit,
		weight:  weight,
	}
}

func (val *Voter) Address() common.Address { return val.address }
func (val *Voter) Deposit() uint64         { return val.deposit }
func (val *Voter) Weight() *big.Int        { return val.weight }

func (val *Voter) EncodeRLP(w io.Writer) error {
	w.Write(val.address.Bytes())
	return nil
}

type Voters interface {
	NextProposer() *Voter
	At(i int) *Voter
	Get(addr common.Address) *Voter
	Len() int
	Contains(addr common.Address) bool
	Hash() common.Hash
}

var ErrInvalidParams = errors.New("A validator set needs at least one validator")

func NewVoters(voterList []*Voter) (*voters, error) {
	if len(voterList) == 0 {
		return nil, ErrInvalidParams
	}

	set := &voters{
		voters: voterList,
	}

	return set, nil
}

type voters struct {
	voters []*Voter
}

// Update updates the weight and the proposer based on the set of voters
func (voters *voters) NextProposer() *Voter {
	proposer := voters.voters[0]

	for _, validator := range voters.voters {
		validator.weight = validator.weight.Add(validator.weight, big.NewInt(int64(validator.deposit)))
		if validator.weight.Cmp(proposer.weight) > 0 {
			proposer = validator
		}
	}

	// decrement the validator weight since he has been selected
	proposer.weight.Sub(proposer.weight, big.NewInt(int64(proposer.deposit)))

	return proposer
}

func (voters *voters) At(i int) *Voter {
	if i < 0 || i >= len(voters.voters) {
		return nil
	}
	return voters.voters[i]
}

func (voters *voters) Get(addr common.Address) *Voter {
	for _, validator := range voters.voters {
		if validator.Address() == addr {
			return validator
		}
	}
	return nil
}

func (voters *voters) Len() int {
	return len(voters.voters)
}

func (voters *voters) GetRlp(i int) []byte {
	enc, _ := rlp.EncodeToBytes(voters.voters[i])
	return enc
}

func (voters *voters) Hash() common.Hash {
	return DeriveSha(voters)
}

func (voters *voters) Contains(addr common.Address) bool {
	voter := voters.Get(addr)
	return voter != nil
}

func NewDeposit(amount uint64, timeUnix int64) *Deposit {
	return &Deposit{
		amount:              amount,
		availableAtTimeUnix: timeUnix,
	}
}

// Deposit represents the validator deposits at stake
type Deposit struct {
	amount              uint64
	availableAtTimeUnix int64
}

func (dep *Deposit) Amount() uint64 {
	return dep.amount
}

func (dep *Deposit) AvailableAtTimeUnix() int64 {
	return dep.availableAtTimeUnix
}
