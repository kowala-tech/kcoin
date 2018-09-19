# Exchanges mocked server

## What is this?

In order to be able to test the oracles in a controlled way we need a mock
server that is able to fake the returned data about currency values.

This allows to set the values that it will return once you create a
request as it were real exchanges as Exrates.

## How to use it?

First of all we will need to build the docker image if it is not done yet,
from base directory we can do:

``` docker build -t kowalatech/mock-exchange -f mock-exchange/Dockerfile .```

Once it finishes we will have an image ready to run.

To run it we can do:

```docker run --rm -it -p 8080:8080 -d kowalatech/mock-exchange serve```

### Available requests

First of all we will define what data our service will return, this is done calling
`/api/fetch` endpoint.

```
POST /api/endpoint
{
	"sell": [
		{"amount":1,"rate":2},
		{"amount":1,"rate":2}
	],
	"buy": [
		{"amount":1,"rate":2},
		{"amount":1,"rate":2}
	]
}
```

In this we can see that we have defined some fake data for sell prices and buy prices.

#### Getting data

To get data we use the endpoint `/api/exrates/get` (in the future there will be other exchanges)
we get the data that we first defined but formated as the real exchange data.

```
GET /api/exrates/get
```

Returns:

```
{"SELL":[{"amount":1,"rate":2},{"amount":1,"rate":2}],"BUY":[{"amount":1,"rate":2},{"amount":1,"rate":2}]}
```
