Then('I should see the setup form') do
  page.find('h3', text: 'Create A New Account')
end

Then('I should see the {string} input field') do |textbox_text|
  find_textbox_with_text(textbox_text)
end

When('I enter a text less than {int} characters in the {string} input field') do |number_characters, textbox_text|
  enter_text_on_textbox(textbox_text, random_string_less_than_num_characters(number_characters))
  lose_input_focus(find_textbox_with_text(textbox_text))
end


Then('The {string} button should be disabled') do |button_text|
  expect(find_button_with_text(button_text)[:disabled]).to eq('true')
end

When('I enter a username that has already been used') do
  enter_text_on_textbox('Username', 'kowala-test-account-1')
end

When('I enter a valid username') do
  enter_text_on_textbox('Username', random_string(8))
end

Then('I should not see an error message') do
end

Then('The {string} button should be enabled') do |button_text|
  expect(find_button_with_text(button_text)[:disabled]).to be_nil
end

Then('I should see the password input fields') do
  expect(find_input_by_name('password')[:type]).to eq('password')
  expect(find_input_by_name('passwordConfirmation')[:type]).to eq('password')
end

def complete_username_step
  next_button_text = 'Next'
  enter_text_on_textbox('Username', random_string(8))
  find_button_with_text(next_button_text).click
  find_button_with_text(next_button_text).click
end

Given('I have entered a valid username and proceeded to the password step') do
  complete_username_step
end

When('I enter a password less than {int} characters') do |number_characters|
  textbox_text = 'Password'
  enter_text_on_textbox(textbox_text,
                        random_string_less_than_num_characters(number_characters))
  lose_input_focus(find_textbox_with_text(textbox_text))
end

def enter_valid_password
  textbox_text = 'Password'
  valid_password = '123456789'
  enter_text_on_textbox(textbox_text, valid_password)
  lose_input_focus(find_textbox_with_text(textbox_text))
end

When('I enter a valid password') do
  enter_valid_password
end

When('I enter an invalid password confirmation') do
  textbox_text = 'Password Confirmation'
  enter_text_on_textbox(textbox_text, random_string(8))
  lose_input_focus(find_textbox_with_text(textbox_text))
end

def enter_valid_password_confirmation
  textbox_text = 'Password Confirmation'
  enter_text_on_textbox(textbox_text, '123456789')
  lose_input_focus(find_textbox_with_text(textbox_text))
end

When('I enter an valid password confirmation') do
  enter_valid_password_confirmation
end

def complete_password_step
  next_button_text = 'Next'
  enter_valid_password
  enter_valid_password_confirmation
  find_button_with_text(next_button_text).click
end

Given('I have entered a valid username and passwords and proceeded to the pin step') do
  complete_username_step
  complete_password_step
end

Then('I should see the PIN input fields') do
  find_textbox_with_text('PIN')
  find_textbox_with_text('PIN Confirmation')
end

When('I enter a PIN less than {int} characters') do |number_characters|
  textbox_text = 'PIN'
  enter_text_on_textbox(textbox_text, random_string_less_than_num_characters(number_characters))
  lose_input_focus(find_textbox_with_text(textbox_text))
end

When('I enter a PIN more than {int} characters') do |number_characters|
  textbox_text = 'PIN'
  enter_text_on_textbox(textbox_text, random_string_more_than_num_characters(number_characters))
  lose_input_focus(find_textbox_with_text(textbox_text))
end

def entervalid_pin
  textbox_text = 'PIN'
  valid_pin = '1234'
  enter_text_on_textbox(textbox_text, valid_pin)
  lose_input_focus(find_textbox_with_text(textbox_text))
end

When('I enter a valid PIN') do
  entervalid_pin
end

When('I enter an invalid PIN confirmation') do
  textbox_text = 'PIN Confirmation'
  enter_text_on_textbox(textbox_text, random_string(4))
  lose_input_focus(find_textbox_with_text(textbox_text))
end

When('I enter a valid PIN confirmation') do
  textbox_text = 'PIN Confirmation'
  valid_pin = '1234'
  enter_text_on_textbox(textbox_text, valid_pin)
  lose_input_focus(find_textbox_with_text(textbox_text))
end

When('I enter a valid PIN and PIN confirmation') do
  valid_pin = '1234'
  enter_text_on_textbox('PIN', valid_pin)
  enter_text_on_textbox('PIN Confirmation', valid_pin)
end

Then('I should see a loading indicator') do
  all('div').each { |div| puts div[:role] }
end
