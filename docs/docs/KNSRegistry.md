# KNSRegistry.sol

**contract KNSRegistry is [KNS](KNS.md), [Initializable](Initializable.md)**

**KNSRegistry**

## Structs
### Record

```js
struct Record {
  address owner,
  address resolver,
  uint64 ttl
}
```

## Contract Members
**Constants & Variables**

```js
mapping(bytes32 => struct KNSRegistry.Record) internal records;
```

## Modifiers

- [only_owner](#only_owner)

### only_owner

```js
modifier only_owner(bytes32 node) internal
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| node | bytes32 |  | 

## Functions

- [initialize](#initialize)
- [setOwner](#setowner)
- [setSubnodeOwner](#setsubnodeowner)
- [setResolver](#setresolver)
- [setTTL](#setttl)
- [owner](#owner)
- [resolver](#resolver)
- [ttl](#ttl)

### initialize

initialize function for Proxy Pattern.

```js
function initialize(address _owner) public isInitializer
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| _owner | address | address of the root owner | 

### setOwner

:small_red_triangle: overrides [KNS.setOwner](KNS.md#setowner)

Transfers ownership of a node to a new address. May only be called by the current owner of the node.

```js
function setOwner(bytes32 node, address owner) public only_owner
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| node | bytes32 | The node to transfer ownership of. | 
| owner | address | The address of the new owner. | 

### setSubnodeOwner

:small_red_triangle: overrides [KNS.setSubnodeOwner](KNS.md#setsubnodeowner)

Transfers ownership of a subnode keccak256(node, label) to a new address. May only be called by the owner of the parent node.

```js
function setSubnodeOwner(bytes32 node, bytes32 label, address owner) public only_owner
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| node | bytes32 | The parent node. | 
| label | bytes32 | The hash of the label specifying the subnode. | 
| owner | address | The address of the new owner. | 

### setResolver

:small_red_triangle: overrides [KNS.setResolver](KNS.md#setresolver)

Sets the resolver address for the specified node.

```js
function setResolver(bytes32 node, address resolver) public only_owner
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| node | bytes32 | The node to update. | 
| resolver | address | The address of the resolver. | 

### setTTL

:small_red_triangle: overrides [KNS.setTTL](KNS.md#setttl)

Sets the TTL for the specified node.

```js
function setTTL(bytes32 node, uint64 ttl) public only_owner
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| node | bytes32 | The node to update. | 
| ttl | uint64 | The TTL in seconds. | 

### owner

:small_red_triangle: overrides [KNS.owner](KNS.md#owner)

Returns the address that owns the specified node.

```js
function owner(bytes32 node) public view
returns(address)
```

**Returns**

address of the owner.

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| node | bytes32 | The specified node. | 

### resolver

:small_red_triangle: overrides [KNS.resolver](KNS.md#resolver)

Returns the address of the resolver for the specified node.

```js
function resolver(bytes32 node) public view
returns(address)
```

**Returns**

address of the resolver.

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| node | bytes32 | The specified node. | 

### ttl

:small_red_triangle: overrides [KNS.ttl](KNS.md#ttl)

Returns the TTL of a node, and any records associated with it.

```js
function ttl(bytes32 node) public view
returns(uint64)
```

**Returns**

ttl of the node.

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| node | bytes32 | The specified node. | 

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
