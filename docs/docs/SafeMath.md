# SafeMath.sol

**SafeMath**

## Functions

- [mul](#mul)
- [div](#div)
- [sub](#sub)
- [add](#add)

### mul

Multiplies two numbers, throws on overflow.

```js
function mul(uint256 _a, uint256 _b) internal pure
returns(c uint256)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| _a | uint256 |  | 
| _b | uint256 |  | 

### div

Integer division of two numbers, truncating the quotient.

```js
function div(uint256 _a, uint256 _b) internal pure
returns(uint256)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| _a | uint256 |  | 
| _b | uint256 |  | 

### sub

Subtracts two numbers, throws on overflow (i.e. if subtrahend is greater than minuend).

```js
function sub(uint256 _a, uint256 _b) internal pure
returns(uint256)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| _a | uint256 |  | 
| _b | uint256 |  | 

### add

Adds two numbers, throws on overflow.

```js
function add(uint256 _a, uint256 _b) internal pure
returns(c uint256)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| _a | uint256 |  | 
| _b | uint256 |  | 

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
