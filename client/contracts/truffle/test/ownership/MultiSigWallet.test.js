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

contract('MultiSigWallet', ([_, admin, owner, owner2, owner3, owner4, notOwner]) => {
  beforeEach(async () => {
    this.multisig = await MultiSig.new([owner, owner2, owner3], 2, { from: admin });
  });

  describe('Owner operations', async () => {
    it('Should add owner to the wallet', async () => {
      // when
      const addOwnerData = this.multisig.contract.addOwner.getData(owner4);
      const transactionID = await getParamFromTxEvent(
        await this.multisig.submitTransaction(this.multisig.address, 0, addOwnerData, { from: owner }),
        'transactionId',
        null,
        'Submission',
      );
      await this.multisig.confirmTransaction(transactionID, { from: owner2 });

      // then
      const owners = await this.multisig.getOwners();
      await owners[3].should.be.equal(owner4);
    });

    it('Should not add owner to the wallet via not owner when submitting transaction', async () => {
      // when
      const addOwnerData = this.multisig.contract.addOwner.getData(owner4);
      const expectedAddOwnerFail = this.multisig.submitTransaction(this.multisig.address, 0, addOwnerData, { from: notOwner });

      // then
      await expectedAddOwnerFail.should.eventually.be.rejectedWith(EVMError('revert'));
    });

    it('Should not add owner to the wallet via not owner when confirming transaction', async () => {
      // when
      const addOwnerData = this.multisig.contract.addOwner.getData(owner4);
      const transactionID = await getParamFromTxEvent(
        await this.multisig.submitTransaction(this.multisig.address, 0, addOwnerData, { from: owner }),
        'transactionId',
        null,
        'Submission',
      );
      const expectedAddOwnerFail = this.multisig.confirmTransaction(transactionID, { from: notOwner });

      // then
      await expectedAddOwnerFail.should.eventually.be.rejectedWith(EVMError('revert'));
    });

    it('Should remove owner from the wallet', async () => {
      // given
      const addOwnerData = this.multisig.contract.addOwner.getData(owner4);
      const transactionID = await getParamFromTxEvent(
        await this.multisig.submitTransaction(this.multisig.address, 0, addOwnerData, { from: owner }),
        'transactionId',
        null,
        'Submission',
      );
      await this.multisig.confirmTransaction(transactionID, { from: owner2 });

      let owners = await this.multisig.getOwners();
      await owners.length.should.be.equal(4);

      // when
      const removeOwnerData = this.multisig.contract.removeOwner.getData(owner4);
      const transactionID2 = await getParamFromTxEvent(
        await this.multisig.submitTransaction(this.multisig.address, 0, removeOwnerData, { from: owner }),
        'transactionId',
        null,
        'Submission',
      );
      await this.multisig.confirmTransaction(transactionID2, { from: owner2 });

      // then
      owners = await this.multisig.getOwners();
      await owners.length.should.be.equal(3);
    });

    it('Should not remove owner from the wallet via not owner when submitting transaction', async () => {
      // given
      const addOwnerData = this.multisig.contract.addOwner.getData(owner4);
      const transactionID = await getParamFromTxEvent(
        await this.multisig.submitTransaction(this.multisig.address, 0, addOwnerData, { from: owner }),
        'transactionId',
        null,
        'Submission',
      );
      await this.multisig.confirmTransaction(transactionID, { from: owner2 });

      const owners = await this.multisig.getOwners();
      await owners.length.should.be.equal(4);

      // when
      const removeOwnerData = this.multisig.contract.removeOwner.getData(owner4);

      const expectedRemoveOwnerFail = this.multisig.submitTransaction(this.multisig.address, 0, removeOwnerData, { from: notOwner });

      // then
      await expectedRemoveOwnerFail.should.eventually.be.rejectedWith(EVMError('revert'));
    });

    it('Should not remove owner from the wallet via not owner when confirming transaction', async () => {
      // given
      const addOwnerData = this.multisig.contract.addOwner.getData(owner4);
      const transactionID = await getParamFromTxEvent(
        await this.multisig.submitTransaction(this.multisig.address, 0, addOwnerData, { from: owner }),
        'transactionId',
        null,
        'Submission',
      );
      await this.multisig.confirmTransaction(transactionID, { from: owner2 });

      const owners = await this.multisig.getOwners();
      await owners.length.should.be.equal(4);

      // when
      const removeOwnerData = this.multisig.contract.removeOwner.getData(owner4);
      const transactionID2 = await getParamFromTxEvent(
        await this.multisig.submitTransaction(this.multisig.address, 0, removeOwnerData, { from: owner }),
        'transactionId',
        null,
        'Submission',
      );

      const expectedRemoveOwnerFail = this.multisig.confirmTransaction(transactionID2, { from: notOwner });

      // then
      await expectedRemoveOwnerFail.should.eventually.be.rejectedWith(EVMError('revert'));
    });

    it('Should replace owner to the wallet', async () => {
      // when
      const replaceOwnerData = this.multisig.contract.replaceOwner.getData(owner2, owner4);
      const transactionID = await getParamFromTxEvent(
        await this.multisig.submitTransaction(this.multisig.address, 0, replaceOwnerData, { from: owner }),
        'transactionId',
        null,
        'Submission',
      );
      await this.multisig.confirmTransaction(transactionID, { from: owner2 });

      // then
      const owners = await this.multisig.getOwners();
      await owners[1].should.be.equal(owner4);
    });

    it('Should not replace owner via not owner when submitting transaction', async () => {
      // when
      const replaceOwnerData = this.multisig.contract.replaceOwner.getData(owner2, owner4);
      const expectedReplaceOwnerFail = this.multisig.submitTransaction(this.multisig.address, 0, replaceOwnerData, { from: notOwner });

      // then
      await expectedReplaceOwnerFail.should.eventually.be.rejectedWith(EVMError('revert'));
    });

    it('Should not replace owner via not owner when confirming transaction', async () => {
      // when
      const replaceOwnerData = this.multisig.contract.replaceOwner.getData(owner2, owner4);
      const transactionID = await getParamFromTxEvent(
        await this.multisig.submitTransaction(this.multisig.address, 0, replaceOwnerData, { from: owner }),
        'transactionId',
        null,
        'Submission',
      );
      const expectedReplaceOwnerFail = this.multisig.confirmTransaction(transactionID, { from: notOwner });

      // then
      await expectedReplaceOwnerFail.should.eventually.be.rejectedWith(EVMError('revert'));
    });
  });

  describe('Requirments change', async () => {
    it('Should change requirements', async () => {
      // when
      const requirmentsData = this.multisig.contract.changeRequirement.getData(3);
      const transactionID = await getParamFromTxEvent(
        await this.multisig.submitTransaction(this.multisig.address, 0, requirmentsData, { from: owner }),
        'transactionId',
        null,
        'Submission',
      );
      await this.multisig.confirmTransaction(transactionID, { from: owner2 });

      // then
      const required = await this.multisig.required();
      await required.should.be.bignumber.equal(3);
    });

    it('Should not change requirements via not owner when submitting transaction', async () => {
      // when
      const requirmentsData = this.multisig.contract.changeRequirement.getData(3);
      const expectedChangeRequirementsFail = this.multisig.submitTransaction(this.multisig.address, 0, requirmentsData, { from: notOwner });

      // then
      await expectedChangeRequirementsFail.should.eventually.be.rejectedWith(EVMError('revert'));
    });

    it('Should not change requirements via not owner when submitting transaction', async () => {
      // when
      const requirmentsData = this.multisig.contract.changeRequirement.getData(3);
      const transactionID = await getParamFromTxEvent(
        await this.multisig.submitTransaction(this.multisig.address, 0, requirmentsData, { from: owner }),
        'transactionId',
        null,
        'Submission',
      );
      const expectedChangeRequirementsFail = this.multisig.confirmTransaction(transactionID, { from: notOwner });

      // then
      await expectedChangeRequirementsFail.should.eventually.be.rejectedWith(EVMError('revert'));
    });
  });

  describe('Transactions operations', async () => {
    it('Should not confirm already confirmed transaction', async () => {
      // when
      const requirmentsData = this.multisig.contract.changeRequirement.getData(3);
      const transactionID = await getParamFromTxEvent(
        await this.multisig.submitTransaction(this.multisig.address, 0, requirmentsData, { from: owner }),
        'transactionId',
        null,
        'Submission',
      );
      await this.multisig.confirmTransaction(transactionID, { from: owner2 });
      const expectedConfirmationFailure = this.multisig.confirmTransaction(transactionID, { from: owner3 });

      // then
      await expectedConfirmationFailure.should.eventually.be.rejectedWith(EVMError('revert'));
    });

    it('Should expected 3 confirmation after change in requirments', async () => {
      // given
      const requirmentsData = this.multisig.contract.changeRequirement.getData(3);
      const transactionID = await getParamFromTxEvent(
        await this.multisig.submitTransaction(this.multisig.address, 0, requirmentsData, { from: owner }),
        'transactionId',
        null,
        'Submission',
      );
      await this.multisig.confirmTransaction(transactionID, { from: owner2 });

      const required = await this.multisig.required();
      await required.should.be.bignumber.equal(3);

      // when
      const addOwnerData = this.multisig.contract.addOwner.getData(owner4);
      const transactionID2 = await getParamFromTxEvent(
        await this.multisig.submitTransaction(this.multisig.address, 0, addOwnerData, { from: owner }),
        'transactionId',
        null,
        'Submission',
      );
      await this.multisig.confirmTransaction(transactionID2, { from: owner2 });
      await this.multisig.confirmTransaction(transactionID2, { from: owner3 });

      // then
      const owners = await this.multisig.getOwners();
      await owners[3].should.be.equal(owner4);
    });

    it('Should revoke confirmation', async () => {
      // given
      const requirmentsData = this.multisig.contract.changeRequirement.getData(3);
      const transactionID = await getParamFromTxEvent(
        await this.multisig.submitTransaction(this.multisig.address, 0, requirmentsData, { from: owner }),
        'transactionId',
        null,
        'Submission',
      );

      // when
      await this.multisig.revokeConfirmation(transactionID, { from: owner });

      // then
      const confirmations = await this.multisig.getConfirmationCount(transactionID);
      await confirmations.should.be.bignumber.equal(0);
    });

    it('Should not execute unconfirmed transaction', async () => {
      // given
      const requirmentsData = this.multisig.contract.changeRequirement.getData(3);
      const transactionID = await getParamFromTxEvent(
        await this.multisig.submitTransaction(this.multisig.address, 0, requirmentsData, { from: owner }),
        'transactionId',
        null,
        'Submission',
      );
      await this.multisig.confirmTransaction(transactionID, { from: owner2 });

      const required = await this.multisig.required();
      await required.should.be.bignumber.equal(3);

      // when
      const addOwnerData = this.multisig.contract.addOwner.getData(owner4);
      const transactionID2 = await getParamFromTxEvent(
        await this.multisig.submitTransaction(this.multisig.address, 0, addOwnerData, { from: owner }),
        'transactionId',
        null,
        'Submission',
      );
      await this.multisig.confirmTransaction(transactionID2, { from: owner2 });

      const expectedExecutionFailure = this.multisig.executeTransaction(transactionID2, { from: owner3 });

      // then
      await expectedExecutionFailure.should.eventually.be.rejectedWith(EVMError('revert'));
    });

    it('Should not execute transaction via not owner', async () => {
      // given
      const requirmentsData = this.multisig.contract.changeRequirement.getData(3);
      const transactionID = await getParamFromTxEvent(
        await this.multisig.submitTransaction(this.multisig.address, 0, requirmentsData, { from: owner }),
        'transactionId',
        null,
        'Submission',
      );
      await this.multisig.confirmTransaction(transactionID, { from: owner2 });

      const required = await this.multisig.required();
      await required.should.be.bignumber.equal(3);

      // when
      const addOwnerData = this.multisig.contract.addOwner.getData(owner4);
      const transactionID2 = await getParamFromTxEvent(
        await this.multisig.submitTransaction(this.multisig.address, 0, addOwnerData, { from: owner }),
        'transactionId',
        null,
        'Submission',
      );
      await this.multisig.confirmTransaction(transactionID2, { from: owner2 });

      const expectedExecutionFailure = this.multisig.executeTransaction(transactionID2, { from: notOwner });

      // then
      await expectedExecutionFailure.should.eventually.be.rejectedWith(EVMError('revert'));
    });
  });
});
