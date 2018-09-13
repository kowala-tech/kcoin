# Block

A block represents the atomic unit of a blockchain. The block in Kowala is the
collection of relevant pieces of information - block header - together with
information corresponding to transactions and the consensus election - block
body.

## Block Specification

Note that there are several hashes of the root nodes of trie structures in the
block header. The purpose of this construction is to make the protocol
light-client friendly in as many ways as possible.
Example: TxHash hash allows an efficient and secure verification of the
transactions that compose the block.

### Block Header

| Field          | Description                                                                                                        |
| -------------- | ------------------------------------------------------------------------------------------------------------------ |
| Number         | Number of ancestor blocks. The genesis block has a number of zero                                                  |
| ParentHash     | Hash of the previous block's header ("parent block")                                                               |
| Coinbase       | Account registered by the validator(proposer) who's responsible for the block creation                             |
| Root           | Hash of the root node of the state trie, after all transactions are executed and finalisations applied             |
| TxHash         | Hash of the root node of the trie structure populated with the transactions of the block                           |
| ReceiptHash    | Hash of the root node of the trie structure populated with the receipts of each transaction that compose the block |
| Extra          | Extra - an arbitrary byte array containing data relevant to this block                                             |
| ValidatorsHash | Hash of the current set of validators for the current block                                                        |
| Time           | Time at the block inception                                                                                        |
| Bloom          | The bloom filter(space-efficient probabilistic data structure) for the logs of the block                           |
| ResourceUsage  | Total of computational resource (in compute units) used in transactions in this block                              |

#### Context

- Number, is particularly important for the sync operations - we rely on it to
  know if synchronisations are necessary. The number assumes more relevance in
  proof-of-stake compared to proof-of-work because the last one relies on
  difficulty, which means that there isn't block finality and the blockchain can
  fork. (https://github.com/kowala-tech/kcoin/blob/master/kusd/sync.go#L165)

- Root, is especially important for the block validation. Example: As soon as
  non-validator nodes receive the block, they need to process the block, and the
  state root that results from the various state changes during the block
  processing must match the received root hash to make sure that we end up with
  the same state in the nodes across the network -
  https://github.com/kowala-tech/kUSD/blob/master/core/block_validator.go#L94.

- TxHash, allows an efficient and secure verification of the transactions that
  compose the block.

- In pratical terms, the Extra field, has been used for example during
  Ethereum's dao hard fork to allow fast/light syncers to correctly pick the side
  they want and is also used in the clique consensus (currently not available in
  this codebase) to include a signature. Kowala is not using this field at the
  moment but it can be useful in the future.

- Time Kowala is currently using this value to synchronise the validators upon a
  new election round.

- Bloom, allows anyone to efficiently search the blockchain for certain
  transactions (or watch new blocks for certain transactions). Example:

  - User who wants to know every single transfer of some specific ERC20 token.
    Checking the to address of all transactions ever as well as the data of that
    transaction, to see if calls transfer on this contract would take forever!
    Instead, because the ERC20 specification shoot of a log for every transfer, you
    can just search the blockchain for these logs!

- ResourceUsage - The fact that the block can handle a certain limit does not mean
  that we will have enough transactions to fill the block.

### Block Body

The block body contains the set of transactions that were mined and consensus
related elements.

</br></br>
