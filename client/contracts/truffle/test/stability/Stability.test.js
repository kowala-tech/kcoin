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
  EVMError, ether
} = require('../helpers/testUtils.js');

const Stability = artifacts.require('Stability.sol');
const SysVar = artifacts.require('SystemVars.sol');

contract('Stability', ([_, owner, sub1, anotherAccount]) => {
  beforeEach(async () => {
    this.sysvar = await SysVar.new(ether(10), 100, { from: owner });
    this.stability = await Stability.new(ether(0.5), this.sysvar.address, { from: owner });
  });

  it('Should subscribe to stability contract service', async () => {
    // when
    await this.stability.subscribe({ from: sub1, value: ether(2) });

    // then
    const subCount = await this.stability.getSubscriptionCount();
    await subCount.should.be.bignumber.equal(1);
  });

  it('Should not subscribe to stability contract service when deposit is lower than minDeposit', async () => {
    // when
    const expectedSubscribeFailure = this.stability.subscribe({ from: sub1, value: ether(0.4) });

    // then
    await expectedSubscribeFailure.should.eventually.be.rejectedWith(EVMError('revert'));
  });

  it('Should unsubscribe from stability contract service', async () => {
    // given
    await this.stability.subscribe({ from: sub1, value: ether(1) });

    // when
    await this.stability.unsubscribe({ from: sub1 });

    // then
    const subCount = await this.stability.getSubscriptionCount();
    await subCount.should.be.bignumber.equal(0);
  });

  it('Should not unsubscribe from stability contract service when trying to unsubscribe different account than yours', async () => {
    // given
    await this.stability.subscribe({ from: sub1, value: ether(1) });

    // when
    const expectedUnsubscribeFailure = this.stability.unsubscribe({ from: anotherAccount });

    // then
    await expectedUnsubscribeFailure.should.eventually.be.rejectedWith(EVMError('revert'));
  });

  it('Should increase deposit after subscribing the same address', async () => {
    // given
    await this.stability.subscribe({ from: sub1, value: ether(1) });

    // when
    await this.stability.subscribe({ from: sub1, value: ether(1) });

    // then
    const subDeposit = await this.stability.getSubscriptionAtIndex(0);
    await subDeposit[1].should.be.bignumber.equal(ether(2));
  });
  it('Should not unsubscribe from stability contract service when trying to unsubscribe different account than yours', async () => {
    // given
    const sysvar = await SysVar.new(ether(0.5), 100, { from: owner });
    const stability = await Stability.new(ether(0.5), sysvar.address, { from: owner });
    await stability.subscribe({ from: sub1, value: ether(1) });

    // when
    const expectedUnsubscribeFailure = stability.unsubscribe({ from: anotherAccount });

    // then
    await expectedUnsubscribeFailure.should.eventually.be.rejectedWith(EVMError('revert'));
  });
});

