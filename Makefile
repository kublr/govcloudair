all: test

GOOS ?=
GOBINARY ?= go

SOURCE=$$($(GOBINARY) list ./... | grep -v /vendor/ | grep -v /gen/)
SOURCE_FILES?=$$(find . -name '*.go' | grep -v /vendor/ | grep -v /gen/)

deps:
	$(GOBINARY) get -u github.com/golang/dep/cmd/dep
	$(GOBINARY) get -u golang.org/x/lint/golint
.PHONY: deps

deps-install:
	$(ENVVAR) GOOS=$(GOOS) dep ensure -v
.PHONY: deps-install

deps-update:
	$(ENVVAR) GOOS=$(GOOS) dep ensure
.PHONY: deps-update

codestyle:
	@echo "==> Running codestyle checks"
	GOOS=$(GOOS) gofmt -l -e -d -s $(SOURCE_FILES)
	GOOS=$(GOOS) test -z "$(shell gofmt -l $(SOURCE_FILES))"

	GOOS=$(GOOS) golint $(SOURCE)

	GOOS=$(GOOS) $(GOBINARY) vet $(SOURCE)
.PHONY: codestyle

test: codestyle
	@echo "==> Running tests"
	GOOS=$(GOOS) $(GOBINARY) test -v $(SOURCE)
.PHONY: test
