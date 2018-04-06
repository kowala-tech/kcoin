## Consensus Algorithm

1.  A new block is proposed by the correct proposer each round and broadcasted to the other validators. The block includes a batch of transactions from the local transaction pool. If a proposer is byzantine it might broadcast different proposals to different validators.

2.  Two phases of voting called pre-vote and pre-commit. A majority (+2/3) of pre-commits for the same block at the same round results in a commit. If a validator does not receive
    a correct proposal within ProposalTimeout, it pre-votes for nil instead

In asynchronous environments with Byzantine validators, a single stage
of voting, where each validator casts only one vote, is not sufficient to ensure
safety.

In essence, because validators can act fraudulently, and because there
are no guarantees on message delivery time, a rogue validator can co-ordinate
some validators to commit a value while others, having not seen the commit,
go to a new round, within which they commit a different value.

A single stage of voting allows validators to tell each other what they
know about the proposal. But to tolerate Byzantine faults (which amounts,
22
essentially to lies, fraud, deceit, etc.), they must also tell each other what
they know about what other validators have professed to know about the
proposal. In other words, a second stage ensures that enough validators
witnessed the result of the first stage.

When a validator receives a polka (read: more than two-thirds pre-votes
for a single block), it has received a signal that the network is prepared to
commit the block, and serves as justification for the validator to sign and
broadcast a pre-commit vote for that block. Sometimes, due to network
asynchrony, a validator may not receive a polka, or there may not have been
one. In that case, the validator is not justified in signing a pre-commit for
that block, and must therefore sign and publish a pre-commit vote for nil.
That is, it is considered malicious behaviour to sign a pre-commit without
justification from a polka.

A pre-commit is a vote to actually commit a block. A pre-commit for nil
is a vote to actually move to the next round. If a validator receives more than
two-thirds pre-commits for a single block, it commits that block, computes
the resulting state, and moves on to round 0 at the next height. If a validator
receives more than two-thirds pre-commits for nil, it moves on to the next
round.

A pre-vote for a block is thus a vote to prepare the network to commit
the block. A pre-vote for nil is a vote to prepare the network to move to the
next round

The outcome of a round is either a commit, or a decision to move to the next round. With a new round comes the next proposer.

In contrast to algorithms which require a form of leader election, Tendermint
has a new lead0er (the proposer) for each round.

## Consensus Timeouts

The beginning of each round has a weak dependence on synchrony as it
utilizes local clocks to determine when to skip a proposer

* If a proposal is not received in sufficient time, the proposer should be skipped.

After the proposal, rounds proceed in a fully asynchronous manner

@TODO (rgeraldes) - check if we skip automatically to the next round if the proposer does not send the proposal in time.

Clocks do not need to be synced
across validators, as they are reset each time a validator observes votes from
two-thirds or more others.

## Validators

TBD - number of validators.

* Responsible for maintaining a full copy of the replicated state
* Proposing/voting new blocks
* take turns proposing new blocks in rounds (proposer)

Locking mechanism - prevents any malicious coalition of less than one third of the validators from compromising safety.

Tendermint ensures that no two validators commit a different block at the same height, presuming less that one-third of the validators are malicious. This is achieved using a locking mechanism which determines how a validator may pre-vote or pre-commit depending on the previous pre-votes and pre-commits at the same height.

To round-skip safely, a small number of locking rules are introduced which
force validators to justify their votes. While we don’t necessarily require
them to broadcast their justifications in real time, we do expect them to
keep the data, such that it can be brought forth as evidence in the event that
safety is compromised by sufficient Byzantine failures. This accountability
mechanism enables Tendermint to provide stronger guarantees in the face of
such failure than eg. PBFT, which provides no guarantees if a third or more
of the validators are Byzantine.

In essence, a pre-commit must
be justified by a polka, and a validator is considered locked on the last block
it pre-commit. There are two rules of locking:

Prevote-the-Lock: a validator must pre-vote for the block they are
locked on, and propose it if they are the proposer. This prevents validators
from pre-committing one block in one round, and then contributing
to a polka for a different block in the next round, thereby
compromising safety.
• Unlock-on-Polka: a validator may only release a lock after seeing a
polka at a round greater than that at which it locked. This allows
validators to unlock if they pre-committed something the rest of the
network doesn’t want to commit, thereby protecting liveness, but does
it in a way that does not compromise safety, by only allowing unlocking
if there has been a polka in a round after that in which the validator
became locked

pre-commit at a new height until they see a polka.
These rules can be understood more intuitively by way of examples. Consider
four validators, A, B, C, D, and suppose there is a proposal for blockX
at round R. Suppose there is a polka for blockX, but A doesn’t see it, and
pre-commits nil, while the others pre-commit for blockX. Now suppose the
only one to see all pre-commits is D, while the others, say, don’t see D’s
pre-commit (they only see their two pre-commits and A’s pre-commit nil).
D will now commit the block, while the others go to round R + 1. Since any
of the validators might be the new proposer, if they can propose and vote
for any new block, say blockY , then they might commit it and compromise
safety, since D already committed blockX. Note that there isn’t even any
Byzantine behaviour here, just asynchrony!
Locking solves the problem by forcing validators to stick with the block
they pre-committed, since other validators might have committed based on
those pre-commits (as D did in this example). In essence, once more than
two-thirds pre-commit a block in a round, the network is locked on that block,
which is to say it must be impossible to produce a valid polka for a different
block at a higher round. This is direct motivation for Prevote-the-Lock.
Prevote-the-Lock is not sufficient, however. There must be a way to
unlock, lest we sacrifice liveness. Consider a round where A and B precommitted
blockX while C and D pre-committed nil - a split vote. They all
move to the next round, and blockY is proposed, which C and D prevote for. Suppose A is Byzantine, and prevotes for blockY as well (despite being
locked on blockX), resulting in a polka. Suppose B does not see the polka
and pre-commits nil, while A goes off-line and C and D pre-commit blockY .
They move to the next round, but B is still locked on blockX, while C and
D are now locked on blockY , and since A is offline, they can never get a
polka. Hence, we’ve compromised liveness with less than a third (here, only
one) Byzantine validators.
The obvious justification for unlocking is a polka. Once B sees the polka
for blockY (which C and D used to jusitfy their pre-commits for blockY ),
it ought to be able to unlock, and hence pre-commit blockY . This is the
motivation for Unlock-on-Polka, which allows validators to unlock (and precommit
a new block), if they have seen a polka in a round greater than that
in which they locked.

For simplicity, a validator is considered to have locked on nil at round
-1 at each height, so that Unlock-on-Polka implies that a validator cannot
pre-commit at a new height until they see a polka

### Validator registry

In essence, validators must make a security deposit (“they must bond some
stake”) in order to participate in consensus.

Explain limit number of validators.

### Proposer selection

Cycling of proposers is necessary for Byzantine tolerance.

Tendermint preserves safety via the voting and locking mechanisms, and maintains liveness by cycling proposers, so if one won’t process any transactions, others can pick up.

Proposers are ordered via a simple, deterministic round robin, so only
a single proposer is valid for a given round, and every validator knows the
correct proposer.

## Liveness

The network may halt altogether if one-third or more of the validators are offline or partitioned.

TBD - When more than two-thirds pre-commit for block b,
we fire b on channel di
, signalling the commit, and terminating the protocol.

## Other

Consider a correct validator having committed block B at height H and
round R. To commit a block means the validator witnessed pre-commits
for block B in round R from more than two-thirds of validators. Suppose
another block C is committed at height H. We have two options: either it
was committed in round R, or round S > R.
If it was committed in round R, then more than two-thirds of validators
must have pre-committed for it in round R, which means that at least a third
of validators pre-committed for both blocks B and C in round R, which is
clearly Byzantine. Suppose block C was instead committed in round S > R.
30
Since more than two-thirds pre-committed for B, they are locked on B in
round S, and thus must pre-vote for B. To pre-commit for block C, they
must witness a polka for C, which requires more than two-thirds to pre-vote
for C. However, since more than two-thirds are locked on and required to
pre-vote for B, a polka for C would require at least one third of validators to
violate Prevote-the-Lock, which is clearly Byzantine. Thus, to violate state
machine safety, at least one third of validators must be Byzantine. Therefore,
Tendermint satisfies state machine safety when less than a third of validators
are Byzantine.

## References
