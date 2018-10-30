/* global artifacts */
/* eslint-disable max-len */

const argv = require('yargs')
    .usage('Usage: $0 option contractAddr option admin option privateKey option file \n e.g $0 -c 0x123 -a 0xAd123 -k 9876655 -f ./build/contracts/SampleContract.json')
    .alias('c', 'contractAddr') // contractAddr option
    .nargs('c', 1)
    .describe('c', 'contract address to update')
    .alias('d', 'domain') // domain option
    .nargs('d', 1)
    .describe('d', 'domain of a contract to be updated')
    .alias('a', 'admin') // admin of the proxy contract (who updates the contract)
    .nargs('a', 1)
    .describe('a', 'admin of the proxy contract (who updates the contract)')
    .alias('f', 'file') // file option
    .nargs('f', 1)
    .describe('f', 'path to a JSON file with ABI and Bytecode for the new version of a contract. Usually a file from build diretory after truffle compile')
    .demandOption(['a','f'])
    .help('h')
    .alias('h', 'help')
    .epilog('Copyright Kowala 2018')
    .argv;

const Web3 = require('web3');
const namehash = require('eth-ens-namehash');
const assert = require('chai').assert;
const truffleContract = require('truffle-contract');

const fs = require('fs');
const path = require('path');
const JSON5 = require('json5');

const network = global.process.argv[4];
const source = fs.readFileSync(path.resolve(__dirname, '../../truffle.js'), 'utf8');
const truffleJS = JSON5.parse(source.slice(17, -2));
const host = truffleJS.networks[network]['host'];
const port = truffleJS.networks[network]['port'];

const web3 = new Web3(new Web3.providers.HttpProvider(`http:\/\/${host}:${port}`));

const {
  AdminUpgradeabilityProxy,
  PublicResolver,
  readABIAndByteCode,
  getParamFromTxEvent,
} = require('../helpers/helpers.js');

const MultiSig = artifacts.require('MultiSigWallet.sol');
const multiSigAddr = '0x0e5d0Fd336650E663C710EF420F85Fb081E21415';
const prAddress = '0x01e1056f6a829E53dadeb8a5A6189A9333Bd1d63';

module.exports = async () => {
  try {
    let proxyAddr;
    let admin;
    const file = argv.file;

    if (!(await web3.utils.isAddress(argv.admin))) throw 'Admin field should be an address';
    else { admin = argv.admin; }
    if (argv.domain !== undefined && argv.contractAddr === undefined) {
      const publicResolver = await PublicResolver.at(prAddress);
      proxyAddr = await publicResolver.addr(namehash(argv.domain));
    } else if (argv.domain === undefined && argv.contractAddr !== undefined && await web3.utils.isAddress(argv.contractAddr)) {
      proxyAddr = argv.contractAddr;
    } else {
      throw 'domain or contract address should be populated';
    }

    const sig = await MultiSig.at(multiSigAddr);
    const adminProxy = await AdminUpgradeabilityProxy.at(proxyAddr);
    console.log('Loading contract`s ABI and Bytecode');
    const contractInternals = await readABIAndByteCode(file);
    console.log('Creating contract object');
    const ContractVersion2 = truffleContract({
      abi: contractInternals[0],
      bytecode: contractInternals[1],
    });
    console.log('Setting provider');
    ContractVersion2.setProvider(web3.currentProvider);
    if (typeof ContractVersion2.currentProvider.sendAsync !== 'function') {
      ContractVersion2.currentProvider.sendAsync = function () {
        return ContractVersion2.currentProvider.send.apply(ContractVersion2.currentProvider, arguments);
      };
    }
    ContractVersion2.defaults({
      gas: 4712388,
      gasPrice: 1,
    });
    console.log('Creating contract instance');
    const contractVersion2 = await ContractVersion2.new({ from: admin });
    console.log('Getting update data for a transaction');
    const upgradeData = adminProxy.contract.upgradeTo.getData(contractVersion2.address);
    console.log('Submitting transaction');
    const transactionID = await getParamFromTxEvent(
      await sig.submitTransaction(adminProxy.address, 0, upgradeData, { from: admin }),
      'transactionId',
      null,
      'Submission',
    );
    console.log('Transaction submitted');
  } catch (err) { console.log(err); }
};
