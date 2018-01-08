pragma solidity ^0.4.18;

contract NetworkStats {
    // Total supply of wei. Must be updated every block.
    uint256 public totalSupplyWei = 0;
    // Reward calculated for the last block. Must be updated every block.
    uint256 public lastBlockReward = 0;
    // Price established by the price oracle for the last block. Must be updated every block.
    uint256 public lastPrice = 0;
}