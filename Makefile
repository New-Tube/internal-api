.PHONY: all
all: build test run

.PHONY: build
build:
	go mod tidy
	go build .

.PHONY: run
run: ./internal-api

.PHONY: test
test: build
	echo "no tests"
