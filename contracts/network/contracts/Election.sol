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
    // period in days
    uint public unbondingPeriod;
    address public genesisValidator;

    // validatorsChecksum is a representation of the current set of validators
    bytes32 public validatorsChecksum;

    // Deposit represents the collateral - staked tokens
    struct Deposit {
        uint amount;
        uint availableAt;
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

    // onlyNewCandidate required the sender to be a new candidate
    modifier onlyNewCandidate {
        require(!isValidator(msg.sender));
        _;
    }

    function Election(uint _baseDeposit, uint _maxValidators, uint _unbondingPeriod, address _genesis) public {
        require(_maxValidators >= 1);

        baseDeposit = _baseDeposit * 1 ether;
        maxValidators = _maxValidators;
        unbondingPeriod = _unbondingPeriod * 1 days;
        genesisValidator = _genesis;
    
        _insertValidator(_genesis, baseDeposit);
    }

    function isGenesisValidator(address code) public view returns (bool isIndeed) {
        return code == genesisValidator;
    }

    function isValidator(address code) public view returns (bool isIndeed) {
        return validators[code].isValidator;
    }

    function getValidatorCount() public view returns (uint count) {
        return validatorIndex.length;
    }

    function getValidatorAtIndex(uint index) public view returns (address code, uint deposit) {
        code = validatorIndex[index];
        Validator validator = validators[code];
        deposit = validator.deposits[validator.deposits.length - 1].amount;
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

    function _updateChecksum() private {
        validatorsChecksum = keccak256(validatorIndex);
    }

    function _insertValidator(address code, uint deposit) private {
        Validator sender = validators[code];
        sender.index = validatorIndex.push(code) - 1;
        sender.isValidator = true;
        sender.deposits.push(Deposit({amount:deposit, availableAt: 0}));

        for (uint index = sender.index; index > 0; index--) {
            Validator target = validators[validatorIndex[index - 1]];
            Deposit collateral = target.deposits[target.deposits.length - 1];
            if (deposit <= collateral.amount) {
                break;
            }
            validatorIndex[index] = validatorIndex[index - 1];
            validatorIndex[index - 1] = code; 
            // update indexes
            target.index = index;
            sender.index = index - 1;
        }

        _updateChecksum();
    }

    function setBaseDeposit(uint deposit) public onlyOwner {
        baseDeposit = deposit;
    }

    function _deleteValidator(address account) private {
        Validator validator = validators[account];
        for (uint index = validator.index; index < validatorIndex.length - 1; index++) {
            validatorIndex[index] = validatorIndex[index + 1];
        }
        validatorIndex.length--;

        validator.isValidator = false;
        validator.deposits[validator.deposits.length - 1].availableAt = now + unbondingPeriod;

        _updateChecksum();
    }

    // _deleteSmallestBidder removes the validator with the smallest deposit
    function _deleteSmallestBidder() private {
        _deleteValidator(validatorIndex[validatorIndex.length - 1]);
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

    function getDepositCount() public view returns (uint count) {
        return validators[msg.sender].deposits.length; 
    }

    function getDepositAtIndex(uint index) public view returns (uint amount, uint availableAt) {
        Deposit deposit = validators[msg.sender].deposits[index];
        return (deposit.amount / 1 ether, deposit.availableAt);
    }

    // join registers a new candidate as validator
    function join() public payable onlyNewCandidate onlyWithMinDeposit {
        if (!_hasAvailability()) {
            _deleteSmallestBidder();
        }
        _insertValidator(msg.sender, msg.value);
    }

    // leave deregisters the msg sender from the validator set
    function leave() public onlyValidator {
        _deleteValidator(msg.sender);
    }

    function _removeDeposits(address code, uint index) private {
        if (index == 0) return;

        Validator validator = validators[code];
        uint lo = 0;
        uint hi = index;
        while (hi < validator.deposits.length) {
            validator.deposits[lo] = validator.deposits[hi];
            lo++;
            hi++;
        }
        validator.deposits.length = lo;
    }

    // redeemDeposits transfers locked deposit(s) back the user account if they
    // are past the unbonding period
    function redeemDeposits() public {
        uint refund = 0;
        uint i = 0;
        Deposit[] deposits = validators[msg.sender].deposits;
        
        for (; i < deposits.length && deposits[i].availableAt != 0; i++) {
            if (now < deposits[i].availableAt) {
                // @NOTE (rgeraldes) - no need to iterate further since the 
                // release date (if is different than 0) of the following deposits
                // will always be past than the current one.
                break;
            }
            refund += deposits[i].amount;
        }
        
        _removeDeposits(msg.sender, i);

        if (refund > 0) {
            msg.sender.transfer(refund);
        }
    }

    // fallback function
    function() payable public {}
}