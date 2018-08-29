/* global web3, assert */
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

const getParamFromTxEvent = async (transaction, paramName, contractFactory, eventName) => {
  assert.isObject(transaction);
  let logs = transaction.logs;
  if (eventName != null) {
    logs = logs.filter(l => l.event === eventName);
  }
  assert.equal(logs.length, 1, 'too many logs found!');
  const param = logs[0].args[paramName];
  if (contractFactory != null) {
    const contract = contractFactory.at(param);
    assert.isObject(contract, `getting ${paramName} failed for ${param}`);
    return contract;
  }
  return param;
};


const EVMError = message => `VM Exception while processing transaction: ${message}`;
const kcoin = n => new web3.BigNumber(web3.toWei(n, 'ether'));

module.exports = {
  kcoin,
  EVMError,
  addSeconds,
  mineNBlocks,
  mineBlock,
  getParamFromTxEvent,
};
