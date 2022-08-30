# Setting SHELL to bash allows bash commands to be executed by recipes.
# Options are set to exit when a recipe line exits non-zero or a piped command fails.
SHELL = /usr/bin/env bash
.SHELLFLAGS = -o pipefail -ec

.PHONY: all
all: build

##@ General

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk commands is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

.PHONY: run
run: ## Run Spencer (only suitable for development).
	go run -race ./cmd/spencer

.PHONY: fmt
fmt: gofumpt ## Run gofumpt against code.
	$(GOFUMPT) -w .

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

.PHONY: lint
lint: golangci ## Run golangci-lint against code.
	$(GOLANGCILINT) run

##@ Build

.PHONY: build
build: fmt vet lint ## Compile Spencer
	go build -o $(LOCALBIN)/spencer ./cmd/spencer

##@ Installation

.PHONY: install
install: build ## Install Spencer.
	cp $(LOCALBIN)/spencer $(shell go env GOBIN)/

.PHONY: uninstall
uninstall: ## Uninstall Spencer.
	-rm $(shell go env GOBIN)/spencer

##@ Build Dependencies

## Location to install dependencies to
LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

## Tool Binaries
GOFUMPT ?= $(LOCALBIN)/gofumpt
GOLANGCILINT ?= $(LOCALBIN)/golangci-lint

## Tool Versions
GOFUMPT_VERSION ?= v0.3.1
GOLANGCILINT_VERSION ?= v1.49.0

.PHONY: gofumpt
gofumpt: $(GOFUMPT) ## Download gofumpt locally if necessary.
$(GOFUMPT): | $(LOCALBIN)
	GOBIN=$(LOCALBIN) go install mvdan.cc/gofumpt@$(GOFUMPT_VERSION)

.PHONY: golangci
golangci: | $(LOCALBIN) ## Download golangci-lint locally if necessary.
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(LOCALBIN) $(GOLANGCILINT_VERSION)

##@ Clean

.PHONY: clean
clean: uninstall ## Remove compiled binaries and build tools
	-rm -rf $(LOCALBIN)
