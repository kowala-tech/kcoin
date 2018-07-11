/* global artifacts, web3 */
/* eslint-disable max-len */

const ENS = artifacts.require('./ens/ENSRegistry.sol');
const FIFSRegistrar = artifacts.require('./ens/FIFSRegistrar.sol');
const PublicResolver = artifacts.require('./ens/PublicResolver.sol');
const namehash = require('eth-ens-namehash');

module.exports = (deployer) => {
  const domain = 'kowala';
  const tld = 'test';

  const rootNode = {
    namehash: namehash(tld),
    sha3: web3.sha3(tld),
  };


  deployer.deploy(ENS)
    .then(() => deployer.deploy(FIFSRegistrar, ENS.address, rootNode.namehash))
    .then(() => ENS.at(ENS.address).setSubnodeOwner('0x0', rootNode.sha3, FIFSRegistrar.address))
    .then(() => deployer.deploy(PublicResolver, ENS.address))
    .then(() => FIFSRegistrar.at(FIFSRegistrar.address).register(web3.sha3(domain), web3.eth.accounts[0]))
    .then(() => ENS.at(ENS.address).setResolver(namehash(`${domain}.${tld}`), PublicResolver.address));
};
