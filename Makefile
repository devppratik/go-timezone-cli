# Ensure go modules are enabled:
export GO111MODULE=on
export GOPROXY=https://proxy.golang.org

# Disable CGO so that we always generate static binaries:
export CGO_ENABLED=0

# Constants:
GOPATH := $(shell go env GOPATH)

.PHONY: build
build:
	go build -o strman

.PHONY: install
install:
	go build -o ${GOPATH}/bin/tmz
	mkdir -p ${GOPATH}/pkg/tmz/
	cp pkg/data/country.json ${GOPATH}/pkg/tmz/
	cp pkg/data/abbr.json ${GOPATH}/pkg/tmz/

.PHONY: mod
mod:
	go mod tidy

.PHONY: vendor
vendor:
	go mod vendor

.PHONY: test
test: mod vendor
	go test ./tests -v

.PHONY: lint
lint: getlint
	$(GOPATH)/bin/golangci-lint run

.PHONY: clean
clean:
	rm -rf \
		*-darwin-amd64 \
		*-linux-amd64 \
		*-windows-amd64 \
		*.sha256 \
		$(NULL)