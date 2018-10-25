# MiningToken.sol

**contract MiningToken is [CappedToken](CappedToken.md)**

**MiningToken**

## Constructor

```js
constructor(uint256 _cap) public
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| _cap | uint256 |  | 

## Functions

- [initialize](#initialize)
- [getOwner](#getowner)

### initialize

```js
function initialize(string _name, string _symbol, uint256 _cap, uint8 _decimals) public isInitializer
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| _name | string |  | 
| _symbol | string |  | 
| _cap | uint256 |  | 
| _decimals | uint8 |  | 

### getOwner

```js
function getOwner() public
returns(address)
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
- [PublicResolver](PublicResolver.md)
- [SafeMath](SafeMath.md)
- [strings](strings.md)
- [SystemVars](SystemVars.md)
- [Token](Token.md)
- [TokenMock](TokenMock.md)
- [TokenReceiver](TokenReceiver.md)
- [ValidatorMgr](ValidatorMgr.md)
