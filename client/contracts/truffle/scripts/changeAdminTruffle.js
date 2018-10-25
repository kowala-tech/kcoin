/* eslint-disable max-len */

const Web3 = require('web3');

const web3 = new Web3(new Web3.providers.HttpProvider('http://0.0.0.0:30503'));

const gov1 = '0x2429f4aa5cf9d23fea0961780ffb4ff8916a26a0';
const gov2 = '0xf861e10641952a42f9c527a43ab77c3030ee2c8f';

const multisigcreatorAddr = '0xFF9DFBD395cD1C4a4F23C16aa8a5c44109Bc17DF';
const multisigAddr = '0x0e5d0Fd336650E663C710EF420F85Fb081E21415';
const proxyFactoryAddr = '0x7A5727E94bbb559e0eAfC399354Dd30dBD51d2aa';

const knsRegistryAddr = '0x4195B06a6e4d5bEDDe15165e01A64F324f03D5d1';
const knsRegistrarAddr = '0xe1adb6075619f52Fc00BDD50dEf1B754b9e7bd17';
const knsResolverAddr = '0x01e1056f6a829E53dadeb8a5A6189A9333Bd1d63';

const namehash = require('eth-ens-namehash');

const {
  AdminUpgradeabilityProxy,
  PublicResolver,
} = require('./helpers.js');




module.exports = async () => {
  try {
    const domain1 = 'systemvars.kowala';
    const domain3 = 'oraclemgr.kowala';
    const domain4 = 'validatormgr.kowala';
    const domain5 = 'miningtoken.kowala';
    const domain6 = 'stability.kowala';

    const publicResolver = await PublicResolver.at(knsResolverAddr);

    const sysvarAddr = await publicResolver.addr(namehash(domain1));
    console.log(sysvarAddr);
    const oraclemgrAddr = await publicResolver.addr(namehash(domain3));
    console.log(oraclemgrAddr);
    const validatormgrAdrr = await publicResolver.addr(namehash(domain4));
    console.log(validatormgrAdrr);
    const miningtokenAddr = await publicResolver.addr(namehash(domain5));
    console.log(miningtokenAddr);
    const stabilityAddr = await publicResolver.addr(namehash(domain6));
    console.log(stabilityAddr);

    const knsRegistryProxy = await AdminUpgradeabilityProxy.at(knsRegistryAddr);
    const knsResolverProxy = await AdminUpgradeabilityProxy.at(knsResolverAddr);
    const knsRegistrarProxy = await AdminUpgradeabilityProxy.at(knsRegistrarAddr);
    
    const sysvarProxy = await AdminUpgradeabilityProxy.at(sysvarAddr);
    const oraclemgrProxy = await AdminUpgradeabilityProxy.at(oraclemgrAddr);
    const validatormgrProxy = await AdminUpgradeabilityProxy.at(validatormgrAdrr);
    const miningtokenProxy = await AdminUpgradeabilityProxy.at(miningtokenAddr);
    const stabilityProxy = await AdminUpgradeabilityProxy.at(stabilityAddr);

    await sysvarProxy.changeAdmin(multisigAddr, { from: multisigcreatorAddr });
    console.log(sysvarProxy.admin({ from: multisigAddr }));
    await oraclemgrProxy.changeAdmin(multisigAddr, { from: multisigcreatorAddr });
    console.log(oraclemgrProxy.admin({ from: multisigAddr }));
    await validatormgrProxy.changeAdmin(multisigAddr, { from: multisigcreatorAddr });
    console.log(validatormgrProxy.admin({ from: multisigAddr }));
    await miningtokenProxy.changeAdmin(multisigAddr, { from: multisigcreatorAddr });
    console.log(miningtokenProxy.admin({ from: multisigAddr }));
    await stabilityProxy.changeAdmin(multisigAddr, { from: multisigcreatorAddr });
    console.log(stabilityProxy.admin({ from: multisigAddr }));
    await knsRegistryProxy.changeAdmin(multisigAddr, { from: multisigcreatorAddr });
    console.log(knsRegistryProxy.admin({ from: multisigAddr }));
    await knsResolverProxy.changeAdmin(multisigAddr, { from: multisigcreatorAddr });
    console.log(knsResolverProxy.admin({ from: multisigAddr }));
    await knsRegistrarProxy.changeAdmin(multisigAddr, { from: multisigcreatorAddr });
  } catch (err) { console.log(err); }
};
