# Receipt

The purpose of the receipt is to provide information about the transaction
execution - contains information that is only available once a transaction has
been executed.

## Receipt Specification

| Field                   | Description                                                                                              |
| ----------------------- | -------------------------------------------------------------------------------------------------------- |
| PostState               | Post transaction state                                                                                   |
| Status                  | Transaction execution status (failed/succeeded)                                                          |
| CumulativeResourceUsage | Sum of computational resources used by this transaction and all preceding transactions in the same block |
| Bloom                   |                                                                                                          |
| Logs                    |                                                                                                          |
| TxHash                  | Transaction Hash                                                                                         |
| ContractAddress         | Address assigned to the new contract (only for contract creation transactions)                           |
| ResourceUsage           | Computational resources usage (in compute units)                                                         |
