pragma solidity ^0.4.18;

import "./ownable.sol";

contract NetworkStats is Ownable {
    // Total supply of wei. Must be updated every block.
    uint256 public totalSupplyWei = 1 ether;
    // Reward calculated for the last block. Must be updated every block.
    uint256 public lastBlockReward = 0;
    // Price established by the price oracle for the last block. Must be updated every block.
    uint256 public lastPrice = 0;
}