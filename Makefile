.PHONY: build
build:
	@echo "Building the application"
	@go build -o bin/asi-user-management cmd/asi-user-management/main.go

.PHONY: run
run:
	@echo "Running the application"
	@go run cmd/asi-user-management/main.go
