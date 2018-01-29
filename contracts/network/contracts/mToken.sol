pragma solidity ^0.4.18;

import "./ownable.sol";

// Simplified version of ERC20.
contract ERC20SimpleInterface {
    // Token name.
    string public name;
    // Symbol.
    string public symbol;
    // Decimal places.
    uint8 public decimals;

    // Get the total token supply.
    function totalSupply() public view returns (uint256 totalTokenSupply);

    // Get the account balance of another account with address addr.
    function balanceOf(address addr) public view returns (uint256 balance);

    // Send amount tokens to address toAddr. Returns true on success.
    function transfer(address toAddr, uint256 amount) public returns (bool success);

    // Triggered when tokens are transferred.
    event Transfer(address indexed fromAddr, address indexed toAddr, uint256 amount);
}

// MintableInterface token.
contract MintableInterface is Ownable {
    // Return maximum supply of tokens.
    function maximumSupply() public view returns (uint256 maxTokenSupply);

    // Mint new mTokens.
    function mintTokens(address addr, uint256 amount) onlyOwner public returns (bool success);

    // Trigger when minting new tokens.
    event Mint(address indexed addr, uint256 amount);
}

// DelegableInterface token
contract DelegableInterface {
    // Delegate mount tokens to an delegateAddr.
    function delegate(address delegateAddr, uint256 amount) public returns (bool success);

    // Revoke previously delegated amount tokens to delegateAddr.
    function revoke(address delegateAddr, uint256 amount) public returns (bool success);

    // Amount of tokens delegated from delegatorAddr.
    function delegatedFrom(address delegatorAddr) public view returns (uint256 amount);

    // Amount of tokens delegated to delegateAddr.
    function delegatedTo(address delegateAddr) public view returns (uint256 amount);

    // Trigger when delegation happens.
    event Delegation(address indexed ownerAddr, address indexed delegateAddr, uint256 amount);
    // Triggered when a revocation happend.
    event Revocation(address indexed ownerAddr, address indexed delegateAddr, uint256 amount);
}

contract mToken is Ownable, ERC20SimpleInterface, MintableInterface, DelegableInterface {
    // Amount of tokens owned by address.
    mapping(address => uint256) ownedTokens;
    // Total token supply.
    uint256 totalTokens = 0;
    // Maximum token supply.
    uint256 internal maxTokens = 0;
    // Amount of tokens hold by delegates.
    mapping(address => uint256) delegatesTokens;
    // Amount of tokens delegated.
    mapping(address => uint256) delegatedTokens;
    // Amount of tokens delegated ( tokenDelegations[delegate][delegator] ).
    mapping(address => mapping(address => uint256)) tokenDelegations;

    function() public {
        // Don't receive any ether.
        revert();
    }

    // Implement ERC20SimpleInterface.
    // Get the total token supply.
    function totalSupply() public view returns (uint256 totalTokenSupply) {
        return totalTokens;
    }

    // Implement MintableInterface.
    // Return maximum supply of tokens.
    function maximumSupply() public view returns (uint256 maxTokenSupply) {
        return maxTokens;
    }

    // Mint new mTokens.
    function mintTokens(address addr, uint256 amount) onlyOwner public returns (bool success) {
        if (totalTokens + amount > maxTokens) {
            return false;
        }
        ownedTokens[addr] += amount;
        totalTokens += amount;
        Mint(addr, amount);
        return true;
    }

    // Implement DelegableInterface.
    // Delegate mount tokens to an delegateAddr.
    function delegate(address delegateAddr, uint256 amount) public returns (bool success) {
        if (ownedTokens[msg.sender] - delegatedTokens[msg.sender] < amount) {
            return false;
        }
        delegatedTokens[msg.sender] += amount;
        delegatesTokens[delegateAddr] += amount;
        tokenDelegations[delegateAddr][msg.sender] += amount;
        Delegation(msg.sender, delegateAddr, amount);
        return true;
    }

    // Revoke previously delegated amount tokens to delegateAddr.
    function revoke(address delegateAddr, uint256 amount) public returns (bool success) {
        if (delegatedTokens[msg.sender] < amount) {
            return false;
        }
        if (tokenDelegations[delegateAddr][msg.sender] < amount) {
            return false;
        }
        delegatedTokens[msg.sender] -= amount;
        delegatesTokens[delegateAddr] -= amount;
        tokenDelegations[delegateAddr][msg.sender] -= amount;
        Revocation(msg.sender, delegateAddr, amount);
        return true;
    }

    // Amount of tokens delegated from delegatorAddr.
    function delegatedFrom(address delegatorAddr) public view returns (uint256 amount) {
        return tokenDelegations[msg.sender][delegatorAddr];
    }

    // Amount of tokens delegated to delegateAddr.
    function delegatedTo(address delegateAddr) public view returns (uint256 amount) {
        return tokenDelegations[delegateAddr][msg.sender];
    }

    // Get the account balance of another account with address addr.
    // Balance needs to include delegated tokens.
    function balanceOf(address addr) public view returns (uint256 balance) {
        return ownedTokens[addr] + delegatesTokens[addr] - delegatedTokens[addr];
    }

    // Send amount tokens to address toAddr. Returns true on success.
    // Can only transfer owned tokens.
    function transfer(address toAddr, uint256 amount) public returns (bool success) {
        if (this.balanceOf(msg.sender) < amount) {
            return false;
        }
        ownedTokens[msg.sender] -= amount;
        ownedTokens[toAddr] += amount;
        Transfer(msg.sender, toAddr, amount);
        return true;
    }
}

