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

const MultiSig = artifacts.require('MultiSigWallet.sol');

contract('System Vars', ([_, owner, owner2, owner3, owner4, owner5, notOwner, anotherAccount]) => {
  beforeEach(async () => {
    this.multisig = await MultiSig.new([owner, owner2, owner3], 2);
  });

  it('Should add owner to the wallet', async () => {
    // when
    console.log('------------1------------');
    console.log(this.multisig);
    await this.multisig.addOwner(owner4, { from: this.multisig.web3.contract });
    console.log('------------2------------');
    // then
    console.log('------------3------------');
    const owners = await this.multisig.getOwners();
    console.log('------------4------------');
    await owners[3].should.be.equal(owner4);
  });
});

