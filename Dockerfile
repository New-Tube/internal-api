FROM --platform=amd64 ubuntu:22.04 as builder

WORKDIR /build

RUN apt-get update
RUN apt-get install -y wget

RUN wget "https://go.dev/dl/go1.21.5.linux-amd64.tar.gz"

RUN tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin

RUN apt-get install -y protobuf-compiler make

RUN go version

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

COPY . . 

ENV GOOS=linux
ENV GOARCH=amd64
RUN PATH="$PATH:$(go env GOPATH)/bin" make build

FROM --platform=amd64 ubuntu:22.04

WORKDIR /app

COPY --from=builder /build/internal-api /app/internal-api
COPY --from=builder /build/.env /app/.env

EXPOSE 5050

ENTRYPOINT [ "./internal-api" ]

