# Make sure you install `make` and `git` tool

BIN        ?= go-xn$(shell go env GOEXE)
VERSION     = $(shell git describe --tags `git rev-list --tags --max-count=1`)
COMMIT_ID   = $(shell git rev-parse --short HEAD)
BUILD_TIME  = $(shell date +'%Y-%m-%d %T')
LDFLAGS    += -X "github.com/Pengxn/go-xn/src/cmd.version=${VERSION}"
LDFLAGS    += -X "github.com/Pengxn/go-xn/src/cmd.commitID=${COMMIT_ID}"
LDFLAGS    += -X "github.com/Pengxn/go-xn/src/cmd.buildTime=${BUILD_TIME}"

all: build

build: clean
	@go build -o build/$(BIN) -v -ldflags '-w -s $(LDFLAGS)'

clean:
	@if [ -f "build/$(BIN)" ]; then \
		rm -rf build/$(BIN); \
	fi;

generate:
	@rm -rf src/rpc/proto/*.pb.go
	@go generate ./...

cover:
	@go tool cover -html=coverage.txt -o coverage.html

test:
	@go test ./... -v -coverprofile=coverage.txt

.PHONY: build clean generate cover test
