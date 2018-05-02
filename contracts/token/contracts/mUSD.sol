pragma solidity 0.4.21;

import "./CappedToken.sol";

contract mUSD is CappedToken {
    function mUSD() public CappedToken(1073741824) {
        name = "mUSD";
        symbol = "mUSD";
        decimals = 18;
    }
}