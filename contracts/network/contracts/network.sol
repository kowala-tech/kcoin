pragma solidity ^0.4.18;

import "./ownable.sol";

contract Network is Ownable {
    // minimum deposit that a candidate has to do in order to 
    // secure a place in the elections (if there are positions available)
    uint public minDeposit;       
    
    // minimum deposit hard limits (safety)
    uint public minDepositUpperBound;
    uint public minDepositLowerBound;

    // @NOTE (rgeraldes) - this field is used to know the address of the 
    // genesis validator (does not need to make a deposit)
    address public genesis;

    // onlyWithMinDeposit requires a minimum deposit to proceed
    modifier onlyWithMinDeposit {
        require(msg.value >= minDeposit);
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
    function setMinDeposit(uint deposit) public onlyOwner {
        require(deposit >= minDepositLowerBound && deposit <= minDepositUpperBound);
        minDeposit = deposit;
    }

    function Network(uint _minDeposit, address _genesis) public {
        // initial minimum deposit values
        minDeposit = _minDeposit;
        minDepositLowerBound = minDeposit / 2;
        minDepositUpperBound = minDeposit * 2;

        // initial validator
        genesis = _genesis;
        // @TODO (register the genesis validator)
    }
}

    /*
    // Total supply of wei. Must be updated every block and initialized to the correct value.
    uint256 public totalSupplyWei = 1 ether;
    // Reward calculated for the last block. Must be updated every block.
    uint256 public lastBlockReward = 0;
    // Price established by the price oracle for the last block. Must be updated every block.
    uint256 public lastPrice = 0;
    

    // unbondingPeriod is a predetermined period of time that coins remain locked
    // from the moment a validator leaves the consensus elections
    uint public unbondingPeriod = 4 weeks;

    // @NOTE (rgeraldes) - easy and efficient way to identify changes in the contract
    // current checksum of the validatorsChecksum
    bytes32 public validatorsChecksum;

    address public genesis;

    // Validator represents a consensus validator      
    struct Validator {
        uint index;
        bool isValidator;

        // @NOTE (rgeraldes) - users can have more than one collateral
        // Example: user leaves and re-enters the election. At this point
        // the initial collateral will have a release date and the validator 
        // will have a new collateral for the current election.
        Collateral[] collaterals; 
    }

    // Collateral - staking tokens bonded
    struct Collateral {
        uint deposit;
        uint releasedAt;
    }

    mapping (address => Validator) private validators;
    address[] private validatorIndex; 

    // maxValidators - hard limits 
    uint public maxValidatorsUpperBound = 500;
    uint public maxValidatorsLowerBound = 100;

    // maximum number of validators at one time
    uint public maxValidators = 100;

    


    function Network() public {
        address investor1 = 0xd6e579085c82329c89fca7a9f012be59028ed53f;
        uint investment1 = 100;

        // genesis validators
        genesis[investor1] = investment1;

        // @NOTE(rgeraldes) - be able to vote from the start
        _registerValidator(investor1, investment1);
    }

    // modifiers


    // onlyValidator requires the sender to be a validator
    modifier onlyValidator {
        require(isValidator(msg.sender));
        _;
    }

    

    

    

    function _updateCollateral(address account, uint deposit) private {

    }

    function _updateValidatorsChecksum() private {
        validatorsChecksum = keccak256(validatorIndex);
    }

    
    // _deleteVoter removes the record from the voterIndex
    // @NOTE (rgeraldes) - the record in the voters map still exists - contains info on the locked collateral.
    function _deleteVoter(address addr) private {
        uint rowToDelete = voters[addr].index;
        address keyToMove = voterIndex[voterIndex.length - 1];
        voterIndex[rowToDelete] = keyToMove;
        voters[keyToMove].index = rowToDelete;
        voterIndex.length--;
        voters[addr].isVoter = false;
    }

    // _insertVoter adds a new voter record to the voterIndex array and voters map
    function _insertVoter(address account, uint deposit) private {
        voters[addr].collateral.push();
        voters[addr].index = voterIndex.push(addr) - 1;
        voters[addr].isVoter = true;
    }

    function _registerVoter(address account, uint deposit) private onlyWithMinDeposit {
        _insertVoter(account, deposit);
        _updateVotersChecksum();
    }

    function _unbondCollateral(address account) private {
        // @NOTE (rgeraldes) - the current active collateral is the last one.
        // Note that the validator can have multiple collaterals since he could
        // have left the 
        voters[account].collaterals[0].releasedAt = now + unbondingPeriod;
    }

    function _releaseCollateral(address account, uint index) private {
        account.transfer(voters[account].deposit);
    }

    function _deregisterVoter(address account) private {
        _deleteVoter(account);
        _unbondCollateral(account);
        _updateVotersChecksum();
    }   


    // functions restricted to the owner

    function setMaxVotersUpperBound(uint max) public onlyOwner {
        require(max >= maxVotersLowerBound);
        maxVotersUpperBound = max;
    }

    function setMaxVotersLowerBound(uint min) public onlyOwner {
        require(max <= maxVotersUpperBound);
        maxVotersLowerBound = min;
    }

    function setMaxVoters(uint max) public onlyOwner(owner) {
        require(max >= maxVotersLowerBound && max <= maxVotersUpperBound);

        if (max > maxVoters) {
            maxVoters = max;
        } else {
            // @NOTE (rgeraldes) - if the new max if smaller than the current one
            // we will have to remove (old - new) validators from the election if 
            // all the validator positions are occupied 
            
        }
    }

    

    // transactions

    // deposit increments the stake of a voter or registers a new candidate
    function deposit() public payable {
        if (isVoter(msg.sender)) {
            // @NOTE (rgeraldes) - validator wants to increase its stake
            _updateCollateral(msg.sender, msg.value);
        } else {
            // @NOTE (rgeraldes) - new candidate
            require(msg.value >= minDeposit && (voterIndex.length < maxVoters || voters));
            _registerVoter(msg.sender, msg.value);
        }
    }

    // leave deregisters the voter 
    function leave() public onlyVoter {
        _deregisterVoter(msg.sender);
    }

    // withdraw refunds the user account with a collateral(s).
    function withdraw() public onlyVoter {
        collaterals = voters[msg.sender].collaterals;
        for (uint i = 0; i < collaterals.length && collaterals[i].releasedAt != 0; i++) {
            if (now < collaterals[i].releasedAt) {
                // @NOTE (rgeraldes) - no need to iterate further since the 
                // release date of the following collaterals will always be
                // bigger than the current one.
                break;
            }
            _releaseCollateral(msg.sender, i);
        }
    }

    // helpers

    // availability provides information on the current number of positions available
    function availability() public view returns (bool available) {
        return voterIndex.length < maxVoters;
    }

    function isGenesisVoter(address account) public view returns (bool isIndeed) {
        return account = genesis;
    }

    function isVoter(address addr) public view returns (bool isIndeed) {
        return voters[addr].isVoter;
    }

    function getVoter(address addr) public view returns (uint deposit, uint index) {
        require(isVoter(addr));
        return (voters[addr].deposit, voters[addr].index);
    }

    function getVoterCount() public view returns (uint count) {
        return voterIndex.length;
    }

    function getVoterAtIndex(uint index) public view returns (address addr, uint deposit) {
        addr = voterIndex[index];
        deposit = voters[addr].deposit;
    } */