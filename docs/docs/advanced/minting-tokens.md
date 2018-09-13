# Minting mining tokens

## Overview

In the early stages of any real-currency network, or on a test net, it's
necessary to generate and allocate mining tokens in order to validate on the
network. 

The sum of tokens for each network is 2^30, which is roughly one billion.  Only
a few of these tokens are pre-minted (they're used by the genesis validator to
start the network); the rest must be minted and allocated to token holders by a
set of temporary governor accounts which are initially defined in the genesis
block, and then controlled by vote of the governors themselves. Eventually all
the tokens will be minted, and the governors can retire from their roles.

Tokens are held by any valid address in the network. They can be held by humans
via any kind of software, hardware or brain wallet key, and the can also be
held by smart contracts.

## Governance keys

Each governor is represented by an account, or _governance key_, that is able
to make transactions on the network. In the case of main nets, these are
typically hardware keys that kept under very tight security.

The governors are represented on the network as a multisig contract. There can
be any number of governance keys, and the governors may choose between them the
number of agreeing votes (called _confirmations_) required for any given
transaction (a process sometimes called M of N). The minting of tokens is such
a transaction, and it therefore requires a previously specified number of
confirmations in order to take place.

For example, a group of 5 governors might decide that they require 3 votes
between them in order to issue some number of tokens to an account. In order to
actually mint tokens, one governor would propose a transaction that mints X
transactions to address Y.  Once two other governors confirm this initial
proposition, the tokens would be automatically minted. If there are not enough
votes, the tokens will never be minted.

## Minting procedure

The process consists of a series of data transactions made to a multisig contract
that's embedded into the network genesis. Creating, signing and sending complex
data transactions is rather difficult for most humans, so we're made some
tooling in the console that makes things easier.

The basic process is as described above: one governance key submits a
transaction that will mint some tokens to an account. This initial submissions
results in a _minting ID_: a number, which is an arbitrary reference to the
suggested minting operation.  The remaining governance keys can confirm the
transaction by referencing the minting ID &mdash; or ignore it if they don't want to
confirm it.

The operations that constitute this process take place in the Javascript console.
In order to use them, you'll need to have a running node and access to your secure
keys (which may be on USB devices).

For the sake of simplicity, the examples given below assume 2 of 3 multisig setup.
That is, one governor submits a transaction, and only one other governor has to
confirm it.

### Transaction submission

#### Syntax

 `mtoken.mint(governance_key_address, recipient_address, amount_of_tokens)`

#### Example usage in the console 

```
> mtoken.mint(eth.coinbase, "0x1f1f1480f77b2565ae7f3a5580fd3da79b59b09b", 10); 
```

This would submit a transaction of 10 mtokens to the address ending `b09b`. The
transaction will not be actioned until it's confirmed by another governance
key.

### Listing pending and complete transactions

#### Syntax

`mtoken.mintList()`

#### Example usage in the console

```
> mtoken.mintList()
```

For which the output might be something like:

```
[{
    amount: 42,
    confirmed: false,
    id: 0,
    to: "0x1f1f1480f77b2565ae7f3a5580fd3da79b59b09b"
}, {
    amount: 44,
    confirmed: true,
    id: 1,
    to: "0x1f1f1480f77b2565ae7f3a5580fd3da79b59b09b"
}]
```

The `id` field is the minting ID, which will be needed in the confirmation
step. Once a transaction is submitted, it will appear at the end of this list
for verification. The `confirmed` field indicated whether or not another
governance key has confirmed the transaction (and thus actioned it).

### Confirming

#### Syntax

 `mtoken.confirm(governance_key_address, minting_id)`

#### Example usage

```
> mtoken.confirm(eth.coinbase, 1)
```

This would confirm the transaction with minting ID `1`, and action it. In the
example output above, it would issue 42 tokens to the address ending `b09b`.

