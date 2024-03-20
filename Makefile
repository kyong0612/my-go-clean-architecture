include .env

# Default Shell
SHELL := /bin/bash

# Type of OS: Linux or Darwin.
OSTYPE := $(shell uname -s | tr A-Z a-z)
ARCH := $(shell uname -m)

# --- Tooling & Variables ----------------------------------------------------------------
AIR := go run github.com/cosmtrek/air@v1.51.0
GOLANGCI_LINT := go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.56.2
GOTESTSUM := go run gotest.tools/gotestsum@v1.11.0
TPARSE := go run github.com/mfridman/tparse@v0.13.2
ATLAS := docker run --rm --net=host --mount type=bind,source="$(PWD)"/migrations,target=/migrations arigaio/atlas:0.19.2

# ~~~ Development Environment ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
.PHONY: dev-air
dev-air: 
	$(AIR) -c .air.app.toml

.PHONY: compose-up
compose-up:
	@docker compose up -d --build

.PHONY: compose-down
compose-down:
	@docker compose down

.PHONY: compose-teardown
compose-teardown:
	@docker compose down --remove-orphans -v

.PHONY: compose-reset
compose-reset: compose-teardown compose-up

.PHONY: gen-sqlc
generate-sqlc:
	@go run github.com/sqlc-dev/sqlc/cmd/sqlc@v1.25.0 generate

# ~~~ Code Actions ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
.PHONY: lint
lint:
	@echo "Applying linter"
	$(GOLANGCI_LINT) run --fix

TESTS_ARGS := --format testname --jsonfile gotestsum.json.out
TESTS_ARGS += --max-fails 2
TESTS_ARGS += -- ./...
TESTS_ARGS += -test.parallel 2
TESTS_ARGS += -test.count    1
TESTS_ARGS += -test.failfast
TESTS_ARGS += -test.coverprofile   coverage.out
TESTS_ARGS += -test.timeout        5s
TESTS_ARGS += -race

tests:
	@$(GOTESTSUM) $(TESTS_ARGS) -short

tests-complete: ## Run Tests & parse details
	make tests
	@cat gotestsum.json.out | $(TPARSE) -all -notests

# ~~~ Docker ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~



# # ~~~ Database Migrations ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 
BASE_DSN := postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)
DB_DSN := $(BASE_DSN)/$(DB_DATABASE)
DEV_DB_DSN := $(BASE_DSN)/migrate-dev-database

.PHONY: migrate-new
migrate-new:
	@read -p "Please provide name for the migration: " Name; \
	$(ATLAS) migrate new $${Name}

.PHONY: migrate-hash
migrate-hash:
	$(ATLAS) migrate hash

.PHONY: migrate-apply
migrate-apply:
	$(ATLAS) migrate apply \
		-u $(DB_DSN)?sslmode=disable \
		--dir "file://migrations" 
		# --dev-url $(DEV_DB_DSN)?sslmode=disable

.PHONY: schema-inspect
schema-inspect:
	$(ATLAS) schema inspect \
		--url $(DB_DSN)?sslmode=disable \
		--format '{{ mermaid . }}' > doc/db-schema.mmd


# migrate-up:
# 	migrate  -database $(DB_DSN) -path=misc/migrations up ${NN}

# .PHONY: migrate-down
# migrate-down: $(MIGRATE) ## Apply all (or N down) migrations.
# 	@read -p "How many migration you wants to perform (default value: [all]): " N; \
# 	migrate  -database $(DB_DSN) -path=misc/migrations down ${NN}

# .PHONY: migrate-drop
# migrate-drop: $(MIGRATE) ## Drop everything inside the database.
# 	migrate  -database $(DB_DSN) -path=misc/migrations drop

# .PHONY: migrate-create
# migrate-create: 
# 	@read -p "Please provide name for the migration: " Name; \
# 	migrate create -ext sql -dir misc/migrations $${Name}

