pragma solidity ^0.4.18;

// Ownable.
contract OwnableInterface {
    // Transfer ownership.
    function transferOwnership(address addr) public;

    // Log ownership transfers.
    event OwnershipTransfer(address oldAddr, address newAddr);
}

contract Ownable is OwnableInterface {
    // Contract owner address.
    address contractOwner;

    function Ownable() public {
        // Set contract owner.
        contractOwner = msg.sender;
    }

    // Only allow owner.
    modifier onlyOwner {
        require(msg.sender == contractOwner);
        _;
    }

    // Transfer ownership.
    function transferOwnership(address addr) public onlyOwner {
        contractOwner = addr;
        OwnershipTransfer(contractOwner, addr);
    }

    // Log ownership transfers.
    event OwnershipTransfer(address oldAddr, address newAddr);
}