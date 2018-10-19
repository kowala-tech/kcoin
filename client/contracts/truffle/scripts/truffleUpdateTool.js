/* global artifacts, assert */
/* eslint-disable max-len */

// const Web3 = require('web3');
// const web3 = new Web3(new Web3.providers.HttpProvider('http://0.0.0.0:30503'));

const MultiSig = artifacts.require('MultiSigWallet.sol');
const SysVars1 = artifacts.require('SystemVars.sol');
const SysVars2 = artifacts.require('SystemVars2.sol');

const governor1 = '0xf861e10641952a42f9c527a43ab77c3030ee2c8f';
const governor2 = '0xa1d4755112491db5ddf0e10b9253b5a0f6783759';
const namehash = require('eth-ens-namehash');
const assert = require('chai').assert;

const {
  AdminUpgradeabilityProxy, UpgradeabilityProxyFactory,
  PublicResolver, KNSRegistry
} = require('./helpers.js');


const getParamFromTxEvent = async (transaction, paramName, contractFactory, eventName) => {
  assert.isObject(transaction);
  let logs = transaction.logs;
  if (eventName != null) {
    logs = logs.filter(l => l.event === eventName);
  }
  assert.equal(logs.length, 1, 'too many logs found!');
  const param = logs[0].args[paramName];
  if (contractFactory != null) {
    const contract = contractFactory.at(param);
    assert.isObject(contract, `getting ${paramName} failed for ${param}`);
    return contract;
  }
  return param;
};
const multiSigAddr = '0x0e5d0Fd336650E663C710EF420F85Fb081E21415';

module.exports = async () => {
  try {
    const domain = 'systemvars.kowala';
    const sig = await MultiSig.at(multiSigAddr);
    const owners = await sig.getOwners();
    console.log(owners);
    const prAddress = '0x01e1056f6a829E53dadeb8a5A6189A9333Bd1d63';
    const publicResolver = await PublicResolver.at(prAddress);
    const svAddr = await publicResolver.addr(namehash(domain));
    const adminProxy = await AdminUpgradeabilityProxy.at(svAddr);
    const contract1 = await SysVars1.at(adminProxy.address);
    const sys2 = await SysVars2.new({ from: governor1 });
    const upgradeData = await adminProxy.contract.upgradeTo.getData(sys2.address);
    console.log("submitting");
    const transactionID = await getParamFromTxEvent(
      await sig.submitTransaction(adminProxy.address, 0, upgradeData, { from: governor1 }),
      'transactionId',
      null,
      'Submission',
    );
    console.log("confirming");
    const tmp = await sig.confirmTransaction(transactionID, { from: governor2 });
    console.log(tmp);
    console.log("confirmed");
    
    // const contract2 = await SysVars2.at(adminProxy.address);
    // const tmp = await contract2.helloProxy({ from: governor2 });

  } catch (err) { console.log(err); }
};
