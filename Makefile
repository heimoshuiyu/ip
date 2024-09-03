.PHONY: linux
linux:
	CGO_ENABLED=0 go build -v -ldflags '-extldflags=-static'
