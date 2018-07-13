# Oracle landscape

## [Oraclize](http://www.oraclize.it/Oraclize)

Oraclize stores the TLS notary secret in AWS virtual machine. By involving AWS, Oraclizehas made it harder for their data retrieval service to lie about the contents of a webpage. Major disadvantage is that Amazon, themselves, or anyone able to hack/subpoena Amazon’s AWS platform, can gain the ability to fake the proofs of honesty by stealing the AWS oracle’s private key. If you do catch someone faking Oraclize proofs, there’s no way to prove it to anyone else.

## [TownCrier](http://www.town-crier.org/)

The main difference to Oraclize is that TownCrier executes code in an Intel SGX enclave which is a protected address space–trusted hardware. Any processes running in the enclave are protected from hardware attacks and software running on the same host. Additionally, any remote client can verify the software running in the enclave by requesting a hash of the enclave state which is signed by the enclave’s hardware protected private key. Anyone with the enclave’s public key can then verify signed attestations made by the enclave about the program state. The signed attestation proves to users that TownCrier could have not tampered the retrieved data as long as users believe that Intel’s trusted hardware implementation is trustworthy. The main idea is that by using Intel SGX we don’t need to trust the operators. Even the operators of the oracle cannot tamper with its operation or, for that matter, see the data it's processing. Note that we still need to trust Intel’s implementation but it’s a major step over Oraclize solution and something necessary/critical to make sure that we not subject to hardware/software attacks.

**Note** This solution requires the users to have processors with Intel SGX support - latest Intel CPUS include it. Alternatively, there are cloud providers that are offering SGX support at the moment such as [IBM](https://www.ibm.com/blogs/bluemix/2017/12/data-use-protection-ibm-cloud-ibm-intel-fortanix-partner-keep-enterprises-secure-core/).
