# Capped token (CappedToken.sol)

**contract CappedToken is [MintableToken](MintableToken.md)**

**CappedToken**

Mintable token with a token cap.

## Contract Members
**Constants & Variables**

```js
uint256 public cap;
```

## Functions

- [mint](#mint)

### mint

:small_red_triangle: overrides [MintableToken.mint](MintableToken.md#mint)

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
