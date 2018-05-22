#!/bin/sh
set -e

cd /kcoin

if [[ $@ = *"--testnet"* ]]; then
  ./kcoin init /kcoin/testnet_genesis.json
else
  ./kcoin init /kcoin/genesis.json
fi


./kcoin "$@"
