pragma solidity ^0.4.18;

import "./ownable.sol";

// PriceOracle makes a relationship between a cryptocurrency
// and a fiat one. the smallest units are always used:
// cryptocurrency: 1 crypto/10^cryptoDecimals
// fiat: 1 fiat/10^fiatDecimals
contract PriceOracle is Ownable {
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
    // Amounts of each currency to store the relationship.
    uint256 cryptoAmount;
    uint256 fiatAmount;

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

    // Return the amount of the cryptocurrency corresponding to fiatAmount.
    function priceForFiat(uint256 _fiatAmount) public view returns (uint256 _cryptoAmount) {
        return _fiatAmount * cryptoAmount / fiatAmount;
    }

    // Returns 10**fiatDecimals.
    function oneFiat() public view returns (uint256 _fiatAmount) {
        return uint256(10)**fiatDecimals;
    }

    // Returns the price for one fiat.
    function priceForOneFiat() public view returns (uint256 _cryptoAmount) {
        return priceForFiat(oneFiat());
    }

    // Return the amount of fiat corresponding to cryptoAmount.
    function priceForCrypto(uint256 _cryptoAmount) public view returns (uint256 _fiatAmount) {
        return _cryptoAmount * fiatAmount / cryptoAmount;
    }

    // Returns 10**cryptoDecimals.
    function oneCrypto() public view returns (uint256 _cryptoAmount) {
        return uint256(10)**cryptoDecimals;
    }

    // Return the price for one crypto.
    function priceForOneCrypto() public view returns (uint256 _fiatAmount) {
        return priceForCrypto(oneCrypto());
    }

    // Set the price.
    function setPrice(uint256 _cryptoAmount, uint256 _fiatAmount) onlyOwner public returns (bool success) {
        cryptoAmount = _cryptoAmount;
        fiatAmount = _fiatAmount;
        NewPrice(cryptoAmount, fiatAmount);
        return true;
    }

    // Triggered when a new price is set.
    event NewPrice(uint256 cryptoPrice, uint256 fiatPrice);
}
