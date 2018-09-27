# Oracle Manager contract (OracleMgr.sol)

**contract OracleMgr is [Pausable](Pausable.md), [Initializable](Initializable.md)**

**OracleMgr**

## Structs
### OraclePrice

```js
struct OraclePrice {
  uint256 price,
  address oracle
}
```

### Oracle

```js
struct Oracle {
  uint256 index,
  bool isOracle,
  bool hasSubmittedPrice
}
```

## Contract Members
**Constants & Variables**

```js
//public members
uint256 public maxNumOracles;
uint256 public syncFrequency;
uint256 public updatePeriod;
uint256 public price;
contract DomainResolver public knsResolver;
//internal members
bytes32 internal nodeNamehash;
//private members
mapping(address => struct OracleMgr.Oracle) private oracleRegistry;
address[] private oraclePool;
struct OracleMgr.OraclePrice[] private prices;
```

## Modifiers

- [onlyOracle](#onlyoracle)
- [onlyNewCandidate](#onlynewcandidate)
- [onlySuperNode](#onlysupernode)
- [onlyOnce](#onlyonce)

### onlyOracle

```js
modifier onlyOracle() internal
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|

### onlyNewCandidate

```js
modifier onlyNewCandidate() internal
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|

### onlySuperNode

```js
modifier onlySuperNode() internal
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|

### onlyOnce

```js
modifier onlyOnce() internal
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|

## Functions

- [initialize](#initialize)
- [isOracle](#isoracle)
- [_hasAvailability](#_hasavailability)
- [_deleteOracle](#_deleteoracle)
- [_insertOracle](#_insertoracle)
- [getOracleCount](#getoraclecount)
- [getOracleAtIndex](#getoracleatindex)
- [getPriceCount](#getpricecount)
- [getPriceAtIndex](#getpriceatindex)
- [registerOracle](#registeroracle)
- [deregisterOracle](#deregisteroracle)
- [submitPrice](#submitprice)

### initialize

```js
function initialize(uint256 _maxNumOracles, uint256 _syncFrequency, uint256 _updatePeriod, address _resolverAddr) public isInitializer
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| _maxNumOracles | uint256 | Maximum numbers of Oracles. | 
| _syncFrequency | uint256 | Synchronize frequency for Oracles. | 
| _updatePeriod | uint256 | Update period. | 
| _resolverAddr | address | Address of KNS Resolver. | 

### isOracle

Checks if given address is Oracle

```js
function isOracle(address identity) public view
returns(isIndeed bool)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| identity | address | Address of an Oracle. | 

### _hasAvailability

Checks availability of OraclePool

```js
function _hasAvailability() private view
returns(available bool)
```

### _deleteOracle

Deletes given oracle

```js
function _deleteOracle(address identity) private
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| identity | address | Address of an Oracle. | 

### _insertOracle

Inserts oracle

```js
function _insertOracle(address identity, uint256 deposit) private
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| identity | address | Address of an Oracle. | 
| deposit | uint256 | Deposit ammount | 

### getOracleCount

Get Oracle count

```js
function getOracleCount() public view
returns(count uint256)
```

### getOracleAtIndex

Get Oracle information

```js
function getOracleAtIndex(uint256 index) public view
returns(code address)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| index | uint256 | index of an Oracle to check. | 

### getPriceCount

Get submissions count

```js
function getPriceCount() public view
returns(count uint256)
```

### getPriceAtIndex

Get submissions information

```js
function getPriceAtIndex(uint256 index) public view
returns(price uint256, oracle address)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| index | uint256 | index of a submission to check. | 

### registerOracle

Registers a new candidate as oracle

```js
function registerOracle() public payable payable whenNotPaused onlyNewCandidate onlySuperNode
```

### deregisterOracle

Deregisters the msg sender from the oracle set

```js
function deregisterOracle() public whenNotPaused onlyOracle
```

### submitPrice

Adds price

```js
function submitPrice(uint256 _price) public whenNotPaused onlyOracle onlyOnce
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| _price | uint256 | price | 

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
