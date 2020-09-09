.PHONY: help unit dist install

SRC_FILES := $(shell find . -iname "*.go" )

all: unit

help:
	@echo "Please use \`make <target>' where <target> is one of"
	@echo "  unit                    run unit tests"
	@echo "  dist                    build artifacts for the local and production architectures"
	@echo "  install                 install artifact on local path"
	@exit 1

unit:
	@echo "go test package"
	go mod tidy
	go test -cover -p 1 ./...

fmt:
	goimports -local github.com/beaconsoftwarellc/ -w .
