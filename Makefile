VEIDEMANN_API_VERSION?=1.0.0-beta18
VEIDEMANN_API_REPOSITORY?=https://github.com/nlnwa/veidemann-api
PROTOC_VERSION=3.13.0
DIRS=commons config contentwriter dnsresolver frontier robotsevaluator browsercontroller eventhandler ooshandler controller report

.PHONY: all
.PHONY: clean
.PHONY: distclean
.PHONY: build

all:	tools veidemann-api build

clean:
	rm -rf $(DIRS)

distclean: clean
	rm -rf tools
	rm -rf veidemann-api

veidemann-api:
	git clone --depth=1 --branch ${VEIDEMANN_API_VERSION} ${VEIDEMANN_API_REPOSITORY}

tools:
	mkdir tools \
	&& cd tools \
	&& wget -q https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip \
	&& unzip protoc-${PROTOC_VERSION}-linux-x86_64.zip \
	&& rm protoc-${PROTOC_VERSION}-linux-x86_64.zip \
	&& go get google.golang.org/protobuf/cmd/protoc-gen-go \
	google.golang.org/grpc/cmd/protoc-gen-go-grpc

$(DIRS): veidemann-api tools
	find veidemann-api -name *.proto -exec 'tools/bin/protoc' '-Iveidemann-api/protobuf' '-Itools/include' '--go_out=.' '--go_opt=paths=source_relative' '--go-grpc_out=.' '--go-grpc_opt=paths=source_relative' '{}' ';'

build: $(DIRS)
