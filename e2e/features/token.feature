Feature: Using mTokens
  As a user
  I want to be able to transfer and mint mTokens

  Background:
    Given I have the following accounts:
      | account | password | tokens | funds | validator |
      | A       | test     | 20     | 10    | true      |
      | B       | test     | 10     | 10    | false     |
      | C       | test     | 0      | 10    | false     |

  Scenario: Transfer tokens: not-empty receiving account
    Given I wait for my node to be synced
    And I start validator with 5 mTokens deposit
    And the token balance of A should be 15 mTokens
    And I unlock the account A with password 'test'
    When I transfer 9 mTokens from A to B
    Then the token balance of A should be 6 mTokens
    And the token balance of B should be 19 mTokens

  Scenario: Transfer tokens: empty receiving account
    Given I wait for my node to be synced
    And I start validator with 5 mTokens deposit
    And the token balance of A should be 15 mTokens
    And I unlock the account A with password 'test'
    When I transfer 9 mTokens from A to C
    Then the token balance of A should be 6 mTokens
    And the token balance of C should be 9 mTokens


  #  Scenario: Mint tokens: consensus established
  #    Given I wait for my node to be synced
  #    And I start validator with 5 mTokens deposit
  #    When 2 of 3 governance accounts mint 1 mToken to C
  #    Then the token balance of C should be 1 mTokens

  #  Scenario: Mint tokens: failed consensus
  #    Given I wait for my node to be synced
  #    Given I start validator with 5 mTokens deposit
  #    When 1 of 3 governance accounts mints 1 mToken to C
  #    Then the token balance of C should be 0 mTokens