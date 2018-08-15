Feature: Using mTokens
  As a user
  I want to be able to transfer and mint mTokens

  Background:
    Given I generate a genesis with 2 required signatures in the multisig contract
    And the network is running
    And I have the following accounts:
      | account | password | tokens | funds | validator |
      | A       | test     | 0      | 10    | false     |

  Scenario: Mint tokens: consensus established
    When 2 of 3 governance accounts mint 1 mToken to A
    Then the token balance of A should be 1 mTokens
  
  Scenario: Mint tokens: failed consensus
    When 1 of 3 governance accounts mints 1 mToken to A
    Then the token balance of A should be 0 mTokens
