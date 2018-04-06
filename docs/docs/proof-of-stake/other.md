# Accountability

A

Futhermore, accountability can be at best eventual in asynchronous networks

\*

Note that it is not possible to cause a violation of safety with two-thirds or fewer byzantine validators using only violations of Unlock on polka - more that a third must violate prevote-the-lock for there to be a polka justifying a commit for the remaining honest votes.

In the case of violating locking rules, following a violation of safety, correct
validators must broadcast all votes they have seen at that height, so that the
evidence can be stitched together. The correct validators, which number
something under two-thirds, were collectively privy to all votes which caused
the two blocks to be committed. Within those votes, if there are not a third
or more validators signing conflicting votes, then there are a third or more
violating Prevote-the-Lock

If a pre-vote or a pre-commit influenced a commit, it must have been seen
by a correct validator. Thus, by collecting all votes, violations of Prevotethe-Lock
can be detected by matching each pre-vote to the most recent precommit
by the same validator, unless there isn’t one.

Similarly, violations of Unlock-on-Polka can be detected by matching each
pre-commit to the polka that justifies it. Note that this means a Byzantine
validator can pre-commit before seeing a polka, and escape accountability if
the appropriate polka eventually occurs. However, such cases cannot actually
contribute to violations of safety if the polka is happening anyways.

The current design provides accountability following a post-crisis broadcast
protocol, but it could be improved to allow accountability in real time.
That is, a commit could be changed to include not just the pre-commits, but
all votes justifying the pre-commits, going all the way back to the beginning
of the height. That way, if safety is violated, the unjustified votes can be
detected immediately.

If a third or more of validators are Byzantine, they can compromise safety
a number of ways, for instance, by proposing two blocks for the same round,
and voting both of them through to commit, or by pre-committing on two
different blocks at the same height but in different rounds by violating the
rules on locking. In each case, there is clear, identifiable evidence that certain
validators misbehaved. In the first instance, they signed two proposals at the
same round, a clear violation of the rules. In the second, they may have prevoted
for a different block in round R than they locked on in R−1, a violation
of the Prevote-the-Lock rule.

## Votes (read again 4.2.2)

If a peer has just entered a
new height, it is sent pre-commits from the previous block, so it may include
them in the next blocks LastCommit if it’s a proposer. If a peer has prevoted
but has yet to pre-commit, or has pre-committed, but has yet to go to
the next round, it is sent pre-votes or pre-commits, respectively. If a peer is
catching up, it is sent the pre-commits for the committed block at its current
height.

# Crisis recovery

In the event of a crisis, such as a fork in the transaction log, or the system
coming to a halt, a traditional consensus system provides little or no
guarantees, and typically requires manual intervention.

Tendermint assures that those responsible for violating safety can be identified,
such that any client who can access at least one honest validator can
discern with cryptographic certainty who the dishonest validators are, and
thereby chose to follow the honest validators onto a new chain with a validator
set excluding those who were Byzantine.

For instance, suppose a third or more validators violate locking rules,
causing two blocks to be committed at height H. The honest validators can
determine who double-signed by gossipping all the votes. At this point, they
cannot use the consensus protocol, because the basic fault assumptions have
been violated. Note that being able to at this point accumulate all votes for
H implies strong assumptions about network connectivity and availability
during the crisis, which, if it cannot be provided by the p2p network, may
require validators use alternative means, such as social media and high availability
services, to communicate evidence. A new blockchain can be started
by the full set of remaining honest nodes, once at least two-thirds of them
have gathered all the evidence.

Alternatively, modifying the Tendermint protocol so that pre-commits
require polka would ensure that those responsible for the fork could be punished
immediately, and would not require an additional publishing period.

More complex uses of Governmint are possible for accommodating various
particularities of crisis, such as permanent crash failures and the compromise
of private keys.

Regardless of how crisis recovery proceeds, its success depends on integration
with clients. If clients do not accept the new blockchain, the service
is effectively offline. Thus, clients must be aware of the rules used by the
particular blockchain to recover. In the cases of safety violation described
above, they must also gather the evidence, determine which validators to
remove, and compute the new state with the remaining validators. In the
case of the liveness violation, they must keep up with Governmint

---

Critical elements of operating
the system in the real world, such as managing validator set changes and
recovering from a crisis, have not yet been discussed.

## Validator set changes

Tendermint Governance module vs Kowala - smart contract for validator set changes

Upgrading the software

a block with the intended effect of updating
the validator set, the application can return a list of validators to update by
specifying their public key and new voting power in response to the EndBlock
message. Validators can be removed by setting their voting power to zero.
This provides a generic means for applications to update the validator set
without having to specify transaction types.

If the block at height H returns an updated validator set, then the block
at height H + 1 will reflect the update. Note, however, that the LastCommit
in block H + 1 must utilize the validator set as it was at H, since it may
contain signatures from a validator that was removed

Changes to voting power are applied for H+1 such that the next proposer
is affected by the update. In particular, the validator that otherwise should
have been the next proposer may be removed. The round robin algorithm
should handle this gracefully, simply moving on to the next proposer in line.

Since the same block is replicated on at least two-thirds of validators, and
the round robin is deterministic, they will all make the same update and
expect the same next proposer.

# P2P

On startup, each Tendermint node receives an initial list of peers to dial.
For each peer, a node maintains a persistent TCP connection over which
multiple subprotocols are multiplexed in a rate-limited fashion

## Software Upgrades

software upgrades on a possibly decentralized network. Software upgrades on the public
Internet are a notoriously challenging operation, requiring careful planning
to maintain backwards compatibility for users that don’t upgrade right away,
and to not upset loyal users of the software by introducing bugs, removing
features, adding complexity, or, perhaps worst of all, updating automatically
without permission.

Upgrades to blockchains are typically differentiated as being soft forks
or hard forks, on account of the scope of the changes. Soft forks are meant
to be backwards compatible, and to use degrees of freedom in the protocol
that may be ignored by users who have not upgraded, but which provide new
features to users which do. Hard forks, on the other hand, are non-backwards
compatible upgrades that, in Bitcoin’s case, may cause violations of safety,
and in Tendermint’s case, cause the system to halt.

Clients should be written with configurable update
parameters, so they can specify whether to update automatically or to require
that they are notified first.

Of course, any software upgrade which is not thoroughly vetted could pose
a danger to the system, and a conservative approach to upgrades should be
taken in general.

# Syncing

Another routine continuously attempts to remove blocks from the pool and add them to the blockchain by validating and executing them, two blocks at a time, against the latest state of the blockchain. Blocks must be validated two blocks at a time because the commit for one block is included as the LastCommit data in the next one.
