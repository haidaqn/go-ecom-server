APP_NAME=server
APP_PATH=cmd/$(APP_NAME)/main.go
APP_WIRE=./internal/wire

run:
	go run $(APP_PATH)

run_wire:
	wire $(APP_WIRE)

# .PHONY: run  

