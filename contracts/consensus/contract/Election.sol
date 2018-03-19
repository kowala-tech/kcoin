pragma solidity 0.4.21;

import "./ownable.sol";

/**
 * @title Election
 * @dev The Election contract manages the consensus validator set and provides basic functions
 * to join, leave and redeem locked funds.
 */
contract Election is Ownable {
    uint public baseDeposit;       
    uint public maxValidators;
    uint public unbondingPeriod;
    address public genesis;

    // validatorsChecksum is a representation of the current set of validators
    bytes32 public validatorsChecksum;

    // Deposit represents the collateral - staked tokens
    struct Deposit {
        uint amount;
        uint releasedAt;
    }

    // Validator represents a consensus validator      
    struct Validator {
        uint index;
        bool isValidator;

        // @NOTE (rgeraldes) - users can have more than one deposit
        // Example: user leaves and re-enters the election. At this point
        // the initial deposit will have a release date and the validator 
        // will have a new deposit for the current election.
        Deposit[] deposits; 
    }
    
    mapping (address => Validator) private validators;
    
    // validatorIndex contains the validator code ordered by the biggest deposit to
    // the smallest deposit.
    address[] validatorIndex;

    // onlyWithMinDeposit requires a minimum deposit to proceed
    modifier onlyWithMinDeposit {
        require(msg.value >= getMinimumDeposit());
        _;
    }

    // onlyValidator requires the sender to be a validator
    modifier onlyValidator {
        require(isValidator(msg.sender));
        _;
    }

    function Election(uint _baseDeposit, uint _maxValidators, uint _unbondingPeriod, address _genesis) public payable {
        require(msg.value >= _baseDeposit);
        require(_maxValidators >= 1);

        baseDeposit = _baseDeposit;
        maxValidators = _maxValidators;
        unbondingPeriod = _unbondingPeriod;
        genesis = _genesis;

        _insertValidator(_genesis, msg.value);
    }

    function _updateChecksum() private {
        validatorsChecksum = keccak256(validatorIndex);
    }

    function _insertValidator(address code, uint deposit) private {
        Validator sender = validators[code];
        sender.index = validatorIndex.push(code) - 1;
        sender.isValidator = true;
        sender.deposits.push(Deposit({amount:deposit, releasedAt: 0}));

        // ordered insert based on the deposit amount
        for (uint index = sender.index; index > 0; index--) {
            Validator target = validators[validatorIndex[index - 1]];
            Deposit collateral = target.deposits[validatorIndex.length - 1];
            if (deposit <= collateral.amount) {
                break;
            }
            validatorIndex[index] = validatorIndex[index - 1];
            validatorIndex[index - 1] = code; 
        }

        _updateChecksum();
    }

    function _deleteValidator(address account) private {
        Validator validator = validators[account];
        for (uint index = validator.index; index < validatorIndex.length - 1; index++) {
            validatorIndex[index] = validatorIndex[index + 1];
        }
        validatorIndex.length--;

        validator.isValidator = false;
        validator.deposits[validator.deposits.length - 1].releasedAt = now + unbondingPeriod;

        _updateChecksum();
    }

    // _deleteSmallestBidder removes the validator with the smallest deposit
    function _deleteSmallestBidder() private {
        _deleteValidator(validatorIndex[validatorIndex.length - 1]);
    }

    function setBaseDeposit(uint deposit) public onlyOwner {
        baseDeposit = deposit;
    }

    function setMaxValidators(uint max) public onlyOwner { 
        if (max < validatorIndex.length) {
            uint toRemove = validatorIndex.length - max;
            for (uint i = 0; i < toRemove; i++) {
                _deleteSmallestBidder();
            }
        }
        maxValidators = max;   
    }

    function getValidatorCount() public view returns (uint count) {
        return validatorIndex.length;
    }

    function getValidatorAtIndex(uint index) public view returns (address code, uint deposit) {
        code = validatorIndex[index];
        Validator validator = validators[code];
        deposit = validator.deposits[validator.deposits.length - 1].amount;
    }

    function isGenesisValidator(address code) public view returns (bool isIndeed) {
        return code == genesis;
    }

    function getDepositCount() public view returns (uint count) {
        return validators[msg.sender].deposits.length; 
    }

    function getDepositAtIndex(uint index) public view returns (uint amount, uint releasedAt) {
        Deposit deposit = validators[msg.sender].deposits[index];
        return (deposit.amount, deposit.releasedAt);
    }

    function _hasAvailability() public view returns (bool available) {
        return (maxValidators - validatorIndex.length) > 0;
    }

    // getMinimumDeposit returns the base deposit if there are positions available or
    // the current smallest deposit required if there aren't positions availabe.
    function getMinimumDeposit() public view returns (uint deposit) {
        // there are positions for validator available
        if (_hasAvailability()) {
            return baseDeposit;
        } else {
            Validator smallestBidder = validators[validatorIndex[validatorIndex.length - 1]];               
            return smallestBidder.deposits[smallestBidder.deposits.length - 1].amount + 1;
        }
    }

    // join registers a new candidate as validator
    function join() public payable onlyWithMinDeposit {
        if (!_hasAvailability()) {
            _deleteSmallestBidder();
        }
        _insertValidator(msg.sender, msg.value);
    }
    
    // @NOTE (rgeraldes) - not sure if it's going to be used in the future as
    // the transaction can fail and we need to look at the transaction status
    // instead.
    function isValidator(address code) public view returns (bool isIndeed) {
        return validators[code].isValidator;
    }

    // leave deregisters the msg sender from the validator set
    function leave() public onlyValidator {
        _deleteValidator(msg.sender);
    }

    /*
    function _removeDeposits(Validator validator, uint index) private {
        if (index == 0) return;
        // left shifts
        uint lo = 0;
        uint hi = index;
        // @TODO (rgeraldes) - review
        while (hi < validator.deposits.length) {
             validator.deposits[lo] = validator.deposits[hi];
             lo++;
             hi++;
        }
        // resize array
        // @TODO (rgeraldes)
        // validator.deposits.length--;
    }
    */

    // withdraw transfer locked deposit(s) back the user account if they
    // are past the unbonding period
    function redeemFunds() public returns {
        uint i = 0;
        Validator validator = validators[msg.sender];
        for (; i < validator.deposits.length && validator.deposits[i].releasedAt != 0; i++) {
            if (now < validator.deposits[i].releasedAt) {
                // @NOTE (rgeraldes) - no need to iterate further since the 
                // release date (if is different than 0) of the following deposits
                // will always be past than the current one.
                break;
            }
            refund += validator.deposits[i].amount;
        }
        
        //_removeDeposits(validator, i);

        if (refund > 0) {
            msg.sender.transfer(refund);
        }
    }
}