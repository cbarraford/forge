.PHONY: test-docker test lint

test-docker:
	@docker-compose run --rm forge go test ./...

test:
	go test ./...

lint:
	@docker-compose run --rm lint run ./...
