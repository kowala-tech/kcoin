Feature: KNS interactions
  As a user
  I want to be able to interact with the KNS

  Scenario: Lookup addresses
    Given The KNS includes the following records:
      | name    | address | owned by me |
      | a.test  | A       | false       |
      | b.test  | B       | false       |
	Then The name 'a.test' should map to A in the KNS
	And The name 'b.test' should map to B in the KNS
 
  Scenario: Failing to alter records not owned by me
    Given The KNS includes the following records:
      | name    | address | owned by me |
      | a.test  | A       | false       |
    When I try to update 'a.test' to B in the KNS
    Then I should see an error from the KNS contract

  Scenario: Successfully update records owned by me
    Given The KNS includes the following records:
      | name    | address | owned by me |
      | a.test  | A       | true        |
    When I try to update 'a.test' to B in the KNS
	Then The name 'a.test' should map to B in the KNS

  Scenario: Successfully update system contracts:
	When I deploy a new mining token contract
	And I update KNS to map the new mining token contract
	Then My client should use the new mining token contract 
	# Note: this can be tested by deploying a fake version of the mining
	# token that always returns a particular value for a particular function.
	# For example balanceOf() could always return -1.
