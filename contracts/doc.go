/*
Package contracts implements the network core contracts


Glossary

Smart contract - Kowala account that contains code that can be executed.

Fallback function - is the function called when one sends funds to a contract
without providing any data. Ex: function() payable {}. This function is required
in such scenario.


Interacting with a contract

In order to interact with a contract using golang, we must generate bindings for
golang. The process is based on the go generate command. By default, the
contract should be included in a folder called contract and the wrapper type
should be created in the root folder. The wrapper type contains the instructions
to generate the bindings,

Example:

wrapper_type.go
//go:generate solc --abi --bin --overwrite -o build contracts/Election.sol
//go:generate abigen -abi build/Election.abi -bin build/Election.bin -pkg
contracts -type ElectionContract -out contracts/election.go

> cd root
> go generate

In order to submit changes to a contract(ex: register a candidate as a
validator), the user has to make a transaction. Each transaction costs a certain
amount of gas(tx fee). If gas prices and gas limits are not provided, the node
will calculate these values by picking default values and executing the
transaction locally against the latest state. In order to read the contract
state the user does not pay anything.

The golang bindings hide most of these details and usually we just need to make
sure that the sender has enough balance for transactions.


Contract Testing

- You should take into account gas consumption while creating contracts.
- Private methods cannot be tested - There's a workaround which involved
creating libraries(different contract) but this option has an negative impact on
the tx fees. It might be the best option in some cases though.


Common Pitfals

- The transaction sender does not have enough funds to deploy the contract and
execute other transactions.
- State might be inconsistent because you forgot to commit a transaction.
- In order to transfer money to a contract without specifying a method, you must
include a fallback function in the contract in order to be able to transfer
funds to the contract. 
- The block has a gas limit block, so, that means that it's possible that an
operation takes more gas than the limit. The limit will probably change in the
future and it's something that we need to benchmark with e2e tests.