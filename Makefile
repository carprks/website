.PHONY: all clean

export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

VERSION ?= latest
LDFLAGS := -X "main.Version=${VERSION}" -X "main.Build=$(shell git rev-parse --short=7 HEAD)"
TAGS ?=
SERVICENAME ?= website
DD := "docker"

.PHONY: all
all: build

.PHONY: build
build:
	go build -v -tags '$(TAGS)' -ldflags '-s -w $(LDFLAGS)' -o $(SERVICENAME)

.PHONY: clean
clean:
	go clean -i ./...
	rm -rf $(SERVICENAME)
	$(DD) rmi "carprks/$(SERVICENAME):$(VERSION)"
	$(DD) rmi "carprks/$(SERVICENAME):latest"

.PHONMY: osx
osx:
	GOOS=darwin go build -v -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $(SERVICENAME)

.PHONY: docker
docker:
	docker build -t "carprks/$(SERVICENAME):$(VERSION)" \
		--build-arg build=$(shell git rev-parse --short=7 HEAD) \
		--build-arg version=$(VERSION) \
		--build-arg SERVICE_NAME=$(SERVICENAME) \
		-f Dockerfile .