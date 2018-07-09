# Account

Accounts play a central role in Kowala - the state of all accounts is the state
of the Kowala network which is updated with every block.

Accounts are essential for users to interact with the Kowala blockchain via
transactions. There are two types of accounts:

- **externally owned accounts** - have balance (think bank account)
- **contract accounts** - have both balance and contract storage.

Externally owned accounts are usually refererred to as accounts and contract
accounts as contracts
Accounts are used, for example, to sign transactions so that the EVM can
securely validate the identity of a transaction sender.

## Address and key files

Every account is defined by a pair of keys, a private key and a public key.
Accounts are indexed by their address which is derived from the public key that
defines the account. Every private key(encrypted)/address pair is encoded in a
keyfile - JSON text file that can be found in the keystore subdirectory of your
Kowala's node data directory (default: .kowala/keystore). **Make sure you backup
your keyfiles regularly! **

## Contract Accounts

The best way to describe contracts is to compare the technology to a vending
machine. Ordinarily, you would go to a lawyer or a notary, pay them, and wait
while you get the document. With smart contracts, you simply drop a kcoin into
the vending machine (i.e. ledger), and your escrow, driver’s license, or
whatever drops into your account. More so, ontracts not only define the rules
and penalties around an agreement in the same way that a traditional contract
does, but also automatically enforce those obligations.

Use cases:

- Function as 'multi-signature' accounts, so that funds are spent only when a
  required percentage of people agree.
- Manage agreements between users, say, if one buys insurance from the other.
- Store information about an application, such as domain registration
  information or membership records (Ex: consensus validators).

</br></br>
