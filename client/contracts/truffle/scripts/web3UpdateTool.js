/* eslint-disable max-len */

const Web3 = require('web3');

// const web3 = new Web3(new Web3.providers.HttpProvider('http://0.0.0.0:30503'));
const web3 = new Web3(new Web3.providers.HttpProvider('http://127.0.0.1:8545'));

const namehash = require('eth-ens-namehash');
const commandLineArgs = require('command-line-args');

const optionDefinitions = [
  { name: 'contractAddr', alias: 'c', type: String },
  { name: 'domain', alias: 'd', type: String },
  { name: 'admin', alias: 'a', type: String },
  { name: 'privateKey', alias: 'k', type: String },
  { name: 'file', alias: 'f', type: String }
];
const options = commandLineArgs(optionDefinitions);

const multiSigAddr = '0x0e5d0Fd336650E663C710EF420F85Fb081E21415';
const publicResolverAddr = '0x01e1056f6a829E53dadeb8a5A6189A9333Bd1d63';

const {
  AdminUpgradabilityProxyAbi,
  PublicResolverABI,
  signTransactionAndSend,
  readABIAndByteCode,
  deployContract,
} = require('./helpers.js');

(async () => {
  try {
    let proxyAddr;
    let admin;
    let pk;

    if (options.admin === undefined || !(await web3.utils.isAddress(options.admin))) throw 'Admin field should be populated';
    else admin = options.admin;
    if (options.privateKey === undefined) throw 'Private key field should be populated';
    else pk = options.privateKey;
    if (options.domain !== undefined && options.contractAddr === undefined) {
      const publicResolver = new web3.eth.Contract(PublicResolverABI, publicResolverAddr);
      proxyAddr = await publicResolver.methods.addr(namehash(options.domain)).call(); 
    } else if (options.domain === undefined && options.contractAddr !== undefined && await web3.utils.isAddress(options.contractAddr)) {
      proxyAddr = options.contractAddr;
    } else {
      throw 'domain or contract address should be populated';
    }

    const contractInternals = await readABIAndByteCode(options.file);
    const contractAddress = await deployContract(contractInternals[1], admin, pk);

    const adminProxy = new web3.eth.Contract(AdminUpgradabilityProxyAbi, proxyAddr);
    console.log('created proxy contract object');
    const data = adminProxy.methods.upgradeTo(contractAddress).encodeABI();
    await signTransactionAndSend(data, proxyAddr, admin, pk);
  } catch (err) {
    console.log(err);
  }
})();
