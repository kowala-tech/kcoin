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
  EVMError, ether, getParamFromTxEvent,
} = require('../helpers/testUtils.js');

const MultiSig = artifacts.require('MultiSigWallet.sol');

contract('System Vars', ([_, owner, owner2, owner3, owner4, owner5, notOwner, anotherAccount]) => {
  beforeEach(async () => {
    this.multisig = await MultiSig.new([owner, owner2, owner3], 2);
  });

  it('Should add owner to the wallet', async () => {
    // when
    const addOwnerData = this.multisig.contract.addOwner.getData(owner4);
    const transactionID = await getParamFromTxEvent(
      await this.multisig.submitTransaction(this.multisig.address, 0, addOwnerData, { from: owner }),
      'transactionId',
      null,
      'Submission',
    );
    await this.multisig.confirmTransaction(transactionID.toNumber(), { from: owner2 });

    // then
    const owners = await this.multisig.getOwners();
    await owners[3].should.be.equal(owner4);
  });
});
