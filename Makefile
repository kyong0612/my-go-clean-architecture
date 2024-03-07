# Database
DB_USER ?= postgres
DB_PASSWORD ?= password
DB_ADDRESS ?= 127.0.0.1:5432
DB_DATABASE ?= article

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

# ~~~ Code Actions ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
.PHONY: lint
lint:
	@echo "Applying linter"
	$(GOLANGCI_LINT) -c .golangci.yml ./...



.PHONY: go-generate
go-generate:  ## Runs go generte ./...
	go generate ./...


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
	@cat gotestsum.json.out | $(TPARSE) -all -notests

# ~~~ Docker ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~



# # ~~~ Database Migrations ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 

# MYSQL_DSN := "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_ADDRESS))/$(MYSQL_DATABASE)"

# migrate-up: $(MIGRATE) ## Apply all (or N up) migrations.
# 	@read -p "How many migration you wants to perform (default value: [all]): " N; \
# 	migrate  -database $(MYSQL_DSN) -path=misc/migrations up ${NN}

# .PHONY: migrate-down
# migrate-down: $(MIGRATE) ## Apply all (or N down) migrations.
# 	@read -p "How many migration you wants to perform (default value: [all]): " N; \
# 	migrate  -database $(MYSQL_DSN) -path=misc/migrations down ${NN}

# .PHONY: migrate-drop
# migrate-drop: $(MIGRATE) ## Drop everything inside the database.
# 	migrate  -database $(MYSQL_DSN) -path=misc/migrations drop

# .PHONY: migrate-create
# migrate-create: $(MIGRATE) ## Create a set of up/down migrations with a specified name.
# 	@read -p "Please provide name for the migration: " Name; \
# 	migrate create -ext sql -dir misc/migrations $${Name}

