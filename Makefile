include .env

GOOSE_DBSTRING=$(STR_MYSQL)
GOOSE_MIGRATION_DIR=sql/schema
GOOSE_DRIVER=mysql

APP_NAME=server
APP_PATH=cmd/$(APP_NAME)/main.go
APP_WIRE=./internal/wire

# test

all:
	@echo "The message is: $(MESSAGE)"
	@echo "The count is: $(COUNT)"

print_vars:
	@echo "MESSAGE (from make): $(MESSAGE)"
	@echo "COUNT (from make): $(COUNT)"


docker_build:
	docker-compose up -d --docker_build
	docker-compose ps

docker_down:
	docker-compose -f environment/docker-compose-dev.yml down

docker_up:
	docker-compose -f environment/docker-compose-dev.yml up

docker_stop:
	docker-compose stop

dev:
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

sqlgen:
	sqlc generate

swag:
	swag init -g ./cmd/server/main.go -o ./cmd/swag/docs

.PHONY: dev downse upse resetse docker_build docker_stop docker_up
.PHONY: air