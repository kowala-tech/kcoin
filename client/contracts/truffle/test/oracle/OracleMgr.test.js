/* global artifacts, contract, it, beforeEach, describe, before, web3 */
/* eslint no-unused-expressions: 0 */
/* eslint consistent-return: 0 */
/* eslint-disable max-len */

process.env.NODE_ENV = 'test';

require('chai')
  .use(require('chai-as-promised'))
  .use(require('chai-bignumber')(web3.BigNumber))
  .should();

const OracleMgr = artifacts.require('OracleMgr.sol');
const ValidatorMgr = artifacts.require('ValidatorMgr.sol');
const { EVMError } = require('../helpers/testUtils.js');

contract('Oracle Manager', ([_, owner, newOwner, notOwner]) => {
  beforeEach(async () => {
    this.validator = await ValidatorMgr.new(1, 2, 3, '0x1234', 1);
    this.oracle = await OracleMgr.new(1, 1, 1, 1, 1, 1, this.validator.address, { from: owner });
  });

  it('should set Oracle owner during creation', async () => {
    // when
    const oracleOwner = await this.oracle.owner();

    // then
    await oracleOwner.should.be.equal(owner);
  });

  it('should transfer ownership by a owner', async () => {
    // when
    await this.oracle.transferOwnership(newOwner, { from: owner });
    const oracleOwner = await this.oracle.owner();
    // then
    await oracleOwner.should.be.equal(newOwner);
  });

  it('should not transfer ownership by not a owner', async () => {
    // when
    const onwershipTransfer = this.oracle.transferOwnership(newOwner);

    // then
    await onwershipTransfer.should.eventually.be.rejectedWith(EVMError('revert'));
  });

  it('should change validator address by a owner', async () => {
    // given
    const newValidator = await ValidatorMgr.new(1, 2, 3, '0x1234', 1);

    // when
    await this.oracle.changeValidator(newValidator.address, { from: owner });

    // then
    const validatorAddress = await this.oracle.getValidatorAddress({ from: owner });
    await validatorAddress.should.be.equal(newValidator.address);
  });

  it('should not change validator address by not a owner', async () => {
    // given
    const newValidator = await ValidatorMgr.new(1, 2, 3, '0x1234', 1);

    // when
    const changeValidator = this.oracle.changeValidator(newValidator.address, { from: notOwner });

    // then
    await changeValidator.should.be.eventually.rejectedWith(EVMError('revert'));
  });
});
