pragma solidity 0.4.24;

import "openzeppelin-solidity/contracts/lifecycle/Pausable.sol";
import "./Consensus.sol";

contract OracleMgr is Pausable {
     
    uint public maxNumOracles;
    uint public syncFrequency;
    uint public updatePeriod;
    Consensus consensus;

    struct OraclePrice {
        uint price;
        address oracle;
    }

    struct Oracle {
        uint index;
        bool isOracle;
        bool hasSubmittedPrice;
    }
    
    mapping (address => Oracle) private oracleRegistry;
    
    // oraclePool contains the oracle identity ordered by the biggest deposit to
    // the smallest deposit.
    address[] private oraclePool;

    OraclePrice[] private prices;

    modifier onlyOracle {
        require(isOracle(msg.sender));
        _;
    }

    modifier onlyNewCandidate {
        require(!isOracle(msg.sender));
        _;
    }

    modifier onlySuperNode {
        require(consensus.isSuperNode(msg.sender));
        _;
    }

    modifier onlyOnce {
        require(!oracleRegistry[msg.sender].hasSubmittedPrice);
        _;
    }

    function OracleMgr(
        uint _maxNumOracles,
        uint _syncFrequency,
        uint _updatePeriod,
        address _consensusAddr) 
    public {
        require(_maxNumOracles > 0);

        // sync enabled
        if (_syncFrequency > 0) {
            require(_updatePeriod > 0 && _updatePeriod <= _syncFrequency);
        }
        
        maxNumOracles = _maxNumOracles;
        syncFrequency = _syncFrequency;
        updatePeriod = _updatePeriod;
        consensus = Consensus(_consensusAddr);
    }

    function isOracle(address identity) public view returns (bool isIndeed) {
        return oracleRegistry[identity].isOracle;
    }

    function _hasAvailability() private view returns (bool available) {
        return (maxNumOracles - oraclePool.length) > 0;
    }

    function _deleteOracle(address identity) private {
        Oracle oracle = oracleRegistry[identity];
        uint rowToDelete = oracle.index;
        delete oracleRegistry[identity];
        
        // replace the deprecated record with the last element
        address keyToMove = oraclePool[oraclePool.length-1];
        oraclePool[rowToDelete] = keyToMove;
        oracleRegistry[keyToMove].index = rowToDelete;
        oraclePool.length--;       
    }

    function _insertOracle(address identity, uint deposit) private {
        Oracle oracle = oracleRegistry[identity];
        oracle.index = oraclePool.push(identity) - 1;
        oracle.isOracle = true;
    }

    function getOracleCount() public view returns (uint count) {
        return oraclePool.length;
    }

    function getOracleAtIndex(uint index) public view returns (address code) {
        code = oraclePool[index];
        Oracle oracle = oracleRegistry[code];
    }

    function getPriceCount() public view returns (uint count) {
        return prices.length;
    }

    function getPriceAtIndex(uint index) public view returns (uint price, address oracle) {
        OraclePrice oraclePrice = prices[index];
        price = oraclePrice.price;
        oracle = oraclePrice.oracle;
    }

    // registerOracle registers a new candidate as oracle
    function registerOracle() public payable whenNotPaused onlyNewCandidate onlySuperNode {
        require(_hasAvailability());
        _insertOracle(msg.sender, msg.value);
    }

    // deregisterOracle deregisters the msg sender from the oracle set
    function deregisterOracle() public whenNotPaused onlyOracle {
        _deleteOracle(msg.sender);
    }

    function submitPrice(uint _price) public whenNotPaused onlyOracle onlyOnce {
        oracleRegistry[msg.sender].hasSubmittedPrice = true;
        prices.push(OraclePrice({price: _price, oracle: msg.sender}));
    }
}