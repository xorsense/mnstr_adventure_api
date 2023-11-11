# Environment variables
-include .env

# build related variables

PREFIX?=.
BUILD_DIR?=$(PREFIX)/builds

# database migration tasks
MIGRATION_DIR?=$(PREFIX)/migrations
TABLE_NAME?=untitled

.PHONY: migrations-setup migrations-generate migrations-up migrations-down

migrations-setup:
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go get github.com/pressly/goose
	go mod tidy

migrations-generate: migrations-setup
	goose -dir $(MIGRATION_DIR) postgres $(POSTGRES_URL) create $(TABLE_NAME) sql

migrations-up: migrations-setup
	goose -dir $(MIGRATION_DIR) postgres $(POSTGRES_URL) up

migrations-down: migrations-setup
	goose -dir $(MIGRATION_DIR) postgres $(POSTGRES_URL) down

# mnstr_advntr tasks
MNSTR_ADVNTR_DIR=$(BUILD_DIR)/server/mnstr_advntr

build:
ifeq ($(OS), Windows_NT)
	if not exist builds\server\mnstr_advntr mkdir builds\server\mnstr_advntr
else
	mkdir -p $(MNSTR_ADVNTR_DIR)
endif
	go build -o $(MNSTR_ADVNTR_DIR)/mnstr_advntr .

run: build
	$(MNSTR_ADVNTR_DIR)/mnstr_advntr
