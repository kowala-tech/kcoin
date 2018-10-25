pragma solidity 0.4.24;

import "../token/TokenReceiver.sol";

contract BalanceContract is TokenReceiver {
    address owner;
    address public from;
    uint public value;
    bytes data;
    constructor () public {
        owner = msg.sender;
    }

    function tokenFallback(address _from, uint _value, bytes _data) public {
        from = _from;
        value = _value;
        data = _data;
    }
}