/* global artifacts, assert */
/* eslint-disable max-len */

const argv = require('yargs')
    .usage('Usage: $0 option maxNumValidators \n e.g $0 -n 600')
    .alias('n', 'maxNumValidators') // contractAddr option
    .nargs('n', 1)
    .describe('n', 'number of validators')
    .demandOption(['n'])
    .help('h')
    .alias('h', 'help')
    .epilog('Copyright Kowala 2018')
    .argv;

const Web3 = require('web3');
const namehash = require('eth-ens-namehash');
const assert = require('chai').assert;

const web3 = new Web3(new Web3.providers.HttpProvider('http://0.0.0.0:30503'));

const MultiSig = artifacts.require('MultiSigWallet.sol');
const ValidatorMgr = artifacts.require('ValidatorMgr.sol');

const governor1 = '0xf861e10641952a42f9c527a43ab77c3030ee2c8f';
const governor3 = '0xa1d4755112491db5ddf0e10b9253b5a0f6783759';

const {
  AdminUpgradeabilityProxy,
  PublicResolver,
  getParamFromTxEvent,
} = require('../helpers/helpers.js');

const multiSigAddr = '0x0e5d0fd336650e663c710ef420f85fb081e21415';
const prAddress = '0x01e1056f6a829E53dadeb8a5A6189A9333Bd1d63';

module.exports = async () => {
  try {
    const sig = await MultiSig.at(multiSigAddr);
    const publicResolver = await PublicResolver.at(prAddress);
    const validatorAddr = await publicResolver.addr(namehash('validatormgr.kowala'));
    const proxyAdmin = await AdminUpgradeabilityProxy.at(validatorAddr);
    const implValAddr = await proxyAdmin.implementation({ from: multiSigAddr });
    const validator = await ValidatorMgr.at(implValAddr);
    const upgradeData = validator.contract.setMaxValidators.getData(argv.maxNumValidators);
    console.log("submitting");
    const transactionID = await getParamFromTxEvent(
      await sig.submitTransaction(validator.address, 0, upgradeData, { from: governor1 }),
      'transactionId',
      null,
      'Submission',
    );
    console.log("submitted");
    console.log("confirming");
    await sig.confirmTransaction(transactionID, { from: governor3 });
    console.log("confirmed");
    console.log(`New number of validators: ${await validator.maxNumValidators()}`);
  } catch (err) { console.log(err); }
};
