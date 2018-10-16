/* global artifacts */
/* eslint-disable max-len */

// const Web3 = require('web3');
// const web3 = new Web3(new Web3.providers.HttpProvider('http://0.0.0.0:30503'));

const MyContractV0 = artifacts.require('MyContractV0.sol');
const MyContractV1 = artifacts.require('MyContractV1.sol');

const acc1 = '0xEC057cCeA42C0b39086C5bc21FA0DD873bDa78A7';
const acc2 = '0xD1eFC9Ed211F8943B07a5Ce739d6C3a05e7EF4FD';
// const { Contracts } = require('zos-lib');

// const AdminUpgradeabilityProxy = Contracts.getFromNodeModules('zos-lib', 'AdminUpgradeabilityProxy');
const {
  AdminUpgradeabilityProxy, UpgradeabilityProxyFactory,
  PublicResolver, KNSRegistry,
  account1, governor1
} = require('./helpers.js');

module.exports = async () => {
  try {
    const proxyFactory = await UpgradeabilityProxyFactory.new({ from: acc2 });
    const myContractV0 = await MyContractV0.new();

    const logs = await proxyFactory.createProxy(acc1, myContractV0.address, { from: acc2 });
    const logs1 = logs.logs;
    const proxyAddr = logs1.find(l => l.event === 'ProxyCreated').args.proxy;
    const adminProxy = await AdminUpgradeabilityProxy.at(proxyAddr);
    const contract1 = await MyContractV0.at(adminProxy.address);
    const value = 42;
    await contract1.initialize(value, { from: acc2 });
    console.log((await contract1.value({ from: acc2 })).toString());

    // const myContractV1 = await MyContractV1.new();
    // await adminProxy.upgradeTo(myContractV1.address, { from: acc1 });
    // const contract2 = await MyContractV1.at(proxyAddr);
    // await contract2.add(5, { from: acc2 });
    // console.log(await contract2.value({ from: acc2 }));
  } catch (err) { console.log(err); }
};
