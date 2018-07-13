# Generated from the Google Spreadsheet. Editing may not be a good idea!

Feature: Setup a new account
  As a New User
  I want to Setup a new account
  So that I can create a wallet and use kUSD

  Background:
     Given I do not have an account
       And I am on the "setup" page
#   @ignore
  Scenario: Form renders
      Then I should see the setup form
       And I should see the "Username" input field
#   @ignore
  Scenario: Enter invalid username (length)
      When I enter a text less than 8 characters in the "Username" input field
      Then I should see a "Username must be at least 8 characters" error message
       And The "Next" button should be disabled
#   @ignore
  Scenario: Enter invalid username (uniqueness)
      When I enter a username that has already been used
      And click the "Next" button
      Then I should see a "Username is already taken" error message
      And The "Next" button should be disabled
#   @ignore
  Scenario: Enter valid username
      When I enter a valid username
      And click the "Next" button
      Then I should not see an error message
      And The "Next" button should be enabled
# @ignore
  Scenario: Advancing to the password input fields
      When I enter a valid username
       And I click the "Next" button
       And I click the "Next" button
      Then I should see the password input fields
# @ignore
  Scenario: Return to username field with "Back" button
      When I enter a valid username
       And I click the "Next" button
       And I click the "Next" button
       And I click the "Back" button
      Then I should see the "Username" input field
# @ignore
  Scenario: Enter invalid password (length)
      Given I have entered a valid username and proceeded to the password step
      When I enter a password less than 8 characters
      Then I should see a "Password must be at least 8 characters" error message
      And The "Next" button should be disabled
# @ignore
  Scenario: Enter invalid password confirmation (match)
      Given I have entered a valid username and proceeded to the password step
      When I enter a valid password
       And I enter an invalid password confirmation
      Then I should see a "Passwords do not match." error message       
       And The "Next" button should be disabled
# @ignore
  Scenario: Enter valid password and password confirmation
      Given I have entered a valid username and proceeded to the password step  
      When I enter a valid password
       And I enter an valid password confirmation
      Then I should not see an error message
       And The "Next" button should be enabled
# @ignore
  Scenario: Advancing to the PIN input fields
      Given I have entered a valid username and passwords and proceeded to the pin step    
      Then I should see the PIN input fields
# @ignore
  Scenario: Return to password input fields with "Back" button
       Given I have entered a valid username and passwords and proceeded to the pin step    
       And I click the "Back" button
      Then I should see the password input fields
# @ignore
  Scenario: Enter invalid PIN (length - short)
      Given I have entered a valid username and passwords and proceeded to the pin step
      When I enter a PIN less than 4 characters
      Then I should see a "PIN must 4 digits." error message
       And The "Finish" button should be disabled
# @ignore
  Scenario: Enter invalid PIN (length - long)
      Given I have entered a valid username and passwords and proceeded to the pin step
      When I enter a PIN more than 4 characters
      Then I should see a "PIN must 4 digits." error message
       And The "Finish" button should be disabled
# @ignore
  Scenario: Enter invalid PIN confirmation (match)
      Given I have entered a valid username and passwords and proceeded to the pin step
      When I enter a valid PIN
       And I enter an invalid PIN confirmation
      Then I should see a "PINs do not match." error message
       And The "Finish" button should be disabled
# @ignore
  Scenario: Enter valid PIN and PIN confirmation
      Given I have entered a valid username and passwords and proceeded to the pin step  
      When I enter a valid PIN
       And I enter a valid PIN confirmation
      Then I should not see an error message
       And The "Finish" button should be enabled
@ignore
  Scenario: Completing signup
      Given I have entered a valid username and passwords and proceeded to the pin step    
      When I enter a valid PIN and PIN confirmation
       And I click the "Finish" button       
    #    Then I should see a "Your account is being created" message
    #    Then I should see a loading indicator
    #    Then I should be redirected to the "login" page