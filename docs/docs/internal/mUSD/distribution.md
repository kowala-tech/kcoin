# Token Distribution

Token distribution from a technical point of view, is the activity of updating the balances of the token contract (blockchain state) to match the investments made by the investors hence the distribution just happens as soon as the production network is available as this network will represent the final state. There are some alternatives to this model, such as the EOS model, which relied on an Ethereum crowdsale to sell their tokens given that their blockchain was not available at the time - the state is mapped from Ethereum to EOS's blockchain later on. Kowala's use case is different, since the crowdsale is being done off-chain. This means that we don't have means to identify the investors addresses since the transactions were not done on a blockchain that supports crowdsales such as Ethereum. Given this scenario we can only provide the tokens as soon as our main network is running. Alternatively, we can create a contract in Ethereum based on the addresses provided by the investors and later on map this to our blockchain but in the end it's the same unless we need funds from a crowdsale such as in EOS case which is not the case for kowala since we have a private sale off-chain and we already have the investors funds.

# Investor role

In order to distribute the tokens, Kowala requires an address per investor. Each investor is responsible for providing an address. Nonetheless, Kowala should provide some recommendations. This activity can be done at any point in time but we should do it as soon we have a reliable UI. An investor should not be forced to use our software to generate a new key (account) as this might be seen as a security risk on their end. From our end, we just need to know the address which will be linked to the tokens acquired by the investor.

</br>
</br>
