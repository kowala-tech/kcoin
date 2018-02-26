/*
Package trie implements Merkle Patricia Tries.

Motivation

Hash trees allow efficient and secure verification of the contents of large data
structures and are an essential part of blockchain.

Example: The block header contains a summary of the entire set of transactions
included in the block. In order to verify if a transaction was included in the
block, we just need the block header - contains the summary of the entire set of
transactions. We don't need the entire set of transactions!

It's possible to build a blockchain without merkle trees, for example, by
creating giant block headers with all the transactions but this option does not
scale, especially for devices such as smartphones. One of the biggest challenges
of running a blockchain node on a smartphone is the bandwidth required.
If we had to the download the full set of transactions, this would required a lot
of bandwidth, especially in peak transactions times.

The merkle trees can also help ensure that data blocks received from other peers
are received undamaged and unaltered, and even to check that the other peers do
not lie and send fake blocks.

Example: The consensus proposal includes a summary of all the block fragments
that are going to be sent. The non-proposer validators verify if the block
fragments are the ones expected.

Block Header

The block header contains three trees:
1. Transactions
2. Receipts (pieces of data showing the effect of each transaction)
3. State

This allows the following queries(examples):
- Has this transaction been included in a particular block? (1)
- Tell me all instances of an event of type X (eg. a crowdfunding contract
reaching its goal) emitted by this address in the past 30 days (2)
- What is the current balance of my account? Does this account exist? (3)

Patricia Trees

Unlike transaction history, however, the state needs to be frequently updated:
the balance and nonce of accounts is often changed, and what’s more, new
accounts are frequently inserted, and keys in storage are frequently inserted
and deleted. What is thus desired is a data structure where we can quickly
calculate the new tree root after an insert, update edit or delete operation,
without recomputing the entire tree. The Patricia tree, in simple terms, is
perhaps the closest that we can come to achieving all of these properties
simultaneously.


References:
* https://blog.ethereum.org/2015/11/15/merkling-in-ethereum/
*/

package trie
