.PHONY: build
build:
	go build ./...

test: build
	go test ./... -bench=.

# Test with color output.
# go get -u github.com/rakyll/gotest
testc: build
	gotest -v ./... -bench=.