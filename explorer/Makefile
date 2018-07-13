.PHONY: start stop build

start:
	docker-compose up -d

stop:
	docker-compose stop

build:
	docker build . -f web.Dockerfile -t kowalatech/kexplorer
	docker build . -f sync.Dockerfile -t kowalatech/kexplorersync
