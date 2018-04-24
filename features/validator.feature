Feature: Joining network as a validator
  As a user
  I want to be able to join validators set

  Scenario: Start validator
    Given I have my node running
    And I have an account in my node
    When I start validator with 5000001 deposit
    Then I should be a validator
