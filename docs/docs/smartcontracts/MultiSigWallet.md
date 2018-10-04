# MultiSigWallet.sol

**MultiSigWallet**

## Constructor

Contract constructor sets initial owners and required number of confirmations.

```js
constructor(uint256 , uint256 _required) public
```

**Arguments**

## Structs
### Transaction

```js
struct Transaction {
  address destination,
  uint256 value,
  bytes data,
  bool executed
}
```

## Contract Members
**Constants & Variables**

```js
uint256 public constant MAX_OWNER_COUNT;
mapping(uint256 => struct MultiSigWallet.Transaction) public transactions;
mapping(uint256 => mapping(address => bool)) public confirmations;
mapping(address => bool) public isOwner;
address[] public owners;
uint256 public required;
uint256 public transactionCount;
```

**Events**

```js
event Confirmation(address indexed sender, uint256 indexed transactionId);
event Revocation(address indexed sender, uint256 indexed transactionId);
event Submission(uint256 indexed transactionId);
event Execution(uint256 indexed transactionId);
event ExecutionFailure(uint256 indexed transactionId);
event Deposit(address indexed sender, uint256 value);
event OwnerAddition(address indexed owner);
event OwnerRemoval(address indexed owner);
event RequirementChange(uint256 required);
```

| Name        | Type           | Description  |
| ------------- |------------- | -----|
|  | uint256 | _owners List of initial owners. | 
| _required | uint256 | Number of required confirmations. | 

## Modifiers

- [onlyWallet](#onlywallet)
- [ownerDoesNotExist](#ownerdoesnotexist)
- [ownerExists](#ownerexists)
- [transactionExists](#transactionexists)
- [confirmed](#confirmed)
- [notConfirmed](#notconfirmed)
- [notExecuted](#notexecuted)
- [notNull](#notnull)
- [validRequirement](#validrequirement)

### onlyWallet

```js
modifier onlyWallet() internal
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|

### ownerDoesNotExist

```js
modifier ownerDoesNotExist(address owner) internal
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| owner | address |  | 

### ownerExists

```js
modifier ownerExists(address owner) internal
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| owner | address |  | 

### transactionExists

```js
modifier transactionExists(uint256 transactionId) internal
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| transactionId | uint256 |  | 

### confirmed

```js
modifier confirmed(uint256 transactionId, address owner) internal
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| transactionId | uint256 |  | 
| owner | address |  | 

### notConfirmed

```js
modifier notConfirmed(uint256 transactionId, address owner) internal
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| transactionId | uint256 |  | 
| owner | address |  | 

### notExecuted

```js
modifier notExecuted(uint256 transactionId) internal
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| transactionId | uint256 |  | 

### notNull

```js
modifier notNull(address _address) internal
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| _address | address |  | 

### validRequirement

```js
modifier validRequirement(uint256 ownerCount, uint256 _required) internal
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| ownerCount | uint256 |  | 
| _required | uint256 |  | 

## Functions

- [](#)
- [addOwner](#addowner)
- [removeOwner](#removeowner)
- [replaceOwner](#replaceowner)
- [changeRequirement](#changerequirement)
- [submitTransaction](#submittransaction)
- [confirmTransaction](#confirmtransaction)
- [revokeConfirmation](#revokeconfirmation)
- [executeTransaction](#executetransaction)
- [external_call](#external_call)
- [isConfirmed](#isconfirmed)
- [addTransaction](#addtransaction)
- [getConfirmationCount](#getconfirmationcount)
- [getTransactionCount](#gettransactioncount)
- [getOwners](#getowners)
- [getConfirmations](#getconfirmations)
- [getTransactionIds](#gettransactionids)

### 

Fallback function allows to deposit kUSD.

```js
function () public payable payable
```

### addOwner

Allows to add a new owner. Transaction has to be sent by wallet.

```js
function addOwner(address owner) public onlyWallet ownerDoesNotExist notNull validRequirement
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| owner | address | Address of new owner. | 

### removeOwner

Allows to remove an owner. Transaction has to be sent by wallet.

```js
function removeOwner(address owner) public onlyWallet ownerExists
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| owner | address | Address of owner. | 

### replaceOwner

Allows to replace an owner with a new owner. Transaction has to be sent by wallet.

```js
function replaceOwner(address owner, address newOwner) public onlyWallet ownerExists ownerDoesNotExist
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| owner | address | Address of owner to be replaced. | 
| newOwner | address | Address of new owner. | 

### changeRequirement

Allows to change the number of required confirmations. Transaction has to be sent by wallet.

```js
function changeRequirement(uint256 _required) public onlyWallet validRequirement
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| _required | uint256 | Number of required confirmations. | 

### submitTransaction

Allows an owner to submit and confirm a transaction.

```js
function submitTransaction(address destination, uint256 value, bytes data) public
returns(transactionId uint256)
```

**Returns**

Returns transaction ID.

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| destination | address | Transaction target address. | 
| value | uint256 | Transaction ether value. | 
| data | bytes | Transaction data payload. | 

### confirmTransaction

Allows an owner to confirm a transaction.

```js
function confirmTransaction(uint256 transactionId) public ownerExists transactionExists notConfirmed
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| transactionId | uint256 | Transaction ID. | 

### revokeConfirmation

Allows an owner to revoke a confirmation for a transaction.

```js
function revokeConfirmation(uint256 transactionId) public ownerExists confirmed notExecuted
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| transactionId | uint256 | Transaction ID. | 

### executeTransaction

Allows anyone to execute a confirmed transaction.

```js
function executeTransaction(uint256 transactionId) public ownerExists confirmed notExecuted
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| transactionId | uint256 | Transaction ID. | 

### external_call

```js
function external_call(address destination, uint256 value, uint256 dataLength, bytes data) private
returns(bool)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| destination | address |  | 
| value | uint256 |  | 
| dataLength | uint256 |  | 
| data | bytes |  | 

### isConfirmed

Returns the confirmation status of a transaction.

```js
function isConfirmed(uint256 transactionId) public view
returns(bool)
```

**Returns**

Confirmation status.

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| transactionId | uint256 | Transaction ID. | 

### addTransaction

Adds a new transaction to the transaction mapping, if transaction does not exist yet.

```js
function addTransaction(address destination, uint256 value, bytes data) internal notNull
returns(transactionId uint256)
```

**Returns**

Returns transaction ID.

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| destination | address | Transaction target address. | 
| value | uint256 | Transaction ether value. | 
| data | bytes | Transaction data payload. | 

### getConfirmationCount

Returns number of confirmations of a transaction.

```js
function getConfirmationCount(uint256 transactionId) public view
returns(count uint256)
```

**Returns**

Number of confirmations.

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| transactionId | uint256 | Transaction ID. | 

### getTransactionCount

Returns total number of transactions after filers are applied.

```js
function getTransactionCount(bool pending, bool executed) public view
returns(count uint256)
```

**Returns**

Total number of transactions after filters are applied.

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| pending | bool | Include pending transactions. | 
| executed | bool | Include executed transactions. | 

### getOwners

Returns list of owners.

```js
function getOwners() public view
returns(address[])
```

**Returns**

List of owner addresses.

### getConfirmations

Returns array with owner addresses, which confirmed transaction.

```js
function getConfirmations(uint256 transactionId) public view
returns(_confirmations address[])
```

**Returns**

Returns array of owner addresses.

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| transactionId | uint256 | Transaction ID. | 

### getTransactionIds

Returns list of transaction IDs in defined range.

```js
function getTransactionIds(uint256 from, uint256 to, bool pending, bool executed) public view
returns(_transactionIds uint256[])
```

**Returns**

Returns array of transaction IDs.

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| from | uint256 | Index start position of transaction array. | 
| to | uint256 | Index end position of transaction array. | 
| pending | bool | Include pending transactions. | 
| executed | bool | Include executed transactions. | 

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
