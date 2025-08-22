APP_NAME ?= go-starter
PORT ?= 8080


.PHONY: run build test tidy lint docker-up docker-down


run:
go run ./cmd/api


build:
go build -o bin/$(APP_NAME) ./cmd/api


test:
go test ./...


tidy:
go mod tidy


lint:
@which golangci-lint >/dev/null 2>&1 || (echo "Install golangci-lint first: https://golangci-lint.run/" && exit 1)
golangci-lint run ./...


docker-up:
docker compose up --build -d


docker-down:
docker compose down -v