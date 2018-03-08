pragma solidity ^0.4.18;

import "./ownable.sol";

contract Network is Ownable {

    // @NOTE (rgeraldes) - to be confirmed by Hélio
    /*
    // Total supply of wei. Must be updated every block and initialized to the correct value.
    uint256 public totalSupplyWei = 1 ether;
    // Reward calculated for the last block. Must be updated every block.
    uint256 public lastBlockReward = 0;
    // Price established by the price oracle for the last block. Must be updated every block.
    uint256 public lastPrice = 0;
    */

    // minimum deposit that a candidate has to do in order to 
    // secure a place in the elections (if there are positions available)
    uint public minDeposit;       
    
    // minimum deposit hard limits (safety)
    uint public minDepositUpperBound;
    uint public minDepositLowerBound;

    // @NOTE (rgeraldes) - this field is used to know the address of the 
    // genesis validator (does not need to make a deposit)
    address public genesis;

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

    // validators contains the information of the validators
    mapping (address => Validator) private validators;

    // validatorIndex is used in order to have an ordered insert
    address[] private validatorIndex; 

    // validatorsChecksum is a representation of the current set of validators.
    bytes32 public validatorsChecksum;

    // maximum number of validators at one time
    uint public maxValidators;

    // maxValidators hard limits (safety)
    uint public maxValidatorsUpperBound;
    uint public maxValidatorsLowerBound;

    // unbondingPeriod is a predetermined period of time that coins remain locked
    // starting from the moment a validator leaves the consensus elections
    uint public unbondingPeriod;

    // onlyWithMinDeposit requires a minimum deposit to proceed
    modifier onlyWithMinDeposit {
        require(msg.value >= minDeposit);
        _;
    } 

    // onlyValidator requires the sender to be a validator
    modifier onlyValidator {
        require(isValidator(msg.sender));
        _;
    }

    // onlyWithinMaxValidatorBounds requires the new value to be within valid bounds
    modifier onlyWithinMaxValidatorBounds(uint max) {
        require(max >= maxValidatorsLowerBound && max <= maxValidatorsUpperBound);
        _;
    }

    // onlyWithinMinDepositBounds requires the new value to be within valid bounds
    modifier onlyWithinMinDepositBounds(uint deposit) {
        require(deposit >= minDepositLowerBound && deposit <= minDepositUpperBound);
        _;
    }


    // setMinDepositUpperBound sets the upper bound of the minimum deposit operation
    function setMinDepositUpperBound(uint max) public onlyOwner {
        require(max >= minDepositLowerBound);
        minDepositUpperBound = max;
    }

    // setMinDepositLowerBound sets the lower bound of the minimum deposit operation
    function setMinDepositLowerBound(uint min) public onlyOwner {
        require(min <= minDepositUpperBound);
        minDepositLowerBound = min;
    }

    // setMinDeposit sets the minimum deposit accepted by the network to join the consensus
    // elections.
    function setMinDeposit(uint deposit) public onlyOwner onlyWithinMinDepositBounds(deposit) {
        minDeposit = deposit;
    }

    function Network(uint _minDeposit, address _genesis, uint _maxValidators, uint _unbondingPeriod) public {
        minDeposit = _minDeposit;
        minDepositLowerBound = minDeposit / 2;
        minDepositUpperBound = minDeposit * 2;
        maxValidators = _maxValidators;
        maxValidatorsLowerBound = maxValidators / 2;
        maxValidatorsUpperBound = maxValidators * 2;
        genesis = _genesis;
        unbondingPeriod = _unbondingPeriod;
        // @TODO (register the genesis validator)
        //  _registerValidator(investor1, investment1);
    }

    function isGenesisValidator(address code) public view returns (bool isIndeed) {
        return code == genesis;
    }

    function isValidator(address code) public view returns (bool isIndeed) {
        return validators[code].isValidator;
    }

    function getValidatorCount() public view returns (uint count) {
        return validatorIndex.length;
    }

    // availability returns the number of positions available
    function availability() public view returns (uint available) {
        return maxValidators - validatorIndex.length;
    }

    function setMaxValidatorsUpperBound(uint max) public onlyOwner {
        require(max >= maxValidatorsLowerBound);
        maxValidatorsUpperBound = max;
    }

    function setMaxValidatorsLowerBound(uint min) public onlyOwner {
        require(min <= maxValidatorsUpperBound);
        maxValidatorsLowerBound = min;
    }

    function setMaxValidators(uint max) public onlyOwner onlyWithinMaxValidatorBounds(max) { 
        if (max < maxValidators) {
            // @NOTE (rgeraldes) - if the new max if smaller than the current one
            // we will have to remove (maxValidators - max) validators from the election if 
            // all the validator positions are occupied. We start by removing the validators
            // which appear in the last positions as it represents the validators with 
            // the smallest stake at play. 
            uint toRemove = (maxValidators - max);
            for (uint i = 0; i < toRemove; i++) {
                _deregisterLastValidator();
            }
        }
        maxValidators = max;
    }

    // deposit increments the deposit of a validator/registers a new candidate
    function deposit() public payable {
        if (isValidator(msg.sender)) {
            _updateDepositAmount();
        } else {
            _newCandidate();
        }
    }

    function _updateDepositAmount() private {
        Deposit deposit = validators[msg.sender].collaterals[validatorIndex.length - 1];
        deposit.amount += msg.value;
    }

    function _newCandidate() private onlyWithMinDeposit {
        // there are no positions available
        if (validatorIndex.length == maxValidators) {
            // value needs to be bigger than the smallest deposit
            // the smallest deposit corresponds to the current deposit 
            // of the last validator in validatorIndex array
            Validator lastValidator = validators[validatorIndex.length];               
            uint smallestDeposit = lastValidator.Deposits[lastValidator.Deposits.length - 1];
            require(msg.value > smallestDeposit);
            _deregisterValidator();
        }
        _registerCandidate(msg.sender, msg.value);
    }


    function _deregisterLastValidator() private {
        lastValidator = validatorIndex[validatorIndex.length - 1];
        _deregisterValidator(lastValidator);
    }

    // leave deregisters the validator
    function leave() public onlyValidator {
        _deregisterValidator(msg.sender);
    }

    // withdraw returns the locked deposit(s) (if they are past the unbonding period) 
    // to the user account.
    function withdraw() public onlyValidator {
        Validator validator = validators[msg.sender];

        for (uint i = 0; i < validator.deposits.length && validator.deposits[i].releasedAt != 0;) {
            if (now < validator.deposits[i].releasedAt) {
                // @NOTE (rgeraldes) - no need to iterate further since the 
                // release date (if is different than 0) of the following deposits
                // will always be bigger than the current one.
                break;
            } 
            _releaseDeposit(validator);
        }
    }

    function getValidator(address account) public view returns (uint deposit, uint index) {
        require(isValidator(account));
        return (voters[addr].deposit, voters[addr].index);
    }

    
    function getValidatorAtIndex(uint index) public view returns (address addr, uint deposit) {
        addr = validatorIndex[index];
        deposit = validators[addr].deposit;
    }

    function _updateChecksum() private {
        validatorsChecksum = keccak256(validatorIndex);
    }


    function _deregisterLastValidator(address code) private {
        _deleteValidator(code);
        _setDepositReleaseDate(code);
        _updateVotersChecksum();
    }   

    function _registerCandidate(address account, uint deposit) {
        _insertValidator(account, deposit);
        _updateVotersChecksum();
    }

    function _setDepositReleaseDate(address account) private {
        // @NOTE (rgeraldes) - the current active collateral is the last one.
        // Note that the validator can have multiple collaterals since he could
        // have left the 
        validators[account].deposits[0].releasedAt = now + unbondingPeriod;
    }

    function _transfer(address account, uint index) private {
        account.transfer(validators[account].deposit);
    }

    function _deleteValidator(address account) private {
        uint rowToDelete = validators[account].index;
        address keyToMove = validatorIndex[validatorIndex.length - 1];
        validatorIndex[rowToDelete] = keyToMove;
        validators[keyToMove].index = rowToDelete;
        validatorIndex.length--;
        validators[account].isVoter = false;
    }

    // _insertVoter adds a new voter record to the voterIndex array and voters map
    function _insertValidator(address account, uint deposit) private {
        voters[addr].collateral.push();
        voters[addr].index = voterIndex.push(addr) - 1;
        voters[addr].isVoter = true;
    }
    
    /*
    function remove(uint index)  returns(uint[]) {
        if (index < array.length) return;

        for (uint i = index; i<array.length-1; i++){
            array[i] = array[i+1];
        }
        array.length--;
        return array;
    }
    */

}