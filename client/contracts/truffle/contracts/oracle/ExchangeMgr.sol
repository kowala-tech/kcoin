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

    string[] public whitelist;

    struct Exchange {
        uint index;
        bool isExchange;
        bool isWhitelisted;
    }
    
    /*
     * Modifiers
     */
    
    modifier onlyNewCandidate(string name) {
        require(!isExchange(name), "exchange already exists");
        _;
    }

    modifier onlyExchange(string name) {
        require(isExchange(name), "given name is not an exchange");
        _;
    }

    modifier onlyWhitelistedExchange(string name) {
        require(isWhitelistedExchange(name), "given name is not a whitelisted exchange");
        _;
    }

    modifier onlyBlacklistedExchange(string name) {
        require(isBlacklistedExchange(name), "given name is not a blacklisted exchange");
        _;
    }
    
    /*
     * Public functions
     */

    /**
     * @dev Adds and whitelists an exchange.
     * @param name exchange name.
     */
    function addExchange(string name) public whenNotPaused onlyOwner onlyNewCandidate(name) {
        Exchange exchange = exchangeRegistry[name];
        exchange.isExchange = true;
        emit Addition(name);
        whitelistExchange(name);
    }

    /**
     * @dev Removes an exchange.
     * @param name exchange name.
     */
    function removeExchange(string name) public whenNotPaused onlyOwner onlyExchange(name) {
        delete exchangeRegistry[name];
        emit Removal(name);
    }

    /**
     * @dev Whitelists an exchange.
     * @param name exchange name.
     */
    function whitelistExchange(string name) public whenNotPaused onlyOwner onlyBlacklistedExchange(name) {
        exchangeRegistry[name].index = whitelist.push(name) - 1;
        exchangeRegistry[name].isWhitelisted = true;
        emit Whitelisted(name);
    }

    /**
     * @dev Blacklists an exchange.
     * @param name exchange name.
     */
    function blacklistExchange(string name) public whenNotPaused onlyOwner onlyWhitelistedExchange(name) {
        Exchange exchange = exchangeRegistry[name];
        uint rowToDelete = exchange.index;
        exchange.isWhitelisted = false;

         // replace the deprecated record with the last element
        string keyToMove = whitelist[whitelist.length-1];
        whitelist[rowToDelete] = keyToMove;
        exchangeRegistry[keyToMove].index = rowToDelete;
        whitelist.length--;    
        emit Blacklisted(name);
    }

    /**
     * @dev checks whether the given name is a whitelisted exchange or not
     * @param name exchange name.
     */
    function isExchange(string name) public view returns (bool isIndeed) {
        return exchangeRegistry[name].isExchange;
    }

    /**
     * @dev checks whether the given name is a whitelisted exchange or not
     * @param name exchange name.
     */
    function isWhitelistedExchange(string name) public view returns (bool isIndeed) {
        return isExchange(name) && exchangeRegistry[name].isWhitelisted;
    }

    /**
     * @dev checks whether the given name is a blacklisted exchange or not
     * @param name exchange name.
     */
    function isBlacklistedExchange(string name) public view returns (bool isIndeed) {
        return isExchange(name) && !exchangeRegistry[name].isWhitelisted;
    }

    /**
     * @dev get whitelisted exchange count
     */
    function getWhitelistedExchangeCount() public view returns (uint count) {
        return whitelist.length;
    }

    /**
     * @dev get whitelisted exchange information
     * @param index index of a given exchange in the whitelist
     */
    function getWhitelistedExchangeAtIndex(uint index) public view returns (string name) {
        return whitelist[index];
    }
}
