PROGRAM = identity
LABEL   = $(shell git rev-parse --abbrev-ref HEAD)

.PHONY: build test

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

test:
	go test -cover github.com/camphotos/identity/pkg/models
	go test -cover github.com/camphotos/identity/pkg/repository



