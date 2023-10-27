.PHONY: build clean help
BIN_FILE=project-layout

all: build

build:
	go build -ldflags "-w -s" -o ${BIN_FILE}

clean:
	rm -rf ${BIN_FILE}
	go clean -i .

help:
	@echo "make: compile packages and dependencies"
	@echo "make clean: remove object files and cached files"
