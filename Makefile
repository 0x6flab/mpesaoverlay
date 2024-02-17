MO_DOCKER_IMAGE_NAME_PREFIX ?= ghcr.io/0x6flab/mpesaoverlay
BUILD_DIR = build
SERVICES = cli grpc mqtt
DOCKERS = $(addprefix docker_,$(SERVICES))
DOCKERS_DEV = $(addprefix docker_dev_,$(SERVICES))
CGO_ENABLED ?= 0
GOARCH ?= amd64
VERSION ?= $(shell git describe --tags --abbrev=0 2>/dev/null || echo "v0.1.0")
COMMIT ?= $(shell git rev-parse HEAD)
TIME ?= $(shell date +%F_%T)

define compile_service
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) GOARM=$(GOARM) \
	go build -ldflags "-s -w \
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
		--tag=$(MO_DOCKER_IMAGE_NAME_PREFIX)/$(svc):$(VERSION) \
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

lint:
	golangci-lint run

test:
	go install github.com/vektra/mockery/v2@v2.42.0
	go generate ./...
	go test -v -race -covermode=atomic -coverprofile cover.out $(shell go list ./... | grep -v "example|cmd|cli|mocks")

version:
	standard-version
	goreleaser release --clean --release-notes CHANGELOG.md
	git push --follow-tags origin main

proto:
	go install github.com/anjmao/go2proto@latest
	go2proto -f grpc/ -p pkg/mpesa/request.go
	mv grpc/output.proto grpc/requests.proto
	sed -i 's,package proto;,package mpesaoverlay.grpc;\noption go_package = "./grpc";,g' grpc/requests.proto
	sed -i 's/uint8/uint32/g' grpc/requests.proto
	go2proto -f grpc/ -p pkg/mpesa/response.go
	mv grpc/output.proto grpc/responses.proto
	sed -i 's,package proto;,package mpesaoverlay.grpc;\noption go_package = "./grpc";,g' grpc/responses.proto
	sed -i 's/uint8/uint32/g' grpc/responses.proto
	protoc -I. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative grpc/*.proto

cert:
	openssl req -x509 -nodes -days 365 -newkey rsa:4096 -keyout docker/certs/cert.key -out docker/certs/cert.crt -subj "/C=KE/ST=Nairobi/L=Nairobi/O=0x6flab/CN=mpesaoverlay" && \
	chmod 644 docker/certs/cert.key && \
	chmod 600 docker/certs/cert.crt

run:
	docker-compose -f docker/docker-compose.yaml --env-file .env up

logs:
	docker-compose -f docker/docker-compose.yaml --env-file .env logs -f

stop:
	docker-compose -f docker/docker-compose.yaml --env-file .env down -v
