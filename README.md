[![Gitter chat](https://badges.gitter.im/kowala/kcoin.png)](https://gitter.im/kowala-tech/kcoin) [![Build Status](http://ci.kowala.io/api/badges/kowala-tech/kcoin/status.svg)](http://ci.kowala.io/kowala-tech/kcoin) [![Public testnet](https://img.shields.io/badge/public-testnet-981071.svg)](http://testnet.kowala.io)

_Note: This readme contains technical documentation for working with the source code and performaing low-level technical tasks. The end-user documentation can be found at [docs.kowala.io](http://docs.kowala.io)._

## kCoin mining client

Official implementation of the Kowala protocol based on the [go-ethereum client](https://github.com/ethereum/go-ethereum/). The **`kcoin`** client is the main client for the Kowala network.

It is the entry point into the Kowala network, and is capable of running a full node. The client offers
a gateway (Endpoints, WebSocket, IPC) to the Kowala network to other processes.

## Running kcoin

### Building the source

	make kcoin

### Configuration

As an alternative to passing the numerous flags to the `kcoin` binary, you can also pass a configuration file via:

```
$ kcoin --config /path/to/your_config.toml
```

To get an idea how the file should look like you can use the `dumpconfig` subcommand to export your existing configuration:

```
$ kcoin --your-favourite-flags dumpconfig
```

or check the [config sample](https://github.com/kowala-tech/kcoin/blob/master/sample-kowala.toml).

### Client Options

[Client page]()

### Docker quick start

One of the quickest ways to get Kowala up and running on your machine is by picking a currency and using Docker:

```
docker run -d --name kcoin-node -v /Users/alice/kcoin:/root \
		   -p 11223:11223 -p 22334:22334 \
		   kowalatech/kusd --fast --cache=512
```

## Networks

### Public

There aren't public networks at the moment.

### Testnet

http://testnet.kowala.io/

## Creating a Private Blockchain Network

### Genesis State

#### Validators

1. Generate/request a new account for each genesis validator

```
$ kcoin --config /path/to/your_config.toml account new
Address: {c7f1d574658e7b0f37244366c40c8002d78c734f}
```

2. Fill in the validator details in the network contracts

   1. pre-fund the validators in [here](https://github.com/kowala-tech/kcoin/blob/feature/tendermint/contracts/network/contracts/mUSD.sol#L10)
   2. mark the validators as genesis validators and also as voters [here](https://github.com/kowala-tech/kcoin/blob/feature/tendermint/contracts/network/contracts/network.sol#L96)

#### Network Contracts - Owner

1. Generate a new account - this account will be selected (on the next step) as the owner of the network contracts.

```
$ kcoin --config /path/to/your_config.toml account new
Address: {c7f1d574658e7b0f37244366c40c8002d78c734f}
```

#### Network

1. Set `totalSupplyWei` field present in [here](https://github.com/kowala-tech/kcoin/blob/feature/tendermint/contracts/network/contracts/network.sol#L5) to the correct pre-minted amount of kcoin (only needed to calculate the blockcap for the reward)

#### File

1. run the code generation on the `contracts/network` sub-package.

```
$ go generate
```

2. The first step consists in creating the genesis of your new network. By far, the easiest way to do it, is by running the puppeth client.

   1. Rebuild the puppeth client
	  `$ cd cmd $ go install ./puppeth/...`

   2. Run the client
	  `$ puppeth`

   3. Specify a network name

   4. Select the option "2. Configure new genesis"

   5. Select "1. Tendermint - proof-of-stake"

   6. Fill in the account of the owner of the network contracts

   7. Fill in any additional information (until the process is complete)

   8. Select "2. Save existing genesis" and fill in the file path to save the genesis into.

```
	$ Which file to save the genesis into? (default = test.json)
	> /src/github.com/kowala-tech/kcoin/assets/test.json
	INFO [01-16|16:49:37] Exported existing genesis block
```

3. Initialize the blockchain based on the genesis file created on the previous step.

```
$ kcoin --config /path/to/your_config.toml init path/to/genesis.json
```

4. Set the `mapAddress` (based on the genesis file) variable in [here](https://github.com/kowala-tech/kcoin/blob/feature/tendermint/contracts/network/data_layouts.go#L94) to the correct address.

5. Rebuild the kcoin client

```
$ make kcoin
```

### Bootstrap Node

In order to have nodes find each other, you need to start a bootstrap node.

1. Generate a node key.

```
$ bootnode --genkey=boot.key
```

2. Start the bootnode using the given node key.

```
$ bootnode --nodekey=boot.key
```

As soon as the bootnode is running, it will display an enode URL that other nodes can use to connect
and gather info on other nodes.

## Proof-of-Stake (PoS)

### Protocol

[Tendermint](https://github.com/tendermint/tendermint)

### Running a PoS validator

Make sure that you have an account available:

```
$ kcoin --config /path/to/your_config.toml account new
Address: {c7f1d574658e7b0f37244366c40c8002d78c734f}
```

To start a kcoin instance for block validation, run it with all your usual flags, extended by:

```
$ kcoin --config /path/to/your_config.toml --validate --deposit 4000 --unlock 0xc7f1d574658e7b0f37244366c40c8002d78c734f –-coinbase 0xc7f1d574658e7b0f37244366c40c8002d78c734f
```

## Mining client metrics

Start the client with `--metrics` to collect performance metrics. This will expose a [Prometheus](https://prometheus.io/) HTTP endpoint at `/metrics` on `http://localhost:8080`.

Prometheus endpoint address can be specified using flag `--metrics-prometheus-address`.
