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
GOLANGCI_LINT := go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.56.2
GOTESTSUM := go run gotest.tools/gotestsum@v1.11.0
TPARSE := go run github.com/mfridman/tparse@v0.13.2

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


