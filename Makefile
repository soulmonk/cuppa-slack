PROJECT_NAME := $(shell basename "$(PWD)")

TAG ?= test

.PHONY: build
build: ## Build the application and produce a binary
	go build .

.PHONY: test
test: ## Test our application
	go test ./...

.PHONY: docker-build
docker-build: ## Build a docker image
	docker build -t="${PROJECT_NAME}:${TAG}" --platform linux/arm64 .

.DEFAULT_GOAL := help
.PHONY: help
all: help
help: ## Prints this help
	@echo "Choose a command run in \"$(PROJECT_NAME)\":"
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'