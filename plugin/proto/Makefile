.PHONY: all env

export GO111MODULE=on
export GOPROXY=https://goproxy.cn

all:env
	protoc \
		-I . \
		--go_out=":./" \
		--validate_out="lang=go:./"  \
		msg.proto
env:
	go get github.com/golang/protobuf/protoc-gen-go@v1.3.4
	go get github.com/golang/protobuf/proto@v1.3.4
	go get google.golang.org/grpc
	go get github.com/envoyproxy/protoc-gen-validate
	go get github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc
venv:
	docker run \
		-it \
		--rm \
		--net=host \
		--name=venv \
		-w /lab \
		-v ${PWD}:/lab \
		registry.cn-hangzhou.aliyuncs.com/sqjian/venv:ubuntu20_04 bash
test:
	go test -v
