# Token Standards

In the beginning, compatibility with projects and services was a major issue and concern throughout the community - most of the ICOs to this day have implemented a token using the ERC20 standard but this standard comes with several problems. There are currently at least three improvement proposals to tackle ERC20's biggest flaws. The main proposal is ERC223 which we are going to detail below. Most of the services have been built based on the ERC20 standard so having backwards compability is a must. Note that we can implement mUSD as an ERC20 token but I think that we should strive for better standards as soon as possible - this does not mean that ERC223 will be the final answer.

## ERC223

```
Specification: https://github.com/Dexaran/ERC223-token-standard/
```

ERC223 tokens are backwards compatible with ERC20 tokens. This means that ERC223 supports every ERC20 functionality and contracts or services working with ERC20 tokens will work with ERC223 tokens correctly. Main improvements over ERC20:

* Eliminates the lost token problem - when people send funds to a contract by using the transfer function (instead of approve and transferFrom), funds are lost and cannot be retrieved. The ERC20 standard doesn’t include a refund function for these cases and the transferred tokens are lost. So far this implementation has resulted in the loss of millions; there's a list in [here](https://github.com/Dexaran/ERC223-token-standard/).
* Allowing developers to handle incoming token transactions with additional capability to reject and revert non-supported tokens to their user, which is not possible with ERC20.
* Transfers become just one single transaction, instead of two, costing half as much gas as they did with ERC20.

## mUSD

We currently have an implementation in the ERC223 standard which is backwards compatible with ERC20. We need to verify the exchanges' take on this topic going forward given that we may list mUSD. Our token is currently set to have 18 decimal places which is the "standard". Main motivation for 18 decimal places is to ensure that in any on-chain exchange, the price as expressed in computer units (ie. wei or equivalent) is the same number as the price as expressed in human units (ie. ether or equivalent), reducing the risk of confusion or bugs. For more information on this topic please check [this link](https://github.com/ethereum/EIPs/issues/724).

</br>
</br>
