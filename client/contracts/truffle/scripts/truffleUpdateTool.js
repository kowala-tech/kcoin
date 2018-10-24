/* global artifacts, assert */
/* eslint-disable max-len */

const Web3 = require('web3');

// const web3 = new Web3(new Web3.providers.HttpProvider('http://127.0.0.1:30504'));
const web3 = new Web3(new Web3.providers.HttpProvider('http://127.0.0.1:8545'));

const MultiSig = artifacts.require('MultiSigWallet.sol');
const SysVars1 = artifacts.require('SystemVars.sol');
const SysVars2 = artifacts.require('SystemVars2.sol');
const MyContractV2 = artifacts.require('MyContractV2.sol');
const KNS = artifacts.require('KNSRegistry.sol');
const KNSV1 = artifacts.require('KNSRegistryV1.sol');

const governor1 = '0xf861e10641952a42f9c527a43ab77c3030ee2c8f';
const governor2 = '0x7dd43075b89c129bcd2cca1e2d680a6f3f30b5d9';
const validator1 = '0xa4a06cb7bcaa0162082f170d4b6f4b5360da9e11';
const validator2 = '0xa0b5cd08648cf5d29269356377ac06c538d7de96';

const acc1 = '0x7A9E64e6baA8BD021C9FAa72869e8A4aF953D906';
const acc2 = '0xD1eFC9Ed211F8943B07a5Ce739d6C3a05e7EF4FD';
const acc3 = '0x1f3A8b01516e6031870cCC7810214D85D39165cb';
const acc4 = '0xa028EEB91b1661589888D26a4c9814294241dA69';
const acc5 = '0xfaA8C086182086f6cA16610E58baE5A14fc5Fc26';

const namehash = require('eth-ens-namehash');
const assert = require('chai').assert;

const {
  AdminUpgradeabilityProxy,AdminUpgradabilityProxyAbi,
  UpgradeabilityProxyFactory,
  PublicResolver, KNSRegistry,
  signTransactionAndSend,
  readABIAndByteCode,
} = require('./helpers.js');

const kcoin = n => new web3.utils.BN(web3.utils.toWei(n, 'ether'));
// const {
//   AdminUpgradabilityProxyAbi,
//   PublicResolverABI,
//   signTransactionAndSend,
//   readABIAndByteCode,
//   deployContract,
// } = require('./helpers.js');
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
// const multiSigAddr = '0x0e5d0Fd336650E663C710EF420F85Fb081E21415';

module.exports = async () => {
  try {
    const sig = await MultiSig.new([acc2, acc3, acc4], 2, { from: acc1 });
    // const contractV2 = await MyContractV2.new();

    // const domain = 'systemvars.kowala';
    // const prAddress = '0x01e1056f6a829E53dadeb8a5A6189A9333Bd1d63';
    // const publicResolver = await PublicResolver.new();
    // const svAddr = await publicResolver.addr(namehash(domain));
    const proxyFactory = await UpgradeabilityProxyFactory.new({ from: acc1 });
    // KNS Proxy
    const kns = await KNS.new();
    const logs = await proxyFactory.createProxy(sig.address, kns.address, { from: acc1 });
    const logs1 = logs.logs;
    const knsProxyAddress = logs1.find(l => l.event === 'ProxyCreated').args.proxy;
    const knsProxy = await AdminUpgradeabilityProxy.at(knsProxyAddress);
    let knsContract = await KNS.at(knsProxyAddress);
    await knsContract.initialize(acc5);
    console.log(await knsContract.owner(0x0));
    const knsv1 = await KNSV1.new();
    // await knsProxy.upgradeTo(knsv1.address, { from: acc1 });

    const upgradeData = knsProxy.contract.upgradeTo.getData(knsv1.address);


    console.log("submitting");
    // const addOwnerData = contractV2.contract.setValue.getData(5);
    const transactionID = await getParamFromTxEvent(
      await sig.submitTransaction(knsProxy.address, 0, upgradeData, { from: acc2 }),
      'transactionId',
      null,
      'Submission',
    );
    console.log("confirming");
    const tmp = await sig.confirmTransaction(transactionID, { from: acc3 });
    console.log(tmp);
    // const transactionID = await getParamFromTxEvent(
    //   await sig.submitTransaction(adminProxy.address, 0, upgradeData, { from: acc2 }),
    //   'transactionId',
    //   null,
    //   'Submission',
    // );
    // console.log("confirming");
    // const tmp = await sig.confirmTransaction(transactionID, { from: acc3 });
    // console.log(tmp);
    console.log("confirmed");
    knsContract = await KNSV1.at(knsProxyAddress);

    const hello = await knsContract.helloProxy();
    console.log(hello);
    // const contract2 = await SysVars2.at(sys1Proxy);
    // console.log("8");
    // const tmp2 = await contract2.helloProxy();
    // console.log("9");
    // console.log(tmp2);
  } catch (err) { console.log(err); }
};
