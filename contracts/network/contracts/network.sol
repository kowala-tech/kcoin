pragma solidity ^0.4.21;

import "./ownable.sol";

contract Network is Ownable {

    // @NOTE (rgeraldes) - to be confirmed by Hélio Rosa
    /*
    // Total supply of wei. Must be updated every block and initialized to the correct value.
    uint256 public totalSupplyWei = 1 ether;
    // Reward calculated for the last block. Must be updated every block.
    uint256 public lastBlockReward = 0;
    // Price established by the price oracle for the last block. Must be updated every block.
    uint256 public lastPrice = 0;
    */

    // base deposit represents the deposit that a candidate has to do in order to 
    // secure a place in the elections (if there are positions available)
    uint public baseDeposit;       
    
    // base deposit hard limits (safety)
    uint public baseDepositUpperBound;
    uint public baseDepositLowerBound;

    // onlyWithinBaseDepositBounds requires the new deposit value to be within valid bounds
    modifier onlyWithinBaseDepositBounds(uint deposit) {
        require(deposit >= baseDepositLowerBound && deposit <= baseDepositUpperBound);
        _;
    }

    // setBaseDepositLowerBound sets the lower bound of the minimum deposit operation
    function setBaseDepositLowerBound(uint min) public onlyOwner {
        require(min < baseDepositUpperBound && min >= 0);
        baseDepositLowerBound = min;
    }

    // setBaseDepositUpperBound sets the upper bound of the minimum deposit operation
    function setBaseDepositUpperBound(uint max) public onlyOwner {
        require(max > baseDepositLowerBound);
        baseDepositUpperBound = max;
    }

    // setBaseDeposit sets the minimum deposit accepted by the network to join the consensus
    // elections.
    function setBaseDeposit(uint deposit) public onlyOwner onlyWithinBaseDepositBounds(deposit) {
        baseDeposit = deposit;
    }

    // maxValidators represents the maximum number of validators allowed
    // in a consensus election at once
    uint public maxValidators;

    // maxValidators hard limits (safety)
    uint public maxValidatorsUpperBound;
    uint public maxValidatorsLowerBound;

    // onlyWithinMaxValidatorBounds requires the new value to be within valid bounds
    modifier onlyWithinMaxValidatorBounds(uint max) {
        require(max >= maxValidatorsLowerBound && max <= maxValidatorsUpperBound);
        _;
    }

    function setMaxValidatorsLowerBound(uint min) public onlyOwner {
        require(min <= maxValidatorsUpperBound && min >= 0);
        maxValidatorsLowerBound = min;
    }

    function setMaxValidatorsUpperBound(uint max) public onlyOwner {
        require(max >= maxValidatorsLowerBound);
        maxValidatorsUpperBound = max;
    }

    // validatorIndex contains the validator code ordered by the biggest deposit to
    // the smallest deposit.
    address[] validatorIndex;

    // availability returns the number of positions for validator available
    function availability() public view returns (uint available) {
        return maxValidators - validatorIndex.length;
    }

    function setMaxValidators(uint max) public onlyOwner onlyWithinMaxValidatorBounds(max) { 
        if (availability() == 0) {
            uint toRemove = (maxValidators - max);
            for (uint i = 0; i < toRemove; i++) {
                _deleteLastValidator();
            }
        }
        maxValidators = max;
    }

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

    function getDepositCount() public view returns (uint count) {
        return validators[msg.sender].deposits.length; 
    }

    function getDepositAtIndex(uint index) public view returns (uint amount, uint releasedAt) {
        // @TODO (rgeraldes) index bounds
        Deposit deposit = validators[msg.sender].deposits[index];
        return (deposit.amount, deposit.releasedAt);
    }

    mapping (address => Validator) private validators;

    function getValidatorCount() public view returns (uint count) {
        return validatorIndex.length;
    }

    function getValidatorAtIndex(uint index) public view returns (address code, uint deposit) {
        // @TODO (rgeraldes) index bounds
        code = validatorIndex[index];
        Validator validator = validators[code];
        deposit = validator.deposits[validator.deposits.length - 1].amount;
    }

    // @NOTE (rgeraldes) - event filtering is a possibility in the future (issue #140)
    // validatorsChecksum is a representation of the current set of validators
    bytes32 public validatorsChecksum;

    function _updateChecksum() private {
        validatorsChecksum = keccak256(validatorIndex);
    }

    function _insertValidator(address code, uint deposit) private {
        Validator sender = validators[code];
        sender.index = validatorIndex.push(code) - 1;
        sender.isValidator = true;
        sender.deposits.push(Deposit({amount:deposit, releasedAt: 0}));

        // @TODO (rgeraldes) - complete
        /*      
        for (index = val.index; index > 0; index--) {
            if (deposit > oldDeposit) {
                // replace
                validatorIndex[index] = old validator
                validatorIndex[index-1] = code
            }
        }
        */
        
        _updateChecksum();
    }

    // genesis stores the registration code of the genesis validator
    address public genesis;

    function isGenesisValidator(address code) public view returns (bool isIndeed) {
        return code == genesis;
    }

    // unbondingPeriod is a predetermined period of time that coins remain locked
    // starting from the moment a validator leaves the consensus elections
    uint public unbondingPeriod;

    function Network(uint _baseDeposit, address _genesis, uint _maxValidators, uint _unbondingPeriod) public payable {
        require(msg.value >= _baseDeposit);
        require(_maxValidators >= 1);
        require(_unbondingPeriod >= 0);


        baseDeposit = _baseDeposit;
        maxValidators = _maxValidators;
        genesis = _genesis;
        unbondingPeriod = _unbondingPeriod;

        baseDepositLowerBound = baseDeposit / 2;
        baseDepositUpperBound = baseDeposit * 2;
        maxValidatorsLowerBound = maxValidators / 2;
        maxValidatorsUpperBound = maxValidators * 2;

        _insertValidator(_genesis, msg.value);
    }

    // getMinimumDeposit returns the base deposit if there are positions available or
    // the current smallest deposit required if there aren't positions availabe.
    function getMinimumDeposit() public view returns (uint deposit) {
        // there are positions for validator available
        if (availability() > 0) {
            return baseDeposit;
        } else {
            Validator displacedValidator = validators[validatorIndex[validatorIndex.length - 1]];               
            return displacedValidator.deposits[displacedValidator.deposits.length - 1].amount + 1;
        }
    }

    // onlyWithMinDeposit requires a minimum deposit to proceed
    modifier onlyWithMinDeposit {
        require(msg.value >= getMinimumDeposit());
        _;
    }

    function _deleteValidator(address account) private {
        // left shifts
        Validator validator = validators[account];
        for (uint index = validator.index; index < validatorIndex.length - 1; index++) {
            validatorIndex[index] = validatorIndex[index + 1];
        }
        
        // resize array
        validatorIndex.length--;

        // update status
        validator.isValidator = false;
        // update the release date of the current deposit
        validator.deposits[validator.deposits.length - 1].releasedAt = now + unbondingPeriod;
    }

    // _deleteLastValidator removes the validator with the smallest deposit
    function _deleteLastValidator() private {
        _deleteValidator(validatorIndex[validatorIndex.length - 1]);
    }

    // deposit registers a new candidate as validator
    function deposit() public payable onlyWithMinDeposit {
        if (availability() == 0) {
            _deleteLastValidator();
        }
        _insertValidator(msg.sender, msg.value);
    }

    function isValidator(address code) public view returns (bool isIndeed) {
        return validators[code].isValidator;
    }

    // onlyValidator requires the sender to be a validator
    modifier onlyValidator {
        require(isValidator(msg.sender));
        _;
    }

    // leave deregisters the msg sender from the validator set
    function leave() public onlyValidator {
        _deleteValidator(msg.sender);
    }

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

    // withdraw transfer locked deposit(s) back the user account if they
    // are past the unbonding period
    function withdrawFunds() public onlyValidator {
        uint refund = 0;
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
        
        _removeDeposits(validator, i);

        if (refund > 0) {
            msg.sender.transfer(refund);
        }
    }
    
}