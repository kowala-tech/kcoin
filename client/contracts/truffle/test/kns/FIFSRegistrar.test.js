/* global artifacts, contract, it, describe, beforeEach, web3 */
/* eslint no-unused-expressions: 1 */
/* eslint consistent-return: 0 */
/* eslint-disable max-len */

const FIFSRegistrar = artifacts.require('FIFSRegistrar.sol');
const KNS = artifacts.require('KNSRegistry.sol');

const { EVMError } = require('../helpers/testUtils.js');
const namehash = require('eth-ens-namehash');

contract('FIFSRegistrar', (accounts) => {
  let registrar;
  let kns;

  beforeEach(async () => {
    kns = await KNS.new();
    registrar = await FIFSRegistrar.new(kns.address, 0);

    await kns.setOwner(0, registrar.address, { from: accounts[0] });
  });

  it('should allow registration of names', async () => {
    // when
    await registrar.register(web3.sha3('eth'), accounts[0], { from: accounts[0] });

    // then
    const nsOwner = await kns.owner(0);
    const nsSubnodeOwner = await kns.owner(namehash('eth'));

    await nsOwner.should.be.equal(registrar.address);
    await nsSubnodeOwner.should.be.equal(accounts[0]);
  });

  describe('transferring names', async () => {
    beforeEach(async () => {
      await registrar.register(web3.sha3('eth'), accounts[0], { from: accounts[0] });
    });

    it('should allow transferring name to your own', async () => {
      // when
      await registrar.register(web3.sha3('eth'), accounts[1], { from: accounts[0] });

      // then
      const nsSubnodeOwner = await kns.owner(namehash('eth'));
      await nsSubnodeOwner.should.be.equal(accounts[1]);
    });

    it('forbids transferring the name you do not own', async () => {
      // when
      const transfer = registrar.register(web3.sha3('eth'), accounts[1], { from: accounts[1] });

      // then
      await transfer.should.be.rejectedWith(EVMError('revert'));
    });
  });
});

