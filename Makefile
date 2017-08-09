.PHONY build test

test:
	go test

build: test
	go build
