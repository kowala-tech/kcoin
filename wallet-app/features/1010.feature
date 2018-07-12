# Generated from the Google Spreadsheet. Editing may not be a good idea!

Feature: Log in to my account
  As a New User
  I want to Log in to my account
  So that I can access my wallet from my current device

Assumptions / Technical Notes 
  Background:
     Given I am not logged in
       And I am on the "root" page

#   @ignore
  Scenario: Log in link
      Then I should see a "log in" link

#   @ignore
  Scenario: Going to log in page
      When I click the "log in" link
      Then I should see the log in form
  
#   @ignore
  Scenario: Incorrect credentials
   Given I am not logged in
   And I am on the "root" page
      When I click the "log in" link
       And I enter my username
       And I enter an incorrect password
       And I click the "LOGIN" button
      Then I should see a "Invalid password" message
       And I should remain in the "login" page

  @ignore
  Scenario: Returning to setup page
      When I click the "log in" link
      Then I should see a "back" button
       And I click the "back" button
      Then I should be redirected to the "setup" page

# @ignore
  Scenario: Correct credentials
      When I click the "log in" link
       And I enter my username
       And I enter my password
       And I click the "LOGIN" button
       And I should be redirected to the "wallet" page