.PHONY: build test lint proto run clean help

# ============================================
# RMN Backend Monorepo - Makefile
# ============================================

# Build all services
build:
	@echo "Building all services..."
	@for dir in services/*/; do \
		if [ -f "$$dir/main.go" ]; then \
			echo "  Building $$dir..."; \
			cd $$dir && go build -o bin/app . && cd ../..; \
		elif [ -d "$$dir/cmd" ]; then \
			echo "  Building $$dir..."; \
			cd $$dir && go build -o bin/app ./cmd && cd ../..; \
		fi \
	done
	@echo "Build complete!"

# Test all services
test:
	@echo "Testing all services..."
	@for dir in services/*/; do \
		if [ -f "$$dir/go.mod" ]; then \
			echo "  Testing $$dir..."; \
			cd $$dir && go test ./... && cd ../..; \
		fi \
	done
	@echo "Tests complete!"

# Lint all services
lint:
	@echo "Linting..."
	golangci-lint run ./...

# Generate proto (Kafka events)
proto:
	@echo "Generating proto..."
	cd proto && buf generate
	@echo "Proto generation complete!"

# Run a specific service
# Usage: make run SERVICE=api-gateway
run:
	@if [ -z "$(SERVICE)" ]; then \
		echo "Usage: make run SERVICE=<service-name>"; \
		echo "Available services:"; \
		ls -d services/*/ | xargs -n1 basename; \
		exit 1; \
	fi
	@echo "Running $(SERVICE)..."
	@if [ -f "services/$(SERVICE)/main.go" ]; then \
		cd services/$(SERVICE) && go run .; \
	elif [ -d "services/$(SERVICE)/cmd" ]; then \
		cd services/$(SERVICE) && go run ./cmd; \
	else \
		echo "No main.go or cmd/ found for $(SERVICE)"; exit 1; \
	fi

# Start local infrastructure
infra-up:
	@echo "Starting local infrastructure..."
	docker-compose up -d

# Stop local infrastructure
infra-down:
	@echo "Stopping local infrastructure..."
	docker-compose down

# Go workspace sync
workspace-sync:
	@echo "Syncing Go workspace..."
	go work sync

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@for dir in services/*/; do \
		rm -rf $$dir/bin; \
	done
	@echo "Clean complete!"

# Deploy to staging
# Usage: make deploy-stg SERVICE=api-gateway
deploy-stg:
	@if [ -z "$(SERVICE)" ]; then \
		echo "Usage: make deploy-stg SERVICE=<service-name>"; \
		exit 1; \
	fi
	cd services/$(SERVICE)/chart && helm upgrade --install $(SERVICE) . -f values-stg.yaml

# Show help
help:
	@echo "RMN Backend Monorepo Commands:"
	@echo ""
	@echo "  make build          - Build all services"
	@echo "  make test           - Run tests for all services"
	@echo "  make lint           - Run linter"
	@echo "  make proto          - Generate Kafka event code from proto"
	@echo "  make run SERVICE=x  - Run a specific service"
	@echo "  make infra-up       - Start local infrastructure (Docker)"
	@echo "  make infra-down     - Stop local infrastructure"
	@echo "  make workspace-sync - Sync Go workspace"
	@echo "  make clean          - Clean build artifacts"
	@echo "  make deploy-stg     - Deploy service to staging"
	@echo ""
