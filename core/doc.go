/*
Package core implements the most relevant data structures and core elements of
the blockchain

* Data Structures

Blockchain

BlockChain represents the canonical chain given a database with a genesis
block. The Blockchain manages chain imports, reverts, chain reorganisations.
Assuming that less that 1/3 of the validators are Byzantine, Tendermint
(consensus protocol) guarantees that safety will never be violated - validators
will never commit conflicting blocks at the same weight. Therefore, the Kowala
blockchain never forks.
Tendermint prioritizes safety and finality over liveness. For more on settlement
finality: https://blog.ethereum.org/2016/05/09/on-settlement-finality/.


Transaction Pool (tx_pool.go)

The transaction pool contains all of the transactions that have been submitted
to the Kowala network but have not been allocated to a block. The transaction
pool removes the transaction as soon as it receives a block with the
transaction, freeing up space.


Voting Table (voting_table.go)

The voting table helps collect signatures from validators of a consensus
election round for a predefined sub election (pre-vote or pre-commit).
Voting tables also let us know if there's a majority (+2/3) - impacts the
validator's decisions.

*/

package core
