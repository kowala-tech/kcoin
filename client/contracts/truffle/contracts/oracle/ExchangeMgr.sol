pragma solidity 0.4.24;

import "openzeppelin-solidity/contracts/lifecycle/Pausable.sol";
import "zos-lib/contracts/migrations/Initializable.sol";

/**
@title Exchange Manager contract
*/
contract ExchangeMgr is Pausable, Initializable {

    /*
     * Events
     */

    event Whitelisted(string exchange);
    event Blacklisted(string exchange);
    event Addition(string exchange);
    event Removal(string exchange);

    /*
     * Storage
     */

    mapping (string => Exchange) private exchangeRegistry;

    // whitelist contains the list of whitelisted exchanges
    string[] public whitelist;

    struct Exchange {
        uint index;
        bool isExchanged;
    }
    
    /*
     * Modifiers
     */
    
    modifier onlyNewCandidate(string exchange) {
        require(!isExchange(exchange));
        _;
    }
    
    /*
     * Public functions
     */

    function addExchange(string name) public onlyOwner {
        Exchange exchange = exchangeRegistry[name];
        exchange.isExchange = true;
        emit Addition(name);
        whitelistExchange(name);
    }

    function removeExchange(string exchange) public onlyOwner {
        delete oracleRegistry[exchange];
        emit Removal(exchange);
    }

    function whitelistExchange(string exchange) public onlyExchange(exchange) {
        exchange.index = whitelist.push(exchange) - 1;
        emit Whitelisted(exchange);
    }

    /**
     * @dev Blacklists given exchange
     * @param identity Address of an Oracle.
     */
    function blacklistExchange(string name) public onlyExchange(exchange) {
        Exchange exchange = exchangeRegistry[name];
        uint rowToDelete = exchange.index;
        exchange.index = 0;

         // replace the deprecated record with the last element
        address keyToMove = whitelist[whitelist.length-1];
        whitelistl[rowToDelete] = keyToMove;
        exchangeRegistry[keyToMove].index = rowToDelete;
        whitelist.length--;    
        emit Blacklisted(exchange);
    }

    function isExchange(string name) public view returns (bool isIndeed) {
        return exchangeRegistry[name].isExchange;
    }

    function isWhitelistedExchange(string name) public view returns (bool isIndeed) {
        return oracleRegistry[name].index > 0;
    }

    /**
     * @dev Checks if given exchange name is whitelisted
     * @param identity Address of an Oracle.
     */
    function isWhitelistedExchange(string exchange) public view returns (bool isIndeed) {
        return exchangeRegistry[exchange].isWhitelisted;
    }

    /**
     * @dev Get whitelisted exchange count
     */
    function getWhitelistedExchangeCount() public view returns (uint count) {
        return oraclePool.length;
    }

    /**
     * @dev Get whitelisted exchange information
     * @param index index of an Oracle to check.
     */
    function getWhitelistedExchangeAtIndex(uint index) public view returns (string name) {
        return whitelist[index];
    }
}