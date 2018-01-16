pragma solidity ^0.4.18;

import "./mToken.sol";

contract mUSD is mToken {
    function mUSD() public {
        name = "mUSD";
        symbol = "mUSD";
        maxTokens = 1073741824;
        // Initial token holders
        mintTokens(0xd6e579085c82329c89fca7a9f012be59028ed53f, 100);
        mintTokens(0x497dc8a0096cf116e696ba9072516c92383770ed, 100);
        mintTokens(0x259be75d96876f2ada3d202722523e9cd4dd917d, 100);
    }
}