pragma solidity 0.4.24;

/**
 * Mock contract for PublicResolver.
 */
contract DomainResolver {

    address domainAddr;

    constructor(address _domainAddr){
      domainAddr = _domainAddr;
    }

    function addr(bytes32 node) public view returns (address){
      return domainAddr;
    }
}
