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
