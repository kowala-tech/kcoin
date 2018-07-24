# Kowala Testnet

This section will help you get started with the latest [test network](http://zygote.kowala.io/).

The [Kowala Testnet](http://testnet.kowala.io/) is the perfect place for testing the Kowala Protocol, kCoins and mTokens without the fear of making a mistake and losing real money. You can spin up your own node and connect to the network, see how the network is performing using the [networks stats dashboard](http://testnet.kowala.io/stats/), and get free kUSD to play with using the [Kowala's faucet](http://faucet.testnet.kowala.io/).

## Requirements

You can connect to the testnet using the official kcoin docker image. We strongly advise you to use the official kcoin docker image, but advanced users are of course free to clone the source code from [kcoin repository](https://github.com/kowala-tech/kcoin) and compile their own clients. In order to use the docker image, make sure that your setup meets the following requirements:

- there aren't hardware requirements yet - the current test network requires minimal resources.

- [install Docker](https://www.docker.com/community-edition) if you don't already have it installed. There are Linux, Mac and Windows versions available. Docker is a virtualization system that allows software compiled for Linux to run on any platform. We use it to automatically fetch and run pre-compiled, up-to-date versions (called 'images') of the kUSD mining client.

- We frequently push updates to the client. Make sure that you have the latest docker image by running:

```
docker pull kowalatech/kusd
```

## Connecting to the Kowala Network

With docker installed, it's time to fire up the the terminal. To download the mining client, start it, generate a wallet and connect to the testnet, run:

```
docker run --rm -it kowalatech/kusd --testnet --new-account console
```

Or if you already have an account:
```
docker run --rm -it kowalatech/kusd --testnet console
```

Note that this command will create a disposable version of the mining client — all accounts will be deleted as soon as you terminate the process. If want to persist the state, you can leave off the --rm flag.

By now, you should be in the interactive console. Your node will take some time to synchronize with the network.
You can confirm the sync status by running the following command:

```
kcoin.syncing
```

## Getting kUSD to play with

The official Docker image will create an account for you when it starts, and you can get its public address via the console:

```
> kcoin.coinbase
```

The console will output your public address. For example:

```
> kcoin.coinbase
"0xe2ac86cbae1bbbb47d157516d334e70859a1be45"
```

You can use that address in the [Coin Faucet](<(http://faucet.testnet.kowala.io/)>) to acquire some free money. The Coin Faucet will send you money (the operation should take ~ 1 second), and you can check your balance with:

```
> kcoin.getBalance(kcoin.coinbase)
```

## Network Status

The KCoin network status monitor is a web-based application to monitor the health of the testnet/mainnet through a group of nodes.

You can visit [here](https://zygote.kowala.tech/stats/).

## BlockChain Explorer

Easy viewer for investors and developers.

You can visit [here](https://explorer.zygote.kowala.tech/).

</br></br>
