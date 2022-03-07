EXECUTABLE ?= helm-file-utils
OUTPUT_DIR ?= ./dist

.PHONY: build
build:
	CGO_ENABLED=0 go build -o $(OUTPUT_DIR)/$(EXECUTABLE) ./cmd

.PHONY: tidy
tidy:
	rm -fr go.sum
	go mod tidy -compat=1.17

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: tests
tests:
	go test --short -race -v ./...

.PHONY: clean
clean:
	rm -rf $(OUTPUT_DIR)

.PHONY: version
version:
	@echo $(VERSION)

.PHONY: addlicense
addlicense:
	go install github.com/google/addlicense@latest && \
		addlicense -c "True North <info@true-north.hr>" -l apache -v ./

