Feature: Managing accounts
  As a user
  I want to be able to manage my account

  Background:
    Given the network is running

  Scenario: I can unlock my account
    Given I created an account with password '12345'
    When I try to unlock my account with password '12345'
    Then I should get my account unlocked

  Scenario: My account doesn't unlock using a wrong password
    Given I created an account with password '12345'
    When I try to unlock my account with password 'wrong'
    Then I should get an error unlocking the account
