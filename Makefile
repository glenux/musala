
GOFILES=$(wildcard *.go)
NAME=musala-push

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

npm: 
	npm install


MJML_TEMPLATES=$(wildcard templates/*.mjml)
MJML_OUTPUT=$(patsubst %.mjml,%.mjml.html,$(MJML_TEMPLATES))

%.mjml.html: %.mjml
	npx mjml $< --config.minify > $@

templates: $(MJML_OUTPUT)

.PHONY: templates

