/* global artifacts, web3 */
/* eslint-disable max-len */

const SimpleContract = artifacts.require('SimpleContract.sol');
const Tx = require('ethereumjs-tx');

const account1 = '0x2429f4aa5cf9d23fea0961780ffb4ff8916a26a0';
const privateKey1 = 'ae288a2d269e706e151c5c572b590d7d7fc478c243fde209b019b89054fd2737';

const UpgradabilityProxyFactoryAbi = [{'anonymous':false,'inputs':[{'indexed':false,'name':'proxy','type':'address'}],'name':'ProxyCreated','type':'event'},{'constant':false,'inputs':[{'name':'admin','type':'address'},{'name':'implementation','type':'address'}],'name':'createProxy','outputs':[{'name':'','type':'address'}],'payable':false,'stateMutability':'nonpayable','type':'function'},{'constant':false,'inputs':[{'name':'admin','type':'address'},{'name':'implementation','type':'address'},{'name':'data','type':'bytes'}],'name':'createProxyAndCall','outputs':[{'name':'','type':'address'}],'payable':true,'stateMutability':'payable','type':'function'}];
const AdminUpgradabilityProxyAbi = [{'inputs':[{'name':'_implementation','type':'address'}],'payable':false,'stateMutability':'nonpayable','type':'constructor'},{'payable':true,'stateMutability':'payable','type':'fallback'},{'anonymous':false,'inputs':[{'indexed':false,'name':'previousAdmin','type':'address'},{'indexed':false,'name':'newAdmin','type':'address'}],'name':'AdminChanged','type':'event'},{'anonymous':false,'inputs':[{'indexed':false,'name':'implementation','type':'address'}],'name':'Upgraded','type':'event'},{'constant':true,'inputs':[],'name':'admin','outputs':[{'name':'','type':'address'}],'payable':false,'stateMutability':'view','type':'function'},{'constant':true,'inputs':[],'name':'implementation','outputs':[{'name':'','type':'address'}],'payable':false,'stateMutability':'view','type':'function'},{'constant':false,'inputs':[{'name':'newAdmin','type':'address'}],'name':'changeAdmin','outputs':[],'payable':false,'stateMutability':'nonpayable','type':'function'},{'constant':false,'inputs':[{'name':'newImplementation','type':'address'}],'name':'upgradeTo','outputs':[],'payable':false,'stateMutability':'nonpayable','type':'function'},{'constant':false,'inputs':[{'name':'newImplementation','type':'address'},{'name':'data','type':'bytes'}],'name':'upgradeToAndCall','outputs':[],'payable':true,'stateMutability':'payable','type':'function'}] ;

module.exports = (callback) => {
  const proxyFactoryContract = web3.eth.contract(UpgradabilityProxyFactoryAbi).at('0x7A5727E94bbb559e0eAfC399354Dd30dBD51d2aa');

  // SimpleContract.deployed().then((instance) => {
  //   return instance.readX();
  // }).then((response) => {
  //   console.log(response);
  // });
  // const txAddr = proxyFactoryContract.createProxy(web3.eth.coinbase, SimpleContract.address, { from: '0x2429f4aa5cf9d23fea0961780ffb4ff8916a26a0' });

  web3.eth.getTransactionCount(account1, (err, txCount) => {
    const txObject = {
      nonce: web3.toHex(txCount),
      gasLimit: web3.toHex(800000), // Raise the gas limit to a much higher amount
      gasPrice: web3.toHex(web3.toWei('10', 'gwei')),
      to: proxyFactoryContract.address,
      data: proxyFactoryContract.abi.createProxy(web3.eth.coinbase, SimpleContract.address, { from: '0x2429f4aa5cf9d23fea0961780ffb4ff8916a26a0' }).encodeABI(),
    };

    const tx = new Tx(txObject);
    tx.sign(privateKey1);

    const serializedTx = tx.serialize();
    const raw = '0x' + serializedTx.toString('hex');

    web3.eth.sendSignedTransaction(raw, (err, txHash) => {
      console.log('err:', err, 'txHash:', txHash);
      // Use this txHash to find the contract on Etherscan!
    });
  });

  // const txObject = web3.eth.getTransaction(txAddr);
  // const tx = new Tx(txObject);
  // tx.sign(privateKey1);

  // const serializedTx = tx.serialize();
  // const raw = '0x' + serializedTx.toString('hex');

  // // Broadcast the transaction
  // web3.eth.sendSignedTransaction(raw, (err, txHash) => {
  //   console.log('txHash:', txHash);
  //   // Now go check etherscan to see the transaction!
  // });
  // console.log(HashTx);
  // const logs = web3.eth.getTransactionReceipt(HashTx);
  // console.log(logs);
};
