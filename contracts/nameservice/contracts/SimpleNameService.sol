pragma solidity 0.4.21;

import "github.com/kowala-tech/kcoin/contracts/lifecycle/contracts/Pausable.sol" as pausable;

// NOTE (rgeraldes) - // 0x0 uninitialized address

contract SimpleNameService is pausable.Pausable {
    struct Entry {
        address addr;
        bool isDomain;
    }

    mapping(string=>Entry) registry;

    modifier whenDomain(string domain) {
        require(registry[domain].isDomain);
        _;
    }

    function register(string domain, address addr) public {
        registry[domain] = Entry(addr, true);
    }

    function lookup(string domain) whenDomain(domain) public view returns (address addr) {
        return registry[domain].addr;
    }
}