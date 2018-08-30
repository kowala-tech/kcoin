/* global artifacts, contract, it, beforeEach, describe, before, web3, assert */
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
const web3Abi = require('web3-eth-abi');

const Token = artifacts.require('MintableToken.sol');
const TokenMock = artifacts.require('TokenMock.sol');
const BalanceContract = artifacts.require('BalanceContract.sol');

contract('Token', ([_, owner, newOwner, notOwner]) => {
  const overloadedTransfer3ArgumentsAbi = {
    constant: false,
    inputs: [
      {
        name: '_to',
        type: 'address',
      },
      {
        name: '_value',
        type: 'uint256',
      },
      {
        name: '_data',
        type: 'bytes',
      },
    ],
    name: 'transfer',
    outputs: [
      {
        name: 'success',
        type: 'bool',
      },
    ],
    payable: false,
    stateMutability: 'nonpayable',
    type: 'function',
  };

  const overloadedTransfer4ArgumentsAbi = {
    constant: false,
    inputs: [
      {
        name: '_to',
        type: 'address',
      },
      {
        name: '_value',
        type: 'uint256',
      },
      {
        name: '_data',
        type: 'bytes',
      },
      {
        name: '_custom_fallback',
        type: 'string',
      },
    ],
    name: 'transfer',
    outputs: [
      {
        name: 'success',
        type: 'bool',
      },
    ],
    payable: false,
    stateMutability: 'nonpayable',
    type: 'function',
  };


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

  it('Should transfer tokens to a contract', async () => {
    // given
    const balanceContract = await BalanceContract.new({ from: owner });
    await this.token.mint(owner, 15, { from: owner });

    // when
    await this.token.transfer(balanceContract.address, 15, { from: owner });

    // then
    const balance = await balanceContract.value();
    await balance.should.be.bignumber.equal(15);
  });

  it('Should not transfer tokens to a contract without tokenFallback implementation', async () => {
    // given
    const tokenNotReceiver = await TokenMock.new({ from: owner });
    await this.token.mint(owner, 15, { from: owner });

    // when
    const failedTransfer = this.token.transfer(tokenNotReceiver.address, 15, { from: owner });

    // then
    await failedTransfer.should.eventually.be.rejectedWith(EVMError('revert'));
  });

  it('Should transfer tokens to new owner with additional data', async () => {
    // given
    await this.token.mint(owner, 10, { from: owner });

    const transferMethodTransactionData = web3Abi.encodeFunctionCall(
      overloadedTransfer3ArgumentsAbi,
      [
        newOwner,
        10,
        '0x00',
      ],
    );

    // when
    await web3.eth.sendTransaction({
      from: owner, to: this.token.address, data: transferMethodTransactionData, value: 0,
    });

    // then
    const balance = await this.token.balanceOf(newOwner);
    await balance.should.be.bignumber.equal(10);
  });

  it('Should not transfer tokens to new owner with additional data from not a owner', async () => {
    // when
    await this.token.mint(owner, 10, { from: owner });

    const transferMethodTransactionData = web3Abi.encodeFunctionCall(
      overloadedTransfer3ArgumentsAbi,
      [
        newOwner,
        10,
        '0x00',
      ],
    );

    // then
    await web3.eth.sendTransaction({
      from: notOwner, to: this.token.address, data: transferMethodTransactionData, value: 0,
    }, (error) => {
      if (error) { assert.equal(error, 'Error: VM Exception while processing transaction: revert'); }
    });
  });

  it('Should transfer tokens to new owner with additional data and fallback', async () => {
    // given
    await this.token.mint(owner, 10, { from: owner });

    const transferMethodTransactionData = web3Abi.encodeFunctionCall(
      overloadedTransfer4ArgumentsAbi,
      [
        newOwner,
        10,
        '0x00',
        'String Fallback',
      ],
    );

    // when
    await web3.eth.sendTransaction({
      from: owner, to: this.token.address, data: transferMethodTransactionData, value: 0,
    });

    // then
    const balance = await this.token.balanceOf(newOwner);
    await balance.should.be.bignumber.equal(10);
  });

  it('Should not transfer tokens to new owner with additional data and fallback from not a owner', async () => {
    // when
    await this.token.mint(owner, 10, { from: owner });

    const transferMethodTransactionData = web3Abi.encodeFunctionCall(
      overloadedTransfer4ArgumentsAbi,
      [
        newOwner,
        10,
        '0x00',
        'String Fallback',
      ],
    );

    // then
    await web3.eth.sendTransaction({
      from: notOwner, to: this.token.address, data: transferMethodTransactionData, value: 0,
    }, (error) => {
      if (error) { assert.equal(error, 'Error: VM Exception while processing transaction: revert'); }
    });
  });

  it('Should finish minting', async () => {
    // given
    await this.token.mint(owner, 5, { from: owner });

    // when
    await this.token.finishMinting({ from: owner });
    const exptectedMintingFailure = this.token.mint(owner, 10, { from: owner });

    // then
    await exptectedMintingFailure.should.eventually.be.rejectedWith(EVMError('revert'));
  });
});
