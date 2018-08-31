pragma solidity 0.4.24;

import "openzeppelin-solidity/contracts/lifecycle/Pausable.sol";
import "./PriceProvider.sol";
import "zos-lib/contracts/migrations/Initializable.sol";


/**
 * @title Stability contract supports network utility
 */
contract Stability is Pausable, Initializable{

    uint constant ONE = 1 ether;
    
    uint public minDeposit;
    PriceProvider priceProvider;
    
    struct Subscription {
        uint index;
        bool hasSubscription;
        uint deposit;
    }

    mapping (address => Subscription) private subscriptionRegistry;

    address[] private subscriptions;

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

    /**
     * initialize function for Proxy Pattern.
     * @param _minDeposit minimum deposit required to subscribe to the service
     * @param _priceProviderAddr address of system variables contract
     */
    function initialize(uint _minDeposit, address _priceProviderAddr) isInitializer public {
        minDeposit = _minDeposit;
        priceProvider = PriceProvider(_priceProviderAddr);
    }

    function getSubscriptionCount() public view returns (uint count) {
        return subscriptions.length;
    }

    function getSubscriptionAtIndex(uint index) public view returns (address code, uint deposit) {
        code = subscriptions[index];
        Subscription subs = subscriptionRegistry[code];
        deposit = subs.deposit;
    }

    function _hasSubscription(address identity) private view returns (bool isIndeed) {
        return subscriptionRegistry[identity].hasSubscription;
    }

    function _insertSubscription() private onlyWithMinDeposit {
        Subscription subs = subscriptionRegistry[msg.sender];
        subs.index = subscriptions.push(msg.sender) - 1;
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
        uint rowToDelete = subs.index;
        msg.sender.transfer(subs.deposit);
        delete subscriptionRegistry[msg.sender];

        // replace the deprecated record with the last element
        address keyToMove = subscriptions[subscriptions.length-1]; 
        subscriptions[rowToDelete] = keyToMove;
        subscriptionRegistry[keyToMove].index = rowToDelete;
        subscriptions.length--;
    }
}