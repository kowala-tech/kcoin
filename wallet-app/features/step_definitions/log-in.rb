Given('I am not logged in') do
end

Then('I should see a {string} link') do |link_text|
  page.find('button', text: /#{link_text}$/i)
end

Then('I should see a {string} button') do |link_text|
  page.find('button', text: /#{link_text}$/i)
end

Then('I should see the log in form') do
  page.find('h3', text: 'Login To Your Account')
end

When('I enter my username') do
  enter_text_on_textbox('Username', 'kowala-test-account-1')
end

When('I enter my password') do
  enter_text_on_textbox('Password', 'Password1234')
end

When('I enter an incorrect password') do
  enter_text_on_textbox('Password', random_string(8))
end

Then('I should remain in the {string} page') do |target_page|
  expect(page).to have_current_path(lookup_url(target_page))
end
