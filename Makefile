MAKEFILE_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
APP_SERVICE := app

.PHONY: build
build:
	cd $(MAKEFILE_DIR)/src && go build -o $(MAKEFILE_DIR)/database-mcp

.PHONY: build-docker
build-docker:
	docker compose run --rm -v $(MAKEFILE_DIR):/out $(APP_SERVICE) sh -c "go build -o /out/database-mcp"

.PHONY: format
format:
	docker compose run --rm $(APP_SERVICE) make format

.PHONY: lint
lint:
	docker run -t --rm -v $(MAKEFILE_DIR):/app -w /app/src golangci/golangci-lint:v2.9.0 golangci-lint run

.PHONY: test
test:
	docker compose run --rm $(APP_SERVICE) make test
