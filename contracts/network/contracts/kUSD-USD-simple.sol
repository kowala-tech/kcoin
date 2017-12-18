pragma solidity ^0.4.18;

import "price-oracle.sol";

contract kUSD_USD is PriceOracle {
    function() public { revert(); }

    function kUSD_USD(uint256 cryptoAmount, uint256 fiatAmount) public {
        cryptoName = "kUSD cryptocurrency";
        cryptoSymbol = "kUSD";
        cryptoDecimals = 18;
        fiatName = "US Dollar";
        fiatSymbol = "USD";
        fiatDecimals = 4;

        setPrice(cryptoAmount, fiatAmount);
    }
}