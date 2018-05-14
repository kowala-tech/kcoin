#!/bin/sh
set -e

cd /kcoin

./kcoin init /kcoin/genesis.json
./kcoin "$@"
