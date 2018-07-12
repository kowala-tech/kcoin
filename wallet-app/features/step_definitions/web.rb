Given('I am on the {string} page') do |page|
  visit lookup_url(page)
end

Given('I do not have an account') do
end

Given("I have an account") do
end

def random_string_less_than_num_characters(num_characters)
  random_string(num_characters - 1)
end

def random_string_more_than_num_characters(num_characters)
  random_string(num_characters + 1)
end

def random_string(num_characters)
  rand(36**num_characters).to_s(36)
end

def lose_input_focus(element)
  element.send_keys(:tab)
end

def find_button_with_text(button_text)
  page.find('button', text: /#{button_text}$/i)
end

def find_textbox_with_text(textbox_text)
  page.find('label', text: /#{textbox_text}$/i).sibling('div').first('input')
end

def find_input_by_name(name)
  page.find("input[name=#{name}]")
end

def enter_text_on_textbox(textbox_to_find, text_to_enter)
  find_textbox_with_text(textbox_to_find).set(text_to_enter)
end

And /^(?:I )?click the "([^"]+)" button$/ do |button_text|
  find_button_with_text(button_text).click
end

And /^(?:I )?click the "([^"]+)" link$/ do |button_text|
  find_button_with_text(button_text).click
end

Then('I should see a {string} error message') do |message|
  expect(page.find('p', text: /#{message}/i)[:class]).to match(/error/i)
end

Then('I should see a {string} message') do |message|
  page.find('p', text: /#{message}/i)
end

Then('I should be redirected to the {string} page') do |target_page|
  expect(page).to have_current_path(lookup_url(target_page))
end

def lookup_url(page_name)
  case page_name
  when 'setup'
    '/'
  when 'root'
    '/'
  when 'login'
    '/login'
  when 'wallet'
    '/wallet'
  else
    ''
  end
end

After do
  if Capybara.current_driver == :selenium && !Capybara.current_url.start_with?('data:')
    page.execute_script <<-JAVASCRIPT
      localStorage.clear();
      sessionStorage.clear();
    JAVASCRIPT
  end
end
