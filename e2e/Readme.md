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

# How to run the e2e suite with a WIP image

Just build the docker image of your WIP code (or your PR code) in the same docker host the e2e will run and tag it with the tag that will run, usually `dev`

# Example for Drone CI

This is a simple example of how to run the e2e suite using the local code from a PR.

```
pipeline:
  docker-image-for-e2e:
    image: docker
    environment:
    - DOCKER_HOST=tcp://docker:2375
    commands: 
    - docker build -t kowalatech/my-service:dev -f Dockerfile .
    when:
      event: [pull_request]

  e2e:
    image: kowalatech/e2e:latest
    pull: true
    environment:
    - DOCKER_HOST=tcp://docker:2375
    - DOCKER_PUBLIC_IP=docker
    commands: 
    - /e2e
    when:
      event: [pull_request]

services:
  docker:
    image: docker:dind
    command: [ '-l', 'fatal' ]
    privileged: true

```
