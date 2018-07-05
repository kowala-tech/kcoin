IMAGE_NAME     := kowala-app-web
CUCUMBER_IMAGE_NAME     := kowala-app-web-tests

install:
	@yarn install --network-concurrency 1

build:
	@yarn run build

test:
	@yarn run test

lint:
	@yarn run lint

start:
	KOWALA_NETWORK="kusd-zygote" yarn run start

storybook:
	@yarn run storybook

build-storybook:
	@yarn run storybook

clean:
	@rm -rf node_modules/
	@rm -rf dist/
	@rm -rf coverage/

cucumber-test:
	@cucumber --tags 'not @ignore'

test-e2e:
	@docker-compose exec $(CUCUMBER_IMAGE_NAME) make cucumber-test

.started:
	@docker-compose build
	@docker-compose up -d
	@touch .started
	@echo "Docker containers are now running."

start-docker-images: .started

# Alias for watch
serve: watch

# Start the website on port 3000
watch: .started
	@docker-compose exec $(IMAGE_NAME) make install
	@docker-compose exec $(IMAGE_NAME) yarn run start:server

stop:
	@docker-compose kill
	@docker-compose stop
	@docker-compose down
	@docker-compose rm -f
	-@rm .started
