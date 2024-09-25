appName:=test-project
osArch:=$(shell uname -m)
envfile:=.env

include $(envfile)

ifeq ($(osArch),arm64)
	dynamicFlag:=--tags dynamic
endif

export $(shell sed 's/=.*//' $(envfile))

lint:
	@golangci-lint run

format:
	@go fmt -n ./...

run:
	@go run $(dynamicFlag) main.go $(command)

debug:
	@dlv --listen=127.0.0.1:8181 --headless --api-version=2 --accept-multiclient --check-go-version=false debug ./src --log -- ${command}
