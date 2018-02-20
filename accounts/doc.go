/*
Package accounts manages wallets and its account(s).
Here we focus on externally owned accounts which are usually referred to as accounts - think bank account.
Accounts play a central role in Kowala - the state of all accounts is the state of the Kowala network which is updated with every block.
The accounts package also includes the ABI package - defines the structures and methods that you will use to interact with binary contract.

Account

Accounts are essential for users to interact with the Kowala blockchain via transactions.
There are two types of accounts: externally owned accounts and contract accounts - accounts have balance and contracts have both balance and contract storage.
Externally owned accounts are usually refererred to as accounts and contract accounts as contracts - contracts are covered in the contracts package.
Accounts are used, for example, to sign transactions so that the EVM can securely validate the identity of a transaction sender.

Example 1: send funds (console)
> eth.sendTransaction({from:eth.coinbase, to:"0xcfff0fdae894be2ed95e02f514b3fbfc1bf41656", value: web3.toWei(0.05, "ether")})

Account Address & Keyfiles

Every account is defined by a pair of keys, a private key and a public key.
Accounts are indexed by their address which is derived from the public key that defines the account.
Every private key(encrypted)/address pair is encoded in a keyfile - JSON text file that can be found in the keystore subdirectory of your Kowala's node data directory (default: .kowala/keystore).
Make sure you backup your keyfiles regularly!

Creating an account

Once you have the kusd client installed, creating an account is merely a case of executing the "kusd account new" command in a terminal.

Wallet

A wallet represents a software or hardware wallet that might contain one or more accounts (derived from the same seed).
The current codebase (1.73) supports the following hardware wallets: Ledger Blue & Ledger Nano S

Application Binary Interface (/abi)

A Kowala smart contract is bytecode deployed on the Kowala blockchain. There could be several functions in a contract.
An ABI is necessary so that you can specify which function in the contract to invoke, as well as get a guarantee that the function will return data in the format you are expecting.

Example 1: generating an abi for the network contract
//go:generate abigen -abi build/Network.abi -bin build/Network.bin -pkg network -type NetworkContract -out gen_network.go
*/

package accounts
