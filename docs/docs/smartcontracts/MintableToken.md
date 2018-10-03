# MintableToken.sol

**contract MintableToken is [Token](Token.md), [Ownable](Ownable.md)**

**MintableToken**

## Contract Members
**Constants & Variables**

```js
bool public mintingFinished;
```

**Events**

```js
event Mint(address indexed to, uint256 amount);
event MintFinished();
```

## Modifiers

- [canMint](#canmint)

### canMint

```js
modifier canMint() internal
```

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|

## Functions

- [mint](#mint)
- [finishMinting](#finishminting)

### mint

Function to mint tokens

```js
function mint(address _to, uint256 _amount) public onlyOwner canMint
returns(bool)
```

**Returns**

A boolean that indicates if the operation was successful.

**Arguments**

| Name        | Type           | Description  |
| ------------- |------------- | -----|
| _to | address | The address that will receive the minted tokens. | 
| _amount | uint256 | The amount of tokens to mint. | 

### finishMinting

Function to stop minting new tokens.

```js
function finishMinting() public onlyOwner canMint
returns(bool)
```

**Returns**

True if the operation was successful.

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
