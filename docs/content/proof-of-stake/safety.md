Topic: Network Safety
Handle: proof-of-stake/network-safety

An accountable Byzantine Fault Tolerant algorithm is one that can identify all Byzantine validators when there is a violation of safety.

Committing and finalizing blocks depends on a supermajority — a >⅔ quorum — of all validators signing off on the proposed block. **Note that accountability can only apply when between one-third of two-thirds of validators are Byzantine(malicious)**. If more than two-thirds are Byzantine, they can completely dominate the protocol, and we have no guarantee that a correct validator will receive any evidence of their misdeeds.

BFT systems can **only tolerate up to a ⅓ of failures**, where failures can include arbitrary or malicious behaviour:

* crash fault - if a third or more of validators crash, the network halts, as no validator is able to make progress without hearing from more than two-thirds of the validator set. The network remains available for reads, but no new commits can be made - transactions are not confirmed. As soon as validators come back on-line, they can carry on from where they left in a round. Kowala consensus state-machine employs a write-ahead log, such that a recovered validator can quickly return to the step it was in when it crashed, ensuring it doesn’t accidentally violate a rule. For situations such as this one, the system coming to a halt, a traditional consensus system provides little or no guarantees, and typically requires manual intervention. For instance, increase the number of max validators and add nodes on-the-fly in order to be able to have a >⅔ quorum.

* In a byzantine fault, it can behave arbitrarily (ex: send different and contradictory messages to different peers). Crash faults are easier to handle, as no process can lie to another process. Byzantine failures are more complicated. In a system of 2f + 1 processes, if f are Byzantine, they can co-ordinate to say arbitrary things to the other f + 1 processes. For instance, suppose we are trying to agree on the value of a single bit, and f = 1, so we have N = 3 processes, A, B, and C, where C is Byzantine. C can tell A that the value is 0 and tell B that it’s 1. If A agrees that its 0, and B agrees that its 1, then they will both think they have a majority and commit, thereby violating the safety condition.

<figure>
    <img src="/assets/images/byzantine.png" >
    <figcaption>A Byzantine process, C, tells A one thing and B another, causing them to come to different conclusions about the network. Here, simple majority vote results in a violation of safety due to only a single Byzantine process.</figcaption>
</figure>

Hence, the upper bound on faults tolerated by a Byzantine system is strictly lower than a non-Byzantine one. In fact, it can be shown that the upper limit on f for Byzantine faults is f < N/3. Thus, to tolerate a single Byzantine process, we require at least N = 4. Then the faulty process can’t split the vote the way it was able to when N = 3. Systems which only tolerate crash faults can operate via simple majority rule, and therefore typically tolerate simultaneous failure of up to half of the system. If the number of failures the system can tolerate is f, such systems must have at least 2f + 1 processes.

Given that <⅔ of the validators are byzantine, there are only two ways for violation of safety to occur, and both are accountable:

** 1. Double Signing**

Byzantine proposer makes two conflicting proposals within a round, and byzantine validators vote for both of them. In the event that they are found to double-sign proposals or votes, validators publish evidence of the transgression in the form of a transaction, which the application state can use to change the validator set by removing the transgressor, burning its deposit. This has the effect of associating an explicit economic cost with Byzantine behaviour, and enables one to estimate the cost of violating safety by bribing a third or more of the validators to be Byzantine.

** 2. Locking rules violation **

Byzantine validators violate locking rules after some validators have already committed, causing other validators to commit a different block on a later round.

** Other forms of attacks **

Note that a consensus protocol may specify more behaviours to be punished than just double signing. In particular, we are interested in punishing any strong signalling behaviour which is unjustified - typically, any reported change in state that is not based on the reported state of others. For instance, in a version of Tendermint where all pre-commits must come with the polka that justifies them, validators may be punished for broadcasting unjustified pre-commits. Note, however, that we cannot just punish for any unexpected behaviour - for instance, a validator proposing when it is not their round to propose may be a basis for optimizations which pre-empt asynchrony or crashed nodes. A detailed plan for such situations will be shared soon.

Following a violation of safety, the delayed delivery of critical messages may make it impossible to determine which validators were Byzantine until some time after the safety violation is detected. In fact, if correct processes can receive evidence of Byzantine behaviour, but fail irreversibly before they are able to gossip it, there may be cases where accountability is permanently compromised, though in practice such situations should be surmountable with advanced backup solutions like a write ahead log/transaction journal.

# Crisis recovery

In the event of a crisis, such as a fork in the transaction log, a traditional consensus system provides little or no guarantees, and typically requires manual intervention. Tendermint assures that those responsible for violating safety can be identified, such that any client who can access at least one honest validator can discern with cryptographic certainty who the dishonest validators are, and thereby chose to follow the honest validators onto a new chain with a validator set excluding those who were Byzantine. For instance, suppose a third or more validators violate locking rules, causing two blocks to be committed at height H. The honest validators can determine who double-signed by gossipping all the votes. At this point, they cannot use the consensus protocol, because the basic fault assumptions have been violated. Note that being able to at this point accumulate all votes for H implies strong assumptions about network connectivity and availability during the crisis, which, if it cannot be provided by the p2p network, may require validators use alternative means, such as social media and high availability services, to communicate evidence. A new blockchain can be started by the full set of remaining honest nodes, once at least two-thirds of them have gathered all the evidence. Alternatively, modifying the Tendermint protocol so that pre-commits require polka would ensure that those responsible for the fork could be punished immediately, and would not require an additional publishing period.

There are more scenarios that will be addressed such as permanent crash failures and the compromise of private keys.

Regardless of how crisis recovery proceeds, its success depends on integration with clients. If clients do not accept the new blockchain, the service is effectively offline. Thus, clients must be aware of the rules used by the particular blockchain to recover. In the cases of safety violation described above, they must also gather the evidence, determine which validators to remove, and compute the new state with the remaining validators.
