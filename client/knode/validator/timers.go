package validator

import (
	"fmt"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/log"
	"github.com/kowala-tech/kcoin/client/params"
	"math/big"
	"time"
)

type timers struct {
	parentStageCreatedAt time.Time
	round                uint64
	blockNumber          *big.Int

	blockStarts time.Time
	roundStarts time.Time

	proposingDelay time.Duration
	votingDelay    time.Duration

	newElection timer
	proposing   timer
	preVoting   timer
	preCommit   timer
	commit      timer
}

type timer struct {
	endsAt   time.Time
	deadline time.Time
}

type deadlines struct {
	endsAt   *time.Timer
	deadline *time.Timer
}

const millisecondFormat = "2006-01-02T15:04:05.000"

func newTimers(parentStageCreatedAt time.Time, blockNumber *big.Int, round uint64) *timers {
	return (&timers{
		parentStageCreatedAt: parentStageCreatedAt,
		round:                round,
		blockNumber:          blockNumber,

		proposingDelay: time.Duration(params.ProposeDeltaDuration) * time.Millisecond,
		votingDelay:    time.Duration(params.PreVoteDeltaDuration) * time.Millisecond,
	}).setNewElection().
		setProposing().
		setPreVoting().
		setPreCommit().
		setCommit()
}

func (t *timers) setNewElection() *timers {
	if t.blockNumber.Cmp(big.NewInt(0)) == 0 || t.blockNumber.Cmp(big.NewInt(1)) == 0 {
		t.parentStageCreatedAt = time.Now()
	}

	t.parentStageCreatedAt = t.roundStartsAt(t.parentStageCreatedAt, t.round)

	t.newElection = timer{
		endsAt: t.parentStageCreatedAt,
	}

	return t
}

func (t *timers) setProposing() *timers {
	proposeShouldEnds := time.Duration(getProposeDeadline(t.round))
	proposeDeadlineAt := time.Duration(uint64(proposeShouldEnds) - params.TimeoutDeltaDuration)

	t.proposing = timer{
		endsAt:   t.parentStageCreatedAt.Add(proposeShouldEnds * time.Millisecond),
		deadline: t.parentStageCreatedAt.Add(proposeDeadlineAt * time.Millisecond),
	}

	return t
}

func (t *timers) setPreVoting() *timers {
	preVoteShouldEnds := time.Duration(params.PreVoteDuration + t.round*params.PreVoteDeltaDuration)
	preVoteDeadlineAt := time.Duration(uint64(preVoteShouldEnds) - params.TimeoutDeltaDuration)

	t.preVoting = timer{
		endsAt:   t.proposing.endsAt.Add(preVoteShouldEnds * time.Millisecond),
		deadline: t.proposing.endsAt.Add(preVoteDeadlineAt * time.Millisecond),
	}

	return t
}

func (t *timers) setPreCommit() *timers {
	preCommitShouldEnds := time.Duration(params.PreCommitDuration + t.round*params.PreCommitDeltaDuration)
	preCommitDeadlineAt := time.Duration(uint64(preCommitShouldEnds) - params.TimeoutDeltaDuration)

	t.preCommit = timer{
		endsAt:   t.preVoting.endsAt.Add(preCommitShouldEnds * time.Millisecond),
		deadline: t.preVoting.endsAt.Add(preCommitDeadlineAt * time.Millisecond),
	}

	return t
}

func (t *timers) setCommit() *timers {
	commitShouldEnds := time.Duration(params.CommitDuration + t.round*params.CommitDeltaDuration)
	commitDeadlineAt := time.Duration(uint64(commitShouldEnds) - params.TimeoutDeltaDuration)

	t.commit = timer{
		endsAt:   t.preCommit.endsAt.Add(commitShouldEnds * time.Millisecond),
		deadline: t.preCommit.endsAt.Add(commitDeadlineAt * time.Millisecond),
	}

	return t
}

func (t *timers) getProposingTimers() deadlines {
	now := time.Now()
	state := t.proposing

	log.Debug("newProposal state will wait until", "time", state.endsAt.Format(millisecondFormat), "number", t.blockNumber.Int64(), "round", t.round)
	log.Debug("newProposal state will deadline", "time", state.deadline.Format(millisecondFormat), "number", t.blockNumber.Int64(), "round", t.round)

	return deadlines{
		endsAt:   time.NewTimer(state.endsAt.Sub(now)),
		deadline: time.NewTimer(state.deadline.Sub(now)),
	}
}

func (t *timers) getPreVotingTimers() deadlines {
	now := time.Now()
	state := t.preVoting

	t.logStateTimers(state, "preVoting")

	return deadlines{
		endsAt:   time.NewTimer(state.endsAt.Sub(now)),
		deadline: time.NewTimer(state.deadline.Sub(now)),
	}
}

func (t *timers) getPreCommitTimers() deadlines {
	now := time.Now()
	state := t.preCommit

	t.logStateTimers(state, "preCommit")

	return deadlines{
		endsAt:   time.NewTimer(state.endsAt.Sub(now)),
		deadline: time.NewTimer(state.deadline.Sub(now)),
	}
}

func (t *timers) getCommitTimers() deadlines {
	now := time.Now()
	state := t.commit

	t.logStateTimers(state, "commit")

	return deadlines{
		endsAt:   time.NewTimer(state.endsAt.Sub(now)),
		deadline: time.NewTimer(state.deadline.Sub(now)),
	}
}

func (t *timers) WaitProposingDelay() {
	log.Debug("proposing will delay", "duration", t.proposingDelay.Nanoseconds(), "number", t.blockNumber.Int64(), "round", t.round)
	time.Sleep(t.proposingDelay)
}

func (t *timers) WaitVotingDelay() {
	log.Debug("voting will delay", "duration", t.votingDelay.Nanoseconds(), "number", t.blockNumber.Int64(), "round", t.round)
	time.Sleep(t.votingDelay)
}

func (t *timers) getRoundStartTimer() *time.Timer {
	now := time.Now()
	state := t.newElection

	t.logStateTimers(state, "newElection")

	return time.NewTimer(state.endsAt.Sub(now))
}

func (t *timers) roundStartsAt(blockStarts time.Time, round uint64) time.Time {
	interval := time.Second
	if round != 0 {
		interval = time.Duration(params.ProposeDeltaDuration) * time.Millisecond
	}

	// round to params.ProposeDeltaDuration for != 0 round
	blockStarts = common.CeilTimeByModule(blockStarts, interval)

	return roundStartsAtTime(blockStarts, round)
}

func (t *timers) logStateTimers(state timer, name string) {
	log.Debug(fmt.Sprintf("%s state will wait until", name), "time", state.endsAt.Format(millisecondFormat), "number", t.blockNumber.Int64(), "round", t.round)

	if !state.deadline.Equal(time.Time{}) {
		log.Debug(fmt.Sprintf("%s state will deadline", name), "time", state.deadline.Format(millisecondFormat), "number", t.blockNumber.Int64(), "round", t.round)
	}
}

func roundStartsAtTime(blockStarts time.Time, round uint64) time.Time {
	return blockStarts.Add(time.Duration(params.BlockTime*(round+1)) * time.Millisecond)
}

func getProposeDeadline(round uint64) uint64 {
	return params.ProposeDuration + round*params.ProposeDeltaDuration
}
