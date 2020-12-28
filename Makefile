.PHONY: run-dev test

run-dev:
	go run cmd/switchboard/main.go dev-config.yml

build:
	go build -o switchboard cmd/switchboard/main.go

test:
	go test ./...
