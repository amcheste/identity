PROGRAM = identity
LABEL   = $(shell git rev-parse --abbrev-ref HEAD)

.PHONY: build test unit-test integration-test

all: build test

build:
	go build -o identity cmd/api/main.go

run:
	go run cmd/api/main.go

start:
	docker compose up --build -d

logs:
	docker compose logs

stop:
	 docker compose down --volumes

unit-test:
	go test -cover github.com/camphotos/identity/pkg/models
	go test -cover github.com/camphotos/identity/pkg/repository
	go test -cover github.com/camphotos/identity/pkg/handlers

integration-test:
	go test github.com/camphotos/identity/integration

test: unit-test integration-test



