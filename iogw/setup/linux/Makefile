WSK=https://github.com/apache/openwhisk-cli/releases/download/1.0.0/OpenWhisk_CLI-1.0.0-linux-amd64.tgz
ifeq ($(origin IOGW_VER), undefined)
IOGW_VER?=$(shell git branch --show-current)
IOGW_VER_PKG=0.0-$(IOGW_VER)
else
IOGW_VER_PKG=$(IOGW_VER)
endif

HERE:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
.PHONY: all
all: bin/iogw bin/wsk
	docker run --rm -v $(HERE):/mnt pagopa/builder-fpm-nsis bash /mnt/build.sh $(IOGW_VER_PKG)

bin/iogw:
	mkdir -p bin
	GOOS=linux GOARCH=amd64 go build \
	-o bin/iogw -ldflags "-X main.Version=$(IOGW_VER) -X main.DockerHubUser=$(DOCKER_USER)" ../..

bin/wsk:
	mkdir -p bin
	curl -sL $(WSK) | tar xzvf - -C bin wsk

.PHONY: clean
clean:
	rm -rfv bin deb rpm

