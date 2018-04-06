# Security

---

# Byzantine Validators

In a crash fault, a process simply halts. In a byzantine fault, it can behave arbitrarily (ex: send different and contradictory messages to different peers). Crash faults are easier to handle, as no process can lie to another process. Byzantine failures are more complicated. In a system of 2f + 1 processes, if f are Byzantine, they can co-ordinate to say arbitrary things to the other f + 1 processes. For instance, suppose we are trying to agree on the value of a single bit, and f = 1, so we have N = 3 processes, A, B, and C, where C is Byzantine. C can tell A that the value is 0 and tell B that it’s 1. If A agrees that its 0, and B agrees that its 1, then they will both think they have a majority and commit, thereby violating the safety condition. Hence, the upper bound on faults tolerated by a Byzantine system is strictly lower than a non-Byzantine one. In fact, it can be shown that the upper limit on f for Byzantine faults is f < N/3. Thus, to tolerate a single Byzantine process, we require at least N = 4. Then the faulty process can’t split the vote the way it was able to when N = 3. Systems which only tolerate crash faults can operate via simple majority rule, and therefore typically tolerate simultaneous failure of up to half of the system. If the number of failures the system can tolerate is f, such systems must have at least 2f + 1 processes.

// @TODO (rgeraldes) - add picture

---

## Accountability

An accountable Byzantine Fault Tolerant algorithm is one that can identify all Byzantine validators when there is a violation of safety. Note that accountability can only apply when between one-third of two-thirds of validators are Byzantine. If more than two-thirds are Byzantine, they can completely dominate the protocol, and we have no guarantee that a correct validator will receive any evidence of their misdeeds.

---

## Attacks

There are only two ways for violation of safety to occur, and both are accountable.

1.  Double Signing

Byzantine Proposer makes two conflicting proposals within a round, and Byzantine validators vote for both of them.

In the case of conflicting proposals and conflicting votes, it is trivial to detect the conflict by receiving both messages, and to identify culprits via their signatures - as soon as an honest validator receives conflicting votes or conflicting proposals he reports the specific validator in the form of a transaction to the validator registry - holds the validator deposit.

In the event that they are found to double-sign proposals or votes, validators publish evidence of the transgression in the form of a transaction, which the application state can use to change the validator set by removing the transgressor, burning its deposit.
This has the effect of associating an explicit economic cost with
Byzantine behaviour, and enables one to estimate the cost of violating safety
by bribing a third or more of the validators to be Byzantine.

2.  Byzantine validators violate locking rules after some validators have already committed, causing other validators to commit a different block on a later round.

---

# Network Liveness

If a third or more of validators crash, the network halts, as no validator is able to make progress without hearing from more than two-thirds of the validator set. The network remains available for reads, but no new commits can be made. As soon as validators come back on-line, they can carry on from where they left in a round. The consensus state-machine should employ a write-ahead log, such that a recovered validator can quickly return to the step it was in when it crashed, ensuring it doesn’t accidentally violate a rule.

---

## Reporting a Byzantine Validator

Following a violation of safety, the delayed delivery of critical messages may make it impossible to determine which validators were Byzantine until some time after the safety violation is detected. In fact, if correct processes can receive evidence of Byzantine behaviour, but fail irreversibly before they are able to gossip it, there may be cases where accountability is permanently compromised, though in practice such situations should be surmountable with advanced backup solutions.

TBD

Note that a consensus protocol may specify more behaviours to be punished
than just double signing. In particular, we are interested in punishing
any strong signalling behaviour which is unjustified - typically, any reported
change in state that is not based on the reported state of others. For instance,
in a version of Tendermint where all pre-commits must come with the polka
that justifies them, validators may be punished for broadcasting unjustified
pre-commits. Note, however, that we cannot just punish for any unexpected
behaviour - for instance, a validator proposing when it is not their round
to propose may be a basis for optimizations which pre-empt asynchrony or
crashed nodes.

In fact, a generalization of Tendermint along these two lines, of 1) looser
forms of justification and 2) allowing validators to propose before their term,
gives rise to a family of protocols similar in nature to that proposed by Vlad
Zamfir, under the guise Casper, as the consensus mechanism for a future version
of ethereum [109]. A more formal account of the relationship between the
protocols, and of the characteristics of anti-Byzantine justifications, remains
for future work.
