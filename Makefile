VEIDEMANN_API_VERSION?=1.0.0-beta12
VEIDEMANN_API_REPOSITORY?=git@github.com:nlnwa/veidemann-api.git
PROTOC_VERSION=3.11.4

.PHONY: all
.PHONY: clean
.PHONY: distclean

all:	tools veidemann-api build

clean:
	rm -rf protobuf include veidemann-api commons config contentwriter dnsresolver frontier robotsevaluator veidemann_api browsercontroller eventhandler ooshandler controller report

distclean: clean
	rm -rf tools

tools:
	mkdir tools
	cd tools \
	&& wget -q https://github.com/google/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip \
	&& unzip protoc-${PROTOC_VERSION}-linux-x86_64.zip \
	&& rm protoc-${PROTOC_VERSION}-linux-x86_64.zip
	go get -u google.golang.org/grpc
	go get -u github.com/golang/protobuf/protoc-gen-go

build: veidemann-api
	find veidemann-api -name *.proto -exec 'tools/bin/protoc' '-Iveidemann-api/protobuf' '-Itools/include' '--go_out=plugins=grpc,paths=source_relative:.' '{}' ';'

veidemann-api: tools
	git clone --depth=1 --branch ${VEIDEMANN_API_VERSION} ${VEIDEMANN_API_REPOSITORY}
