/* global artifacts, contract, it, beforeEach, describe, before, web3 */
/* eslint no-unused-expressions: 0 */
/* eslint consistent-return: 0 */
/* eslint-disable max-len */

process.env.NODE_ENV = 'test';

require('chai')
  .use(require('chai-as-promised'))
  .use(require('chai-bignumber')(web3.BigNumber))
  .should();

const ValidatorMgr = artifacts.require('ValidatorMgr.sol');
const { EVMError } = require('../helpers/testUtils.js');

contract('Validator Manager', ([_, owner, newOwner, notOwner]) => {
  beforeEach(async () => {
    this.validator = await ValidatorMgr.new(1, 2, 3, '0x1234', 1, { from: owner });
  });

  it('should set Validator owner during creation', async () => {
    // when
    const validatorOwner = await this.validator.owner();

    // then
    await validatorOwner.should.be.equal(owner);
  });

  it('should transfer ownership by a owner', async () => {
    // when
    await this.validator.transferOwnership(newOwner, { from: owner });
    const validatorOwner = await this.validator.owner();
    // then
    await validatorOwner.should.be.equal(newOwner);
  });

  it('should not transfer ownership by not a owner', async () => {
    // when
    const onwershipTransfer = this.validator.transferOwnership(newOwner);

    // then
    await onwershipTransfer.should.eventually.be.rejectedWith(EVMError('revert'));
  });

  it('should change mining token address by a owner', async () => {
    // given
    const newAddress = '0x0987';

    // when
    await this.validator.changeAddresOfMiningToken(newAddress, { from: owner });

    // then
    const miningTokenAddr = await this.validator.miningTokenAddr({ from: owner });
    await miningTokenAddr.should.be.equal('0x0000000000000000000000000000000000000987');
  });

  it('should not change mining token address by not a owner', async () => {
    // given
    const newAddress = '0x0987';
    
    // when
    const miningTokenAddr = this.validator.changeAddresOfMiningToken(newAddress, { from: notOwner });

    // then
    await miningTokenAddr.should.be.eventually.rejectedWith(EVMError('revert'));
  });
});
