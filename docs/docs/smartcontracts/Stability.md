# Stability contract supports network utility (Stability.sol)

**contract Stability is [Pausable](Pausable.md), [Initializable](Initializable.md)**

**Stability**

## Structs
### Subscription

```js
struct Subscription {
  uint256 index,
  bool hasSubscription,
  uint256 deposit
}
```

## Contract Members
**Constants & Variables**

```js
//internal members
uint256 internal constant ONE;
contract PriceProvider internal priceProvider;
//public members
uint256 public minDeposit;
//private members
mapping(address => struct Stability.Subscription) private subscriptionRegistry;
address[] private subscriptions;
```

## Modifiers

- [onlySubscriber](#onlysubscriber)
- [onlyWithMinDeposit](#onlywithmindeposit)
- [whenPriceGreaterEqualOne](#whenpricegreaterequalone)

### onlySubscriber

```js
modifier onlySubscriber() internal
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|

### onlyWithMinDeposit

```js
modifier onlyWithMinDeposit() internal
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|

### whenPriceGreaterEqualOne

```js
modifier whenPriceGreaterEqualOne() internal
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|

## Functions

- [initialize](#initialize)
- [getSubscriptionCount](#getsubscriptioncount)
- [getSubscriptionAtIndex](#getsubscriptionatindex)
- [_hasSubscription](#_hassubscription)
- [_insertSubscription](#_insertsubscription)
- [subscribe](#subscribe)
- [unsubscribe](#unsubscribe)

### initialize

```js
function initialize(uint256 _minDeposit, address _priceProviderAddr) public isInitializer
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| _minDeposit | uint256 | minimum deposit required to subscribe to the service | 
| _priceProviderAddr | address | address of system variables contract | 

### getSubscriptionCount

```js
function getSubscriptionCount() public view
returns(count uint256)
```

### getSubscriptionAtIndex

```js
function getSubscriptionAtIndex(uint256 index) public view
returns(code address, deposit uint256)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| index | uint256 |  | 

### _hasSubscription

```js
function _hasSubscription(address identity) private view
returns(isIndeed bool)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| identity | address |  | 

### _insertSubscription

```js
function _insertSubscription() private onlyWithMinDeposit
```

### subscribe

Subscribe to the stability contract service

```js
function subscribe() public payable payable whenNotPaused
```

### unsubscribe

Unsubscribe the service

```js
function unsubscribe() public onlySubscriber whenPriceGreaterEqualOne
```

## Contracts

- [BalanceContract](BalanceContract.md)
- [CappedToken](CappedToken.md)
- [Consensus](Consensus.md)
- [ConsensusMock](ConsensusMock.md)
- [DomainResolver](DomainResolver.md)
- [DomainResolverMock](DomainResolverMock.md)
- [ExchangeMgr](ExchangeMgr.md)
- [FIFSRegistrar](FIFSRegistrar.md)
- [Initializable](Initializable.md)
- [KNS](KNS.md)
- [KNSRegistry](KNSRegistry.md)
- [KNSRegistryV1](KNSRegistryV1.md)
- [KRC223](KRC223.md)
- [Math](Math.md)
- [Migrations](Migrations.md)
- [MiningToken](MiningToken.md)
- [MintableToken](MintableToken.md)
- [MultiSigWallet](MultiSigWallet.md)
- [NameHash](NameHash.md)
- [OracleMgr](OracleMgr.md)
- [Ownable](Ownable.md)
- [Pausable](Pausable.md)
- [PriceProvider](PriceProvider.md)
- [PublicResolver](PublicResolver.md)
- [SafeMath](SafeMath.md)
- [Stability](Stability.md)
- [strings](strings.md)
- [SystemVars](SystemVars.md)
- [Token](Token.md)
- [TokenMock](TokenMock.md)
- [TokenReceiver](TokenReceiver.md)
- [ValidatorMgr](ValidatorMgr.md)
