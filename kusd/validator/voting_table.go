package validator

import (
	"github.com/kowala-tech/kUSD/core"
	"github.com/kowala-tech/kUSD/core/types"
	"github.com/kowala-tech/kUSD/event"
)

// VotingTables represents the voting tables available for each election round
type VotingTables = [2]core.VotingTable

func NewVotingTables(eventMux *event.TypeMux, voters types.ValidatorList) VotingTables {
	majorityFunc := func() {
		go eventMux.Post(core.NewMajorityEvent{})
	}
	tables := VotingTables{}
	tables[0] = core.NewVotingTable(types.PreVote, voters, majorityFunc)
	tables[1] = core.NewVotingTable(types.PreCommit, voters, majorityFunc)
	return tables
}
