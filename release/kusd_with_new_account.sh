#!/bin/sh
set -e
trap "exit" INT

cd /kusd

./kusd init /kusd/genesis.json

echo "test" > password.txt
address=$(./kusd account new --password password.txt | tail -c42 | head -c40)

./kusd $(eval "echo $@")
