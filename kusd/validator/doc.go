/*

The validator package implements a consensus state machine.

Validator (think miner in proof-of-work)

Validators are the consensus protocol participants and they are responsible for
committing new blocks in the blockchain. The validators participate in the
consensus protocol by proposing blocks (proposer) and by broadcasting signed
votes.

There are two types of validators:

* genesis validators - included in the genesis block (origin). The genesis
validators don't need to make a deposit for the first block for now (flow might
change in the future).

* non-genesis validators or just validators - join the elections after the
genesis block.


Validator Selection

The kowala network will have a limited number of positions (to be defined)
available for validators - the number of validators might increase over time.
The validators are determined by who has the biggest deposit - the top x number
of validator candidates with the most stake will become Kowala validators.

A limited number of validators is fundamental in order to achieve a consensus in
a short period of time (ex: 1 second) - one of the top priorities for Kowala.
More validators mean more time that is necessary to reach a consensus. We are
actively benchmarking the Kowala blockchain and we will provide specific numbers
as soon as possible. The current values can be found in the network contract:
https://github.com/kowala-tech/kUSD/blob/master/contracts/network/contracts/network.sol.

There are currently two solutions to enable more validators:

* The cosmos network team introduced the 'delegator': candidates can delegate
their own tokens to validators and receive part of the incentives. Validators
may want to establish terms of service and limits on liability for delegators or
have delegators to operate at their own risk.

* by implementing sharding and adding new shards - a new set of validators would
be necessary.


Registering a Validator

In order to join the elections, a validator needs to send a transaction, calling
the deposit method of the network smart contract. The deposit value needs to be
equal or bigger than the minimum deposit if there are open positions for
validators. If there are no positions available he will have to deposit a value
bigger than the current lowest deposit to guarantee a place in the validator
set. On the other hand, if the proposer wants to leave the election he will
submit a transaction, calling the withdraw method of the network smart contract
and the funds will be transferred back to his account. We're actively improving
this scenario.


Proposer

The proposer is the validator node in charge of proposing a new block for a
specific election round. The voting power of each validator is equivalent to the
tokens that he has at stake. The network smart contract contains a list of all
the elections's validators and based on that information, each validator node
will be able to figure out the current proposer - the validator with the most
voting power. The voting power is updated on each election round.

Since Kowala has a dynamic validator set, we set the voting power of all
validators to 0 every time the validator set changes - the new validators are
not aware of the previous voting weights. We're actively looking at this
scenario.


Sync

A validator must be in sync with the network before joining the elections - the
initial sync is a one shot type of update loop. If a user starts the validator
before the sync completes, the start will get delayed until the sync is over.
We're actively studying the sync scenarios during the consensus (ex: a specific
node is delayed).

There's an exception to the rule: a genesis validator with no peers, should be
able to start the validations by himself - he cannot sync since he's the
"creator". We currently rely on a forced sync with no peers to identify the node
as a genesis validator and enable the start of the consensus and acceptance of
transactions by the node. The forced sync is triggered after a certain time
interval(sync.go - forceSyncCycle). Note that the time interval is more than
enough to connect to peers, if they exist.

In order to speedup the sync times, there's a mechanism that disables accepting
transactions until there's a successful sync.


Consensus States

1. Not logged in yet - The validator verifies if a deposit is necessary. If so,
he deposits tokens and subscribes to new block events in order to confirm that
the registration was successful - latest state.

2. New election - the consensus state is initialized and the validators sync
their start time. Note that for block 0 and round 0, the new voting round starts
only after there's a transaction.

3. New election round - the validators voting power is updated.

4. New proposal/timeout - a block is proposed (proposer) / process the block fragments
(non-proposer).

5. Vote - pre-vote sub election.

6. Wait for a majority/timeout.

7. Vote - pre-commit sub election.

8. Wait for a majority/timeout.

9. Commit - the block and state are committed. The validator verifies if is part
of the next election.

10. Logged out - We have left the consensus elections

References:

* Cosmos Network - https://cosmos.network/validators
*/

package validator
