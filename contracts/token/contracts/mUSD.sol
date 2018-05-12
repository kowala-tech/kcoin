pragma solidity 0.4.21;

import "./CappedToken.sol";

contract mUSD is CappedToken {
    function mUSD() public CappedToken(1073741824 ether) {
        name = "mUSD";
        symbol = "mUSD";
        decimals = 18;

        mint(0x80eDa603028fe504B57D14d947c8087c1798D800, 161061273.6 ether);
    }
}