default: build

.PHONY: clean
clean: ## remove build artifact
		rm -rf ./build

.PHONY: build
build: ## build pwbook executable
		@go build -o ./build/pwbook ./main.go

.PHONY: install-deps
install-deps: ## install dependencies
		@go get ./...

.PHONY: install
install: install-deps ## install pwbook
		@go install .

.PHONY: test
test: ## run tests
		@go test -v -cover ./pwbook/...

.PHONY: test-e2e
test-e2e: install ## run e2e tests
		@go test -v ./e2e/...
