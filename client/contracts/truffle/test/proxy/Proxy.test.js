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
const ValidatorMgr1 = artifacts.require('ValidatorMgr1.sol');
const UpgradeabilityProxy = Contracts.getFromNodeModules('zos-lib', 'UpgradeabilityProxyFactory');
const AdminUpgradeabilityProxy = Contracts.getFromNodeModules('zos-lib', 'AdminUpgradeabilityProxy');
const namehash = require('eth-ens-namehash');

contract('Proxy Functionality', ([_, admin, owner, anotherAccount]) => {
  it('should access contracts via proxy', async () => {
    const proxyFactory = await UpgradeabilityProxy.new();

    // KNS Proxy
    const kns = await KNS.new();
    let logs = await proxyFactory.createProxy(admin, kns.address, { from: admin });
    const logs1 = logs.logs;
    const knsProxyAddress = logs1.find(l => l.event === 'ProxyCreated').args.proxy;
    const knsProxy = await AdminUpgradeabilityProxy.at(knsProxyAddress);
    let knsContract = new KNS(knsProxyAddress);
    await knsContract.initialize(owner);

    // Registrar Proxy
    const registrar = await FIFSRegistrar.new(knsProxyAddress, namehash('kowala'));
    logs = await proxyFactory.createProxy(admin, registrar.address, { from: admin });
    const logs2 = logs.logs;
    const registrarProxyAddress = logs2.find(l => l.event === 'ProxyCreated').args.proxy;
    const registrarProxy = await AdminUpgradeabilityProxy.at(registrarProxyAddress);
    const registrarContract = await FIFSRegistrar.at(registrarProxyAddress);
    await registrarContract.initialize(knsProxyAddress, namehash('kowala'));

    // Resolver Proxy
    const resolver = await PublicResolver.new(knsProxyAddress, { from: admin });
    logs = await proxyFactory.createProxy(admin, resolver.address);
    const logs3 = logs.logs;
    const resolverProxyAddress = logs3.find(l => l.event === 'ProxyCreated').args.proxy;
    const resolverProxy = await AdminUpgradeabilityProxy.at(resolverProxyAddress);
    const resolverContract = await PublicResolver.at(resolverProxyAddress);
    await resolverContract.initialize(knsProxyAddress);

    // validator proxy
    const validator = await ValidatorMgr.new(1, 2, 3, 1, resolverProxyAddress, { from: admin });
    logs = await proxyFactory.createProxy(admin, validator.address);
    const logs4 = logs.logs;
    const validatorProxyAddress = logs4.find(l => l.event === 'ProxyCreated').args.proxy;
    const validatorProxy = await AdminUpgradeabilityProxy.at(validatorProxyAddress);
    let validatorContract = await ValidatorMgr.at(validatorProxyAddress);
    await validatorContract.initialize(1, 2, 3, 1, resolverProxyAddress);

    // 
    await knsContract.setSubnodeOwner(0, web3.sha3('kowala'), registrarProxyAddress, { from: owner });
    await registrarContract.register(web3.sha3('validator'), owner);
    await knsContract.setResolver(namehash('validator.kowala'), resolverProxyAddress, { from: owner });
    await resolverContract.setAddr(namehash('validator.kowala'), validatorProxyAddress, { from: owner });
    
    await validatorContract.registerValidator(anotherAccount, 2);
    const tmp = await validatorContract.getValidatorAtIndex(0);

    const validator1 = await ValidatorMgr1.new();
    await validatorProxy.upgradeTo(validator1.address, { from: admin });
    validatorContract = await ValidatorMgr1.at(await resolverContract.addr(namehash('validator.kowala')));

    const hello1 = await validatorContract.helloProxy();
    await hello1.should.be.equal('HelloProxy');

    const tmp1 = await validatorContract.getValidatorAtIndex(0);

    const knsv1 = await KNSV1.new();
    await knsProxy.upgradeTo(knsv1.address, { from: admin });
    knsContract = await KNSV1.at(knsProxyAddress);

    const hello = await knsContract.helloProxy();
    await hello.should.be.equal('HelloProxy');

    const resolverStorage = await knsContract.resolver(namehash('validator.kowala'));
    await resolverStorage.should.be.equal(resolverProxyAddress);

    const validatorEnsAddr = await resolverContract.addr(namehash('validator.kowala'));
    await validatorEnsAddr.should.be.equal(validatorProxyAddress);
  });
});
