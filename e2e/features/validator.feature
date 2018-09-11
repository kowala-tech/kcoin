Feature: Joining network as a validator
  As a user
  I want to be able to join validators set

  Background:
    Given the network is running
    And I have the following accounts:
      | account | password | tokens | funds | validator |
      | A       | test     | 20     | 10    | true      |
      | B       | test     | 10     | 10    | false     |

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

  Scenario: Once again
    Given: I wait for my node to be synced
    And I start validator with 5 mTokens deposit
    And I wait for my node to be synced
    And crash my node validator
    And I restart the validator
    And I wait for my node to be synced
    And I start validator with 5 mTokens deposit
    And the token balance of A should be 15 mTokens
