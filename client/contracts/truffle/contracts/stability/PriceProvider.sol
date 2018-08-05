pragma solidity 0.4.24;

/**
 * @title PriceProvider interface
 */

 contract PriceProvider {
     function price() public view returns (uint256);
 }