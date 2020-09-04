.PHONY: test build install

all: test

test:
	@echo "testing..."
	@go test -v ./...

build:
	@echo "building..."
	@go build -v ./...

install:
	@echo "installing..."
	@go install ./cmd/env-context