/* global artifacts, contract, it, beforeEach, web3 */
/* eslint no-unused-expressions: 0 */
/* eslint consistent-return: 0 */
/* eslint-disable max-len */

const ENS = artifacts.require('ENSRegistry.sol');

const { EVMError } = require('../helpers/testUtils.js');
const namehash = require('eth-ens-namehash');

contract('ENS', (accounts) => {
  let ens;
  const owner = '0x0000000000000000000000000000000000001234';
  const node = '0x0000000000000000000000000000000000000000000000000000000000000000';

  beforeEach(async () => {
    ens = await ENS.new();
  });

  it('should allow ownership transfers', async () => {
    // when
    const result = await ens.setOwner(0, '0x1234', { from: accounts[0] });
    const { args } = result.logs[0];
    const ensOwner = await ens.owner(0);

    // then
    await ensOwner.should.be.equal(owner);
    await result.logs.length.should.be.equal(1);
    await args.node.should.be.equal(node);
    await args.owner.should.be.equal(owner);
  });

  it('should prohibit transfers by non-owners', async () => {
    // when
    const transferByNonOwner = ens.setOwner(1, '0x1234', { from: accounts[0] });

    // then
    await transferByNonOwner.should.be.rejectedWith(EVMError('revert'));
  });

  it('should allow setting resolvers', async () => {
    // when
    const result = await ens.setResolver(0, '0x1234', { from: accounts[0] });
    const { args } = result.logs[0];
    const resolver = await ens.resolver(0);

    // then
    await resolver.should.be.equal(owner);
    await result.logs.length.should.be.equal(1);
    await args.node.should.be.equal(node);
    await args.resolver.should.be.equal(owner);
  });

  it('should prevent setting resolvers by non-owners', async () => {
    // then
    const resolverByNonOwner = ens.setResolver(1, '0x1234', { from: accounts[0] });

    // then
    await resolverByNonOwner.should.be.rejectedWith(EVMError('revert'));
  });

  it('should allow setting the TTL', async () => {
    // when
    const result = await ens.setTTL(0, 3600, { from: accounts[0] });
    const { args } = result.logs[0];
    const ttl = await ens.ttl(0);

    // then
    await ttl.should.be.bignumber.equal(3600);
    await result.logs.length.should.be.equal(1);
    await args.node.should.be.equal(node);
    await args.ttl.should.be.bignumber.equal(3600);
  });

  it('should prevent setting the TTL by non-owners', async () => {
    // when
    const ttlByNonOwner = ens.setTTL(1, 3600, { from: accounts[0] });

    // then
    await ttlByNonOwner.should.be.rejectedWith(EVMError('revert'));
  });

  it('should allow the creation of subnodes', async () => {
    // when
    const result = await ens.setSubnodeOwner(0, web3.sha3('eth'), accounts[1], { from: accounts[0] });
    const { args } = result.logs[0];
    const ensSubnodeOwner = await ens.owner(namehash('eth'));

    // then
    await ensSubnodeOwner.should.be.equal(accounts[1]);
    await result.logs.length.should.be.equal(1);
    await args.node.should.be.equal(node);
    await args.label.should.be.equal(web3.sha3('eth'));
    await args.owner.should.be.equal(accounts[1]);
  });

  it('should prohibit subnode creation by non-owners', async () => {
    // when
    const subnodeByNonOwner = ens.setSubnodeOwner(0, web3.sha3('eth'), accounts[1], { from: accounts[1] });

    // then
    subnodeByNonOwner.should.be.rejectedWith(EVMError('revert'));
  });
});
