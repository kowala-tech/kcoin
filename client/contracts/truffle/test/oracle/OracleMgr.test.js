/* global artifacts, contract, it, beforeEach, describe, before, web3 */
/* eslint no-unused-expressions: 0 */
/* eslint consistent-return: 0 */
/* eslint-disable max-len */

process.env.NODE_ENV = 'test';

require('chai')
  .use(require('chai-as-promised'))
  .use(require('chai-bignumber')(web3.BigNumber))
  .should();
const namehash = require('eth-ens-namehash');
const { EVMError } = require('../helpers/testUtils.js');
const { Contracts } = require('zos-lib');

const OracleMgr = artifacts.require('OracleMgr.sol');
const ConsensusMock = artifacts.require('ConsensusMock.sol');
const DomainResolverMock = artifacts.require('DomainResolverMock.sol');
const PublicResolver = artifacts.require('PublicResolver.sol');
const KNS = artifacts.require('KNSRegistry.sol');
const FIFSRegistrar = artifacts.require('FIFSRegistrar.sol');
const UpgradeabilityProxy = Contracts.getFromNodeModules('zos-lib', 'UpgradeabilityProxyFactory');
const AdminUpgradeabilityProxy = Contracts.getFromNodeModules('zos-lib', 'AdminUpgradeabilityProxy');

contract('OracleMgr', ([_, admin, owner, newOwner, newOwner2, newOwner3, newOwner4, notOwner]) => {
  describe('KNS functionality', async () => {
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
      this.oracle = await OracleMgr.new(1, 1, 1, this.resolver.address, { from: owner });
    });

    it('should set Consensus address using KNS', async () => {
      // given
      const knsResolverAddr = await this.oracle.knsResolver();
      const resolver = await PublicResolver.at(knsResolverAddr);

      // when
      const consensusAddr = await resolver.addr(namehash('validator.kowala'));

      // then
      await consensusAddr.should.be.equal(this.consensus.address);
    });
  });

  describe('Functionality of', async () => {
    beforeEach(async () => {
      this.consensus = await ConsensusMock.new(true);
      this.resolver = await DomainResolverMock.new(this.consensus.address);
      this.oracle = await OracleMgr.new(3, 1, 1, this.resolver.address, { from: owner });
    });
    describe('registration', async () => {
      it('should register oracle', async () => {
        // when
        await this.oracle.registerOracle({ from: newOwner });

        // then
        const oracleCount = await this.oracle.getOracleCount();
        await oracleCount.should.be.bignumber.equal(1);
      });

      it('should not register already registered oracle', async () => {
        // given
        await this.oracle.registerOracle({ from: newOwner });

        // when
        const expectedRegistrationFailure = this.oracle.registerOracle({ from: newOwner });

        // then
        await expectedRegistrationFailure.should.eventually.be.rejectedWith(EVMError('revert'));
      });

      it('should not register oracle above oracle limit', async () => {
        // given
        await this.oracle.registerOracle({ from: newOwner });
        await this.oracle.registerOracle({ from: newOwner2 });
        await this.oracle.registerOracle({ from: newOwner3 });

        // when
        const expectedRegistrationFailure = this.oracle.registerOracle({ from: newOwner4 });

        // then
        await expectedRegistrationFailure.should.eventually.be.rejectedWith(EVMError('revert'));
      });

      it('should not register oracle when not super node', async () => {
        const consensus = await ConsensusMock.new(false);
        const resolver = await DomainResolverMock.new(consensus.address);
        const oracleWithoutSuperNode = await OracleMgr.new(3, 1, 1, resolver.address, { from: owner });
        // when
        const expectedRegistrationFailure = oracleWithoutSuperNode.registerOracle({ from: newOwner });

        // then
        await expectedRegistrationFailure.should.eventually.be.rejectedWith(EVMError('revert'));
      });
    });

    describe('deregistration', async () => {
      it('should deregister oracle', async () => {
        // given
        await this.oracle.registerOracle({ from: newOwner });

        // when
        await this.oracle.deregisterOracle({ from: newOwner });

        // then
        const oracleCount = await this.oracle.getOracleCount();
        await oracleCount.should.be.bignumber.equal(0);
      });

      it('should not deregister deregistered oracle', async () => {
        // given
        await this.oracle.registerOracle({ from: newOwner });
        await this.oracle.deregisterOracle({ from: newOwner });

        // when
        const exptectedFailedDeregistration = this.oracle.deregisterOracle({ from: newOwner });

        // then
        await exptectedFailedDeregistration.should.eventually.be.rejectedWith(EVMError('revert'));
      });

      it('should not deregister non-oracle', async () => {
        // when
        const exptectedFailedDeregistration = this.oracle.deregisterOracle({ from: notOwner });

        // then
        await exptectedFailedDeregistration.should.eventually.be.rejectedWith(EVMError('revert'));
      });
    });

    describe('setters and getters', async () => {
      it('should have proper values at creation', async () => {
        // when
        const maxNumOracles = await this.oracle.maxNumOracles();
        const syncFrequency = await this.oracle.syncFrequency();
        const updatePeriod = await this.oracle.updatePeriod();
        const knsResolver = await this.oracle.knsResolver();

        // then
        await maxNumOracles.should.be.bignumber.equal(3);
        await syncFrequency.should.be.bignumber.equal(1);
        await updatePeriod.should.be.bignumber.equal(1);
        await knsResolver.should.be.equal(this.resolver.address);
      });

      it('should get oracle`s price', async () => {
        // given
        await this.oracle.registerOracle({ from: newOwner });
        await this.oracle.submitPrice(10, { from: newOwner });

        // when
        const oraclePrice = await this.oracle.getPriceAtIndex(0, { from: newOwner });
        // then
        await oraclePrice[0].should.be.bignumber.equal(10);
      });

      it('should get oracle at index', async () => {
        // given
        await this.oracle.registerOracle({ from: newOwner });

        // when
        const oracleAtIndex = await this.oracle.getOracleAtIndex(0, { from: newOwner });

        // then
        await oracleAtIndex.should.be.equal(newOwner);
      });

      it('should not set price twice', async () => {
        // given
        await this.oracle.registerOracle({ from: newOwner });
        await this.oracle.submitPrice(10, { from: newOwner });

        // when
        const expectedPriceFailure = this.oracle.submitPrice(10, { from: newOwner });

        // then
        await expectedPriceFailure.should.eventually.be.rejectedWith(EVMError('revert'));
      });

      it('should not set price if not oracle', async () => {
        // given
        await this.oracle.registerOracle({ from: newOwner });

        // when
        const expectedPriceFailure = this.oracle.submitPrice(10, { from: notOwner });

        // then
        await expectedPriceFailure.should.eventually.be.rejectedWith(EVMError('revert'));
      });

      it('should get price count', async () => {
        // given
        await this.oracle.registerOracle({ from: newOwner });
        await this.oracle.registerOracle({ from: newOwner2 });

        // when
        await this.oracle.submitPrice(10, { from: newOwner });
        await this.oracle.submitPrice(15, { from: newOwner2 });

        // then
        const priceCount = await this.oracle.getPriceCount();
        await priceCount.should.be.bignumber.equal(2);
      });
    });
  });
});
