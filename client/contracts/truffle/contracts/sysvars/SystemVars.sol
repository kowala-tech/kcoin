pragma solidity 0.4.24;

import "openzeppelin-solidity/contracts/math/Math.sol";

contract SystemVars {
    
    uint constant public initialMintedAmount = 42 ether;
    uint constant public initialCap  = 82 ether;
    uint constant public stabilizedPrice = 1 ether;
    uint constant public adjustmentFactor = 10000;
    uint constant public lowSupplyMetric = 1000000 ether;
    uint constant public maxUnderNormalConditions = 1e12;
    uint constant public defaultOracleReward = 1 ether;
    uint constant public oracleDeductionFraction = 4;

    uint public prevCurrencyPrice;
    uint public currencyPrice;
    uint public currencySupply;
    uint public mintedReward;

    function SystemVars(uint initialPrice, uint initialSupply) public {
        prevCurrencyPrice = initialPrice;
        currencyPrice = initialPrice;
        mintedReward = initialSupply;
        currencySupply = initialSupply;
    }

    function _hasEnoughSupply() private view returns (bool) {
        return currencySupply >= lowSupplyMetric;
    }

    function _cap() private view returns (uint amount) {
        return (((block.number + 1) > 1) && _hasEnoughSupply()) ? currencySupply/10000 : initialCap;
    }

    function mintedAmount() public view returns (uint) {
        if ((block.number + 1) == 1) return initialMintedAmount;

        uint adjustment = mintedReward/adjustmentFactor;
        if ((currencyPrice > prevCurrencyPrice) && (prevCurrencyPrice > stabilizedPrice)) {
            return Math.min256(mintedReward + adjustment, _cap());
        }
        return Math.max256(mintedReward - adjustment, maxUnderNormalConditions);
    }

    function oracleDeduction(uint mintedAmount) public view returns (uint) {
        return oracleDeductionFraction * mintedAmount / 100;
    }

    function oracleReward() public view returns (uint) {
        return Math.min256(defaultOracleReward, this.balance);
    }
}