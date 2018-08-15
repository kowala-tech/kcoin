Feature: Using mTokens
  As a user
  I want to be able to transfer and mint mTokens

  Background:
    Given the network is running
    And I have the following accounts:
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
