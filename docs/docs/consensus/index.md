# Proof-of-Stake Overview

A consensus protocol ensures that every transaction is replicated and recorded
in all the machines in the network in the same order. This is critical for an
autonomous entity like a blockchain, because even slight differences in
information on a single node could have significant consequences in terms of
network desynchronization.

There are several approaches to achieving consensus or "agreement", the most
well-known of which is probably _proof-of-work (PoW)_, which is used by Bitcoin
and many other cryptocurrencies. In a proof-of-work consensus, one node on the
network solves a difficult cryptographical problem and wins the right to add the
text block to the chain, which will typically include a financial bonus for the
winning node, called a _block reward_.

The Kowala Protocol utilises an approach called _proof-of-stake (PoS)_, which
is designed to be lighter and faster than proof-of-work. In a proof-of-stake
consensus, a number of known, stake-holding nodes take turns to propose the
next block in the chain. Depending on the implementation, the block reward can
be allocated entirely to the proposing node or shared between one of more of
the stakeholing nodes.

In this section of the documentation we will look at a special category of
consensus protocols, proof-of-stake protocols, and describe Kowala's approach to
the problem. We'll also go into detail on the main properties of the project
elements related to the consensus protocol. This module is heavily based on the
work of [Ethan Buchman](https://atrium.lib.uoguelph.ca/xmlui/handle/10214/9769)
as well as on [other resources](#https://blog.cosmos.network/tendermint/home)
provided by the Tendermint/Cosmos team and resources provided by the [Ethereum
project](https://www.ethereum.org/).

## Understanding Proof-of-Stake

This section is broken down into several parts:

1.  In [Kowala Protocol](/consensus/kowala) we introduce the project's protocol
    and include a breakdown of how the Kowala Protocol consensus differs from
    other implementations like Tendermint.
2.  In [Proof-of-Stake Pitfalls](/consensus/pitfalls), we're going to go
    through the major obstacles of implementing a PoS system.
3.  The [Network Safety](/consensus/safety) section outlines the faults that
    could be potentially expected and how we cope with them.
4.  Finally, the [Algorithm](/consensus/algorithm) section details the actual
    process of achieving consensus, and includes a breakdown of how the Kowala
    Protocol consensus differs from other implementations like Tendermint.

It might also be worth referring to the [glossary](/glossary).

</br></br>
