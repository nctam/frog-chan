# variables
BINARY_FILE := bot
MAIN_PATH := ./cmd/frog-chan/main.go

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## run: unit test
.PHONY: test
test:
	go test ./...

## build: build to binary base on platform your machine
.PHONY: build
build:
	go build -o ${BINARY_FILE} ${MAIN_PATH}

## clean: clean build file
clean:
	rm -rf ./${BINARY_FILE}

## push: push to remote git
.PHONY: push
push: test
	git push
