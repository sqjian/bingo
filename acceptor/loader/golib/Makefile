ARTIFACT:=${OUTPUT}/$(notdir ${CURDIR})

ifeq ($(OS),Windows_NT)
ARTIFACT:=${ARTIFACT}.exe
else
ARTIFACT:=${ARTIFACT}
endif

all:
	go test -c -o ${ARTIFACT}
