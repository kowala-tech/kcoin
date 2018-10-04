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
