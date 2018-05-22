pragma solidity 0.4.21;

import "./CappedToken.sol";

contract MiningToken is CappedToken {
    function MiningToken(string _name, string _symbol, uint _cap, uint8 _decimals) public CappedToken(_cap) {
        name = _name;
        symbol = _symbol;
        decimals = _decimals;
    }
}