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
  EVMError, ether,
} = require('../helpers/testUtils.js');

const SysVar = artifacts.require('SystemVars.sol');

contract('System Vars', ([_, owner, anotherAccount]) => {
  beforeEach(async () => {
    this.sysvar = await SysVar.new(ether(10), 100, { from: owner });
  });

  it('Should get current system`s currency price', async () => {
    // when
    const price = await this.sysvar.price();

    // then
    await price.should.be.bignumber.equal(ether(10));
  });

  it('Should get mintedAmount', async () => {
    // when
    const mintedAmount = await this.sysvar.mintedAmount();

    // then
    await mintedAmount.should.be.bignumber.equal(1e12);
  });

  it('Should get oracle deduction', async () => {
    // when
    const deduction = await this.sysvar.oracleDeduction(100);

    // then
    await deduction.should.be.bignumber.equal(4);
  });

  it('Should get oracle reward', async () => {
    // when
    const reward = await this.sysvar.oracleReward();

    // then
    await reward.should.be.bignumber.equal(0);
  });
});

