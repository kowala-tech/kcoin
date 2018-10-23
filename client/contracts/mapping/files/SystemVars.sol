pragma solidity 0.4.24;

import "openzeppelin-solidity/contracts/math/Math.sol";
import "zos-lib/contracts/migrations/Initializable.sol";

/**
 * @title System Variables
 */
contract SystemVars is Initializable{
    
    uint constant INITIAL_MINTED_AMOUNT = 42 ether;
    uint constant INITIAL_CAP  = 82 ether;
    uint constant STABILIZED_PRICE = 1 ether;
    uint constant ADJUSTMENT_FACTOR = 10000;
    uint constant LOW_SUPPLY_METRIC = 1000000 ether;
    uint constant MAX_UNDER_NORMAL_CONDITIONS = 1e12;
    uint constant DEFAULT_ORACLE_REWARD = 1 ether;
    uint constant ORACLE_DEDUCTION_FRACTION = 4;

    uint public prevCurrencyPrice;
    uint public currencyPrice;
    uint public currencySupply;
    uint public mintedReward;

    /**
     * Constructor.
     * @param _initialPrice initial price for the system's currency
     * @param _initialSupply minted amount on the genesis block
     */
    function SystemVars(uint _initialPrice, uint _initialSupply) public {
        prevCurrencyPrice = _initialPrice;
        currencyPrice = _initialPrice;
        mintedReward = _initialSupply;
        currencySupply = _initialSupply;
    }

    /**
     * initialize function for Proxy Pattern.
     * @param _initialPrice initial price for the system's currency
     * @param _initialSupply minted amount on the genesis block
     */
    function initialize(uint _initialPrice, uint _initialSupply) isInitializer public {
        prevCurrencyPrice = _initialPrice;
        currencyPrice = _initialPrice;
        mintedReward = _initialSupply;
        currencySupply = _initialSupply;
    }

    function _hasEnoughSupply() private view returns (bool) {
        return currencySupply >= LOW_SUPPLY_METRIC;
    }

    function _cap() private view returns (uint amount) {
        return (((block.number + 1) > 1) && _hasEnoughSupply()) ? currencySupply/10000 : INITIAL_CAP;
    }

    /**
     * @dev Get the current system's currency price
     */
    function price() public view returns (uint price) {
        return currencyPrice;
    }

    /**
     * @dev Get the amount of coins that should be minted
     */
    function mintedAmount() public view returns (uint) {
        if ((block.number + 1) == 1) return INITIAL_MINTED_AMOUNT;

        uint adjustment = mintedReward/ADJUSTMENT_FACTOR;
        if ((currencyPrice > prevCurrencyPrice) && (prevCurrencyPrice > STABILIZED_PRICE)) {
            return Math.min256(mintedReward + adjustment, _cap());
        }
        return Math.max256(mintedReward - adjustment, MAX_UNDER_NORMAL_CONDITIONS);
    }

    /**
     * @dev Get the oracle deduction
     * @param mintedAmount the minted amount for the current block.
     */
    function oracleDeduction(uint mintedAmount) public view returns (uint) {
        return ORACLE_DEDUCTION_FRACTION * mintedAmount / 100;
    }

    /**
     * @dev Get the oracle reward
     */
    function oracleReward() public view returns (uint) {
        return Math.min256(DEFAULT_ORACLE_REWARD, this.balance);
    }

    function revertFunc() public {
        revert();
    }
}