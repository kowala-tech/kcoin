# Token.sol

**contract Token is [KRC223](KRC223.md), [Initializable](Initializable.md)**

**Token**

## Contract Members
**Constants & Variables**

```js
//internal members
mapping(address => uint256) internal balances;
//public members
string public name;
string public symbol;
uint8 public decimals;
uint256 public totalSupply;
```

## Functions

- [name](#name)
- [symbol](#symbol)
- [decimals](#decimals)
- [totalSupply](#totalsupply)
- [transfer](#transfer)
- [transfer](#transfer)
- [transfer](#transfer)
- [isContract](#iscontract)
- [transferToAddress](#transfertoaddress)
- [transferToContract](#transfertocontract)
- [balanceOf](#balanceof)

### name

:small_red_triangle: overrides [KRC223.name](KRC223.md#name)

```js
function name() public view
returns(_name string)
```

### symbol

:small_red_triangle: overrides [KRC223.symbol](KRC223.md#symbol)

```js
function symbol() public view
returns(_symbol string)
```

### decimals

:small_red_triangle: overrides [KRC223.decimals](KRC223.md#decimals)

```js
function decimals() public view
returns(_decimals uint8)
```

### totalSupply

:small_red_triangle: overrides [KRC223.totalSupply](KRC223.md#totalsupply)

```js
function totalSupply() public view
returns(_totalSupply uint256)
```

### transfer

:small_red_triangle: overrides [KRC223.transfer](KRC223.md#transfer)

```js
function transfer(address _to, uint256 _value, bytes _data) public
returns(success bool)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| _to | address |  | 
| _value | uint256 |  | 
| _data | bytes |  | 

### transfer

:small_red_triangle: overrides [KRC223.transfer](KRC223.md#transfer)

```js
function transfer(address _to, uint256 _value, bytes _data, string _custom_fallback) public
returns(success bool)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| _to | address |  | 
| _value | uint256 |  | 
| _data | bytes |  | 
| _custom_fallback | string |  | 

### transfer

:small_red_triangle: overrides [KRC223.transfer](KRC223.md#transfer)

```js
function transfer(address _to, uint256 _value) public
returns(success bool)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| _to | address |  | 
| _value | uint256 |  | 

### isContract

```js
function isContract(address _addr) private view
returns(is_contract bool)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| _addr | address |  | 

### transferToAddress

```js
function transferToAddress(address _to, uint256 _value, bytes _data) private
returns(success bool)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| _to | address |  | 
| _value | uint256 |  | 
| _data | bytes |  | 

### transferToContract

```js
function transferToContract(address _to, uint256 _value, bytes _data) private
returns(success bool)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| _to | address |  | 
| _value | uint256 |  | 
| _data | bytes |  | 

### balanceOf

:small_red_triangle: overrides [KRC223.balanceOf](KRC223.md#balanceof)

```js
function balanceOf(address _owner) public view
returns(balance uint256)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| _owner | address |  | 

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
