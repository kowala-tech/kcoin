/* global artifacts, web3 */
/* eslint-disable max-len */

const SimpleContract = artifacts.require('SimpleContract.sol');
const NameHash = artifacts.require('NameHash.sol');
const OracleMgr = artifacts.require('OracleMgr.sol');
const ValidatorMgr = artifacts.require('ValidatorMgr.sol');
const KNSRegistryV1 = artifacts.require('KNSRegistryV1.sol');

module.exports = (deployer) => {
  deployer.deploy(NameHash);
  deployer.link(NameHash, OracleMgr);
  deployer.link(NameHash, ValidatorMgr);

  deployer.deploy(SimpleContract);
  deployer.deploy(KNSRegistryV1);
};
