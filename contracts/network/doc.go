/*
Package network implememnts the main network smart contracts

Contracts

Network - network.sol

Terms

Collateral - staked tokens.
Validator - consensus participant.

Use Cases

* The contract owner should be able to specify the genesis validators and their
initial collateral as arguments for the contract creation.

* In order to achieve short block times, there needs to be a limit to the number
validators. More validators means more time required for block times in the
tendermint consensus protocol. We've included a variable called maxVoters that
can only be modified by the contract owner. It might make sense to decrease the
number of validators at a certain point in time due to external circumstances.
In this case, the validators with the smallest collateral will be removed from
the election if all the validator positions are occupied. We've also included
hard limits - that can be managed - for safety.

* Everyone in the network should be able to be a validator as long as:
1 - There are positions available and the deposit is bigger than the minimum
required deposit.
2 - There are no positions available but the deposit is bigger than the current
smallest deposit. In this case, the new candidate replaces the validator with
the smallest collateral in the election.

* Validators should be able to increase their stake at any time. In order to
do so, they must make a deposit, and this deposit gets added to their current
collateral.

* Validators should be able to join the elections at any time, even if part of
their tokens are under the unbond period, meaning that one can have more than one
collateral at a time.

* Validators should be able to leave an election at any time - they still need
to wait for the confirmation! Afterwards the coins remain locked for a
predetermined period of time called the unbonding period (4 weeks for now),
after which the user - not a validator at this point, except if rejoined the
elections - is free to transfer or spend those coins. The funds are not
automatically transfered to the user account after the unbonding period - The
user needs to send a "withdraw" transaction in order to have access to them
again. This situation happens because future execution of events in the
contracts does not offer guarantees of the transaction execution and for that
reason the validator needs to submit a request to get the funds back to his
account - lazy evaluation.

* The user can withdraw multiple collaterals (staked tokens) with just one "withdraw"
transaction. Scenario:
1. Validator leaves the consensus elections > collateral with a release date.
2. Validator joins the election (before the first collateral is released) > new
collateral (with no release date) + previous collateral.
3. Validator leaves the consensus elections > two collaterals with a release
date.
4. Validator joins the election (before the first collateral is released) > new
collateral (with no release date) + previous collaterals (2).
5. The unbonding period is over for the first two collaterals and the user can
withdraw both.
6. By sending an withdraw transation, his account gets refunded with the
collaterals which unbonding period has expired.
Note that the solidity transfer() function doesn't result in a transaction. It
results in a message call inside the original transaction initiated by an
external account. The blockchain will record a single transaction no matter how
many transfer() or call() invocations there are in the code. The gas cost will
be deducted from the external account that initiated the transaction.


Minimum Deposit

Unbonding Period

Transactions

* Deposit - the deposit transaction is used to register a new candidate or to
increase the stake of a current voter.

* Leave - the validator wants to leave the election - the current collateral
gets a release date.

* Withdraw - After the unbond period the validator needs to request the funds
in order to have them back in his account.

References

* https://blog.ethereum.org/2014/07/05/stake/
- Vitalik Buterin
* https://ethereum.stackexchange.com/questions/38387/contract-address-transfer-method-gas-cost
- medvedev1088


*/

package network
