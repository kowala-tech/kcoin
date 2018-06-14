Feature: Joining network as a validator
  As a user
  I want to be able to join validators set

  Background:
    Given I have the following accounts:
      | account | password | funds |
      | A       | test     | 20    |
      | B       | test     | 10    |

  Scenario: Start validator
    Given I have my node running using account A
    When I start validator with 5 kcoins deposit
    And I wait for my node to be synced
    Then the balance of A should be around 15 kcoins

  Scenario: Stop mining
    Given I have my node running using account A
    And I start validator with 5 kcoins deposit
    And I wait for my node to be synced
    And the balance of A should be around 15 kcoins
    When I withdraw my node from validation
    Then there should be 5 kcoins available to me after 5 days

   Scenario: Mining rewards: basic
    Given I have my node running using account A
    And I start validator with 5 kcoins deposit
    And I wait for my node to be synced
    And the balance of A should be around 15 kcoins
    When I unlock the account A with password 'test'
    And I transfer 10 kcoin from A to B
    Then the balance of A should be greater 5 kcoins
    And the balance of B should be 20 kcoins
