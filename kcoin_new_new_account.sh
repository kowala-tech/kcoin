#!/bin/sh
set -e
trap "exit" INT

cd /kcoin

./kcoin init /kcoin/genesis.json

echo "test" > password.txt
./kcoin account new --password password.txt

./kcoin --verbosity 2 "$@"
