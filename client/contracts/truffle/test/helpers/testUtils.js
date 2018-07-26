/* global web3 */
/* eslint no-unused-expressions: 0 */

require('chai')
  .use(require('chai-as-promised'))
  .use(require('chai-bignumber')())
  .should();


const mineBlock = () => (
  new Promise((resolve, reject) =>
    web3.currentProvider.sendAsync({
      jsonrpc: '2.0',
      method: 'evm_mine',
      id: new Date().getTime(),
    }, (error, result) => (error ? reject(error) : resolve(result.result))))
);

const mineNBlocks = async (n) => {
  for (let i = 0; i < n; i += 1) {
    mineBlock();
  }
};

const addSeconds = seconds => (
  new Promise((resolve, reject) =>
    web3.currentProvider.sendAsync({
      jsonrpc: '2.0',
      method: 'evm_increaseTime',
      params: [seconds],
      id: new Date().getTime(),
    }, (error, result) => (error ? reject(error) : resolve(result.result))))
    .then(mineBlock)
);

const EVMError = message => `VM Exception while processing transaction: ${message}`;
const ether = n => new web3.BigNumber(web3.toWei(n, 'ether'));

module.exports = {
  ether,
  EVMError,
  addSeconds,
  mineNBlocks,
  mineBlock,
};
