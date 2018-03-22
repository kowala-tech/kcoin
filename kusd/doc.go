/*
Package kusd implements a full node service.

Full Node

A full node is a node with a full copy (full/fast syncmodes) of the Kowala
Blockchain.


Archive Node

An archive node (full node) synchronizes the blockchain by downloading the full
chain from the genesis block to the current head block, executing all the
transactions contained within. As the node crunches through the transactions,
all past historical state is stored on disk, and can be queried for each and
every block.


Validator Node

A validator node is a full node with the validator module enabled - active role
on the consensus. There are two types of validators:

* genesis validators - validators that are included in the genesis block.
* non-genesis validators or just validators - validators that just join the
elections after the genesis block.

Note: The genesis validators must use by default the full sync mode. More on
sync below.


Block Synchronization

The current codebase supports two sync modes:

* Fast - Instead of processing the entire blockchain one link at a time, and
replay all transactions that ever happened in history, fast syncing downloads
the transactions receipts along the blocks, and pulls an entire recent state
database.

* Full - Processes the entire blockchain one link at a time, replaying all the
transactions that ever happened in history.

The light sync mode will be implemented soon.

*/

package kusd
