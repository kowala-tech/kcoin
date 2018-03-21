Feature: Send and receive transactions
  In order to use my coins
  As a kowala user
  I need to be able to send and receive transactions

  Scenario: Send 1 kUSD
    Given I have an account A with 10 kUSD
    And I have an account B with 5 kUSD
    When I transfer 1 kUSD from A to B
    Then the balance of A is 9 kUSD
    And the balance of B is 6 kUSD
