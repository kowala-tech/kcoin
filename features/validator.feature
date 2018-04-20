Feature: Joining network as a validator
  As a user
  I want to be able to join validators set

  Scenario: Start validator
    Given I have the following accounts:
      | account | funds |
      | A       | 10    |
    When I start validator with 1 deposit and coinbase A
    Then I should be a validator
    And the balance of A should be around 9 kcoins
