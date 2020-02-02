Feature: Sending and receiving transactions
  As a user
  I want to be able to send and receive transactions

  Background:
    Given the network is running

  Scenario: Send 1 kcoin successfully
    Given I have the following accounts:
      | account | password | funds |
      | A       | test     | 10    |
      | B       | test     | 5     |
    When I unlock the account A with password 'test'
    And I transfer 1 kcoin from A to B
    Then the balance of A should be around 9 kcoins
    And the balance of B should be 6 kcoins
    And only one transaction should be done
    And the transaction hash the same

  Scenario: Not enough funds
    Given I have the following accounts:
      | account | password | funds |
      | A       | test     | 10    |
      | B       | test     | 5     |
    When I unlock the account A with password 'test'
    And I try to transfer 11 kcoins from A to B
    Then the transaction should fail
    And the balance of A should be 10 kcoins
    And the balance of B should be 5 kcoins
