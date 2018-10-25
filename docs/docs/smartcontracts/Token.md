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
- [PublicResolver](PublicResolver.md)
- [SafeMath](SafeMath.md)
- [strings](strings.md)
- [SystemVars](SystemVars.md)
- [Token](Token.md)
- [TokenMock](TokenMock.md)
- [TokenReceiver](TokenReceiver.md)
- [ValidatorMgr](ValidatorMgr.md)
