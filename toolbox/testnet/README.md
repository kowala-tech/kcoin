# Testnet tool for tests

## Pre-requisites

- A working docker installation.

To be sure that this will work on your machine launch the unit tests with:

```go test ./...```

If it passes it will work.


## How to use

Basically the idea is to have a testnet created so you can test different
services on the blockchain. So to launch it just do:

```
	testnet := NewTestnet(dockerEngine)
	err = testnet.Start()
	if err != nil {
	    t.Fatalf("Error %s", err)
	}
```

Then you will have a testnet running in several containers, you can check with
docker ps what is there.

If everything has gone well

```
    testnet.IsValidating()
```

Should be true.

By default the testnet includes:

- a bootnode
- a validator with rpc port exposed.

To be able to create more containers that work in the same environment as the
testnet, for example a wallet backend, we need to execute it in the same network
as the testnet. To get the testnet network we can use

```
testnet.GetNetworkID()
```

To stop the testnet just do:

```
testnet.Stop()
```

For more information check the integration tests of the wallet-backend repo.