DOCKERHUB_IMAGE ?= kowalatech/wallet-tools

.PHONY: run-docker build-docker

run-docker: build-docker
	@docker run --rm -it -p 443:443 -p 8080:80 $(DOCKERHUB_IMAGE)

build-docker:
	@docker build -t $(DOCKERHUB_IMAGE) . 
