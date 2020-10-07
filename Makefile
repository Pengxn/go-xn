# Make sure you install `make` and `git` tool

# OS=Windows_NT for git-bash on Windows OS
ifeq ($(OS), Windows_NT)
	BIN = go-xn.exe
else
	BIN = go-xn
endif

VERSION = $(shell git describe --tags `git rev-list --tags --max-count=1`)
COMMIT_ID = $(shell git rev-parse --short HEAD)
BUILD_TIME = $(shell date +'%Y-%m-%d %T')
LDFLAGS += -X "github.com/Pengxn/go-xn/src/cmd.Version=${VERSION}"
LDFLAGS += -X "github.com/Pengxn/go-xn/src/cmd.commitID=${COMMIT_ID}"
LDFLAGS += -X "github.com/Pengxn/go-xn/src/cmd.buildTime=${BUILD_TIME}"

all: build

build: clean
	@go build -o build/$(BIN) -tags=jsoniter -ldflags '$(LDFLAGS)'

clean:
	@if [ -f "build/$(BIN)" ]; then \
		rm -rf build/$(BIN); \
	fi;

cover:
	@go tool cover -html=coverage.txt -o coverage.html

test:
	@go test ./... -v -coverprofile=coverage.txt

web:
	@if [ -d "build/web" ]; then \
		rm -rf build/web; \
	fi;
    # Copy web folder and fyj.ini file
	@cp -r web/ build/web

.PHONY: build clean cover test web
