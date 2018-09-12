Feature: Joining network as a validator
  As a user
  I want to be able to join validators set

  Background:
    Given the network is running
    And I have the following accounts:
      | account | password | tokens | funds | validator |
      | A       | test     | 20     | 10    | true      |
      | B       | test     | 10     | 10    | false     |

  Scenario: Start validator
    Given I wait for my node to be synced
    When I start validator with 5 mTokens deposit
    Then the token balance of A should be 15 mTokens

  Scenario: Stop mining
    Given I wait for my node to be synced
    And I start validator with 5 mTokens deposit
    And the token balance of A should be 15 mTokens
    When I withdraw my node from validation
    Then there should be 5 mTokens available to me after 5 days

   Scenario: Mining rewards: basic
    Given I wait for my node to be synced
    And I start validator with 5 mTokens deposit
    And the token balance of A should be 15 mTokens
    And I unlock the account A with password 'test'
    When I transfer 10 mTokens from A to B
    Then the token balance of A should be 5 mTokens
    And the token balance of B should be 20 mTokens

  Scenario: Re-Start mining
    Given I wait for my node to be synced
    And I start validator with 5 mTokens deposit
    And I wait for my node to be synced
    And the token balance of A should be 15 mTokens
    When I withdraw my node from validation
    Then there should be 5 mTokens available to me after 5 days
    And I wait for my node to be synced
    And I start validator with 5 mTokens deposit
    And the token balance of A should be 10 mTokens
