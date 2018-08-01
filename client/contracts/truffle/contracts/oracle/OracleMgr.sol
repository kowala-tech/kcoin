pragma solidity 0.4.24;

import "openzeppelin-solidity/contracts/lifecycle/Pausable.sol";
import "../consensus/mgr/ValidatorMgr.sol";
/**
* @title Oracle Manager contract
*/
contract OracleMgr is Pausable {

    uint public baseDeposit;       
    uint public maxNumOracles;
    uint public freezePeriod;
    uint public syncFrequency;
    uint public updatePeriod;
    uint public price;
    ValidatorMgr validatorMgr;
    bytes4 sig = bytes4(keccak256("isSuperNode(address)"));

    struct Deposit {
        uint amount;
        uint availableAt;
    }

    struct Oracle {
        uint index;
        bool isOracle;
        Deposit[] deposits; 
    }
    
    mapping (address => Oracle) private oracleRegistry;
    
    // oraclePool contains the oracle identity ordered by the biggest deposit to
    // the smallest deposit.
    address[] private oraclePool;

    modifier onlyWithMinDeposit {
        require(msg.value >= getMinimumDeposit());
        _;
    }

    modifier onlyOracle {
        require(isOracle(msg.sender));
        _;
    }

    modifier onlyNewCandidate {
        require(!isOracle(msg.sender));
        _;
    }

    modifier onlySuperNode {
        require(validatorMgr.isSuperNode(msg.sender));
        _;
    }

    modifier onlyValidPrice(uint _price) {
        require(_price > 0);
        _;
    }

    /**
     * Constructor.
     * @param _initialPrice Initial Price.
     * @param _baseDeposit base deposit for Oracle.
     * @param _maxNumOracles Maximum numbers of Oracles.
     * @param _freezePeriod Freeze period for Oracle's deposit.
     * @param _syncFrequency Synchronize frequency for Oracles.
     * @param _updatePeriod Update period.
     */
    function OracleMgr(
        uint _initialPrice, 
        uint _baseDeposit,
        uint _maxNumOracles,
        uint _freezePeriod,
        uint _syncFrequency,
        uint _updatePeriod,
        address _validatorMgrAddr) 
    public {
        require(_initialPrice > 0);
        require(_maxNumOracles > 0);
        require(_syncFrequency >= 0);

        // sync enabled
        if (_syncFrequency > 0) {
            require(_updatePeriod > 0 && _updatePeriod <= _syncFrequency);
        }
        
        price = _initialPrice;
        baseDeposit = _baseDeposit;
        maxNumOracles = _maxNumOracles;
        freezePeriod = _freezePeriod * 1 days;
        syncFrequency = _syncFrequency;
        updatePeriod = _updatePeriod;
        validatorMgr = ValidatorMgr(_validatorMgrAddr);
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
    function _hasAvailability() public view returns (bool available) {
        return (maxNumOracles - oraclePool.length) > 0;
    }

    /**
     * @dev Deletes given oracle
     * @param identity Address of an Oracle.
     */
    function _deleteOracle(address identity) private {
        Oracle oracle = oracleRegistry[identity];
        for (uint index = oracle.index; index < oraclePool.length - 1; index++) {
            oraclePool[index] = oraclePool[index + 1];
        }
        oraclePool.length--;

        oracle.isOracle = false;
        oracle.deposits[oracle.deposits.length - 1].availableAt = now + freezePeriod;
    }

    /**
     * @dev removes the oracle with the smallest deposit
     */
    function _deleteSmallestBidder() private {
        _deleteOracle(oraclePool[oraclePool.length - 1]);
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
        oracle.deposits.push(Deposit({amount: deposit, availableAt: 0}));

        for (uint index = oracle.index; index > 0; index--) {
            Oracle target = oracleRegistry[oraclePool[index - 1]];
            Deposit collateral = target.deposits[target.deposits.length - 1];
            if (deposit <= collateral.amount) {
                break;
            }
            oraclePool[index] = oraclePool[index - 1];
            oraclePool[index - 1] = identity; 
            // update indexes
            target.index = index;
            oracle.index = index - 1;
        }
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
    function getOracleAtIndex(uint index) public view returns (address code, uint deposit) {
        code = oraclePool[index];
        Oracle oracle = oracleRegistry[code];
        deposit = oracle.deposits[oracle.deposits.length - 1].amount;
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
            Oracle smallestBidder = oracleRegistry[oraclePool[oraclePool.length - 1]];               
            return smallestBidder.deposits[smallestBidder.deposits.length - 1].amount + 1;
        }
    }
    
    /**
     * @dev Get deposit count
     */
    function getDepositCount() public view returns (uint count) {
        return oracleRegistry[msg.sender].deposits.length; 
    }
    /**
     * @dev Get deposit at given index
     * @param index index of an Oracle
     */
    function getDepositAtIndex(uint index) public view returns (uint amount, uint availableAt) {
        Deposit deposit = oracleRegistry[msg.sender].deposits[index];
        return (deposit.amount, deposit.availableAt);
    }

    /**
     * @dev Registers a new candidate as oracle
     */
    function registerOracle() public payable whenNotPaused onlyNewCandidate onlyWithMinDeposit onlySuperNode {
        if (!_hasAvailability()) {
            _deleteSmallestBidder();
        }
        _insertOracle(msg.sender, msg.value);
    }

    /**
     * @dev Deregisters the msg sender from the oracle set
     */
    function deregisterOracle() public whenNotPaused onlyOracle {
        _deleteOracle(msg.sender);
    }

    /**
     * @dev Remove deposit from an Oracle
     * @param identity Address of an Oracle.
     * @param index index of an Oracle
     */
    function _removeDeposits(address identity, uint index) private {
        if (index == 0) return;

        Oracle oracle = oracleRegistry[identity];
        uint lo = 0;
        uint hi = index;
        while (hi < oracle.deposits.length) {
            oracle.deposits[lo] = oracle.deposits[hi];
            lo++;
            hi++;
        }
        oracle.deposits.length = lo;
    }

    /**
     * @dev transfers locked deposit(s) back the user account if they are past the freeze period
     */
    function releaseDeposits() public whenNotPaused {
        uint refund = 0;
        uint i = 0;
        Deposit[] deposits = oracleRegistry[msg.sender].deposits;
        
        for (; i < deposits.length && deposits[i].availableAt != 0; i++) {
            if (now < deposits[i].availableAt) {
                break;
            }
            refund += deposits[i].amount;
        }
        
        _removeDeposits(msg.sender, i);

        if (refund > 0) {
            msg.sender.transfer(refund);
        }
    }
    
    /**
     * @dev Adds price
     * @param _price price
     */
    function addPrice(uint _price) public whenNotPaused onlyOracle onlyValidPrice(_price) {
        price = _price;
    }
}