pragma solidity 0.4.24;

import "openzeppelin-solidity/contracts/lifecycle/Pausable.sol";
import "./Consensus.sol";
import "../kns/DomainResolver.sol";
import "zos-lib/contracts/migrations/Initializable.sol";
import {NameHash} from "../utils/NameHash.sol";

/**
* @title Oracle Manager contract
*/
contract OracleMgr is Pausable, Initializable {
     
    uint public maxNumOracles;
    uint public syncFrequency;
    uint public updatePeriod;
    uint public price;
    DomainResolver public knsResolver;
    bytes32 nodeNamehash;
    bytes4 sig = bytes4(keccak256("isSuperNode(address)"));

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
        require(Consensus(knsResolver.addr(nodeNamehash)).isSuperNode(msg.sender));
        _;
    }

    modifier onlyOnce {
        require(!oracleRegistry[msg.sender].hasSubmittedPrice);
        _;
    }

    /**
     * Constructor.
     * @param _maxNumOracles Maximum numbers of Oracles.
     * @param _syncFrequency Synchronize frequency for Oracles.
     * @param _updatePeriod Update period.
     * @param _resolverAddr Address of KNS Resolver.
     */
    function OracleMgr(
        uint _maxNumOracles,
        uint _syncFrequency,
        uint _updatePeriod,
        address _resolverAddr) 
    public {
        require(_maxNumOracles > 0);

        // sync enabled
        if (_syncFrequency > 0) {
            require(_updatePeriod > 0 && _updatePeriod <= _syncFrequency);
        }
        
        maxNumOracles = _maxNumOracles;
        syncFrequency = _syncFrequency;
        updatePeriod = _updatePeriod;
        knsResolver = DomainResolver(_resolverAddr);
        nodeNamehash = NameHash.namehash("validatormgr.kowala");
    }

    /**
     * initialize function for Proxy Pattern.
     * @param _maxNumOracles Maximum numbers of Oracles.
     * @param _syncFrequency Synchronize frequency for Oracles.
     * @param _updatePeriod Update period.
     * @param _resolverAddr Address of KNS Resolver.
     */
    function initialize (
        uint _maxNumOracles,
        uint _syncFrequency,
        uint _updatePeriod,
        address _resolverAddr)
    isInitializer 
    public {
        require(_maxNumOracles > 0);

        // sync enabled
        if (_syncFrequency > 0) {
            require(_updatePeriod > 0 && _updatePeriod <= _syncFrequency);
        }
        
        maxNumOracles = _maxNumOracles;
        syncFrequency = _syncFrequency;
        updatePeriod = _updatePeriod;
        knsResolver = DomainResolver(_resolverAddr);
        nodeNamehash = NameHash.namehash("validatormgr.kowala");
    }

    /**
     * @dev Checks if given address is Oracle
     * @param identity Address of an Oracle.
     */
    function isOracle(address identity) public view returns (bool isIndeed) {
        return oracleRegistry[identity].isOracle;
    }

    /**
     * @dev Checks availability of OraclePool
     */
    function _hasAvailability() private view returns (bool available) {
        return (maxNumOracles - oraclePool.length) > 0;
    }

    /**
     * @dev Deletes given oracle
     * @param identity Address of an Oracle.
     */
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

    /**
     * @dev Inserts oracle
     * @param identity Address of an Oracle.
     * @param deposit Deposit ammount
     */
    function _insertOracle(address identity, uint deposit) private {
        Oracle oracle = oracleRegistry[identity];
        oracle.index = oraclePool.push(identity) - 1;
        oracle.isOracle = true;
    }

    /**
     * @dev Get Oracle count
     */
    function getOracleCount() public view returns (uint count) {
        return oraclePool.length;
    }

    /**
     * @dev Get Oracle information
     * @param index index of an Oracle to check.
     */
    function getOracleAtIndex(uint index) public view returns (address code) {
        code = oraclePool[index];
        Oracle oracle = oracleRegistry[code];
    }

    /**
     * @dev Get submissions count
     */
    function getPriceCount() public view returns (uint count) {
        return prices.length;
    }

    /**
     * @dev Get submissions information
     * @param index index of a submission to check.
     */
    function getPriceAtIndex(uint index) public view returns (uint price, address oracle) {
        OraclePrice oraclePrice = prices[index];
        price = oraclePrice.price;
        oracle = oraclePrice.oracle;
    }

    /**
     * @dev Registers a new candidate as oracle
     */
    function registerOracle() public payable whenNotPaused onlyNewCandidate onlySuperNode {
        require(_hasAvailability());
        _insertOracle(msg.sender, msg.value);
    }

    /**
     * @dev Deregisters the msg sender from the oracle set
     */
    function deregisterOracle() public whenNotPaused onlyOracle {
        _deleteOracle(msg.sender);
    }

    /**
     * @dev Adds price
     * @param _price price
     */
    function submitPrice(uint _price) public whenNotPaused onlyOracle onlyOnce {
        oracleRegistry[msg.sender].hasSubmittedPrice = true;
        prices.push(OraclePrice({price: _price, oracle: msg.sender}));
    }
}