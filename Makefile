.PHONY: test build static shared clean cross

NAME := env-context

OUT := bin

#PLATFORMS := $(shell go tool dist list | xargs -n1 -i echo "{}.platform" | xargs)

all: test

clean:
	@echo "cleaning..."
	@rm -rf $(OUT)
	@go clean

test:
	@echo "testing..."
	@go test -v ./...

build:
	@echo "building $(NAME)..."
	@go build -o $(OUT)/$(NAME) ./cmd/$(NAME)

static: $(OUT)
	@echo "building static lib..."
	@go build -o $(OUT)/libc$(NAME).a -buildmode=c-archive libc.go

shared: $(OUT)
	@echo "building shared libs..."
	@go install -buildmode=shared -linkshared std
	@go build -buildmode=shared -linkshared ./...
	@go build -o $(OUT)/libc$(NAME).o -buildmode=c-shared libc.go

%.platform:
	@echo "building for $(*)..."
	@$(MAKE) -e GOOS=$(*D) GOARCH=$(*F) static shared build

$(OUT):
	@mkdir -p $(OUT)

cross: $(OUT) $(patsubst %, %.platform, $(PLATFORMS))

