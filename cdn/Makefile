SASS_STYLE     ?= compressed
SASS_PATHS     ?= sass:assets/styles
SASS_INCLUDES  ?= -I node_modules/foundation-sites/scss -I node_modules/hover.css/scss -I node_modules/sass-burger

build: clean install build-font-icons copy-fonts build-css copy-images copy-pdfs

install:
	@npm install

serve: build
	@npm run serve

build-font-icons:
	@fontcustom compile fonts/fontcustom/vectors -c fonts/fontcustom/fontcustom.yml

build-css:
	@sass --style $(SASS_STYLE) $(SASS_INCLUDES) --update $(SASS_PATHS) -E "UTF-8"
	@for f in $(ls assets/styles/*.css); do gzip $f > $f.gz; done;

watch-css:
	@sass --style $(SASS_STYLE) $(SASS_INCLUDES) --watch $(SASS_PATHS) -E "UTF-8"

copy-fonts:
	@cp -r fonts/ assets/fonts/

copy-images:
	@cp -r images/ assets/

copy-pdfs:
	@cp -r pdfs/ assets/

clean:
	@rm -rf assets/

docker:
	docker build -t cdn .

docker-build: docker
	docker run --rm -v $(shell pwd):/cdn -it cdn make build

docker-serve: docker
	docker run --rm -v $(shell pwd):/cdn -p 9090:9090 -it cdn make serve
