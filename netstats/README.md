# kCoin Network Stats

This is a visual interface for tracking kCoin network status. It uses WebSockets to receive stats from running nodes and output them through an angular interface. 

## Prerequisite
* node
* npm or yarn
* grunt-cli

## Build the application

```bash
grunt
```

## Run

You'll need to set the `WS_SECRET` environment variable. This will be required for all clients sending stats.

```bash
npm start
```

or, for yarn users:

```bash
yarn start
```

If you have `make` installed, you can also run

```bash
make
```

to start the server with a `WS_SECRET` of `abc123` (see Makefile).

## Interacting with the server once it's running

See the interface at http://localhost:3000.
