pragma solidity ^0.4.18;

import "./mToken.sol";

contract mUSD is mToken {
    function mUSD() public {
        name = "mUSD";
        symbol = "mUSD";
        maxTokens = 1073741824;
    }
}