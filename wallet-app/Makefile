DOCKER_NAME = kowala-app-web

start: .docker-started docker-install docker-serve

install:
	@yarn install

build:
	@yarn run build

serve:
	@yarn run start

watch:
	@yarn run start

test:
	@yarn run test
	@cucumber --tags 'not @ignore'

lint:
	@yarn run lint

storybook:
	@yarn run storybook

clean:
	@rm -rf node_modules/
	@rm -rf dist/
	@rm -rf coverage/

docker-install: .docker-started
	@docker-compose exec $(DOCKER_NAME) make install

docker-serve: .docker-started
	@docker-compose exec $(DOCKER_NAME) make serve

docker-build: .docker-started
	@docker-compose exec $(DOCKER_NAME) make build

docker-test: .docker-started
	@docker-compose exec $(DOCKER_NAME) make test

docker-storybook: .docker-started
	@docker-compose exec $(DOCKER_NAME) make storybook

docker-bash: .docker-started
	@docker-compose exec $(DOCKER_NAME) /bin/bash

docker-stop:
	@docker-compose stop
	@docker-compose down
	@rm .docker-started

docker-clean: docker-stop
	@docker-compose kill
	@docker-compose rm -f

.docker-started:
	@docker-compose build
	@docker-compose up -d
	@touch .docker-started
	@echo "Docker containers are now running."
