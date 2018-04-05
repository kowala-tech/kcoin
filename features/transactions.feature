Feature: Send and receive transactions
In order to use my coins
As a kowala user
I need to be able to send and receive transactions

  Scenario: Send 1 kcoin
    Given I have the following accounts:
      | account | funds |
      | A       | 10    |
      | B       | 5     |
    When I transfer 1 kcoin from A to B
    Then the last transaction is successful
    And the balance of A is around 9 kcoin
    And the balance of B is 6 kcoin

  Scenario: Not enough funds
    Given I have the following accounts:
      | account | funds |
      | A       | 10    |
      | B       | 5     |
    When I transfer 11 kcoin from A to B
    Then the last transaction failed
    And the balance of A is 10 kcoin
    And the balance of B is 5 kcoin
