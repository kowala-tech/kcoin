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

#### Validators

1. Generate/Request a new account for each genesis validator

```
$ kusd --config /path/to/your_config.toml account new
Address: {c7f1d574658e7b0f37244366c40c8002d78c734f}
```

2. Fill in the validator details in the network contracts
   2.1. pre-fund the validators in [here](https://github.com/kowala-tech/kUSD/blob/feature/tendermint/contracts/network/contracts/mUSD.sol#L10)
   2.2. mark the validators as genesis validators in [here](https://github.com/kowala-tech/kUSD/blob/feature/tendermint/contracts/network/contracts/network.sol#L96)

3. run the code generation on the `contracts/network` sub-package

````
$ go generate
```

#### Network Contracts - Owner

1. Generate a new account - this account will be selected (on the next step) as the owner of the network contracts.

```
$ kusd --config /path/to/your_config.toml account new
Address: {c7f1d574658e7b0f37244366c40c8002d78c734f}
```

#### File

1. The first step consists in creating the genesis of your new network. By far, the easiest way to do it, is by running the puppeth client.

1.1. Rebuild the puppeth client

   ```
        $ cd cmd
        $ go install ./puppeth/...
   ```

1.2. Run the client

    ```
        $ puppeth
    ```

1.3. Specify a network name

1.4. Select the option "2. Configure new genesis"

1.5. Select "1. Tendermint - proof-of-stake"

1.6. Fill in the account of the owner of the network contracts

1.7. Fill in any additional information (until the process is complete)

1.8. Select "2. Save existing genesis" and fill in the file path to save the genesis into.

```
    $ Which file to save the genesis into? (default = test.json)
    > /src/github.com/kowala-tech/kUSD/assets/test.json
    INFO [01-16|16:49:37] Exported existing genesis block
```

2. Initialize the blockchain based on the genesis file created on the previous step.

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
````
