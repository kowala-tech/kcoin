pragma solidity 0.4.24;

contract DomainResolverMock {
    address domainAddr;

    constructor(address _domainAddr){
        domainAddr = _domainAddr;
    }

    function addr(bytes32 node) public view returns (address){
        return domainAddr;
    }
}
