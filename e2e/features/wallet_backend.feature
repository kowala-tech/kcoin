Feature: Using the wallet backend
   As a wallet user
   I want to perform wallet operations

  Background:
    Given the network is running
    And the wallet backend node is running

  Scenario: I can get the block height
    Given I check the current block height in the wallet backend API
    When I wait for 2 blocks
    Then the new block height in the wallet backend API has increased by at least 2

  Scenario: I can get the list of transactions
    Given I have the following accounts:
      | account | password | funds |
      | A       | test     | 10    |
      | B       | test     | 5     |
    When I unlock the account A with password 'test'
    And I transfer 1 kcoin from A to B
    Then the transactions of A in the wallet backend API contains the last transaction
    And the transactions of B in the wallet backend API contains the last transaction

  Scenario: I can get the balance of my account
    Given I have the following accounts:
      | account | password | funds |
      | A       | test     | 10    |
      | B       | test     | 5     |
    When I unlock the account A with password 'test'
    And I transfer 1 kcoin from A to B
    Then the balance of A using the wallet backend should be around 9 kcoins
    And the balance of B using the wallet backend should be 6 kcoins

  Scenario: I can send transactions
    Given I have the following accounts:
      | account | password | funds |
      | A       | test     | 10    |
      | B       | test     | 5     |
    When I unlock the account A with password 'test'
    And I transfer 1 kcoin from A to B using the wallet API
    Then the balance of A should be around 9 kcoins
    And the balance of B should be 6 kcoins
