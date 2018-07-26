/* global artifacts, web3 */
/* eslint-disable max-len */

const KNS = artifacts.require('./ens/KNSRegistry.sol');
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


  deployer.deploy(KNS)
    .then(() => deployer.deploy(FIFSRegistrar, KNS.address, rootNode.namehash))
    .then(() => KNS.at(KNS.address).setSubnodeOwner('0x0', rootNode.sha3, FIFSRegistrar.address))
    .then(() => deployer.deploy(PublicResolver, KNS.address))
    .then(() => FIFSRegistrar.at(FIFSRegistrar.address).register(web3.sha3(domain), web3.eth.accounts[0]))
    .then(() => KNS.at(KNS.address).setResolver(namehash(`${domain}.${tld}`), PublicResolver.address));
};
