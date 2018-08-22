# Kowala toolbox

## Packages

- Testnet (For creating fast testnets, useful for tests.) [Documentation](testnet/README.md)

## Creating a testnet

```
go run cmd/testnet/testnet.go
```

Then we will have a testnet running on docker, use docker ps for more info.
It is based on 2 containers, one is a validator and another one is a bootnode.