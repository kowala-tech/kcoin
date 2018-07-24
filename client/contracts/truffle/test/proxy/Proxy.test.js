/* global artifacts, contract, it, beforeEach, web3 */
/* eslint no-unused-expressions: 0 */
/* eslint consistent-return: 0 */
/* eslint-disable max-len */

require('chai')
  .use(require('chai-as-promised'))
  .should();

const KNS = artifacts.require('KNSRegistry.sol');
const FIFSRegistrar = artifacts.require('FIFSRegistrar.sol');
const PublicResolver = artifacts.require('PublicResolver.sol');
const ValidatorMgr = artifacts.require('ValidatorMgr.sol');
const UpgradeabilityProxy = artifacts.require('UpgradeabilityProxyFactory.sol');
const AdminUpgradeabilityProxy = artifacts.require('AdminUpgradeabilityProxy.sol');
const namehash = require('eth-ens-namehash');

contract('Proxy Functionality', (accounts) => {
  beforeEach(async () => {
    this.kns = await KNS.new();
    this.registrar = await FIFSRegistrar.new(this.kns.address, namehash('kowala'));
    this.resolver = await PublicResolver.new(this.kns.address);
    await this.kns.setSubnodeOwner(0, web3.sha3('kowala'), this.registrar.address, { from: accounts[0] });
  });
  it('should access contracts via proxy', async () => {
    const proxyFactory = await UpgradeabilityProxy.new();

    // KNS Proxy
    const logs = await proxyFactory.createProxy(accounts[0], this.kns.address);
    const logs1 = logs.logs;
    let proxyAddress = logs1.find(l => l.event === 'ProxyCreated').args.proxy;
    const knsProxy = await AdminUpgradeabilityProxy.at(proxyAddress);
    const knsImplAddr = await knsProxy.implementation();
    const knsContract = await KNS.at(knsImplAddr);
    await knsContract.initialize(accounts[0]);

    await knsContract.owner(0).should.eventually.be.equal(accounts[0]);

    // Registrar Proxy
    const logs2 = await proxyFactory.createProxy(accounts[0], this.registrar.address);
    const logs3 = logs2.logs;
    proxyAddress = logs3.find(l => l.event === 'ProxyCreated').args.proxy;
    const registrarProxy = await AdminUpgradeabilityProxy.at(proxyAddress);
    const registrarImplAddr = await registrarProxy.implementation();
    const registrarContract = await FIFSRegistrar.at(registrarImplAddr);
    await registrarContract.initialize(knsImplAddr, namehash('kowala'));

    await registrarImplAddr.should.be.equal(await knsContract.owner(namehash('kowala')));

    // Resolver Proxy
    const logs4 = await proxyFactory.createProxy(accounts[0], this.resolver.address);
    const logs5 = logs4.logs;
    proxyAddress = logs5.find(l => l.event === 'ProxyCreated').args.proxy;
    const resolverProxy = await AdminUpgradeabilityProxy.at(proxyAddress);
    const resolverImplAddr = await resolverProxy.implementation();
    const resolverContract = await PublicResolver.at(resolverImplAddr);
    await resolverContract.initialize(knsImplAddr);

    await resolverImplAddr.should.be.equal(await this.resolver.address);
  });
});
