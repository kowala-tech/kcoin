# strings.sol

**strings**

## Structs
### slice

```js
struct slice {
  uint256 _len,
  uint256 _ptr
}
```

## Functions

- [memcpy](#memcpy)
- [toSlice](#toslice)
- [len](#len)
- [toSliceB32](#tosliceb32)
- [copy](#copy)
- [toString](#tostring)
- [len](#len)
- [empty](#empty)
- [compare](#compare)
- [equals](#equals)
- [nextRune](#nextrune)
- [nextRune](#nextrune)
- [ord](#ord)
- [keccak](#keccak)
- [startsWith](#startswith)
- [beyond](#beyond)
- [endsWith](#endswith)
- [until](#until)
- [findPtr](#findptr)
- [rfindPtr](#rfindptr)
- [find](#find)
- [rfind](#rfind)
- [split](#split)
- [split](#split)
- [rsplit](#rsplit)
- [rsplit](#rsplit)
- [count](#count)
- [contains](#contains)
- [concat](#concat)
- [join](#join)

### memcpy

```js
function memcpy(uint256 dest, uint256 src, uint256 len) private
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| dest | uint256 |  | 
| src | uint256 |  | 
| len | uint256 |  | 

### toSlice

```js
function toSlice(string self) internal
returns(struct strings.slice)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| self | string |  | 

### len

```js
function len(bytes32 self) internal
returns(uint256)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| self | bytes32 |  | 

### toSliceB32

```js
function toSliceB32(bytes32 self) internal
returns(ret struct strings.slice)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| self | bytes32 |  | 

### copy

```js
function copy(struct strings.slice self) internal
returns(struct strings.slice)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| self | struct strings.slice |  | 

### toString

```js
function toString(struct strings.slice self) internal
returns(string)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| self | struct strings.slice |  | 

### len

```js
function len(struct strings.slice self) internal
returns(uint256)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| self | struct strings.slice |  | 

### empty

```js
function empty(struct strings.slice self) internal
returns(bool)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| self | struct strings.slice |  | 

### compare

```js
function compare(struct strings.slice self, struct strings.slice other) internal
returns(int256)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| self | struct strings.slice |  | 
| other | struct strings.slice |  | 

### equals

```js
function equals(struct strings.slice self, struct strings.slice other) internal
returns(bool)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| self | struct strings.slice |  | 
| other | struct strings.slice |  | 

### nextRune

```js
function nextRune(struct strings.slice self, struct strings.slice rune) internal
returns(struct strings.slice)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| self | struct strings.slice |  | 
| rune | struct strings.slice |  | 

### nextRune

```js
function nextRune(struct strings.slice self) internal
returns(ret struct strings.slice)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| self | struct strings.slice |  | 

### ord

```js
function ord(struct strings.slice self) internal
returns(ret uint256)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| self | struct strings.slice |  | 

### keccak

```js
function keccak(struct strings.slice self) internal
returns(ret bytes32)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| self | struct strings.slice |  | 

### startsWith

```js
function startsWith(struct strings.slice self, struct strings.slice needle) internal
returns(bool)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| self | struct strings.slice |  | 
| needle | struct strings.slice |  | 

### beyond

```js
function beyond(struct strings.slice self, struct strings.slice needle) internal
returns(struct strings.slice)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| self | struct strings.slice |  | 
| needle | struct strings.slice |  | 

### endsWith

```js
function endsWith(struct strings.slice self, struct strings.slice needle) internal
returns(bool)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| self | struct strings.slice |  | 
| needle | struct strings.slice |  | 

### until

```js
function until(struct strings.slice self, struct strings.slice needle) internal
returns(struct strings.slice)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| self | struct strings.slice |  | 
| needle | struct strings.slice |  | 

### findPtr

```js
function findPtr(uint256 selflen, uint256 selfptr, uint256 needlelen, uint256 needleptr) private
returns(uint256)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| selflen | uint256 |  | 
| selfptr | uint256 |  | 
| needlelen | uint256 |  | 
| needleptr | uint256 |  | 

### rfindPtr

```js
function rfindPtr(uint256 selflen, uint256 selfptr, uint256 needlelen, uint256 needleptr) private
returns(uint256)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| selflen | uint256 |  | 
| selfptr | uint256 |  | 
| needlelen | uint256 |  | 
| needleptr | uint256 |  | 

### find

```js
function find(struct strings.slice self, struct strings.slice needle) internal
returns(struct strings.slice)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| self | struct strings.slice |  | 
| needle | struct strings.slice |  | 

### rfind

```js
function rfind(struct strings.slice self, struct strings.slice needle) internal
returns(struct strings.slice)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| self | struct strings.slice |  | 
| needle | struct strings.slice |  | 

### split

```js
function split(struct strings.slice self, struct strings.slice needle, struct strings.slice token) internal
returns(struct strings.slice)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| self | struct strings.slice |  | 
| needle | struct strings.slice |  | 
| token | struct strings.slice |  | 

### split

```js
function split(struct strings.slice self, struct strings.slice needle) internal
returns(token struct strings.slice)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| self | struct strings.slice |  | 
| needle | struct strings.slice |  | 

### rsplit

```js
function rsplit(struct strings.slice self, struct strings.slice needle, struct strings.slice token) internal
returns(struct strings.slice)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| self | struct strings.slice |  | 
| needle | struct strings.slice |  | 
| token | struct strings.slice |  | 

### rsplit

```js
function rsplit(struct strings.slice self, struct strings.slice needle) internal
returns(token struct strings.slice)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| self | struct strings.slice |  | 
| needle | struct strings.slice |  | 

### count

```js
function count(struct strings.slice self, struct strings.slice needle) internal
returns(count uint256)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| self | struct strings.slice |  | 
| needle | struct strings.slice |  | 

### contains

```js
function contains(struct strings.slice self, struct strings.slice needle) internal
returns(bool)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| self | struct strings.slice |  | 
| needle | struct strings.slice |  | 

### concat

```js
function concat(struct strings.slice self, struct strings.slice other) internal
returns(string)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| self | struct strings.slice |  | 
| other | struct strings.slice |  | 

### join

```js
function join(struct strings.slice self, struct strings.slice[] parts) internal
returns(string)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| self | struct strings.slice |  | 
| parts | struct strings.slice[] |  | 

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
