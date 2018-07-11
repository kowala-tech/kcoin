# Consensus - Validator Registry

This section covers the steps required to start a consensus validator . To read
more about Kowala's consensus engine, please refer to the
[Consensus](/consensus/overview) section.

## Requirements

- your node is already running as specified in the [testnet
  guide](/getting-started/testnet) and you have access to the console.

## Validator Registration

** Note: if your client is not synced, it might take some time to get your
transaction posted - the validator only submits the transaction as soon as your
node is synced with the network. **

You must have an unlocked account with enough mining tokens to cover the initial
deposit required to participate in the consensus.** If you wish to join the
consensus let us know in our [Telegram](https://t.co/MpSK3z1aWw) or [Gitter
channel](https://gitter.im/kowala-tech/Lobby), and we will mint enough mUSD
tokens to your account.**

If you don't have an account yet, you can create a new one and unlock it by
running the following commands:

```
> personal.newAccount()
Passphrase: *****
Repeat passphrase: *****
"0x3dc0bd5208f2e7dcb896f99546e0f9586d58257c"

> personal.unlockAccount("0x3dc0bd5208f2e7dcb896f99546e0f9586d58257c")
Unlock account 0x3dc0bd5208f2e7dcb896f99546e0f9586d58257c
Passphrase: *****
true
```

To verify your current mining token balance run:

```
> mtoken.getBalance("0x3dc0bd5208f2e7dcb896f99546e0f9586d58257c")
2e+24
```

To verify the minimum deposit required run:

```
> validator.getMinimumDeposit()
1e+24
```

As soon as all the required steps were accomplished, you should select the
account to be used:

```
validator.setCoinbase("0x3dc0bd5208f2e7dcb896f99546e0f9586d58257c")
```

and the deposit value. For instance, transfer all the mUSD tokens to maximize the
rewards,

```
validator.setDeposit(mtoken.getBalance("0x3dc0bd5208f2e7dcb896f99546e0f9586d58257c"))
```

and now we are ready to start the consensus validator:

```
validator.start()
```

At this point, your mining tokens are locked in the consensus contract and you
are part of the consensus.

## Validator Deregistration

In order to leave the consensus, you must run:

```
validator.stop()
```

**Note: there's a freeze period and your mUSD tokens will remain locked for
a pre-defined period of time**. In order to verify the status of your locked
deposit(s) run:

```
validator.getDeposits()
```

As soon as you leave the consensus, you can join right away even if you have a
locked deposit, as long as you have enough mining tokens for the new registration.

## Redeem Unlocked Deposits

As soon as your deposit(s) is past the freeze period you must request them in
order to have them back to your account:

```
validator.redeemDeposits()
```

</br></br>
