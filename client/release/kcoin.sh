#!/bin/sh
set -e

DO_INIT=true;
TESTNET="";
NEW_ACC=false;
NEW_ACC_PASS="";
GENESIS_PATH="";

command="";

while [[ $# -gt 0 ]] ;
do
    opt="$1";
    shift;        

    case "$opt" in
	  "--testnet")
		TESTNET="--testnet"
		command="$command$opt " # make sure this passed to the binary
		;;
	  "--dev")
		TESTNET="--dev"
		command="$command$opt " # make sure this passed to the binary
		;;
	  "version")
		DO_INIT=false
		command="$command$opt " # make sure this passed to the binary
		;;
	  "--new-account")
		NEW_ACC=true
		;;
	  "--new-account-password="*)
		NEW_ACC_PASS="${opt#*=}"
		;;
	  "--genesis-path="*)
		GENESIS_PATH="${opt#*=}"
		;;
	  *)
		command="$command$opt " # make sure this passed to the binary
		;;
	esac
done

cd /kcoin

case $DO_INIT in
	(true) ./kcoin init $TESTNET "$GENESIS_PATH"
		;;
esac

case $NEW_ACC in
	(true)
		echo "$NEW_ACC_PASS" > .password
		./kcoin account new --password .password
		rm .password
		;;
esac

./control --ipc /root/.kcoin/kusd/kcoin.ipc &
status=$?
if [ $status -ne 0 ]; then
  echo "Failed to start control panel: $status"
  exit $status
fi


./kcoin $command
