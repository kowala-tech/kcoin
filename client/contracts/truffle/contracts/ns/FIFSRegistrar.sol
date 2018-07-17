pragma solidity 0.4.24;

import "./NS.sol";

/**
 * A registrar that allocates subdomains to the first person to claim them.
 */
contract FIFSRegistrar {
    NS ns;
    bytes32 rootNode;

    modifier only_owner(bytes32 subnode) {
        address currentOwner = ns.owner(keccak256(rootNode, subnode));
        require(currentOwner == 0 || currentOwner == msg.sender);
        _;
    }

    /**
     * Constructor.
     * @param nsAddr The address of the NS registry.
     * @param node The node that this registrar administers.
     */
    constructor(NS nsAddr, bytes32 node) public {
        ns = nsAddr;
        rootNode = node;
    }

    /**
     * Register a name, or change the owner of an existing registration.
     * @param subnode The hash of the label to register.
     * @param owner The address of the new owner.
     */
    function register(bytes32 subnode, address owner) public only_owner(subnode) {
        ns.setSubnodeOwner(rootNode, subnode, owner);
    }
}
