# Genesis

The genesis block is the first block of a block chain, the starting point. In order to generate a genesis file, Kowala provides a [command line client](https://github.com/kowala-tech/kcoin/tree/dev/cmd/genesis) that generates the genesis json file based on a configuration file/cmdline flags.

## Client

Make sure that you have the latest version of the client by running the following command on the root dir of [the kcoin repository](https://github.com/kowala-tech/kcoin).

```bash
$ make genesis
```

### Config

In order to use the genesis client you must provide a configuration file as the one in the following command:

```bash
$ ./genesis --config /Users/ricardogeraldes/code/src/github.com/kowala-tech/kcoin/cmd/genesis/sample-config.toml
```

A config sample is available in the [kcoin repository](https://github.com/kowala-tech/kcoin/blob/dev/cmd/genesis/sample-config.toml).
We detail the purpose of each configuration field on the following sections.

#### Governance

Kowala is the owner of the core contracts in the form of a multi signature wallet. The main idea is that we should not be subject to a single key. In this scenario, even if a set of keys get compromised, the attackers won't be able to make requests because a given number of authorizations is necessary. By default the max number of governors is set to 50, but it can be changed to a different number.

| Field            | Description                                                                                                                                                                                                                                         |
| ---------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| origin           | This address is used to get a specific address for the core contracts - this value should remain the same during the dev phase. Changing this value means that we will have to change the hardcoded addresses used by the golang contract bindings. |
| governors        | Set of addresses that govern the life cycle of kowala's core contracts. **len(governors) must be > 0.**                                                                                                                                             |
| numConfirmations | Number of governors' approvals necessary to post a transaction. **Value must be <= len(governors).**                                                                                                                                                |

#### Consensus

| Field            | Description                                                                                                                                                                                                                        |
| ---------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| engine           | Specifies the consensus engine used in the network. "tendermint" is the only available option at the moment.                                                                                                                       |
| maxNumValidators | Maximum number of validators. **Value must be > 0.**                                                                                                                                                                               |
| validators       | Set of initial validators aka genesis validators. Note: the code currently supports one validator but this might change in the future. **len(validators) must be > 0 and the genesis validator must be listed as a token holder**. |
| freezePeriod     | Period of time, in days, that coins remained locked as soon as a validator decides to leave the consensus.                                                                                                                         |
| baseDeposit      | Minimum deposit, in mUSD, required to be a consensus validator. Note: the minimum deposit value increases as soon as the consensus is full.                                                                                        |
| miningToken      | Set of fields related to the mining token.                                                                                                                                                                                         |

Mining Token Fields:

| Field    | Description                                                                                     |
| -------- | ----------------------------------------------------------------------------------------------- |
| Name     | Mining token name. Ex: "mUSD".                                                                  |
| Symbol   | Mining token symbol. Ex: "$".                                                                   |
| Cap      | Total supply.                                                                                   |
| Decimals | Mining token decimals. 18 by default. Relevant: https://github.com/ethereum/EIPs/issues/724.    |
| Holders  | Set of fields related to token holders. We use this field to specify the initial token holders. |

Note that the process of distributing tokens remains available until Kowala considers the minting process over, in the form of a transaction.

Token Holder fields:

| Field     | Description                  |
| --------- | ---------------------------- |
| Address   | Address of the token holder. |
| NumTokens | Initial balance in mUSD.     |

#### Authenticated Data Feed System

The following fields are under review and they might change in the future depending on the economic context of the oracle activity.

| Field         | Description                                                                                         |
| ------------- | --------------------------------------------------------------------------------------------------- |
| maxNumOracles | Maximum number of oracles at one time. **Value must be > 0**.                                       |
| freezePeriod  | Period of time, in days, that coins remained locked as soon as an oracle decides stop the activity. |
| baseDeposit   | Minimum deposit, in kUSD, required to be an oracle.                                                 |

#### Mint KUSD

In order to **mint kUSD** to a specific account you must include a description as follows:

```
[[prefundedAccounts]]
accountAddress = "0xd6e579085c82329c89fca7a9f012be59028ed53f"
balance = 10
```

#### Sample

```
[genesis]
fileName = "genesis.json"
network = "test"
extraData = "Kowala's first block!"

[genesis.governance]
origin = "0x259be75d96876f2ada3d202722523e9cd4dd917d"
governors = ["0x259be75d96876f2ada3d202722523e9cd4dd917d"]
numConfirmations = 4

[genesis.consensus]
engine = "tendermint"
maxNumValidators = 1000
freezePeriod = 0
baseDeposit = 0
[[genesis.consensus.validators]]
address = "0xd6e579085c82329c89fca7a9f012be59028ed53f"
deposit = 123456

[genesis.consensus.token]
name = "mUSD"
symbol = "mUSD"
cap = 1073741824
decimals = 18
[[genesis.consensus.token.holders]]
address = "0xd6e579085c82329c89fca7a9f012be59028ed53f"
numTokens = 123456

[genesis.datafeed]
maxNumOracles = 1000
freezePeriod = 0
baseDeposit = 0

[[prefundedAccounts]]
accountAddress = "0xd6e579085c82329c89fca7a9f012be59028ed53f"
balance = 10
```

</br>
</br>
