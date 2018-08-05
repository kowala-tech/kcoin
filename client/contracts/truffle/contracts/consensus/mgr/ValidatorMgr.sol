pragma solidity 0.4.24;

import "openzeppelin-solidity/contracts/lifecycle/Pausable.sol";
import "../../token/ERC223.sol";
// @NOTE (rgeraldes) - https://github.com/kowala-tech/kcoin/client/issues/284
//import "github.com/kowala-tech/kcoin/client/contracts/token/contracts/TokenReceiver.sol" as receiver; 

/**
 * @title Validator Manager for PoS consensus
 */
contract ValidatorMgr is Pausable {
    uint public baseDeposit;       
    uint public maxNumValidators;
    uint public freezePeriod;
    bytes32 public validatorsChecksum;
    address public miningTokenAddr;
    uint public superNodeAmount;

    struct Deposit {
        uint amount;
        uint availableAt;
    }

    struct Validator {
        uint index;
        bool isValidator;
        bool isGenesis;

        // @NOTE (rgeraldes) - users can have more than one deposit
        // Example: user leaves and re-enters the election. At this point
        // the initial deposit will have a release date and the validator 
        // will have a new deposit for the current election.
        Deposit[] deposits; 
    }

    struct TKN {
        address sender;
        uint value;
        //bytes data;
        //bytes4 sig;
    }
    
    mapping (address => Validator) private validatorRegistry;
    
    // validatorPool contains the validator code ordered by the biggest deposit to
    // the smallest deposit.
    address[] validatorPool;

    TKN tkn;

    modifier onlyWithMinDeposit {
        require(tkn.value >= getMinimumDeposit());
        _;
    }

    modifier onlyValidator {
        require(isValidator(msg.sender));
        _;
    }

    modifier onlyNewCandidate {
        require(!isValidator(tkn.sender));
        _;
    }

    /**
     * Constructor.
     * @param _baseDeposit base deposit for Validator
     * @param _maxNumValidators Maximum numbers of Validators.
     * @param _freezePeriod Freeze period for Validator's deposits.
     * @param _miningTokenAddr Address of mining token.
     * @param _superNodeAmount Amount required to be considered a super node.
     */
    function ValidatorMgr(uint _baseDeposit, uint _maxNumValidators, uint _freezePeriod, address _miningTokenAddr, uint _superNodeAmount) public {
        require(_maxNumValidators >= 1);

        baseDeposit = _baseDeposit;
        maxNumValidators = _maxNumValidators;
        freezePeriod = _freezePeriod * 1 days;
        miningTokenAddr = _miningTokenAddr;
        superNodeAmount = _superNodeAmount;
    }

    /**
     * @dev Checks if given address is Genesis Validator
     * @param code Address of an Validator.
     */
    function isGenesisValidator(address code) public view returns (bool isIndeed) {
        return validatorRegistry[code].isGenesis;
    }

    /**
     * @dev Checks if given address is Validator
     * @param code Address of a Validator to check.
     */
    function isValidator(address code) public view returns (bool isIndeed) {
        return validatorRegistry[code].isValidator;
    }

    /**
     * @dev Checks if given address is super node
     * @param code Address of a supernode to check.
     */
    function isSuperNode(address code) public view returns (bool isIndeed) {
        if (!isValidator(code)) return false;

        Deposit[] deposits = validatorRegistry[code].deposits;
        return deposits[deposits.length - 1].amount >= superNodeAmount;
    }

    /**
     * @dev Get Validator count
     */
    function getValidatorCount() public view returns (uint count) {
        return validatorPool.length;
    }

    /**
     * @dev Get Validator information
     * @param index index of an Validator to check.
     */
    function getValidatorAtIndex(uint index) public view returns (address code, uint deposit) {
        code = validatorPool[index];
        Validator validator = validatorRegistry[code];
        deposit = validator.deposits[validator.deposits.length - 1].amount;
    }

    function _hasAvailability() public view returns (bool available) {
        return (maxNumValidators - validatorPool.length) > 0;
    }

    /**
     * @dev returns the base deposit if there are positions available or
            the current smallest deposit required if there aren't positions available.
     */
    function getMinimumDeposit() public view returns (uint deposit) {
        // there are positions for validator available
        if (_hasAvailability()) {
            return baseDeposit;
        } else {
            Validator smallestBidder = validatorRegistry[validatorPool[validatorPool.length - 1]];               
            return smallestBidder.deposits[smallestBidder.deposits.length - 1].amount + 1;
        }
    }

    /**
     * @dev updates the checksum
     */
    function _updateChecksum() private {
        validatorsChecksum = keccak256(validatorPool);
    }

    /**
     * @dev Add new validator
     * @param code Address of an Validator.
     * @param deposit amount to deposit
     */
    function _insertValidator(address code, uint deposit) private {
        Validator sender = validatorRegistry[code];
        sender.index = validatorPool.push(code) - 1;
        sender.isValidator = true;
        if (block.number == 0) sender.isGenesis = true;
        
        sender.deposits.push(Deposit({amount:deposit, availableAt: 0}));
        for (uint index = sender.index; index > 0; index--) {
            Validator target = validatorRegistry[validatorPool[index - 1]];
            Deposit collateral = target.deposits[target.deposits.length - 1];
            if (deposit <= collateral.amount) {
                break;
            }
            validatorPool[index] = validatorPool[index - 1];
            validatorPool[index - 1] = code; 
            // update indexes
            target.index = index;
            sender.index = index - 1;
        }

        _updateChecksum();
    }

    /**
     * @dev Set new base deposit for Validators
     */
    function setBaseDeposit(uint deposit) public onlyOwner {
        baseDeposit = deposit;
    }

    /**
     * @dev Set maximum numbers of Validators
     * @param max number of max Validators
     */
    function setMaxValidators(uint max) public onlyOwner { 
        if (max < validatorPool.length) {
            uint toRemove = validatorPool.length - max;
            for (uint i = 0; i < toRemove; i++) {
                _deleteSmallestBidder();
            }
        }
        maxNumValidators = max;   
    }

    /**
     * @dev Delete Validator
     * @param account address of a Validator
     */
    function _deleteValidator(address account) private {
        Validator validator = validatorRegistry[account];
        for (uint index = validator.index; index < validatorPool.length - 1; index++) {
            validatorPool[index] = validatorPool[index + 1];
        }
        validatorPool.length--;

        validator.isValidator = false;
        validator.deposits[validator.deposits.length - 1].availableAt = now + freezePeriod;

        _updateChecksum();
    }

    /**
     * @dev removes the Validator with the smallest deposit
     */
    function _deleteSmallestBidder() private {
        _deleteValidator(validatorPool[validatorPool.length - 1]);
    }
    
    /**
     * @dev Get deposit count
     */
    function getDepositCount() public view returns (uint count) {
        return validatorRegistry[msg.sender].deposits.length; 
    }

    /**
     * @dev Get deposit at given index
     * @param index index of a Validator to get deposit
     */
    function getDepositAtIndex(uint index) public view returns (uint amount, uint availableAt) {
        Deposit deposit = validatorRegistry[msg.sender].deposits[index];
        return (deposit.amount, deposit.availableAt);
    }

    /**
     * @dev Register new Validator
     */
    function _registerValidator() private whenNotPaused onlyNewCandidate onlyWithMinDeposit {
        if (!_hasAvailability()) {
            _deleteSmallestBidder();
        }
        _insertValidator(tkn.sender, tkn.value);
    }

    /**
     * @dev deregister Validator
     */
    function deregisterValidator() public whenNotPaused onlyValidator {
        _deleteValidator(msg.sender);
    }

    /**
     * @dev remove deposit
     * @param code address of a Validator
     * @param index index of a deposit
     */
    function _removeDeposits(address code, uint index) private {
        if (index == 0) return;

        Validator validator = validatorRegistry[code];
        uint lo = 0;
        uint hi = index;
        while (hi < validator.deposits.length) {
            validator.deposits[lo] = validator.deposits[hi];
            lo++;
            hi++;
        }
        validator.deposits.length = lo;
    }

    /**
     * @dev transfers locked deposit(s) back the user account if they are past the freeze period
     */
    function releaseDeposits() public whenNotPaused {
        uint refund = 0;
        uint i = 0;
        Deposit[] deposits = validatorRegistry[msg.sender].deposits;
        for (; i < deposits.length && deposits[i].availableAt != 0; i++) {
            if (now < deposits[i].availableAt) {
                // no need to iterate further since the release date
                // (if is different than 0) of the following deposits
                // will always be past than the current one.
                break;
            }
            refund += deposits[i].amount;
        }

        _removeDeposits(msg.sender, i);

        if (refund > 0) {
            ERC223 mtoken = ERC223(miningTokenAddr);
            mtoken.transfer(msg.sender, refund);
        }
    }
    
     /**
     * @dev Register Validator
     * @param _from from address
     * @param _value value to send
     * @param _data data to sent
     */
    function registerValidator(address _from, uint _value, bytes _data) public {
        //uint32 u = uint32(_data[3]) + (uint32(_data[2]) << 8) + (uint32(_data[1]) << 16) + (uint32(_data[0]) << 24);
        // SSTORE problem - expensive
        tkn = TKN(_from, _value/*, _data, bytes4(u)*/);
        _registerValidator();
    }
    
}