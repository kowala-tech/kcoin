pragma solidity ^0.4.18;

import "./mToken.sol";

contract mUSD is mToken {
    function mUSD() public {
        name = "mUSD";
        symbol = "mUSD";
        maxTokens = 1073741824;
        mintTokens(0XC57BF12BB34F6FD85BDBF0CACA983528422BF7A2, 100);
    }
}