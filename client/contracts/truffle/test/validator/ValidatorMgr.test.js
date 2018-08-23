/* global artifacts, contract, it, beforeEach, describe, before, web3 */
/* eslint no-unused-expressions: 0 */
/* eslint consistent-return: 0 */
/* eslint-disable max-len */

process.env.NODE_ENV = 'test';

require('chai')
  .use(require('chai-as-promised'))
  .use(require('chai-bignumber')(web3.BigNumber))
  .should();

const {
  EVMError,
} = require('../helpers/testUtils.js');

const ValidatorMgr = artifacts.require('ValidatorMgr.sol');
const PublicResolver = artifacts.require('PublicResolver.sol');
const KNS = artifacts.require('KNSRegistry.sol');
const FIFSRegistrar = artifacts.require('FIFSRegistrar.sol');
const MiningTokenMock = artifacts.require('TokenMock.sol');
const DomainResolverMock = artifacts.require('DomainResolverMock.sol');
const namehash = require('eth-ens-namehash');

contract('ValidatorMgr', ([_, owner, newOwner, newOwner2, newOwner3, newOwner4, notOwner]) => {
  describe('KNS functionality', async () => {
    beforeEach(async () => {
      this.kns = await KNS.new({ from: owner });
      this.registrar = await FIFSRegistrar.new(this.kns.address, namehash('kowala'));
      this.resolver = await PublicResolver.new(this.kns.address);

      await this.kns.setSubnodeOwner(0, web3.sha3('kowala'), this.registrar.address, { from: owner });
      await this.registrar.register(web3.sha3('miningtoken'), owner, { from: owner });
      await this.kns.setResolver(namehash('miningtoken.kowala'), this.resolver.address, { from: owner });
      this.miningToken = await MiningTokenMock.new();
      await this.resolver.setAddr(namehash('miningtoken.kowala'), this.miningToken.address, { from: owner });
      this.validator = await ValidatorMgr.new(100, 100, 0, 200, this.resolver.address, { from: owner });
    });

    it('should set MiningToken Address from KNS during creation', async () => {
      // given
      const knsResolverAddr = await this.validator.knsResolver();
      const resolver = await PublicResolver.at(knsResolverAddr);

      // when
      const miningTokenAddr = await resolver.addr(namehash('miningtoken.kowala'));

      // then
      await miningTokenAddr.should.be.equal(this.miningToken.address);
    });

    it('should set MiningToken Address from KNS during creation', async () => {
      // given
      const knsResolverAddr = await this.validator.knsResolver();
      const resolver = await PublicResolver.at(knsResolverAddr);

      // when
      const miningTokenAddr = await resolver.addr(namehash('miningtoken.kowala'));

      // then
      await miningTokenAddr.should.be.equal(this.miningToken.address);
    });
  });
  describe('Functionality of', async () => {
    beforeEach(async () => {
      this.miningToken = await MiningTokenMock.new();
      this.resolver = await DomainResolverMock.new(this.miningToken.address);
      // const byteCode = '0x60806040526000805460a060020a60ff021916905534801561002057600080fd5b5060405160a0806112958339810160409081528151602083015191830151606084015160809094015160008054600160a060020a0319163317815592949192841161006a57600080fd5b60018590556002849055620151808302600355600682905560078054600160a060020a031916600160a060020a038316179055604080517f09879962000000000000000000000000000000000000000000000000000000008152602060048201819052601260248301527f6d696e696e67746f6b656e2e6b6f77616c6100000000000000000000000000006044830152915173__NameHash______________________________9263098799629260648082019391829003018186803b15801561013357600080fd5b505af4158015610147573d6000803e3d6000fd5b505050506040513d602081101561015d57600080fd5b5051600555505050505061111f806101766000396000f3006080604052600436106101535763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663035cf14281146101585780630a3cb6631461017f5780632086ca251461019457806326833148146101a95780633ed0a373146101be5780633f4ba83a146101ef5780635c975abb14610206578063694746251461022f5780636a911ccf146102445780637071688a14610259578063715018a61461026e5780637d0e81bf146102835780638456cb59146102a45780638da5cb5b146102b95780639363a141146102ea57806397584b3e146102ff5780639abee7d0146103145780639bb2ea5a14610338578063a2207c6a14610350578063aded41ec14610365578063b774cb1e1461037a578063c22a933c1461038f578063cefddda9146103a7578063e7a60a9c146103c8578063f2fde38b14610403578063facd743b14610424575b600080fd5b34801561016457600080fd5b5061016d610445565b60408051918252519081900360200190f35b34801561018b57600080fd5b5061016d6104cf565b3480156101a057600080fd5b5061016d6104d5565b3480156101b557600080fd5b5061016d6104db565b3480156101ca57600080fd5b506101d66004356104e1565b6040805192835260208301919091528051918290030190f35b3480156101fb57600080fd5b50610204610527565b005b34801561021257600080fd5b5061021b61059d565b604080519115158252519081900360200190f35b34801561023b57600080fd5b5061016d6105ad565b34801561025057600080fd5b506102046105b3565b34801561026557600080fd5b5061016d6105e9565b34801561027a57600080fd5b506102046105f0565b34801561028f57600080fd5b5061021b600160a060020a036004351661065c565b3480156102b057600080fd5b506102046106ca565b3480156102c557600080fd5b506102ce610745565b60408051600160a060020a039092168252519081900360200190f35b3480156102f657600080fd5b5061016d610754565b34801561030b57600080fd5b5061021b61076a565b34801561032057600080fd5b50610204600160a060020a0360043516602435610779565b34801561034457600080fd5b506102046004356107c9565b34801561035c57600080fd5b506102ce610817565b34801561037157600080fd5b50610204610826565b34801561038657600080fd5b5061016d610a39565b34801561039b57600080fd5b50610204600435610a3f565b3480156103b357600080fd5b5061021b600160a060020a0360043516610a5b565b3480156103d457600080fd5b506103e0600435610a81565b60408051600160a060020a03909316835260208301919091528051918290030190f35b34801561040f57600080fd5b50610204600160a060020a0360043516610aee565b34801561043057600080fd5b5061021b600160a060020a0360043516610b11565b60008061045061076a565b1561045f5760015491506104cb565b60098054600891600091600019810190811061047757fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190206002810180549192509060001981019081106104b357fe5b90600052602060002090600202016000015460010191505b5090565b60035481565b60025481565b60065481565b3360009081526008602052604081206002018054829182918590811061050357fe5b90600052602060002090600202019050806000015481600101549250925050915091565b600054600160a060020a0316331461053e57600080fd5b60005460a060020a900460ff16151561055657600080fd5b6000805474ff0000000000000000000000000000000000000000191681556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b339190a1565b60005460a060020a900460ff1681565b60015481565b60005460a060020a900460ff16156105ca57600080fd5b6105d333610b11565b15156105de57600080fd5b6105e733610b32565b565b6009545b90565b600054600160a060020a0316331461060757600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b60008061066883610b11565b151561067757600091506106c4565b50600160a060020a038216600090815260086020526040902060065460029091018054909190829060001981019081106106ad57fe5b906000526020600020906002020160000154101591505b50919050565b600054600160a060020a031633146106e157600080fd5b60005460a060020a900460ff16156106f857600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a1781556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff6259190a1565b600054600160a060020a031681565b3360009081526008602052604090206002015490565b60095460025460009190031190565b60408051808201909152600160a060020a0383168082526020909101829052600a805473ffffffffffffffffffffffffffffffffffffffff19169091179055600b8190556107c5610c2e565b5050565b600080548190600160a060020a031633146107e357600080fd5b60095483101561081057505060095481900360005b8181101561081057610808610ca9565b6001016107f8565b5050600255565b600754600160a060020a031681565b6000805481908190819060a060020a900460ff161561084457600080fd5b33600090815260086020526040812090945084935060020191505b81548310801561088f5750818381548110151561087857fe5b906000526020600020906002020160010154600014155b156108f15781838154811015156108a257fe5b9060005260206000209060020201600101544210156108c0576108f1565b81838154811015156108ce57fe5b90600052602060002090600202016000015484019350828060010193505061085f565b6108fb3384610cda565b6000841115610a3357600754600554604080517f3b3b57de000000000000000000000000000000000000000000000000000000008152600481019290925251600160a060020a0390921691633b3b57de916024808201926020929091908290030181600087803b15801561096e57600080fd5b505af1158015610982573d6000803e3d6000fd5b505050506040513d602081101561099857600080fd5b5051604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152336004820152602481018790529051919250600160a060020a0383169163a9059cbb916044808201926020929091908290030181600087803b158015610a0657600080fd5b505af1158015610a1a573d6000803e3d6000fd5b505050506040513d6020811015610a3057600080fd5b50505b50505050565b60045481565b600054600160a060020a03163314610a5657600080fd5b600155565b600160a060020a0316600090815260086020526040902060010154610100900460ff1690565b6000806000600984815481101515610a9557fe5b6000918252602080832090910154600160a060020a031680835260089091526040909120600281018054929550909250906000198101908110610ad457fe5b906000526020600020906002020160000154915050915091565b600054600160a060020a03163314610b0557600080fd5b610b0e81610d81565b50565b600160a060020a031660009081526008602052604090206001015460ff1690565b600160a060020a038116600090815260086020526040902080545b60095460001901811015610bd0576009805460018301908110610b6c57fe5b60009182526020909120015460098054600160a060020a039092169183908110610b9257fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055600101610b4d565b6009805490610be3906000198301611069565b5060018201805460ff1916905560035460028301805442909201916000198101908110610c0c57fe5b906000526020600020906002020160010181905550610c29610dfe565b505050565b60005460a060020a900460ff1615610c4557600080fd5b600a54610c5a90600160a060020a0316610b11565b15610c6457600080fd5b610c6c610445565b600b541015610c7a57600080fd5b610c8261076a565b1515610c9057610c90610ca9565b600a54600b546105e791600160a060020a031690610e4e565b600980546105e791906000198101908110610cc057fe5b600091825260209091200154600160a060020a0316610b32565b60008080831515610cea57610d7a565b505050600160a060020a038216600090815260086020526040812090825b6002830154811015610d6c5760028301805482908110610d2457fe5b90600052602060002090600202018360020183815481101515610d4357fe5b600091825260209091208254600290920201908155600191820154908201559182019101610d08565b81610a30600285018261108d565b5050505050565b600160a060020a0381161515610d9657600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6009604051808280548015610e3c57602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610e1e575b50506040519081900390206004555050565b600160a060020a038216600081815260086020526040812060098054600181810183559184527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af8101805473ffffffffffffffffffffffffffffffffffffffff1916909517909455928155828101805460ff19169093179092558080431515610ee35760018401805461ff0019166101001790555b6040805180820190915285815260006020808301828152600280890180546001818101835591865293909420945192029093019081559151910155835492505b60008311156110615760086000600960018603815481101515610f4257fe5b6000918252602080832090910154600160a060020a031683528201929092526040019020600281018054919350906000198101908110610f7e57fe5b90600052602060002090600202019050806000015485111515610fa057611061565b600980546000198501908110610fb257fe5b60009182526020909120015460098054600160a060020a039092169185908110610fd857fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a031602179055508560096001850381548110151561101957fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039290921691909117905582825560001990920180845591610f23565b610a30610dfe565b815481835581811115610c2957600083815260209020610c299181019083016110b9565b815481835581811115610c2957600202816002028360005260206000209182019101610c2991906110d3565b6105ed91905b808211156104cb57600081556001016110bf565b6105ed91905b808211156104cb57600080825560018201556002016110d95600a165627a7a72305820490ca4f309376b2114d02f35efa0d84543c6258905854e3a736e9efab50a23330029';
      // // const contractData = await ValidatorMgr.new.getData(100, 3, 0, 3, this.resolver.address, { data: byteCode });
      // const gasEstimate = await web3.eth.estimateGas({ data: byteCode });
      // console.log(gasEstimate);
      this.validator = await ValidatorMgr.new(100, 3, 0, 3, this.resolver.address, { from: owner });
    });

    describe('registration', async () => {
      it('should register validator', async () => {
        // when
        await this.validator.registerValidator(newOwner, 100, { from: newOwner });

        // then
        const validatorCount = await this.validator.getValidatorCount();
        await validatorCount.should.be.bignumber.equal(1);
      });

      it('should not register validator when deposit is lower than base deposit', async () => {
        // when
        const exptectedFailedRegistration = this.validator.registerValidator(newOwner, 99, { from: newOwner });

        // then
        await exptectedFailedRegistration.should.eventually.be.rejectedWith(EVMError('revert'));
      });

      it('should not register validator with the same address', async () => {
        // given
        await this.validator.registerValidator(newOwner, 100, { from: newOwner });

        // when
        const exptectedFailedRegistration = this.validator.registerValidator(newOwner, 100, { from: newOwner });

        // then
        await exptectedFailedRegistration.should.eventually.be.rejectedWith(EVMError('revert'));
      });

      it('should not register validators above the limit', async () => {
        // given
        await this.validator.registerValidator(newOwner, 130, { from: newOwner });
        await this.validator.registerValidator(newOwner2, 120, { from: newOwner2 });
        await this.validator.registerValidator(newOwner3, 110, { from: newOwner3 });

        // when
        const exptectedFailedRegistration = this.validator.registerValidator(newOwner4, 100, { from: newOwner4 });

        // then
        await exptectedFailedRegistration.should.eventually.be.rejectedWith(EVMError('revert'));
      });
    });

    describe('deregistration', async () => {
      it('should deregister validator', async () => {
        // given
        await this.validator.registerValidator(newOwner, 100, { from: newOwner });

        // when
        await this.validator.deregisterValidator({ from: newOwner });

        // then
        const validatorCount = await this.validator.getValidatorCount();
        await validatorCount.should.be.bignumber.equal(0);
      });

      it('should not deregister deregistered validator', async () => {
        // given
        await this.validator.registerValidator(newOwner, 100, { from: newOwner });

        // when
        await this.validator.deregisterValidator({ from: newOwner });
        const exptectedFailedDeregistration = this.validator.deregisterValidator({ from: newOwner });

        // then
        await exptectedFailedDeregistration.should.eventually.be.rejectedWith(EVMError('revert'));
      });

      it('should not deregister non-validator', async () => {
        // when
        const exptectedFailedDeregistration = this.validator.deregisterValidator({ from: notOwner });

        // then
        await exptectedFailedDeregistration.should.eventually.be.rejectedWith(EVMError('revert'));
      });
    });

    describe('release deposits', async () => {
      it('should release deposit', async () => {
        // given
        const initialBalance = await this.miningToken.balanceOf(newOwner, { from: newOwner });
        await this.validator.registerValidator(newOwner, 100, { from: newOwner });
        await this.validator.deregisterValidator({ from: newOwner });

        // when
        await this.validator.releaseDeposits({ from: newOwner });

        // then
        const balanceAfterRelease = await this.miningToken.balanceOf(newOwner);
        await initialBalance.should.be.bignumber.equal(0);
        await balanceAfterRelease.should.be.bignumber.equal(100);
      });

      it('should not release deposit twice', async () => {
        // given
        await this.validator.registerValidator(newOwner, 100, { from: newOwner });
        await this.validator.deregisterValidator({ from: newOwner });
        await this.validator.releaseDeposits({ from: newOwner });
        const balanceFirstAfterRelease = await this.miningToken.balanceOf(newOwner);

        // when
        await this.validator.releaseDeposits({ from: newOwner });
        const balanceAfterSecondRelease = await this.miningToken.balanceOf(newOwner);

        // then
        await balanceAfterSecondRelease.should.be.bignumber.equal(balanceFirstAfterRelease);
      });
    });
    describe('setters and getters', async () => {
      it('should get validator`s deposit', async () => {
        // given
        await this.validator.registerValidator(newOwner, 150, { from: newOwner });

        // when
        const validatorDeposit = await this.validator.getDepositAtIndex(0, { from: newOwner });

        // then
        await validatorDeposit[0].should.be.bignumber.equal(150);
      });

      it('should not get validator`s deposit when calling by not owner', async () => {
        // given
        await this.validator.registerValidator(newOwner, 150, { from: newOwner });

        // when
        const expectedValidatorDepositFailure = this.validator.getDepositAtIndex(0, { from: notOwner });

        // then
        await expectedValidatorDepositFailure.should.eventually.be.rejectedWith(EVMError('invalid opcode'));
      });

      it('should get validator at index', async () => {
        // given
        await this.validator.registerValidator(newOwner, 100, { from: newOwner });

        // when
        const validatorAtIndex = await this.validator.getValidatorAtIndex(0, { from: newOwner });

        // then
        await validatorAtIndex[0].should.be.equal(newOwner);
      });

      it('should get minimal required deposit when there are no positions left', async () => {
        // given
        await this.validator.registerValidator(newOwner, 130, { from: newOwner });
        await this.validator.registerValidator(newOwner2, 120, { from: newOwner2 });
        await this.validator.registerValidator(newOwner3, 110, { from: newOwner3 });

        // when
        const minDeposit = await this.validator.getMinimumDeposit();

        // then
        await minDeposit.should.be.bignumber.equal(111);
      });

      it('should get base deposit when there are still positions available', async () => {
        // given
        await this.validator.registerValidator(newOwner, 130, { from: newOwner });
        await this.validator.registerValidator(newOwner2, 120, { from: newOwner });

        // when
        const minDeposit = await this.validator.getMinimumDeposit();

        // then
        await minDeposit.should.be.bignumber.equal(100);
      });

      it('should set base deposit', async () => {
        // given
        const currentBaseDeposit = await this.validator.getMinimumDeposit();

        // when
        await this.validator.setBaseDeposit(150, { from: owner });
        const newBaseDeposit = await this.validator.getMinimumDeposit();

        // then
        await currentBaseDeposit.should.be.bignumber.equal(100);
        await newBaseDeposit.should.be.bignumber.equal(150);
      });

      it('should not set base deposit by not owner', async () => {
        // when
        const exptectedFailedDepositSet = this.validator.setBaseDeposit(150, { from: notOwner });

        // then
        await exptectedFailedDepositSet.should.eventually.be.rejectedWith(EVMError('revert'));
      });
    });
  });
});
