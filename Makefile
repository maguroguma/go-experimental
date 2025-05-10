.PHONY: help
.PHONY: generate test
.PHONY: up down

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

generate: ## go generate の実行
	go generate ./...

test: ## go test の実行
	go test -cover ./...

up: ## Docker Compose の起動
	docker compose -f compose.yaml up -d

down: ## Docker Compose の停止
	docker compose -f compose.yaml down

