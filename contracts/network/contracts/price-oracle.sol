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
    function priceForFiat(uint256 _fiatAmount) public view returns (uint256 _cryptoAmount);

    // Return the amount of fiat corresponding to cryptoAmount.
    function priceForCrypto(uint256 _cryptoAmount) public view returns (uint256 _fiatAmount);

    // Set the price.
    function setPrice(uint256 _cryptoAmount, uint256 _fiatAmount) public returns (bool success);

    // Triggered when a new price is set.
    event NewPrice(uint256 cryptoPrice, uint256 fiatPrice);
}

// Simple implementation.
contract PriceOracle is Ownable, PriceOracleInterface {
    // Amounts of each currency to store the relationship.
    uint256 cryptoAmount = 0;
    uint256 fiatAmount = 0;

    // Initialize.
    function PriceOracle(
        string _cryptoName,
        string _cryptoSymbol,
        uint8 _cryptoDecimals,
        uint256 _cryptoAmount,
        string _fiatName,
        string _fiatSymbol,
        uint8 _fiatDecimals,
        uint256 _fiatAmount
    ) public {
        cryptoName = _cryptoName;
        cryptoSymbol = _cryptoSymbol;
        cryptoDecimals = _cryptoDecimals;
        cryptoAmount = _cryptoAmount;
        fiatName = _fiatName;
        fiatSymbol = _fiatSymbol;
        fiatDecimals = _fiatDecimals;
        fiatAmount = _fiatAmount;
    }

    // Return the amount of the crytocurrency corresponding to fiatAmount.
    function priceForFiat(uint256 _fiatAmount) public view returns (uint256 _cryptoAmount) {
        return _fiatAmount * cryptoAmount / fiatAmount;
    }

    // Return the amount of fiat corresponding to cryptoAmount.
    function priceForCrypto(uint256 _cryptoAmount) public view returns (uint256 _fiatAmount) {
        return _cryptoAmount * fiatAmount / cryptoAmount;
    }

    // Set the price.
    function setPrice(uint256 _cryptoAmount, uint256 _fiatAmount) onlyOwner public returns (bool success) {
        cryptoAmount = _cryptoAmount;
        fiatAmount = _fiatAmount;
        NewPrice(cryptoAmount, fiatAmount);
        return true;
    }
}

