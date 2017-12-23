package validator

// ProposalPreEvent is posted when a proposal is promoted by the validator.
type ProposalPreEvent struct{}

// PreVoteMajorityEvent is posted when a party wins majority on the pre vote election.
type PreVoteMajorityEvent struct{}

// PreCommitMajorityEvent is posted when a party wins majority on the pre commit election.
type PreCommitMajorityEvent struct{}

// VoteCountDoneEvent is posted as soon as all the pre commit votes are available.
type VoteCountDoneEvent struct{}
