pragma solidity ^0.4.18;

import "./ownable.sol";

contract NetworkStats is Ownable {
    // Total mined wei
    uint256 public totalMinedWei = 0;
}