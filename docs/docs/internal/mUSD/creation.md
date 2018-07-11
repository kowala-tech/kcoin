# Token Creation

The mUSD token is a capped token and is part of the network core contracts. The network core contracts are included in the genesis block state during the genesis block generation. This means that as soon as the production network starts, the token contract will be available with the investor balances. The cap of the mUSD contract has been set to 1073741824 tokens as defined in the whitepaper. Given the cap defined, we may mint a maximum number of 1073741824 tokens and no more than that as ruled by the contract.

```
contract mUSD is CappedToken {
    constructor() public CappedToken(1073741824) {
        name = "mUSD";
        symbol = "mUSD";
        decimals = 18;
    }
}
```

</br>
</br>
