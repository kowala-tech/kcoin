/*

Package consensus implements consensus protocols - the implementations satisfy
the Engine Interface (consensus.go)

The consensus protocols are responsible for operations such as checking whether
a header conforms to the consensus rules, running any post-transaction state
modifications (Ex: Block rewards) and assembling the final block. The kowala
dev/research team is currently focused on Proof-of-Stake protocols.

We're actively improving the protocols - the codebase supports the following
protocols at the moment:

Proof-of-Stake

* Tendermint Consensus Engine - Tendermint is a mostly asynchronous BFT
consensus protocol. Participants in the protocol are called validators and they
take turns on proposing blocks of transactions and voting on them. Assuming that
less than one third of the validators are Byzantine, Tendermint guarantees that
safety will never be violated. Block finality(every block is final) in
Tendermint can be achieved in 1 second.

*/

package consensus
