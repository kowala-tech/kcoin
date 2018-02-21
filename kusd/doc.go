/*
Package kusd implements a full node service.

Full Node

A full node is a node with a full copy (full/fast syncmodes) of the Kowala Blockchain.

Archive Node

An archive node (full node) synchronizes the blockchain by downloading the full chain from the genesis block to the current head block,
executing all the transactions contained within. As the node crunches through the transactions, all past historical state is stored on
disk, and can be queried for each and every block.
To run an archive node, download the correct genesis and start kusd with:
1. kusd --datadir=<your_data_dir> init <path_to_genesis> - init the chain with the correct genesis block
2. kusd --networkid=<value_in_the_genesis_file> --datadir=<same_as_above> --cache=1024 --syncmode=full - init the archive node

Note that an archival node can be turned into a validator node by executing the following operations in the console.

Validator Node



Block Synchronization

The current codebase supports two sync modes:

* Fast - Instead of processing the entire blockchain one link at a time, and replay all transactions that ever
happened in history, fast syncing downloads the transactions receipts along the blocks, and pulls an entire
recent state database.

* Full - Processes the entire blockchain one link at a time, replaying all the transactions that ever happened in
history. The genesis validators must use by default the full sync mode.

In order to speedup the sync times, there's a mechanism that rejects transactions until the peer is synced (handler.go - acceptTxs)

*/

package kusd
