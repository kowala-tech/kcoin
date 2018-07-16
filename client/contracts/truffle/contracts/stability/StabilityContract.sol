pragma solidity 0.4.24;

import "openzeppelin-solidity/contracts/lifecycle/Pausable.sol";

contract StabilityContract is Pausable {
    
    uint public numSubscribers = 0;
    uint public numSubscriptions = 0;
    uint public minDeposit;
    uint public totalBalance = 0;

    struct Deposit {
        uint amount;
        uint availableAt;
    }

    struct Subscription {
        uint index;
        bool isActive;
        Deposit[] deposits;
    }

    struct Subscriber {
        bool isSubscriber;
        Subscription[] subscriptions;
    }

    mapping (address => Subscription) private subscriptionRegistry;

    address[] subscriptionPool;

    modifier onlyWithMinDeposit {
        require(msg.value >= minDeposit);
        _;
    }

    function StabilityContract(
        uint _minDeposit)
    public {
        require(_minDeposit >= 0);

        minDeposit = _minDeposit;
    } 

    function _insertSubscription(address identity, uint deposit) private {
        
        // @TODO (case if account already exists)
        //Subscription subscription = subscriptionRegistry[identity];

    }

    function subscribe() public payable whenNotPaused onlyWithMinDeposit {
        _insertSubscription(msg.sender, msg.value);
    }

    function releaseDeposits() public whenNotPaused {

    }
}