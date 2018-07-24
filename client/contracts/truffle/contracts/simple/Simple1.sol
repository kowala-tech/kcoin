pragma solidity 0.4.24;

import "../proxy/Initializable.sol";

contract Simple1 is Initializable {
    uint256 public value;

    function initialize(uint256 _value) isInitializer public {
        value = _value;
    }
}