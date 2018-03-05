/*
Package types implements the blockchain core types

Specifications

This section contains the specifications of the main blockchain types.

Block

A block represents the atomic unit of a blockchain. The block in Kowala is the
collection of relevant pieces of information - block header - together with
information corresponding to transactions and the consensus election - block body.


Block Header

For more information on tries, please check the trie package.
Note that there are several hashes of the root nodes of trie structures in the
block header. The purpose of this construction is to make the protocol
light-client friendly in as many ways as possible - for more information on the
topic, please check the third reference (Ethereum - Design Rationale).

* Number - Number of ancestor blocks. The genesis block has a number of zero.
The number is particularly important for the sync operations - we rely on it to
know if synchronisations are necessary. The number assumes more relevance in proof-
of-stake compared to proof-of-work because the last one relies on difficulty, which
means that there isn't block finality and the blockchain can fork.
(https://github.com/kowala-tech/kUSD/blob/master/kusd/sync.go#L165)

* Parent Hash - is the hash of the previous block's header ("parent block").

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
transactions of the block. This hash allows an efficient and secure verification
of the transactions that compose the block.

* ReceiptHash - hash of the root node of the trie structure populated with the
receipts of each transaction that compose the block.

* Extra - an arbitrary byte array containing data relevant to this block. In
pratical terms this field has been used for example during Ethereum's dao hard
fork to allow fast/light syncers to correctly pick the side they want and is also
used in the clique consensus (currently not available in this codebase) to include
a signature. Kowala is not using this field at the moment but it can be useful in
the future.

* ValidatorsHash - contains a hash of the current set of validators for the
block. Tracking changes in the validator set can be a time consuming task -
especially for a high number of validators - and with this hash we can compare
with previous summaries to see if there have been changes. The network contract
(smart contract) carries the current set of validators and also this hash.
This field is extremely important for the light clients, since the light client
needs to keep track of the validator set in order to verify block headers.

* Time - time at the block inception. Kowala is currently using this value to
synchronise the validators' start time for a new election round.

* Bloom - the bloom filter(space-efficient probabilistic data structure) for
the logs of the block - allows anyone to efficiently search the blockchain for
certain transactions (or watch new blocks for certain transactions). Example:
User who wants to know every single transfer of some specific ERC20 token.
Checking the to address of all transactions ever as well as the data of that
transaction, to see if calls transfer on this contract would take forever!
Instead, because the ERC20 specification shoot of a log for every transfer, you
can just search the blockchain for these logs!


Pending Confirmation

* GasLimit - current limit of gas expenditure per block. This limit defines the
maximum amount of gas (computational effort) that all the trasactions included in
the block can consume. Its purpose is to keep block propagation and processing
time low. Note that this value in bitcoin is constant but it's variable in
Ethereum.

* GasUsed - total gas used in transactions in this block. The fact that the block
can handle a certain limit does not mean that we will have enough transactions
to fill the block.


Block Body

The block body contains the set of transactions that were mined.


Transaction

Transaction refers to the signed data package that stores a message to be sent
from an externally owned account to another account on the blockchain. When you
interact with the  Kowala blockchain, you are executing transactions and
updating its state.

* AccountNonce - represents the number of transactions sent from a given address. The
nonce increases by 1 each time you send a transaction. The nonce is used to
enforce some rules such as: Transactions must be in order. This field prevents
double-spends as the nonce is the order transactions go in. The transaction with
a nonce value of 3 cannot be mined before a transaction with a nonce value of 2.

* Recipient - The address to which we are directing this message.

* Amount - Total kUSD that you want to send. If you are executing a transaction to send
kUSD to another person or a contract, you set this value.

* Payload - is the data field in a transaction - Either a byte string containing the
associated data of the message, on in the case of a contract-creation transaction, the
initialisation code.


Pending Confirmation

* GasLimit - The maximum amount of gas that this transaction can consume.

* Price - If you want to spend less on a transaction, you can do so by lowering the amount
you pay per unit of gas. The price you pay for each unit increases or decreases
how quickly your transaction will be mined.


References

* Ethereum Yellow Paper - https://github.com/ethereum/yellowpaper
* Tendermint Core - https://github.com/tendermint/tendermint
* Ethereum - Design Rationale -
https://github.com/ethereum/wiki/wiki/Design-Rationale
* Bloom example -
https://www.digitalcurrencygrid.com/2018/01/21/my-favorite-line-in-ethereums-java-implementation/

*/
package types
