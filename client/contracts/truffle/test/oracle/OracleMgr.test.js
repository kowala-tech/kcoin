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
const PublicResolver = artifacts.require('PublicResolver.sol');
const KNS = artifacts.require('KNSRegistry.sol');
const FIFSRegistrar = artifacts.require('FIFSRegistrar.sol');
const namehash = require('eth-ens-namehash');
const { EVMError } = require('../helpers/testUtils.js');

contract('Oracle Manager', ([_, owner, newOwner, notOwner]) => {
  beforeEach(async () => {
    this.validator = await ValidatorMgr.new(1, 2, 3, '0x1234', 1);
    this.kns = await KNS.new({ from: owner });
    this.registrar = await FIFSRegistrar.new(this.kns.address, namehash('kowala'));
    this.resolver = await PublicResolver.new(this.kns.address);
    await this.kns.setSubnodeOwner(0, web3.sha3('kowala'), this.registrar.address, { from: owner });
    await this.registrar.register(web3.sha3('validator'), owner, { from: owner });
    await this.kns.setResolver(namehash('validator.kowala'), this.resolver.address, { from: owner });
    await this.resolver.setAddr(namehash('validator.kowala'), this.validator.address, { from: owner });
    this.oracle = await OracleMgr.new(1, 1, 1, 1, 1, 1, this.validator.address, this.kns.address, { from: owner });
  });

  it('should set ValidatorMgr address using KNS', async () => {
    // given
    const knsResolverAddr = await this.oracle.knsResolver();
    const resolver = await PublicResolver.at(knsResolverAddr);

    // when
    const validatorAddress = await this.oracle.getValidatorAddress({ from: owner });

    // then
    const validatorAddrFromResolver = await resolver.addr(namehash('validator.kowala'));
    await validatorAddress.should.be.equal(validatorAddrFromResolver);
  });
});
