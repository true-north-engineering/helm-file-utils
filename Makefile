EXECUTABLE ?= helm-file-utils
GO ?= go
ARCHS ?= amd64 386
LDFLAGS ?= -X 'main.version=$(VERSION)'
GOFILES := $(shell find . -name "*.go" -type f)
OUTPUT_DIR ?= ./bin/release/

ifneq ($(TAG),)
	VERSION ?= $(TAG)
else
	VERSION ?= $(shell git describe --tags --always || git rev-parse --short HEAD)
endif

VERSION := $(VERSION)

.PHONY: build
build: clean $(EXECUTABLE)

.PHONY: $(EXECUTABLE)
$(EXECUTABLE): $(GOFILES)
	CGO_ENABLED=0 GOOS=linux $(GO) build -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -v -a -installsuffix cgo -o ./bin/file-utils ./cmd
	@echo "Done."

.PHONY: build_darwin
build_darwin:
	GOOS=darwin $(GO) build -v -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o ./bin/release/$@ ./cmd

.PHONY: unit_tests
unit_tests:
	$(GO) test --short -race -v ./...

.PHONY: clean
clean:
	rm -rf ${OUTPUT_DIR}

.PHONY: version
version:
	@echo $(VERSION)
