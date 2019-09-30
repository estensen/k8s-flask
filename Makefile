.PHONY: help

help: #Self-documents targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'


build: ## Build containers
	DOCKER_BUILDKIT=1 docker build -t estensen/books books/.
	DOCKER_BUILDKIT=1 docker build -t estensen/club club/.

run: ## Run containers
	docker run -p 5000:5000 estensen/books
	docker run -p 8080:8080 estensen/club

stop: ## Stop containers
	docker stop estensen/books
	docker stop estensen/club

push: build upload ## Push to Docker hub

upload: ## Upload to Docker Hub
	docker push estensen/books
	docker push estensen/club

compose: ## Run all containers
	docker-compose up

up: build compose ## Build and run containers
