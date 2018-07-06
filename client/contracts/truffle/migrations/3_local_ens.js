/* global artifacts, web3 */

const ENS = artifacts.require("./ENSRegistry.sol");
const FIFSRegistrar = artifacts.require('./FIFSRegistrar.sol');
const PublicResolver = artifacts.require('./PublicResolver.sol');
const namehash = require('../node_modules/eth-ens-namehash');

module.exports = function(deployer) {
  var domain = 'kowala';
  var tld = 'test';

  var rootNode = {
    namehash: namehash(tld),
    sha3: web3.sha3(tld)
  };

  deployer.deploy(ENS)
    .then(function() {
      return deployer.deploy(FIFSRegistrar, ENS.address, rootNode.namehash);
    })
    .then(function() {
      return ENS.at(ENS.address).setSubnodeOwner('0x0', rootNode.sha3, FIFSRegistrar.address);
    })
    .then(function() {
      return deployer.deploy(PublicResolver, ENS.address);
    })
    .then(function() {
      return FIFSRegistrar.at(FIFSRegistrar.address).register(web3.sha3(domain), web3.eth.accounts[0]);
    })
    .then(function() {
      return ENS.at(ENS.address).setResolver(namehash(domain + '.' + tld), PublicResolver.address);
    });
};