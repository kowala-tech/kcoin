/* global artifacts, contract, it, describe, beforeEach, web3, assert */
/* eslint no-unused-expressions: 1 */
/* eslint consistent-return: 0 */
/* eslint-disable max-len */


const KNS = artifacts.require('KNSRegistry.sol');
const PublicResolver = artifacts.require('PublicResolver.sol');

const { EVMError } = require('../helpers/testUtils.js');
const utils = require('../helpers/Utils.js');
const namehash = require('eth-ens-namehash');

contract('PublicResolver', (accounts) => {
  let node;
  let kns;
  let resolver;

  beforeEach(async () => {
    node = namehash('eth');
    kns = await KNS.new();
    resolver = await PublicResolver.new(kns.address);
    await kns.setSubnodeOwner(0, web3.sha3('eth'), accounts[0], { from: accounts[0] });
  });

  describe('fallback function', async () => {
    it('forbids calls to the fallback function with 0 value', async () => {
      try {
        await web3.eth.sendTransaction({
          from: accounts[0],
          to: resolver.address,
          gas: 3000000,
        });
      } catch (error) {
        return utils.ensureException(error);
      }

      assert.fail('transfer did not fail');
    });

    it('forbids calls to the fallback function with 1 value', async () => {
      try {
        await web3.eth.sendTransaction({
          from: accounts[0],
          to: resolver.address,
          gas: 3000000,
          value: 1,
        });
      } catch (error) {
        return utils.ensureException(error);
      }

      assert.fail('transfer did not fail');
    });
  });

  describe('supportsInterface function', async () => {
    it('supports known interfaces', async () => {
      // then
      await resolver.supportsInterface('0x3b3b57de').should.eventually.be.true;
      await resolver.supportsInterface('0x691f3431').should.eventually.be.true;
      await resolver.supportsInterface('0x2203ab56').should.eventually.be.true;
      await resolver.supportsInterface('0xc8690233').should.eventually.be.true;
      await resolver.supportsInterface('0x59d1d43c').should.eventually.be.true;
    });

    it('does not support a random interface', async () => {
      // then
      await resolver.supportsInterface('0x3b3b57df').should.eventually.be.false;
    });
  });


  describe('addr', async () => {
    it('permits setting address by owner', async () => {
      // when
      await resolver.setAddr(node, accounts[1], { from: accounts[0] });
      const nodeAddr = resolver.addr(node);

      // then
      await nodeAddr.should.eventually.be.equal(accounts[1]);
    });

    it('can overwrite previously set address', async () => {
      // given
      await resolver.setAddr(node, accounts[1], { from: accounts[0] });

      // when
      await resolver.setAddr(node, accounts[0], { from: accounts[0] });

      // then
      const nodeAddr = resolver.addr(node);
      await nodeAddr.should.eventually.be.equal(accounts[0]);
    });

    it('can overwrite to same address', async () => {
      // given
      await resolver.setAddr(node, accounts[1], { from: accounts[0] });

      // when
      await resolver.setAddr(node, accounts[1], { from: accounts[0] });

      // then
      const nodeAddr = resolver.addr(node);
      await nodeAddr.should.eventually.be.equal(accounts[1]);
    });

    it('forbids setting new address by non-owners', async () => {
      // then
      const setAddr = resolver.setAddr(node, accounts[1], { from: accounts[1] });
      await setAddr.should.be.eventually.rejectedWith(EVMError('revert'));
    });

    it('forbids writing same address by non-owners', async () => {
      // given
      await resolver.setAddr(node, accounts[1], { from: accounts[0] });

      // when
      const setAddr = resolver.setAddr(node, accounts[1], { from: accounts[1] });

      // then
      await setAddr.should.be.eventually.rejectedWith(EVMError('revert'));
    });

    it('forbids overwriting existing address by non-owners', async () => {
      // given
      await resolver.setAddr(node, accounts[1], { from: accounts[0] });

      // when
      const setAddr = resolver.setAddr(node, accounts[0], { from: accounts[1] });

      // then
      await setAddr.should.be.eventually.rejectedWith(EVMError('revert'));
    });

    it('returns zero when fetching nonexistent addresses', async () => {
      // then
      await resolver.addr(node).should.eventually.be.equal('0x0000000000000000000000000000000000000000');
    });
  });

  describe('content', async () => {
    it('permits setting content by owner', async () => {
      // when
      await resolver.setContent(node, 'hash1', { from: accounts[0] });

      // then
      const nodeContent = web3.toUtf8(await resolver.content(node));
      await nodeContent.should.be.equal('hash1');
    });

    it('can overwrite previously set content', async () => {
      // given
      await resolver.setContent(node, 'hash1', { from: accounts[0] });

      // when
      await resolver.setContent(node, 'hash2', { from: accounts[0] });

      // then
      const nodeContent = web3.toUtf8(await resolver.content(node));
      await nodeContent.should.be.equal('hash2');
    });

    it('can overwrite to same content', async () => {
      // given
      await resolver.setContent(node, 'hash1', { from: accounts[0] });

      // when
      await resolver.setContent(node, 'hash1', { from: accounts[0] });

      // then
      const nodeContent = web3.toUtf8(await resolver.content(node));
      await nodeContent.should.be.equal('hash1');
    });

    it('forbids setting content by non-owners', async () => {
      // when
      const nodeContent = resolver.setContent(node, 'hash1', { from: accounts[1] });

      // then
      await nodeContent.should.eventually.be.rejectedWith(EVMError('revert'));
    });

    it('forbids writing same content by non-owners', async () => {
      // given
      await resolver.setContent(node, 'hash1', { from: accounts[0] });

      // when
      const nodeContent = resolver.setContent(node, 'hash1', { from: accounts[1] });

      // then
      await nodeContent.should.eventually.be.rejectedWith(EVMError('revert'));
    });

    it('returns empty when fetching nonexistent content', async () => {
      // when
      const nodeContent = await resolver.content(node);

      // then
      await nodeContent.should.be.equal('0x0000000000000000000000000000000000000000000000000000000000000000');
    });
  });

  describe('name', async () => {
    it('permits setting name by owner', async () => {
      // when
      await resolver.setName(node, 'name1', { from: accounts[0] });

      // then
      const name = resolver.name(node);
      await name.should.eventually.be.equal('name1');
    });

    it('can overwrite previously set names', async () => {
      // given
      await resolver.setName(node, 'name1', { from: accounts[0] });

      // when
      await resolver.setName(node, 'name2', { from: accounts[0] });

      // then
      const name = resolver.name(node);
      await name.should.eventually.be.equal('name2');
    });

    it('forbids setting name by non-owners', async () => {
      // then
      await resolver.setName(node, 'name2', { from: accounts[1] })
        .should.eventually.be.rejectedWith(EVMError('revert'));
    });

    it('returns empty when fetching nonexistent name', async () => {
      // when
      const nodeName = await resolver.name(node);

      // then
      await nodeName.should.be.equal('');
    });
  });

  describe('pubkey', async () => {
    it('returns empty when fetching nonexistent values', async () => {
      // given
      const nonexistentValues = [
        '0x0000000000000000000000000000000000000000000000000000000000000000',
        '0x0000000000000000000000000000000000000000000000000000000000000000'];

      // when
      const pubkey = resolver.pubkey(node);

      // then
      await pubkey.should.eventually.be.deep.equal(nonexistentValues);
    });

    it('permits setting public key by owner', async () => {
      // given
      const keys = [
        '0x1000000000000000000000000000000000000000000000000000000000000000',
        '0x2000000000000000000000000000000000000000000000000000000000000000'];
      await resolver.setPubkey(node, 1, 2, { from: accounts[0] });

      // when
      const pubkey = resolver.pubkey(node);

      // then
      await pubkey.should.eventually.be.deep.equal(keys);
    });

    it('can overwrite previously set value', async () => {
      // given
      const keys = [
        '0x3000000000000000000000000000000000000000000000000000000000000000',
        '0x4000000000000000000000000000000000000000000000000000000000000000'];
      await resolver.setPubkey(node, 1, 2, { from: accounts[0] });

      // when
      await resolver.setPubkey(node, 3, 4, { from: accounts[0] });
      const pubkey = resolver.pubkey(node);

      // then
      await pubkey.should.eventually.be.deep.equal(keys);
    });

    it('can overwrite to same value', async () => {
      // given
      const keys = [
        '0x1000000000000000000000000000000000000000000000000000000000000000',
        '0x2000000000000000000000000000000000000000000000000000000000000000'];
      await resolver.setPubkey(node, 1, 2, { from: accounts[0] });

      // when
      await resolver.setPubkey(node, 1, 2, { from: accounts[0] });
      const pubkey = resolver.pubkey(node);

      // then
      await pubkey.should.eventually.be.deep.equal(keys);
    });

    it('forbids setting value by non-owners', async () => {
      // then
      await resolver.setPubkey(node, 1, 2, { from: accounts[1] })
        .should.eventually.be.rejectedWith(EVMError('revert'));
    });

    it('forbids writing same value by non-owners', async () => {
      // when
      await resolver.setPubkey(node, 1, 2, { from: accounts[0] });

      // then
      await resolver.setPubkey(node, 1, 2, { from: accounts[1] })
        .should.eventually.be.rejectedWith(EVMError('revert'));
    });

    it('forbids overwriting existing value by non-owners', async () => {
      // when
      await resolver.setPubkey(node, 1, 2, { from: accounts[0] });

      // then
      await resolver.setPubkey(node, 3, 4, { from: accounts[1] })
        .should.eventually.be.rejectedWith(EVMError('revert'));
    });
  });

  describe('ABI', async () => {
    it('returns a contentType of 0 when nothing is available', async () => {
      // given
      const result = await resolver.ABI(node, 0xFFFFFFFF);

      // when
      const ABI = [result[0].toNumber(), result[1]];

      // then
      await ABI.should.be.deep.equal([0, '0x']);
    });

    it('returns an ABI after it has been set', async () => {
      // given
      await resolver.setABI(node, 0x1, 'foo', { from: accounts[0] });

      // when
      const result = await resolver.ABI(node, 0xFFFFFFFF);
      const ABI = [result[0].toNumber(), result[1]];

      // then
      await ABI.should.be.deep.equal([1, '0x666f6f']);
    });

    it('returns the first valid ABI', async () => {
      // given
      await resolver.setABI(node, 0x2, 'foo', { from: accounts[0] });
      await resolver.setABI(node, 0x4, 'bar', { from: accounts[0] });

      // when
      const result1 = await resolver.ABI(node, 0x7);
      const result2 = await resolver.ABI(node, 0x5);
      const ABI1 = [result1[0].toNumber(), result1[1]];
      const ABI2 = [result2[0].toNumber(), result2[1]];

      // then
      await ABI1.should.be.deep.equal([2, '0x666f6f']);
      await ABI2.should.be.deep.equal([4, '0x626172']);
    });

    it('allows deleting ABIs', async () => {
      // given
      await resolver.setABI(node, 0x1, 'foo', { from: accounts[0] });
      const result1 = await resolver.ABI(node, 0xFFFFFFFF);
      const beforeABI = [result1[0].toNumber(), result1[1]];

      // when
      await resolver.setABI(node, 0x1, '', { from: accounts[0] });
      const result2 = await resolver.ABI(node, 0xFFFFFFFF);
      const afterABI = [result2[0].toNumber(), result2[1]];

      // then
      await beforeABI.should.be.deep.equal([1, '0x666f6f']);
      await afterABI.should.be.deep.equal([0, '0x']);
    });

    it('rejects invalid content types', async () => {
      await resolver.setABI(node, 0x3, 'foo', { from: accounts[0] })
        .should.be.rejectedWith(EVMError('revert'));
    });

    it('forbids setting value by non-owners', async () => {
      await resolver.setABI(node, 0x1, 'foo', { from: accounts[1] })
        .should.be.rejectedWith(EVMError('revert'));
    });
  });

  describe('text', async () => {
    const url = 'https://ethereum.org';
    const url2 = 'https://github.com/ethereum';

    it('permits setting text by owner', async () => {
      // when
      await resolver.setText(node, 'url', url, { from: accounts[0] });
      const text = resolver.text(node, 'url');

      // then
      await text.should.eventually.be.equal(url);
    });

    it('can overwrite previously set text', async () => {
      // given
      await resolver.setText(node, 'url', url, { from: accounts[0] });

      // when
      await resolver.setText(node, 'url', url2, { from: accounts[0] });
      const text = resolver.text(node, 'url');

      // then
      await text.should.eventually.be.equal(url2);
    });

    it('can overwrite to same text', async () => {
      // given
      await resolver.setText(node, 'url', url, { from: accounts[0] });

      // when
      await resolver.setText(node, 'url', url, { from: accounts[0] });
      const text = resolver.text(node, 'url');

      // then
      await text.should.eventually.be.equal(url);
    });

    it('forbids setting new text by non-owners', async () => {
      // then
      await resolver.setText(node, 'url', url, { from: accounts[1] })
        .should.be.rejectedWith(EVMError('revert'));
    });

    it('forbids writing same text by non-owners', async () => {
      // when
      await resolver.setText(node, 'url', url, { from: accounts[0] });

      // then
      await resolver.setText(node, 'url', url, { from: accounts[1] })
        .should.be.rejectedWith(EVMError('revert'));
    });
  });

  describe('multihash', async () => {
    const hash1 = '0x0000000000000000000000000000000000000000000000000000000000000001';
    const hash2 = '0x0000000000000000000000000000000000000000000000000000000000000002';

    it('permits setting multihash by owner', async () => {
      // when
      await resolver.setMultihash(node, hash1, { from: accounts[0] });
      const hashResult = resolver.multihash(node);

      // then
      await hashResult.should.eventually.be.equal(hash1);
    });

    it('can overwrite previously set multihash', async () => {
      // given
      await resolver.setMultihash(node, hash1, { from: accounts[0] });

      // when
      await resolver.setMultihash(node, hash2, { from: accounts[0] });
      const hashResult = resolver.multihash(node);

      // then
      await hashResult.should.eventually.be.equal(hash2);
    });

    it('can overwrite to same multihash', async () => {
      // given
      await resolver.setMultihash(node, hash1, { from: accounts[0] });

      // when
      await resolver.setMultihash(node, hash1, { from: accounts[0] });
      const hashResult = resolver.multihash(node);

      // then
      await hashResult.should.eventually.be.equal(hash1);
    });

    it('forbids setting multihash by non-owners', async () => {
      // then
      await resolver.setMultihash(node, hash1, { from: accounts[1] })
        .should.be.rejectedWith(EVMError('revert'));
    });

    it('forbids writing same multihash by non-owners', async () => {
      // given
      await resolver.setMultihash(node, hash1, { from: accounts[0] });

      // when
      const hashResult = resolver.setMultihash(node, hash1, { from: accounts[1] });

      // then
      await hashResult.should.eventually.be.rejectedWith(EVMError('revert'));
    });

    it('returns empty when fetching nonexistent multihash', async () => {
      // when
      const hashResult = await resolver.multihash(node);

      // then
      await hashResult.should.be.equal('0x');
    });
  });
});
