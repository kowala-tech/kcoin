/* global artifacts, contract, it, beforeEach, describe, cd, web3 */
/* eslint no-unused-expressions: 0 */
/* eslint consistent-return: 0 */
/* eslint-disable max-len */

process.env.NODE_ENV = 'test';

require('chai')
  .use(require('chai-as-promised'))
  .use(require('chai-bignumber')(web3.BigNumber))
  .should();

const { EVMError } = require('../helpers/testUtils.js');

const ExchangeMgr = artifacts.require('ExchangeMgr.sol');

contract('ExchangeMgr', ([_, owner, newOwner, notOwner]) => {
  describe('Functionality of', async () => {
    beforeEach(async () => {
      this.exchange = await ExchangeMgr.new({ from: owner });
    });
    describe('Adding Exchange', async () => {
      it('should add exchange', async () => {
        // when
        await this.exchange.addExchange('Kowala', { from: owner });

        // then
        const isExchange = await this.exchange.isExchange('Kowala');
        await isExchange.should.be.true;
      });

      it('should not add exchange which is already added', async () => {
        // given
        await this.exchange.addExchange('Kowala', { from: owner });

        // when
        const expectedFailure = this.exchange.addExchange('Kowala');

        // then
        await expectedFailure.should.be.rejectedWith(EVMError('revert'));
      });

      it('should not add exchange to whitelist by not a owner', async () => {
        // when
        const expectedFailure = this.exchange.addExchange('Kowala');

        // then
        await expectedFailure.should.be.rejectedWith(EVMError('revert'));
      });

      it('should add exchange and that exchange should be whitelisted automatically', async () => {
        // when
        await this.exchange.addExchange('Kowala', { from: owner });

        // then
        const isWhitelistedExchange = await this.exchange.isWhitelistedExchange('Kowala');
        await isWhitelistedExchange.should.be.true;
      });
    });

    describe('Remove Exchange', async () => {
      it('should remove exchange', async () => {
        // given
        await this.exchange.addExchange('Kowala', { from: owner });

        // when
        await this.exchange.removeExchange('Kowala', { from: owner });

        // then
        const isExchange = await this.exchange.isExchange('Kowala');
        await isExchange.should.be.false;
      });

      it('should not remove non exchange', async () => {
        // when
        const expectedFail = this.exchange.removeExchange('Kowala', { from: owner });

        // then
        await expectedFail.should.be.rejectedWith(EVMError('revert'));
      });

      it('should not remove exchange by not owner', async () => {
        // given
        await this.exchange.addExchange('Kowala', { from: owner });

        // when
        const expectedFail = this.exchange.removeExchange('Kowala', { from: notOwner });

        // then
        await expectedFail.should.be.rejectedWith(EVMError('revert'));
      });
    });

    describe('Blackisting an Exchange', async () => {
      it('should blacklist an exchange', async () => {
        // given
        await this.exchange.addExchange('Kowala', { from: owner });

        // when
        await this.exchange.blacklistExchange('Kowala', { from: owner });

        // then
        const isBlacklisted = await this.exchange.isBlacklistedExchange('Kowala');
        await isBlacklisted.should.be.true;
      });

      it('should not blacklist an exchange via not owner', async () => {
        // given
        await this.exchange.addExchange('Kowala', { from: owner });

        // when
        const expectedFail = this.exchange.blacklistExchange('Kowala', { from: notOwner });

        // then
        await expectedFail.should.be.rejectedWith(EVMError('revert'));
      });

      it('should not blacklist not whitelisted exchange', async () => {
        // when
        const expectedFail = this.exchange.blacklistExchange('Kowala', { from: owner });

        // then
        await expectedFail.should.be.rejectedWith(EVMError('revert'));
      });

      it('should blacklist an exchange', async () => {
        // given
        await this.exchange.addExchange('Kowala', { from: owner });
        await this.exchange.blacklistExchange('Kowala', { from: owner });

        // when
        const expectedFail = this.exchange.blacklistExchange('Kowala', { from: owner });

        // then
        await expectedFail.should.be.rejectedWith(EVMError('revert'));
      });
    });

    describe('Whitelisting an Exchange', async () => {
      it('should whitelist an exchange', async () => {
        // given
        await this.exchange.addExchange('Kowala', { from: owner });
        await this.exchange.blacklistExchange('Kowala', { from: owner });

        // when
        await this.exchange.whitelistExchange('Kowala', { from: owner });

        // then
        const isWhitelisted = await this.exchange.isWhitelistedExchange('Kowala');
        await isWhitelisted.should.be.true;
      });

      it('should not whitelist an exchange via not owner', async () => {
        // given
        await this.exchange.addExchange('Kowala', { from: owner });
        await this.exchange.blacklistExchange('Kowala', { from: owner });

        // when
        const expectedFail = this.exchange.whitelistExchange('Kowala', { from: notOwner });

        // then
        await expectedFail.should.be.rejectedWith(EVMError('revert'));
      });

      it('should not blacklist not whitelisted exchange', async () => {
        // when
        const expectedFail = this.exchange.whitelistExchange('Kowala', { from: owner });

        // then
        await expectedFail.should.be.rejectedWith(EVMError('revert'));
      });

      it('should blacklist an exchange', async () => {
        // given
        await this.exchange.addExchange('Kowala', { from: owner });
        await this.exchange.blacklistExchange('Kowala', { from: owner });
        await this.exchange.whitelistExchange('Kowala', { from: owner });

        // when
        const expectedFail = this.exchange.whitelistExchange('Kowala', { from: owner });

        // then
        await expectedFail.should.be.rejectedWith(EVMError('revert'));
      });
    });

    describe('Getters and Setters', async () => {
      it('should get whitelisted count', async () => {
        // given
        await this.exchange.addExchange('Kowala', { from: owner });
        await this.exchange.addExchange('Tech', { from: owner });

        // when
        const count = await this.exchange.getWhitelistedExchangeCount();

        // then
        await count.should.be.bignumber.equal(2);
      });

      it('should get whitelisted exchange at index', async () => {
        // given
        await this.exchange.addExchange('Kowala', { from: owner });
        await this.exchange.addExchange('Tech', { from: owner });

        // when
        const exchange = await this.exchange.getWhitelistedExchangeAtIndex(1);

        // then
        await exchange.should.be.equal('Tech');
      });

      it('should fail getting whitelisted exchange at index out of bounds', async () => {
        // given
        await this.exchange.addExchange('Kowala', { from: owner });
        await this.exchange.addExchange('Tech', { from: owner });

        // when
        const expectedFail = this.exchange.getWhitelistedExchangeAtIndex(2);

        // then
        await expectedFail.should.be.rejectedWith(EVMError('invalid opcode'));
      });
    });
  });
});
