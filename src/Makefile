export PROJECT = github.com/gufranmirza/provider-microservice
export BINARY_NAME = "provider-microservice"


# global command
.PHONY: all
all: dependencies api-docs build

.PHONY: dependencies
dependencies:
	go mod download

.PHONY: build
build:
	env GOOS=linux GOARCH=amd64 go build -o ../bin/${BINARY_NAME} .

.PHONY: start
start:
	../bin/${BINARY_NAME} serve

api-docs:
	swag init --output=./web/docs