.PHONY: clean build test deps

build:
	go build

test:
	go test -v ./test/automated/...

clean:
	go clean -i github.com/thehungry-dev/rag...

deps:
	go get && go mod tidy
