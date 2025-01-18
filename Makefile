.PHONY: all build run test clean setup

all: build

setup: 
	sh ./scripts/setup.sh

build:
	go build -o bin/server ./cmd/server

run: build
	./bin/server

test:
	go test ./test/...

clean:
	rm -rf bin/
