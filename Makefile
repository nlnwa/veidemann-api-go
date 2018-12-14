VEIDEMANN_API_VERSION=v1.0.0-beta1
PROTOC_VERSION=3.6.1
GRPC_WEB_VERSION=1.0.3

.PHONY: all
.PHONY: clean
.PHONY: distclean

all:	tools veidemann-api build

clean:
	rm -rf protobuf include veidemann-api config veidemann_api

distclean: clean
	rm -rf tools

tools:
	mkdir tools
	cd tools \
	&& wget -q https://github.com/google/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip \
	&& unzip protoc-${PROTOC_VERSION}-linux-x86_64.zip \
	&& rm protoc-${PROTOC_VERSION}-linux-x86_64.zip
	go get google.golang.org/grpc
	go get github.com/golang/protobuf/protoc-gen-go
	go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options
	go get google.golang.org/genproto/googleapis/api/annotations

build: veidemann-api
	find veidemann-api -name *.proto -exec 'tools/bin/protoc' '-Iveidemann-api/protobuf' '-Iveidemann-api/include' '-Itools/include' '--go_out=plugins=grpc,paths=source_relative:.' '{}' ';'

veidemann-api: tools
	git clone --depth=1 --branch ${VEIDEMANN_API_VERSION} git@github.com:nlnwa/veidemann-api.git
