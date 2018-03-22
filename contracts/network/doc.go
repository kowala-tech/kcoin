/*
Package network implements the network related contracts

This package includes the consensus election contract.

Right now, these contracts are included in the genesis block via the genesis
file creation process (using puppeth) or by using the default genesis for the
kowala networks - pre allocated data is present in here /core/genesis_alloc.go.
The process consists on running the contract constructor against an empty state
using the runtime.Create method and include the resulting code and storage in a
new genesis account. The initial deposit for the genesis validator is also added
in the contract genesis account as balance.


Consensus Election Contract

The election contract manages the election validators. During its creation
you are asked to provide the address of the genesis validator, a base deposit,
the maximum number of validators and an unbonding period:

Maximum Number of validators - in order to guarantee low block times, there
needs to be a limit to the number validators. More validators mean more time
required to achieve consensus in the Tendermint protocol.  Imagine a real
election with 10 participants and another one with 1 million participants. It
takes much more time to get to a conclusion on the last election. This number
can be modified to different values as it might make sense to decrease/increase
the number at a certain point in time due to external circumstances. In case
that the new number is less than the current number of validators, the
validators with the smallest deposit will be the ones leaving the election
first. Note that the value must be greater or equal than one in order to
accomodate the genesis validator.

Base deposit - represents the deposit that a candidate has to do in order to
secure a place in the elections (if there are positions available). The base
deposit solves the nothing at stake problem. This value can be modified by the
contract owner.

Unbonding period - is a predetermined period of time that coins remain locked,
starting from the moment a validator decides to leave the consensus elections.
The unbonding period prevents long-range double-spend attacks. A user can avoid
long-range attacks by syncing their blockchain periodically within the boulds of
the unbounding period. This value should be provided in days. This value can be
modified by the contract owner. Tendermint is aiming 4 weeks.

We also must provide the address of the genesis validator in order to add that
identity to the validator list. Initially we thought about adding the genesis
validator deposit by calling the constructor method with a value (payable) but
it does not make much sense for the genesis block as we start with an empty state
and the account is included from the start. Instead the contract balance must
cover the genesis deposit. We will probably need to revisit this topic as soon
as the token contracts are ready.

The consensus election contract contains three main methods:

Join - any user in the network should be able to join the election as long as it
has the funds to cover the minimum deposit. The minimum deposit is calculated in
the following way:
1. If There are positions available, the minimum deposit is equal to the base
deposit.
2. If not, the minimum deposit is equal to the smallest bid in the election + 1.
The only exception to the rule is that the same identity cannot be used for
different validators.

Leave - allows a validator to leave the election. This call will trigger
operations such as setting a release date for the current deposit (current time
+ unbonding time). Note that an ex-validator can join again the election as soon
as we wants as long as he has the funds to do it. These means that even if there
are locked funds, as soon as he joins he will have a new deposit resulting on
multiple deposits.

RedeemDeposits - requests a transfer of the unlocked funds - current date is
past the release date -  back to the validator account. The network user can
redeem multiple deposits with just one request as long as he has those deposits
unlocked. In order to confirm this information, an user can use the console to
verify the current deposits of the validator.

*/

package network
