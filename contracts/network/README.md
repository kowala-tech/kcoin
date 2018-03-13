# mUSD

## mToken as mUSD

The contract mUSD simply sets the descriptive fields for the token.

## Mintable token

`mToken` is implemented as a mintable token, that is, there is a maximum amount of tokens that can be minted. `totalSupply` holds the amount of minted tokens, `maximumSupply` returns the maximum amount of tokens, `mintTokens` allows the contract owner to mint tokens (dependent on `Ownable`) and the event `Mint` is triggered when new tokens are minted.

## Address triplets

A token holder (that has the intention of mining) needs a triplet of addresses:
* management address - used for token management
* mining address - used for mining
* receiver address - used to receive funds

Mining and receiver addresses can be set with a propose/accept mechanism.

To set a new mining address, the management address needs to call the `proposeMiningAddress(address miningAddr)` method. This address change needs to be accepted by the mining address (by calling the `acceptMiningAddress(address ownerAddr)` method).

To set a new receiver address, the process is similar. The management address proposes a new receiver (by calling the `proposeReceiverAddress(address receiverAddr)` method) and the mining address accepts the change (by calling the `acceptReceiverAddress(address receiverAddr)` method).

The contract owner can also use `initializeAccount(address ownerAddr, address miningAddr, address receiverAddr)` to initialize an account.

## Delegation

mTokens are also delegable in a very simplistic sense. A can delegate the control of X mTokens to B (and X "becomes" part of the balance of B). A should also be able to revoke this delegation at any time.

Example: A holds 100, B: 10, C: 0.

* A delegates 10 tokens to B
* A delegates 10 tokens to C
* B delegates 5 tokens to C

balances:

* `availableTo(A)` returns 80 (100 - 10 - 10)
* `availableTo(B)` returns 15 (10 + 10 - 5)
* `availableTo(C)` returns 15 (10 + 15)

# Contracts

Contains the address of special purpose contracts (and allow for those to be changed by the contract owner).

# Network

Contains network information and is used to save the total amount of minted wei.
It also contains the logic to handle network validators.

# kUSD-USD

The kUSD-USD simple oracle is derived from `PriceOracle`. It sets the name, symbol and decimal places for kUSD and USD. This basic implementation hold the amount of cryptocurrency and fiat in order to retain better precision This should be implemented according to the whitepaper description (of the price determining transactions).

# Data layouts

The types matching the contracts layout (and some helper methods) can be found [here](./data_layout.go)
