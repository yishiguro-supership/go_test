include ./dynamodb/dynamodb.mk

.DEFAULT_GOAL := help

start: ## local service start
	$(MAKE) dynamodb-start

restart: ## local service restart
	$(MAKE) dynamodb-restart

stop: ## local service stop
	$(MAKE) dynamodb-stop

help: ## Help
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort |\
		awk 'BEGIN {FS = ":[^:]*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## dev build
	@go build

clean: ## dev clean
	@go clean

run: ## dev run
	@go run myproject

wire: ## wire
	wire ./injector/
