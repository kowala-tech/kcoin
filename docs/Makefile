.PHONY: server-with-docker

PWD   := $(shell pwd)

server-with-docker:
	@docker run --rm -i -t -p 8000:8000 -v $(PWD):/documents kowalatech/mkdocs 
