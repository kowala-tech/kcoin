pragma solidity 0.4.24;

import "openzeppelin-solidity/contracts/math/Math.sol";

contract System {
    
    uint constant public initialMintedAmount = 42 ether;
    uint constant public initialCap  = 82 ether;
    uint constant public stabilizedPrice = 1 ether;
    uint constant public adjustmentFactor = 10000;
    uint constant public lowSupplyMetric = 1000000 ether;
    uint constant public maxUnderNormalConditions = 1e12;
    uint constant public defaultOracleReward = 1 ether;
    uint constant public oracleDeductionFraction = 4;

    uint public prevCurrencyPrice = 1 ether;
    uint public currencyPrice = 1 ether;
    uint public currencySupply = 0;
    uint public prevMintedAmount = 0;

    function _hasLowSupply() private view returns (bool) {
        return currencySupply < lowSupplyMetric;
    }

    function _cap() private view returns (uint amount) {
        return ((block.number > 1) && !_hasLowSupply) ? currencySupply/10000 : initialCap;
    }

    function mintedAmount() public view returns (uint) {
        if (block.number == 1) return initialMintedAmount;

        uint adjustment = prevMintedAmount/adjustmentFactor;
        if ((currencyPrice > prevCurrencyPrice) && (prevCurrencyPrice > stabilizedPrice)) {
            return Math.min256(prevMintedAmount + adjustment, _cap());
        }
        return Math.max256(prevMintedAmount - adjustment, maxUnderNormalConditions);
    }

    function oracleDeduction(uint mintedAmount) public view returns (uint) {
        return oracleDeductionFraction * mintedAmount / 100;
    }

    function oracleReward() public view returns (uint) {
        return Math.min256(defaultOracleReward, this.balance);
    }
}