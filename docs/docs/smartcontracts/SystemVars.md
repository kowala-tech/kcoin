# System Variables (SystemVars.sol)

**contract SystemVars is [Initializable](Initializable.md)**

**SystemVars**

## Contract Members
**Constants & Variables**

```js
//internal members
uint256 internal constant INITIAL_MINTED_AMOUNT;
uint256 internal constant INITIAL_CAP;
uint256 internal constant STABILIZED_PRICE;
uint256 internal constant ADJUSTMENT_FACTOR;
uint256 internal constant LOW_SUPPLY_METRIC;
uint256 internal constant MAX_UNDER_NORMAL_CONDITIONS;
uint256 internal constant DEFAULT_ORACLE_REWARD;
uint256 internal constant ORACLE_DEDUCTION_FRACTION;
//public members
uint256 public prevCurrencyPrice;
uint256 public currencyPrice;
uint256 public currencySupply;
uint256 public mintedReward;
```

## Functions

- [initialize](#initialize)
- [_hasEnoughSupply](#_hasenoughsupply)
- [_cap](#_cap)
- [price](#price)
- [mintedAmount](#mintedamount)
- [oracleDeduction](#oraclededuction)
- [oracleReward](#oraclereward)

### initialize

```js
function initialize(uint256 _initialPrice, uint256 _initialSupply) public isInitializer
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| _initialPrice | uint256 | initial price for the system's currency | 
| _initialSupply | uint256 | minted amount on the genesis block | 

### _hasEnoughSupply

```js
function _hasEnoughSupply() private view
returns(bool)
```

### _cap

```js
function _cap() private view
returns(amount uint256)
```

### price

Get the current system's currency price

```js
function price() public view
returns(price uint256)
```

### mintedAmount

Get the amount of coins that should be minted

```js
function mintedAmount() public view
returns(uint256)
```

### oracleDeduction

Get the oracle deduction

```js
function oracleDeduction(uint256 mintedAmount) public view
returns(uint256)
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| mintedAmount | uint256 | the minted amount for the current block. | 

### oracleReward

Get the oracle reward

```js
function oracleReward() public view
returns(uint256)
```

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
