APP_NAME := biller
CMD_DIR := ./cmd/bill
OUTPUT_DIR := ./bin
OUTPUT_FILE := $(OUTPUT_DIR)/$(APP_NAME)

test:
	@echo "Running tests..."
	go test -v ./...

test-coverage:
	@echo "Running tests..."
	go test -cover ./...

build:
	@echo "Building $(APP_NAME)..."
	go build -o $(OUTPUT_FILE) $(CMD_DIR)

run:
	@echo "Running $(APP_NAME)..."
	go run $(CMD_DIR)

clean:
	@echo "Cleaning up..."
	go clean
	rm -rf $(OUTPUT_DIR)/*

tidy:
	@echo "Tidying Go modules..."
	go mod tidy

test-and-build: test build

all: tidy test-and-build