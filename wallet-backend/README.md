# Wallet backend

## How to run

The app is dockerized, so the best way to run it should be building with docker
with command:

```
docker build -t backend .
```

Then we will have a created image called backend which we can run with:

```
docker run --rm -it -d -p 80:80 backend
```

We will have then a server listening in the 80 port.

To get information of the available parameters you can run:

```
docker run --rm -it backend --help
```


## API endpoints

### Blocknumber

You can get the last block number with:

```
http://localhost/api/blockheight
```

We will receive a response like:

```
{"block_height":2055401}
```


This can be accessed too by websockets, in the endpoint:

```
ws://localhost/ws
```

And send throught it
```
{"action":"blockheight"}
```

### Balance

We can take the balance of a specific account using:

```
http://localhost/api/balance/accountnum
```

We will receive a response like:

```
{"balance":7000000000000000000}
```

### Transactions of an account

We can take the transactios of a specific account using:
```
http://localhost/api/transactions/accountnum
```

And we can receive something like:

```
{"transactions":[{"hash":"0x6d7216643e4aabd748b1e15c019dfca7f98baf23b0d4a8c43cfe6f60d710f533","from":"0xD6e579085c82329C89fca7a9F012bE59028ED53F","to":"0x2a4d42ddEFb0e82be965cE545F8d2f882cDc997b","amount":1000000000000000000},{"hash":"0x36e5b8bc3d8cdee55dde895b10b75dc1ce65e3552575c5b06518eaf72e6e496a","from":"0xD6e579085c82329C89fca7a9F012bE59028ED53F","to":"0x2a4d42ddEFb0e82be965cE545F8d2f882cDc997b","amount":1000000000000000000}]
```

We can specify a block range when asking for transactions like:

```
http://localhost/api/transactions/accountnum/from/{blocknum}/to/{blocknum}
```
