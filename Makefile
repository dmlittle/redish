BIN_DIR  ?= ./bin
PKG_NAME ?= redish
LDFLAGS  ?= "-w -s"

COVERAGE_PROFILE ?= coverage.out

GOTOOLS := \
	golang.org/x/tools/cmd/cover \

default: build

.PHONY: build
build:
	@echo "---> Building"
	go build -o $(BIN_DIR)/$(PKG_NAME) -ldflags $(LDFLAGS) ./cmd/main.go

.PHONY: clean
clean:
	@echo "---> Cleaning"
	rm -rf $(BIN_DIR) ./vendor

.PHONY: enforce
enforce:
	@echo "---> Enforcing coverage"
	./scripts/coverage.sh $(COVERAGE_PROFILE)

.PHONY: html
html:
	@echo "---> Generating HTML coverage report"
	go tool cover -html $(COVERAGE_PROFILE)

.PHONY: lint
lint:
	@echo "---> Linting"
	$(BIN_DIR)/golangci-lint run


.PHONY: setup
setup:
	@echo "--> Installing development tools"
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(BIN_DIR) v1.16.0
	go get -u $(GOTOOLS)

.PHONY: test
test:
	@echo "---> Testing"
	ENVIRONMENT=test go test ./pkg/... -race -coverprofile $(COVERAGE_PROFILE)
