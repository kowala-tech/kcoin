#!/bin/sh
set -e
trap "exit" INT

cd /kusd

./kusd init /kusd/genesis.json

echo "test" > password.txt
./kusd account new --password password.txt

./kusd "$@"
