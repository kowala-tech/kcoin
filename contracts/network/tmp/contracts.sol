pragma solidity ^0.4.18;

import "./ownable.sol";

// Updatable list of network addresses.
contract Contracts is Ownable {
    // mToken contract.
    address public mToken;
    // kUSD/USD price oracle.
    address public priceOracle;
    // Network contract
    address public networkContract;

    function Contracts(address _mToken, address _priceOracle, address _networkContract) public {
        mToken = _mToken;
        priceOracle = _priceOracle;
        networkContract = _networkContract;
    }

    // Set mToken contract address.
    function setMToken(address addr) onlyOwner public {
        mToken = addr;
    }

    // Set price oracle address.
    function setPriceOracle(address addr) onlyOwner public {
        priceOracle = addr;
    }

    // Set network contract address.
    function setNetworkContract(address addr) onlyOwner public {
        networkContract = addr;
    }
}