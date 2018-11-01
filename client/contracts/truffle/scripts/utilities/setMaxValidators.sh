#!/bin/sh
npm install;
npm run truffle compile;
npm run truffle exec -- --network kcoin_test scripts/setMaxValidators.js -n $1;
