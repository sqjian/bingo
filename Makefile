.PHONY: all test venv clean

all: venv
test:
	make -C example/echo
venv:
	make -f venv.mk
clean:
	$(clean)
