/* global artifacts, web3 */
/* eslint-disable max-len */

const MyContractV0 = artifacts.require('MyContractV0.sol');
const MyContractV1 = artifacts.require('MyContractV1.sol');

const KNS = artifacts.require('./kns/KNSRegistry.sol');
const FIFSRegistrar = artifacts.require('./kns/FIFSRegistrar.sol');
const PublicResolver = artifacts.require('./kns/PublicResolver.sol');
const NameHash = artifacts.require('./utils/NameHash.sol');
const OracleMgr = artifacts.require('OracleMgr.sol');
const ValidatorMgr = artifacts.require('ValidatorMgr.sol');
const KNSRegistryV1 = artifacts.require('KNSRegistryV1.sol');
const namehash = require('../node_modules/eth-ens-namehash');

module.exports = (deployer) => {
  const domain = 'kowala';
  const tld = 'test';
  const rootNode = {
    namehash: namehash(tld),
    sha3: web3.sha3(tld),
  };

  deployer.deploy(NameHash);
  deployer.link(NameHash, OracleMgr);
  deployer.link(NameHash, ValidatorMgr);

  deployer.deploy(KNS)
    .then(() => deployer.deploy(FIFSRegistrar, KNS.address, rootNode.namehash))
    .then(() => KNS.at(KNS.address).setSubnodeOwner('0x0', rootNode.sha3, FIFSRegistrar.address))
    .then(() => deployer.deploy(PublicResolver, KNS.address))
    .then(() => FIFSRegistrar.at(FIFSRegistrar.address).register(web3.sha3(domain), web3.eth.accounts[0]))
    .then(() => KNS.at(KNS.address).setResolver(namehash(`${domain}.${tld}`), PublicResolver.address));

  deployer.deploy(MyContractV0);
  deployer.deploy(MyContractV1);
  deployer.deploy(KNSRegistryV1);
};
