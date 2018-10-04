# Exchange Manager contract (ExchangeMgr.sol)

**contract ExchangeMgr is [Pausable](Pausable.md), [Initializable](Initializable.md)**

**ExchangeMgr**

## Structs
### Exchange

```js
struct Exchange {
  uint256 index,
  bool isExchange,
  bool isWhitelisted
}
```

## Contract Members
**Constants & Variables**

```js
//private members
mapping(string => struct ExchangeMgr.Exchange) private exchangeRegistry;
//public members
string[] public whitelist;
```

**Events**

```js
event Whitelisted(string exchange);
event Blacklisted(string exchange);
event Addition(string exchange);
event Removal(string exchange);
```

## Modifiers

- [onlyNewCandidate](#onlynewcandidate)
- [onlyExchange](#onlyexchange)
- [onlyWhitelistedExchange](#onlywhitelistedexchange)
- [onlyBlacklistedExchange](#onlyblacklistedexchange)

### onlyNewCandidate

```js
modifier onlyNewCandidate(string name) internal
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| name | string |  | 

### onlyExchange

```js
modifier onlyExchange(string name) internal
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| name | string |  | 

### onlyWhitelistedExchange

```js
modifier onlyWhitelistedExchange(string name) internal
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| name | string |  | 

### onlyBlacklistedExchange

```js
modifier onlyBlacklistedExchange(string name) internal
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| name | string |  | 

## Functions

- [addExchange](#addexchange)
- [removeExchange](#removeexchange)
- [whitelistExchange](#whitelistexchange)
- [blacklistExchange](#blacklistexchange)
- [isExchange](#isexchange)
- [isWhitelistedExchange](#iswhitelistedexchange)
- [isBlacklistedExchange](#isblacklistedexchange)
- [getWhitelistedExchangeCount](#getwhitelistedexchangecount)
- [getWhitelistedExchangeAtIndex](#getwhitelistedexchangeatindex)

### addExchange

Adds and whitelists an exchange.

```js
function addExchange(string name) public whenNotPaused onlyOwner onlyNewCandidate
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| name | string | exchange name. | 

### removeExchange

Removes an exchange.

```js
function removeExchange(string name) public whenNotPaused onlyOwner onlyExchange
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| name | string | exchange name. | 

### whitelistExchange

Whitelists an exchange.

```js
function whitelistExchange(string name) public whenNotPaused onlyOwner onlyBlacklistedExchange
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| name | string | exchange name. | 

### blacklistExchange

Blacklists an exchange.

```js
function blacklistExchange(string name) public whenNotPaused onlyOwner onlyWhitelistedExchange
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| name | string | exchange name. | 

### isExchange

checks whether the given name is an whitelisted exchange or not

```js
function isExchange(string name) public view
returns(isIndeed bool)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| name | string | exchange name. | 

### isWhitelistedExchange

checks whether the given name is an whitelisted exchange or not

```js
function isWhitelistedExchange(string name) public view
returns(isIndeed bool)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| name | string | exchange name. | 

### isBlacklistedExchange

checks whether the given name is a blacklisted exchange or not

```js
function isBlacklistedExchange(string name) public view
returns(isIndeed bool)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| name | string | exchange name. | 

### getWhitelistedExchangeCount

get whitelisted exchange count

```js
function getWhitelistedExchangeCount() public view
returns(count uint256)
```

### getWhitelistedExchangeAtIndex

get whitelisted exchange information

```js
function getWhitelistedExchangeAtIndex(uint256 index) public view
returns(name string)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| index | uint256 | index of a given exchange in the whitelist | 

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
