/* global artifacts, contract, it, beforeEach, web3 */
/* eslint no-unused-expressions: 0 */
/* eslint consistent-return: 0 */
/* eslint-disable max-len */

const ENS = artifacts.require('ENSRegistry.sol');
const FIFSRegistrar = artifacts.require('FIFSRegistrar.sol');
const PublicResolver = artifacts.require('PublicResolver.sol');
const ValidatorMgr = artifacts.require('ValidatorMgr.sol');

const namehash = require('eth-ens-namehash');

contract('ENS Functionality', (accounts) => {
  beforeEach(async () => {
    this.ens = await ENS.new();
    this.registrar = await FIFSRegistrar.new(this.ens.address, namehash('eth'));
    this.resolver = await PublicResolver.new(this.ens.address);
    await this.ens.setSubnodeOwner(0, web3.sha3('eth'), this.registrar.address, { from: accounts[0] });
  });

  it('should register a validator domain', async () => {
    // when
    await this.registrar.register(web3.sha3('validator'), accounts[0], { from: accounts[0] });
    // then
    const ensSubnodeOwner = await this.ens.owner(namehash('validator.eth'));
    await ensSubnodeOwner.should.be.equal(accounts[0]);
  });

  it('should add resolver to validator.eth domain', async () => {
    // given
    await this.registrar.register(web3.sha3('validator'), accounts[0], { from: accounts[0] });

    // when
    await this.ens.setResolver(namehash('validator.eth'), this.resolver.address);
    const resolver = await this.ens.resolver(namehash('validator.eth'));

    // then
    await resolver.should.be.equal(this.resolver.address);
  });

  it('should add validator address to validator.eth domain', async () => {
    // given
    const validator = await ValidatorMgr.new(1, 2, 3, '0x1234');
    await this.registrar.register(web3.sha3('validator'), accounts[0], { from: accounts[0] });
    await this.ens.setResolver(namehash('validator.eth'), this.resolver.address);

    // when
    await this.resolver.setAddr(namehash('validator.eth'), validator.address);

    // then
    const validatorEnsAddr = await this.resolver.addr(namehash('validator.eth'));
    await validatorEnsAddr.should.be.equal(validator.address);
  });
});
