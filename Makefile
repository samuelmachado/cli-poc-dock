MODULE_NAME=$(shell grep ^module go.mod | cut -d " " -f2)
GIT_COMMIT_HASH=$(shell git rev-parse HEAD)
LD_FLAGS=-ldflags="-X $(MODULE_NAME)/internal/config.gitCommitHash=$(GIT_COMMIT_HASH)"

.PHONY: version
version:
	@go run cmd/cli-poc-dock/main.go version

.PHONY: build
build:
	@go mod tidy
	@go build -o ./bin/dock ./cmd/cli-poc-dock


.PHONY: test
test:
	@go test -race ./... -coverprofile=coverage.out
	@go tool cover -html=coverage.out