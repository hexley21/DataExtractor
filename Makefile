MAKEFLAGS += --no-print-directory
CLI_NAME=datex

build:
ifeq ($(OS),Windows_NT) 
	@$(MAKE) build/batch
else
	@$(MAKE) build/bash
endif


# PLATFORM SPECIFIC
build/batch:
	set CGO_ENABLED=0&& go build -o ./bin/${CLI_NAME}.exe ./main.go \

build/bash:
	CGO_ENABLED=0 go build -o ./bin/${CLI_NAME} ./main.go; \
