# Transaction

Transaction refers to the signed data package that stores a message to be sent
from an externally owned account to another account on the blockchain. When you
interact with the Kowala blockchain, you are executing transactions and updating
the blockchain state. A given transaction is only valid once. This is done using
sequence numbers (Nonce).

## Transaction Specification

| Field        | Description                                                                            |
| ------------ | -------------------------------------------------------------------------------------- |
| AccountNonce | Represents the number of transactions sent from a given address                        |
| ComputeLimit | The compute units that this transaction can use (not required for simple transactions) |
| Recipient    | The address to which we are directing this message (nil for contract creation)         |
| Amount       | Total kUSD that you want to send                                                       |
| Payload      | Data field in a transaction                                                            |

### Context

- AccountNonce, is used to enforce some rules such as:

  - Transactions must be in order. This field prevents double-spends as the
    nonce is the order transactions go in. The transaction with a nonce value of 3
    cannot be mined before a transaction with a nonce value of 2.

- Amount, If you are executing a transaction to send kUSD to another person or a
  contract, you set this value.

- Payload is either a byte string containing the associated data of the message,
  or in the case of a contract-creation transaction, the initialisation code.

</br></br>
