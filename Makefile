PROGRAM = identity
LABEL   = $(shell git rev-parse --abbrev-ref HEAD)

.PHONY: build test unit integration acceptance

all: build test

build:
	CGO_ENABLED=0 GOOS=linux go build -o main cmd/api/main.go

run:
	go run cmd/api/main.go

start:
	docker compose up --build -d

logs:
	docker compose logs

stop:
	 docker compose down --volumes

unit:
	go test -cover github.com/camphotos/identity/pkg/models
	go test -cover github.com/camphotos/identity/pkg/repository
	go test -cover github.com/camphotos/identity/pkg/handlers

integration:
	go test github.com/camphotos/identity/integration

acceptance:
	make start
	sleep 1
	go test github.com/camphotos/identity/acceptance
	make stop

test: unit integration acceptance



