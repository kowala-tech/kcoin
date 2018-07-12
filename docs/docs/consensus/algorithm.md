# Algorithm

The actual process of advancing the blockchain is described here in general
terms. It's quite similar to Tendermint's approach, and the differences are
outlined below.

## Procedure

1.  A new block is proposed by the correct proposer each round and broadcast to
    the other validators. The block includes a batch of transactions from the
    local transaction pool. If a proposer is Byzantine it might broadcast
    different proposals to different validators.

2.  Two phases of voting, called _pre-vote_ and _pre-commit_, then commence. A
    majority (>2/3) of pre-commits for the same block at the same round results
    in a commit. If a validator does not receive a correct proposal within
    ProposalTimeout, it pre-votes for _nil_ instead

In an asynchronous environments with Byzantine validators, a single stage of
voting, where each validator casts only one vote, is not sufficient to ensure
safety. In essence, because validators can act fraudulently, and because there
are no guarantees on message delivery time, a rogue validator can co-ordinate
some validators to commit a value while others, having not seen the commit, go
to a new round, within which they commit a different value. A single stage of
voting allows validators to tell each other what they know about the proposal.
But to tolerate Byzantine faults (which amounts, essentially to lies, fraud,
deceit, etc.), they must also tell each other what they know about what other
validators have professed to know about the proposal. In other words, a second
stage ensures that enough validators witnessed the result of the first stage.

When a validator receives a more than two-thirds pre-votes for a single block,
it has received a signal that the network is prepared to commit the block, and
serves as justification for the validator to sign and broadcast a pre-commit
vote for that block. Sometimes, due to network asynchrony, a validator may not
receive a polka, or there may not have been one. In that case, the validator is
not justified in signing a pre-commit for that block, and must therefore sign
and publish a pre-commit vote for nil. That is, it is considered malicious
behaviour to sign a pre-commit without justification from a polka.

A pre-commit is a vote to actually commit a block. A pre-commit for nil is a
vote to actually move to the next round. If a validator receives more than
two-thirds pre-commits for a single block, it commits that block, computes the
resulting state, and moves on to round 0 at the next height. If a validator
receives more than two-thirds pre-commits for nil, it moves on to the next
round. A pre-vote for a block is thus a vote to prepare the network to commit
the block. A pre-vote for nil is a vote to prepare the network to move to the
next round.

The outcome of a round is either a commit, or a decision to move to the next
round. With a new round comes the next proposer.

</br></br>
