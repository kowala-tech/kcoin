# Overview

** Note: Kowala's oracle client is currently under development. **

The most compelling applications of smart contracts – such as financial
instruments – require access to data about real-world state and events and smart
contracts lack network access. Oracles aim to meet this need. The end goal of
Kowala's authenticated data feed system is to incentivize any network user to
provide the exchange prices in a decentralized and timely manner in order to
keep the coin stable in return of rewards.

In this section we detail Kowala's authenticated data feed system. Kowala's
system has taken inspiration from projects such as [Town
Crier](http://www.town-crier.org/) for Intel SGX's use cases aswell as from
[IPFS Desktop](https://github.com/ipfs-shipyard/ipfs-desktop) to deliver a very
simple application for the users.

Kowala's Oracle client executes code in an Intel SGX enclave which is a
protected address space–trusted hardware. Any processes running in the enclave
are protected from hardware attacks and software running on the same host.
Additionally, any remote client can verify the software running in the enclave
by requesting a hash of the enclave state which is signed by the enclave’s
hardware protected private key. Anyone with the enclave’s public key can then
verify signed attestations made by the enclave about the program state. The
signed attestation proves to users that TownCrier could have not tampered the
retrieved data as long as users believe that Intel’s trusted hardware
implementation is trustworthy. By using Intel SGX we don’t need to trust the
operators. Even the operators of the oracle cannot tamper with its operation or,
for that matter, see the data it's processing. Note that we still need to trust
Intel’s implementation but it’s a major step over Oraclize solution and
something necessary/critical to make sure that we not subject to
hardware/software attacks.

</br></br>
