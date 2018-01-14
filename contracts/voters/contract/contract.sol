pragma solidity ^0.4.19;

contract VoterRegistry {
    // maximum number of voters at one time
    uint maxVoters = 3;

    // current number of voters at one time
    uint count = 0;

    // minimum deposit value to participate in the consensus
    uint minDeposit = 0;

    // Voter represents a consensus validator
    struct Voter {
        uint deposit;       // amount at stake
        bool isVoter;       // helper
    }

    mapping (address => Voter) public voters;

    function VoterRegistry() public {
        // add genesis validators?
    }

    function minimumDeposit() public view returns (uint) { return minDeposit; }
    function availability() public view returns (bool) { return count < maxVoters; }
    function isVoter(address addr) public constant returns (bool) { return voters[addr].isVoter; }

    function _newVoter(address _coinbase, uint _amount) private {
        count++;
        voters[_coinbase] = Voter(_amount, true);
    }

    function _deleteVoter(address _coinbase) private {
        count--;
        voters[msg.sender] = Voter(0, false);
    }

    function deposit() public payable {
        // revert call if the sender is already a voter
        require(!isVoter(msg.sender));

        // revert the call if there are no spots available
        require(count < maxVoters);

        // if the deposit is not higher, send the money back
        require(msg.value >= minDeposit);

       _newVoter(msg.sender, msg.value);
    }

    function withdraw() public {
        address beneficiary = msg.sender;

        // revert call if the sender is not a voter
        require(isVoter(beneficiary));

        // withdraw locked money
        uint amount = voters[beneficiary].deposit;
        
        _deleteVoter(beneficiary);
        
        beneficiary.transfer(amount);        
    }
}