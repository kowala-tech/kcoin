package types

import (
	"errors"
	"math/big"

	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/rlp"
	"io"
)

// Validator represents a consensus Voter
type Voter struct {
	address common.Address
	deposit uint64
	weight  *big.Int
}

// NewVoter returns a new Voter instance
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

// Voters represent a set of voters
// it allows to iterate over to next Proposer
// base on Voter deposit and weight
type Voters interface {
	NextProposer() *Voter
	At(i int) *Voter
	Get(addr common.Address) *Voter
	Len() int
	Contains(addr common.Address) bool
	Hash() common.Hash
}

var ErrInvalidParams = errors.New("voters set needs at least one voter")

// NewVoter validates that a list of voters is valid returning a new type if so
func NewVoters(voterList []*Voter) (*voters, error) {
	if len(voterList) == 0 {
		return nil, ErrInvalidParams
	}

	set := &voters{
		voters: voterList,
	}

	return set, nil
}

// voters is a list of Voter
type voters struct {
	voters []*Voter
}

// NextProposer returns the next proposer based on the round and weight of the each voters
func (voters *voters) NextProposer() *Voter {
	proposer := voters.voters[0]

	for _, voter := range voters.voters {

		// add more chance for each voter to be the next Proposer by adding their deposit amount as weight
		voter.weight = voter.weight.Add(voter.weight, big.NewInt(int64(voter.deposit)))

		if voter.weight.Cmp(proposer.weight) > 0 {
			proposer = voter
		}
	}

	// decrement this Voter weight since he has been selected as next proposer
	proposer.weight.Sub(proposer.weight, big.NewInt(int64(proposer.deposit)))

	return proposer
}

// At returns Voter at position or nil if not found
func (voters *voters) At(i int) *Voter {
	if i < 0 || i >= len(voters.voters) {
		return nil
	}
	return voters.voters[i]
}

// Get returns the Voter at index position, nil if outside boundaries or not found
func (voters *voters) Get(addr common.Address) *Voter {
	for _, voter := range voters.voters {
		if voter.Address() == addr {
			return voter
		}
	}
	return nil
}

// Len returns the amount of voters in this set
// needed for hash thru interface DerivableList interface
func (voters *voters) Len() int {
	return len(voters.voters)
}

// GetRlp returns encoded bytes for one voter
// needed for hash thru interface DerivableList interface
func (voters *voters) GetRlp(i int) []byte {
	enc, _ := rlp.EncodeToBytes(voters.voters[i])
	return enc
}

// Hash returns a unique Hash value for this set of Voters
func (voters *voters) Hash() common.Hash {
	return DeriveSha(voters)
}

// Contains returns is ones Voter address is part of this set
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

// Deposit represents the voter deposits at stake
type Deposit struct {
	amount              uint64
	availableAtTimeUnix int64
}

// Amount at stake
func (dep *Deposit) Amount() uint64 {
	return dep.amount
}

// AvailableAtTimeUnix when this deposit is available to withdraw
func (dep *Deposit) AvailableAtTimeUnix() int64 {
	return dep.availableAtTimeUnix
}
