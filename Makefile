build:
	@go build -o bin/$(BINARY_NAME) -v
	@echo "Build complete"
run:
	@go run .
clean:
	@rm -f bin/$(BINARY_NAME)
	@echo "Clean complete"
test:
	@go test -v ./...
	@echo "Test complete"
	