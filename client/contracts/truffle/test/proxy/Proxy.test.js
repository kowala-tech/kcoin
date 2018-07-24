/* global artifacts, contract, it, beforeEach, web3 */
/* eslint no-unused-expressions: 0 */
/* eslint consistent-return: 0 */
/* eslint-disable max-len */

require('chai')
  .use(require('chai-as-promised'))
  .should();

const KNS = artifacts.require('KNSRegistry.sol');
// const KNSV1 = artifacts.require('KNSRegistryV1.sol');
// const FIFSRegistrar = artifacts.require('FIFSRegistrar.sol');
// const PublicResolver = artifacts.require('PublicResolver.sol');
// const ValidatorMgr = artifacts.require('ValidatorMgr.sol');
const UpgradeabilityProxy = artifacts.require('UpgradeabilityProxyFactory.sol');
const AdminUpgradeabilityProxy = artifacts.require('AdminUpgradeabilityProxy.sol');
// const namehash = require('eth-ens-namehash');

contract('Proxy Functionality', (accounts) => {
  // beforeEach(async () => {
  //   this.kns = await KNS.new();
  //   this.registrar = await FIFSRegistrar.new(this.kns.address, namehash('kowala'));
  //   this.resolver = await PublicResolver.new(this.kns.address);
  //   await this.kns.setSubnodeOwner(0, web3.sha3('kowala'), this.registrar.address, { from: accounts[0] });
  // });
  it('should access contracts via proxy', async () => {
    // const proxyFactory = await UpgradeabilityProxy.new();

    // KNS Proxy
    const kns = await KNS.new();
    await kns.initialize(accounts[0]);
    const proxy = await AdminUpgradeabilityProxy.new(kns.address, { from: accounts[0] });
    const proxyAddress = proxy.address;
    console.log('----1----');
    const knsContract = new KNS(proxyAddress);
    console.log('----2----');
    console.log(knsContract);
    await knsContract.owner('0x0', { from: accounts[0] });
    console.log('----3----');
    // const logs = await proxyFactory.createProxy(accounts[0], kns.address);
    // const logs1 = logs.logs;
    // let proxyAddress = logs1.find(l => l.event === 'ProxyCreated').args.proxy;
    // const knsProxy = await AdminUpgradeabilityProxy.at(proxyAddress);
    // let knsImplAddr = await knsProxy.implementation();
    // let knsContract = await KNS.at(proxyAddress);
    // await knsContract.initialize(accounts[0]);
    // console.log(await knsContract.owner(web3.sha3('kowala')));

    // // Registrar Proxy
    // const registrar = await FIFSRegistrar.new(knsContract.address, namehash('kowala'));
    // const logs2 = await proxyFactory.createProxy(accounts[0], this.registrar.address);
    // const logs3 = logs2.logs;
    // proxyAddress = logs3.find(l => l.event === 'ProxyCreated').args.proxy;
    // const registrarProxy = await AdminUpgradeabilityProxy.at(proxyAddress);
    // const registrarImplAddr = await registrarProxy.implementation();
    // const registrarContract = await FIFSRegistrar.at(registrarImplAddr);
    // await registrarContract.initialize(knsImplAddr, namehash('kowala'));

    // // await registrarImplAddr.should.be.equal(await knsContract.owner(namehash('kowala')));

    // // Resolver Proxy
    // const logs4 = await proxyFactory.createProxy(accounts[0], this.resolver.address);
    // const logs5 = logs4.logs;
    // proxyAddress = logs5.find(l => l.event === 'ProxyCreated').args.proxy;
    // const resolverProxy = await AdminUpgradeabilityProxy.at(proxyAddress);
    // const resolverImplAddr = await resolverProxy.implementation();
    // const resolverContract = await PublicResolver.at(resolverImplAddr);
    // await resolverContract.initialize(knsImplAddr);

    // const validator = await ValidatorMgr.new(1, 2, 3, '0x1234', 1);
    // await registrarContract.register(web3.sha3('validator'), accounts[0], { from: accounts[0] });
    // await knsContract.setResolver(namehash('validator.kowala'), resolverImplAddr);

    // await resolverContract.setAddr(namehash('validator.kowala'), validator.address);

    // // console.log(knsImplAddr);
    // const knsv1 = await KNSV1.new();
    // await knsProxy.upgradeTo(knsv1.address);
    // knsImplAddr = await knsProxy.implementation();
    // knsContract = await KNS.at(knsImplAddr);
    // // console.log(knsImplAddr);
    // console.log(await knsContract.owner(web3.sha3('kowala')));
    // // assert
    // // console.log("setResolver");
    // // await knsContract.setResolver(namehash('validator.kowala'), resolverImplAddr);
    // // console.log("assert setResolver");
    // const resolverStorage = await knsContract.resolver(namehash('validator.kowala'));
    // await resolverStorage.should.be.equal(resolverImplAddr);

    // const validatorEnsAddr = await resolverContract.addr(namehash('validator.kowala'));
    // await validatorEnsAddr.should.be.equal(validator.address);
  });
});
