pragma solidity 0.4.24;

/**
 * Abstract contract for PublicResolver.
 */
contract DomainResolver {
    /**
     * Returns the address associated with an KNS node.
     * @param node The KNS node to query.
     * @return The associated address.
     */
    function addr(bytes32 node) public view returns (address);
}
