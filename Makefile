default: build

.PHONY: clean
clean: ## remove build artifact
		rm -rf pwbook

.PHONY: build
build: ## build executable
		@go build ./cmd/pwbook/pwbook.go

.PHONY: test
test: ## run tests
		@go test -v -cover ./internal/...

.PHONY: test-e2e
test-e2e: ## run e2e tests
		@go get ./...
		@go install ./cmd/pwbook
		@go test -v ./e2e/...