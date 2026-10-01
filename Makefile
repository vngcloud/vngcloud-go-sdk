.DEFAULT_GOAL := help

# Packages that are safe to run without credentials or network access.
# ./test/... is excluded on purpose: it holds live integration tests that call
# the real VNGCloud APIs and need credentials in test/env.yaml.
UNIT_PKGS = $(shell go list ./... | grep -v '/test$$' | grep -v '/test/')

# Pinned to the same minor version CI uses (golangci-lint-action version: v2.6).
GOLANGCI_LINT_VERSION ?= v2.6.1
LOCALBIN ?= $(CURDIR)/bin
GOLANGCI_LINT = $(LOCALBIN)/golangci-lint-$(GOLANGCI_LINT_VERSION)

# Lint only code that is new relative to this revision (old findings are ignored).
LINT_BASE ?= origin/main

##@ General

.PHONY: help
help: ## Show this help.
	@awk 'BEGIN {FS = ":.*##"; printf "Usage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-18s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) }' $(MAKEFILE_LIST)

##@ Verify

.PHONY: build
build: ## Compile all packages.
	go build ./...

.PHONY: vet
vet: ## Run go vet on all packages except ./test (integration tests).
	go vet $(UNIT_PKGS)

.PHONY: test
test: ## Run unit tests (excludes ./test/..., see test-integration).
	go test $(UNIT_PKGS)

.PHONY: test-integration
test-integration: ## Opt-in: live integration tests against real VNGCloud APIs (needs test/env.yaml). Never run by verify-fast.
	go test ./test/... -v

.PHONY: lint
lint: golangci-lint ## Run golangci-lint on new code only (vs LINT_BASE, default origin/main).
	$(GOLANGCI_LINT) run --new-from-rev=$(LINT_BASE) ./...

.PHONY: verify-fast
verify-fast: vet lint test ## Quick pre-finish check: vet + lint (new code) + unit tests, no network or credentials.

##@ Tools

.PHONY: golangci-lint
golangci-lint: $(GOLANGCI_LINT) ## Install the pinned golangci-lint into ./bin.

$(GOLANGCI_LINT):
	@mkdir -p $(LOCALBIN)
	GOBIN=$(LOCALBIN) go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)
	mv $(LOCALBIN)/golangci-lint $(GOLANGCI_LINT)
