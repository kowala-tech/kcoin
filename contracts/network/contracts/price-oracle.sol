pragma solidity ^0.4.18;

import "./ownable.sol";

// PriceOracle makes a relationship between a cryptocurrency
// and a fiat one. the smallest units are always used:
// cryptocurrency: 1 crypto/10^cryptoDecimals
// fiat: 1 fiat/10^fiatDecimals
contract PriceOracleInterface {
    // Cryptocurrency name.
    string public cryptoName;
    // Cryptocurrency symbol.
    string public cryptoSymbol;
    // Cryptocurrency decimal places.
    uint8 public cryptoDecimals;
    // Fiat name.
    string public fiatName;
    // Fiat symbol.
    string public fiatSymbol;
    // Fiat decimal places.
    uint8 public fiatDecimals;

    // Return the amount of the crytocurrency corresponding to fiatAmount.
    function priceForFiat(uint256 fiatAmount) public view returns (uint256 cryptoAmount);

    // Return the amount of fiat corresponding to cryptoAmount.
    function priceForCrypto(uint256 cryptoAmount) public view returns (uint256 fiatAmount);

    // Set the price.
    function setPrice(uint256 cryptoAmount, uint256 fiatAmount) public returns (bool success);

    // Triggered when a new price is set.
    event NewPrice(uint256 cryptoPrice, uint256 fiatPrice);
}

// Simple implementation.
contract PriceOracle is Ownable, PriceOracleInterface {
    // Amounts of each currency to store the relationship.
    uint256 lastCryptoAmount = 0;
    uint256 lastFiatAmount = 0;

    // Initialize.
    function PriceOracle(uint256 cryptoAmount, uint256 fiatAmount) public {
        lastCryptoAmount = cryptoAmount;
        lastFiatAmount = fiatAmount;
    }

    // Return the amount of the crytocurrency corresponding to fiatAmount.
    function priceForFiat(uint256 fiatAmount) public view returns (uint256 cryptoAmount) {
        return fiatAmount * lastCryptoAmount / lastFiatAmount;
    }

    // Return the amount of fiat corresponding to cryptoAmount.
    function priceForCrypto(uint256 cryptoAmount) public view returns (uint256 fiatAmount) {
        return cryptoAmount * lastFiatAmount / lastCryptoAmount;
    }

    // Set the price.
    function setPrice(uint256 cryptoAmount, uint256 fiatAmount) onlyOwner public returns (bool success) {
        lastCryptoAmount = cryptoAmount;
        lastFiatAmount = fiatAmount;
        NewPrice(lastCryptoAmount, lastFiatAmount);
        return true;
    }
}

