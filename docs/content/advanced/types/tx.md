# Transaction

Transaction refers to the signed data package that stores a message to be sent from an externally owned account to another account on the blockchain. When you interact with the Kowala blockchain, you are executing transactions and updating the blockchain state. A given transaction is only valid once. This is done using sequence numbers (Nonce).

## Transaction Specification

| Field        | Description                                                     |
| ------------ | --------------------------------------------------------------- |
| AccountNonce | Represents the number of transactions sent from a given address |
| Recipient    | The address to which we are directing this message              |
| Amount       | Total kUSD that you want to send                                |
| Payload      | Data field in a transaction                                     |
| GasLimit     | The maximum amount of gas that this transaction can consume     |
| Price        | Price per unit of gas                                           |

### Context

* Nonce, is used to enforce some rules such as:

  * Transactions must be in order. This field prevents double-spends as the nonce is the order transactions go in. The transaction with a nonce value of 3 cannot be mined before a transaction with a nonce value of 2.

* Amount, If you are executing a transaction to send kUSD to another person or a contract, you set this value.

* Payload is Either a byte string containing the associated data of the message, or in the case of a contract-creation transaction, the initialisation code.

* Price - If you want to spend less on a transaction, you can do so by lowering the amount you pay per unit of gas. The price you pay for each unit increases or decreases how quickly your transaction will be mined.
