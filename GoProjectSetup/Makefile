# Build configuration
BINARY_DIR=build
VERSION=1.0.0
# Shrinks the Go binary size for production by stripping debug symbols
LDFLAGS=-ldflags="-s -w -X main.Version=$(VERSION)"

.PHONY: all build build-api build-ws test clean run-api run-ws

all: build

## build: Compiles both binaries with production optimization flags
build: build-api build-ws

build-api:
	@echo "[BUILD] Building Production API Microservice..."
	@mkdir -p $(BINARY_DIR)
	go build $(LDFLAGS) -o $(BINARY_DIR)/api ./cmd/api

build-ws:
	@echo "[BUILD] Building Production WebSocket Microservice..."
	@mkdir -p $(BINARY_DIR)
	go build $(LDFLAGS) -o $(BINARY_DIR)/ws ./cmd/ws

## test: Runs tests with data-race detection
test:
	@echo "[TEST] Running unit tests with race detection..."
	go test -race -v ./...

## run-api: Local test execution for API (POSIX env variables syntax)
run-api: build-api
	@echo "[RUN] Starting API..."
	@API_PORT=8080 ./$(BINARY_DIR)/api

## run-ws: Local test execution for WebSocket (POSIX env variables syntax)
run-ws: build-ws
	@echo "[RUN] Starting WebSocket..."
	@WS_PORT=8081 ./$(BINARY_DIR)/ws

## clean: Drops local build artifacts
clean:
	@echo "[CLEAN] Cleaning build artifacts..."
	@rm -rf $(BINARY_DIR)
