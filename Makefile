
GOFILES=$(wildcard *.go)
BINDIR=bin
NAME=musala
MJML_TEMPLATES=$(wildcard templates/*.mjml)
MJML_OUTPUT=$(patsubst %.mjml,%.mjml.html,$(MJML_TEMPLATES))

all: help

%.mjml.html: %.mjml
	npx mjml $< --config.minify > $@

.PHONY: build build-binaries build-templates
build: build-binaries build-templates  ## build executable

build-binaries: build-templates
	cd $(BINDIR) && go build ../...

build-templates: $(MJML_OUTPUT)

install: ## install binaries
	go install ./...

.PHONY: shellcheck
shellcheck: ## run shellcheck validation
	scripts/validate/shellcheck

.PHONY: help
help: ## print this help
	@awk 'BEGIN {FS = ":.*?## "} \
		/^[a-zA-Z_-]+:.*?## / \
		{sub("\\\\n",sprintf("\n%22c"," "), \
		$$2);printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' \
		$(MAKEFILE_LIST)

.PHONY: clean clean-templates clean-binaries
clean: clean-templates clean-binaries ## remove build artifacts

clean-templates:
	rm -f ./templates/*.mjml.html

clean-binaries:
	rm -rf $(BINDIR)/*

test: build 
	./test.sh

npm: 
	npm install



