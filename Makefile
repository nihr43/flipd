.PHONY: fmt build

build:
	GO111MODULE=on CGO_ENABLED=0 go build -o flipd

fmt:
	gofmt -l -w ./
