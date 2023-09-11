BUILD_DIR = build
SERVICES = cli
CGO_ENABLED ?= 0
GOARCH ?= amd64
VERSION ?= $(shell git describe --tags --abbrev=0 2>/dev/null || echo "0.1.0")

define compile_service
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) GOARM=$(GOARM) \
	go build -mod=vendor -ldflags "-s -w \
	-X 'github.com/0x6flab/mpesaoverlay.BuildTime=$(TIME)' \
	-X 'github.com/0x6flab/mpesaoverlay.Version=$(VERSION)' \
	-X 'github.com/0x6flab/mpesaoverlay.Commit=$(COMMIT)'" \
	-o ${BUILD_DIR}/mpesa-$(1) cmd/$(1)/main.go
endef

$(SERVICES):
	$(call compile_service,$@)

all: $(SERVICES)

.PHONY: all $(SERVICES)

clean:
	rm -rf ${BUILD_DIR}

install:
	cp ${BUILD_DIR}/* $(GOBIN)

test:
	go test -mod=vendor -v -race -count 1 -tags test $(shell go list ./... | grep -v 'vendor\|cmd')

changelog:
	git log $(shell git describe --tags --abbrev=0)..HEAD --pretty=format:"- %s"
