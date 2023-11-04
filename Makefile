# build related variables
PREFIX?=$(shell pwd)
BUILD_DIR?=${PREFIX}/builds

# database migration tasks
MIGRATION_DIR?=$(shell pwd)/migrations
TABLE_NAME?=untitled

.PHONY: migrations-setup migrations-generate migrations-up migrations-down

migrations-setup:
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go get github.com/pressly/goose
	go mod tidy

migrations-generate: migrations-setup
	goose -dir ${MIGRATION_DIR} postgres ${POSTGRES_URL} create ${TABLE_NAME} sql

migrations-up: migrations-setup
	goose -dir ${MIGRATION_DIR} postgres ${POSTGRES_URL} up

migrations-down: migrations-setup
	goose -dir ${MIGRATION_DIR} postgres ${POSTGRES_URL} down

# mnstr_advntr tasks
MNSTR_ADVNTR_DIR=${BUILD_DIR}/client/mnstr_advntr

build-mnstr-advntr:
	mkdir -p ${MNSTR_ADVNTR_DIR}
	go build -o ${MNSTR_ADVNTR_DIR}/mnstr_advntr ./cmd/client/mnstr_advntr

run-mnstr_advntr: build-mnstr-advntr
	${MNSTR_ADVNTR_DIR}/mnstr_advntr
