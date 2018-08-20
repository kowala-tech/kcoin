/* global artifacts, contract, it, beforeEach, describe, before, web3 */
/* eslint no-unused-expressions: 0 */
/* eslint consistent-return: 0 */
/* eslint-disable max-len */

process.env.NODE_ENV = 'test';

require('chai')
  .use(require('chai-as-promised'))
  .use(require('chai-bignumber')(web3.BigNumber))
  .should();

const {
  EVMError,
} = require('../helpers/testUtils.js');

const ValidatorMgr = artifacts.require('ValidatorMgr.sol');
const PublicResolver = artifacts.require('PublicResolver.sol');
const KNS = artifacts.require('KNSRegistry.sol');
const FIFSRegistrar = artifacts.require('FIFSRegistrar.sol');
const MiningTokenMock = artifacts.require('TokenMock.sol');
const DomainResolverMock = artifacts.require('DomainResolverMock.sol');
const namehash = require('eth-ens-namehash');

contract('Validator Manager', ([_, owner, newOwner, newOwner2, newOwner3, notOwner]) => {
  describe('KNS functionality', async () => {
    beforeEach(async () => {
      this.kns = await KNS.new({ from: owner });
      this.registrar = await FIFSRegistrar.new(this.kns.address, namehash('kowala'));
      this.resolver = await PublicResolver.new(this.kns.address);

      await this.kns.setSubnodeOwner(0, web3.sha3('kowala'), this.registrar.address, { from: owner });
      await this.registrar.register(web3.sha3('miningtoken'), owner, { from: owner });
      await this.kns.setResolver(namehash('miningtoken.kowala'), this.resolver.address, { from: owner });
      this.miningToken = await MiningTokenMock.new();
      await this.resolver.setAddr(namehash('miningtoken.kowala'), this.miningToken.address, { from: owner });
      this.validator = await ValidatorMgr.new(100, 100, 0, 200, this.resolver.address, { from: owner });
    });

    it('should set MiningToken Address from KNS during creation', async () => {
      // given
      const knsResolverAddr = await this.validator.knsResolver();
      const resolver = await PublicResolver.at(knsResolverAddr);

      // when
      const miningTokenAddr = await resolver.addr(namehash('miningtoken.kowala'));

      // then
      await miningTokenAddr.should.be.equal(this.miningToken.address);
    });

    it('should set MiningToken Address from KNS during creation', async () => {
      // given
      const knsResolverAddr = await this.validator.knsResolver();
      const resolver = await PublicResolver.at(knsResolverAddr);

      // when
      const miningTokenAddr = await resolver.addr(namehash('miningtoken.kowala'));

      // then
      await miningTokenAddr.should.be.equal(this.miningToken.address);
    });
  });
  describe('Functionality of', async () => {
    beforeEach(async () => {
      this.miningToken = await MiningTokenMock.new();
      this.resolver = await DomainResolverMock.new(this.miningToken.address);
      this.validator = await ValidatorMgr.new(100, 3, 0, 3, this.resolver.address, { from: owner });
    });

    describe('registration', async () => {
      it('should register validator', async () => {
        // when
        await this.validator.registerValidator(newOwner, 100, { from: newOwner });

        // then
        const validatorCount = await this.validator.getValidatorCount();
        await validatorCount.should.be.bignumber.equal(1);
      });

      it('should not register validator when deposit is lower than base deposit', async () => {
        // when
        const exptectedFailedRegistration = this.validator.registerValidator(newOwner, 99, { from: newOwner });

        // then
        await exptectedFailedRegistration.should.eventually.be.rejectedWith(EVMError('revert'));
      });

      it('should not register validator with the same address', async () => {
        // given
        await this.validator.registerValidator(newOwner, 100, { from: newOwner });

        // when
        const exptectedFailedRegistration = this.validator.registerValidator(newOwner, 100, { from: newOwner });

        // then
        await exptectedFailedRegistration.should.eventually.be.rejectedWith(EVMError('revert'));
      });
    });

    describe('deregistration', async () => {
      it('should deregister validator', async () => {
        // given
        await this.validator.registerValidator(newOwner, 100, { from: newOwner });

        // when
        await this.validator.deregisterValidator({ from: newOwner });

        // then
        const validatorCount = await this.validator.getValidatorCount();
        await validatorCount.should.be.bignumber.equal(0);
      });

      it('should not deregister deregistered validator', async () => {
        // given
        await this.validator.registerValidator(newOwner, 100, { from: newOwner });

        // when
        await this.validator.deregisterValidator({ from: newOwner });
        const exptectedFailedDeregistration = this.validator.deregisterValidator({ from: newOwner });

        // then
        await exptectedFailedDeregistration.should.eventually.be.rejectedWith(EVMError('revert'));
      });

      it('should not deregister non-validator', async () => {
        // when
        const exptectedFailedDeregistration = this.validator.deregisterValidator({ from: notOwner });

        // then
        await exptectedFailedDeregistration.should.eventually.be.rejectedWith(EVMError('revert'));
      });
    });

    describe('release deposits', async () => {
      it('should release deposit', async () => {
        // given
        const initialBalance = await this.miningToken.balanceOf(newOwner, { from: newOwner });
        await this.validator.registerValidator(newOwner, 100, { from: newOwner });
        await this.validator.deregisterValidator({ from: newOwner });

        // when
        await this.validator.releaseDeposits({ from: newOwner });

        // then
        const balanceAfterRelease = await this.miningToken.balanceOf(newOwner);
        await initialBalance.should.be.bignumber.equal(0);
        await balanceAfterRelease.should.be.bignumber.equal(100);
      });

      it('should not release deposit twice', async () => {
        // given
        await this.validator.registerValidator(newOwner, 100, { from: newOwner });
        await this.validator.deregisterValidator({ from: newOwner });
        await this.validator.releaseDeposits({ from: newOwner });
        const balanceFirstAfterRelease = await this.miningToken.balanceOf(newOwner);

        // when
        await this.validator.releaseDeposits({ from: newOwner });
        const balanceAfterSecondRelease = await this.miningToken.balanceOf(newOwner);

        // then
        await balanceAfterSecondRelease.should.be.bignumber.equal(balanceFirstAfterRelease);
      });
    });
    describe('setters and getters', async () => {
      it('should get validator`s deposit', async () => {
        // given
        await this.validator.registerValidator(newOwner, 150, { from: newOwner });

        // when
        const validatorDeposit = await this.validator.getDepositAtIndex(0, { from: newOwner });

        // then
        await validatorDeposit[0].should.be.bignumber.equal(150);
      });

      it('should not get validator`s deposit when calling by not owner', async () => {
        // given
        await this.validator.registerValidator(newOwner, 150, { from: newOwner });

        // when
        const expectedValidatorDepositFailure = this.validator.getDepositAtIndex(0, { from: notOwner });

        // then
        await expectedValidatorDepositFailure.should.eventually.be.rejectedWith(EVMError('invalid opcode'));
      });

      it('should get validator at index', async () => {
        // given
        await this.validator.registerValidator(newOwner, 100, { from: newOwner });

        // when
        const validatorAtIndex = await this.validator.getValidatorAtIndex(0, { from: newOwner });

        // then
        await validatorAtIndex[0].should.be.equal(newOwner);
      });

      it('should get minimal required deposit when there are no positions left', async () => {
        // given
        await this.validator.registerValidator(newOwner, 130, { from: newOwner });
        await this.validator.registerValidator(newOwner2, 120, { from: newOwner });
        await this.validator.registerValidator(newOwner3, 110, { from: newOwner });

        // when
        const minDeposit = await this.validator.getMinimumDeposit();

        // then
        await minDeposit.should.be.bignumber.equal(111);
      });

      it('should get base deposit when there are still positions available', async () => {
        // given
        await this.validator.registerValidator(newOwner, 130, { from: newOwner });
        await this.validator.registerValidator(newOwner2, 120, { from: newOwner });

        // when
        const minDeposit = await this.validator.getMinimumDeposit();

        // then
        await minDeposit.should.be.bignumber.equal(100);
      });

      it('should set base deposit', async () => {
        // given
        const currentBaseDeposit = await this.validator.getMinimumDeposit();

        // when
        await this.validator.setBaseDeposit(150, { from: owner });
        const newBaseDeposit = await this.validator.getMinimumDeposit();

        // then
        await currentBaseDeposit.should.be.bignumber.equal(100);
        await newBaseDeposit.should.be.bignumber.equal(150);
      });

      it('should not set base deposit by not owner', async () => {
        // when
        const exptectedFailedDepositSet = this.validator.setBaseDeposit(150, { from: notOwner });

        // then
        await exptectedFailedDepositSet.should.eventually.be.rejectedWith(EVMError('revert'));
      });
    });
  });
});
