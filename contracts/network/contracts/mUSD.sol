pragma solidity ^0.4.18;

import "./mToken.sol";

contract mUSD is mToken {
    function mUSD() public {
        name = "mUSD";
        symbol = "mUSD";
        maxTokens = 1073741824;
        // Initial token holders
        mintTokens(0x0D4CA5AF584E49AB6D08EB0A8C6AD73A41AA74D8, 100);
        mintTokens(0x29EE62EB3A8322E7FDDB548E8A1FA62871027CD4, 100);
        mintTokens(0x98328A8723275E9588CFC6ABD71E93C3000BD7B5, 100);
        mintTokens(0xAE1B3B25B26E71343EDA6744F88D9D98DF141D2F, 100);
        mintTokens(0xB28FC698F28A8ADC2F38CC8A16B87FA709ADE0FF, 100);
        mintTokens(0xC57BF12BB34F6FD85BDBF0CACA983528422BF7A2, 100);
    }
}