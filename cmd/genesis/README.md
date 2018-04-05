# Usage

## From config file
To create a genesis json file from config file the command to execute is

```
genesis -c sample-config.toml
```

A typical config file should be in toml and should like like this:

```
[genesis]
network = "test" # Only supported test and main
maxNumValidators = "1"
unbondingPeriod = "1"
walletAddressGenesisValidator = "0xfe9bed356e7bc4f7a8fc48cc19c958f4e640ac62"

# Optional values
consensusEngine = "tendermint" # Only supported tendermint by the moment
smartContractsOwner = "0x1234ed356e7bc4f7a8fc48cc19c958f4e640ac62"
extraData = "extraData"

[[prefundedAccounts]]
walletAddress = "0xfe9bed356e7bc4f7a8fc48cc19c958f4e640ac62"
balance = 15

[[prefundedAccounts]]
walletAddress = "0x1234ed356e7bc4f7a8fc48cc19c958f4e640ac62"
balance = 15
```

## From params
If we don't want to create the file we can execute the same in this format.

```
genesis -n test -v 1 -p 1 -g 0xfe9bed356e7bc4f7a8fc48cc19c958f4e640ac62 -a 0xfe9bed356e7bc4f7a8fc48cc19c958f4e640ac62:13,0xfe9bed356e7bc4f7a8fc48cc19c958f4e640ac65:13
```

To get more information about the available flags just do:

```
genesis -h
```