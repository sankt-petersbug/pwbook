default: build

.PHONY: clean
clean: ## remove build artifact
		rm -rf ./build

.PHONY: build
build: ## build executable
		@go build -o ./build/pwbook ./main.go

.PHONY: test
test: ## run tests
		@go test -v -cover ./pwbook/...

.PHONY: test-e2e
test-e2e: ## run e2e tests
		@go get ./...
		@go install .
		@go test -v ./e2e/...
