BUILD=9999.99.99
VER=1.0.0
ARTIFACT:=objs/$(notdir ${CURDIR})

ifeq ($(OS),Windows_NT)
DETECTED_OS := Windows
ARTIFACT:=$(ARTIFACT).exe
else
DETECTED_OS := $(shell sh -c 'uname 2>/dev/null || echo Unknown')
BUILD := $(shell date +%F_%T)
endif

LDFLAGS=-ldflags "-X main._os=$(DETECTED_OS) -X main._date=$(BUILD) -X main._ver=$(VER)"

.PHONY: all acceptor plugin echo

all: acceptor plugin echo
	find . -name *.so | rg -v '^./objs.*$$' | xargs -n 1 -i mv {} objs

echo: export GOPROXY=https://goproxy.cn
echo: export GO111MODULE=on
echo:
	go build $(LDFLAGS) -v -o $(ARTIFACT) main.go

acceptor:
	make -C acceptor/golib/httpserver

plugin:
	make -C plugin/golib/print