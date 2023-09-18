MO_DOCKER_IMAGE_NAME_PREFIX ?= ghcr.io/0x6flab/mpesaoverlay
BUILD_DIR = build
SERVICES = cli overlay
DOCKERS = $(addprefix docker_,$(SERVICES))
DOCKERS_DEV = $(addprefix docker_dev_,$(SERVICES))
CGO_ENABLED ?= 0
GOARCH ?= amd64
VERSION ?= $(shell git describe --tags --abbrev=0 2>/dev/null || echo "0.1.0")
COMMIT ?= $(shell git rev-parse HEAD)
TIME ?= $(shell date +%F_%T)

define compile_service
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) GOARM=$(GOARM) \
	go build -mod=vendor -ldflags "-s -w \
	-X 'github.com/0x6flab/mpesaoverlay.BuildTime=$(TIME)' \
	-X 'github.com/0x6flab/mpesaoverlay.Version=$(VERSION)' \
	-X 'github.com/0x6flab/mpesaoverlay.Commit=$(COMMIT)'" \
	-o ${BUILD_DIR}/mpesa-$(1) cmd/$(1)/main.go
endef

define make_docker
	$(eval svc=$(subst docker_,,$(1)))

	docker build \
		--no-cache \
		--build-arg SVC=$(svc) \
		--build-arg GOARCH=$(GOARCH) \
		--build-arg GOARM=$(GOARM) \
		--build-arg GOOS=$(GOOS) \
		--build-arg CGO_ENABLED=$(CGO_ENABLED) \
		--build-arg VERSION=$(VERSION) \
		--build-arg COMMIT=$(COMMIT) \
		--build-arg TIME=$(TIME) \
		--tag=$(MO_DOCKER_IMAGE_NAME_PREFIX)/$(svc):latest \
		-f docker/Dockerfile .
endef

define make_docker_dev
	$(eval svc=$(subst docker_dev_,,$(1)))

	docker build \
		--no-cache \
		--build-arg SVC=$(svc) \
		--tag=$(MO_DOCKER_IMAGE_NAME_PREFIX)/$(svc):latest \
		-f docker/Dockerfile.dev ./build
endef

define docker_push
	for svc in $(SERVICES); do \
		docker push $(MO_DOCKER_IMAGE_NAME_PREFIX)/$$svc:$(VERSION); \
		docker push $(MO_DOCKER_IMAGE_NAME_PREFIX)/$$svc:latest; \
	done
endef

$(SERVICES):
	$(call compile_service,$@)

all: $(SERVICES)

.PHONY: all $(SERVICES) dockers dockers_dev

$(DOCKERS):
	$(call make_docker,$(@),$(GOARCH),$(GOARM),$(GOOS),$(CGO_ENABLED),$(VERSION),$(COMMIT),$(TIME))

$(DOCKERS_DEV):
	$(call make_docker_dev,$(@),$(GOARCH),$(GOARM),$(GOOS),$(CGO_ENABLED),$(VERSION),$(COMMIT),$(TIME))

dockers: $(DOCKERS)
dockers_dev: $(DOCKERS_DEV)

docker_push:
	$(call docker_push)

clean:
	rm -rf ${BUILD_DIR}

install:
	cp ${BUILD_DIR}/* $(GOBIN)

test:
	go test -mod=vendor -v -race -count 1 -tags test $(shell go list ./... | grep -v 'vendor\|cmd')

changelog:
	git log $(shell git describe --tags --abbrev=0)..HEAD --pretty=format:"- %s"

proto:
	# go install github.com/anjmao/go2proto@latest
	go2proto -f overlay/ -p pkg/request.go
	mv overlay/output.proto overlay/requests.proto
	sed -i 's,package proto;,package mpesaoverlay.overlay;\noption go_package = "./overlay";,g' overlay/requests.proto
	sed -i 's/uint8/uint32/g' overlay/requests.proto
	go2proto -f overlay/ -p pkg/response.go
	mv overlay/output.proto overlay/responses.proto
	sed -i 's,package proto;,package mpesaoverlay.overlay;\noption go_package = "./overlay";,g' overlay/responses.proto
	sed -i 's/uint8/uint32/g' overlay/responses.proto
	protoc -I. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative overlay/*.proto

run:
	docker-compose -f docker/docker-compose.yml --env-file docker/.env up -d

stop:
	docker-compose -f docker/docker-compose.yml --env-file docker/.env down
