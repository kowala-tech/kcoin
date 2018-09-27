# Ownable (Ownable.sol)

**Ownable**

The Ownable contract has an owner address, and provides basic authorization control
functions, this simplifies the implementation of "user permissions".

## Contract Members
**Constants & Variables**

```js
address public owner;
```

**Events**

```js
event OwnershipRenounced(address indexed previousOwner);
event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);
```

## Modifiers

- [onlyOwner](#onlyowner)

### onlyOwner

Throws if called by any account other than the owner.

```js
modifier onlyOwner() internal
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|

## Functions

- [renounceOwnership](#renounceownership)
- [transferOwnership](#transferownership)
- [_transferOwnership](#_transferownership)

### renounceOwnership

Renouncing to ownership will leave the contract without an owner.
It will not be possible to call the functions with the `onlyOwner`
modifier anymore.

```js
function renounceOwnership() public onlyOwner
```

### transferOwnership

Allows the current owner to transfer control of the contract to a newOwner.

```js
function transferOwnership(address _newOwner) public onlyOwner
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| _newOwner | address | The address to transfer ownership to. | 

### _transferOwnership

Transfers control of the contract to a newOwner.

```js
function _transferOwnership(address _newOwner) internal
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| _newOwner | address | The address to transfer ownership to. | 

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
