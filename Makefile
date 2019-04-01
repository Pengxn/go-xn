# Get OS name: Linux, MINGW64_NT-10.0 or others
OS_NAME = $(shell uname -s)

ifeq ($(OS_NAME), Linux)
	BIN = go-xn
else
	BIN = go-xn.exe
endif

all: build

build: clean
	@go build -o build/$(BIN) -tags=jsoniter

clean:
	@if [ -f "build/$(BIN)" ]; then rm -rf build/$(BIN); fi;

test:
	@cd build && ./$(BIN)

web:
	# Delete 'build/web'
	@if [ -d "build/web" ]; then rm -rf build/web; fi;
    # Copy web folder and fyj.ini file
	@cp -r web/ build/web

.PHONY: build test clean web
