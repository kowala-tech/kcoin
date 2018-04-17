Topic: Proof-of-Stake pitfalls
Handle: proof-of-stake/pitfalls

### Nothing-at-stake problem

Validators can effectively break safety by voting for multiple conflicting blocks at a given block height without incurring cost for doing so. Native PoS implementations are vulnerable to these attacks. Catastrophically, since there’s no incentive to ever converge on a unique chain and every incentive to sign duplicitously on multiple conflicting chains at once, **the economically optimal strategy becomes voting on as many forks as you can find in order to reap more block rewards**. Below is a diagram that demonstrates this:

<figure>
    <img src="/assets/images/nothing-at-stake.png" >
    <figcaption>Expected value for voting on competing chains is greater than expected value for voting on a single chain in naive PoS design.</figcaption>
</figure>

In Proof-of-Work, the “penalty” for mining on multiple chains is that a miner must split up their physical hashing power (scarce resource) to do this.
This attack vector is mitigated by making sure that a validator candidate pays a deposit up-front and with with in-protocol penalties. Even with penalties there's the possibility of having consensus agents that tries to break safety - byzantine validators.

### Long Range Attacks

The long range attack draws from the right that users have to withdraw their security deposits. A fundamental problem arises with this because it means an attacker can build up a fork from an arbitrarily long range without fear of being slashed. **Once security deposits are unbonded, the incentive not to vote on a long-range fork from some block height ago is removed.** In other words, when more than ⅔ of the validators have unbonded, they could maliciously create a second chain which included the past validator set, which could result in arbitrary alternative transactions.

Long range attacks in PoS are rectified under the weak subjectivity model, which requires the following of new nodes which come onto the network:

* Must be currently bonded.
* Only trust validating nodes that currently have security deposits.
* Unbonded deposits must go through a ‘thawing’ period. After unbonding, tokens need time to ‘thaw’ for weeks to months to account for synchrony assumptions (i.e. delayed messages).
* Forbid reverting back before N blocks, where N is the length of the security deposit. This rule invalidates any long range fork.
* Optionally store the validator set on a PoW chain.

Tendermint adopt a simple locking mechanism (colloquially called ‘freezing’ in Tendermint) which “locks” stake for a certain period of time (weeks to months of ‘thawing’) in order to prevent any malicious coalition of validators from violating safety.

### Cartel formation

A third, final hurdle facing any economic paradigm that’s worth any value faces the very real problem of oligopolies; decentralized protocols with native cryptocurrencies are no exception.
Tendermint relies on extra-protocol governance measures to combat oligopolistic validators. While there are no in-protocol measures for censorship-resistance, the rationale behind relying on out-of-band social information to tackle cartel formation is that the users would eventually and inevitably notice cartels forming, socially gossip about it, and then either abandon or vote to reorganize the blockchain that’s under attack.
