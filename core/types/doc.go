/*

Block

A block represents the atomic unit of a blockchain. The block in Kowala is the
collection of relevant pieces of information - block header - together with
information corresponding to transactions and the consensus election.

Block Header

For more information on tries, please check the trie package.

* Number - Number of ancestor blocks. The genesis block has a number of zero.
The number is particularly important fo for the sync operations - we rely on it to
know if synchronisations are necessary. The number assumes more relevance in proof-
of-stake compared to proof-of-work because the last one relies on difficulty, which
means that there isn't block finality and the blockchain can fork.
(https://github.com/kowala-tech/kUSD/blob/master/kusd/sync.go#L165)

* Coinbase - account registered by the validator(proposer) who's responsible for
the block creation.

* Root - hash of the root node of the state trie, after all transactions are
executed and finalisations applied. The root hash is especially important for
the block validation. Example: As soon as non-validator nodes receive the block,
they need to process the block, and the state root that results from the various
state changes during the block processing must match the received root hash to
make sure that we end up with the same state in the nodes across the network -
https://github.com/kowala-tech/kUSD/blob/master/core/block_validator.go#L94.

* TxHash - hash of the root node of the trie structure populated with the
receipts of each transaction of the block. This hash allows an efficient and
secure verification of the transactions that compose the block.

* Extra - an arbitrary byte array containing data relevant to this block.

* GasLimit - current limit of gas expenditure per block. This limit defines the
maximum amount of gas (computational effort) that all the trasactions included in
the block can consume. Its purpose is to keep block propagation and processing
time low. Note that this value in bitcoin is constant but it's variable in
Ethereum.


Block Body
*



References

* Ethereum Yellow Paper - https://github.com/ethereum/yellowpaper
* Tendermint Core - https://github.com/tendermint/tendermint



*/
package types
