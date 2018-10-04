# Validator Manager for PoS consensus (ValidatorMgr.sol)

**contract ValidatorMgr is [Pausable](Pausable.md), [Initializable](Initializable.md)**

**ValidatorMgr**

## Structs
### Deposit

```js
struct Deposit {
  uint256 amount,
  uint256 availableAt
}
```

### Validator

```js
struct Validator {
  uint256 index,
  bool isValidator,
  bool isGenesis,
  struct ValidatorMgr.Deposit[] deposits
}
```

### TKN

```js
struct TKN {
  address sender,
  uint256 value
}
```

## Contract Members
**Constants & Variables**

```js
//public members
uint256 public baseDeposit;
uint256 public maxNumValidators;
uint256 public freezePeriod;
bytes32 public validatorsChecksum;
uint256 public superNodeAmount;
contract DomainResolver public knsResolver;
//internal members
bytes32 internal nodeNamehash;
address[] internal validatorPool;
struct ValidatorMgr.TKN internal tkn;
//private members
mapping(address => struct ValidatorMgr.Validator) private validatorRegistry;
```

## Modifiers

- [onlyWithMinDeposit](#onlywithmindeposit)
- [onlyValidator](#onlyvalidator)
- [onlyNewCandidate](#onlynewcandidate)

### onlyWithMinDeposit

```js
modifier onlyWithMinDeposit() internal
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|

### onlyValidator

```js
modifier onlyValidator() internal
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

## Functions

- [initialize](#initialize)
- [isGenesisValidator](#isgenesisvalidator)
- [isValidator](#isvalidator)
- [isSuperNode](#issupernode)
- [getValidatorCount](#getvalidatorcount)
- [getValidatorAtIndex](#getvalidatoratindex)
- [_hasAvailability](#_hasavailability)
- [getMinimumDeposit](#getminimumdeposit)
- [_updateChecksum](#_updatechecksum)
- [_insertValidator](#_insertvalidator)
- [setBaseDeposit](#setbasedeposit)
- [setMaxValidators](#setmaxvalidators)
- [_deleteValidator](#_deletevalidator)
- [_deleteSmallestBidder](#_deletesmallestbidder)
- [getDepositCount](#getdepositcount)
- [getDepositAtIndex](#getdepositatindex)
- [_registerValidator](#_registervalidator)
- [deregisterValidator](#deregistervalidator)
- [_removeDeposits](#_removedeposits)
- [releaseDeposits](#releasedeposits)
- [registerValidator](#registervalidator)

### initialize

```js
function initialize(uint256 _baseDeposit, uint256 _maxNumValidators, uint256 _freezePeriod, uint256 _superNodeAmount, address _resolverAddr) public isInitializer
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| _baseDeposit | uint256 | base deposit for Validator | 
| _maxNumValidators | uint256 | Maximum numbers of Validators. | 
| _freezePeriod | uint256 | Freeze period for Validator's deposits. | 
| _superNodeAmount | uint256 | Amount required to be considered a super node. | 
| _resolverAddr | address | Address of KNS Resolver. | 

### isGenesisValidator

Checks if given address is Genesis Validator

```js
function isGenesisValidator(address code) public view
returns(isIndeed bool)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| code | address | Address of an Validator. | 

### isValidator

Checks if given address is Validator

```js
function isValidator(address code) public view
returns(isIndeed bool)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| code | address | Address of a Validator to check. | 

### isSuperNode

Checks if given address is super node

```js
function isSuperNode(address code) public view
returns(isIndeed bool)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| code | address | Address of a supernode to check. | 

### getValidatorCount

Get Validator count

```js
function getValidatorCount() public view
returns(count uint256)
```

### getValidatorAtIndex

Get Validator information

```js
function getValidatorAtIndex(uint256 index) public view
returns(code address, deposit uint256)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| index | uint256 | index of an Validator to check. | 

### _hasAvailability

```js
function _hasAvailability() public view
returns(available bool)
```

### getMinimumDeposit

returns the base deposit if there are positions available or
the current smallest deposit required if there aren't positions available.

```js
function getMinimumDeposit() public view
returns(deposit uint256)
```

### _updateChecksum

updates the checksum

```js
function _updateChecksum() private
```

### _insertValidator

Add new validator

```js
function _insertValidator(address code, uint256 deposit) private
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| code | address | Address of an Validator. | 
| deposit | uint256 | amount to deposit | 

### setBaseDeposit

Set new base deposit for Validators

```js
function setBaseDeposit(uint256 deposit) public onlyOwner
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| deposit | uint256 |  | 

### setMaxValidators

Set maximum numbers of Validators

```js
function setMaxValidators(uint256 max) public onlyOwner
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| max | uint256 | number of max Validators | 

### _deleteValidator

Delete Validator

```js
function _deleteValidator(address account) private
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| account | address | address of a Validator | 

### _deleteSmallestBidder

removes the Validator with the smallest deposit

```js
function _deleteSmallestBidder() private
```

### getDepositCount

Get deposit count

```js
function getDepositCount() public view
returns(count uint256)
```

### getDepositAtIndex

Get deposit at given index

```js
function getDepositAtIndex(uint256 index) public view
returns(amount uint256, availableAt uint256)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| index | uint256 | index of a Validator to get deposit | 

### _registerValidator

Register new Validator

```js
function _registerValidator() private whenNotPaused onlyNewCandidate onlyWithMinDeposit
```

### deregisterValidator

deregister Validator

```js
function deregisterValidator() public whenNotPaused onlyValidator
```

### _removeDeposits

remove deposit

```js
function _removeDeposits(address code, uint256 index) private
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| code | address | address of a Validator | 
| index | uint256 | index of a deposit | 

### releaseDeposits

transfers locked deposit(s) back the user account if they are past the freeze period

```js
function releaseDeposits() public whenNotPaused
```

### registerValidator

Register Validator

```js
function registerValidator(address _from, uint256 _value) public
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| _from | address | from address | 
| _value | uint256 | value to send | 

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
