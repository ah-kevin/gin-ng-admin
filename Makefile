NOW = $(shell date -u '+%Y%m%d%I%M%S')
RELEASE_VERSION = v0.0.0
APP 			= ng-admin
SERVER_BIN  	= ./cmd/${APP}/${APP}
RELEASE_ROOT 	= release
RELEASE_SERVER 	= release/${APP}
GIT_COUNT 		= $(shell git rev-list --all --count)
GIT_HASH        = $(shell git rev-parse --short HEAD)
RELEASE_TAG     = $(RELEASE_VERSION).$(GIT_COUNT).$(GIT_HASH)

.PHONY: init
# init env
init:
	go get -u github.com/google/wire/cmd/wire
	go get -u github.com/swaggo/swag/cmd/swag

.PHONY: start
# start
start:
	air -c .air.toml

.PHONY: clean
# clean
clean:
	rm -rf $(SERVER_BIN)

.PHONY: generate
# generate client code
generate:
	go generate ./...

.PHONY: swagger
# generate swagger file
swagger:
	@swag init --parseDependency --generalInfo ./cmd/${APP}/main.go --output ./internal/admin/swagger

.PHONY: build
# build
build:
	go build -ldflags "-X main.Version=$(VERSION)" -o  $(SERVER_BIN) ./cmd/${APP}

.PHONY: test
# test
test:
	go test -v ./... -cover

.PHONY: all
# generate all
all:
	make generate
	make swagger
	make build

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
