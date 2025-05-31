ENV ?= config/local.yml

export ENV

run:
	@echo "Running services with ENV=$(ENV)..."
	@CONFIG_PATH=$(ENV) go run ./services/auth/cmd/main.go &
	@CONFIG_PATH=$(ENV) go run ./services/users/cmd/main.go &
	@CONFIG_PATH=$(ENV) go run ./services/api-gateway/cmd/main.go

.PHONY: run