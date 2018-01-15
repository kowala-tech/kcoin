## Kowala

Official implementation of the Kowala protocol. The **`kusd`** client is the main client for the Kowala network.
It is the entry point into the Kowala network, and is capable of running a full node(default). The client offers
a gateway (Endpoints, WebSocket, IPC) to the Kowala network to other processes.

## Running kusd

### Building the source

    make kusd

### Configuration

As an alternative to passing the numerous flags to the `kusd` binary, you can also pass a configuration file via:

```
$ kusd --config /path/to/your_config.toml
```

To get an idea how the file should look like you can use the `dumpconfig` subcommand to export your existing configuration:

```
$ kusd --your-favourite-flags dumpconfig
```

or check the [config sample](https://github.com/kowala-tech/kUSD/blob/master/sample-kowala.toml).

### Client Options

[Client page]()

### Docker quick start

One of the quickest ways to get Kowala up and running on your machine is by using Docker:

```
docker run -d --name kusd-node -v /Users/alice/kusd:/root \
           -p 11223:11223 -p 22334:22334 \
           kusd/client-go --fast --cache=512
```

## Networks

### Public

There aren't public networks at the moment.

### Test

http://testnet.kowala.io/

## Creating a Private Blockchain Network

### Genesis State

1. Generate a new account - this account will be selected (on the next step) as the owner of the network contracts.

```
$ kusd --config /path/to/your_config.toml account new
Address: {c7f1d574658e7b0f37244366c40c8002d78c734f}
```

2. The second step consists in creating the genesis of your new network. By far, the easiest way to do it, is by
   running the puppeth client.

```

```

3. Initialize the blockchain based on the genesis file given.

```
$ kusd --config /path/to/your_config.toml init path/to/genesis.json
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
$ kusd --config /path/to/your_config.toml account new
Address: {c7f1d574658e7b0f37244366c40c8002d78c734f}
```

To start a kusd instance for block validation, run it with all your usual flags, extended by:

```
$ kusd --config /path/to/your_config.toml --validate --deposit 4000 --unlock 0xc7f1d574658e7b0f37244366c40c8002d78c734f –-coinbase 0xc7f1d574658e7b0f37244366c40c8002d78c734f
```

## Core Contributors

[Core Team Members](https://github.com/orgs/kowala-tech/people)

## Contact us

Feel free to email us at support@kowala.tech.
