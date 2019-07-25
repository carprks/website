.PHONY: all clean

export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

VERSION ?= latest
BUILD=$(shell git rev-parse --short=7 HEAD)
LDFLAGS := -ldflags "-w -s -X github.com/carprks/website/backend/website.Version=${VERSION} -X github.com/carprks/website/backend/website.Build=${BUILD}"
TAGS ?=
SERVICENAME ?= website
DD := "docker"

.PHONY: all
all: build

.PHONY: build
build:
	go build -v -tags '$(TAGS)' ${LDFLAGS} -o ${SERVICENAME} .

.PHONY: clean
clean:
	go clean -i ./...
	rm -rf $(SERVICENAME)
	$(DD) rmi "carprks/$(SERVICENAME):$(VERSION)"
	$(DD) rmi "carprks/$(SERVICENAME):latest"

.PHONY: osx
osx:
	GOOS=darwin go build -v -tags '$(TAGS)' ${LDFLAGS} -o ${SERVICENAME} .

.PHONY: docker
docker:
	docker build -t "carprks/$(SERVICENAME):$(VERSION)" \
		--build-arg build=${BUILD} \
		--build-arg version=$(VERSION) \
		--build-arg SERVICE_NAME=$(SERVICENAME) \
		-f Dockerfile .
