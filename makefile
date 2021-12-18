.PHONY: all build run gotool clean help

BINARY="baize"

all: gotoll build

build:
	set CGO_ENABLED=0  GOOS=linux  GOARCH=amd64 go build -o ${BINARY}

run:
	@go run ./main.go

gotool:
	go fmt ./
	go vet ./

clean:
	@if [ -f ${BINARY} ] ; then rm $${BINARY} ; fi

help:
	@echo "make - 格式化 Go 代码,并编译成二进制文件"