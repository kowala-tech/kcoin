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

const Token = artifacts.require('MintableToken.sol');

contract('Mining Token', ([_, owner, newOwner, notOwner]) => {
  beforeEach(async () => {
    this.token = await Token.new({ from: owner });
  });
  it('Should transfer tokens to new owner', async () => {
    // given
    await this.token.mint(owner, 10, { from: owner });

    // when
    await this.token.transfer(newOwner, 5, { from: owner });

    // then
    const balance = await this.token.balanceOf(newOwner);
    await balance.should.be.bignumber.equal(5);
  });

  it('Should not transfer tokens to new owner from not a owner', async () => {
    // given
    await this.token.mint(owner, 10, { from: owner });

    // when
    const failedTransfer = this.token.transfer(newOwner, 5, { from: notOwner });

    // then
    await failedTransfer.should.eventually.be.rejectedWith(EVMError('revert'));
  });
});
