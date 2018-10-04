# TokenMock.sol

**TokenMock**

## Contract Members
**Constants & Variables**

```js
//public members
uint256 public totalSupply;
//private members
mapping(address => uint256) private balances;
```

**Events**

```js
event Transfer(address indexed from, address indexed to, uint256 value, bytes indexed data);
```

## Functions

- [balanceOf](#balanceof)
- [name](#name)
- [symbol](#symbol)
- [decimals](#decimals)
- [totalSupply](#totalsupply)
- [transfer](#transfer)
- [transfer](#transfer)
- [transfer](#transfer)

### balanceOf

```js
function balanceOf(address who) public view
returns(uint256)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| who | address |  | 

### name

```js
function name() public view
returns(_name string)
```

### symbol

```js
function symbol() public view
returns(_symbol string)
```

### decimals

```js
function decimals() public view
returns(_decimals uint8)
```

### totalSupply

```js
function totalSupply() public view
returns(_supply uint256)
```

### transfer

```js
function transfer(address to, uint256 value) public
returns(ok bool)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| to | address |  | 
| value | uint256 |  | 

### transfer

```js
function transfer(address to, uint256 value, bytes data) public
returns(ok bool)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| to | address |  | 
| value | uint256 |  | 
| data | bytes |  | 

### transfer

```js
function transfer(address to, uint256 value, bytes data, string custom_fallback) public
returns(ok bool)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| to | address |  | 
| value | uint256 |  | 
| data | bytes |  | 
| custom_fallback | string |  | 

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
