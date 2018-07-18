/* global assert */
/* eslint no-unused-expressions: 0 */

require('chai')
  .use(require('chai-as-promised'))
  .use(require('chai-bignumber')())
  .should();

function isException(error) {
  const strError = error.toString();
  return strError.includes('invalid opcode') || strError.includes('invalid JUMP') || strError.includes('revert');
}

function ensureException(error) {
  assert(isException(error), error.toString());
}

module.exports = {
  zeroAddress: '0x0000000000000000000000000000000000000000',
  ensureException,
};
