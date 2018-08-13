pragma solidity 0.4.24;

contract PriceProviderMock {

    uint mockPrice;

    function PriceProviderMock(uint _price) public {
        mockPrice = _price;
    }
    
    function price() public view returns (uint256) {
        return mockPrice;
    }
}