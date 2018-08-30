/* global artifacts, contract, it, beforeEach, web3 */
/* eslint no-unused-expressions: 0 */
/* eslint consistent-return: 0 */
/* eslint-disable max-len */

require('chai')
  .use(require('chai-as-promised'))
  .should();

const KNS = artifacts.require('KNSRegistry.sol');
const FIFSRegistrar = artifacts.require('FIFSRegistrar.sol');
const PublicResolver = artifacts.require('PublicResolver.sol');
const ValidatorMgr = artifacts.require('ValidatorMgr.sol');
const namehash = require('eth-ens-namehash');

contract('KNS Functionality', (accounts) => {
  const domain = 'kowala';
  const tld = 'validator';
  const domaintld = 'validator.kowala';

  beforeEach(async () => {
    this.kns = await KNS.new();
    this.registrar = await FIFSRegistrar.new(this.kns.address, namehash(domain));
    this.resolver = await PublicResolver.new(this.kns.address);
    await this.kns.setSubnodeOwner(0, web3.sha3(domain), this.registrar.address, { from: accounts[0] });
  });

  it('should register a domain', async () => {
    // when
    await this.registrar.register(web3.sha3(tld), accounts[0], { from: accounts[1] });
    // then
    const knsSubnodeOwner = await this.kns.owner(namehash(domaintld));
    await knsSubnodeOwner.should.be.equal(accounts[0]);
  });

  it('should add resolver to validator.kowala domain', async () => {
    // given
    await this.registrar.register(web3.sha3(tld), accounts[0], { from: accounts[1] });

    // when
    await this.kns.setResolver(namehash(domaintld), this.resolver.address);
    const resolver = await this.kns.resolver(namehash(domaintld));

    // then
    await resolver.should.be.equal(this.resolver.address);
  });

  it('should add validator address to validator.kowala domain', async () => {
    // given
    const validator = await ValidatorMgr.new(1, 2, 3, 1, this.resolver.address);
    await this.registrar.register(web3.sha3(tld), accounts[0], { from: accounts[1] });
    await this.kns.setResolver(namehash(domaintld), this.resolver.address);

    // when
    await this.resolver.setAddr(namehash(domaintld), validator.address, { from: accounts[0] });

    // then
    const validatorEnsAddr = await this.resolver.addr(namehash(domaintld));
    await validatorEnsAddr.should.be.equal(validator.address);
  });
});
