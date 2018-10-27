/* global artifacts, assert */
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
    .demandOption(['a','k','f'])
    .help('h')
    .alias('h', 'help')
    .epilog('Copyright Kowala 2018')
    .argv;

const Web3 = require('web3');
const namehash = require('eth-ens-namehash');
const assert = require('chai').assert;
const truffleContract = require('truffle-contract');

const web3 = new Web3(new Web3.providers.HttpProvider('http://127.0.0.1:8545'));

const MultiSig = artifacts.require('MultiSigWallet.sol');

const governor1 = '0xf861e10641952a42f9c527a43ab77c3030ee2c8f';
const governor3 = '0xa1d4755112491db5ddf0e10b9253b5a0f6783759';


const {
  AdminUpgradeabilityProxy,
  PublicResolver,
  readABIAndByteCode,
  getParamFromTxEvent,
} = require('./helpers.js');

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
      proxyAddr = await publicResolver.methods.addr(namehash(argv.domain)).call();
    } else if (argv.domain === undefined && argv.contractAddr !== undefined && await web3.utils.isAddress(argv.contractAddr)) {
      proxyAddr = argv.contractAddr;
    } else {
      throw 'domain or contract address should be populated';
    }

    const sig = await MultiSig.at(multiSigAddr);
    const adminProxy = await AdminUpgradeabilityProxy.at(proxyAddr);
    console.log(await adminProxy.admin({ from: multiSigAddr }));

    const contractInternals = await readABIAndByteCode(file);

    const ContractVersion2 = truffleContract({
      abi: contractInternals[0],
      bytecode: contractInternals[1],
    });
    ContractVersion2.setProvider(web3.currentProvider);

    const contractVersion2 = await ContractVersion2.new();
    const upgradeData = adminProxy.contract.upgradeTo.getData(contractVersion2.address);
    console.log("submitting");
    const transactionID = await getParamFromTxEvent(
      await sig.submitTransaction(adminProxy.address, 0, upgradeData, { from: admin }),
      'transactionId',
      null,
      'Submission',
    );
    console.log("submitted");

    console.log("confirming");
    await sig.confirmTransaction(transactionID, { from: governor3 });
    console.log("confirmed");
  } catch (err) { console.log(err); }
};
