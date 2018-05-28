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

./control --ipc /root/.kcoin/kcoin.ipc &
status=$?
if [ $status -ne 0 ]; then
  echo "Failed to start control panel: $status"
  exit $status
fi

./kcoin "$@"
