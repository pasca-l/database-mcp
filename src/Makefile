MAKEFILE_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

.PHONY: format
format:
	go fmt ./...

.PHONY: test
test:
	go test -v -race -cover -covermode=atomic ./...
