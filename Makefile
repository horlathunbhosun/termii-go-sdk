# Load environment variables from .env file
include .env
export

OUTPUT = main

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	@echo "Cleaning..."
	@go clean
	@rm ${OUTPUT}
	@echo "Cleaned!"

build-local:
	go build -o $(OUTPUT) ./cmd/termii.go

run: build-local
	@echo ">> Running application ..."
	@echo ">> TERMII_API_KEY: $(TERMII_API_KEY)"
	./$(OUTPUT)


