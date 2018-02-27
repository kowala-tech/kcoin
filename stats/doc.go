/*
Package stats implements the network stats reporting service.

We're actively improving the stats. The current implementation contains elements
and terms based on the proof-of-work protocol - it's a goal of ours to modify
the stats to monitor metrics relevant in proof-of-stake.


Service

The service relies on a netstats url. If the url is not present in the
configuration, the service is not enabled. The url has the following format:

<hostname>:<security_token>@<domain>:<port>

Example:
"Ricardo's node:DVagynuHLdn9sK6c@testnet.kowala.io:80"

The stats reporting service depends on the kusd service and changes in the
reports content must be followed with an update in the netstats, otherwise, the
reporting will not work.


Netstats

Kowala has a slighty modified version of netstats -
https://github.com/kowala-tech/kUSD-netstats.


Events

The reporting service subscribes chain head events for block reports and
transaction events to report the current tx pool stats. Every 15 seconds,
there's also a status update that includes all possible data to report -
includes latencies, block, tx pool, and local node stats. There's also a history
report that retrieves the most recent batch of blocks and reports it to the
stats server.


Reports

Block report

| Field       | Description                                   |
|-------------|-----------------------------------------------|
| Number      | block number                                  |
| Hash        | block hash                                    |
| Parent Hash | parent block hash                             |
| Timestamp   | block creation timestamp                      |
| Miner       | block proposer                                |
| GasUsed     | gas used                                      |
| GasLimit    | block gas limit                               |
| Txs         | transaction count                             |
| TxHash      | hash of the root node of the transaction trie |
| Root        | hash of the root node of the state trie       |

TX Pool report

| Field   | Description                    |
|---------|--------------------------------|
| Pending | number of pending transactions |

Local Node Status report

| Field     | Description                         |
|-----------|-------------------------------------|
| Active    | is always true                      |
| Mining    | true if the validator is validating |
| Peer      | current peer count                  |
| Gas Price | suggested gas price                 |
| Syncing   | true if the sync is enabled         |
| Uptime    | is always 100                       |

Latency report

| Field   | Description                                      |
|---------|--------------------------------------------------|
| latency | RTT time - based on a ping request to the server |
*/

package stats
