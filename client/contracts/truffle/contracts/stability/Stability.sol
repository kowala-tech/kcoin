pragma solidity 0.4.24;

import "openzeppelin-solidity/contracts/lifecycle/Pausable.sol";
import "./PriceProvider.sol";

/**
 * @title Stability contract supports network utility
 */
contract Stability is Pausable {

    uint constant ONE = 1 ether;
    
    uint minDeposit;
    PriceProvider priceProvider;
    
    struct Subscription {
        uint index;
        bool hasSubscription;
        uint deposit;
    }

    mapping (address => Subscription) private subscriptionRegistry;

    address[] private subscriptionPool;

    modifier onlySubscriber {
        require(_hasSubscription(msg.sender));
        _;
    }

    modifier onlyWithMinDeposit {
        require(msg.value >= minDeposit);
        _;
    }

    modifier whenPriceGreaterEqualOne {
        require(priceProvider.price() >= ONE);
        _;
    }

    /**
     * Constructor
     * @param _minDeposit minimum deposit required to subscribe to the service
     * @param _priceProviderAddr address of system variables contract
     */
    function Stability(uint _minDeposit, address _priceProviderAddr) public {
        minDeposit = _minDeposit;
        priceProvider = PriceProvider(_priceProviderAddr);
    }

    function _hasSubscription(address identity) private view returns (bool isIndeed) {
        return subscriptionRegistry[identity].hasSubscription;
    }

    function _insertSubscription() private onlyWithMinDeposit {
        Subscription subs = subscriptionRegistry[msg.sender];
        subs.index = subscriptionPool.push(msg.sender) - 1;
        subs.hasSubscription = true;
        subs.deposit = msg.value;
    }

     /**
     * @dev Subscribe to the stability contract service
     */
    function subscribe() public payable whenNotPaused {
        if (_hasSubscription(msg.sender)) {
            Subscription subs = subscriptionRegistry[msg.sender];
            subs.deposit += msg.value;
            return;
        }
        
        _insertSubscription();
    }

    /**
     * @dev Unsubscribe the service
     */
    function unsubscribe() public onlySubscriber whenPriceGreaterEqualOne {
        Subscription subs = subscriptionRegistry[msg.sender];
        msg.sender.transfer(subs.deposit);
        subs.deposit = 0;
        subs.hasSubscription = false;
    }
}