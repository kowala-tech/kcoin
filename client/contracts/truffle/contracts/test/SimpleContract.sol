pragma solidity 0.4.24;

import "zos-lib/contracts/migrations/Initializable.sol";

contract SimpleContract is Initializable {
    uint x;

    function initialize(uint _x) isInitializer public {
        x = _x;
    }

    function setX(uint _x) public {
        x = _x;
    }

    function readX() public view returns(uint){
        return x;
    }

}
