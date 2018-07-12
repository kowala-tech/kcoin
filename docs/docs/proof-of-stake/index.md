# Proof-of-stake (PoS)

A consensus protocol ensures that every transaction is replicated and recorded in all the machines in the network in the same order.

In this section of the documentation we will look at a special category of consensus protocols, proof-of-stake protocols, and describe Kowala's approach to the problem and the main properties of the project elements related to the consensus protocol. This section is heavily based on the work of [Ethan Buchman](https://atrium.lib.uoguelph.ca/xmlui/handle/10214/9769) as well as on [other resources](#https://blog.cosmos.network/tendermint/home) provided by the Tendermint/Cosmos team and resources provided by the [Ethereum project](https://www.ethereum.org/).

The Kowala project has its own implementation of the [Tendermint protocol](https://github.com/tendermint/tendermint). Tendermint is a weakly synchronous, Byzantine fault tolerant, state machine replication protocol, with optimal Byzantine fault tolerance and additional accountability guarantees in the event the BFT assumptions are violated. There are varying ways to implement Proof-of-Stake algorithms, but the two major tenets in Proof-of-Stake design are chain-based PoS and Byzantine Fault Tolerant-based PoS. Kowala implements a hybrid of both - strictly choosing consistency over availability. Some of the main properties of the Kowala project are:

* The codebase is based on [Ethereum's go client](https://github.com/ethereum/go-ethereum/) - As of Jan 2018 Ethereum tokens account for 90% Market share. We believe that most of these projects will benefit from a stable coin context. It's not uncommon for volatile cryptocurrencies to fall or rise more than 50% in a day. While traders and crypto supporters are used to those fluctuations, they could be catastrophic for the average person - No one wants to wake up and find half of their savings missing. No one can live with a currency that changes value while they’re in the grocery store.
* Fast-finality (1 second confirmations) - We believe that fast confirmations will be essential for mass adoption.
* On-chain dynamic validator set management (registry/in-protocol penalties) via genesis smart contracts.

The order at which information is presented is intentional:

* [Glossary](http://docs.kowala.tech/proof-of-stake/glossary/) of terms related to consensus protocols.
* Major pitfalls of native proof-of-stake implementations.
* Protocol accountability - identifying all byzantine validators when there is a violation of safety.
* Overall picture of the consensus algorithm.
