BINARY=./tmp/main

test:
	@echo "Running tests..."
	@go test -v ./...

run: build
	@echo "Running the application..."
	@air

clean:
	@echo "Cleaning up..."
	@rm -f $(BINARY)

create-tmp:
	@mkdir -p ./tmp

migration:
	@migrate create -ext sql -dir internal/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run internal/migrate/main.go up

migrate-down:
	@go run internal/migrate/main.go down
force-version:
	@read -p "Enter the version to force: " version; \
	go run internal/migrate/main.go force $$version

fmt:
	@echo "Formatting code..."
	@go fmt ./...

install-deps:
	@echo "Installing dependencies..."
	@go mod download

all: install-deps fmt test build

.PHONY: test run clean create-tmp migration migrate-up migrate-down force-version build fmt install-deps all
