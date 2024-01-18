# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

.DEFAULT_GOAL := help

launch: ## Build and launch the application
	@echo "==> Launching application"
	docker-compose up --build 
.PHONY: launch

test: ## Test the code 
	@echo "==> Running tests"
	go test -v ./...
.PHONY: test

gen: generate fmt tidy ## Go generate and format and clean all files 
.PHONY: gen

generate: ## Go generate all files 
	@echo "==> Go generate"
	go generate ./...
.PHONY: generate

fmt: ## Go format all files 
	@echo "==> Running gofumpt"
	go run mvdan.cc/gofumpt@latest -l -w .
.PHONY: fmt

tidy: ## Cleans the Go module.
	@echo "==> Tidying module"
	go mod tidy
.PHONY: tidy

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
.PHONY: help
