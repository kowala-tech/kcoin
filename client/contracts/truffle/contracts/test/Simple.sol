pragma solidity 0.4.24;

contract Simple {
    uint256 public x;
    
    constructor(uint _x) public {
        x = _x;
    }

    function viewX() view public returns(uint) {
        return x;
    }
}