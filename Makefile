APP_NAME:=mocro-be
APP_PATH:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
SCRIPT_PATH:=$(APP_PATH)/scripts
COMPILE_OUT:=$(APP_PATH)/bin/$(APP_NAME)

run:export EGO_DEBUG=true
run:
	@go run main.go --config=./config/local.toml