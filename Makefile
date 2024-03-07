# Database
MYSQL_USER ?= user
MYSQL_PASSWORD ?= password
MYSQL_ADDRESS ?= 127.0.0.1:3306
MYSQL_DATABASE ?= article

# Default Shell
SHELL := /bin/bash

# Type of OS: Linux or Darwin.
OSTYPE := $(shell uname -s | tr A-Z a-z)
ARCH := $(shell uname -m)



# --- Tooling & Variables ----------------------------------------------------------------
GOLANGCI_LINT := go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.56.2
GOTESTSUM := go run gotest.tools/gotestsum@v1.11.0
TPARSE := go run github.com/mfridman/tparse@v0.13.2
# ~~~ Code Actions ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

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


