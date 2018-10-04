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

- [DomainResolverMock](DomainResolverMock.md)
- [ValidatorMgr](ValidatorMgr.md)
- [SafeMath](SafeMath.md)
- [MintableToken](MintableToken.md)
- [Ownable](Ownable.md)
- [KRC223](KRC223.md)
- [KNSRegistry](KNSRegistry.md)
- [Token](Token.md)
- [OracleMgr](OracleMgr.md)
- [NameHash](NameHash.md)
- [KNS](KNS.md)
- [Pausable](Pausable.md)
- [TokenMock](TokenMock.md)
- [strings](strings.md)
- [Math](Math.md)
- [BalanceContract](BalanceContract.md)
- [PublicResolver](PublicResolver.md)
- [MultiSigWallet](MultiSigWallet.md)
- [KNSRegistryV1](KNSRegistryV1.md)
- [ExchangeMgr](ExchangeMgr.md)
- [Migrations](Migrations.md)
- [SystemVars](SystemVars.md)
- [FIFSRegistrar](FIFSRegistrar.md)
- [PriceProvider](PriceProvider.md)
- [Initializable](Initializable.md)
- [MiningToken](MiningToken.md)
- [ConsensusMock](ConsensusMock.md)
- [Stability](Stability.md)
- [TokenReceiver](TokenReceiver.md)
- [DomainResolver](DomainResolver.md)
- [CappedToken](CappedToken.md)
- [Consensus](Consensus.md)
