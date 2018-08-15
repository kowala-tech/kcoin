Feature: Using the faucet
  As a user
  I want to be able to use the faucet

  Background:
    Given the network is running

  Scenario: Open the faucet
    Given I have the following accounts:
      | account | password | funds |
      | A       | test     | 10    |
    And the faucet node is running using the account A and password 'test'
    When I fetch / on the faucet
    Then the status code is 200
