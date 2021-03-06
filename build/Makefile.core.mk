# PRODUCTION CONFIGURATION
SHELL := /bin/bash
PROJECT_NAME := ArteCol

# This will output the help for each task
.PHONY: help
help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

refresh-dependencies: ## Update dependencies of the project
	go get $(HOST_PROJECT_ROOT)

update-dependencies:
	go get -u $(HOST_PROJECT_ROOT)

run: ## Run project
	go run ./cmd/$(PROJECT_NAME)
