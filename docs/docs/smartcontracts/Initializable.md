# Initializable (Initializable.sol)

**Initializable**

Simple helper contract to support initialization outside of the constructor.
To use it, replace the constructor with a function that has the
`isInitializer` modifier.
WARNING: This helper does not support multiple inheritance.
WARNING: It is the developer's responsibility to ensure that an initializer
is actually called.
Use `Migratable` for more complex migration mechanisms.

## Contract Members
**Constants & Variables**

```js
bool public initialized;
```

## Modifiers

- [isInitializer](#isinitializer)

### isInitializer

Modifier to use in the initialization function of a contract.

```js
modifier isInitializer() internal
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|

## Functions

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
