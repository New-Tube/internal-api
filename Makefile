.PHONY: all
all: build \
	test \
	run

.PHONY: build
build:
	export PATH="$PATH:$(go env GOPATH)/bin"
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		protos/**/*.proto
	go mod tidy
	go build .

.PHONY: run
run:
	./internal-api

.PHONY: test
test: build
	echo "no tests"
