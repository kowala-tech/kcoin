pragma solidity ^0.4.18;

contract Network {
    // Total supply of wei. Must be updated every block and initialized to the correct value.
    uint256 public totalSupplyWei = 1 ether;
    // Reward calculated for the last block. Must be updated every block.
    uint256 public lastBlockReward = 0;
    // Price established by the price oracle for the last block. Must be updated every block.
    uint256 public lastPrice = 0;

    // Voter represents a consensus validator      
    struct Voter {
        uint deposit; // amount at stake
        uint index;
    }

    mapping (address => uint) private genesis; // investors (genesis voters)
    mapping (address => Voter) private voters;
    address[] private voterIndex; 

    // maximum number of voters at one time
    uint public constant MAX_VOTERS = 100;
    // minimum deposit value to participate in the consensus
    uint public minDeposit = 100000;

    //event LogNewVoter(address indexed addr, uint index, uint deposit);
    //event LogDeleteVoter(address indexed addr, uint index);

    function Network() public {
        address investor1 = 0xd6e579085c82329c89fca7a9f012be59028ed53f;
        address investor2 = 0x497dc8a0096cf116e696ba9072516c92383770ed;
        address investor3 = 0xd46d2023a7dde27037de5387b38b17ce1e93e3d2;
        uint investment1 = 100;
        uint investment2 = 100;
        uint investment3 = 100;

        // genesis validators
        genesis[investor1] = investment1;
        genesis[investor2] = investment2;
        genesis[investor3] = investment3;

        // @NOTE(rgeraldes) - be able to vote from the start
        _insertVoter(investor1, investment1);
        _insertVoter(investor2, investment2);
    }

    function isGenesisVoter(address addr) public view returns (bool isIndeed) {
        return genesis[addr] > 0;
    }

    function isVoter(address addr) public view returns (bool isIndeed) {
        if (voterIndex.length == 0) {
            return false; 
        }
        return (voterIndex[voters[addr].index] == addr);
    }

    
    function _insertVoter(address addr, uint deposit) private {
        voters[addr].deposit = deposit;
        voters[addr].index = voterIndex.push(addr) - 1;
    }

    function getVoter(address addr) public view returns (uint deposit, uint index) {
        require(isVoter(addr));
        return (voters[addr].deposit, voters[addr].index);
    }

    function _deleteVoter(address addr) private {
        uint rowToDelete = voters[addr].index;
        address keyToMove = voterIndex[voterIndex.length - 1];
        voterIndex[rowToDelete] = keyToMove;
        voters[keyToMove].index = rowToDelete;
        voterIndex.length--;
    }

    function getVoterCount() public view returns (uint count) {
        return voterIndex.length;
    }

    function getVoterAtIndex(uint index) public view returns (address addr, uint deposit) {
        addr = voterIndex[index];
        deposit = voters[addr].deposit;
    }

    function deposit() public payable {
        require(!isVoter(msg.sender));
        require(msg.value >= minDeposit);
        if (!isGenesisVoter(msg.sender)) {
            require(voterIndex.length < MAX_VOTERS);
        } 
        _insertVoter(msg.sender, msg.value);
    }

    function withdraw() public {
        require(isVoter(msg.sender));

        // withdraw locked money
        msg.sender.transfer(voters[msg.sender].deposit);
        
        _deleteVoter(msg.sender);
    }

    function availability() public view returns (bool available) {
        return voterIndex.length < MAX_VOTERS;
    }

}