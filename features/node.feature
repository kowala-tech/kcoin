Feature: Joining network
  As a node maintainer
  I want to be able to connect to a network

  Scenario: Connect a node to the network
    When I start a new node
    Then My node should sync with the network
 
  Scenario: Disconnect and reconnect a node
    Given I start a new node
    And My node is already synchronised
    When I disconnect my node for 2 blocks and reconnect it
    Then My node should sync with the network

  Scenario: Wrong network ID
    When I start a new node with a different network ID
    Then My node should not sync with the network

  Scenario: Wrong chain ID
    When I start a new node with a different chain ID
    Then My node should not sync with the network

  # Scenario: Wrong protocol version (Note: might be hard to test)
  #   When I start a new node with an old protocol version
  #   Then My node should not sync with the network
