/* global artifacts, web3 */
/* eslint-disable max-len */

const NS = artifacts.require('./ens/NSRegistry.sol');
const FIFSRegistrar = artifacts.require('./ens/FIFSRegistrar.sol');
const PublicResolver = artifacts.require('./ens/PublicResolver.sol');
const namehash = require('../node_modules/eth-ens-namehash');

module.exports = (deployer) => {
  const domain = 'kowala';
  const tld = 'test';

  const rootNode = {
    namehash: namehash(tld),
    sha3: web3.sha3(tld),
  };


  deployer.deploy(NS)
    .then(() => deployer.deploy(FIFSRegistrar, NS.address, rootNode.namehash))
    .then(() => NS.at(NS.address).setSubnodeOwner('0x0', rootNode.sha3, FIFSRegistrar.address))
    .then(() => deployer.deploy(PublicResolver, NS.address))
    .then(() => FIFSRegistrar.at(FIFSRegistrar.address).register(web3.sha3(domain), web3.eth.accounts[0]))
    .then(() => NS.at(NS.address).setResolver(namehash(`${domain}.${tld}`), PublicResolver.address));
};
