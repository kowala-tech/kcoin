/* eslint-disable max-len */

const Web3 = require('web3');

// const web3 = new Web3(new Web3.providers.HttpProvider('http://0.0.0.0:30503'));
const web3 = new Web3(new Web3.providers.HttpProvider('http://127.0.0.1:8545'));

const namehash = require('eth-ens-namehash');

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
    .alias('k', 'privateKey') // admin's private key
    .nargs('k', 1)
    .describe('k', 'admin`s private key to sign a transaction')
    .alias('f', 'file') // file option
    .nargs('f', 1)
    .describe('f', 'path to a JSON file with ABI and Bytecode for the new version of a contract. Usually a file from build diretory after truffle compile')
    .demandOption(['a','k','f'])
    .help('h')
    .alias('h', 'help')
    .epilog('Copyright Kowala 2018')
    .argv;

const multiSigAddr = '0x0e5d0Fd336650E663C710EF420F85Fb081E21415';
const publicResolverAddr = '0x01e1056f6a829E53dadeb8a5A6189A9333Bd1d63';

const {
  AdminUpgradabilityProxyAbi,
  PublicResolverABI,
  signTransactionAndSend,
  readABIAndByteCode,
  deployContract,
} = require('../helpers/helpers.js');

(async () => {
  try {
    let proxyAddr = argv.contractAddr;
    let admin = argv.admin;
    let pk = argv.privateKey;
    let file = argv.file;

    if (!(await web3.utils.isAddress(argv.admin))) throw 'Admin field should be an address'
    if (argv.domain !== undefined && argv.contractAddr === undefined) {
      const publicResolver = new web3.eth.Contract(PublicResolverABI, publicResolverAddr);
      proxyAddr = await publicResolver.methods.addr(namehash(argv.domain)).call();
    } else if (argv.domain === undefined && argv.contractAddr !== undefined && await web3.utils.isAddress(argv.contractAddr)) {
      proxyAddr = argv.contractAddr;
    } else {
      throw 'domain or contract address should be populated';
    }

    const contractInternals = await readABIAndByteCode(file);
    const contractAddress = await deployContract(contractInternals[1], admin, pk);
    console.log('creating proxy contract object');
    const adminProxy = new web3.eth.Contract(AdminUpgradabilityProxyAbi, proxyAddr);
    console.log('created proxy contract object');
    console.log('upgrading...');
    const data = adminProxy.methods.upgradeTo(contractAddress).encodeABI();
    await signTransactionAndSend(data, proxyAddr, admin, pk);
    console.log('upgraded');
  } catch (err) {
    console.log(err);
  }
})();
