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
