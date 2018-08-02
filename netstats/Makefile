IMAGE_NAME      ?= kowalatech/netstats
URL             ?= "https://zygote.kowala.tech"
EXPLORER_URL    ?= "https://explorer.zygote.kowala.tech"
FAUCET_URL      ?= "https://faucet.zygote.kowala.tech"
CDN_URL         ?= "https://cdn.kowala.io"
WS_SECRET       ?= "abc123"

.PHONY: run run-docker build-docker

run: build-docker run-docker

run-docker:
	@docker run --rm -it -p 3000:3000 $(IMAGE_NAME)

build-docker:
	@docker build \
		-t $(IMAGE_NAME) \
		--build-arg URL=$(URL) \
		--build-arg CDN_URL=$(CDN_URL) \
		--build-arg EXPLORER_URL=$(EXPLORER_URL) \
		--build-arg FAUCET_URL=$(FAUCET_URL) \
		--build-arg WS_SECRET=$(WS_SECRET) \
		.
