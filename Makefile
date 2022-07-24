.PHONY: version

version:
	@go run cmd/cli-poc-dock/main.go version
build:
	@go mod tidy
	@go build -o ./bin/dock ./cmd/cli-poc-dock