pragma solidity 0.4.24;

import "openzeppelin-solidity/contracts/lifecycle/Pausable.sol";
import "../sysvars/SystemVars.sol";

/**
 * @title Stability contract supports network utility
 */
contract Stability is Pausable {

    uint constant one = 1 ether;
    
    uint minDeposit;
    SystemVars sysvars;
    
    struct Subscription {
        uint index;
        bool hasSubscription;
        uint deposit;
    }

    mapping (address => Subscription) private subscriptionRegistry;

    address[] private subscriptionPool;

    modifier onlySubscriber {
        require(hasSubscription(msg.sender));
        _;
    }

    modifier onlyWithMinDeposit {
        require(msg.value >= minDeposit);
        _;
    }

    modifier whenPriceGreaterEqualOne {
        require(sysvars.price() >= one);
        _;
    }

    /**
     * Constructor.
     * @param _systemVarsAddr Address of system variables contract.
     * @param _minDeposit minimum deposit required.
     */
    function Stability(uint _minDeposit, address _systemVarsAddr) public {
        minDeposit = _minDeposit;
        sysvars = SystemVars(_systemVarsAddr);
    }

    function hasSubscription(address identity) public view returns (bool isIndeed) {
        return subscriptionRegistry[identity].hasSubscription;
    }

    function _insertSubscription() private onlyWithMinDeposit {
        Subscription subs = subscriptionRegistry[msg.sender];
        subs.index = subscriptionPool.push(msg.sender) - 1;
        subs.hasSubscription = true;
        subs.deposit = msg.value;
    }

    function subscribe() public payable whenNotPaused {
        if (hasSubscription(msg.sender)) {
            Subscription subs = subscriptionRegistry[msg.sender];
            subs.deposit += msg.value;
            return;
        }
        
        _insertSubscription();
    }

    function unsubscribe() public onlySubscriber whenPriceGreaterEqualOne {
        Subscription subs = subscriptionRegistry[msg.sender];
        msg.sender.transfer(subs.deposit);
        subs.deposit = 0;
        subs.hasSubscription = false;
    }
}