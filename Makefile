
GOFILES=$(wildcard *.go)
NAME=musala-mail

all: build

.PHONY: build
build:  ## build executable
	go build ./...

install: ## install binaries
	go install ./...

.PHONY: shellcheck
shellcheck: ## run shellcheck validation
	scripts/validate/shellcheck

.PHONY: help
help: ## print this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: clean
clean: ## remove build artifacts
	rm -rf ./_build/*

test: build 
	./test.sh
