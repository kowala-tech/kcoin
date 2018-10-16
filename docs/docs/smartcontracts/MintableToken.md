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
