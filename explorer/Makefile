.PHONY: mongo

mongo:
	@docker run --name mongoexplorer -p 27017:27017 -d mongo || docker start mongoexplorer

stopmongo:
	@docker stop mongoexplorer
