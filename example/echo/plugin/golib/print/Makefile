ARTIFACT:=objs/$(notdir ${CURDIR}).so

export GOPROXY=https://goproxy.cn
export GO111MODULE=on

.PHONY: print
print:
ifeq ($(OS),Windows_NT)
	@echo "not support windows"
else
	go build -buildmode=plugin -o ${ARTIFACT} *.go
endif
