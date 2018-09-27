# PublicResolver.sol

**contract PublicResolver is [Initializable](Initializable.md)**

**PublicResolver**

## Structs
### PublicKey

```js
struct PublicKey {
  bytes32 x,
  bytes32 y
}
```

### Record

```js
struct Record {
  address addr,
  bytes32 content,
  string name,
  struct PublicResolver.PublicKey pubkey,
  mapping(string => string) text,
  mapping(uint256 => bytes) abis,
  bytes multihash
}
```

## Contract Members
**Constants & Variables**

```js
bytes4 internal constant INTERFACE_META_ID;
bytes4 internal constant ADDR_INTERFACE_ID;
bytes4 internal constant CONTENT_INTERFACE_ID;
bytes4 internal constant NAME_INTERFACE_ID;
bytes4 internal constant ABI_INTERFACE_ID;
bytes4 internal constant PUBKEY_INTERFACE_ID;
bytes4 internal constant TEXT_INTERFACE_ID;
bytes4 internal constant MULTIHASH_INTERFACE_ID;
contract KNS internal kns;
mapping(bytes32 => struct PublicResolver.Record) internal records;
```

**Events**

```js
event AddrChanged(bytes32 indexed node, address a);
event ContentChanged(bytes32 indexed node, bytes32 hash);
event NameChanged(bytes32 indexed node, string name);
event ABIChanged(bytes32 indexed node, uint256 indexed contentType);
event PubkeyChanged(bytes32 indexed node, bytes32 x, bytes32 y);
event TextChanged(bytes32 indexed node, string indexedKey, string key);
event MultihashChanged(bytes32 indexed node, bytes hash);
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
- [setAddr](#setaddr)
- [setContent](#setcontent)
- [setMultihash](#setmultihash)
- [setName](#setname)
- [setABI](#setabi)
- [setPubkey](#setpubkey)
- [setText](#settext)
- [text](#text)
- [pubkey](#pubkey)
- [ABI](#abi)
- [name](#name)
- [content](#content)
- [multihash](#multihash)
- [addr](#addr)
- [supportsInterface](#supportsinterface)

### initialize

initialize function for Proxy Pattern.

```js
function initialize(KNS knsAddr) public isInitializer
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| knsAddr | KNS | The address of the KNS registry. | 

### setAddr

```js
function setAddr(bytes32 node, address addr) public only_owner
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| node | bytes32 | The node to update. | 
| addr | address | The address to set. | 

### setContent

```js
function setContent(bytes32 node, bytes32 hash) public only_owner
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| node | bytes32 | The node to update. | 
| hash | bytes32 | The content hash to set | 

### setMultihash

```js
function setMultihash(bytes32 node, bytes hash) public only_owner
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| node | bytes32 | The node to update. | 
| hash | bytes | The multihash to set | 

### setName

```js
function setName(bytes32 node, string name) public only_owner
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| node | bytes32 | The node to update. | 
| name | string | The name to set. | 

### setABI

```js
function setABI(bytes32 node, uint256 contentType, bytes data) public only_owner
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| node | bytes32 | The node to update. | 
| contentType | uint256 | The content type of the ABI | 
| data | bytes | The ABI data. | 

### setPubkey

```js
function setPubkey(bytes32 node, bytes32 x, bytes32 y) public only_owner
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| node | bytes32 | The KNS node to query | 
| x | bytes32 | the X coordinate of the curve point for the public key. | 
| y | bytes32 | the Y coordinate of the curve point for the public key. | 

### setText

```js
function setText(bytes32 node, string key, string value) public only_owner
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| node | bytes32 | The node to update. | 
| key | string | The key to set. | 
| value | string | The text data value to set. | 

### text

```js
function text(bytes32 node, string key) public view
returns(string)
```

**Returns**

The associated text data.

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| node | bytes32 | The KNS node to query. | 
| key | string | The text data key to query. | 

### pubkey

```js
function pubkey(bytes32 node) public view
returns(x bytes32, y bytes32)
```

**Returns**

x, y the X and Y coordinates of the curve point for the public key.

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| node | bytes32 | The KNS node to query | 

### ABI

```js
function ABI(bytes32 node, uint256 contentTypes) public view
returns(contentType uint256, data bytes)
```

**Returns**

contentType The content type of the return value

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| node | bytes32 | The KNS node to query | 
| contentTypes | uint256 | A bitwise OR of the ABI formats accepted by the caller. | 

### name

```js
function name(bytes32 node) public view
returns(string)
```

**Returns**

The associated name.

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| node | bytes32 | The KNS node to query. | 

### content

```js
function content(bytes32 node) public view
returns(bytes32)
```

**Returns**

The associated content hash.

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| node | bytes32 | The KNS node to query. | 

### multihash

```js
function multihash(bytes32 node) public view
returns(bytes)
```

**Returns**

The associated multihash.

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| node | bytes32 | The KNS node to query. | 

### addr

```js
function addr(bytes32 node) public view
returns(address)
```

**Returns**

The associated address.

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| node | bytes32 | The KNS node to query. | 

### supportsInterface

```js
function supportsInterface(bytes4 interfaceID) public pure
returns(bool)
```

**Returns**

True if the contract implements the requested interface.

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| interfaceID | bytes4 | The ID of the interface to check for. | 

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
