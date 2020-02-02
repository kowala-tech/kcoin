Feature: Joining network
  As a node maintainer
  I want to be able to connect to a network

  Background:
    Given the network is running

  Scenario: Connect a node to the network
    When I start a new node
    Then my node should sync with the network

  Scenario: Disconnect and reconnect a node
    Given I start a new node
    And my node is already synchronised
    When I disconnect my node for 2 blocks and reconnect it
    Then my node should sync with the network

  Scenario: Wrong network ID
    When I start a new node with a different network ID
    Then my node should not sync with the network

  Scenario: Wrong chain ID
    When I start a new node with a different chain ID
    Then my node should not sync with the network

  # Scenario: Wrong protocol version (Note: might be hard to test)
  #   When I start a new node with an old protocol version
  #   Then my node should not sync with the network
