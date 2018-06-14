#!/bin/sh
set -e

TESTNET=false;
NEW_ACC=false;
NEW_ACC_PASS="";

command="";

while [[ $# -gt 0 ]] ;
do
    opt="$1";
    shift;        

    case "$opt" in
	  "--testnet")
		TESTNET=true
		command="$command$opt " # make sure this passed to the binary
		;;
	  "--new-account")
		NEW_ACC=true
		;;
	  "--new-account-password="*)
		NEW_ACC_PASS="${opt#*=}"
		;;
	  *)
		command="$command$opt " # make sure this passed to the binary
		;;
	esac
done

cd /kcoin

case $TESTNET in
	(true)  ./kcoin init --testnet;;
	(false) ./kcoin init;;
esac

case $NEW_ACC in
	(true)
		echo "$NEW_ACC_PASS" > .password
		./kcoin account new --password .password
		rm .password
		;;
esac

./control --ipc /root/.kcoin/kcoin.ipc &
status=$?
if [ $status -ne 0 ]; then
  echo "Failed to start control panel: $status"
  exit $status
fi


./kcoin $command
