pragma solidity 0.4.24;

import "./CappedToken.sol";

contract MiningToken is CappedToken {
    function MiningToken(string _name, string _symbol, uint _cap, uint8 _decimals) public CappedToken(_cap) {
        name = _name;
        symbol = _symbol;
        decimals = _decimals;
    }

    function initialize(string _name, string _symbol, uint _cap, uint8 _decimals) isInitializer public {
        require(_cap > 0);
        cap = _cap;
        name = _name;
        symbol = _symbol;
        decimals = _decimals;
    }
}