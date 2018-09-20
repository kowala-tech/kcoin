Feature: Using the mocked exchange backend
  As a user I want to use the backend
  To perform fake controlled requests

  Background:
    Given the network is running
    And the mocked exchange is running

  Scenario: I fetch the data I want and I get it later
    When I fetch the exchange mock with data:
      | type | amount | rate |
      | buy  | 1      | 2    |
      | buy  | 1.2    | 2.9  |
      | sell | 1      | 2    |
      | sell | 1.35   | 3.4  |
