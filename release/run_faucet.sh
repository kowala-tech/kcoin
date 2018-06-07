#!/bin/sh
set -e

./faucet \
  --genesis /faucet/genesis.json \
  $@
