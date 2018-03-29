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
    Then the balance of A is eventually around 9 kcoin
    And the balance of B is eventually 6 kcoin
