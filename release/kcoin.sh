#!/bin/sh
set -e

cd /kcoin

if [[ $@ = *"--testnet"* ]]; then
  ./kcoin init /kcoin/testnet_genesis.json
else
  ./kcoin init /kcoin/genesis.json
fi


./control --ipc /root/.kcoin/kcoin.ipc &
status=$?
if [ $status -ne 0 ]; then
  echo "Failed to start control panel: $status"
  exit $status
fi

./kcoin "$@"
