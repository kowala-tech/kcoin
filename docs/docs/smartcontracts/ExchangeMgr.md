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

checks whether the given name is a whitelisted exchange or not

```js
function isExchange(string name) public view
returns(isIndeed bool)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| name | string | exchange name. | 

### isWhitelistedExchange

checks whether the given name is a whitelisted exchange or not

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
