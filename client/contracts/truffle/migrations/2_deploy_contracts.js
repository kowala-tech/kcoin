/* global artifacts, web3 */

const ENS = artifacts.require('./ens/ENSRegistry.sol');
const FIFSRegistrar = artifacts.require('./ens/FIFSRegistrar.sol');
const PublicResolver = artifacts.require('./ens/PublicResolver.sol');
const namehash = require('eth-ens-namehash');

// /**
//  * Calculate root node hashes given the top level domain(tld)
//  *
//  * @param {string} tld plain text tld, for example: 'eth'
//  */
// function getRootNodeFromTLD(tld) {
//   return {
//     namehash: namehash(tld),
//     sha3: web3.sha3(tld),
//   };
// }

// /**
//  * Deploy the ENS and FIFSRegistrar
//  *
//  * @param {Object} deployer truffle deployer helper
//  * @param {string} tld tld which the FIFS registrar takes charge of
//  */
// function deployFIFSRegistrar(deployer, tld) {
//   const rootNode = getRootNodeFromTLD(tld);

//   // Deploy the ENS first
//   deployer.deploy(ENS)
//     .then(() => { return deployer.deploy(FIFSRegistrar, ENS.address, rootNode.namehash); })
//     .then(() => {
//       // Transfer the owner of the `rootNode` to the FIFSRegistrar
//       ENS.at(ENS.address).setSubnodeOwner('0x0', rootNode.sha3, FIFSRegistrar.address);
//     });
// }

// module.exports = (deployer) => {
//   const tld = 'eth';
//   deployFIFSRegistrar(deployer, tld);
// };

module.exports = (deployer) => {
  const domain = 'kowala';
  const tld = 'test';

  const rootNode = {
    namehash: namehash(tld),
    sha3: web3.sha3(tld),
  };


  deployer.deploy(ENS)
    .then(() => deployer.deploy(FIFSRegistrar, ENS.address, rootNode.namehash))
    .then(() => ENS.at(ENS.address).setSubnodeOwner('0x0', rootNode.sha3, FIFSRegistrar.address))
    .then(() => deployer.deploy(PublicResolver, ENS.address))
    .then(() => FIFSRegistrar.at(FIFSRegistrar.address).register(web3.sha3(domain), web3.eth.accounts[0]))
    .then(() => ENS.at(ENS.address).setResolver(namehash(`${domain}.${tld}`), PublicResolver.address));
};
