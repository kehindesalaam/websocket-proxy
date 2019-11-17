.PHONY: all clean build test test-unit test-integration pipeline-set

NAME=message-socket-service
REPO=github.com/terawork-com/$(NAME)
VERSION ?= dev
BUILD_DIR ?= bin
BINARY=$(BUILD_DIR)/$(NAME)
BUILD_FLAGS=-ldflags="-s -w -X main.Version=$(VERSION)"

NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

DOCKER_COMPOSE_EXISTS := $(shell command -v docker-compose 2> /dev/null)

#-----------------------------------------------------------------------------------------------------------------------
# Functions
#-----------------------------------------------------------------------------------------------------------------------
#---- normal enviroment ----
define godog
	@godog $(1)
endef
define go
	@go $(1)
endef
#---- docker enviroment ----
ifdef DOCKER_COMPOSE_EXISTS
define godog
	@docker-compose exec api sh -c "godog ${1}"
endef
define go
	@docker-compose exec api sh -c 'go ${1}'
endef
endif

# Rules
#-----------------------------------------------------------------------------------------------------------------------
all: build

clean:
	@echo "$(OK_COLOR)==> Cleaning Bin $(NO_COLOR)"
	@rm -Rf bin/

build:
	@echo "$(OK_COLOR)==> Building Binary$(NO_COLOR)"
	@go build -v $(BUILD_FLAGS) -o "$(BINARY)" $(BINARY_SRC)

test: test-unit test-integration

test-unit:
	@echo "$(OK_COLOR)==> Running tests$(NO_COLOR)\n"
	@go test -race -tags=unit ./...

test-integration:
	${call godog, --tags="~@notImplemented" --strict}

lint:
	@echo "$(OK_COLOR)==> Checking code standards...$(NO_COLOR)"
	@golangci-lint -c .golangci-pkg.yml run

fix-lint:
	@echo "$(OK_COLOR)==> Fix gofmt and goimports...$(NO_COLOR)"
	@gofmt -s -w ./..
	@goimports -l -w ./..

run-compile-daemon:
	@echo ">> running app with CompileDaemon"
	@CompileDaemon -exclude-dir=vendor -build='make build' -command='${BINARY}' -graceful-kill

setup: build up

ssh:
	@docker-compose exec mss /bin/bash

up:
	@echo "$(OK_COLOR)==> Starting containers...$(NO_COLOR)"
	@docker-compose up -d