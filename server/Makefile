
# Test target
test:
	go test -v ./...

# Lint target
lint:
	golangci-lint run

# Build target
build:
	go build -o $(BIN_DIR)/$(BINARY) $(SRC_DIR)

swag-init:
	swag init