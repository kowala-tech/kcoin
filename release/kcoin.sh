#!/bin/sh
set -e

cd /kcoin

case "$@" in 
  *"--testnet"*)
    ./kcoin init /kcoin/testnet_genesis.json
    ;;
  *)
    ./kcoin init /kcoin/genesis.json
    ;;
esac

./kcoin "$@"
