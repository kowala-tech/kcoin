/* global artifacts, contract, it, describe, beforeEach, web3 */
/* eslint no-unused-expressions: 1 */
/* eslint consistent-return: 0 */
/* eslint-disable max-len */

const FIFSRegistrar = artifacts.require('FIFSRegistrar.sol');
const ENS = artifacts.require('ENSRegistry.sol');

const { EVMError } = require('../helpers/testUtils.js');
const namehash = require('eth-ens-namehash');

contract('FIFSRegistrar', (accounts) => {
  let registrar;
  let ens;

  beforeEach(async () => {
    ens = await ENS.new();
    registrar = await FIFSRegistrar.new(ens.address, 0);

    await ens.setOwner(0, registrar.address, { from: accounts[0] });
  });

  it('should allow registration of names', async () => {
    // when
    await registrar.register(web3.sha3('eth'), accounts[0], { from: accounts[0] });

    // then
    const ensOwner = await ens.owner(0);
    const ensSubnodeOwner = await ens.owner(namehash('eth'));

    await ensOwner.should.be.equal(registrar.address);
    await ensSubnodeOwner.should.be.equal(accounts[0]);
  });

  describe('transferring names', async () => {
    beforeEach(async () => {
      await registrar.register(web3.sha3('eth'), accounts[0], { from: accounts[0] });
    });

    it('should allow transferring name to your own', async () => {
      // when
      await registrar.register(web3.sha3('eth'), accounts[1], { from: accounts[0] });

      // then
      const ensSubnodeOwner = await ens.owner(namehash('eth'));
      await ensSubnodeOwner.should.be.equal(accounts[1]);
    });

    it('forbids transferring the name you do not own', async () => {
      // when
      const transfer = registrar.register(web3.sha3('eth'), accounts[1], { from: accounts[1] });

      // then
      await transfer.should.be.rejectedWith(EVMError('revert'));
    });
  });
});

