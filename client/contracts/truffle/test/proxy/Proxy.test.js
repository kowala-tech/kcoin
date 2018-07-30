/* global artifacts, contract, it, beforeEach, describe, before, web3 */
/* eslint no-unused-expressions: 0 */
/* eslint consistent-return: 0 */
/* eslint-disable max-len */

process.env.NODE_ENV = 'test';

require('chai')
  .use(require('chai-as-promised'))
  .use(require('chai-bignumber')(web3.BigNumber))
  .should();

const { Contracts } = require('zos-lib');

const KNS = artifacts.require('KNSRegistry.sol');
const KNSV1 = artifacts.require('KNSRegistryV1.sol');
const FIFSRegistrar = artifacts.require('FIFSRegistrar.sol');
const PublicResolver = artifacts.require('PublicResolver.sol');
const ValidatorMgr = artifacts.require('ValidatorMgr.sol');
const UpgradeabilityProxy = Contracts.getFromNodeModules('zos-lib', 'UpgradeabilityProxyFactory');
const AdminUpgradeabilityProxy = Contracts.getFromNodeModules('zos-lib', 'AdminUpgradeabilityProxy');
const namehash = require('eth-ens-namehash');

contract('Proxy Functionality', ([_, admin, owner, anotherAccount]) => {
  it('should access contracts via proxy', async () => {
    const proxyFactory = await UpgradeabilityProxy.new();

    // KNS Proxy
    const kns = await KNS.new();
    const logs = await proxyFactory.createProxy(admin, kns.address, { from: admin });
    const logs1 = logs.logs;
    const knsProxyAddress = logs1.find(l => l.event === 'ProxyCreated').args.proxy;
    const knsProxy = await AdminUpgradeabilityProxy.at(knsProxyAddress);
    let knsContract = new KNS(knsProxyAddress);
    await knsContract.initialize(owner);

    // Registrar Proxy
    const registrar = await FIFSRegistrar.new(knsProxyAddress, namehash('kowala'));
    const logs2 = await proxyFactory.createProxy(admin, registrar.address, { from: admin });
    const logs3 = logs2.logs;
    const registrarProxyAddress = logs3.find(l => l.event === 'ProxyCreated').args.proxy;
    const registrarProxy = await AdminUpgradeabilityProxy.at(registrarProxyAddress);
    const registrarContract = await FIFSRegistrar.at(registrarProxyAddress);
    await registrarContract.initialize(knsProxyAddress, namehash('kowala'));

    // Resolver Proxy
    const resolver = await PublicResolver.new(knsProxyAddress, { from: admin });
    const logs4 = await proxyFactory.createProxy(admin, resolver.address);
    const logs5 = logs4.logs;
    const resolverProxyAddress = logs5.find(l => l.event === 'ProxyCreated').args.proxy;
    const resolverProxy = await AdminUpgradeabilityProxy.at(resolverProxyAddress);
    const resolverContract = await PublicResolver.at(resolverProxyAddress);
    
    await resolverContract.initialize(knsProxyAddress);
    await knsContract.setSubnodeOwner(0, web3.sha3('kowala'), registrarProxyAddress, { from: owner });
    const validator = await ValidatorMgr.new(1, 2, 3, '0x1234', 1);
    await registrarContract.register(web3.sha3('validator'), owner);
    await knsContract.setResolver(namehash('validator.kowala'), resolverProxyAddress, { from: owner });
    await resolverContract.setAddr(namehash('validator.kowala'), validator.address, { from: owner });
    const knsv1 = await KNSV1.new();
    await knsProxy.upgradeTo(knsv1.address, { from: admin });
    knsContract = await KNSV1.at(knsProxyAddress);

    const hello = await knsContract.helloProxy();
    await hello.should.be.equal('HelloProxy');

    const resolverStorage = await knsContract.resolver(namehash('validator.kowala'));
    await resolverStorage.should.be.equal(resolverProxyAddress);

    const validatorEnsAddr = await resolverContract.addr(namehash('validator.kowala'));
    await validatorEnsAddr.should.be.equal(validator.address);
  });
});
