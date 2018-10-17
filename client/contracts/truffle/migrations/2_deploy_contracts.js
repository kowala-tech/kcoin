/* global artifacts, web3 */
/* eslint-disable max-len */

const MyContractV0 = artifacts.require('MyContractV0.sol');
const MyContractV1 = artifacts.require('MyContractV1.sol');
const NameHash = artifacts.require('NameHash.sol');
const OracleMgr = artifacts.require('OracleMgr.sol');
const ValidatorMgr = artifacts.require('ValidatorMgr.sol');
const KNSRegistryV1 = artifacts.require('KNSRegistryV1.sol');

module.exports = (deployer) => {
  deployer.deploy(NameHash);
  deployer.link(NameHash, OracleMgr);
  deployer.link(NameHash, ValidatorMgr);

  deployer.deploy(MyContractV0);
  deployer.deploy(MyContractV1);
  deployer.deploy(KNSRegistryV1);
};
