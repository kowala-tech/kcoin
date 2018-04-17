#!/bin/sh
set -e
trap "exit" INT

cd /kcoin

./kcoin init /kcoin/genesis.json
./kcoin --verbosity 2 "$@"
