const Tx = require('ethereumjs-tx');
const Web3 = require('web3');

const web3 = new Web3('http://0.0.0.0:30503');

const account1 = '0x2429f4aa5cf9d23fea0961780ffb4ff8916a26a0';
const privateKey1 = Buffer.from('ae288a2d269e706e151c5c572b590d7d7fc478c243fde209b019b89054fd2737', 'hex');
const privateKey2 = '0xae288a2d269e706e151c5c572b590d7d7fc478c243fde209b019b89054fd2737';
const UpgradabilityProxyFactoryAbi = [{'anonymous':false,'inputs':[{'indexed':false,'name':'proxy','type':'address'}],'name':'ProxyCreated','type':'event'},{'constant':false,'inputs':[{'name':'admin','type':'address'},{'name':'implementation','type':'address'}],'name':'createProxy','outputs':[{'name':'','type':'address'}],'payable':false,'stateMutability':'nonpayable','type':'function'},{'constant':false,'inputs':[{'name':'admin','type':'address'},{'name':'implementation','type':'address'},{'name':'data','type':'bytes'}],'name':'createProxyAndCall','outputs':[{'name':'','type':'address'}],'payable':true,'stateMutability':'payable','type':'function'}];
const AdminUpgradabilityProxyAbi = [{'inputs':[{'name':'_implementation','type':'address'}],'payable':false,'stateMutability':'nonpayable','type':'constructor'},{'payable':true,'stateMutability':'payable','type':'fallback'},{'anonymous':false,'inputs':[{'indexed':false,'name':'previousAdmin','type':'address'},{'indexed':false,'name':'newAdmin','type':'address'}],'name':'AdminChanged','type':'event'},{'anonymous':false,'inputs':[{'indexed':false,'name':'implementation','type':'address'}],'name':'Upgraded','type':'event'},{'constant':true,'inputs':[],'name':'admin','outputs':[{'name':'','type':'address'}],'payable':false,'stateMutability':'view','type':'function'},{'constant':true,'inputs':[],'name':'implementation','outputs':[{'name':'','type':'address'}],'payable':false,'stateMutability':'view','type':'function'},{'constant':false,'inputs':[{'name':'newAdmin','type':'address'}],'name':'changeAdmin','outputs':[],'payable':false,'stateMutability':'nonpayable','type':'function'},{'constant':false,'inputs':[{'name':'newImplementation','type':'address'}],'name':'upgradeTo','outputs':[],'payable':false,'stateMutability':'nonpayable','type':'function'},{'constant':false,'inputs':[{'name':'newImplementation','type':'address'},{'name':'data','type':'bytes'}],'name':'upgradeToAndCall','outputs':[],'payable':true,'stateMutability':'payable','type':'function'}];
const SimpleContractAbi = [{"constant":true,"inputs":[],"name":"initialized","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_x","type":"uint256"}],"name":"setX","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"readX","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"}];

// async function main() {
//   const proxyFactoryContract = new web3.eth.Contract(UpgradabilityProxyFactoryAbi, '0x7A5727E94bbb559e0eAfC399354Dd30dBD51d2aa');
//   console.log("----1---");
//   console.log(proxyFactoryContract.options.address);
//   console.log("----2---");
//   const tmp = await proxyFactoryContract.methods.createProxy('0x2429f4aa5cf9d23fea0961780ffb4ff8916a26a0', '0x490fc7f8453e56f84abfc8fd17d74ad3fb6e819f').send({from: '0x2429f4aa5cf9d23fea0961780ffb4ff8916a26a0'});
//   console.log(tmp);
//   // const proxyContract = new web3.eth.Contract(AdminUpgradabilityProxyAbi, proxyAddr);
//   console.log("----3---");
//   console.log(proxyContract.options.address);
//   proxyContract.methods.admin().send({gasLimit: 5000000, from: web3.eth.coinbase}).then((a) => console.log(a));
// }
// main();
// const proxyFactoryContract = new web3.eth.Contract(UpgradabilityProxyFactoryAbi, '0x7A5727E94bbb559e0eAfC399354Dd30dBD51d2aa');
// console.log("----1---");
// console.log(proxyFactoryContract.options.address);
// console.log("----2---");
// proxyFactoryContract.methods.createProxy(account1, '0x490fc7f8453e56f84abfc8fd17d74ad3fb6e819f').send({gasLimit: 5000000, from: account1})
//   .then((err, proxyAddr) => {
//     const proxyContract = new web3.eth.Contract(AdminUpgradabilityProxyAbi, proxyAddr);
//     console.log(proxyContract);
//   });
// // const proxyContract = new web3.eth.Contract(AdminUpgradabilityProxyAbi, proxyAddr);
// console.log("----3---");
// console.log(proxyContract.options.address);
// proxyContract.methods.admin().send({gasLimit: 5000000, from: web3.eth.coinbase}).then((a) => console.log(a));

const contract = new web3.eth.Contract(UpgradabilityProxyFactoryAbi, '0x7A5727E94bbb559e0eAfC399354Dd30dBD51d2aa');
const txObject = {
  gasLimit: web3.utils.toHex(1000000), // Raise the gas limit to a much higher amount
  // gasPrice: web3.utils.toHex(web3.utils.toWei('10', 'gwei')),
  to: '0x7A5727E94bbb559e0eAfC399354Dd30dBD51d2aa',
  from: '0x2429f4aa5cf9d23fea0961780ffb4ff8916a26a0',
  data: contract.methods.createProxy('0x2429f4aa5cf9d23fea0961780ffb4ff8916a26a0', '0x490fc7f8453e56f84abfc8fd17d74ad3fb6e819f').encodeABI(),
};
web3.eth.accounts.signTransaction(txObject, privateKey2, function (error, result) {
  web3.eth.sendSignedTransaction(result.rawTransaction).on('receipt', console.log);
});


// web3.eth.getTransactionCount(account1, (err, txCount) => {
//   const txObject = {
//     nonce: txCount,
//     chaindId: 2,
//     gasLimit: web3.utils.toHex(1000000), // Raise the gas limit to a much higher amount
//     // gasPrice: web3.utils.toHex(web3.utils.toWei('10', 'gwei')),
//     to: '0x7A5727E94bbb559e0eAfC399354Dd30dBD51d2aa',
//     data: contract.methods.createProxy('0x2429f4aa5cf9d23fea0961780ffb4ff8916a26a0', '0x490fc7f8453e56f84abfc8fd17d74ad3fb6e819f').encodeABI(),
//   };

//   const tx = new Tx(txObject);
//   tx.sign(privateKey1);

//   const serializedTx = tx.serialize();
//   const raw = '0x' + serializedTx.toString('hex');
// //   web3.eth.sendSignedTransaction(raw).then((err, response) => {
// //     console.log(response);
// //     console.log(err);
// //  });
//   web3.eth.sendSignedTransaction(raw).on('receipt', console.log);
// });

