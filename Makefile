# Conditionally load .env if it exists
ifneq (,$(wildcard .env))
    include .env
    export
endif

# Default values (override with CLI arguments)
APP_NAME ?= financing-system
APP_VERSION ?= 1.0.0
HTTP_PORT ?= 8080

# Makefile variables
APP_SRC_DIR = cmd/app/main.go
BUILD_DIR = bin
LOG_FILE = logs/output.log
SWAGGER_OUTPUT ?= docs
SWAGGER_MAIN ?= internal/controller/http/v1/router.go

# Define colors
RESET=\033[0m
GREEN=\033[32m
YELLOW=\033[33m
BLUE=\033[34m
RED=\033[31m
CYAN=\033[36m
MAGENTA=\033[35m

# Function to log messages with timestamp and color
define log
	@echo "$$(date '+[%Y-%m-%d %H:%M:%S]') $(1)$(2)$(RESET)" | tee -a $(LOG_FILE)
endef

# Function to set build target
define buildTarget
	$(call log,$(YELLOW),"🔨 Building $(1)...")
	@go build -o $(BUILD_DIR)/$(1) $(2)
	$(call log,$(GREEN),"✅ Build complete: $(BUILD_DIR)/$(1)")
endef

# Makefile targets
.PHONY: build run test fmt clean deps lint swag help

# Ensure logs directory exists
build-dir:
	@mkdir -p $(BUILD_DIR)

# Ensure logs directory exists
logs-dir:
	@mkdir -p logs

# Ensure Swagger docs directory exists
swagger-dir:
	@mkdir -p $(SWAGGER_OUTPUT)

# Initialize Swagger documentation
swag: swagger-dir
	$(call log,$(YELLOW),"📄 Generating Swagger docs for $(SWAGGER_MAIN)...")
	@swag init -g $(SWAGGER_MAIN) -o $(SWAGGER_OUTPUT) | tee -a $(LOG_FILE)
	$(call log,$(GREEN),"✅ Swagger documentation generated in $(SWAGGER_OUTPUT).")

# Build the application
build: clean logs-dir build-dir
	$(call buildTarget,$(APP_NAME),$(APP_SRC_DIR))

# Run the application
run: build
	$(call log,$(CYAN),"🚀 Running $(APP_NAME) on port $(HTTP_PORT)...")
	@./$(BUILD_DIR)/$(APP_NAME) | tee -a $(LOG_FILE)

# Run tests
test:
	$(call log,$(BLUE),"🧪 Running tests...")
	@go test ./... | tee -a $(LOG_FILE)

# Format code
fmt:
	$(call log,$(BLUE),"🎨 Formatting code...")
	@go fmt ./...

# Clean generated files
clean:
	$(call log,$(RED),"🧹 Cleaning up...")
	@rm -rf $(BUILD_DIR)

# Install dependencies
deps:
	$(call log,$(BLUE),"📦 Downloading dependencies...")
	@go mod tidy | tee -a $(LOG_FILE)

# Lint the code
lint:
	$(call log,$(BLUE),"🔍 Linting code...")
	@golangci-lint run | tee -a $(LOG_FILE)

# Display available Makefile commands with descriptions
help:
	@echo ""
	@echo "$(BLUE)📌 Available Makefile Commands:$(RESET)"
	@echo "-------------------------------------------------"
	@echo "  $(GREEN)1. build           $(YELLOW)- Compile the Go application$(RESET)"
	@echo "  $(GREEN)2. run             $(YELLOW)- Build and run the application$(RESET)"
	@echo "  $(GREEN)3. test            $(YELLOW)- Run all tests in the project$(RESET)"
	@echo "  $(GREEN)4. fmt             $(YELLOW)- Format the Go code$(RESET)"
	@echo "  $(GREEN)5. clean           $(YELLOW)- Remove compiled files and clean workspace$(RESET)"
	@echo "  $(GREEN)6. deps            $(YELLOW)- Install and tidy Go dependencies$(RESET)"
	@echo "  $(GREEN)7. lint            $(YELLOW)- Run linting (requires golangci-lint)$(RESET)"
	@echo "  $(GREEN)8. swag           $(YELLOW)- Run swagger init (requires swaggo)$(RESET)"
	@echo "  $(GREEN)9. help           $(YELLOW)- Show this help message$(RESET)"
	@echo "-------------------------------------------------"
	@echo "📌 Usage: $(CYAN)make [command]$(RESET)"
	@echo ""

# Prevent Make from treating the extra argument as a target
%:
	@: