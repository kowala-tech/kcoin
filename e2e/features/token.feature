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
    Given I start validator with 5 mTokens deposit
    And I wait for my node to be synced
    And the deposit of A should be 15 mTokens
    And I unlock the account A with password 'test'
    When I transfer 10 mTokens from A to B
    Then the deposit of A should be 5 mTokens
    And the deposit of B should be 20 mTokens

  Scenario: Transfer tokens: empty receiving account
    Given I start validator with 5 mTokens deposit
    And I wait for my node to be synced
    And the deposit of A should be 15 mTokens
    And I unlock the account A with password 'test'
    When I transfer 10 mTokens from A to C
    Then the deposit of A should be 5 mTokens
    And the deposit of B should be 10 mTokens

  Scenario: Mint tokens: consensus established
    Given I start validator with 5 mTokens deposit
    And I wait for my node to be synced
    When 2 of 3 governance accounts mint 1 mToken to C
    Then the deposit of C should be 1 mTokens

  Scenario: Mint tokens: failed consensus
    Given I start validator with 5 mTokens deposit
    And I wait for my node to be synced
    When 1 of 3 governance accounts mints 1 mToken to C
    Then the deposit of C should be 0 mTokens