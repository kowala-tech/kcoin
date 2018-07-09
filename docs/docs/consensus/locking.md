# Locking Mechanism

Prevents any malicious coalition of less than one third of the validators from compromising safety.

Tendermint ensures that no two validators commit a different block at the same height, presuming less that one-third of the validators are malicious. This is achieved using a locking mechanism which determines how a validator may pre-vote or pre-commit depending on the previous pre-votes and pre-commits at the same height.

To round-skip safely, a small number of locking rules are introduced which force validators to justify their votes. While we don’t necessarily require them to broadcast their justifications in real time, we do expect them to keep the data, such that it can be brought forth as evidence in the event that safety is compromised by sufficient Byzantine failures. This accountability mechanism enables Tendermint to provide stronger guarantees in the face of such failure than eg. PBFT, which provides no guarantees if a third or more
of the validators are Byzantine.

In essence, a pre-commit must be justified by a polka, and a validator is considered locked on the last block it pre-commit. There are two rules of locking:

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

</br></br>
