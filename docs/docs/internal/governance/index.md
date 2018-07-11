# Governance

Blockchain governance is one of the most important topics on the space. The
ability of a blockchain to succeed over time is based on its ability to evolve.
In order to evolve there are many decisions that must be taken and this is
where governance comes in.

## Early Stages

While Kowala is in the early alpha stage, there must be a set of governance
transactions that can be done by the Kowala team to ensure that the system
keeps running as expected at all times - in practical terms, the core network
contracts will be manageable. In case of a major bug or an important feature
change, the core team can deploy a smart contract upgrade through the name
service mechanism which will be explained below. The core team should not have
the power to re-establish state though - the updates should only be only done
on the functional level and not to state, but problems happen, and we need to
account for such scenario. The core contracts will be managed by a multi
signature wallet owned by the core team - during the genesis block creation, we
will specify the list of addresses that will own the wallet - this is the
recommended way to own important assets, as long as we use a proven
implementation of a multi signature wallet. We've decided to use the [gnosis's
multi signature wallet](https://github.com/gnosis/MultiSigWallet) as it's being
used by multiple companies.

The failsafes mentioned above rely on the following parts:

* A multi signature wallet owned by kowala that has the ability to request
  certain operations such as pausing a service (Ex: stopping the oracle
  service). The keys must be stored offline (ex: hardware wallet - our client
  supports some; we can use them to sign txs) and we should define an internal
  process to handle all the situations.
* A domain name system - contract updgrades by pointing a domain to a new
  contract(address); also solves the problem of having to remember an address
  vs a domain name (like IP vs domain name in browsers). The upgrade model is
  similar to this
  [one](https://medium.com/cardstack/upgradable-contracts-in-solidity-d5af87f0f913)
  except for the name service contract which will be updated via the
  proxy-delegatecall mechanism since the contract is the name service itself.

## Proven Concept (removing failsafes > use forks)

At some point in time, kowala's governace should be terminated on the contract
level. By then, Kowala has to implement a process similar to Ethereum in order
to improve the system. In their current model, an improvement/software is
proposed, discussed, accepted, coded into clients and finally released. The
upgrade is then subject to political analysis.

<figure> <img src="/assets/images/evolution.jpg" > </figure>

As soon as clients are released, the users are informed about what's contained
in the new release and they have a choice to run the hard fork client or not.
In case of disagreement they continue with the legacy chain. In the current
model, if the community rejects the change, it can always refuse to adopt it.
The development is centralized with Kowala but ultimately, the user is the one
deciding which chain to use.

## Edge cases

Multiple faulty exchanges/tremendous volatility - users could report manually
and be rewarded via oracle as suggested by our internal reviewers.

