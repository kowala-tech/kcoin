pragma solidity ^0.4.18;

contract NetworkStats {
    // Total supply of wei. Must be updated every block and initialized to the correct value.
    uint256 public totalSupplyWei = 1 ether;
    // Reward calculated for the last block. Must be updated every block.
    uint256 public lastBlockReward = 0;
    // Price established by the price oracle for the last block. Must be updated every block.
    uint256 public lastPrice = 0;
    // Tendermint validators.
	address[] public tendermintValidators;

    function NetworkStats() public {
        tendermintValidators.push(0x0D4CA5AF584E49AB6D08EB0A8C6AD73A41AA74D8);
        tendermintValidators.push(0x29EE62EB3A8322E7FDDB548E8A1FA62871027CD4);
        tendermintValidators.push(0x98328A8723275E9588CFC6ABD71E93C3000BD7B5);
        tendermintValidators.push(0xAE1B3B25B26E71343EDA6744F88D9D98DF141D2F);
        tendermintValidators.push(0xB28FC698F28A8ADC2F38CC8A16B87FA709ADE0FF);
        tendermintValidators.push(0xC57BF12BB34F6FD85BDBF0CACA983528422BF7A2);    }
}
