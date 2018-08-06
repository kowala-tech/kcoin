DOCKERHUB_IMAGE ?= kowalatech/wallet-tools
APP_URL         ?= "https://wallet-tools.zygote.kowala.io" 
CDN_URL         ?= "https://cdn.kowala.io"

.PHONY: run run-docker build-docker

run: build-docker run-docker 

run-docker: 
	@docker run --rm -it -p 443:443 -p 8080:80 $(DOCKERHUB_IMAGE)

build-docker:
	@docker build -t $(DOCKERHUB_IMAGE) --build-arg APP_URL=$(APP_URL) --build-arg CDN_URL=$(CDN_URL) . 
