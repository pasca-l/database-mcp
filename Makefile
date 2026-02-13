MAKEFILE_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
APP_SERVICE := app

.PHONY: start
start:
	docker compose up --build -d
	docker compose exec $(APP_SERVICE) sh -c "go run main.go"

.PHONY: format
format:
	docker compose run --rm $(APP_SERVICE) make format

.PHONY: lint
lint:
	docker run -t --rm -v $(MAKEFILE_DIR):/app -w /app/src golangci/golangci-lint:v2.9.0 golangci-lint run

.PHONY: test
test:
	docker compose run --rm $(APP_SERVICE) make test
