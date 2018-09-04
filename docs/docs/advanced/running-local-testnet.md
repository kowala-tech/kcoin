# Running local testnet

## Creating a testnet

In the toolbox/ folder we will find a tool for creating local testnets for testing, completely isolated.

To execute a local testnet we only need to run:

```
go run cmd/testnet/testnet.go
```

Then we will have a testnet running on docker, use `docker ps` for more info.
It is based on 2 containers, one is a validator and another one is a bootnode.

It exposes port **30503** with an rpc endpoint.

## Connecting to the console with truffle.

Once the network is running it is possible to connect using truffle console.
We can go to client/contracts/truffle and run:

```
truffle console --network kcoin
```

And we will have a truffle console, with all the features of a geth console plus some benefits from truffle.