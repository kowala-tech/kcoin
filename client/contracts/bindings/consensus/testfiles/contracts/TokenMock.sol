pragma solidity 0.4.24;

contract TokenMock {
    uint public totalSupply;
  
    mapping (address => uint) private balances;

    function balanceOf(address who) public view returns (uint) {
        return balances[who];
    }
  
    function name() public view returns (string _name) {
        return "mock";
    }
    
    function symbol() public view returns (string _symbol) {
        return "mock";
    }
    
    function decimals() public view returns (uint8 _decimals) {
        return 18;
    }
    
    function totalSupply() public view returns (uint256 _supply) {
        return 0;
    }

    function transfer(address to, uint value) public returns (bool ok) {
        balances[to] += value;
    }

    function transfer(address to, uint value, bytes data) public returns (bool ok) {
        balances[to] += value;
    }
    
    function transfer(address to, uint value, bytes data, string custom_fallback) public returns (bool ok) {
        balances[to] += value;
    }
  
    event Transfer(address indexed from, address indexed to, uint value, bytes indexed data);
}