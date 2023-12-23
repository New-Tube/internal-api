.PHONY: all
all: build run

.PHONY: build
build:
	go mod tidy
	go build .

.PHONY: run
run: ./internal-api
