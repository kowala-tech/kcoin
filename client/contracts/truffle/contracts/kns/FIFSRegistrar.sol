pragma solidity 0.4.24;

import "./KNS.sol";
import "zos-lib/contracts/migrations/Initializable.sol";

/**
 * A registrar that allocates subdomains to the first person to claim them.
 */
contract FIFSRegistrar is Initializable{
    KNS kns;
    bytes32 rootNode;

    modifier only_owner(bytes32 subnode) {
        address currentOwner = kns.owner(keccak256(rootNode, subnode));
        require(currentOwner == 0 || currentOwner == msg.sender);
        _;
    }

    /**
     * Constructor.
     * @param knsAddr The address of the KNS registry.
     * @param node The node that this registrar administers.
     */
    constructor(KNS knsAddr, bytes32 node) public {
        kns = knsAddr;
        rootNode = node;
    }

    /**
     * @dev initialize function for Proxy Pattern.
     * @param knsAddr The address of the KNS registry.
     * @param node The node that this registrar administers.
     */
    function initialize(KNS knsAddr, bytes32 node) isInitializer public {
        kns = knsAddr;
        rootNode = node;
    }

    /**
     * Register a name, or change the owner of an existing registration.
     * @param subnode The hash of the label to register.
     * @param owner The address of the new owner.
     */
    function register(bytes32 subnode, address owner) public only_owner(subnode) {
        kns.setSubnodeOwner(rootNode, subnode, owner);
    }
}
