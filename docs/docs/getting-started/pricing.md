# Kowala Pricing

This section covers the network pricing.

- With Kowala you pay per transaction. The transaction fee is debited from the
  source account - be advised that the amount debited from the source account will
  be slightly larger than that credited to the target account.

- Kowala does not support transaction prioritisation - first come, first served
  as the major payment networks; transaction prioritisation was born out of the need
  to prioritise transactions due to network congestion (Ex: Ethereum network).

- Different operations require different computational efforts. The price of the
  compute unit is currently set to <b>k$0.0000004</b>.

With that said, that are two types of transactions:

## kUSD Transactions

The simplest transactions are kUSD transfer transactions.

These transactions have a fixed base cost, <b>k$0.0084.</b>

## Contract Transactions

Contract transactions have variable costs because the blockchain state is not
static and different operations require different computational efforts. For that
reason, users must provide a compute limit - Kowala provide tools to calculate an
estimated limit - which is the maximum amount of computational resource the sender is
willing to pay for this transaction.

Additionally, the following points must be considered:

- The miners will stop the transaction execution as soon as they run out of
  computational resource (limit).

- If there's any computational resource left over (compute limit - resource usage),
  the correspondent cost will be immediatly refunded to the user's account.

</br></br>
