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

### Client Options

[CLI page](https://github.com/ethereum/go-ethereum/wiki/Command-Line-Options)

### Docker quick start

One of the quickest ways to get Kowala up and running on your machine is by using Docker:

```
docker run -d --name kusd-node -v /Users/alice/kusd:/root \
           -p 11223:11223 -p 22334:22334 \
           kusd/client-go --fast --cache=512
```

## Proof-of-Stake (PoS)

## Protocol

[Tendermint](https://github.com/tendermint/tendermint)

### Running a PoS validator

Make sure that you have an unlocked account available:

```
kusd account new
kusd --unlock "0x407d73d8a49eeb85d32cf465507dd71d507100c1"
```

To start a kusd instance for block validation, run it with all your usual flags, extended by:

```
$ kusd <usual-flags> --validate --deposit=4000 --coinbase=0x0000000000000000000000000000000000000000
```
