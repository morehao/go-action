.PHONY: build-opensource build-enterprise build-all run-opensource run-enterprise clean test

BIN_DIR := bin

build-opensource:
	@echo "=== Building OpenSource Edition ==="
	go build -o $(BIN_DIR)/server-opensource .

build-enterprise:
	@echo "=== Building Enterprise Edition ==="
	go build -tags enterprise -o $(BIN_DIR)/server-enterprise .

build-all: build-opensource build-enterprise
	@echo "=== Both editions built successfully ==="

run-opensource: build-opensource
	@echo "=== Running OpenSource Edition ==="
	./$(BIN_DIR)/server-opensource

run-enterprise: build-enterprise
	@echo "=== Running Enterprise Edition ==="
	./$(BIN_DIR)/server-enterprise

test:
	@echo "=== Testing OpenSource Edition ==="
	go test ./...
	@echo "=== Testing Enterprise Edition ==="
	go test -tags enterprise ./...

clean:
	rm -rf $(BIN_DIR)

verify:
	@echo "=== Verify: No enterprise tag leaks ==="
	! grep -r "//go:build enterprise" --include="*.go" . || (echo "ERROR: enterprise tag found!"; exit 1)
