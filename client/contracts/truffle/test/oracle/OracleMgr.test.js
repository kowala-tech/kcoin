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

const OracleMgr = artifacts.require('OracleMgr.sol');
const ConsensusMock = artifacts.require('ConsensusMock.sol');
const PublicResolver = artifacts.require('PublicResolver.sol');
const KNS = artifacts.require('KNSRegistry.sol');
const FIFSRegistrar = artifacts.require('FIFSRegistrar.sol');
const UpgradeabilityProxy = Contracts.getFromNodeModules('zos-lib', 'UpgradeabilityProxyFactory');
const AdminUpgradeabilityProxy = Contracts.getFromNodeModules('zos-lib', 'AdminUpgradeabilityProxy');
const namehash = require('eth-ens-namehash');
const { EVMError } = require('../helpers/testUtils.js');

contract('Oracle Manager', ([_, admin, owner, newOwner, notOwner]) => {
  beforeEach(async () => {
    this.proxyFactory = await UpgradeabilityProxy.new();
    this.knsContract = await KNS.new({ from: owner });
    this.logs = await this.proxyFactory.createProxy(admin, this.knsContract.address, { from: admin });
    this.logs1 = this.logs.logs;
    this.knsProxyAddress = this.logs1.find(l => l.event === 'ProxyCreated').args.proxy;
    this.knsProxy = await AdminUpgradeabilityProxy.at(this.knsProxyAddress);
    this.kns = new KNS(this.knsProxyAddress);
    await this.kns.initialize(owner);
    this.registrar = await FIFSRegistrar.new(this.knsProxyAddress, namehash('kowala'));
    this.resolver = await PublicResolver.new(this.knsProxyAddress);

    await this.kns.setSubnodeOwner(0, web3.sha3('kowala'), this.registrar.address, { from: owner });
    await this.registrar.register(web3.sha3('validator'), owner, { from: owner });
    await this.kns.setResolver(namehash('validator.kowala'), this.resolver.address, { from: owner });
    this.consensus = await ConsensusMock.new(true);
    await this.resolver.setAddr(namehash('validator.kowala'), this.consensus.address, { from: owner });
    this.oracle = await OracleMgr.new(1, 1, 1, this.consensus.address, this.knsProxyAddress, { from: owner });
  });

  it('should set Consensus address using KNS', async () => {
    // given
    const knsResolverAddr = await this.oracle.knsResolver();
    const resolver = await PublicResolver.at(knsResolverAddr);

    // when
    await this.oracle.registerOracle({ from: owner });
    const numberOfOracles = await this.oracle.getOracleCount();

    // then
    await numberOfOracles.should.be.bignumber.equal(1);
  });
});
