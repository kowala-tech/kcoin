/*
Package contracts implements the main network contracts.

Contracts play a central role in Kowala - the state of all accounts is the state
of the Kowala network which is updated with every block.
The Kowala dev team is currently focused and actively improving contracts
concerned with the consensus validation, price stabilisation and rewards.

Smart Contracts

The best way to describe smart contracts is to compare the technology to a
vending machine.
Ordinarily, you would go to a lawyer or a notary, pay them, and wait while you
get the document.
With smart contracts, you simply drop a bitcoin into the vending machine (i.e.
ledger), and your escrow, driver’s license, or whatever drops into your
account. More so, smart contracts not only define the rules and penalties around
an agreement in the same way that a traditional contract does, but also
automatically enforce those obligations.

Smart contracts can:

* Function as 'multi-signature' accounts, so that funds are spent only when a
required percentage of people agree.
* Manage agreements between users, say, if one buys insurance from the other.
* Store information about an application, such as domain registration
information or membership records (Ex: consensus validators).

Kowala Network contracts:

| File             | Description                                    |
|------------------|------------------------------------------------|
| network.sol      | consensus validators / slashing rules contract |
| mtoken.sol       | mtoken contract                                |
| kUSD.sol         | kUSD contract                                  |
| price_oracle.sol | oracle prices                                  |

References
* https://blockgeeks.com/guides/smart-contracts/

*/

package contracts
