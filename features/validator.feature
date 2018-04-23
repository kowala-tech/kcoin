Feature: Joining network as a validator
  As a user
  I want to be able to join validators set

  Scenario: Start validator
    Given I have the following accounts:
      | account |  funds  |
      | A       | 1000000 |
    When I start validator with 1 deposit and coinbase A
    Then I should be a validator
    And the balance of A should be around 9 kcoins

  Scenario: Stop validator
    Given I have the following accounts:
      | account |  funds  |
      | A       | 1000000 |
    When I start validator with 1 deposit and coinbase A
    And I stop validation
    And I wait for the unbonding period to be over
    Then I should not be a validator
    And the balance of A should be around 9 kcoins

  # Scenario: Mining rewards: basic
  #   Given There is a network
  #   And I have an existing node connected to the network
  #   And My node is validating with all the issues tokens
  #   And There are no other validators
  #   And The current block reward is 100 # if it's easier this could be market price
  #   And I have the following accounts:
	# 	  | account | funds |
	# 	  | A       | 0     |
	# 	  | B       | 0     |
	# 	And My node pays out rewards to the following addresses
	# 	  | account | share |
	# 	  | A       | 80    |
	# 	  | B       | 20     |
  #   And there are no other transactions
  #   And there is no stability fee
  #   When A new block is mined
  #   Then the balance of A should be 80 kcoins
  #   And the balance of B should be 20 kcoins
