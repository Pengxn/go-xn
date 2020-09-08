# Make sure you install `make` and `git` tool

# OS=Windows_NT for git-bash on Windows OS
ifeq ($(OS), Windows_NT)
	BIN = go-xn.exe
else
	BIN = go-xn
endif

VERSION = $(shell git describe --tags `git rev-list --tags --max-count=1`)
COMMIT_ID = $(shell git rev-parse --short HEAD)
BUILD_DATE = $(shell date +'%Y-%m-%d')
BUILD_TIME = $(shell date +'%T')

all: build

build: clean
	@go build -o build/$(BIN) -tags=jsoniter -ldflags \
	"-X github.com/Pengxn/go-xn/src/cmd.Version=${VERSION} \
	-X github.com/Pengxn/go-xn/src/cmd.commitID=${COMMIT_ID} \
	-X github.com/Pengxn/go-xn/src/cmd.buildDate=${BUILD_DATE} \
	-X github.com/Pengxn/go-xn/src/cmd.buildTime=${BUILD_TIME}"

clean:
	@if [ -f "build/$(BIN)" ]; then \
		rm -rf build/$(BIN); \
	fi;

cover:
	@go tool cover -html=coverage.txt -o coverage.html

test:
	@go test ./... -v -coverprofile=coverage.txt

web:
	# Delete 'build/web'
	@if [ -d "build/web" ]; then \
		rm -rf build/web; \
	fi;
    # Copy web folder and fyj.ini file
	@cp -r web/ build/web

.PHONY: build clean cover test web
