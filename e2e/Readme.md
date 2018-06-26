# E2E

This project contains all the e2e testing suite for all kowala related projects.

# Config

set the following env variables for a better control of how things run

- `DOCKER_HOST` with the docker host. Usually `tcp://docker:2375`
- `DOCKER_PUBLIC_IP` it's the IP address to be used to connect to services running in that docker host. Usually `docker`.

# How it works

The e2e tests are written using gherkin and runs using [godog](https://github.com/DATA-DOG/godog)

Each scenario starts with a simple private kcoin netork (bootnode, genesis validator and rpc).

The services are usually built using the public docker image of each service using the tag `:dev`.
