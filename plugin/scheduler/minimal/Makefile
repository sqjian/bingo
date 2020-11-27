ARTIFACT:=${OUTPUT}/$(notdir ${CURDIR})

ifeq ($(OS),Windows_NT)
ARTIFACT:=${ARTIFACT}.exe
else
ARTIFACT:=${ARTIFACT}
endif

.PHONY: test

test:
	go test -c -o ${ARTIFACT}
