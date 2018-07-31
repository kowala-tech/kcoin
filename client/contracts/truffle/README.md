# KNS Explained
KNS is the Kowala Name Service, a distributed, open, and extensible naming system based on the Kowala blockchain.

KNS can be used to resolve Kowala addresses, from human-readable form like “kowala.tech” into machine-readable identifiers, including Kowala addresses.

KNS is very similar to DNS, the Internet’s Domain Name Service, but has significantly different architecture, due to the capabilities and constraints provided by the Kowala blockchain. Like DNS, KNS operates on a system of dot-separated hierarchial names called domains, with the owner of a domain having full control over the distribution of subdomains.

First component of the KNS is **registry**. The registry is a central directory of KNS where all of the KNS domains are kept.
Every name in KNS can be found by looking it up in the KNS registry, and it’s the only component you need the address for.

**Registrars** are the second component of KNS, and are responsible for allocating new names to users. Registrars don’t have any special permissions — they just use their ability to tell the registry to create subdomains.

**Resolvers** are contracts that can tell you the resource associated with a name — such as an Kowala address.


## Namehash
KNS does not operate directly on names — there are a number of reasons for this, one of which is that parsing and processing text on the blockchain is very inefficient. Instead, KNS operates on secure hashes.

KNS hashes names using a system called namehash. Here’s the full definition of namehash:
* namehash('') -> '0x0000000000000000000000000000000000000000000000000000000000000000'
* namehash('a.xyz') -> sha3(namehash('xyz'), sha3('a'))


## How to use KNS?
To use KNS you will need to deploy all the components to the network.

Deploying KNS with *kowala.tech* domain and assign contract under this domain will look something like this (truffle test pseudo-code):
1. Deploy KNS Registry
	`KNS.new();`
2. Deploy FIFSRegistrar with parameters 
	1. **nsAddr** The address of the KNS registry.
	2. **node** The node that this registrar administers.

In our example it will be like this
`FIFSRegistrar.new(this.kns.address, namehash('tech'));`

3. Deploy Resolver with one parameter
	1. **nsAddr** The address of the KNS registry.

`PublicResolver.new(this.kns.address);`

4. Add root domain to the kns registry
	`kns.setSubnodeOwner(0, web3.sha3('tech'), this.registrar.address);`
    * First parameter: **node** The parent node.
    * Second parameter: **label** The hash of the label specifying the subnode.
    * Third parameter: **owner** The address of the new owner.

5. Register *kowala* domain under *.tech* root domain.
	`registrar.register(web3.sha3('kowala'), accounts[0]);`
    * First parameter: **subnode** The hash of the label to register.
    * Second parameter: **owner** The address of the new owner.

6. When we have a new domain under root domain, we should add resolver to that domain.
	`kns.setResolver(namehash('kowala.tech'), resolver.address);`
	* First parameter: **node** The node to update.
	* Second parameter: **resolver** The address of the resolver.
7. Now we can use our resolver to set our domain to point to address of a contract.
	`resolver.setAddr(namehash('kowala.tech'), kowalaContract.address);`
	* First parameter: **node** The node to update.
	* Second parameter: **addr** The address to set.

8. Having set everything up, we can now use simple function from our resolver to translate domain name to an address
	`resolver.addr(namehash('kowala.tech'));`


# Proxy Explained
All smart contracts that are deployed to the Blockchain network are immutable. Generally this is a good thing but it tends to be a double-edge sword. Code needs to be well tested, secure and take into account all corner cases. But, unfortunately this is not always a case and sometimes some bugs are discovered, not necessary in the smart contracts we developed, but in the modules, libraries it uses. Due to the immutability, we cannot update the code of smart contract by itself, we need another solution. To a rescue comes Proxy Pattern.

A proxy architecture pattern is such that all message calls go through a Proxy contract that will redirect them to the latest deployed contract logic. To upgrade, a new version of your contract is deployed, and the Proxy is updated to reference the new contract address.

**User** -> **Proxy Contract**(*storage layer*) -> **Logic Contracts**(*Logic layer*).

Proxy pattern is needed to handle our KNS contracts. KNS solves similar problem and it is use to handle our other contracts, but we don’t have a way to update our KNS contracts if some vulnerability is discovered. Hence use of Proxy contracts for our KNS system.

To use tell our system to use proxy contracts of KNS instead of KNS contracts we need to do the following.

1. Deploy new ProxyFactory contract
2. Deploy a new KNS contract
3. Create new proxy from ProxyFactory for the KNS contract
4. Point KNS contract to a proxyKNS contract
5. Use KNS which in fact uses address of our proxy which addresses all calls to a KNS logic

This steps in a js code would look like this
```
const proxyFactory = await UpgradeabilityProxy.new();  
const kns = await KNS.new();
const logs = await proxyFactory.createProxy(admin, kns.address, { from: admin });
const logs1 = logs.logs;
const knsProxyAddress = logs1.find(l => l.event === 'ProxyCreated').args.proxy;
const knsProxy = await AdminUpgradeabilityProxy.at(knsProxyAddress);
let knsContract = new KNS(knsProxyAddress);
await knsContract.initialize(owner);
```

Notice the initialize function. Most contracts require some sort of initialization function, but upgradeable contracts can't use constructors because the proxy won't be able to call them. This is why we need to use the initializable pattern provided by zos-lib.

## Update contract via proxy
These are the steps to update our contract.

1. Deploy new version of KNS contract
2. Call on proxy `upgrateTo();` function and provide address of a new deployed contract.
3. Point KNS contract to a new KNSV1 at proxy address.

This steps in a js code would look like this

```
const knsv1 = await KNSV1.new();
await knsProxy.upgradeTo(knsv1.address, { from: admin });
knsContract = await KNSV1.at(knsProxyAddress);
```
