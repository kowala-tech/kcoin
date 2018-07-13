#encoding: UTF-8
require 'cucumber'
require 'rspec'
require 'selenium-webdriver'
require 'capybara'
require 'capybara/dsl'

Capybara.register_driver :selenium do |app|
  
    arguments = ["headless","disable-gpu", "no-sandbox", "window-size=1920,1080", "privileged"]
    # arguments = ["start-maximized"]
    preferences = {
        'download.prompt_for_download': false,
    }
    options = Selenium::WebDriver::Chrome::Options.new(args: arguments, prefs: preferences)
    Capybara::Selenium::Driver.new(app, browser: :chrome, options: options)

end

Capybara.run_server = false
Capybara.default_driver = :selenium
Capybara.javascript_driver = :selenium
Capybara.default_selector = :css
Capybara.default_max_wait_time = 10

Capybara.app_host = 'http://kowala-react-web:3000'
# Capybara.app_host = 'http://localhost:3000'

World(Capybara::DSL)
puts "running on browser: chrome"