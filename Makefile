.PHONY: fmt build deb

build: fmt
	GO111MODULE=on CGO_ENABLED=0 go build -o flipd

fmt:
	gofmt -l -w ./

deb: build
	mkdir -p bin
	cp flipd bin/
	dpkg-deb --root-owner-group --build .
	dpkg-name --overwrite ..deb
