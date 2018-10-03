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

- [KNSRegistryV1](KNSRegistryV1.md)
- [ValidatorMgr](ValidatorMgr.md)
- [Math](Math.md)
- [NameHash](NameHash.md)
- [SystemVars](SystemVars.md)
- [Stability](Stability.md)
- [Token](Token.md)
- [TokenMock](TokenMock.md)
- [TokenReceiver](TokenReceiver.md)
- [SafeMath](SafeMath.md)
- [CappedToken](CappedToken.md)
- [FIFSRegistrar](FIFSRegistrar.md)
- [Initializable](Initializable.md)
- [KNSRegistry](KNSRegistry.md)
- [ExchangeMgr](ExchangeMgr.md)
- [KRC223](KRC223.md)
- [PublicResolver](PublicResolver.md)
- [MultiSigWallet](MultiSigWallet.md)
- [DomainResolver](DomainResolver.md)
- [PriceProvider](PriceProvider.md)
- [BalanceContract](BalanceContract.md)
- [MiningToken](MiningToken.md)
- [MintableToken](MintableToken.md)
- [strings](strings.md)
- [Pausable](Pausable.md)
- [Migrations](Migrations.md)
- [Ownable](Ownable.md)
- [Consensus](Consensus.md)
- [OracleMgr](OracleMgr.md)
- [ConsensusMock](ConsensusMock.md)
- [DomainResolverMock](DomainResolverMock.md)
- [KNS](KNS.md)
