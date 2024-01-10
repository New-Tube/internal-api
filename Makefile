.PHONY: all
all: build \
	test \
	run

.PHONY: build
build:
	cd internal-api-protos && \
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		*.proto
	go mod tidy
	go build .

.PHONY: run
run:
	./internal-api

.PHONY: test
test: build
	docker run --name test_db -p 5432:5432 -e POSTGRES_PASSWORD=password -e POSTGRES_USER=new-tube -d postgres && \
	sleep 5 && \
	cd endpoints && \
	go test || \
	docker stop test_db && \
	docker rm test_db

.PHONY: testv
testv: build
	docker run --name test_db -p 5432:5432 -e POSTGRES_PASSWORD=password -e POSTGRES_USER=new-tube -d postgres && \
	sleep 5 && \
	cd endpoints && \
	go test -v || \
	docker stop test_db && \
	docker rm test_db
