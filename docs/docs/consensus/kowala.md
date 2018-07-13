# Kowala Protocol

The Kowala Project has its own implementation of the [Tendermint
protocol](https://github.com/tendermint/tendermint). Tendermint is a weakly
synchronous, Byzantine fault tolerant (BFT) state machine replication protocol,
with optimal Byzantine fault tolerance and additional accountability guarantees
in the event that the BFT assumptions are violated. There are varying ways to
implement Proof-of-Stake algorithms, but the two major tenets in Proof-of-Stake
design are _chain-based PoS_ and _Byzantine fault tolerant-based PoS_. The
Kowala protocol implements a hybrid of both — we use the Tendermint protocol for
fast finality - strictly choosing consistency over availability - while
leveraging Ethereum's smart contract platform to manage the dynamic validator
registry and to govern part of the protocol - for instance, the application of
punishments to faulty validators.

Some of the main/unique properties of the Kowala project are:

- **Fast-finality (1 second confirmations!)**.

There are certain sacrifices that must be made to get a certain experience with
fault tolerant consensus protocols. The following picture summarises the
tradeoffs that must be made when choosing a fault tolerant protocol:

<figure> <img src="/assets/images/consensus_triangle.jpg" ><figcaption>Figure 1 -
Vlad Zamfir's triangle.</figcaption></figure>
</br>

Tendermint falls somewhere along the bottom-left corner while Kowala is focused
on the top corner - we believe that fast confirmations will be essential to
rival the current payment networks(Visa, Mastercard, Amex) and to achieve mass
adoption. Based on that, we choose to create finalized blocks and have fast
finality via tendermint protocol. This way, we can guarantee a smooth experience
while overcoming the major points of failure in today's financial systems:
breaches in security, network downtime, or network outages. Note that a "small
number of nodes" is relative. We can extend the number of nodes to a certain
extent by imposing high tier hardware. The speed of progress does not depend on
system parameters, but instead depends on real network speed.

<hr>

- **Our codebase is based on [Ethereum's go client](https://github.com/ethereum/go-ethereum/)**.

Some of the reasons why we've decided to focus on the ethereum client:

- As of Jan 2018 Ethereum tokens account for 90% Market share.

We believe that most of these projects will benefit from a stable coin context -
it's not uncommon for volatile cryptocurrencies to fall or rise more than 50% in
a day. While traders and crypto supporters are used to those fluctuations, they
could be catastrophic for the average person — no one wants to wake up and
find half of their savings missing. No one can live with a currency that changes
value while they’re in the grocery store. Kowala provides a similar application
programming interface that facilitates porting any project from Ethereum to
Kowala.

- Smart contracts may be revolutionary and they are fundamental for Kowala's use
  cases. The first set of standards are also being discussed by this community.
- Ethereum accounts for most of the scalability research underway (+100k tx/s).

<hr>

- **Chain-based Proof-of-Stake features**.

Kowala network is composed by a set of core contracts that automatically
coordinate specific operations, vital to the network. For instance, the
validator manager contract handles the validators' registration/deregistration
or punishments. We achieve this by embeddeding the managers' contracts in the
genesis block. By following this approach not only the implementation is simpler
but also it becomes easier to upgrade very sensitive parts of the system.

</br></br>
