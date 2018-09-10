# Transaction Pool

Kowala's transaction pool contains all currently known transactions.

Since Kowala has a fixed compute unit price, there aren't stale transactions as
in Ethereum due to network congestion, thus the concept of replacing stale
transactions does not exist - there's no urgency factor. Additionally,
transactions are not sorted by price - all transactions are treated in the same
way; miners process all the transactions which decreases the propagation time of
a transaction which leads to faster transaction confirmations.

The transaction pool ensures that its content is valid with regard to the chain
state upon a new block. The process consists in removing any transactions that
have been included in the block or that have been invalidated because of another
transaction, update all the accounts to the latest pending nonce and moving
future transactions over to the executable list.

## Transaction life cycle

1 - Transactions enter the pool when they are received from the network or
submitted locally.

Note: by default the distiction between local and remote transactions is
enabled, but it can be disabled via config.

2 - Insert the transaction into the pending queue:

- If the transaction is already known, discard it.
- If the transaction fails basic validation, discard it:
  - heuristic limit - size.
  - transactions cannot be negative.
  - ensure that transaction does not exceed the block compute capacity.
  - make sure that the transaction is signed properly.
  - ensure that the transaction has a valid nonce.
  - sender should have enough balance to cover the (maximum) costs.
  - compute limit must cover the transaction's intrinsic computational effort.
- If the new transaction overlaps a known transaction, discard it.
- If the transaction is local, add it to the local pool.
- Add the transaction to the local disk journal.

3 - Transaction is promoted to the set of pending transactions - contains the
expected account nonce.

Kowala transactions contain a field called account nonce, which is a transaction
counter that should be attached for each transaction. The account nonce prevents
replay attacks and each node processes transactions from a specific account in a
strict order according to the value of its nonce. Let's say that the latest
transaction nonce was 4: if we send a new transaction from the same account with
the same nonce or below, the transaction pool will reject it; If we send a new
transaction with a nonce of 6, there's a gap and thus the new transaction will
not be processed until this gap is closed. If the gap stays open for too long,
the transaction is removed from the pool.

If there's a sequential increasing list of transactions starting at the expected
nonce, all of them are promoted.

4 - Transaction is executed and included in a block by the proposer if its
execution does not trigger errors.

5 - Block is committed by the miners - transaction confirmation.

6 - Transaction pool removes the transaction.
