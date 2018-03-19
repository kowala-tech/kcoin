/*
Package consensus implements the consensus election contract.
The contract has an owner and it's responsible for managing the consensus
validator set.


Glossary

Base deposit - represents the deposit that a candidate has to do in order to
secure a place in the elections (if there are positions available).

Minimum deposit -

Unbonding period - is a predetermined period of time that coins remain locked,
starting from the moment a validator decides to leave the consensus elections.




*/

package consensus
