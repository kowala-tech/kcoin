# FIFSRegistrar.sol

**contract FIFSRegistrar is [Initializable](Initializable.md)**

**FIFSRegistrar**

## Contract Members
**Constants & Variables**

```js
contract KNS internal kns;
bytes32 internal rootNode;
```

## Modifiers

- [only_owner](#only_owner)

### only_owner

```js
modifier only_owner(bytes32 subnode) internal
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| subnode | bytes32 |  | 

## Functions

- [initialize](#initialize)
- [register](#register)

### initialize

initialize function for Proxy Pattern.

```js
function initialize(KNS knsAddr, bytes32 node) public isInitializer
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| knsAddr | KNS | The address of the KNS registry. | 
| node | bytes32 | The node that this registrar administers. | 

### register

```js
function register(bytes32 subnode, address owner) public only_owner
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| subnode | bytes32 | The hash of the label to register. | 
| owner | address | The address of the new owner. | 

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
