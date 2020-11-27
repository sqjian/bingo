IMG:='registry.cn-hangzhou.aliyuncs.com/sqjian/venv:ubuntu20_04'
all: run

run: pull
	docker run \
			-it \
			--rm \
			--net=host \
			-v ${PWD}:/bingo:rw \
			-w /bingo \
			${IMG} \
			bash
pull:
	docker pull ${IMG}
