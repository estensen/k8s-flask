.PHONY: help

help: #Self-documents targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'


build: ## Build container
	DOCKER_BUILDKIT=1 docker build -t k8s-flask .

run: ## Run container
	docker run -p 5000:5000 k8s-flask

up: build run ## Build and run container
