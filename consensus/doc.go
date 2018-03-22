/*

Package consensus implements consensus protocols - the implementations satisfy
the Engine Interface (consensus.go)

Consensus

The consensus protocols are responsible for operations such as checking whether
a header conforms to the consensus rules, running any post-transaction state
modifications (Ex: Block rewards) and assembling the final block. The kowala
dev/research team is currently focused on Proof-of-Stake protocols.

Proof of Work

The algorithm rewards participants who solve cryptographic puzzles in order to
validate transactions and create new blocks. The energy consumption of the
protocol is one of its negative aspects - If things play out a little less
favorably, however, the bitcoin network may draw over 14 Gigawatts of
electricity by 2020, equivalent to the total power generation capacity of a
small country, like Denmark for example.

Proof of Stake

Category of consensus algorithms that depend on a validator's economic stake in
the network. In proof of stake, a set of validators take turns proposing and
voting on the next block, and the weight of each validator's vote depends on the
size of its deposit (i.e. stake). The stake acts as an disincentive for bad actors.

Problems

Some of these protocols suffer from the following problems:

Nothing at stake - participants have nothing to lose by contributing to multiple
blockchain forks, so consensus on a single blockchain is not guaranteed.


Implementations

Tendermint

* Tendermint Consensus Engine - Tendermint is a mostly asynchronous BFT
consensus protocol. Participants in the protocol are called validators and they
take turns on proposing blocks of transactions and voting on them. Assuming that
less than one third of the validators are Byzantine, Tendermint guarantees that
safety will never be violated. Block finality(every block is final) in
Tendermint can be achieved in 1 second.




References

* https://github.com/tendermint/tendermint
* https://medium.com/@jonchoi/ethereum-casper-101-7a851a4f1eb0
* https://motherboard.vice.com/en_us/article/aek3za/bitcoin-could-consume-as-much-electricity-as-denmark-by-2020

*/

package consensus
