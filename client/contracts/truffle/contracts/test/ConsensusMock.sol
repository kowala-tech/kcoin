pragma solidity 0.4.24;

contract ConsensusMock {
    bool superNode;

    function ConsensusMock(bool _superNode) public {
        superNode = _superNode;
    }

    function isSuperNode(address identity) public view returns (bool) {
        return superNode;
    }

}