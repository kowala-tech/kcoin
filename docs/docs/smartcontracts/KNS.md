# KNS.sol

**KNS**

**Events**

```js
event NewOwner(bytes32 indexed node, bytes32 indexed label, address owner);
event Transfer(bytes32 indexed node, address owner);
event NewResolver(bytes32 indexed node, address resolver);
event NewTTL(bytes32 indexed node, uint64 ttl);
```

## Functions

- [setSubnodeOwner](#setsubnodeowner)
- [setResolver](#setresolver)
- [setOwner](#setowner)
- [setTTL](#setttl)
- [owner](#owner)
- [resolver](#resolver)
- [ttl](#ttl)

### setSubnodeOwner

```js
function setSubnodeOwner(bytes32 node, bytes32 label, address owner) external
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| node | bytes32 |  | 
| label | bytes32 |  | 
| owner | address |  | 

### setResolver

```js
function setResolver(bytes32 node, address resolver) external
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| node | bytes32 |  | 
| resolver | address |  | 

### setOwner

```js
function setOwner(bytes32 node, address owner) external
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| node | bytes32 |  | 
| owner | address |  | 

### setTTL

```js
function setTTL(bytes32 node, uint64 ttl) external
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| node | bytes32 |  | 
| ttl | uint64 |  | 

### owner

```js
function owner(bytes32 node) external view
returns(address)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| node | bytes32 |  | 

### resolver

```js
function resolver(bytes32 node) external view
returns(address)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| node | bytes32 |  | 

### ttl

```js
function ttl(bytes32 node) external view
returns(uint64)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| node | bytes32 |  | 

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
