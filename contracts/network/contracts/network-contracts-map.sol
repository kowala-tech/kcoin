pragma solidity ^0.4.18;

import "./ownable.sol";

// Updatable list of network addresses.
contract NetworkContractsMap is Ownable {
    // mToken contract.
    address public mToken;
    // kUSD/USD price oracle.
    address public priceOracle;
    // Network stats and info.
    address public networkStats;

    function NetworkContractsMap(address _mToken, address _priceOracle, address _networkStats) public {
        mToken = _mToken;
        priceOracle = _priceOracle;
        networkStats = _networkStats;
    }

    // Set mToken contract address.
    function setMToken(address addr) onlyOwner public {
        mToken = addr;
    }

    // Set price oracle address.
    function setPriceOracle(address addr) onlyOwner public {
        priceOracle = addr;
    }

    // Set network stats address.
    function setNetworkStats(address addr) onlyOwner public {
        networkStats = addr;
    }
}