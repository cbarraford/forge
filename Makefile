.PHONY: test-docker test lint

GO ?= go
TESTFOLDER := $(shell $(GO) list ./... | grep -E 'gin$$|binding$$|render$$' | grep -v examples)


test-docker:
	@docker-compose run --rm forge go test ./...

test:
	go test ./...

test-cover:
	./scripts/cover.sh

lint:
	@docker-compose run --rm lint run ./...
