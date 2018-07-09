PWD   := $(shell pwd)
BUILD := mkdocs build --clean --strict -d site

server-with-docker:
	@docker run --rm -i -t -p 8000:8000 -v $(PWD):/documents kowalatech/mkdocs 

build-with-docker:
	@docker run --rm -v $(PWD):/documents kowalatech/mkdocs $(BUILD)

build:
	@$(BUILD)
