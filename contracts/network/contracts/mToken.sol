pragma solidity ^0.4.18;

import "./ownable.sol";

contract mToken is Ownable {
    // Token name.
    string public name;
    // Symbol.
    string public symbol;
    // Decimal places.
    uint8 public decimals;
    // Maximum token supply.
    uint256 public maximumSupply = 0;
    // Total token supply.
    uint256 public totalSupply = 0;
    // Amount of tokens owned by address.
    mapping(address => uint256) public ownedBy;
    // all token holders addresses.
    address[] public tokenHolders;
    // token holders index.
    mapping(address => uint256) tokenHoldersIndex;
    // Amount of tokens hold by delegates.
    mapping(address => uint256) delegatesTokens;
    // Amount of tokens delegated.
    mapping(address => uint256) delegatedTokens;
    // Amount of tokens delegated ( tokenDelegations[delegate][delegator] ).
    mapping(address => mapping(address => uint256)) tokenDelegations;

    // Mining addresses proposals minersProposals[miningAddr] = ownerAddr.
    mapping(address => address) minersProposals;
    // Miners/owners minersOwners[miningAddr] = ownerAddr.
    mapping(address => address) public minersOwners;
    // Owners/miners ownersMiners[ownerAddr] = miningAddr.
    mapping(address => address) public ownersMiners;
    // Receiver addresses proposals receiversProposals[receiverAddr] = miningAddr.
    mapping(address => address) receiversProposals;
    // Miners receiver addresses.
    mapping(address => address) public minersReceivers;

    function() public {
        // Don't receive any kUSD.
        revert();
    }

    // Return the total count of token holders.
    function numberTokenHolders() public view returns (uint256 count) {
        return tokenHolders.length;
    }

    // Initialize account. Only for the contract owner.
    function initializeAccount(
        address ownerAddr,
        address miningAddr,
        address receiverAddr
    ) onlyOwner public {
        ownersMiners[ownerAddr] = miningAddr;
        minersOwners[miningAddr] = ownerAddr;
        minersReceivers[miningAddr] = receiverAddr;
    }

    // proposeMiningAddress must be called by the owner address.
    // The mining address change proposal must be accepted by the mining address.
    function proposeMiningAddress(address miningAddr) public {
        minersProposals[miningAddr] = msg.sender;
    }

    // acceptMiningAddress must be called by the mining address. A proposal must exist
    // and the owner addresses (the proposed and the one provided) must match.
    function acceptMiningAddress(address ownerAddr) public returns (bool success) {
        if (minersProposals[msg.sender] != ownerAddr) {
            return false;
        }
        delete minersProposals[msg.sender];
        ownersMiners[ownerAddr] = msg.sender;
        minersOwners[msg.sender] = ownerAddr;
        NewMiningAddress(ownerAddr, msg.sender);
        return true;
    }

    // Propose a new receiver address. Must be called by the owner address
    // and accepted by the the mining address (must be set).
    function proposeReceiverAddress(address receiverAddr) public {
        receiversProposals[receiverAddr] = ownersMiners[msg.sender];
    }

    // Accept receiver address proposal. A change to the receiver address must
    // be accepted by the mining address. The receiver addresses must match
    // (the proposed and the provided one).
    function acceptReceiverAddress(address receiverAddr) public returns (bool success) {
        if (receiversProposals[receiverAddr] != msg.sender) {
            return false;
        }
        delete receiversProposals[receiverAddr];
        minersReceivers[msg.sender] = receiverAddr;
        NewReceiverAddress(msg.sender, receiverAddr);
        return true;
    }

    // proposeAccountAddresses must be called by the owner address.
    function proposeAccountAddresses(address miningAddr, address receiverAddr) public {
        proposeMiningAddress(miningAddr);
        proposeReceiverAddress(receiverAddr);
    }

    // acceptAccountAddresses must be called by the mining address. Both the
    // owner addresses and receiver addresses must match (proposed and provided).
    function acceptAccountAddresses(address ownerAddr, address receiverAddr) public returns (bool success) {
        return acceptMiningAddress(ownerAddr) && acceptReceiverAddress(receiverAddr);
    }

    // Return the mining and receiver addresses for ownerAddr.
    function addressesOf(address ownerAddr) public view returns (address minerAddr, address receiveAddr) {
        address m = ownersMiners[ownerAddr];
        return (m, minersReceivers[m]);
    }

    // Add address to the token holders index.
    function addHolder(address addr) private {
        tokenHolders.push(addr);
        tokenHoldersIndex[addr] = tokenHolders.length;
    }

    // Remove address from the token holders index.
    function removeHolder(address addr) private {
        uint256 addrIdx = tokenHoldersIndex[addr];
        if (addrIdx == 0) {
            return;
        }
        addrIdx--;
        address lastAddr = tokenHolders[tokenHolders.length - 1];
        tokenHolders[addrIdx] = lastAddr;
        tokenHoldersIndex[lastAddr] = addrIdx + 1;
        delete tokenHoldersIndex[addr];
        tokenHolders.length--;
    }

    // Mint new mTokens.
    function mint(address addr, uint256 amount) onlyOwner public returns (bool success) {
        if (totalSupply + amount > maximumSupply) {
            return false;
        }
        if (ownedBy[addr] == 0) {
            addHolder(addr);
        }
        ownedBy[addr] += amount;
        totalSupply += amount;
        Mint(addr, amount);
        return true;
    }

    // Delegate mount tokens to an delegateAddr.
    function delegate(address delegateAddr, uint256 amount) public returns (bool success) {
        if (ownedBy[msg.sender] - delegatedTokens[msg.sender] < amount) {
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
    function availableTo(address addr) public view returns (uint256 balance) {
        return ownedBy[addr] + delegatesTokens[addr] - delegatedTokens[addr];
    }

    // Returns the amount of tokens available for delegation for a given address.
    function availableForDelegationTo(address addr) public view returns (uint256 amount) {
        return availableTo(addr) - delegatedTokens[addr];
    }

    // Returns the amount of tokens available for delegation.
    function availableForDelegation() public view returns (uint256 amount) {
        return availableForDelegationTo(msg.sender);
    }

    // Send amount tokens to address toAddr. Returns true on success.
    // Can only transfer owned tokens that are not under delegation.
    function transfer(address toAddr, uint256 amount) public returns (bool success) {
        if (availableTo(msg.sender) - delegatesTokens[msg.sender] < amount) {
            return false;
        }
        ownedBy[msg.sender] -= amount;
        ownedBy[toAddr] += amount;
        if (ownedBy[msg.sender] == 0) {
            removeHolder(msg.sender);
        }
        if (ownedBy[toAddr] == amount) {
            addHolder(toAddr);
        }
        Transfer(msg.sender, toAddr, amount);
        return true;
    }

    // Events.
    // Triggered when tokens are transferred.
    event Transfer(address indexed fromAddr, address indexed toAddr, uint256 amount);

    // Triggered when minting new tokens.
    event Mint(address indexed addr, uint256 amount);

    // Triggered when a change of mining address happens.
    event NewMiningAddress(address managementAddr, address miningAddr);

    // Triggered when a change of receiver address happens.
    event NewReceiverAddress(address miningAddr, address receiverAddr);

    // Triggered when delegation happens.
    event Delegation(address indexed ownerAddr, address indexed delegateAddr, uint256 amount);

    // Triggered when a revocation happens.
    event Revocation(address indexed ownerAddr, address indexed delegateAddr, uint256 amount);
}

