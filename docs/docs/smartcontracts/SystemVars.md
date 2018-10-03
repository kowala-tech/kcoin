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
