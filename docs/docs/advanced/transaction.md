# Transaction

To satisfy integrity we forbid a correct validator from proposing a block or pre-committing for a block containing a batch of transactions that has already been comitted.
A transaction persists on the tx pool until is committed.
A given transaction is only valid once. This is done using sequence numbers (Nonce)

Once a block is committed, all transactions included in the
block are removed from the mempool, and the remaining transactions are
re-validated by the application logic, as their validity may have changed on
account of other transactions being committed, which the node may not have
had in its mempool.

---

## Transaction structure

Transaction refers to the signed data package that stores a message to be sent
from an externally owned account to another account on the blockchain. When you
interact with the Kowala blockchain, you are executing transactions and
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

## References

* Ethereum Yellow Paper - https://github.com/ethereum/yellowpaper
* Tendermint Core - https://github.com/tendermint/tendermint
* Ethereum - Design Rationale -
  https://github.com/ethereum/wiki/wiki/Design-Rationale
* Bloom example -
  https://www.digitalcurrencygrid.com/2018/01/21/my-favorite-line-in-ethereums-java-implementation/
