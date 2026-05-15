include .env

GOOSE_DBSTRING=$(STR_MYSQL)
GOOSE_MIGRATION_DIR=sql/schema
GOOSE_DRIVER=mysql

APP_NAME=server
APP_PATH=cmd/$(APP_NAME)/main.go
APP_WIRE=./internal/wire

run:
	go run $(APP_PATH)

run_wire:
	wire $(APP_WIRE)

upse:
	set GOOSE_DRIVER=$(GOOSE_DRIVER)&& \
	set GOOSE_DBSTRING=$(GOOSE_DBSTRING)&& \
	goose -dir=$(GOOSE_MIGRATION_DIR) up

downse:
	set GOOSE_DRIVER=$(GOOSE_DRIVER)&& \
	set GOOSE_DBSTRING=$(GOOSE_DBSTRING)&& \
	goose -dir=$(GOOSE_MIGRATION_DIR) down

resetse:
	set GOOSE_DRIVER=$(GOOSE_DRIVER)&& \
	set GOOSE_DBSTRING=$(GOOSE_DBSTRING)&& \
	goose -dir=$(GOOSE_MIGRATION_DIR) reset

.PHONY: run run_wire upse downse resetse

.PHONY: air